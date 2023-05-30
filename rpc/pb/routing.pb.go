// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.9
// source: routing.proto

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

type RoutingErrorCode int32

const (
	RoutingErrorCode_SOME_THING_WENT_WRONG               RoutingErrorCode = 0
	RoutingErrorCode_RREQ_EXISTED                        RoutingErrorCode = 1
	RoutingErrorCode_RREP_EXISTED                        RoutingErrorCode = 2
	RoutingErrorCode_MORE_THAN_ONE_RREQ_EXISTED          RoutingErrorCode = 3
	RoutingErrorCode_FORWARD_RREP_ERROR                  RoutingErrorCode = 4
	RoutingErrorCode_NOT_FOUND_NEIGHBOR_NODE             RoutingErrorCode = 5
	RoutingErrorCode_PARAM_INVALID                       RoutingErrorCode = 6
	RoutingErrorCode_WRONG_NODE                          RoutingErrorCode = 7
	RoutingErrorCode_VALIDATE_INVOICE_SECRET             RoutingErrorCode = 8
	RoutingErrorCode_INSERT_SECRET                       RoutingErrorCode = 9
	RoutingErrorCode_DESTINATION_ADDRESS_FIND_BY_ADDRESS RoutingErrorCode = 10
	RoutingErrorCode_OK                                  RoutingErrorCode = 999
)

// Enum value maps for RoutingErrorCode.
var (
	RoutingErrorCode_name = map[int32]string{
		0:   "SOME_THING_WENT_WRONG",
		1:   "RREQ_EXISTED",
		2:   "RREP_EXISTED",
		3:   "MORE_THAN_ONE_RREQ_EXISTED",
		4:   "FORWARD_RREP_ERROR",
		5:   "NOT_FOUND_NEIGHBOR_NODE",
		6:   "PARAM_INVALID",
		7:   "WRONG_NODE",
		8:   "VALIDATE_INVOICE_SECRET",
		9:   "INSERT_SECRET",
		10:  "DESTINATION_ADDRESS_FIND_BY_ADDRESS",
		999: "OK",
	}
	RoutingErrorCode_value = map[string]int32{
		"SOME_THING_WENT_WRONG":               0,
		"RREQ_EXISTED":                        1,
		"RREP_EXISTED":                        2,
		"MORE_THAN_ONE_RREQ_EXISTED":          3,
		"FORWARD_RREP_ERROR":                  4,
		"NOT_FOUND_NEIGHBOR_NODE":             5,
		"PARAM_INVALID":                       6,
		"WRONG_NODE":                          7,
		"VALIDATE_INVOICE_SECRET":             8,
		"INSERT_SECRET":                       9,
		"DESTINATION_ADDRESS_FIND_BY_ADDRESS": 10,
		"OK":                                  999,
	}
)

func (x RoutingErrorCode) Enum() *RoutingErrorCode {
	p := new(RoutingErrorCode)
	*p = x
	return p
}

func (x RoutingErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoutingErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_routing_proto_enumTypes[0].Descriptor()
}

func (RoutingErrorCode) Type() protoreflect.EnumType {
	return &file_routing_proto_enumTypes[0]
}

func (x RoutingErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoutingErrorCode.Descriptor instead.
func (RoutingErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{0}
}

type RoutingBaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode RoutingErrorCode `protobuf:"varint,1,opt,name=ErrorCode,proto3,enum=channel.RoutingErrorCode" json:"ErrorCode,omitempty"`
}

func (x *RoutingBaseResponse) Reset() {
	*x = RoutingBaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingBaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingBaseResponse) ProtoMessage() {}

func (x *RoutingBaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingBaseResponse.ProtoReflect.Descriptor instead.
func (*RoutingBaseResponse) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{0}
}

func (x *RoutingBaseResponse) GetErrorCode() RoutingErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return RoutingErrorCode_SOME_THING_WENT_WRONG
}

type RREQRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceAddress      string `protobuf:"bytes,1,opt,name=SourceAddress,proto3" json:"SourceAddress,omitempty"`
	DestinationAddress string `protobuf:"bytes,2,opt,name=DestinationAddress,proto3" json:"DestinationAddress,omitempty"`
	BroadcastID        string `protobuf:"bytes,3,opt,name=BroadcastID,proto3" json:"BroadcastID,omitempty"` // hash invoice
	FromAddress        string `protobuf:"bytes,4,opt,name=FromAddress,proto3" json:"FromAddress,omitempty"`
	ToAddress          string `protobuf:"bytes,5,opt,name=ToAddress,proto3" json:"ToAddress,omitempty"`
}

func (x *RREQRequest) Reset() {
	*x = RREQRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RREQRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RREQRequest) ProtoMessage() {}

func (x *RREQRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RREQRequest.ProtoReflect.Descriptor instead.
func (*RREQRequest) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{1}
}

func (x *RREQRequest) GetSourceAddress() string {
	if x != nil {
		return x.SourceAddress
	}
	return ""
}

func (x *RREQRequest) GetDestinationAddress() string {
	if x != nil {
		return x.DestinationAddress
	}
	return ""
}

func (x *RREQRequest) GetBroadcastID() string {
	if x != nil {
		return x.BroadcastID
	}
	return ""
}

func (x *RREQRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *RREQRequest) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

type RREPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceAddress      string `protobuf:"bytes,1,opt,name=SourceAddress,proto3" json:"SourceAddress,omitempty"`
	DestinationAddress string `protobuf:"bytes,2,opt,name=DestinationAddress,proto3" json:"DestinationAddress,omitempty"`
	BroadcastID        string `protobuf:"bytes,3,opt,name=BroadcastID,proto3" json:"BroadcastID,omitempty"` // hash invoice
	FromAddress        string `protobuf:"bytes,4,opt,name=FromAddress,proto3" json:"FromAddress,omitempty"`
	ToAddress          string `protobuf:"bytes,5,opt,name=ToAddress,proto3" json:"ToAddress,omitempty"`
}

func (x *RREPRequest) Reset() {
	*x = RREPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RREPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RREPRequest) ProtoMessage() {}

func (x *RREPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RREPRequest.ProtoReflect.Descriptor instead.
func (*RREPRequest) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{2}
}

func (x *RREPRequest) GetSourceAddress() string {
	if x != nil {
		return x.SourceAddress
	}
	return ""
}

func (x *RREPRequest) GetDestinationAddress() string {
	if x != nil {
		return x.DestinationAddress
	}
	return ""
}

func (x *RREPRequest) GetBroadcastID() string {
	if x != nil {
		return x.BroadcastID
	}
	return ""
}

func (x *RREPRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *RREPRequest) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

type IREQMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	From   string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To     string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *IREQMessage) Reset() {
	*x = IREQMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IREQMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IREQMessage) ProtoMessage() {}

func (x *IREQMessage) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IREQMessage.ProtoReflect.Descriptor instead.
func (*IREQMessage) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{3}
}

func (x *IREQMessage) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *IREQMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *IREQMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type IREPMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Hash      string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Amount    int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	ErrorCode string `protobuf:"bytes,5,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *IREPMessage) Reset() {
	*x = IREPMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IREPMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IREPMessage) ProtoMessage() {}

func (x *IREPMessage) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IREPMessage.ProtoReflect.Descriptor instead.
func (*IREPMessage) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{4}
}

func (x *IREPMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *IREPMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *IREPMessage) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *IREPMessage) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *IREPMessage) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

type InvoiceSecretMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashcode string `protobuf:"bytes,1,opt,name=hashcode,proto3" json:"hashcode,omitempty"`
	Secret   string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	Dest     string `protobuf:"bytes,3,opt,name=dest,proto3" json:"dest,omitempty"`
}

func (x *InvoiceSecretMessage) Reset() {
	*x = InvoiceSecretMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceSecretMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceSecretMessage) ProtoMessage() {}

func (x *InvoiceSecretMessage) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceSecretMessage.ProtoReflect.Descriptor instead.
func (*InvoiceSecretMessage) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{5}
}

func (x *InvoiceSecretMessage) GetHashcode() string {
	if x != nil {
		return x.Hashcode
	}
	return ""
}

func (x *InvoiceSecretMessage) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *InvoiceSecretMessage) GetDest() string {
	if x != nil {
		return x.Dest
	}
	return ""
}

type FwdMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response   string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	PartnerSig string `protobuf:"bytes,2,opt,name=partner_sig,json=partnerSig,proto3" json:"partner_sig,omitempty"`
	ErrorCode  string `protobuf:"bytes,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *FwdMessageResponse) Reset() {
	*x = FwdMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FwdMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FwdMessageResponse) ProtoMessage() {}

