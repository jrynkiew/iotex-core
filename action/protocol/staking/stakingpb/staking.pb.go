// Copyright (c) 2020 IoTeX
// This source code is provided 'as is' and no warranties are given as to title or non-infringement, merchantability
// or fitness for purpose and, to the extent permitted by law, all liability for your use of the code is disclaimed.
// This source code is governed by Apache License 2.0 that can be found in the LICENSE file.

// To compile the proto, run:
//      protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: staking.proto

package stakingpb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Bucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index            uint64               `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	CandidateAddress string               `protobuf:"bytes,2,opt,name=candidateAddress,proto3" json:"candidateAddress,omitempty"`
	StakedAmount     string               `protobuf:"bytes,3,opt,name=stakedAmount,proto3" json:"stakedAmount,omitempty"`
	StakedDuration   uint32               `protobuf:"varint,4,opt,name=stakedDuration,proto3" json:"stakedDuration,omitempty"`
	CreateTime       *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	StakeStartTime   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=stakeStartTime,proto3" json:"stakeStartTime,omitempty"`
	UnstakeStartTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=unstakeStartTime,proto3" json:"unstakeStartTime,omitempty"`
	AutoStake        bool                 `protobuf:"varint,8,opt,name=autoStake,proto3" json:"autoStake,omitempty"`
	Owner            string               `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{0}
}

func (x *Bucket) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Bucket) GetCandidateAddress() string {
	if x != nil {
		return x.CandidateAddress
	}
	return ""
}

func (x *Bucket) GetStakedAmount() string {
	if x != nil {
		return x.StakedAmount
	}
	return ""
}

func (x *Bucket) GetStakedDuration() uint32 {
	if x != nil {
		return x.StakedDuration
	}
	return 0
}

func (x *Bucket) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Bucket) GetStakeStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StakeStartTime
	}
	return nil
}

func (x *Bucket) GetUnstakeStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.UnstakeStartTime
	}
	return nil
}

func (x *Bucket) GetAutoStake() bool {
	if x != nil {
		return x.AutoStake
	}
	return false
}

func (x *Bucket) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type BucketIndices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indices []uint64 `protobuf:"varint,1,rep,packed,name=indices,proto3" json:"indices,omitempty"`
}

func (x *BucketIndices) Reset() {
	*x = BucketIndices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketIndices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketIndices) ProtoMessage() {}

func (x *BucketIndices) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketIndices.ProtoReflect.Descriptor instead.
func (*BucketIndices) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{1}
}

func (x *BucketIndices) GetIndices() []uint64 {
	if x != nil {
		return x.Indices
	}
	return nil
}

type Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress       string `protobuf:"bytes,1,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
	OperatorAddress    string `protobuf:"bytes,2,opt,name=operatorAddress,proto3" json:"operatorAddress,omitempty"`
	RewardAddress      string `protobuf:"bytes,3,opt,name=rewardAddress,proto3" json:"rewardAddress,omitempty"`
	Name               string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Votes              string `protobuf:"bytes,5,opt,name=votes,proto3" json:"votes,omitempty"`
	SelfStakeBucketIdx uint64 `protobuf:"varint,6,opt,name=selfStakeBucketIdx,proto3" json:"selfStakeBucketIdx,omitempty"`
	SelfStake          string `protobuf:"bytes,7,opt,name=selfStake,proto3" json:"selfStake,omitempty"`
}

func (x *Candidate) Reset() {
	*x = Candidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidate) ProtoMessage() {}

func (x *Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidate.ProtoReflect.Descriptor instead.
func (*Candidate) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{2}
}

func (x *Candidate) GetOwnerAddress() string {
	if x != nil {
		return x.OwnerAddress
	}
	return ""
}

func (x *Candidate) GetOperatorAddress() string {
	if x != nil {
		return x.OperatorAddress
	}
	return ""
}

func (x *Candidate) GetRewardAddress() string {
	if x != nil {
		return x.RewardAddress
	}
	return ""
}

