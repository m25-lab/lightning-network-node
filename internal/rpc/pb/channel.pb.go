// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: channel.proto

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

type OpenChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountA       string `protobuf:"bytes,1,opt,name=accountA,proto3" json:"accountA,omitempty"`
	AccountB       string `protobuf:"bytes,2,opt,name=accountB,proto3" json:"accountB,omitempty"`
	AmountA        int64  `protobuf:"varint,3,opt,name=amountA,proto3" json:"amountA,omitempty"`
	AmountB        int64  `protobuf:"varint,4,opt,name=amountB,proto3" json:"amountB,omitempty"`
	AccountChannel string `protobuf:"bytes,5,opt,name=accountChannel,proto3" json:"accountChannel,omitempty"`
	Sequence       int32  `protobuf:"varint,6,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *OpenChannelRequest) Reset() {
	*x = OpenChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenChannelRequest) ProtoMessage() {}

func (x *OpenChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenChannelRequest.ProtoReflect.Descriptor instead.
func (*OpenChannelRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{0}
}

func (x *OpenChannelRequest) GetAccountA() string {
	if x != nil {
		return x.AccountA
	}
	return ""
}

func (x *OpenChannelRequest) GetAccountB() string {
	if x != nil {
		return x.AccountB
	}
	return ""
}

func (x *OpenChannelRequest) GetAmountA() int64 {
	if x != nil {
		return x.AmountA
	}
	return 0
}

func (x *OpenChannelRequest) GetAmountB() int64 {
	if x != nil {
		return x.AmountB
	}
	return 0
}

func (x *OpenChannelRequest) GetAccountChannel() string {
	if x != nil {
		return x.AccountChannel
	}
	return ""
}

func (x *OpenChannelRequest) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type OpenChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *OpenChannelResponse) Reset() {
	*x = OpenChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenChannelResponse) ProtoMessage() {}

func (x *OpenChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenChannelResponse.ProtoReflect.Descriptor instead.
func (*OpenChannelResponse) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{1}
}

func (x *OpenChannelResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type CreateCommitmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountFrom string `protobuf:"bytes,1,opt,name=accountFrom,proto3" json:"accountFrom,omitempty"`
	AmountA     int64  `protobuf:"varint,2,opt,name=amountA,proto3" json:"amountA,omitempty"`
	ToAHashlock string `protobuf:"bytes,3,opt,name=toAHashlock,proto3" json:"toAHashlock,omitempty"`
	HashCode    string `protobuf:"bytes,4,opt,name=hashCode,proto3" json:"hashCode,omitempty"`
	ToBHashlock string `protobuf:"bytes,5,opt,name=toBHashlock,proto3" json:"toBHashlock,omitempty"`
	BlockHeight string `protobuf:"bytes,6,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	CoinLock    uint64 `protobuf:"varint,7,opt,name=coinLock,proto3" json:"coinLock,omitempty"`
}

func (x *CreateCommitmentRequest) Reset() {
	*x = CreateCommitmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommitmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommitmentRequest) ProtoMessage() {}

func (x *CreateCommitmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommitmentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommitmentRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCommitmentRequest) GetAccountFrom() string {
	if x != nil {
		return x.AccountFrom
	}
	return ""
}

func (x *CreateCommitmentRequest) GetAmountA() int64 {
	if x != nil {
		return x.AmountA
	}
	return 0
}

func (x *CreateCommitmentRequest) GetToAHashlock() string {
	if x != nil {
		return x.ToAHashlock
	}
	return ""
}

func (x *CreateCommitmentRequest) GetHashCode() string {
	if x != nil {
		return x.HashCode
	}
	return ""
}

func (x *CreateCommitmentRequest) GetToBHashlock() string {
	if x != nil {
		return x.ToBHashlock
	}
	return ""
}

func (x *CreateCommitmentRequest) GetBlockHeight() string {
	if x != nil {
		return x.BlockHeight
	}
	return ""
}

func (x *CreateCommitmentRequest) GetCoinLock() uint64 {
	if x != nil {
		return x.CoinLock
	}
	return 0
}

type CreateCommitmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateCommitmentResponse) Reset() {
	*x = CreateCommitmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommitmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommitmentResponse) ProtoMessage() {}

func (x *CreateCommitmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommitmentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommitmentResponse) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCommitmentResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type WithdrawHashlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountTo string `protobuf:"bytes,1,opt,name=accountTo,proto3" json:"accountTo,omitempty"`
	Index     string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Secret    string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *WithdrawHashlockRequest) Reset() {
	*x = WithdrawHashlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawHashlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawHashlockRequest) ProtoMessage() {}

