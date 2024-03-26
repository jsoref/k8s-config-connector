// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crdgeneration

import (
	"fmt"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration/crdboilerplate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/provider"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	//kubeschema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
)

const (
	TF2CRDLabel = "cnrm.cloud.google.com/tf2crd"
)

func GenerateTF2CRD(sm *corekccv1alpha1.ServiceMapping, resourceConfig *corekccv1alpha1.ResourceConfig) (*apiextensions.CustomResourceDefinition, error) {
	resource := resourceConfig.Name
	p := provider.Provider()
	r, ok := p.ResourcesMap[resource]
	if !ok {
		return nil, fmt.Errorf("unknown resource %v", resource)
	}
	s := r.Schema
	specFields := make(map[string]*schema.Schema)
	statusFields := make(map[string]*schema.Schema)
	for k, v := range s {
		if isConfigurableField(v) {
			specFields[k] = v
		} else {
			statusFields[k] = v
		}
	}
	openAPIV3Schema := crdboilerplate.GetOpenAPIV3SchemaSkeleton()
	specJSONSchema := tfObjectSchemaToJSONSchema(specFields)
	statusJSONSchema := tfObjectSchemaToJSONSchema(statusFields)
	removeIgnoredFields(resourceConfig, specJSONSchema, statusJSONSchema)
	removeOverwrittenFields(resourceConfig, specJSONSchema)
	markRequiredLocationalFieldsRequired(resourceConfig, specJSONSchema)
	addResourceIDFieldIfSupported(resourceConfig, specJSONSchema)
	handleHierarchicalReferences(resourceConfig, specJSONSchema)

	if len(specJSONSchema.Properties) > 0 {
		openAPIV3Schema.Properties["spec"] = *specJSONSchema
		if len(specJSONSchema.Required) > 0 {
			openAPIV3Schema.Required = slice.IncludeString(openAPIV3Schema.Required, "spec")
		}
	}

	//<<<<<<< HEAD
	statusJSONSchema, err := k8s.RenameStatusFieldsWithReservedNamesIfResourceNotExcluded(resource, statusJSONSchema)
	if err != nil {
		return nil, fmt.Errorf("error renaming status fields with reserved names for %#v: %w", statusJSONSchema, err)
	}
	//=======
	//var err error
	//if k8s.OutputOnlyFieldsAreUnderObservedState(kubeschema.GroupVersionKind{
	//	Kind:    resourceConfig.Kind,
	//	Version: sm.GetVersionFor(resourceConfig),
	//	Group:   sm.Name,
	//}) {
	//	moveComputedFieldsToObservedState(statusOrObservedStateJSONSchema)
	//} else {
	//	statusOrObservedStateJSONSchema, err = k8s.RenameStatusFieldsWithReservedNamesIfResourceNotExcluded(resource, statusOrObservedStateJSONSchema)
	//	if err != nil {
	//		return nil, fmt.Errorf("error renaming status fields with reserved names for %#v: %w", statusOrObservedStateJSONSchema, err)
	//	}
	//}
	populateObservedState(resourceConfig, specJSONSchema, statusJSONSchema)
	for k, v := range statusJSONSchema.Properties {
		openAPIV3Schema.Properties["status"].Properties[k] = v
	}

	group := strings.ToLower(sm.Spec.Name) + "." + APIDomain

	kind := text.SnakeCaseToUpperCamelCase(resource)
	if resourceConfig != nil && resourceConfig.Kind != "" {
		kind = resourceConfig.Kind
	}
	versions := []string{sm.GetVersionFor(resourceConfig)}
	if resourceConfig.V1alpha1ToV1beta1 != nil && *resourceConfig.V1alpha1ToV1beta1 {
		versions = append(versions, k8s.KCCAPIVersionV1Alpha1)
	}
	storageVersion := ""
	if resourceConfig.StorageVersion != nil {
		storageVersion = *resourceConfig.StorageVersion
	}
	crd := GetCustomResourceDefinition(kind, group, versions, storageVersion, openAPIV3Schema, TF2CRDLabel)
	if resourceConfig.AutoGenerated {
		// All the auto-generated Terraform-based CRDs are in Alpha stability.
		crd.ObjectMeta.Labels[k8s.KCCStabilityLabel] = k8s.StabilityLevelAlpha
	} else if allVersionsAreAlpha(versions) {
		// If the CRD versions are all alpha, treat the stability level as alpha also.
		crd.ObjectMeta.Labels[k8s.KCCStabilityLabel] = k8s.StabilityLevelAlpha
	} else {
		crd.ObjectMeta.Labels[k8s.KCCStabilityLabel] = k8s.StabilityLevelStable
	}
	return crd, nil
}