func (x *Candidate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Candidate) GetVotes() string {
	if x != nil {
		return x.Votes
	}
	return ""
}

func (x *Candidate) GetSelfStakeBucketIdx() uint64 {
	if x != nil {
		return x.SelfStakeBucketIdx
	}
	return 0
}

func (x *Candidate) GetSelfStake() string {
	if x != nil {
		return x.SelfStake
	}
	return ""
}

type Candidates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candidates []*Candidate `protobuf:"bytes,1,rep,name=candidates,proto3" json:"candidates,omitempty"`
}

func (x *Candidates) Reset() {
	*x = Candidates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candidates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidates) ProtoMessage() {}

func (x *Candidates) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidates.ProtoReflect.Descriptor instead.
func (*Candidates) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{3}
}

func (x *Candidates) GetCandidates() []*Candidate {
	if x != nil {
		return x.Candidates
	}
	return nil
}

type TotalAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Count  uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *TotalAmount) Reset() {
	*x = TotalAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalAmount) ProtoMessage() {}

func (x *TotalAmount) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalAmount.ProtoReflect.Descriptor instead.
func (*TotalAmount) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{4}
}

func (x *TotalAmount) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *TotalAmount) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_staking_proto protoreflect.FileDescriptor

var file_staking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x03, 0x0a, 0x06,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2a, 0x0a, 0x10,
	0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x6b,
	0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x42, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x75, 0x6e, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x61, 0x75, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x22, 0x29, 0x0a, 0x0d, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x22, 0xf7, 0x01, 0x0a, 0x09,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a,
	0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x65, 0x6c, 0x66, 0x53,
	0x74, 0x61, 0x6b, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x78, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x12, 0x73, 0x65, 0x6c, 0x66, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x66, 0x53,
	0x74, 0x61, 0x6b, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x66,
	0x53, 0x74, 0x61, 0x6b, 0x65, 0x22, 0x42, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x0b, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staking_proto_rawDescOnce sync.Once
	file_staking_proto_rawDescData = file_staking_proto_rawDesc
)

func file_staking_proto_rawDescGZIP() []byte {
	file_staking_proto_rawDescOnce.Do(func() {
		file_staking_proto_rawDescData = protoimpl.X.CompressGZIP(file_staking_proto_rawDescData)
	})
	return file_staking_proto_rawDescData
}

var file_staking_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_staking_proto_goTypes = []interface{}{
	(*Bucket)(nil),              // 0: stakingpb.Bucket
	(*BucketIndices)(nil),       // 1: stakingpb.BucketIndices
	(*Candidate)(nil),           // 2: stakingpb.Candidate
	(*Candidates)(nil),          // 3: stakingpb.Candidates
	(*TotalAmount)(nil),         // 4: stakingpb.TotalAmount
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_staking_proto_depIdxs = []int32{
	5, // 0: stakingpb.Bucket.createTime:type_name -> google.protobuf.Timestamp
	5, // 1: stakingpb.Bucket.stakeStartTime:type_name -> google.protobuf.Timestamp
	5, // 2: stakingpb.Bucket.unstakeStartTime:type_name -> google.protobuf.Timestamp
	2, // 3: stakingpb.Candidates.candidates:type_name -> stakingpb.Candidate
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_staking_proto_init() }
func file_staking_proto_init() {
	if File_staking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_staking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucket); i {
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
		file_staking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketIndices); i {
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
		file_staking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candidate); i {
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
		file_staking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candidates); i {
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
		file_staking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalAmount); i {
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
			RawDescriptor: file_staking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_staking_proto_goTypes,
		DependencyIndexes: file_staking_proto_depIdxs,
		MessageInfos:      file_staking_proto_msgTypes,
	}.Build()
	File_staking_proto = out.File
	file_staking_proto_rawDesc = nil
	file_staking_proto_goTypes = nil
	file_staking_proto_depIdxs = nil
}
