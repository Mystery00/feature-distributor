// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: common/protobuf/project.proto

package pb

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

type CheckProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerKey *string `protobuf:"bytes,1,opt,name=serverKey,proto3,oneof" json:"serverKey,omitempty"`
	ClientKey *string `protobuf:"bytes,2,opt,name=clientKey,proto3,oneof" json:"clientKey,omitempty"`
}

func (x *CheckProjectRequest) Reset() {
	*x = CheckProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_protobuf_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckProjectRequest) ProtoMessage() {}

func (x *CheckProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_protobuf_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckProjectRequest.ProtoReflect.Descriptor instead.
func (*CheckProjectRequest) Descriptor() ([]byte, []int) {
	return file_common_protobuf_project_proto_rawDescGZIP(), []int{0}
}

func (x *CheckProjectRequest) GetServerKey() string {
	if x != nil && x.ServerKey != nil {
		return *x.ServerKey
	}
	return ""
}

func (x *CheckProjectRequest) GetClientKey() string {
	if x != nil && x.ClientKey != nil {
		return *x.ClientKey
	}
	return ""
}

type CheckProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsProjectExist bool    `protobuf:"varint,1,opt,name=isProjectExist,proto3" json:"isProjectExist,omitempty"`
	ProjectId      *int64  `protobuf:"varint,2,opt,name=projectId,proto3,oneof" json:"projectId,omitempty"`
	ProjectKey     *string `protobuf:"bytes,3,opt,name=projectKey,proto3,oneof" json:"projectKey,omitempty"`
	ProjectName    *string `protobuf:"bytes,4,opt,name=projectName,proto3,oneof" json:"projectName,omitempty"`
}

func (x *CheckProjectResponse) Reset() {
	*x = CheckProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_protobuf_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckProjectResponse) ProtoMessage() {}

func (x *CheckProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_protobuf_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckProjectResponse.ProtoReflect.Descriptor instead.
func (*CheckProjectResponse) Descriptor() ([]byte, []int) {
	return file_common_protobuf_project_proto_rawDescGZIP(), []int{1}
}

func (x *CheckProjectResponse) GetIsProjectExist() bool {
	if x != nil {
		return x.IsProjectExist
	}
	return false
}

func (x *CheckProjectResponse) GetProjectId() int64 {
	if x != nil && x.ProjectId != nil {
		return *x.ProjectId
	}
	return 0
}

func (x *CheckProjectResponse) GetProjectKey() string {
	if x != nil && x.ProjectKey != nil {
		return *x.ProjectKey
	}
	return ""
}

func (x *CheckProjectResponse) GetProjectName() string {
	if x != nil && x.ProjectName != nil {
		return *x.ProjectName
	}
	return ""
}

var File_common_protobuf_project_proto protoreflect.FileDescriptor

var file_common_protobuf_project_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x77, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x22, 0xda, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x4a, 0x0a, 0x0b, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_protobuf_project_proto_rawDescOnce sync.Once
	file_common_protobuf_project_proto_rawDescData = file_common_protobuf_project_proto_rawDesc
)

func file_common_protobuf_project_proto_rawDescGZIP() []byte {
	file_common_protobuf_project_proto_rawDescOnce.Do(func() {
		file_common_protobuf_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_protobuf_project_proto_rawDescData)
	})
	return file_common_protobuf_project_proto_rawDescData
}

var file_common_protobuf_project_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_protobuf_project_proto_goTypes = []interface{}{
	(*CheckProjectRequest)(nil),  // 0: CheckProjectRequest
	(*CheckProjectResponse)(nil), // 1: CheckProjectResponse
}
var file_common_protobuf_project_proto_depIdxs = []int32{
	0, // 0: CoreService.CheckProject:input_type -> CheckProjectRequest
	1, // 1: CoreService.CheckProject:output_type -> CheckProjectResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_protobuf_project_proto_init() }
func file_common_protobuf_project_proto_init() {
	if File_common_protobuf_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_protobuf_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckProjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_protobuf_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckProjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_common_protobuf_project_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_common_protobuf_project_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_protobuf_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_protobuf_project_proto_goTypes,
		DependencyIndexes: file_common_protobuf_project_proto_depIdxs,
		MessageInfos:      file_common_protobuf_project_proto_msgTypes,
	}.Build()
	File_common_protobuf_project_proto = out.File
	file_common_protobuf_project_proto_rawDesc = nil
	file_common_protobuf_project_proto_goTypes = nil
	file_common_protobuf_project_proto_depIdxs = nil
}