// allVersionsAreAlpha returns true if all the CRD versions are alpha (v1alpha1, v1alpha2, etc)
func allVersionsAreAlpha(versions []string) bool {
	for _, version := range versions {
		if strings.HasPrefix(version, "v1alpha") {
			continue
		}
		if strings.HasPrefix(version, "v1beta") {
			return false
		}
		if version == "v1" {
			return false
		}
		klog.Fatalf("unhandled version %q", version)
	}
	return true
}

func tfObjectSchemaToJSONSchema(s map[string]*schema.Schema) *apiextensions.JSONSchemaProps {
	jsonSchema := apiextensions.JSONSchemaProps{
		Type:       "object",
		Properties: make(map[string]apiextensions.JSONSchemaProps),
	}
	for k, v := range s {
		key := text.SnakeCaseToLowerCamelCase(k)
		if v.Required {
			jsonSchema.Required = slice.IncludeString(jsonSchema.Required, key)
		}
		js := *tfSchemaToJSONSchema(v)
		description := js.Description
		if description != "" {
			description = ensureEndsInPeriod(description)
		}
		if v.ForceNew {
			description = strings.TrimSpace("Immutable. " + description)
		}
		if v.Deprecated != "" {
			deprecationMsg := ensureEndsInPeriod(fmt.Sprintf("DEPRECATED. %v", v.Deprecated))
			description = cleanupDeprecatedFieldDescription(strings.TrimSpace(fmt.Sprintf("%v %v", deprecationMsg, description)))
		}
		// if the description contains "terraform", ignore the description field
		for _, word := range []string{"terraform", "Terraform"} {
			if !strings.Contains(description, word) {
				continue
			}
			if v.Deprecated != "" {
				panic(fmt.Errorf("about to strip field description since it contains "+
					"the word '%v', but we likely must avoid stripping the "+
					"description entirely since it contains a deprecation message "+
					"that likely should stay included. Suggest changing field's "+
					"description and/or deprecation message to drop the word '%v'. "+
					"Description:\n%v",
					word, word, description))
			}
			description = ""
		}
		js.Description = description
		jsonSchema.Properties[key] = js
	}
	return &jsonSchema
}

func ensureEndsInPeriod(str string) string {
	if !strings.HasSuffix(str, ".") {
		return str + "."
	}
	return str
}

func tfSchemaToJSONSchema(tfSchema *schema.Schema) *apiextensions.JSONSchemaProps {
	jsonSchema := apiextensions.JSONSchemaProps{}
	switch tfSchema.Type {
	case schema.TypeBool:
		jsonSchema.Type = "boolean"
	case schema.TypeFloat:
		jsonSchema.Type = "number"
	case schema.TypeInt:
		jsonSchema.Type = "integer"
	case schema.TypeSet:
		// schema.TypeSet is just like schema.TypeList; the validation for no duplicates happens elsewhere.
		fallthrough
	case schema.TypeList:
		jsonSchema.Type = "array"
		switch v := tfSchema.Elem.(type) {
		case *schema.Resource:
			// MaxItems == 1 actually signifies that this is a nested object, and not actually a
			// list, due to limitations of the TF schema type.
			if tfSchema.MaxItems == 1 {
				jsonSchema = *tfObjectSchemaToJSONSchema(v.Schema)
				break
			}
			jsonSchema.Items = &apiextensions.JSONSchemaPropsOrArray{
				Schema: tfObjectSchemaToJSONSchema(v.Schema),
			}
		case *schema.Schema:
			// List of primitives
			jsonSchema.Items = &apiextensions.JSONSchemaPropsOrArray{
				Schema: tfSchemaToJSONSchema(v),
			}
		default:
			panic("could not parse elem attribute of TF list/set schema")
		}
	case schema.TypeMap:
		// schema.TypeMap is only used for basic map[primitive]primitive resources; maps with schemas for the keys
		// are handled by schema.TypeList with MaxItems == 1
		jsonSchema.Type = "object"
		if mapSchema, ok := tfSchema.Elem.(*schema.Schema); ok {
			jsonSchema.AdditionalProperties = &apiextensions.JSONSchemaPropsOrBool{
				Schema: tfSchemaToJSONSchema(mapSchema),
			}
		}
	case schema.TypeString:
		if tfSchema.Sensitive && isConfigurableField(tfSchema) {
			jsonSchema = crdboilerplate.GetSensitiveFieldSchemaBoilerplate()
		} else {
			jsonSchema.Type = "string"
		}
	case schema.TypeInvalid:
		panic(fmt.Errorf("schema type is invalid"))
	default:
		panic(fmt.Errorf("unknown schema type %v", tfSchema.Type))
	}
	jsonSchema.Description = tfSchema.Description
	return &jsonSchema
}

