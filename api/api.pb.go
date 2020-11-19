// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: api.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	schema "github.com/tweag/trustix/schema"
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

type SubmitResponse_Status int32

const (
	SubmitResponse_OK SubmitResponse_Status = 0
)

// Enum value maps for SubmitResponse_Status.
var (
	SubmitResponse_Status_name = map[int32]string{
		0: "OK",
	}
	SubmitResponse_Status_value = map[string]int32{
		"OK": 0,
	}
)

func (x SubmitResponse_Status) Enum() *SubmitResponse_Status {
	p := new(SubmitResponse_Status)
	*p = x
	return p
}

func (x SubmitResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubmitResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (SubmitResponse_Status) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x SubmitResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SubmitResponse_Status) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SubmitResponse_Status(num)
	return nil
}

// Deprecated: Use SubmitResponse_Status.Descriptor instead.
func (SubmitResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{11, 0}
}

type STHRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *STHRequest) Reset() {
	*x = STHRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *STHRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*STHRequest) ProtoMessage() {}

func (x *STHRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use STHRequest.ProtoReflect.Descriptor instead.
func (*STHRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type GetLogConsistencyProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstSize  *uint64 `protobuf:"varint,1,req,name=FirstSize" json:"FirstSize,omitempty"`
	SecondSize *uint64 `protobuf:"varint,2,req,name=SecondSize" json:"SecondSize,omitempty"`
}

func (x *GetLogConsistencyProofRequest) Reset() {
	*x = GetLogConsistencyProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogConsistencyProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogConsistencyProofRequest) ProtoMessage() {}

func (x *GetLogConsistencyProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogConsistencyProofRequest.ProtoReflect.Descriptor instead.
func (*GetLogConsistencyProofRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetLogConsistencyProofRequest) GetFirstSize() uint64 {
	if x != nil && x.FirstSize != nil {
		return *x.FirstSize
	}
	return 0
}

func (x *GetLogConsistencyProofRequest) GetSecondSize() uint64 {
	if x != nil && x.SecondSize != nil {
		return *x.SecondSize
	}
	return 0
}

type ProofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof [][]byte `protobuf:"bytes,1,rep,name=Proof" json:"Proof,omitempty"`
}

func (x *ProofResponse) Reset() {
	*x = ProofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofResponse) ProtoMessage() {}

func (x *ProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofResponse.ProtoReflect.Descriptor instead.
func (*ProofResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *ProofResponse) GetProof() [][]byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type GetLogAuditProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index    *uint64 `protobuf:"varint,1,req,name=Index" json:"Index,omitempty"`
	TreeSize *uint64 `protobuf:"varint,2,req,name=TreeSize" json:"TreeSize,omitempty"`
}

func (x *GetLogAuditProofRequest) Reset() {
	*x = GetLogAuditProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogAuditProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogAuditProofRequest) ProtoMessage() {}

func (x *GetLogAuditProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogAuditProofRequest.ProtoReflect.Descriptor instead.
func (*GetLogAuditProofRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetLogAuditProofRequest) GetIndex() uint64 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *GetLogAuditProofRequest) GetTreeSize() uint64 {
	if x != nil && x.TreeSize != nil {
		return *x.TreeSize
	}
	return 0
}

type GetLogEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start  *uint64 `protobuf:"varint,1,req,name=Start" json:"Start,omitempty"`
	Finish *uint64 `protobuf:"varint,2,req,name=Finish" json:"Finish,omitempty"`
}

func (x *GetLogEntriesRequest) Reset() {
	*x = GetLogEntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogEntriesRequest) ProtoMessage() {}

func (x *GetLogEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogEntriesRequest.ProtoReflect.Descriptor instead.
func (*GetLogEntriesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetLogEntriesRequest) GetStart() uint64 {
	if x != nil && x.Start != nil {
		return *x.Start
	}
	return 0
}

func (x *GetLogEntriesRequest) GetFinish() uint64 {
	if x != nil && x.Finish != nil {
		return *x.Finish
	}
	return 0
}

type GetMapValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     []byte `protobuf:"bytes,1,req,name=Key" json:"Key,omitempty"`
	MapRoot []byte `protobuf:"bytes,2,req,name=MapRoot" json:"MapRoot,omitempty"`
}

