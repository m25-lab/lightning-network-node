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

	FromAddress string `protobuf:"bytes,1,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	ToAddress   string `protobuf:"bytes,2,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	Signature   string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Payload     string `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
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

func (x *OpenChannelRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *OpenChannelRequest) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *OpenChannelRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *OpenChannelRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
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

	ChannelId     string `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
	FromAddress   string `protobuf:"bytes,2,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	FromHashCode  string `protobuf:"bytes,3,opt,name=fromHashCode,proto3" json:"fromHashCode,omitempty"`
	FromPayload   string `protobuf:"bytes,4,opt,name=fromPayload,proto3" json:"fromPayload,omitempty"`
	FromSignature string `protobuf:"bytes,5,opt,name=fromSignature,proto3" json:"fromSignature,omitempty"`
	ToAddress     string `protobuf:"bytes,6,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	ToHashcode    string `protobuf:"bytes,7,opt,name=toHashcode,proto3" json:"toHashcode,omitempty"`
	ToPayload     string `protobuf:"bytes,8,opt,name=toPayload,proto3" json:"toPayload,omitempty"`
	ToSignature   string `protobuf:"bytes,9,opt,name=toSignature,proto3" json:"toSignature,omitempty"`
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

func (x *CreateCommitmentRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *CreateCommitmentRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *CreateCommitmentRequest) GetFromHashCode() string {
	if x != nil {
		return x.FromHashCode
	}
	return ""
}

func (x *CreateCommitmentRequest) GetFromPayload() string {
	if x != nil {
		return x.FromPayload
	}
	return ""
}

func (x *CreateCommitmentRequest) GetFromSignature() string {
	if x != nil {
		return x.FromSignature
	}
	return ""
}

func (x *CreateCommitmentRequest) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *CreateCommitmentRequest) GetToHashcode() string {
	if x != nil {
		return x.ToHashcode
	}
	return ""
}

func (x *CreateCommitmentRequest) GetToPayload() string {
	if x != nil {
		return x.ToPayload
	}
	return ""
}

func (x *CreateCommitmentRequest) GetToSignature() string {
	if x != nil {
		return x.ToSignature
	}
	return ""
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

type GetChannelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetChannelsRequest) Reset() {
	*x = GetChannelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChannelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelsRequest) ProtoMessage() {}

func (x *GetChannelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelsRequest.ProtoReflect.Descriptor instead.
func (*GetChannelsRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{8}
}

func (x *GetChannelsRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetChannelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*GetChannelResponse `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *GetChannelsResponse) Reset() {
	*x = GetChannelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChannelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelsResponse) ProtoMessage() {}

func (x *GetChannelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelsResponse.ProtoReflect.Descriptor instead.
func (*GetChannelsResponse) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{9}
}

func (x *GetChannelsResponse) GetRows() []*GetChannelResponse {
	if x != nil {
		return x.Rows
	}
	return nil
}

type GetChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetChannelRequest) Reset() {
	*x = GetChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelRequest) ProtoMessage() {}