func (x *WithdrawHashlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawHashlockRequest.ProtoReflect.Descriptor instead.
func (*WithdrawHashlockRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{4}
}

func (x *WithdrawHashlockRequest) GetAccountTo() string {
	if x != nil {
		return x.AccountTo
	}
	return ""
}

func (x *WithdrawHashlockRequest) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *WithdrawHashlockRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type WithdrawHashlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *WithdrawHashlockResponse) Reset() {
	*x = WithdrawHashlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawHashlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawHashlockResponse) ProtoMessage() {}

func (x *WithdrawHashlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawHashlockResponse.ProtoReflect.Descriptor instead.
func (*WithdrawHashlockResponse) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{5}
}

func (x *WithdrawHashlockResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type WithdrawTimelockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountTo string `protobuf:"bytes,1,opt,name=accountTo,proto3" json:"accountTo,omitempty"`
	Index     string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *WithdrawTimelockRequest) Reset() {
	*x = WithdrawTimelockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawTimelockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawTimelockRequest) ProtoMessage() {}

func (x *WithdrawTimelockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawTimelockRequest.ProtoReflect.Descriptor instead.
func (*WithdrawTimelockRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{6}
}

func (x *WithdrawTimelockRequest) GetAccountTo() string {
	if x != nil {
		return x.AccountTo
	}
	return ""
}

func (x *WithdrawTimelockRequest) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

type WithdrawTimelockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *WithdrawTimelockResponse) Reset() {
	*x = WithdrawTimelockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawTimelockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawTimelockResponse) ProtoMessage() {}

func (x *WithdrawTimelockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawTimelockResponse.ProtoReflect.Descriptor instead.
func (*WithdrawTimelockResponse) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{7}
}

func (x *WithdrawTimelockResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_channel_proto protoreflect.FileDescriptor

var file_channel_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc4, 0x01, 0x0a, 0x12, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x41, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x31, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf3, 0x01, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x41, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x41, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x41, 0x48, 0x61, 0x73, 0x68, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x42, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x42, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x22,
	0x36, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x0a, 0x17, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x36,
	0x0a, 0x18, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x0a, 0x17, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x36, 0x0a, 0x18, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xad, 0x02,
	0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3a, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x13, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x69,
	0x6d, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x32, 0x35, 0x2d,
	0x6c, 0x61, 0x62, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_channel_proto_rawDescOnce sync.Once
	file_channel_proto_rawDescData = file_channel_proto_rawDesc
)

func file_channel_proto_rawDescGZIP() []byte {
	file_channel_proto_rawDescOnce.Do(func() {
		file_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_channel_proto_rawDescData)
	})
	return file_channel_proto_rawDescData
}

var file_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_channel_proto_goTypes = []interface{}{
	(*OpenChannelRequest)(nil),       // 0: OpenChannelRequest
	(*OpenChannelResponse)(nil),      // 1: OpenChannelResponse
	(*CreateCommitmentRequest)(nil),  // 2: CreateCommitmentRequest
	(*CreateCommitmentResponse)(nil), // 3: CreateCommitmentResponse
	(*WithdrawHashlockRequest)(nil),  // 4: WithdrawHashlockRequest
	(*WithdrawHashlockResponse)(nil), // 5: WithdrawHashlockResponse
	(*WithdrawTimelockRequest)(nil),  // 6: WithdrawTimelockRequest
	(*WithdrawTimelockResponse)(nil), // 7: WithdrawTimelockResponse
}
var file_channel_proto_depIdxs = []int32{
	0, // 0: ChannelService.OpenChannel:input_type -> OpenChannelRequest
	2, // 1: ChannelService.CreateCommitment:input_type -> CreateCommitmentRequest
	4, // 2: ChannelService.WithdrawHashlock:input_type -> WithdrawHashlockRequest
	6, // 3: ChannelService.WithdrawTimelock:input_type -> WithdrawTimelockRequest
	1, // 4: ChannelService.OpenChannel:output_type -> OpenChannelResponse
	3, // 5: ChannelService.CreateCommitment:output_type -> CreateCommitmentResponse
	5, // 6: ChannelService.WithdrawHashlock:output_type -> WithdrawHashlockResponse
	7, // 7: ChannelService.WithdrawTimelock:output_type -> WithdrawTimelockResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_channel_proto_init() }
func file_channel_proto_init() {
	if File_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenChannelRequest); i {
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
		file_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenChannelResponse); i {
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
		file_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommitmentRequest); i {
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
		file_channel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommitmentResponse); i {
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
		file_channel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawHashlockRequest); i {
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
		file_channel_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawHashlockResponse); i {
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
		file_channel_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawTimelockRequest); i {
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
		file_channel_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawTimelockResponse); i {
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
			RawDescriptor: file_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_channel_proto_goTypes,
		DependencyIndexes: file_channel_proto_depIdxs,
		MessageInfos:      file_channel_proto_msgTypes,
	}.Build()
	File_channel_proto = out.File
	file_channel_proto_rawDesc = nil
	file_channel_proto_goTypes = nil
	file_channel_proto_depIdxs = nil
}