// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: lib/proto/auth.proto

package libproto

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

type ExchangeAuthCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthCode string `protobuf:"bytes,1,opt,name=auth_code,json=authCode,proto3" json:"auth_code,omitempty"`
}

func (x *ExchangeAuthCodeRequest) Reset() {
	*x = ExchangeAuthCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeAuthCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeAuthCodeRequest) ProtoMessage() {}

func (x *ExchangeAuthCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeAuthCodeRequest.ProtoReflect.Descriptor instead.
func (*ExchangeAuthCodeRequest) Descriptor() ([]byte, []int) {
	return file_lib_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeAuthCodeRequest) GetAuthCode() string {
	if x != nil {
		return x.AuthCode
	}
	return ""
}

type ExchangeAuthCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *ExchangeAuthCodeResponse) Reset() {
	*x = ExchangeAuthCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeAuthCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeAuthCodeResponse) ProtoMessage() {}

func (x *ExchangeAuthCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeAuthCodeResponse.ProtoReflect.Descriptor instead.
func (*ExchangeAuthCodeResponse) Descriptor() ([]byte, []int) {
	return file_lib_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeAuthCodeResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type VerifyAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *VerifyAccessTokenRequest) Reset() {
	*x = VerifyAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyAccessTokenRequest) ProtoMessage() {}

func (x *VerifyAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*VerifyAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_lib_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyAccessTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type VerifyAccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role   string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *VerifyAccessTokenResponse) Reset() {
	*x = VerifyAccessTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyAccessTokenResponse) ProtoMessage() {}

func (x *VerifyAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*VerifyAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_lib_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyAccessTokenResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *VerifyAccessTokenResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_lib_proto_auth_proto protoreflect.FileDescriptor

var file_lib_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x36, 0x0a, 0x17,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x3d, 0x0a, 0x18, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x3d, 0x0a, 0x18, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x48, 0x0a, 0x19, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x32, 0xb3, 0x01, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x53, 0x0a, 0x10, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x11, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6c, 0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lib_proto_auth_proto_rawDescOnce sync.Once
	file_lib_proto_auth_proto_rawDescData = file_lib_proto_auth_proto_rawDesc
)

func file_lib_proto_auth_proto_rawDescGZIP() []byte {
	file_lib_proto_auth_proto_rawDescOnce.Do(func() {
		file_lib_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_proto_auth_proto_rawDescData)
	})
	return file_lib_proto_auth_proto_rawDescData
}

var file_lib_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_lib_proto_auth_proto_goTypes = []interface{}{
	(*ExchangeAuthCodeRequest)(nil),   // 0: auth.ExchangeAuthCodeRequest
	(*ExchangeAuthCodeResponse)(nil),  // 1: auth.ExchangeAuthCodeResponse
	(*VerifyAccessTokenRequest)(nil),  // 2: auth.VerifyAccessTokenRequest
	(*VerifyAccessTokenResponse)(nil), // 3: auth.VerifyAccessTokenResponse
}
var file_lib_proto_auth_proto_depIdxs = []int32{
	0, // 0: auth.Auth.ExchangeAuthCode:input_type -> auth.ExchangeAuthCodeRequest
	2, // 1: auth.Auth.VerifyAccessToken:input_type -> auth.VerifyAccessTokenRequest
	1, // 2: auth.Auth.ExchangeAuthCode:output_type -> auth.ExchangeAuthCodeResponse
	3, // 3: auth.Auth.VerifyAccessToken:output_type -> auth.VerifyAccessTokenResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lib_proto_auth_proto_init() }
func file_lib_proto_auth_proto_init() {
	if File_lib_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeAuthCodeRequest); i {
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
		file_lib_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeAuthCodeResponse); i {
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
		file_lib_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyAccessTokenRequest); i {
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
		file_lib_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyAccessTokenResponse); i {
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
			RawDescriptor: file_lib_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lib_proto_auth_proto_goTypes,
		DependencyIndexes: file_lib_proto_auth_proto_depIdxs,
		MessageInfos:      file_lib_proto_auth_proto_msgTypes,
	}.Build()
	File_lib_proto_auth_proto = out.File
	file_lib_proto_auth_proto_rawDesc = nil
	file_lib_proto_auth_proto_goTypes = nil
	file_lib_proto_auth_proto_depIdxs = nil
}
