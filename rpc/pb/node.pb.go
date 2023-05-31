// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.9
// source: node.proto

package pb

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

type NodeInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NodeInfoRequest) Reset() {
	*x = NodeInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfoRequest) ProtoMessage() {}

func (x *NodeInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfoRequest.ProtoReflect.Descriptor instead.
func (*NodeInfoRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

type NodeInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId       string `protobuf:"bytes,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	Endpoint      string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	CoinType      uint64 `protobuf:"varint,3,opt,name=coinType,proto3" json:"coinType,omitempty"`
	PrefixAddress string `protobuf:"bytes,4,opt,name=prefixAddress,proto3" json:"prefixAddress,omitempty"`
	TokenSymbol   string `protobuf:"bytes,5,opt,name=tokenSymbol,proto3" json:"tokenSymbol,omitempty"`
}

func (x *NodeInfoResponse) Reset() {
	*x = NodeInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfoResponse) ProtoMessage() {}

func (x *NodeInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfoResponse.ProtoReflect.Descriptor instead.
func (*NodeInfoResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *NodeInfoResponse) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *NodeInfoResponse) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *NodeInfoResponse) GetCoinType() uint64 {
	if x != nil {
		return x.CoinType
	}
	return 0
}

func (x *NodeInfoResponse) GetPrefixAddress() string {
	if x != nil {
		return x.PrefixAddress
	}
	return ""
}

func (x *NodeInfoResponse) GetTokenSymbol() string {
	if x != nil {
		return x.TokenSymbol
	}
	return ""
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xac, 0x01, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x32, 0x4a, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x15, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x32, 0x35, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67,
	0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_node_proto_goTypes = []interface{}{
	(*NodeInfoRequest)(nil),  // 0: node.NodeInfoRequest
	(*NodeInfoResponse)(nil), // 1: node.NodeInfoResponse
}
var file_node_proto_depIdxs = []int32{
	0, // 0: node.NodeService.NodeInfo:input_type -> node.NodeInfoRequest
	1, // 1: node.NodeService.NodeInfo:output_type -> node.NodeInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfoRequest); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfoResponse); i {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}
