// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.9.0
// source: etcd_bridge_users.proto

package etcdbridgeproto

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etcd_bridge_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_etcd_bridge_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_etcd_bridge_users_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etcd_bridge_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_etcd_bridge_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_etcd_bridge_users_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Status   int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Roles    []string `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etcd_bridge_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_etcd_bridge_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_etcd_bridge_users_proto_rawDescGZIP(), []int{2}
}

func (x *Users) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Users) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Users) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Users) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Users) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Status   int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Roles    []string `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *UsersRequest) Reset() {
	*x = UsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etcd_bridge_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersRequest) ProtoMessage() {}

func (x *UsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_etcd_bridge_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersRequest.ProtoReflect.Descriptor instead.
func (*UsersRequest) Descriptor() ([]byte, []int) {
	return file_etcd_bridge_users_proto_rawDescGZIP(), []int{3}
}

func (x *UsersRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UsersRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UsersRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UsersRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UsersRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UsersLisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchKey string `protobuf:"bytes,1,opt,name=search_key,json=searchKey,proto3" json:"search_key,omitempty"`
	Page      int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *UsersLisRequest) Reset() {
	*x = UsersLisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etcd_bridge_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersLisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersLisRequest) ProtoMessage() {}

func (x *UsersLisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_etcd_bridge_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersLisRequest.ProtoReflect.Descriptor instead.
func (*UsersLisRequest) Descriptor() ([]byte, []int) {
	return file_etcd_bridge_users_proto_rawDescGZIP(), []int{4}
}

func (x *UsersLisRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

func (x *UsersLisRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *UsersLisRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type DelUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelUsersRequest) Reset() {
	*x = DelUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etcd_bridge_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelUsersRequest) ProtoMessage() {}

func (x *DelUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_etcd_bridge_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelUsersRequest.ProtoReflect.Descriptor instead.
func (*DelUsersRequest) Descriptor() ([]byte, []int) {
	return file_etcd_bridge_users_proto_rawDescGZIP(), []int{5}
}

func (x *DelUsersRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UsersListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*Users `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *UsersListResponse) Reset() {
	*x = UsersListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etcd_bridge_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersListResponse) ProtoMessage() {}

func (x *UsersListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_etcd_bridge_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersListResponse.ProtoReflect.Descriptor instead.
func (*UsersListResponse) Descriptor() ([]byte, []int) {
	return file_etcd_bridge_users_proto_rawDescGZIP(), []int{6}
}

func (x *UsersListResponse) GetData() []*Users {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UsersListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_etcd_bridge_users_proto protoreflect.FileDescriptor

var file_etcd_bridge_users_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x74, 0x63, 0x64, 0x5f,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x25,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7d, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x61, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x21,
	0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x51, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x3b, 0x65, 0x74, 0x63, 0x64, 0x62, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_etcd_bridge_users_proto_rawDescOnce sync.Once
	file_etcd_bridge_users_proto_rawDescData = file_etcd_bridge_users_proto_rawDesc
)

func file_etcd_bridge_users_proto_rawDescGZIP() []byte {
	file_etcd_bridge_users_proto_rawDescOnce.Do(func() {
		file_etcd_bridge_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_etcd_bridge_users_proto_rawDescData)
	})
	return file_etcd_bridge_users_proto_rawDescData
}

var file_etcd_bridge_users_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_etcd_bridge_users_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),      // 0: etcd_bridge.LoginRequest
	(*LoginResponse)(nil),     // 1: etcd_bridge.LoginResponse
	(*Users)(nil),             // 2: etcd_bridge.Users
	(*UsersRequest)(nil),      // 3: etcd_bridge.UsersRequest
	(*UsersLisRequest)(nil),   // 4: etcd_bridge.UsersLisRequest
	(*DelUsersRequest)(nil),   // 5: etcd_bridge.DelUsersRequest
	(*UsersListResponse)(nil), // 6: etcd_bridge.UsersListResponse
}
var file_etcd_bridge_users_proto_depIdxs = []int32{
	2, // 0: etcd_bridge.UsersListResponse.data:type_name -> etcd_bridge.Users
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_etcd_bridge_users_proto_init() }
func file_etcd_bridge_users_proto_init() {
	if File_etcd_bridge_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_etcd_bridge_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_etcd_bridge_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_etcd_bridge_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_etcd_bridge_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersRequest); i {
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
		file_etcd_bridge_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersLisRequest); i {
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
		file_etcd_bridge_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelUsersRequest); i {
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
		file_etcd_bridge_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersListResponse); i {
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
			RawDescriptor: file_etcd_bridge_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_etcd_bridge_users_proto_goTypes,
		DependencyIndexes: file_etcd_bridge_users_proto_depIdxs,
		MessageInfos:      file_etcd_bridge_users_proto_msgTypes,
	}.Build()
	File_etcd_bridge_users_proto = out.File
	file_etcd_bridge_users_proto_rawDesc = nil
	file_etcd_bridge_users_proto_goTypes = nil
	file_etcd_bridge_users_proto_depIdxs = nil
}