func (x *GetMapValueRequest) Reset() {
	*x = GetMapValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapValueRequest) ProtoMessage() {}

func (x *GetMapValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapValueRequest.ProtoReflect.Descriptor instead.
func (*GetMapValueRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetMapValueRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GetMapValueRequest) GetMapRoot() []byte {
	if x != nil {
		return x.MapRoot
	}
	return nil
}

type SparseCompactMerkleProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SideNodes             [][]byte `protobuf:"bytes,1,rep,name=SideNodes" json:"SideNodes,omitempty"`
	NonMembershipLeafData []byte   `protobuf:"bytes,2,req,name=NonMembershipLeafData" json:"NonMembershipLeafData,omitempty"`
	BitMask               []byte   `protobuf:"bytes,3,req,name=BitMask" json:"BitMask,omitempty"`
	NumSideNodes          *uint64  `protobuf:"varint,4,req,name=NumSideNodes" json:"NumSideNodes,omitempty"`
}

func (x *SparseCompactMerkleProof) Reset() {
	*x = SparseCompactMerkleProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparseCompactMerkleProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparseCompactMerkleProof) ProtoMessage() {}

func (x *SparseCompactMerkleProof) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparseCompactMerkleProof.ProtoReflect.Descriptor instead.
func (*SparseCompactMerkleProof) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *SparseCompactMerkleProof) GetSideNodes() [][]byte {
	if x != nil {
		return x.SideNodes
	}
	return nil
}

func (x *SparseCompactMerkleProof) GetNonMembershipLeafData() []byte {
	if x != nil {
		return x.NonMembershipLeafData
	}
	return nil
}

func (x *SparseCompactMerkleProof) GetBitMask() []byte {
	if x != nil {
		return x.BitMask
	}
	return nil
}

func (x *SparseCompactMerkleProof) GetNumSideNodes() uint64 {
	if x != nil && x.NumSideNodes != nil {
		return *x.NumSideNodes
	}
	return 0
}

type MapValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Note that the Value field is actually a MapEntry
	// but we need to return the marshaled version
	// as that's what the proof is created from
	Value []byte                    `protobuf:"bytes,1,req,name=Value" json:"Value,omitempty"`
	Proof *SparseCompactMerkleProof `protobuf:"bytes,2,req,name=Proof" json:"Proof,omitempty"`
}

func (x *MapValueResponse) Reset() {
	*x = MapValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapValueResponse) ProtoMessage() {}

func (x *MapValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapValueResponse.ProtoReflect.Descriptor instead.
func (*MapValueResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *MapValueResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *MapValueResponse) GetProof() *SparseCompactMerkleProof {
	if x != nil {
		return x.Proof
	}
	return nil
}

type LogEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leaves []*schema.LogLeaf `protobuf:"bytes,1,rep,name=Leaves" json:"Leaves,omitempty"`
}

func (x *LogEntriesResponse) Reset() {
	*x = LogEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntriesResponse) ProtoMessage() {}

func (x *LogEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntriesResponse.ProtoReflect.Descriptor instead.
func (*LogEntriesResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *LogEntriesResponse) GetLeaves() []*schema.LogLeaf {
	if x != nil {
		return x.Leaves
	}
	return nil
}

type KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,req,name=Key" json:"Key,omitempty"`
	Value []byte `protobuf:"bytes,2,req,name=Value" json:"Value,omitempty"`
}

func (x *KeyValuePair) Reset() {
	*x = KeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValuePair) ProtoMessage() {}

