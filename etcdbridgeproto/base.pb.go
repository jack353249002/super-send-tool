// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.9.0
// source: base.proto

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

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*BaseResponse_ServerInfoList
	//	*BaseResponse_SyncConfList
	//	*BaseResponse_ServerInfo
	//	*BaseResponse_UsersInfo
	//	*BaseResponse_UsersList
	//	*BaseResponse_SyncConfInfo
	//	*BaseResponse_RolesList
	//	*BaseResponse_RolesPermissionsList
	//	*BaseResponse_LoginInfo
	Payload isBaseResponse_Payload `protobuf_oneof:"payload"`
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaseResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (m *BaseResponse) GetPayload() isBaseResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *BaseResponse) GetServerInfoList() *ServerInfoListResponse {
	if x, ok := x.GetPayload().(*BaseResponse_ServerInfoList); ok {
		return x.ServerInfoList
	}
	return nil
}

func (x *BaseResponse) GetSyncConfList() *SyncConfListResponse {
	if x, ok := x.GetPayload().(*BaseResponse_SyncConfList); ok {
		return x.SyncConfList
	}
	return nil
}

func (x *BaseResponse) GetServerInfo() *ServerInfo {
	if x, ok := x.GetPayload().(*BaseResponse_ServerInfo); ok {
		return x.ServerInfo
	}
	return nil
}

func (x *BaseResponse) GetUsersInfo() *Users {
	if x, ok := x.GetPayload().(*BaseResponse_UsersInfo); ok {
		return x.UsersInfo
	}
	return nil
}

func (x *BaseResponse) GetUsersList() *UsersListResponse {
	if x, ok := x.GetPayload().(*BaseResponse_UsersList); ok {
		return x.UsersList
	}
	return nil
}

func (x *BaseResponse) GetSyncConfInfo() *SyncConf {
	if x, ok := x.GetPayload().(*BaseResponse_SyncConfInfo); ok {
		return x.SyncConfInfo
	}
	return nil
}

func (x *BaseResponse) GetRolesList() *RolesListResponse {
	if x, ok := x.GetPayload().(*BaseResponse_RolesList); ok {
		return x.RolesList
	}
	return nil
}

func (x *BaseResponse) GetRolesPermissionsList() *RolesPermissionsListResponse {
	if x, ok := x.GetPayload().(*BaseResponse_RolesPermissionsList); ok {
		return x.RolesPermissionsList
	}
	return nil
}

func (x *BaseResponse) GetLoginInfo() *LoginResponse {
	if x, ok := x.GetPayload().(*BaseResponse_LoginInfo); ok {
		return x.LoginInfo
	}
	return nil
}

type isBaseResponse_Payload interface {
	isBaseResponse_Payload()
}

type BaseResponse_ServerInfoList struct {
	ServerInfoList *ServerInfoListResponse `protobuf:"bytes,3,opt,name=server_info_list,json=serverInfoList,proto3,oneof"`
}

type BaseResponse_SyncConfList struct {
	SyncConfList *SyncConfListResponse `protobuf:"bytes,4,opt,name=sync_conf_list,json=syncConfList,proto3,oneof"`
}

type BaseResponse_ServerInfo struct {
	ServerInfo *ServerInfo `protobuf:"bytes,5,opt,name=server_info,json=serverInfo,proto3,oneof"`
}

type BaseResponse_UsersInfo struct {
	UsersInfo *Users `protobuf:"bytes,6,opt,name=users_info,json=usersInfo,proto3,oneof"`
}

type BaseResponse_UsersList struct {
	UsersList *UsersListResponse `protobuf:"bytes,7,opt,name=users_list,json=usersList,proto3,oneof"`
}

type BaseResponse_SyncConfInfo struct {
	SyncConfInfo *SyncConf `protobuf:"bytes,8,opt,name=sync_conf_info,json=syncConfInfo,proto3,oneof"`
}

type BaseResponse_RolesList struct {
	RolesList *RolesListResponse `protobuf:"bytes,9,opt,name=roles_list,json=rolesList,proto3,oneof"`
}

type BaseResponse_RolesPermissionsList struct {
	RolesPermissionsList *RolesPermissionsListResponse `protobuf:"bytes,10,opt,name=roles_permissions_list,json=rolesPermissionsList,proto3,oneof"`
}

