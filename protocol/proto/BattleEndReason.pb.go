// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BattleEndReason.proto

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

type BattleEndReason int32

const (
	BattleEndReason_BATTLE_END_REASON_NONE       BattleEndReason = 0
	BattleEndReason_BATTLE_END_REASON_ALL_DIE    BattleEndReason = 1
	BattleEndReason_BATTLE_END_REASON_TURN_LIMIT BattleEndReason = 2
)

// Enum value maps for BattleEndReason.
var (
	BattleEndReason_name = map[int32]string{
		0: "BATTLE_END_REASON_NONE",
		1: "BATTLE_END_REASON_ALL_DIE",
		2: "BATTLE_END_REASON_TURN_LIMIT",
	}
	BattleEndReason_value = map[string]int32{
		"BATTLE_END_REASON_NONE":       0,
		"BATTLE_END_REASON_ALL_DIE":    1,
		"BATTLE_END_REASON_TURN_LIMIT": 2,
	}
)

func (x BattleEndReason) Enum() *BattleEndReason {
	p := new(BattleEndReason)
	*p = x
	return p
}

func (x BattleEndReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BattleEndReason) Descriptor() protoreflect.EnumDescriptor {
	return file_BattleEndReason_proto_enumTypes[0].Descriptor()
}

func (BattleEndReason) Type() protoreflect.EnumType {
	return &file_BattleEndReason_proto_enumTypes[0]
}

func (x BattleEndReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BattleEndReason.Descriptor instead.
func (BattleEndReason) EnumDescriptor() ([]byte, []int) {
	return file_BattleEndReason_proto_rawDescGZIP(), []int{0}
}

var File_BattleEndReason_proto protoreflect.FileDescriptor

var file_BattleEndReason_proto_rawDesc = []byte{
	0x0a, 0x15, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x6e, 0x0a, 0x0f, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x42, 0x41,
	0x54, 0x54, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45,
	0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x5f,
	0x44, 0x49, 0x45, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f,
	0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x55, 0x52, 0x4e, 0x5f,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattleEndReason_proto_rawDescOnce sync.Once
	file_BattleEndReason_proto_rawDescData = file_BattleEndReason_proto_rawDesc
)

func file_BattleEndReason_proto_rawDescGZIP() []byte {
	file_BattleEndReason_proto_rawDescOnce.Do(func() {
		file_BattleEndReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattleEndReason_proto_rawDescData)
	})
	return file_BattleEndReason_proto_rawDescData
}

var file_BattleEndReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BattleEndReason_proto_goTypes = []interface{}{
	(BattleEndReason)(0), // 0: BattleEndReason
}
var file_BattleEndReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BattleEndReason_proto_init() }
func file_BattleEndReason_proto_init() {
	if File_BattleEndReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_BattleEndReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattleEndReason_proto_goTypes,
		DependencyIndexes: file_BattleEndReason_proto_depIdxs,
		EnumInfos:         file_BattleEndReason_proto_enumTypes,
	}.Build()
	File_BattleEndReason_proto = out.File
	file_BattleEndReason_proto_rawDesc = nil
	file_BattleEndReason_proto_goTypes = nil
	file_BattleEndReason_proto_depIdxs = nil
}
