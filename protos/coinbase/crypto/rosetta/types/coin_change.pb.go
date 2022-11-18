// (-- api-linter: core::0140::abbreviations=disabled
//     aip.dev/not-precedent: Matching the naming convention of rosetta --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: coinbase/crypto/rosetta/types/coin_change.proto

// The stable release for rosetta types

package types

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// CoinActions are different state changes that a Coin can undergo. It is assumed that a single
// Coin cannot be created or spent more than once.
type CoinChange_CoinAction int32

const (
	// COIN_ACTION_UNSPECIFIED indicating a CoinAction is not set
	CoinChange_COIN_ACTION_UNSPECIFIED CoinChange_CoinAction = 0
	// COIN_CREATED indicating a Coin was created.
	CoinChange_COIN_CREATED CoinChange_CoinAction = 1
	// COIN_SPENT indicating a Coin was spent.
	CoinChange_COIN_SPENT CoinChange_CoinAction = 2
)

// Enum value maps for CoinChange_CoinAction.
var (
	CoinChange_CoinAction_name = map[int32]string{
		0: "COIN_ACTION_UNSPECIFIED",
		1: "COIN_CREATED",
		2: "COIN_SPENT",
	}
	CoinChange_CoinAction_value = map[string]int32{
		"COIN_ACTION_UNSPECIFIED": 0,
		"COIN_CREATED":            1,
		"COIN_SPENT":              2,
	}
)

func (x CoinChange_CoinAction) Enum() *CoinChange_CoinAction {
	p := new(CoinChange_CoinAction)
	*p = x
	return p
}

func (x CoinChange_CoinAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoinChange_CoinAction) Descriptor() protoreflect.EnumDescriptor {
	return file_coinbase_crypto_rosetta_types_coin_change_proto_enumTypes[0].Descriptor()
}

func (CoinChange_CoinAction) Type() protoreflect.EnumType {
	return &file_coinbase_crypto_rosetta_types_coin_change_proto_enumTypes[0]
}

func (x CoinChange_CoinAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoinChange_CoinAction.Descriptor instead.
func (CoinChange_CoinAction) EnumDescriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_coin_change_proto_rawDescGZIP(), []int{0, 0}
}

// CoinChange is used to represent a change in state of a coin identified by a coin_identifier.
// This object is part of the Operation model and must be populated for UTXO-based blockchains.
// Coincidentally, this abstraction of UTXOs allows for supporting both account-based transfers and
// UTXO-based transfers on the same blockchain (when a transfer is account-based, don't populate
// this model).
type CoinChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CoinIdentifier uniquely identifies a Coin.
	CoinIdentifier *CoinIdentifier `protobuf:"bytes,1,opt,name=coin_identifier,json=coinIdentifier,proto3" json:"coin_identifier,omitempty"`
	// CoinActions are different state changes that a Coin can undergo
	CoinAction CoinChange_CoinAction `protobuf:"varint,2,opt,name=coin_action,json=coinAction,proto3,enum=coinbase.crypto.rosetta.types.CoinChange_CoinAction" json:"coin_action,omitempty"`
}

func (x *CoinChange) Reset() {
	*x = CoinChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_coin_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinChange) ProtoMessage() {}

func (x *CoinChange) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_coin_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinChange.ProtoReflect.Descriptor instead.
func (*CoinChange) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_coin_change_proto_rawDescGZIP(), []int{0}
}

func (x *CoinChange) GetCoinIdentifier() *CoinIdentifier {
	if x != nil {
		return x.CoinIdentifier
	}
	return nil
}

func (x *CoinChange) GetCoinAction() CoinChange_CoinAction {
	if x != nil {
		return x.CoinAction
	}
	return CoinChange_COIN_ACTION_UNSPECIFIED
}

// CoinIdentifier uniquely identifies a Coin.
type CoinIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier should be populated with a globally unique identifier of a Coin. In Bitcoin, this
	// identifier would be transaction_hash:index.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *CoinIdentifier) Reset() {
	*x = CoinIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_crypto_rosetta_types_coin_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinIdentifier) ProtoMessage() {}

