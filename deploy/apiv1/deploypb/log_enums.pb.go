// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: google/cloud/deploy/v1/log_enums.proto

package deploypb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type indicates the type of the log entry and can be used as a filter.
type Type int32

const (
	// Type is unspecified.
	Type_TYPE_UNSPECIFIED Type = 0
	// A Pub/Sub notification failed to be sent.
	Type_TYPE_PUBSUB_NOTIFICATION_FAILURE Type = 1
	// Deprecated: This field is never used. Use release_render log type instead.
	//
	// Deprecated: Marked as deprecated in google/cloud/deploy/v1/log_enums.proto.
	Type_TYPE_RENDER_STATUES_CHANGE Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_PUBSUB_NOTIFICATION_FAILURE",
		2: "TYPE_RENDER_STATUES_CHANGE",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":                 0,
		"TYPE_PUBSUB_NOTIFICATION_FAILURE": 1,
		"TYPE_RENDER_STATUES_CHANGE":       2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_deploy_v1_log_enums_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_google_cloud_deploy_v1_log_enums_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_deploy_v1_log_enums_proto_rawDescGZIP(), []int{0}
}

var File_google_cloud_deploy_v1_log_enums_proto protoreflect.FileDescriptor

var file_google_cloud_deploy_v1_log_enums_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x76, 0x31,
	0x2a, 0x66, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24,
	0x0a, 0x20, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x55, 0x42, 0x53, 0x55, 0x42, 0x5f, 0x4e, 0x4f,
	0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x4e,
	0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x45, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x10, 0x02, 0x1a, 0x02, 0x08, 0x01, 0x42, 0x61, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x70, 0x62, 0x3b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_deploy_v1_log_enums_proto_rawDescOnce sync.Once
	file_google_cloud_deploy_v1_log_enums_proto_rawDescData = file_google_cloud_deploy_v1_log_enums_proto_rawDesc
)

func file_google_cloud_deploy_v1_log_enums_proto_rawDescGZIP() []byte {
	file_google_cloud_deploy_v1_log_enums_proto_rawDescOnce.Do(func() {
		file_google_cloud_deploy_v1_log_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_deploy_v1_log_enums_proto_rawDescData)
	})
	return file_google_cloud_deploy_v1_log_enums_proto_rawDescData
}

var file_google_cloud_deploy_v1_log_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_deploy_v1_log_enums_proto_goTypes = []interface{}{
	(Type)(0), // 0: google.cloud.deploy.v1.Type
}
var file_google_cloud_deploy_v1_log_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_deploy_v1_log_enums_proto_init() }
func file_google_cloud_deploy_v1_log_enums_proto_init() {
	if File_google_cloud_deploy_v1_log_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_deploy_v1_log_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_deploy_v1_log_enums_proto_goTypes,
		DependencyIndexes: file_google_cloud_deploy_v1_log_enums_proto_depIdxs,
		EnumInfos:         file_google_cloud_deploy_v1_log_enums_proto_enumTypes,
	}.Build()
	File_google_cloud_deploy_v1_log_enums_proto = out.File
	file_google_cloud_deploy_v1_log_enums_proto_rawDesc = nil
	file_google_cloud_deploy_v1_log_enums_proto_goTypes = nil
	file_google_cloud_deploy_v1_log_enums_proto_depIdxs = nil
}