func (x *GetChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelRequest.ProtoReflect.Descriptor instead.
func (*GetChannelRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{10}
}

func (x *GetChannelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status      string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	FromAddress string `protobuf:"bytes,3,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	ToAddress   string `protobuf:"bytes,4,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	SignatureA  string `protobuf:"bytes,5,opt,name=signatureA,proto3" json:"signatureA,omitempty"`
	SignatureB  string `protobuf:"bytes,6,opt,name=signatureB,proto3" json:"signatureB,omitempty"`
	Payload     string `protobuf:"bytes,7,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *GetChannelResponse) Reset() {
	*x = GetChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelResponse) ProtoMessage() {}

func (x *GetChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelResponse.ProtoReflect.Descriptor instead.
func (*GetChannelResponse) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{11}
}

func (x *GetChannelResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetChannelResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetChannelResponse) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *GetChannelResponse) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *GetChannelResponse) GetSignatureA() string {
	if x != nil {
		return x.SignatureA
	}
	return ""
}

func (x *GetChannelResponse) GetSignatureB() string {
	if x != nil {
		return x.SignatureB
	}
	return ""
}

func (x *GetChannelResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_channel_proto protoreflect.FileDescriptor

var file_channel_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x8c, 0x01, 0x0a, 0x12, 0x4f, 0x70, 0x65,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x31, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc3, 0x02, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x48, 0x61,
	0x73, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72,
	0x6f, 0x6d, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x66, 0x72, 0x6f, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x6f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0x36, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x0a, 0x17, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22,
	0x36, 0x0a, 0x18, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48, 0x61, 0x73, 0x68, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x0a, 0x17, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x36, 0x0a, 0x18, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x46,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x32, 0x86, 0x04, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4f, 0x70,
	0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48, 0x61,
	0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x69,
	0x6d, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x32, 0x35, 0x2d,
	0x6c, 0x61, 0x62, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_channel_proto_goTypes = []interface{}{
	(*OpenChannelRequest)(nil),       // 0: channel.OpenChannelRequest
	(*OpenChannelResponse)(nil),      // 1: channel.OpenChannelResponse
	(*CreateCommitmentRequest)(nil),  // 2: channel.CreateCommitmentRequest
	(*CreateCommitmentResponse)(nil), // 3: channel.CreateCommitmentResponse
	(*WithdrawHashlockRequest)(nil),  // 4: channel.WithdrawHashlockRequest
	(*WithdrawHashlockResponse)(nil), // 5: channel.WithdrawHashlockResponse
	(*WithdrawTimelockRequest)(nil),  // 6: channel.WithdrawTimelockRequest
	(*WithdrawTimelockResponse)(nil), // 7: channel.WithdrawTimelockResponse
	(*GetChannelsRequest)(nil),       // 8: channel.GetChannelsRequest
	(*GetChannelsResponse)(nil),      // 9: channel.GetChannelsResponse
	(*GetChannelRequest)(nil),        // 10: channel.GetChannelRequest
	(*GetChannelResponse)(nil),       // 11: channel.GetChannelResponse
}
var file_channel_proto_depIdxs = []int32{
	11, // 0: channel.GetChannelsResponse.rows:type_name -> channel.GetChannelResponse
	0,  // 1: channel.ChannelService.OpenChannel:input_type -> channel.OpenChannelRequest
	8,  // 2: channel.ChannelService.GetChannels:input_type -> channel.GetChannelsRequest
	10, // 3: channel.ChannelService.GetChannelById:input_type -> channel.GetChannelRequest
	2,  // 4: channel.ChannelService.CreateCommitment:input_type -> channel.CreateCommitmentRequest
	4,  // 5: channel.ChannelService.WithdrawHashlock:input_type -> channel.WithdrawHashlockRequest
	6,  // 6: channel.ChannelService.WithdrawTimelock:input_type -> channel.WithdrawTimelockRequest
	1,  // 7: channel.ChannelService.OpenChannel:output_type -> channel.OpenChannelResponse
	9,  // 8: channel.ChannelService.GetChannels:output_type -> channel.GetChannelsResponse
	11, // 9: channel.ChannelService.GetChannelById:output_type -> channel.GetChannelResponse
	3,  // 10: channel.ChannelService.CreateCommitment:output_type -> channel.CreateCommitmentResponse
	5,  // 11: channel.ChannelService.WithdrawHashlock:output_type -> channel.WithdrawHashlockResponse
	7,  // 12: channel.ChannelService.WithdrawTimelock:output_type -> channel.WithdrawTimelockResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
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
		file_channel_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChannelsRequest); i {
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
		file_channel_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChannelsResponse); i {
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
		file_channel_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChannelRequest); i {
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
		file_channel_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChannelResponse); i {
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
			NumMessages:   12,
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