func removeOverwrittenFields(rc *corekccv1alpha1.ResourceConfig, s *apiextensions.JSONSchemaProps) {
	if rc.MetadataMapping.Name != "" {
		removeField(rc.MetadataMapping.Name, s)
	}
	if rc.MetadataMapping.Labels != "" {
		removeField(rc.MetadataMapping.Labels, s)
	}
	for _, refConfig := range rc.ResourceReferences {
		handleResourceReference(refConfig, s)
	}
	for _, d := range rc.Directives {
		removeField(d, s)
	}
	if !krmtotf.SupportsHierarchicalReferences(rc) {
		// TODO(b/193177782): Delete this if-block once all resources support
		// hierarchical references.
		for _, c := range rc.Containers {
			removeField(c.TFField, s)
		}
	}
}

func removeIgnoredFields(rc *corekccv1alpha1.ResourceConfig, specJSONSchema, statusJSONSchema *apiextensions.JSONSchemaProps) {
	for _, f := range rc.IgnoredFields {
		removedInSpec := removeFieldIfExist(f, specJSONSchema)
		removedInStatus := removeFieldIfExist(f, statusJSONSchema)
		if removedInSpec && removedInStatus {
			panic(fmt.Errorf("found ignored field %s in both spec and status JSON schema for resource %s", f, rc.Name))
		}
		if !removedInSpec && !removedInStatus {
			panic(fmt.Errorf("cannot find ignored field %s in either spec or status JSON schema for resource %s", f, rc.Name))
		}
	}
}

// removeFieldIfExist attempts to remove a field from the provided json schema.
// The function is no-op if a field is not found.
// Returns true if a field is found and removed, returns false if a field is not found.
func removeFieldIfExist(f string, s *apiextensions.JSONSchemaProps) bool {
	if !fieldExists(f, s) {
		return false
	}
	removeField(f, s)
	return true
}

func markRequiredLocationalFieldsRequired(rc *corekccv1alpha1.ResourceConfig, s *apiextensions.JSONSchemaProps) {
	if rc.IDTemplate == "" {
		return
	}

	locationalFields := []string{"region", "zone", "location"}
	for _, field := range locationalFields {
		// It is assumed that locational fields (region, zone, location) would
		// always be at the base level.
		if _, ok := s.Properties[field]; !ok {
			continue
		}
		if !strings.Contains(rc.IDTemplate, fmt.Sprintf("{{%v}}", field)) {
			continue
		}
		s.Required = slice.IncludeString(s.Required, field)
	}
}

func handleResourceReference(refConfig corekccv1alpha1.ReferenceConfig, s *apiextensions.JSONSchemaProps) {
	*s = populateReference(strings.Split(refConfig.TFField, "."), refConfig, s)
}

