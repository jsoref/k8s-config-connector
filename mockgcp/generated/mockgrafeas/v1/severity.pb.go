// Copyright 2021 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mockgrafeas/v1/severity.proto

package grafeas

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Note provider assigned severity/impact ranking.
type Severity int32

const (
	// Unknown.
	Severity_SEVERITY_UNSPECIFIED Severity = 0
	// Minimal severity.
	Severity_MINIMAL Severity = 1
	// Low severity.
	Severity_LOW Severity = 2
	// Medium severity.
	Severity_MEDIUM Severity = 3
	// High severity.
	Severity_HIGH Severity = 4
	// Critical severity.
	Severity_CRITICAL Severity = 5
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "MINIMAL",
		2: "LOW",
		3: "MEDIUM",
		4: "HIGH",
		5: "CRITICAL",
	}
	Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"MINIMAL":              1,
		"LOW":                  2,
		"MEDIUM":               3,
		"HIGH":                 4,
		"CRITICAL":             5,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_mockgrafeas_v1_severity_proto_enumTypes[0].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_mockgrafeas_v1_severity_proto_enumTypes[0]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_mockgrafeas_v1_severity_proto_rawDescGZIP(), []int{0}
}

var File_mockgrafeas_v1_severity_proto protoreflect.FileDescriptor

var file_mockgrafeas_v1_severity_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2a,
	0x5e, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53,
	0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49, 0x4e, 0x49, 0x4d, 0x41, 0x4c,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10,
	0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x05, 0x42,
	0x59, 0x0a, 0x11, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f,
	0x63, 0x6b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x61,
	0x66, 0x65, 0x61, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mockgrafeas_v1_severity_proto_rawDescOnce sync.Once
	file_mockgrafeas_v1_severity_proto_rawDescData = file_mockgrafeas_v1_severity_proto_rawDesc
)

func file_mockgrafeas_v1_severity_proto_rawDescGZIP() []byte {
	file_mockgrafeas_v1_severity_proto_rawDescOnce.Do(func() {
		file_mockgrafeas_v1_severity_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgrafeas_v1_severity_proto_rawDescData)
	})
	return file_mockgrafeas_v1_severity_proto_rawDescData
}

var file_mockgrafeas_v1_severity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mockgrafeas_v1_severity_proto_goTypes = []interface{}{
	(Severity)(0), // 0: mockgrafeas.v1.Severity
}
var file_mockgrafeas_v1_severity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mockgrafeas_v1_severity_proto_init() }
func file_mockgrafeas_v1_severity_proto_init() {
	if File_mockgrafeas_v1_severity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mockgrafeas_v1_severity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgrafeas_v1_severity_proto_goTypes,
		DependencyIndexes: file_mockgrafeas_v1_severity_proto_depIdxs,
		EnumInfos:         file_mockgrafeas_v1_severity_proto_enumTypes,
	}.Build()
	File_mockgrafeas_v1_severity_proto = out.File
	file_mockgrafeas_v1_severity_proto_rawDesc = nil
	file_mockgrafeas_v1_severity_proto_goTypes = nil
	file_mockgrafeas_v1_severity_proto_depIdxs = nil
}