func (x *FwdMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FwdMessageResponse.ProtoReflect.Descriptor instead.
func (*FwdMessageResponse) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{6}
}

func (x *FwdMessageResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *FwdMessageResponse) GetPartnerSig() string {
	if x != nil {
		return x.PartnerSig
	}
	return ""
}

func (x *FwdMessageResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

type FwdMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action       string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Data         string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	From         string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To           string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	HashcodeDest string `protobuf:"bytes,5,opt,name=hashcode_dest,json=hashcodeDest,proto3" json:"hashcode_dest,omitempty"`
	Dest         string `protobuf:"bytes,6,opt,name=dest,proto3" json:"dest,omitempty"`
	Sig          string `protobuf:"bytes,7,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *FwdMessage) Reset() {
	*x = FwdMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FwdMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FwdMessage) ProtoMessage() {}

func (x *FwdMessage) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FwdMessage.ProtoReflect.Descriptor instead.
func (*FwdMessage) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{7}
}

func (x *FwdMessage) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *FwdMessage) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *FwdMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *FwdMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *FwdMessage) GetHashcodeDest() string {
	if x != nil {
		return x.HashcodeDest
	}
	return ""
}

func (x *FwdMessage) GetDest() string {
	if x != nil {
		return x.Dest
	}
	return ""
}

func (x *FwdMessage) GetSig() string {
	if x != nil {
		return x.Sig
	}
	return ""
}

var File_routing_proto protoreflect.FileDescriptor

var file_routing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x4e, 0x0a, 0x13, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x52, 0x52, 0x45,
	0x51, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e,
	0x0a, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x44,
	0x12, 0x20, 0x0a, 0x0b, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x52, 0x52, 0x45, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x72, 0x6f, 0x6d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46,
	0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54,
	0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x49, 0x0a, 0x0b, 0x49, 0x52, 0x45, 0x51,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x0b, 0x49, 0x52, 0x45, 0x50, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x5e, 0x0a, 0x14, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73,
	0x68, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x73,
	0x68, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x74, 0x22, 0x70, 0x0a, 0x12, 0x46, 0x77, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x53, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x0a, 0x46, 0x77, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x64,
	0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x68, 0x63,
	0x6f, 0x64, 0x65, 0x44, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x67, 0x2a, 0xab, 0x02,
	0x0a, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x4f, 0x4d, 0x45, 0x5f, 0x54, 0x48, 0x49, 0x4e, 0x47,
	0x5f, 0x57, 0x45, 0x4e, 0x54, 0x5f, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x52, 0x52, 0x45, 0x51, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x52, 0x52, 0x45, 0x50, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x4f, 0x52, 0x45, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x4f,
	0x4e, 0x45, 0x5f, 0x52, 0x52, 0x45, 0x51, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x52, 0x52, 0x45,
	0x50, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x4e, 0x45, 0x49, 0x47, 0x48, 0x42, 0x4f, 0x52, 0x5f,
	0x4e, 0x4f, 0x44, 0x45, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x52, 0x4f,
	0x4e, 0x47, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x07, 0x12, 0x1b, 0x0a, 0x17, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x45,
	0x43, 0x52, 0x45, 0x54, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54,
	0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x10, 0x09, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x45, 0x53,
	0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53,
	0x5f, 0x46, 0x49, 0x4e, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53,
	0x10, 0x0a, 0x12, 0x07, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0xe7, 0x07, 0x32, 0xe6, 0x02, 0x0a, 0x0e,
	0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x04, 0x52, 0x52, 0x45, 0x51, 0x12, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x52, 0x52, 0x45, 0x51, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x04,
	0x52, 0x52, 0x45, 0x50, 0x12, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x52,
	0x52, 0x45, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x14, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x49, 0x52, 0x45, 0x51, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x49, 0x52, 0x45,
	0x50, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x46, 0x77, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x13, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x46, 0x77, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x46, 0x77, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x32, 0x35, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routing_proto_rawDescOnce sync.Once
	file_routing_proto_rawDescData = file_routing_proto_rawDesc
)

func file_routing_proto_rawDescGZIP() []byte {
	file_routing_proto_rawDescOnce.Do(func() {
		file_routing_proto_rawDescData = protoimpl.X.CompressGZIP(file_routing_proto_rawDescData)
	})
	return file_routing_proto_rawDescData
}

var file_routing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_routing_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_routing_proto_goTypes = []interface{}{
	(RoutingErrorCode)(0),        // 0: channel.RoutingErrorCode
	(*RoutingBaseResponse)(nil),  // 1: channel.RoutingBaseResponse
	(*RREQRequest)(nil),          // 2: channel.RREQRequest
	(*RREPRequest)(nil),          // 3: channel.RREPRequest
	(*IREQMessage)(nil),          // 4: channel.IREQMessage
	(*IREPMessage)(nil),          // 5: channel.IREPMessage
	(*InvoiceSecretMessage)(nil), // 6: channel.InvoiceSecretMessage
	(*FwdMessageResponse)(nil),   // 7: channel.FwdMessageResponse
	(*FwdMessage)(nil),           // 8: channel.FwdMessage
}
var file_routing_proto_depIdxs = []int32{
	0, // 0: channel.RoutingBaseResponse.ErrorCode:type_name -> channel.RoutingErrorCode
	2, // 1: channel.RoutingService.RREQ:input_type -> channel.RREQRequest
	3, // 2: channel.RoutingService.RREP:input_type -> channel.RREPRequest
	4, // 3: channel.RoutingService.RequestInvoice:input_type -> channel.IREQMessage
	8, // 4: channel.RoutingService.ProcessFwdMessage:input_type -> channel.FwdMessage
	6, // 5: channel.RoutingService.ProcessInvoiceSecret:input_type -> channel.InvoiceSecretMessage
	1, // 6: channel.RoutingService.RREQ:output_type -> channel.RoutingBaseResponse
	1, // 7: channel.RoutingService.RREP:output_type -> channel.RoutingBaseResponse
	5, // 8: channel.RoutingService.RequestInvoice:output_type -> channel.IREPMessage
	7, // 9: channel.RoutingService.ProcessFwdMessage:output_type -> channel.FwdMessageResponse
	1, // 10: channel.RoutingService.ProcessInvoiceSecret:output_type -> channel.RoutingBaseResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_routing_proto_init() }
func file_routing_proto_init() {
	if File_routing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingBaseResponse); i {
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
		file_routing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RREQRequest); i {
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
		file_routing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RREPRequest); i {
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
		file_routing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IREQMessage); i {
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
		file_routing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IREPMessage); i {
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
		file_routing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoiceSecretMessage); i {
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
		file_routing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FwdMessageResponse); i {
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
		file_routing_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FwdMessage); i {
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
			RawDescriptor: file_routing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routing_proto_goTypes,
		DependencyIndexes: file_routing_proto_depIdxs,
		EnumInfos:         file_routing_proto_enumTypes,
		MessageInfos:      file_routing_proto_msgTypes,
	}.Build()
	File_routing_proto = out.File
	file_routing_proto_rawDesc = nil
	file_routing_proto_goTypes = nil
	file_routing_proto_depIdxs = nil
}