type BaseResponse_LoginInfo struct {
	LoginInfo *LoginResponse `protobuf:"bytes,11,opt,name=login_info,json=loginInfo,proto3,oneof"`
}

func (*BaseResponse_ServerInfoList) isBaseResponse_Payload() {}

func (*BaseResponse_SyncConfList) isBaseResponse_Payload() {}

func (*BaseResponse_ServerInfo) isBaseResponse_Payload() {}

func (*BaseResponse_UsersInfo) isBaseResponse_Payload() {}

func (*BaseResponse_UsersList) isBaseResponse_Payload() {}

func (*BaseResponse_SyncConfInfo) isBaseResponse_Payload() {}

func (*BaseResponse_RolesList) isBaseResponse_Payload() {}

func (*BaseResponse_RolesPermissionsList) isBaseResponse_Payload() {}

func (*BaseResponse_LoginInfo) isBaseResponse_Payload() {}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x74,
	0x63, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x1a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65,
	0x74, 0x63, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x62, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xad, 0x05, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x4f, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x0e, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x62,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x62,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65,
	0x74, 0x63, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x43,
	0x6f, 0x6e, 0x66, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x62,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x61, 0x0a, 0x16, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x00, 0x52, 0x14, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x74,
	0x63, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42,
	0x13, 0x5a, 0x11, 0x2e, 0x3b, 0x65, 0x74, 0x63, 0x64, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_base_proto_goTypes = []interface{}{
	(*BaseResponse)(nil),                 // 0: etcd_bridge.BaseResponse
	(*ServerInfoListResponse)(nil),       // 1: etcd_bridge.ServerInfoListResponse
	(*SyncConfListResponse)(nil),         // 2: etcd_bridge.SyncConfListResponse
	(*ServerInfo)(nil),                   // 3: etcd_bridge.ServerInfo
	(*Users)(nil),                        // 4: etcd_bridge.Users
	(*UsersListResponse)(nil),            // 5: etcd_bridge.UsersListResponse
	(*SyncConf)(nil),                     // 6: etcd_bridge.SyncConf
	(*RolesListResponse)(nil),            // 7: etcd_bridge.RolesListResponse
	(*RolesPermissionsListResponse)(nil), // 8: etcd_bridge.RolesPermissionsListResponse
	(*LoginResponse)(nil),                // 9: etcd_bridge.LoginResponse
}
var file_base_proto_depIdxs = []int32{
	1, // 0: etcd_bridge.BaseResponse.server_info_list:type_name -> etcd_bridge.ServerInfoListResponse
	2, // 1: etcd_bridge.BaseResponse.sync_conf_list:type_name -> etcd_bridge.SyncConfListResponse
	3, // 2: etcd_bridge.BaseResponse.server_info:type_name -> etcd_bridge.ServerInfo
	4, // 3: etcd_bridge.BaseResponse.users_info:type_name -> etcd_bridge.Users
	5, // 4: etcd_bridge.BaseResponse.users_list:type_name -> etcd_bridge.UsersListResponse
	6, // 5: etcd_bridge.BaseResponse.sync_conf_info:type_name -> etcd_bridge.SyncConf
	7, // 6: etcd_bridge.BaseResponse.roles_list:type_name -> etcd_bridge.RolesListResponse
	8, // 7: etcd_bridge.BaseResponse.roles_permissions_list:type_name -> etcd_bridge.RolesPermissionsListResponse
	9, // 8: etcd_bridge.BaseResponse.login_info:type_name -> etcd_bridge.LoginResponse
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	file_server_info_proto_init()
	file_sync_conf_proto_init()
	file_etcd_bridge_users_proto_init()
	file_etcd_bridge_roles_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
	file_base_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BaseResponse_ServerInfoList)(nil),
		(*BaseResponse_SyncConfList)(nil),
		(*BaseResponse_ServerInfo)(nil),
		(*BaseResponse_UsersInfo)(nil),
		(*BaseResponse_UsersList)(nil),
		(*BaseResponse_SyncConfInfo)(nil),
		(*BaseResponse_RolesList)(nil),
		(*BaseResponse_RolesPermissionsList)(nil),
		(*BaseResponse_LoginInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
