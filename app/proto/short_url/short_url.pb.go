// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: short_url/short_url.proto

package url_shortener_go

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShortUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *ShortUrl) Reset() {
	*x = ShortUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_short_url_short_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortUrl) ProtoMessage() {}

func (x *ShortUrl) ProtoReflect() protoreflect.Message {
	mi := &file_short_url_short_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortUrl.ProtoReflect.Descriptor instead.
func (*ShortUrl) Descriptor() ([]byte, []int) {
	return file_short_url_short_url_proto_rawDescGZIP(), []int{0}
}

func (x *ShortUrl) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *ShortUrl) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type CreateShortUrlReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateShortUrlReq) Reset() {
	*x = CreateShortUrlReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_short_url_short_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortUrlReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortUrlReq) ProtoMessage() {}

func (x *CreateShortUrlReq) ProtoReflect() protoreflect.Message {
	mi := &file_short_url_short_url_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortUrlReq.ProtoReflect.Descriptor instead.
func (*CreateShortUrlReq) Descriptor() ([]byte, []int) {
	return file_short_url_short_url_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShortUrlReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateShortUrlResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl *ShortUrl `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *CreateShortUrlResp) Reset() {
	*x = CreateShortUrlResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_short_url_short_url_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortUrlResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortUrlResp) ProtoMessage() {}

func (x *CreateShortUrlResp) ProtoReflect() protoreflect.Message {
	mi := &file_short_url_short_url_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortUrlResp.ProtoReflect.Descriptor instead.
func (*CreateShortUrlResp) Descriptor() ([]byte, []int) {
	return file_short_url_short_url_proto_rawDescGZIP(), []int{2}
}

func (x *CreateShortUrlResp) GetShortUrl() *ShortUrl {
	if x != nil {
		return x.ShortUrl
	}
	return nil
}

type RedirectToShortUrlReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *RedirectToShortUrlReq) Reset() {
	*x = RedirectToShortUrlReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_short_url_short_url_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectToShortUrlReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectToShortUrlReq) ProtoMessage() {}

func (x *RedirectToShortUrlReq) ProtoReflect() protoreflect.Message {
	mi := &file_short_url_short_url_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectToShortUrlReq.ProtoReflect.Descriptor instead.
func (*RedirectToShortUrlReq) Descriptor() ([]byte, []int) {
	return file_short_url_short_url_proto_rawDescGZIP(), []int{3}
}

func (x *RedirectToShortUrlReq) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type RedirectToShortUrlResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RedirectToShortUrlResp) Reset() {
	*x = RedirectToShortUrlResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_short_url_short_url_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectToShortUrlResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectToShortUrlResp) ProtoMessage() {}

func (x *RedirectToShortUrlResp) ProtoReflect() protoreflect.Message {
	mi := &file_short_url_short_url_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectToShortUrlResp.ProtoReflect.Descriptor instead.
func (*RedirectToShortUrlResp) Descriptor() ([]byte, []int) {
	return file_short_url_short_url_proto_rawDescGZIP(), []int{4}
}

func (x *RedirectToShortUrlResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_short_url_short_url_proto protoreflect.FileDescriptor

var file_short_url_short_url_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4a, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x25, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c,
	0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x34, 0x0a, 0x15, 0x52, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c,
	0x22, 0x2a, 0x0a, 0x16, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0xc1, 0x02, 0x0a,
	0x0c, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x95, 0x01,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c,
	0x12, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1d,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x46, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x3a, 0x01, 0x2a, 0x62, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x92, 0x41, 0x1f, 0x4a, 0x1d, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x16, 0x0a,
	0x14, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x20, 0x55, 0x52, 0x4c, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2e, 0x12, 0x98, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x54, 0x6f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x54, 0x6f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x21,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x54, 0x6f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x92, 0x41, 0x24, 0x4a, 0x22, 0x0a, 0x03,
	0x33, 0x30, 0x32, 0x12, 0x1b, 0x0a, 0x19, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x20,
	0x74, 0x6f, 0x20, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x20, 0x55, 0x52, 0x4c, 0x2e,
	0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x33,
	0x46, 0x61, 0x6e, 0x59, 0x75, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_short_url_short_url_proto_rawDescOnce sync.Once
	file_short_url_short_url_proto_rawDescData = file_short_url_short_url_proto_rawDesc
)

func file_short_url_short_url_proto_rawDescGZIP() []byte {
	file_short_url_short_url_proto_rawDescOnce.Do(func() {
		file_short_url_short_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_short_url_short_url_proto_rawDescData)
	})
	return file_short_url_short_url_proto_rawDescData
}

var file_short_url_short_url_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_short_url_short_url_proto_goTypes = []interface{}{
	(*ShortUrl)(nil),               // 0: short_url.ShortUrl
	(*CreateShortUrlReq)(nil),      // 1: short_url.CreateShortUrlReq
	(*CreateShortUrlResp)(nil),     // 2: short_url.CreateShortUrlResp
	(*RedirectToShortUrlReq)(nil),  // 3: short_url.RedirectToShortUrlReq
	(*RedirectToShortUrlResp)(nil), // 4: short_url.RedirectToShortUrlResp
}
var file_short_url_short_url_proto_depIdxs = []int32{
	0, // 0: short_url.CreateShortUrlResp.short_url:type_name -> short_url.ShortUrl
	1, // 1: short_url.UrlShortener.CreateShortUrl:input_type -> short_url.CreateShortUrlReq
	3, // 2: short_url.UrlShortener.RedirectToShortUrl:input_type -> short_url.RedirectToShortUrlReq
	2, // 3: short_url.UrlShortener.CreateShortUrl:output_type -> short_url.CreateShortUrlResp
	4, // 4: short_url.UrlShortener.RedirectToShortUrl:output_type -> short_url.RedirectToShortUrlResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_short_url_short_url_proto_init() }
func file_short_url_short_url_proto_init() {
	if File_short_url_short_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_short_url_short_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortUrl); i {
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
		file_short_url_short_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortUrlReq); i {
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
		file_short_url_short_url_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortUrlResp); i {
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
		file_short_url_short_url_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectToShortUrlReq); i {
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
		file_short_url_short_url_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectToShortUrlResp); i {
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
			RawDescriptor: file_short_url_short_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_short_url_short_url_proto_goTypes,
		DependencyIndexes: file_short_url_short_url_proto_depIdxs,
		MessageInfos:      file_short_url_short_url_proto_msgTypes,
	}.Build()
	File_short_url_short_url_proto = out.File
	file_short_url_short_url_proto_rawDesc = nil
	file_short_url_short_url_proto_goTypes = nil
	file_short_url_short_url_proto_depIdxs = nil
}
