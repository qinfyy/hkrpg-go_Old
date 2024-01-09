// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: Gateserver.proto

package proto

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

type Gateserver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip              string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	ClientSecretKey string `protobuf:"bytes,239,opt,name=client_secret_key,json=clientSecretKey,proto3" json:"client_secret_key,omitempty"`
	MdkResVersion   string `protobuf:"bytes,1506,opt,name=mdk_res_version,json=mdkResVersion,proto3" json:"mdk_res_version,omitempty"`
	AssetBundleUrl  string `protobuf:"bytes,14,opt,name=asset_bundle_url,json=assetBundleUrl,proto3" json:"asset_bundle_url,omitempty"`
	ExResourceUrl   string `protobuf:"bytes,12,opt,name=ex_resource_url,json=exResourceUrl,proto3" json:"ex_resource_url,omitempty"`
	Port            uint32 `protobuf:"varint,15,opt,name=port,proto3" json:"port,omitempty"`
	LuaUrl          string `protobuf:"bytes,10,opt,name=lua_url,json=luaUrl,proto3" json:"lua_url,omitempty"`
	IfixVersion     string `protobuf:"bytes,20,opt,name=ifix_version,json=ifixVersion,proto3" json:"ifix_version,omitempty"`
	IfixUrl         string `protobuf:"bytes,98,opt,name=ifix_url,json=ifixUrl,proto3" json:"ifix_url,omitempty"`
	RegionName      string `protobuf:"bytes,4,opt,name=region_name,json=regionName,proto3" json:"region_name,omitempty"`
	Msg             string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Unk1            bool   `protobuf:"varint,5,opt,name=unk1,proto3" json:"unk1,omitempty"`
	Unk2            bool   `protobuf:"varint,9,opt,name=unk2,proto3" json:"unk2,omitempty"`
	Unk3            bool   `protobuf:"varint,170,opt,name=unk3,proto3" json:"unk3,omitempty"`
	Unk4            bool   `protobuf:"varint,1397,opt,name=unk4,proto3" json:"unk4,omitempty"`
	Unk5            bool   `protobuf:"varint,1649,opt,name=unk5,proto3" json:"unk5,omitempty"`
}

func (x *Gateserver) Reset() {
	*x = Gateserver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Gateserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gateserver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gateserver) ProtoMessage() {}

func (x *Gateserver) ProtoReflect() protoreflect.Message {
	mi := &file_Gateserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gateserver.ProtoReflect.Descriptor instead.
func (*Gateserver) Descriptor() ([]byte, []int) {
	return file_Gateserver_proto_rawDescGZIP(), []int{0}
}

func (x *Gateserver) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Gateserver) GetClientSecretKey() string {
	if x != nil {
		return x.ClientSecretKey
	}
	return ""
}

func (x *Gateserver) GetMdkResVersion() string {
	if x != nil {
		return x.MdkResVersion
	}
	return ""
}

func (x *Gateserver) GetAssetBundleUrl() string {
	if x != nil {
		return x.AssetBundleUrl
	}
	return ""
}

func (x *Gateserver) GetExResourceUrl() string {
	if x != nil {
		return x.ExResourceUrl
	}
	return ""
}

func (x *Gateserver) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Gateserver) GetLuaUrl() string {
	if x != nil {
		return x.LuaUrl
	}
	return ""
}

func (x *Gateserver) GetIfixVersion() string {
	if x != nil {
		return x.IfixVersion
	}
	return ""
}

func (x *Gateserver) GetIfixUrl() string {
	if x != nil {
		return x.IfixUrl
	}
	return ""
}

func (x *Gateserver) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *Gateserver) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Gateserver) GetUnk1() bool {
	if x != nil {
		return x.Unk1
	}
	return false
}

func (x *Gateserver) GetUnk2() bool {
	if x != nil {
		return x.Unk2
	}
	return false
}

func (x *Gateserver) GetUnk3() bool {
	if x != nil {
		return x.Unk3
	}
	return false
}

func (x *Gateserver) GetUnk4() bool {
	if x != nil {
		return x.Unk4
	}
	return false
}

func (x *Gateserver) GetUnk5() bool {
	if x != nil {
		return x.Unk5
	}
	return false
}

var File_Gateserver_proto protoreflect.FileDescriptor

var file_Gateserver_proto_rawDesc = []byte{
	0x0a, 0x10, 0x47, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc9, 0x03, 0x0a, 0x0a, 0x47, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0xef, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x27,
	0x0a, 0x0f, 0x6d, 0x64, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0xe2, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x64, 0x6b, 0x52, 0x65, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x6c, 0x75, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x75, 0x61, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x66, 0x69, 0x78, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x66,
	0x69, 0x78, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x66, 0x69,
	0x78, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x62, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x66, 0x69,
	0x78, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x31, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x6e, 0x6b, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x32, 0x12,
	0x13, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x33, 0x18, 0xaa, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x75, 0x6e, 0x6b, 0x33, 0x12, 0x13, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x34, 0x18, 0xf5, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x34, 0x12, 0x13, 0x0a, 0x04, 0x75, 0x6e, 0x6b,
	0x35, 0x18, 0xf1, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x35, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Gateserver_proto_rawDescOnce sync.Once
	file_Gateserver_proto_rawDescData = file_Gateserver_proto_rawDesc
)

func file_Gateserver_proto_rawDescGZIP() []byte {
	file_Gateserver_proto_rawDescOnce.Do(func() {
		file_Gateserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_Gateserver_proto_rawDescData)
	})
	return file_Gateserver_proto_rawDescData
}

var file_Gateserver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Gateserver_proto_goTypes = []interface{}{
	(*Gateserver)(nil), // 0: Gateserver
}
var file_Gateserver_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Gateserver_proto_init() }
func file_Gateserver_proto_init() {
	if File_Gateserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Gateserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gateserver); i {
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
			RawDescriptor: file_Gateserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Gateserver_proto_goTypes,
		DependencyIndexes: file_Gateserver_proto_depIdxs,
		MessageInfos:      file_Gateserver_proto_msgTypes,
	}.Build()
	File_Gateserver_proto = out.File
	file_Gateserver_proto_rawDesc = nil
	file_Gateserver_proto_goTypes = nil
	file_Gateserver_proto_depIdxs = nil
}
