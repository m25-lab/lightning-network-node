// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: message.proto

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

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId       string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	ChannelId       string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Action          string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Data            string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	From            string `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	To              string `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
	ReliedMessageId string `protobuf:"bytes,7,opt,name=relied_message_id,json=reliedMessageId,proto3" json:"relied_message_id,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *SendMessageRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SendMessageRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SendMessageRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *SendMessageRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendMessageRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendMessageRequest) GetReliedMessageId() string {
	if x != nil {
		return x.ReliedMessageId
	}
	return ""
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response  string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	ErrorCode string `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *SendMessageResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

type SendSecretMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MySecret        string `protobuf:"bytes,1,opt,name=my_secret,json=mySecret,proto3" json:"my_secret,omitempty"`
	MyHashcode      string `protobuf:"bytes,2,opt,name=my_hashcode,json=myHashcode,proto3" json:"my_hashcode,omitempty"`
	PartnerHashcode string `protobuf:"bytes,3,opt,name=partner_hashcode,json=partnerHashcode,proto3" json:"partner_hashcode,omitempty"`
}

func (x *SendSecretMessage) Reset() {
	*x = SendSecretMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSecretMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSecretMessage) ProtoMessage() {}

func (x *SendSecretMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSecretMessage.ProtoReflect.Descriptor instead.
func (*SendSecretMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *SendSecretMessage) GetMySecret() string {
	if x != nil {
		return x.MySecret
	}
	return ""
}

func (x *SendSecretMessage) GetMyHashcode() string {
	if x != nil {
		return x.MyHashcode
	}
	return ""
}

func (x *SendSecretMessage) GetPartnerHashcode() string {
	if x != nil {
		return x.PartnerHashcode
	}
	return ""
}

type SendSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret    string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	ErrorCode string `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *SendSecretResponse) Reset() {
	*x = SendSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSecretResponse) ProtoMessage() {}

func (x *SendSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSecretResponse.ProtoReflect.Descriptor instead.
func (*SendSecretResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *SendSecretResponse) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *SendSecretResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

type BroadcastNotiMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To       string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	HashCode string `protobuf:"bytes,3,opt,name=hash_code,json=hashCode,proto3" json:"hash_code,omitempty"`
	Index    string `protobuf:"bytes,4,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *BroadcastNotiMessage) Reset() {
	*x = BroadcastNotiMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastNotiMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastNotiMessage) ProtoMessage() {}

func (x *BroadcastNotiMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastNotiMessage.ProtoReflect.Descriptor instead.
func (*BroadcastNotiMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *BroadcastNotiMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *BroadcastNotiMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *BroadcastNotiMessage) GetHashCode() string {
	if x != nil {
		return x.HashCode
	}
	return ""
}

func (x *BroadcastNotiMessage) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

type BroadcastNotiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response  string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	ErrorCode string `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *BroadcastNotiResponse) Reset() {
	*x = BroadcastNotiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastNotiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastNotiResponse) ProtoMessage() {}

func (x *BroadcastNotiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastNotiResponse.ProtoReflect.Descriptor instead.
func (*BroadcastNotiResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *BroadcastNotiResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *BroadcastNotiResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

type CloseChannelMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Sig       string `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
	ChannelId string `protobuf:"bytes,4,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	CoinA     int64  `protobuf:"varint,5,opt,name=coin_a,json=coinA,proto3" json:"coin_a,omitempty"`
	CoinB     int64  `protobuf:"varint,6,opt,name=coin_b,json=coinB,proto3" json:"coin_b,omitempty"`
}

func (x *CloseChannelMessage) Reset() {
	*x = CloseChannelMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseChannelMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseChannelMessage) ProtoMessage() {}

func (x *CloseChannelMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseChannelMessage.ProtoReflect.Descriptor instead.
func (*CloseChannelMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *CloseChannelMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CloseChannelMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *CloseChannelMessage) GetSig() string {
	if x != nil {
		return x.Sig
	}
	return ""
}

func (x *CloseChannelMessage) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *CloseChannelMessage) GetCoinA() int64 {
	if x != nil {
		return x.CoinA
	}
	return 0
}

func (x *CloseChannelMessage) GetCoinB() int64 {
	if x != nil {
		return x.CoinB
	}
	return 0
}

type CloseChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response  string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	ErrorCode string `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *CloseChannelResponse) Reset() {
	*x = CloseChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseChannelResponse) ProtoMessage() {}

func (x *CloseChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseChannelResponse.ProtoReflect.Descriptor instead.
func (*CloseChannelResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *CloseChannelResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *CloseChannelResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xce, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2a, 0x0a,
	0x11, 0x72, 0x65, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x69, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x13, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x7c, 0x0a, 0x11, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x79, 0x48, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x48, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x4b, 0x0a, 0x12, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x6d, 0x0a, 0x14, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x52, 0x0a, 0x15, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x13, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x69, 0x6e, 0x41, 0x12, 0x15, 0x0a,
	0x06, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x69, 0x6e, 0x42, 0x22, 0x51, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xc0, 0x02, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x0d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x12, 0x1d,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1e, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x0c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1c, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1d, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x32, 0x35, 0x2d, 0x6c, 0x61, 0x62,
	0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_message_proto_goTypes = []interface{}{
	(*SendMessageRequest)(nil),    // 0: channel.SendMessageRequest
	(*SendMessageResponse)(nil),   // 1: channel.SendMessageResponse
	(*SendSecretMessage)(nil),     // 2: channel.SendSecretMessage
	(*SendSecretResponse)(nil),    // 3: channel.SendSecretResponse
	(*BroadcastNotiMessage)(nil),  // 4: channel.BroadcastNotiMessage
	(*BroadcastNotiResponse)(nil), // 5: channel.BroadcastNotiResponse
	(*CloseChannelMessage)(nil),   // 6: channel.CloseChannelMessage
	(*CloseChannelResponse)(nil),  // 7: channel.CloseChannelResponse
}
var file_message_proto_depIdxs = []int32{
	0, // 0: channel.MessageService.SendMessage:input_type -> channel.SendMessageRequest
	2, // 1: channel.MessageService.SendSecret:input_type -> channel.SendSecretMessage
	4, // 2: channel.MessageService.BroadcastNoti:input_type -> channel.BroadcastNotiMessage
	6, // 3: channel.MessageService.CloseChannel:input_type -> channel.CloseChannelMessage
	1, // 4: channel.MessageService.SendMessage:output_type -> channel.SendMessageResponse
	3, // 5: channel.MessageService.SendSecret:output_type -> channel.SendSecretResponse
	5, // 6: channel.MessageService.BroadcastNoti:output_type -> channel.BroadcastNotiResponse
	7, // 7: channel.MessageService.CloseChannel:output_type -> channel.CloseChannelResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSecretMessage); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSecretResponse); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastNotiMessage); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastNotiResponse); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseChannelMessage); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseChannelResponse); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
