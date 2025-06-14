// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.9.0
// source: conf.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

// 配置
type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`  //主键
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"` //key
	Val string `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"` //值
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conf.ProtoReflect.Descriptor instead.
func (*Conf) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Conf) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Conf) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Conf) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

// 获取配置列表
type GetConfListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keywords string `protobuf:"bytes,1,opt,name=keywords,proto3" json:"keywords,omitempty"` //搜索关键字
}

func (x *GetConfListRequest) Reset() {
	*x = GetConfListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfListRequest) ProtoMessage() {}

func (x *GetConfListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfListRequest.ProtoReflect.Descriptor instead.
func (*GetConfListRequest) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{1}
}

func (x *GetConfListRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

// 配置列表响应
type ConfListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []*Conf `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`       //配置列表数组
	Code    uint32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      //状态码 0=失败,1=成功
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` //信息
	Count   int64   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`    //总条数
}

func (x *ConfListResponse) Reset() {
	*x = ConfListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfListResponse) ProtoMessage() {}

func (x *ConfListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfListResponse.ProtoReflect.Descriptor instead.
func (*ConfListResponse) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{2}
}

func (x *ConfListResponse) GetData() []*Conf {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ConfListResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConfListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConfListResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 设置配置请求
type SetConfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`  //主键
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"` //key
	Val string `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"` //值
}

func (x *SetConfRequest) Reset() {
	*x = SetConfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetConfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConfRequest) ProtoMessage() {}

func (x *SetConfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConfRequest.ProtoReflect.Descriptor instead.
func (*SetConfRequest) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{3}
}

func (x *SetConfRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetConfRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetConfRequest) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

// 删除配置
type DelConfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelConfRequest) Reset() {
	*x = DelConfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelConfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelConfRequest) ProtoMessage() {}

func (x *DelConfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelConfRequest.ProtoReflect.Descriptor instead.
func (*DelConfRequest) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{4}
}

func (x *DelConfRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 配置响应
type ConfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      //状态码 0=失败,1=成功
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` //信息
}

func (x *ConfResponse) Reset() {
	*x = ConfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfResponse) ProtoMessage() {}

func (x *ConfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfResponse.ProtoReflect.Descriptor instead.
func (*ConfResponse) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{5}
}

func (x *ConfResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConfResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_conf_proto protoreflect.FileDescriptor

var file_conf_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x30, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0x7c, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x44, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa0, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x12, 0x1a, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x2e,
	0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x07, 0x44, 0x65,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x1a, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x6e, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06,
	0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18,
	0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_proto_rawDescOnce sync.Once
	file_conf_proto_rawDescData = file_conf_proto_rawDesc
)

func file_conf_proto_rawDescGZIP() []byte {
	file_conf_proto_rawDescOnce.Do(func() {
		file_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_proto_rawDescData)
	})
	return file_conf_proto_rawDescData
}

var file_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_conf_proto_goTypes = []interface{}{
	(*Conf)(nil),               // 0: super_send.Conf
	(*GetConfListRequest)(nil), // 1: super_send.GetConfListRequest
	(*ConfListResponse)(nil),   // 2: super_send.ConfListResponse
	(*SetConfRequest)(nil),     // 3: super_send.SetConfRequest
	(*DelConfRequest)(nil),     // 4: super_send.DelConfRequest
	(*ConfResponse)(nil),       // 5: super_send.ConfResponse
	(*empty.Empty)(nil),        // 6: google.protobuf.Empty
}
var file_conf_proto_depIdxs = []int32{
	0, // 0: super_send.ConfListResponse.data:type_name -> super_send.Conf
	3, // 1: super_send.ConfService.SetConf:input_type -> super_send.SetConfRequest
	4, // 2: super_send.ConfService.DelConf:input_type -> super_send.DelConfRequest
	1, // 3: super_send.ConfService.GetConfList:input_type -> super_send.GetConfListRequest
	6, // 4: super_send.ConfService.Reload:input_type -> google.protobuf.Empty
	5, // 5: super_send.ConfService.SetConf:output_type -> super_send.ConfResponse
	5, // 6: super_send.ConfService.DelConf:output_type -> super_send.ConfResponse
	2, // 7: super_send.ConfService.GetConfList:output_type -> super_send.ConfListResponse
	5, // 8: super_send.ConfService.Reload:output_type -> super_send.ConfResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_conf_proto_init() }
func file_conf_proto_init() {
	if File_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conf); i {
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
		file_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfListRequest); i {
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
		file_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfListResponse); i {
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
		file_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetConfRequest); i {
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
		file_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelConfRequest); i {
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
		file_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfResponse); i {
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
			RawDescriptor: file_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conf_proto_goTypes,
		DependencyIndexes: file_conf_proto_depIdxs,
		MessageInfos:      file_conf_proto_msgTypes,
	}.Build()
	File_conf_proto = out.File
	file_conf_proto_rawDesc = nil
	file_conf_proto_goTypes = nil
	file_conf_proto_depIdxs = nil
}
