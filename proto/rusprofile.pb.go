// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/rusprofile.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INN         string `protobuf:"bytes,1,opt,name=INN,proto3" json:"INN,omitempty"`
	KPP         string `protobuf:"bytes,2,opt,name=KPP,proto3" json:"KPP,omitempty"`
	CompanyName string `protobuf:"bytes,3,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	Director    string `protobuf:"bytes,4,opt,name=Director,proto3" json:"Director,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rusprofile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rusprofile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_rusprofile_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetINN() string {
	if x != nil {
		return x.INN
	}
	return ""
}

func (x *Response) GetKPP() string {
	if x != nil {
		return x.KPP
	}
	return ""
}

func (x *Response) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *Response) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INN string `protobuf:"bytes,1,opt,name=INN,proto3" json:"INN,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rusprofile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rusprofile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_rusprofile_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetINN() string {
	if x != nil {
		return x.INN
	}
	return ""
}

var File_proto_rusprofile_proto protoreflect.FileDescriptor

var file_proto_rusprofile_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x4e, 0x4e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x4e, 0x4e, 0x12, 0x10, 0x0a, 0x03, 0x4b,
	0x50, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x50, 0x50, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x1b, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x4e, 0x4e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x4e, 0x4e, 0x32, 0x4f, 0x0a, 0x11, 0x52, 0x75, 0x73, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x3a, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f,
	0x49, 0x4e, 0x4e, 0x2f, 0x7b, 0x49, 0x4e, 0x4e, 0x7d, 0x42, 0x12, 0x5a, 0x10, 0x52, 0x75, 0x73,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rusprofile_proto_rawDescOnce sync.Once
	file_proto_rusprofile_proto_rawDescData = file_proto_rusprofile_proto_rawDesc
)

func file_proto_rusprofile_proto_rawDescGZIP() []byte {
	file_proto_rusprofile_proto_rawDescOnce.Do(func() {
		file_proto_rusprofile_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rusprofile_proto_rawDescData)
	})
	return file_proto_rusprofile_proto_rawDescData
}

var file_proto_rusprofile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_rusprofile_proto_goTypes = []interface{}{
	(*Response)(nil), // 0: proto.Response
	(*Request)(nil),  // 1: proto.Request
}
var file_proto_rusprofile_proto_depIdxs = []int32{
	1, // 0: proto.RusprofileWrapper.Get:input_type -> proto.Request
	0, // 1: proto.RusprofileWrapper.Get:output_type -> proto.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_rusprofile_proto_init() }
func file_proto_rusprofile_proto_init() {
	if File_proto_rusprofile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rusprofile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_rusprofile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_proto_rusprofile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rusprofile_proto_goTypes,
		DependencyIndexes: file_proto_rusprofile_proto_depIdxs,
		MessageInfos:      file_proto_rusprofile_proto_msgTypes,
	}.Build()
	File_proto_rusprofile_proto = out.File
	file_proto_rusprofile_proto_rawDesc = nil
	file_proto_rusprofile_proto_goTypes = nil
	file_proto_rusprofile_proto_depIdxs = nil
}