func populateReference(path []string, refConfig corekccv1alpha1.ReferenceConfig, s *apiextensions.JSONSchemaProps) apiextensions.JSONSchemaProps {
	field := text.SnakeCaseToLowerCamelCase(path[0])
	if len(path) > 1 {
		subSchema := s.Properties[field]
		switch subSchema.Type {
		case "array":
			itemSchema := populateReference(path[1:], refConfig, subSchema.Items.Schema)
			subSchema.Items.Schema = &itemSchema
			return *s
		case "object":
			objSchema := populateReference(path[1:], refConfig, &subSchema)
			s.Properties[field] = objSchema
			return *s
		default:
			panic(fmt.Errorf("error parsing reference %v: cannot iterate into type that is not object or array of objects", path))
		}
	}

	// Base case; we have found the field representing the reference
	isList := s.Properties[field].Type == "array"
	var refSchema *apiextensions.JSONSchemaProps
	key := field
	if len(refConfig.Types) == 0 {
		if refConfig.Key != "" {
			key = refConfig.Key
			delete(s.Properties, field)
			if slice.StringSliceContains(s.Required, field) {
				s.Required = slice.RemoveStringFromStringSlice(s.Required, field)
				s.Required = slice.IncludeString(s.Required, key)
			}
		}
		refSchema = GetResourceReferenceSchemaFromTypeConfig(refConfig.TypeConfig)
	} else {
		refSchema = &apiextensions.JSONSchemaProps{
			Type:       "object",
			Properties: map[string]apiextensions.JSONSchemaProps{},
		}
		for _, v := range refConfig.Types {
			if v.JSONSchemaType == "" {
				refSchema.Properties[v.Key] = *GetResourceReferenceSchemaFromTypeConfig(v)
			} else {
				refSchema.Properties[v.Key] = apiextensions.JSONSchemaProps{
					Type: v.JSONSchemaType,
				}
			}
		}
	}

	refSchema.Description = refConfig.Description

	if isList {
		s.Properties[key] = apiextensions.JSONSchemaProps{
			Type: "array",
			Items: &apiextensions.JSONSchemaPropsOrArray{
				Schema: refSchema,
			},
		}
	} else {
		s.Properties[key] = *refSchema
	}

	return *s
}

func getDescriptionForExternalRef(typeConfig corekccv1alpha1.TypeConfig) string {
	targetField := typeConfig.TargetField
	if targetField == "" {
		targetField = "name"
	}
	targetField = text.SnakeCaseToLowerCamelCase(targetField)
	article := text.IndefiniteArticleFor(typeConfig.GVK.Kind)
	if typeConfig.ValueTemplate != "" {
		return fmt.Sprintf(
			"Allowed value: string of the format `%v`, where {{value}} is the `%v` field of %v `%v` resource.",
			typeConfig.ValueTemplate, targetField, article, typeConfig.GVK.Kind,
		)
	}
	return fmt.Sprintf("Allowed value: The `%v` field of %v `%v` resource.", targetField, article, typeConfig.GVK.Kind)
}

func GetResourceReferenceSchemaFromTypeConfig(typeConfig corekccv1alpha1.TypeConfig) *apiextensions.JSONSchemaProps {
	description := getDescriptionForExternalRef(typeConfig)
	return crdboilerplate.GetResourceReferenceSchemaBoilerplate(description)
}

func fieldExists(f string, s *apiextensions.JSONSchemaProps) bool {
	path := strings.Split(f, ".")
	return nestedFieldExists(path, s)
}

func nestedFieldExists(path []string, s *apiextensions.JSONSchemaProps) bool {
	if len(path) == 0 {
		panic("unexpected empty field path")
	}
	// check current level
	field := text.SnakeCaseToLowerCamelCase(path[0])
	subSchema, exists := s.Properties[field]
	if len(path) == 1 {
		return exists
	}
	// go to next level
	switch subSchema.Type {
	case "array":
		return nestedFieldExists(path[1:], subSchema.Items.Schema)
	case "object":
		return nestedFieldExists(path[1:], &subSchema)
	default:
		return false
	}
}

func removeField(tfField string, s *apiextensions.JSONSchemaProps) {
	*s = removeNestedField(strings.Split(tfField, "."), *s)
}

func removeNestedField(path []string, s apiextensions.JSONSchemaProps) apiextensions.JSONSchemaProps {
	field := text.SnakeCaseToLowerCamelCase(path[0])
	if len(path) > 1 {
		subSchema := s.Properties[field]
		switch subSchema.Type {
		case "array":
			itemSchema := removeNestedField(path[1:], *subSchema.Items.Schema)
			subSchema.Items.Schema = &itemSchema
		case "object":
			subSchema = removeNestedField(path[1:], subSchema)
		default:
			panic(fmt.Errorf("error parsing field %v: cannot iterate into type that is not object or array of objects", path))
		}
		s.Properties[field] = subSchema
		return s
	}
	delete(s.Properties, field)
	s.Required = slice.RemoveStringFromStringSlice(s.Required, field)
	return s
}

