// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: singer.proto

package signer

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

// 第一部分，应该是发送给下一签名者来构造顺序签名
type SendSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	R         []byte `protobuf:"bytes,2,opt,name=R,proto3" json:"R,omitempty"`
}

func (x *SendSignature) Reset() {
	*x = SendSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_singer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSignature) ProtoMessage() {}

func (x *SendSignature) ProtoReflect() protoreflect.Message {
	mi := &file_singer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSignature.ProtoReflect.Descriptor instead.
func (*SendSignature) Descriptor() ([]byte, []int) {
	return file_singer_proto_rawDescGZIP(), []int{0}
}

func (x *SendSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SendSignature) GetR() []byte {
	if x != nil {
		return x.R
	}
	return nil
}

type SendSignatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendSignatureResponse) Reset() {
	*x = SendSignatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_singer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSignatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSignatureResponse) ProtoMessage() {}

func (x *SendSignatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_singer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSignatureResponse.ProtoReflect.Descriptor instead.
func (*SendSignatureResponse) Descriptor() ([]byte, []int) {
	return file_singer_proto_rawDescGZIP(), []int{1}
}

type SendIBESASSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X []byte `protobuf:"bytes,1,opt,name=X,proto3" json:"X,omitempty"`
	Y []byte `protobuf:"bytes,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z []byte `protobuf:"bytes,3,opt,name=Z,proto3" json:"Z,omitempty"`
}

func (x *SendIBESASSignature) Reset() {
	*x = SendIBESASSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_singer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendIBESASSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendIBESASSignature) ProtoMessage() {}

func (x *SendIBESASSignature) ProtoReflect() protoreflect.Message {
	mi := &file_singer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendIBESASSignature.ProtoReflect.Descriptor instead.
func (*SendIBESASSignature) Descriptor() ([]byte, []int) {
	return file_singer_proto_rawDescGZIP(), []int{2}
}

func (x *SendIBESASSignature) GetX() []byte {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *SendIBESASSignature) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

func (x *SendIBESASSignature) GetZ() []byte {
	if x != nil {
		return x.Z
	}
	return nil
}

type SendIBESASSignatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendIBESASSignatureResponse) Reset() {
	*x = SendIBESASSignatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_singer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendIBESASSignatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendIBESASSignatureResponse) ProtoMessage() {}

func (x *SendIBESASSignatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_singer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendIBESASSignatureResponse.ProtoReflect.Descriptor instead.
func (*SendIBESASSignatureResponse) Descriptor() ([]byte, []int) {
	return file_singer_proto_rawDescGZIP(), []int{3}
}

var File_singer_proto protoreflect.FileDescriptor

var file_singer_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x52, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x01, 0x52, 0x22, 0x17, 0x0a, 0x15, 0x73, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x0a, 0x13,
	0x73, 0x65, 0x6e, 0x64, 0x49, 0x42, 0x45, 0x53, 0x41, 0x53, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01,
	0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x59, 0x12,
	0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x5a, 0x22, 0x1d, 0x0a,
	0x1b, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x42, 0x45, 0x53, 0x41, 0x53, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xad, 0x01, 0x0a,
	0x06, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x4f,
	0x77, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x1a, 0x1d, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x6e, 0x64,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x59, 0x0a, 0x15, 0x73, 0x65, 0x6e, 0x64, 0x4f, 0x77, 0x6e, 0x49, 0x42, 0x53, 0x41,
	0x53, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x42, 0x45, 0x53, 0x41, 0x53, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x23, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x42, 0x45, 0x53, 0x41, 0x53, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_singer_proto_rawDescOnce sync.Once
	file_singer_proto_rawDescData = file_singer_proto_rawDesc
)

func file_singer_proto_rawDescGZIP() []byte {
	file_singer_proto_rawDescOnce.Do(func() {
		file_singer_proto_rawDescData = protoimpl.X.CompressGZIP(file_singer_proto_rawDescData)
	})
	return file_singer_proto_rawDescData
}

var file_singer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_singer_proto_goTypes = []interface{}{
	(*SendSignature)(nil),               // 0: signer.sendSignature
	(*SendSignatureResponse)(nil),       // 1: signer.sendSignatureResponse
	(*SendIBESASSignature)(nil),         // 2: signer.sendIBESASSignature
	(*SendIBESASSignatureResponse)(nil), // 3: signer.sendIBESASSignatureResponse
}
var file_singer_proto_depIdxs = []int32{
	0, // 0: signer.Signer.sendOwnSignature:input_type -> signer.sendSignature
	2, // 1: signer.Signer.sendOwnIBSASSignature:input_type -> signer.sendIBESASSignature
	1, // 2: signer.Signer.sendOwnSignature:output_type -> signer.sendSignatureResponse
	3, // 3: signer.Signer.sendOwnIBSASSignature:output_type -> signer.sendIBESASSignatureResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_singer_proto_init() }
func file_singer_proto_init() {
	if File_singer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_singer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSignature); i {
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
		file_singer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSignatureResponse); i {
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
		file_singer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendIBESASSignature); i {
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
		file_singer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendIBESASSignatureResponse); i {
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
			RawDescriptor: file_singer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_singer_proto_goTypes,
		DependencyIndexes: file_singer_proto_depIdxs,
		MessageInfos:      file_singer_proto_msgTypes,
	}.Build()
	File_singer_proto = out.File
	file_singer_proto_rawDesc = nil
	file_singer_proto_goTypes = nil
	file_singer_proto_depIdxs = nil
}
