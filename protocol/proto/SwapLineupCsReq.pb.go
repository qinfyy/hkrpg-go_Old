// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SwapLineupCsReq.proto

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

type SwapLineupCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcSlot         uint32          `protobuf:"varint,11,opt,name=src_slot,json=srcSlot,proto3" json:"src_slot,omitempty"`
	PlaneId         uint32          `protobuf:"varint,7,opt,name=plane_id,json=planeId,proto3" json:"plane_id,omitempty"`
	Index           uint32          `protobuf:"varint,15,opt,name=index,proto3" json:"index,omitempty"`
	ExtraLineupType ExtraLineupType `protobuf:"varint,14,opt,name=extra_lineup_type,json=extraLineupType,proto3,enum=ExtraLineupType" json:"extra_lineup_type,omitempty"`
	IsVirtual       bool            `protobuf:"varint,2,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual,omitempty"`
	DstSlot         uint32          `protobuf:"varint,12,opt,name=dst_slot,json=dstSlot,proto3" json:"dst_slot,omitempty"`
}

func (x *SwapLineupCsReq) Reset() {
	*x = SwapLineupCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SwapLineupCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapLineupCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapLineupCsReq) ProtoMessage() {}

func (x *SwapLineupCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SwapLineupCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapLineupCsReq.ProtoReflect.Descriptor instead.
func (*SwapLineupCsReq) Descriptor() ([]byte, []int) {
	return file_SwapLineupCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SwapLineupCsReq) GetSrcSlot() uint32 {
	if x != nil {
		return x.SrcSlot
	}
	return 0
}

func (x *SwapLineupCsReq) GetPlaneId() uint32 {
	if x != nil {
		return x.PlaneId
	}
	return 0
}

func (x *SwapLineupCsReq) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SwapLineupCsReq) GetExtraLineupType() ExtraLineupType {
	if x != nil {
		return x.ExtraLineupType
	}
	return ExtraLineupType_LINEUP_NONE
}

func (x *SwapLineupCsReq) GetIsVirtual() bool {
	if x != nil {
		return x.IsVirtual
	}
	return false
}

func (x *SwapLineupCsReq) GetDstSlot() uint32 {
	if x != nil {
		return x.DstSlot
	}
	return 0
}

var File_SwapLineupCsReq_proto protoreflect.FileDescriptor

var file_SwapLineupCsReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x77, 0x61, 0x70, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x43, 0x73, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69,
	0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5,
	0x01, 0x0a, 0x0f, 0x53, 0x77, 0x61, 0x70, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x72, 0x63, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x3c,
	0x0a, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x64,
	0x73, 0x74, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64,
	0x73, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SwapLineupCsReq_proto_rawDescOnce sync.Once
	file_SwapLineupCsReq_proto_rawDescData = file_SwapLineupCsReq_proto_rawDesc
)

func file_SwapLineupCsReq_proto_rawDescGZIP() []byte {
	file_SwapLineupCsReq_proto_rawDescOnce.Do(func() {
		file_SwapLineupCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SwapLineupCsReq_proto_rawDescData)
	})
	return file_SwapLineupCsReq_proto_rawDescData
}

var file_SwapLineupCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SwapLineupCsReq_proto_goTypes = []interface{}{
	(*SwapLineupCsReq)(nil), // 0: SwapLineupCsReq
	(ExtraLineupType)(0),    // 1: ExtraLineupType
}
var file_SwapLineupCsReq_proto_depIdxs = []int32{
	1, // 0: SwapLineupCsReq.extra_lineup_type:type_name -> ExtraLineupType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SwapLineupCsReq_proto_init() }
func file_SwapLineupCsReq_proto_init() {
	if File_SwapLineupCsReq_proto != nil {
		return
	}
	file_ExtraLineupType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SwapLineupCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapLineupCsReq); i {
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
			RawDescriptor: file_SwapLineupCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SwapLineupCsReq_proto_goTypes,
		DependencyIndexes: file_SwapLineupCsReq_proto_depIdxs,
		MessageInfos:      file_SwapLineupCsReq_proto_msgTypes,
	}.Build()
	File_SwapLineupCsReq_proto = out.File
	file_SwapLineupCsReq_proto_rawDesc = nil
	file_SwapLineupCsReq_proto_goTypes = nil
	file_SwapLineupCsReq_proto_depIdxs = nil
}