func isConfigurableField(tfSchema *schema.Schema) bool {
	return tfSchema.Required || tfSchema.Optional
}

func addResourceIDFieldIfSupported(rc *corekccv1alpha1.ResourceConfig, spec *apiextensions.JSONSchemaProps) {
	if !krmtotf.SupportsResourceIDField(rc) {
		return
	}

	spec.Properties[k8s.ResourceIDFieldName] = apiextensions.JSONSchemaProps{
		Type:        "string",
		Description: generateResourceIDFieldDescription(rc),
	}
}

func generateResourceIDFieldDescription(rc *corekccv1alpha1.ResourceConfig) string {
	targetFieldCamelCase := text.SnakeCaseToLowerCamelCase(rc.ResourceID.TargetField)
	isServerGeneratedResourceID := krmtotf.IsResourceIDFieldServerGenerated(rc)
	return GenerateResourceIDFieldDescription(targetFieldCamelCase, isServerGeneratedResourceID)
}

func handleHierarchicalReferences(rc *corekccv1alpha1.ResourceConfig, spec *apiextensions.JSONSchemaProps) {
	if len(rc.Containers) > 0 {
		// If resource supports resource-level container annotations, mark
		// hierarchical references optional since users can use the annotations
		// to configure the references.
		*spec = *MarkHierarchicalReferencesOptionalButMutuallyExclusive(spec, rc.HierarchicalReferences)
	} else {
		*spec = *MarkHierarchicalReferencesRequiredButMutuallyExclusive(spec, rc.HierarchicalReferences)
	}
}

func cleanupDeprecatedFieldDescription(description string) string {
	return strings.ReplaceAll(description,
		"is deprecated and will be removed in a future major release",
		"is deprecated")
}

func populateObservedState(rc *corekccv1alpha1.ResourceConfig, spec *apiextensions.JSONSchemaProps, status *apiextensions.JSONSchemaProps) {
	if rc.ObservedFields == nil || len(*rc.ObservedFields) == 0 {
		return
	}
	observedStateJSONSchema := apiextensions.JSONSchemaProps{
		Type:        "object",
		Description: "The observed state of the underlying GCP resource.",
		Properties:  make(map[string]apiextensions.JSONSchemaProps),
	}
	for _, path := range *rc.ObservedFields {
		populateObservedField(strings.Split(path, "."), spec, &observedStateJSONSchema)
	}
	status.Properties["observedState"] = observedStateJSONSchema
}

func populateObservedField(observedFieldPath []string, sourceSchema *apiextensions.JSONSchemaProps, observedFieldParent *apiextensions.JSONSchemaProps) apiextensions.JSONSchemaProps {
	field := text.SnakeCaseToLowerCamelCase(observedFieldPath[0])
	if len(observedFieldPath) > 1 {
		subSchema := sourceSchema.Properties[field]
		switch subSchema.Type {
		case "array":
			// TODO(b/312581557): Support the use case when the observed field is a subfield under an array.
			panic(fmt.Errorf("observed fields under an array is not supported"))
		case "object":
			objSchema := apiextensions.JSONSchemaProps{
				Type:        "object",
				Description: subSchema.Description,
				Properties:  make(map[string]apiextensions.JSONSchemaProps),
			}
			observedFieldParent.Properties[field] = populateObservedField(observedFieldPath[1:], &subSchema, &objSchema)
			return *observedFieldParent
		default:
			panic(fmt.Errorf("error parsing observed field %v: cannot iterate into type that is not object or array of objects", observedFieldPath))
		}
	}

	subSchema := sourceSchema.Properties[field]
	switch subSchema.Type {
	case "array":
		// TODO(b/312581569): Support the use case when the observed field is an array.
		panic(fmt.Errorf("observed fields of an array is not supported"))
	case "object":
		// TODO(b/312581569): Support the use case when the observed field is an object.
		panic(fmt.Errorf("observed fields of an object is not supported"))
	default:
		observedFieldParent.Properties[field] = *subSchema.DeepCopy()
	}

	return *observedFieldParent
}