func (x *KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValuePair.ProtoReflect.Descriptor instead.
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *KeyValuePair) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyValuePair) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*KeyValuePair `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{10}
}

func (x *SubmitRequest) GetItems() []*KeyValuePair {
	if x != nil {
		return x.Items
	}
	return nil
}

type SubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *SubmitResponse_Status `protobuf:"varint,1,req,name=status,enum=trustix.SubmitResponse_Status" json:"status,omitempty"`
}

func (x *SubmitResponse) Reset() {
	*x = SubmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponse) ProtoMessage() {}

func (x *SubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponse.ProtoReflect.Descriptor instead.
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{11}
}

func (x *SubmitResponse) GetStatus() SubmitResponse_Status {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return SubmitResponse_OK
}

type FlushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlushRequest) Reset() {
	*x = FlushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushRequest) ProtoMessage() {}

func (x *FlushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushRequest.ProtoReflect.Descriptor instead.
func (*FlushRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{12}
}

type FlushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlushResponse) Reset() {
	*x = FlushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushResponse) ProtoMessage() {}

func (x *FlushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushResponse.ProtoReflect.Descriptor instead.
func (*FlushResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{13}
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x69, 0x78, 0x1a, 0x10, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6c,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a,
	0x53, 0x54, 0x48, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x09,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0a, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x22, 0x4b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x41, 0x75, 0x64, 0x69, 0x74, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x72, 0x65, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x08, 0x54, 0x72, 0x65, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x44, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x06, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x22, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x61, 0x70, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x07, 0x4d, 0x61,
	0x70, 0x52, 0x6f, 0x6f, 0x74, 0x22, 0xac, 0x01, 0x0a, 0x18, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x64, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x64, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x34, 0x0a, 0x15, 0x4e, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x4c, 0x65, 0x61, 0x66, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c, 0x52,
	0x15, 0x4e, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4c, 0x65,
	0x61, 0x66, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x69, 0x74, 0x4d, 0x61, 0x73,
	0x6b, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x07, 0x42, 0x69, 0x74, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x75, 0x6d, 0x53, 0x69, 0x64, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0c, 0x4e, 0x75, 0x6d, 0x53, 0x69, 0x64, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x22, 0x61, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x37,
	0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x52, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x36, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x06, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x61, 0x66, 0x52, 0x06, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x22,
	0x36, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x03, 0x4b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3c, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69,
	0x78, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x5a, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69,
	0x78, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x10, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0x00, 0x22, 0x0e, 0x0a, 0x0c, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xf2, 0x05, 0x0a, 0x0d, 0x54, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x4c, 0x6f,
	0x67, 0x41, 0x50, 0x49, 0x12, 0x25, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x53, 0x54, 0x48, 0x12, 0x13,
	0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x53, 0x54, 0x48, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x04, 0x2e, 0x53, 0x54, 0x48, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x05, 0x46, 0x6c, 0x75, 0x73,
	0x68, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x46, 0x6c, 0x75, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x69, 0x78, 0x2e, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x26, 0x2e, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x41, 0x75, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x12, 0x20, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x41, 0x75, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x1d, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x2e, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x69, 0x78, 0x2e, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x48, 0x4c,
	0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x12, 0x26, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x69, 0x78, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x48, 0x4c, 0x6f, 0x67,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x20, 0x2e, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x69, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x48, 0x4c,
	0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x69, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x69, 0x78, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x77, 0x65, 0x61, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x69, 0x78, 0x2f, 0x61, 0x70, 0x69,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_proto_goTypes = []interface{}{
	(SubmitResponse_Status)(0),            // 0: trustix.SubmitResponse.Status
	(*STHRequest)(nil),                    // 1: trustix.STHRequest
	(*GetLogConsistencyProofRequest)(nil), // 2: trustix.GetLogConsistencyProofRequest
	(*ProofResponse)(nil),                 // 3: trustix.ProofResponse
	(*GetLogAuditProofRequest)(nil),       // 4: trustix.GetLogAuditProofRequest
	(*GetLogEntriesRequest)(nil),          // 5: trustix.GetLogEntriesRequest
	(*GetMapValueRequest)(nil),            // 6: trustix.GetMapValueRequest
	(*SparseCompactMerkleProof)(nil),      // 7: trustix.SparseCompactMerkleProof
	(*MapValueResponse)(nil),              // 8: trustix.MapValueResponse
	(*LogEntriesResponse)(nil),            // 9: trustix.LogEntriesResponse
	(*KeyValuePair)(nil),                  // 10: trustix.KeyValuePair
	(*SubmitRequest)(nil),                 // 11: trustix.SubmitRequest
	(*SubmitResponse)(nil),                // 12: trustix.SubmitResponse
	(*FlushRequest)(nil),                  // 13: trustix.FlushRequest
	(*FlushResponse)(nil),                 // 14: trustix.FlushResponse
	(*schema.LogLeaf)(nil),                // 15: LogLeaf
	(*schema.STH)(nil),                    // 16: STH
}
var file_api_proto_depIdxs = []int32{
	7,  // 0: trustix.MapValueResponse.Proof:type_name -> trustix.SparseCompactMerkleProof
	15, // 1: trustix.LogEntriesResponse.Leaves:type_name -> LogLeaf
	10, // 2: trustix.SubmitRequest.Items:type_name -> trustix.KeyValuePair
	0,  // 3: trustix.SubmitResponse.status:type_name -> trustix.SubmitResponse.Status
	1,  // 4: trustix.TrustixLogAPI.GetSTH:input_type -> trustix.STHRequest
	11, // 5: trustix.TrustixLogAPI.Submit:input_type -> trustix.SubmitRequest
	13, // 6: trustix.TrustixLogAPI.Flush:input_type -> trustix.FlushRequest
	2,  // 7: trustix.TrustixLogAPI.GetLogConsistencyProof:input_type -> trustix.GetLogConsistencyProofRequest
	4,  // 8: trustix.TrustixLogAPI.GetLogAuditProof:input_type -> trustix.GetLogAuditProofRequest
	5,  // 9: trustix.TrustixLogAPI.GetLogEntries:input_type -> trustix.GetLogEntriesRequest
	6,  // 10: trustix.TrustixLogAPI.GetMapValue:input_type -> trustix.GetMapValueRequest
	2,  // 11: trustix.TrustixLogAPI.GetMHLogConsistencyProof:input_type -> trustix.GetLogConsistencyProofRequest
	4,  // 12: trustix.TrustixLogAPI.GetMHLogAuditProof:input_type -> trustix.GetLogAuditProofRequest
	5,  // 13: trustix.TrustixLogAPI.GetMHLogEntries:input_type -> trustix.GetLogEntriesRequest
	16, // 14: trustix.TrustixLogAPI.GetSTH:output_type -> STH
	12, // 15: trustix.TrustixLogAPI.Submit:output_type -> trustix.SubmitResponse
	14, // 16: trustix.TrustixLogAPI.Flush:output_type -> trustix.FlushResponse
	3,  // 17: trustix.TrustixLogAPI.GetLogConsistencyProof:output_type -> trustix.ProofResponse
	3,  // 18: trustix.TrustixLogAPI.GetLogAuditProof:output_type -> trustix.ProofResponse
	9,  // 19: trustix.TrustixLogAPI.GetLogEntries:output_type -> trustix.LogEntriesResponse
	8,  // 20: trustix.TrustixLogAPI.GetMapValue:output_type -> trustix.MapValueResponse
	3,  // 21: trustix.TrustixLogAPI.GetMHLogConsistencyProof:output_type -> trustix.ProofResponse
	3,  // 22: trustix.TrustixLogAPI.GetMHLogAuditProof:output_type -> trustix.ProofResponse
	9,  // 23: trustix.TrustixLogAPI.GetMHLogEntries:output_type -> trustix.LogEntriesResponse
	14, // [14:24] is the sub-list for method output_type
	4,  // [4:14] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*STHRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogConsistencyProofRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofResponse); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogAuditProofRequest); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogEntriesRequest); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapValueRequest); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparseCompactMerkleProof); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapValueResponse); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntriesResponse); i {
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
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValuePair); i {
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
		file_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRequest); i {
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
		file_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponse); i {
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
		file_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushRequest); i {
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
		file_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
