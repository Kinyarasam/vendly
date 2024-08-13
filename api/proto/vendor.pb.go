// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: api/proto/vendor.proto

package vendor

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

type RegisterVendorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RegisterVendorRequest) Reset() {
	*x = RegisterVendorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_vendor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterVendorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterVendorRequest) ProtoMessage() {}

func (x *RegisterVendorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_vendor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterVendorRequest.ProtoReflect.Descriptor instead.
func (*RegisterVendorRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_vendor_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterVendorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterVendorRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RegisterVendorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	VendorId string `protobuf:"bytes,2,opt,name=vendorId,proto3" json:"vendorId,omitempty"`
}

func (x *RegisterVendorResponse) Reset() {
	*x = RegisterVendorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_vendor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterVendorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterVendorResponse) ProtoMessage() {}

func (x *RegisterVendorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_vendor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterVendorResponse.ProtoReflect.Descriptor instead.
func (*RegisterVendorResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_vendor_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterVendorResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RegisterVendorResponse) GetVendorId() string {
	if x != nil {
		return x.VendorId
	}
	return ""
}

var File_api_proto_vendor_proto protoreflect.FileDescriptor

var file_api_proto_vendor_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x22, 0x41, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x4e, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x56,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x49, 0x64, 0x32, 0x60, 0x0a, 0x0d, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1d, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_proto_vendor_proto_rawDescOnce sync.Once
	file_api_proto_vendor_proto_rawDescData = file_api_proto_vendor_proto_rawDesc
)

func file_api_proto_vendor_proto_rawDescGZIP() []byte {
	file_api_proto_vendor_proto_rawDescOnce.Do(func() {
		file_api_proto_vendor_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_vendor_proto_rawDescData)
	})
	return file_api_proto_vendor_proto_rawDescData
}

var file_api_proto_vendor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_vendor_proto_goTypes = []any{
	(*RegisterVendorRequest)(nil),  // 0: vendor.RegisterVendorRequest
	(*RegisterVendorResponse)(nil), // 1: vendor.RegisterVendorResponse
}
var file_api_proto_vendor_proto_depIdxs = []int32{
	0, // 0: vendor.VendorService.RegisterVendor:input_type -> vendor.RegisterVendorRequest
	1, // 1: vendor.VendorService.RegisterVendor:output_type -> vendor.RegisterVendorResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_vendor_proto_init() }
func file_api_proto_vendor_proto_init() {
	if File_api_proto_vendor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_vendor_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterVendorRequest); i {
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
		file_api_proto_vendor_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterVendorResponse); i {
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
			RawDescriptor: file_api_proto_vendor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_vendor_proto_goTypes,
		DependencyIndexes: file_api_proto_vendor_proto_depIdxs,
		MessageInfos:      file_api_proto_vendor_proto_msgTypes,
	}.Build()
	File_api_proto_vendor_proto = out.File
	file_api_proto_vendor_proto_rawDesc = nil
	file_api_proto_vendor_proto_goTypes = nil
	file_api_proto_vendor_proto_depIdxs = nil
}