func (x *CoinIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_crypto_rosetta_types_coin_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinIdentifier.ProtoReflect.Descriptor instead.
func (*CoinIdentifier) Descriptor() ([]byte, []int) {
	return file_coinbase_crypto_rosetta_types_coin_change_proto_rawDescGZIP(), []int{1}
}

func (x *CoinIdentifier) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

var File_coinbase_crypto_rosetta_types_coin_change_proto protoreflect.FileDescriptor

var file_coinbase_crypto_rosetta_types_coin_change_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x88, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x56, 0x0a, 0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72, 0x6f, 0x73, 0x65, 0x74,
	0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0e, 0x63, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x0b, 0x63, 0x6f, 0x69, 0x6e, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x72,
	0x6f, 0x73, 0x65, 0x74, 0x74, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x69, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b,
	0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x17,
	0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x49,
	0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43,
	0x4f, 0x49, 0x4e, 0x5f, 0x53, 0x50, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x22, 0x30, 0x0a, 0x0e, 0x43,
	0x6f, 0x69, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x47, 0x5a,
	0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x73, 0x65, 0x74, 0x74, 0x61,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_crypto_rosetta_types_coin_change_proto_rawDescOnce sync.Once
	file_coinbase_crypto_rosetta_types_coin_change_proto_rawDescData = file_coinbase_crypto_rosetta_types_coin_change_proto_rawDesc
)

func file_coinbase_crypto_rosetta_types_coin_change_proto_rawDescGZIP() []byte {
	file_coinbase_crypto_rosetta_types_coin_change_proto_rawDescOnce.Do(func() {
		file_coinbase_crypto_rosetta_types_coin_change_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_crypto_rosetta_types_coin_change_proto_rawDescData)
	})
	return file_coinbase_crypto_rosetta_types_coin_change_proto_rawDescData
}

var file_coinbase_crypto_rosetta_types_coin_change_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_coinbase_crypto_rosetta_types_coin_change_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_coinbase_crypto_rosetta_types_coin_change_proto_goTypes = []interface{}{
	(CoinChange_CoinAction)(0), // 0: coinbase.crypto.rosetta.types.CoinChange.CoinAction
	(*CoinChange)(nil),         // 1: coinbase.crypto.rosetta.types.CoinChange
	(*CoinIdentifier)(nil),     // 2: coinbase.crypto.rosetta.types.CoinIdentifier
}
var file_coinbase_crypto_rosetta_types_coin_change_proto_depIdxs = []int32{
	2, // 0: coinbase.crypto.rosetta.types.CoinChange.coin_identifier:type_name -> coinbase.crypto.rosetta.types.CoinIdentifier
	0, // 1: coinbase.crypto.rosetta.types.CoinChange.coin_action:type_name -> coinbase.crypto.rosetta.types.CoinChange.CoinAction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coinbase_crypto_rosetta_types_coin_change_proto_init() }
func file_coinbase_crypto_rosetta_types_coin_change_proto_init() {
	if File_coinbase_crypto_rosetta_types_coin_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_crypto_rosetta_types_coin_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinChange); i {
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
		file_coinbase_crypto_rosetta_types_coin_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinIdentifier); i {
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
			RawDescriptor: file_coinbase_crypto_rosetta_types_coin_change_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_crypto_rosetta_types_coin_change_proto_goTypes,
		DependencyIndexes: file_coinbase_crypto_rosetta_types_coin_change_proto_depIdxs,
		EnumInfos:         file_coinbase_crypto_rosetta_types_coin_change_proto_enumTypes,
		MessageInfos:      file_coinbase_crypto_rosetta_types_coin_change_proto_msgTypes,
	}.Build()
	File_coinbase_crypto_rosetta_types_coin_change_proto = out.File
	file_coinbase_crypto_rosetta_types_coin_change_proto_rawDesc = nil
	file_coinbase_crypto_rosetta_types_coin_change_proto_goTypes = nil
	file_coinbase_crypto_rosetta_types_coin_change_proto_depIdxs = nil
}
