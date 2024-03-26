// Copyright 2024 Google LLC
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

package k8s

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestSupportsStateIntoSpecMerge(t *testing.T) {
	tests := []struct {
		name           string
		gvk            schema.GroupVersionKind
		expectedResult bool
	}{
		{
			name: "ComputeInstance should support 'state-into-spec: merge'",
			gvk: schema.GroupVersionKind{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "ComputeInstance",
			},
			expectedResult: true,
		},
		{
			name: "AccessContextManagerServicePerimeterResource should not " +
				"support 'state-into-spec: merge'",
			gvk: schema.GroupVersionKind{
				Group:   "accesscontextmanager.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "AccessContextManagerServicePerimeterResource",
			},
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actualResult := supportsStateIntoSpecMerge(tc.gvk)
			if actualResult != tc.expectedResult {
				t.Fatalf("got %v, want %v", actualResult, tc.expectedResult)
			}
		})
	}
}

func TestOutputOnlyFieldsAreUnderObservedState(t *testing.T) {
	crds, err := crdloader.LoadCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}
	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			gvk := schema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: version.Name,
				Kind:    crd.Spec.Names.Kind,
			}
			t.Run(fmt.Sprintf("%s/%s/%s", gvk.Group, gvk.Version, gvk.Kind), func(t *testing.T) {
				openAPISchema := version.Schema.OpenAPIV3Schema
				prop := findOpenAPIProperty(openAPISchema, "status", "observedState")
				hasObservedState := prop != nil
				got := OutputOnlyFieldsAreUnderObservedState(gvk)

				if hasObservedState != got {
					t.Errorf("status.observedState=%v, but OutputOnlyFieldsAreUnderObservedState=%v", prop, got)
				}
			})
		}
	}
}

func findOpenAPIProperty(schema *apiextensionsv1.JSONSchemaProps, path ...string) *apiextensionsv1.JSONSchemaProps {
	pos := schema
	for _, k := range path {
		p, found := pos.Properties[k]
		if !found {
			return nil
		}
		pos = &p
	}
	return pos
}
