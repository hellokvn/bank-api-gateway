// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.19.4
// source: pkg/common/proto/bank-funds-command.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DepositFundsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DepositFundsRequest) Reset() {
	*x = DepositFundsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositFundsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositFundsRequest) ProtoMessage() {}

func (x *DepositFundsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositFundsRequest.ProtoReflect.Descriptor instead.
func (*DepositFundsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_bank_funds_command_proto_rawDescGZIP(), []int{0}
}

func (x *DepositFundsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DepositFundsRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type DepositFundsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  []string `protobuf:"bytes,2,rep,name=error,proto3" json:"error,omitempty"`
}

func (x *DepositFundsResponse) Reset() {
	*x = DepositFundsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositFundsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositFundsResponse) ProtoMessage() {}

func (x *DepositFundsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositFundsResponse.ProtoReflect.Descriptor instead.
func (*DepositFundsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_bank_funds_command_proto_rawDescGZIP(), []int{1}
}

func (x *DepositFundsResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DepositFundsResponse) GetError() []string {
	if x != nil {
		return x.Error
	}
	return nil
}

type WithdrawFundsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *WithdrawFundsRequest) Reset() {
	*x = WithdrawFundsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawFundsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawFundsRequest) ProtoMessage() {}

func (x *WithdrawFundsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawFundsRequest.ProtoReflect.Descriptor instead.
func (*WithdrawFundsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_bank_funds_command_proto_rawDescGZIP(), []int{2}
}

func (x *WithdrawFundsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WithdrawFundsRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type WithdrawFundsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  []string `protobuf:"bytes,2,rep,name=error,proto3" json:"error,omitempty"`
}

func (x *WithdrawFundsResponse) Reset() {
	*x = WithdrawFundsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawFundsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawFundsResponse) ProtoMessage() {}

func (x *WithdrawFundsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawFundsResponse.ProtoReflect.Descriptor instead.
func (*WithdrawFundsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_bank_funds_command_proto_rawDescGZIP(), []int{3}
}

func (x *WithdrawFundsResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *WithdrawFundsResponse) GetError() []string {
	if x != nil {
		return x.Error
	}
	return nil
}

type TransferFundsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromId string `protobuf:"bytes,1,opt,name=fromId,proto3" json:"fromId,omitempty"`
	ToId   string `protobuf:"bytes,2,opt,name=toId,proto3" json:"toId,omitempty"`
	Amount int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferFundsRequest) Reset() {
	*x = TransferFundsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferFundsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferFundsRequest) ProtoMessage() {}

func (x *TransferFundsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferFundsRequest.ProtoReflect.Descriptor instead.
func (*TransferFundsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_bank_funds_command_proto_rawDescGZIP(), []int{4}
}

func (x *TransferFundsRequest) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *TransferFundsRequest) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *TransferFundsRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferFundsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  []string `protobuf:"bytes,2,rep,name=error,proto3" json:"error,omitempty"`
}

func (x *TransferFundsResponse) Reset() {
	*x = TransferFundsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferFundsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferFundsResponse) ProtoMessage() {}

func (x *TransferFundsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_bank_funds_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferFundsResponse.ProtoReflect.Descriptor instead.
func (*TransferFundsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_bank_funds_command_proto_rawDescGZIP(), []int{5}
}

func (x *TransferFundsResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TransferFundsResponse) GetError() []string {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_pkg_common_proto_bank_funds_command_proto protoreflect.FileDescriptor

var file_pkg_common_proto_bank_funds_command_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2d, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x61, 0x6e,
	0x6b, 0x5f, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22,
	0x3d, 0x0a, 0x13, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44,
	0x0a, 0x14, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x3e, 0x0a, 0x14, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x15, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5a, 0x0a, 0x14, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xce,
	0x02, 0x0a, 0x17, 0x42, 0x61, 0x6e, 0x6b, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x27, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x5f, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x66, 0x75, 0x6e, 0x64, 0x73,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x66, 0x0a, 0x0d, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x75, 0x6e, 0x64, 0x73,
	0x12, 0x28, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x75,
	0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x5f, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x28, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x5f,
	0x66, 0x75, 0x6e, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_common_proto_bank_funds_command_proto_rawDescOnce sync.Once
	file_pkg_common_proto_bank_funds_command_proto_rawDescData = file_pkg_common_proto_bank_funds_command_proto_rawDesc
)

func file_pkg_common_proto_bank_funds_command_proto_rawDescGZIP() []byte {
	file_pkg_common_proto_bank_funds_command_proto_rawDescOnce.Do(func() {
		file_pkg_common_proto_bank_funds_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_common_proto_bank_funds_command_proto_rawDescData)
	})
	return file_pkg_common_proto_bank_funds_command_proto_rawDescData
}

var file_pkg_common_proto_bank_funds_command_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_common_proto_bank_funds_command_proto_goTypes = []interface{}{
	(*DepositFundsRequest)(nil),   // 0: bank_funds_command.DepositFundsRequest
	(*DepositFundsResponse)(nil),  // 1: bank_funds_command.DepositFundsResponse
	(*WithdrawFundsRequest)(nil),  // 2: bank_funds_command.WithdrawFundsRequest
	(*WithdrawFundsResponse)(nil), // 3: bank_funds_command.WithdrawFundsResponse
	(*TransferFundsRequest)(nil),  // 4: bank_funds_command.TransferFundsRequest
	(*TransferFundsResponse)(nil), // 5: bank_funds_command.TransferFundsResponse
}
var file_pkg_common_proto_bank_funds_command_proto_depIdxs = []int32{
	0, // 0: bank_funds_command.BankFundsCommandService.DepositFunds:input_type -> bank_funds_command.DepositFundsRequest
	2, // 1: bank_funds_command.BankFundsCommandService.WithdrawFunds:input_type -> bank_funds_command.WithdrawFundsRequest
	4, // 2: bank_funds_command.BankFundsCommandService.TransferFunds:input_type -> bank_funds_command.TransferFundsRequest
	1, // 3: bank_funds_command.BankFundsCommandService.DepositFunds:output_type -> bank_funds_command.DepositFundsResponse
	3, // 4: bank_funds_command.BankFundsCommandService.WithdrawFunds:output_type -> bank_funds_command.WithdrawFundsResponse
	5, // 5: bank_funds_command.BankFundsCommandService.TransferFunds:output_type -> bank_funds_command.TransferFundsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_common_proto_bank_funds_command_proto_init() }
func file_pkg_common_proto_bank_funds_command_proto_init() {
	if File_pkg_common_proto_bank_funds_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_common_proto_bank_funds_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositFundsRequest); i {
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
		file_pkg_common_proto_bank_funds_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositFundsResponse); i {
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
		file_pkg_common_proto_bank_funds_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawFundsRequest); i {
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
		file_pkg_common_proto_bank_funds_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawFundsResponse); i {
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
		file_pkg_common_proto_bank_funds_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferFundsRequest); i {
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
		file_pkg_common_proto_bank_funds_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferFundsResponse); i {
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
			RawDescriptor: file_pkg_common_proto_bank_funds_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_common_proto_bank_funds_command_proto_goTypes,
		DependencyIndexes: file_pkg_common_proto_bank_funds_command_proto_depIdxs,
		MessageInfos:      file_pkg_common_proto_bank_funds_command_proto_msgTypes,
	}.Build()
	File_pkg_common_proto_bank_funds_command_proto = out.File
	file_pkg_common_proto_bank_funds_command_proto_rawDesc = nil
	file_pkg_common_proto_bank_funds_command_proto_goTypes = nil
	file_pkg_common_proto_bank_funds_command_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BankFundsCommandServiceClient is the client API for BankFundsCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BankFundsCommandServiceClient interface {
	DepositFunds(ctx context.Context, in *DepositFundsRequest, opts ...grpc.CallOption) (*DepositFundsResponse, error)
	WithdrawFunds(ctx context.Context, in *WithdrawFundsRequest, opts ...grpc.CallOption) (*WithdrawFundsResponse, error)
	TransferFunds(ctx context.Context, in *TransferFundsRequest, opts ...grpc.CallOption) (*TransferFundsResponse, error)
}

type bankFundsCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankFundsCommandServiceClient(cc grpc.ClientConnInterface) BankFundsCommandServiceClient {
	return &bankFundsCommandServiceClient{cc}
}

func (c *bankFundsCommandServiceClient) DepositFunds(ctx context.Context, in *DepositFundsRequest, opts ...grpc.CallOption) (*DepositFundsResponse, error) {
	out := new(DepositFundsResponse)
	err := c.cc.Invoke(ctx, "/bank_funds_command.BankFundsCommandService/DepositFunds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankFundsCommandServiceClient) WithdrawFunds(ctx context.Context, in *WithdrawFundsRequest, opts ...grpc.CallOption) (*WithdrawFundsResponse, error) {
	out := new(WithdrawFundsResponse)
	err := c.cc.Invoke(ctx, "/bank_funds_command.BankFundsCommandService/WithdrawFunds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankFundsCommandServiceClient) TransferFunds(ctx context.Context, in *TransferFundsRequest, opts ...grpc.CallOption) (*TransferFundsResponse, error) {
	out := new(TransferFundsResponse)
	err := c.cc.Invoke(ctx, "/bank_funds_command.BankFundsCommandService/TransferFunds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankFundsCommandServiceServer is the server API for BankFundsCommandService service.
type BankFundsCommandServiceServer interface {
	DepositFunds(context.Context, *DepositFundsRequest) (*DepositFundsResponse, error)
	WithdrawFunds(context.Context, *WithdrawFundsRequest) (*WithdrawFundsResponse, error)
	TransferFunds(context.Context, *TransferFundsRequest) (*TransferFundsResponse, error)
}

// UnimplementedBankFundsCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBankFundsCommandServiceServer struct {
}

func (*UnimplementedBankFundsCommandServiceServer) DepositFunds(context.Context, *DepositFundsRequest) (*DepositFundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositFunds not implemented")
}
func (*UnimplementedBankFundsCommandServiceServer) WithdrawFunds(context.Context, *WithdrawFundsRequest) (*WithdrawFundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawFunds not implemented")
}
func (*UnimplementedBankFundsCommandServiceServer) TransferFunds(context.Context, *TransferFundsRequest) (*TransferFundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferFunds not implemented")
}

func RegisterBankFundsCommandServiceServer(s *grpc.Server, srv BankFundsCommandServiceServer) {
	s.RegisterService(&_BankFundsCommandService_serviceDesc, srv)
}

func _BankFundsCommandService_DepositFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositFundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankFundsCommandServiceServer).DepositFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_funds_command.BankFundsCommandService/DepositFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankFundsCommandServiceServer).DepositFunds(ctx, req.(*DepositFundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankFundsCommandService_WithdrawFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawFundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankFundsCommandServiceServer).WithdrawFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_funds_command.BankFundsCommandService/WithdrawFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankFundsCommandServiceServer).WithdrawFunds(ctx, req.(*WithdrawFundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankFundsCommandService_TransferFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferFundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankFundsCommandServiceServer).TransferFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_funds_command.BankFundsCommandService/TransferFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankFundsCommandServiceServer).TransferFunds(ctx, req.(*TransferFundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BankFundsCommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bank_funds_command.BankFundsCommandService",
	HandlerType: (*BankFundsCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DepositFunds",
			Handler:    _BankFundsCommandService_DepositFunds_Handler,
		},
		{
			MethodName: "WithdrawFunds",
			Handler:    _BankFundsCommandService_WithdrawFunds_Handler,
		},
		{
			MethodName: "TransferFunds",
			Handler:    _BankFundsCommandService_TransferFunds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/common/proto/bank-funds-command.proto",
}
