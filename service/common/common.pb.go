// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: common.proto

package common

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

type GetEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sender    string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient string `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	CpuUsed   string `protobuf:"bytes,4,opt,name=cpuUsed,proto3" json:"cpuUsed,omitempty"` //cpu使用情况
	MemUsed   string `protobuf:"bytes,5,opt,name=memUsed,proto3" json:"memUsed,omitempty"` //内存使用情况
	DiskUsed  string `protobuf:"bytes,6,opt,name=diskUsed,proto3" json:"diskUsed,omitempty"`
	Grade     string `protobuf:"bytes,7,opt,name=grade,proto3" json:"grade,omitempty"` //报警等级
}

func (x *GetEmailRequest) Reset() {
	*x = GetEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailRequest) ProtoMessage() {}

func (x *GetEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailRequest.ProtoReflect.Descriptor instead.
func (*GetEmailRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *GetEmailRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *GetEmailRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *GetEmailRequest) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *GetEmailRequest) GetCpuUsed() string {
	if x != nil {
		return x.CpuUsed
	}
	return ""
}

func (x *GetEmailRequest) GetMemUsed() string {
	if x != nil {
		return x.MemUsed
	}
	return ""
}

func (x *GetEmailRequest) GetDiskUsed() string {
	if x != nil {
		return x.DiskUsed
	}
	return ""
}

func (x *GetEmailRequest) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

type GetEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetEmailResponse) Reset() {
	*x = GetEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailResponse) ProtoMessage() {}

func (x *GetEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailResponse.ProtoReflect.Descriptor instead.
func (*GetEmailResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *GetEmailResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetEmailResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xcb, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x70, 0x75, 0x55, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x70, 0x75, 0x55, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x55,
	0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x55, 0x73,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x32, 0x4e, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_proto_goTypes = []interface{}{
	(*GetEmailRequest)(nil),  // 0: common.GetEmailRequest
	(*GetEmailResponse)(nil), // 1: common.GetEmailResponse
}
var file_common_proto_depIdxs = []int32{
	0, // 0: common.EmailService.SendEmail:input_type -> common.GetEmailRequest
	1, // 1: common.EmailService.SendEmail:output_type -> common.GetEmailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailRequest); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
