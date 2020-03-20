// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roykv-for-tikv.proto

package roykvtikv

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{0}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetReply) Reset()         { *m = SetReply{} }
func (m *SetReply) String() string { return proto.CompactTextString(m) }
func (*SetReply) ProtoMessage()    {}
func (*SetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{1}
}

func (m *SetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetReply.Unmarshal(m, b)
}
func (m *SetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetReply.Marshal(b, m, deterministic)
}
func (m *SetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetReply.Merge(m, src)
}
func (m *SetReply) XXX_Size() int {
	return xxx_messageInfo_SetReply.Size(m)
}
func (m *SetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SetReply.DiscardUnknown(m)
}

var xxx_messageInfo_SetReply proto.InternalMessageInfo

func (m *SetReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetReply struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{3}
}

func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (m *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(m, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ExistRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistRequest) Reset()         { *m = ExistRequest{} }
func (m *ExistRequest) String() string { return proto.CompactTextString(m) }
func (*ExistRequest) ProtoMessage()    {}
func (*ExistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{4}
}

func (m *ExistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistRequest.Unmarshal(m, b)
}
func (m *ExistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistRequest.Marshal(b, m, deterministic)
}
func (m *ExistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistRequest.Merge(m, src)
}
func (m *ExistRequest) XXX_Size() int {
	return xxx_messageInfo_ExistRequest.Size(m)
}
func (m *ExistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExistRequest proto.InternalMessageInfo

func (m *ExistRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ExistReply struct {
	Existed              bool     `protobuf:"varint,1,opt,name=existed,proto3" json:"existed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistReply) Reset()         { *m = ExistReply{} }
func (m *ExistReply) String() string { return proto.CompactTextString(m) }
func (*ExistReply) ProtoMessage()    {}
func (*ExistReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{5}
}

func (m *ExistReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistReply.Unmarshal(m, b)
}
func (m *ExistReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistReply.Marshal(b, m, deterministic)
}
func (m *ExistReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistReply.Merge(m, src)
}
func (m *ExistReply) XXX_Size() int {
	return xxx_messageInfo_ExistReply.Size(m)
}
func (m *ExistReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistReply.DiscardUnknown(m)
}

var xxx_messageInfo_ExistReply proto.InternalMessageInfo

func (m *ExistReply) GetExisted() bool {
	if m != nil {
		return m.Existed
	}
	return false
}

type ScanRequest struct {
	StartKey             string   `protobuf:"bytes,1,opt,name=startKey,proto3" json:"startKey,omitempty"`
	StartKeyType         string   `protobuf:"bytes,2,opt,name=startKeyType,proto3" json:"startKeyType,omitempty"`
	EndKey               string   `protobuf:"bytes,3,opt,name=endKey,proto3" json:"endKey,omitempty"`
	EndKeyType           string   `protobuf:"bytes,4,opt,name=endKeyType,proto3" json:"endKeyType,omitempty"`
	KeyPrefix            string   `protobuf:"bytes,5,opt,name=keyPrefix,proto3" json:"keyPrefix,omitempty"`
	Limit                uint64   `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanRequest) Reset()         { *m = ScanRequest{} }
func (m *ScanRequest) String() string { return proto.CompactTextString(m) }
func (*ScanRequest) ProtoMessage()    {}
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{6}
}

func (m *ScanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanRequest.Unmarshal(m, b)
}
func (m *ScanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanRequest.Marshal(b, m, deterministic)
}
func (m *ScanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanRequest.Merge(m, src)
}
func (m *ScanRequest) XXX_Size() int {
	return xxx_messageInfo_ScanRequest.Size(m)
}
func (m *ScanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanRequest proto.InternalMessageInfo

func (m *ScanRequest) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *ScanRequest) GetStartKeyType() string {
	if m != nil {
		return m.StartKeyType
	}
	return ""
}

func (m *ScanRequest) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

func (m *ScanRequest) GetEndKeyType() string {
	if m != nil {
		return m.EndKeyType
	}
	return ""
}

func (m *ScanRequest) GetKeyPrefix() string {
	if m != nil {
		return m.KeyPrefix
	}
	return ""
}

func (m *ScanRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type KVEntry struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVEntry) Reset()         { *m = KVEntry{} }
func (m *KVEntry) String() string { return proto.CompactTextString(m) }
func (*KVEntry) ProtoMessage()    {}
func (*KVEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{7}
}

func (m *KVEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVEntry.Unmarshal(m, b)
}
func (m *KVEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVEntry.Marshal(b, m, deterministic)
}
func (m *KVEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVEntry.Merge(m, src)
}
func (m *KVEntry) XXX_Size() int {
	return xxx_messageInfo_KVEntry.Size(m)
}
func (m *KVEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_KVEntry.DiscardUnknown(m)
}

var xxx_messageInfo_KVEntry proto.InternalMessageInfo

func (m *KVEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVEntry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ScanReply struct {
	Data                 []*KVEntry `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ScanReply) Reset()         { *m = ScanReply{} }
func (m *ScanReply) String() string { return proto.CompactTextString(m) }
func (*ScanReply) ProtoMessage()    {}
func (*ScanReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{8}
}

func (m *ScanReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanReply.Unmarshal(m, b)
}
func (m *ScanReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanReply.Marshal(b, m, deterministic)
}
func (m *ScanReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanReply.Merge(m, src)
}
func (m *ScanReply) XXX_Size() int {
	return xxx_messageInfo_ScanReply.Size(m)
}
func (m *ScanReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanReply.DiscardUnknown(m)
}

var xxx_messageInfo_ScanReply proto.InternalMessageInfo

func (m *ScanReply) GetData() []*KVEntry {
	if m != nil {
		return m.Data
	}
	return nil
}

type MGetRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MGetRequest) Reset()         { *m = MGetRequest{} }
func (m *MGetRequest) String() string { return proto.CompactTextString(m) }
func (*MGetRequest) ProtoMessage()    {}
func (*MGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{9}
}

func (m *MGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MGetRequest.Unmarshal(m, b)
}
func (m *MGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MGetRequest.Marshal(b, m, deterministic)
}
func (m *MGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MGetRequest.Merge(m, src)
}
func (m *MGetRequest) XXX_Size() int {
	return xxx_messageInfo_MGetRequest.Size(m)
}
func (m *MGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MGetRequest proto.InternalMessageInfo

func (m *MGetRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type MGetReply struct {
	Data                 map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MGetReply) Reset()         { *m = MGetReply{} }
func (m *MGetReply) String() string { return proto.CompactTextString(m) }
func (*MGetReply) ProtoMessage()    {}
func (*MGetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{10}
}

func (m *MGetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MGetReply.Unmarshal(m, b)
}
func (m *MGetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MGetReply.Marshal(b, m, deterministic)
}
func (m *MGetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MGetReply.Merge(m, src)
}
func (m *MGetReply) XXX_Size() int {
	return xxx_messageInfo_MGetReply.Size(m)
}
func (m *MGetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MGetReply.DiscardUnknown(m)
}

var xxx_messageInfo_MGetReply proto.InternalMessageInfo

func (m *MGetReply) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetAllRequest struct {
	KeyPrefix            string   `protobuf:"bytes,1,opt,name=keyPrefix,proto3" json:"keyPrefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllRequest) Reset()         { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()    {}
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{11}
}

func (m *GetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRequest.Unmarshal(m, b)
}
func (m *GetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRequest.Marshal(b, m, deterministic)
}
func (m *GetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRequest.Merge(m, src)
}
func (m *GetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRequest.Size(m)
}
func (m *GetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRequest proto.InternalMessageInfo

func (m *GetAllRequest) GetKeyPrefix() string {
	if m != nil {
		return m.KeyPrefix
	}
	return ""
}

type GetAllReply struct {
	Data                 map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetAllReply) Reset()         { *m = GetAllReply{} }
func (m *GetAllReply) String() string { return proto.CompactTextString(m) }
func (*GetAllReply) ProtoMessage()    {}
func (*GetAllReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{12}
}

func (m *GetAllReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllReply.Unmarshal(m, b)
}
func (m *GetAllReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllReply.Marshal(b, m, deterministic)
}
func (m *GetAllReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllReply.Merge(m, src)
}
func (m *GetAllReply) XXX_Size() int {
	return xxx_messageInfo_GetAllReply.Size(m)
}
func (m *GetAllReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllReply proto.InternalMessageInfo

func (m *GetAllReply) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CountRequest struct {
	StartKey             string   `protobuf:"bytes,1,opt,name=startKey,proto3" json:"startKey,omitempty"`
	StartKeyType         string   `protobuf:"bytes,2,opt,name=startKeyType,proto3" json:"startKeyType,omitempty"`
	EndKey               string   `protobuf:"bytes,3,opt,name=endKey,proto3" json:"endKey,omitempty"`
	EndKeyType           string   `protobuf:"bytes,4,opt,name=endKeyType,proto3" json:"endKeyType,omitempty"`
	KeyPrefix            string   `protobuf:"bytes,5,opt,name=keyPrefix,proto3" json:"keyPrefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountRequest) Reset()         { *m = CountRequest{} }
func (m *CountRequest) String() string { return proto.CompactTextString(m) }
func (*CountRequest) ProtoMessage()    {}
func (*CountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{13}
}

func (m *CountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountRequest.Unmarshal(m, b)
}
func (m *CountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountRequest.Marshal(b, m, deterministic)
}
func (m *CountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountRequest.Merge(m, src)
}
func (m *CountRequest) XXX_Size() int {
	return xxx_messageInfo_CountRequest.Size(m)
}
func (m *CountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountRequest proto.InternalMessageInfo

func (m *CountRequest) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *CountRequest) GetStartKeyType() string {
	if m != nil {
		return m.StartKeyType
	}
	return ""
}

func (m *CountRequest) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

func (m *CountRequest) GetEndKeyType() string {
	if m != nil {
		return m.EndKeyType
	}
	return ""
}

func (m *CountRequest) GetKeyPrefix() string {
	if m != nil {
		return m.KeyPrefix
	}
	return ""
}

type CountReply struct {
	Count                uint64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountReply) Reset()         { *m = CountReply{} }
func (m *CountReply) String() string { return proto.CompactTextString(m) }
func (*CountReply) ProtoMessage()    {}
func (*CountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{14}
}

func (m *CountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountReply.Unmarshal(m, b)
}
func (m *CountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountReply.Marshal(b, m, deterministic)
}
func (m *CountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountReply.Merge(m, src)
}
func (m *CountReply) XXX_Size() int {
	return xxx_messageInfo_CountReply.Size(m)
}
func (m *CountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CountReply.DiscardUnknown(m)
}

var xxx_messageInfo_CountReply proto.InternalMessageInfo

func (m *CountReply) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type DelRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelRequest) Reset()         { *m = DelRequest{} }
func (m *DelRequest) String() string { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()    {}
func (*DelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{15}
}

func (m *DelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelRequest.Unmarshal(m, b)
}
func (m *DelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelRequest.Marshal(b, m, deterministic)
}
func (m *DelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelRequest.Merge(m, src)
}
func (m *DelRequest) XXX_Size() int {
	return xxx_messageInfo_DelRequest.Size(m)
}
func (m *DelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelRequest proto.InternalMessageInfo

func (m *DelRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type DelReply struct {
	Deleted              uint64   `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelReply) Reset()         { *m = DelReply{} }
func (m *DelReply) String() string { return proto.CompactTextString(m) }
func (*DelReply) ProtoMessage()    {}
func (*DelReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbcd6879103c807, []int{16}
}

func (m *DelReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelReply.Unmarshal(m, b)
}
func (m *DelReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelReply.Marshal(b, m, deterministic)
}
func (m *DelReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelReply.Merge(m, src)
}
func (m *DelReply) XXX_Size() int {
	return xxx_messageInfo_DelReply.Size(m)
}
func (m *DelReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DelReply.DiscardUnknown(m)
}

var xxx_messageInfo_DelReply proto.InternalMessageInfo

func (m *DelReply) GetDeleted() uint64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func init() {
	proto.RegisterType((*SetRequest)(nil), "roykvtikv.SetRequest")
	proto.RegisterType((*SetReply)(nil), "roykvtikv.SetReply")
	proto.RegisterType((*GetRequest)(nil), "roykvtikv.GetRequest")
	proto.RegisterType((*GetReply)(nil), "roykvtikv.GetReply")
	proto.RegisterType((*ExistRequest)(nil), "roykvtikv.ExistRequest")
	proto.RegisterType((*ExistReply)(nil), "roykvtikv.ExistReply")
	proto.RegisterType((*ScanRequest)(nil), "roykvtikv.ScanRequest")
	proto.RegisterType((*KVEntry)(nil), "roykvtikv.KVEntry")
	proto.RegisterType((*ScanReply)(nil), "roykvtikv.ScanReply")
	proto.RegisterType((*MGetRequest)(nil), "roykvtikv.MGetRequest")
	proto.RegisterType((*MGetReply)(nil), "roykvtikv.MGetReply")
	proto.RegisterMapType((map[string]string)(nil), "roykvtikv.MGetReply.DataEntry")
	proto.RegisterType((*GetAllRequest)(nil), "roykvtikv.GetAllRequest")
	proto.RegisterType((*GetAllReply)(nil), "roykvtikv.GetAllReply")
	proto.RegisterMapType((map[string]string)(nil), "roykvtikv.GetAllReply.DataEntry")
	proto.RegisterType((*CountRequest)(nil), "roykvtikv.CountRequest")
	proto.RegisterType((*CountReply)(nil), "roykvtikv.CountReply")
	proto.RegisterType((*DelRequest)(nil), "roykvtikv.DelRequest")
	proto.RegisterType((*DelReply)(nil), "roykvtikv.DelReply")
}

func init() { proto.RegisterFile("roykv-for-tikv.proto", fileDescriptor_afbcd6879103c807) }

var fileDescriptor_afbcd6879103c807 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdf, 0x8a, 0xda, 0x4e,
	0x14, 0xde, 0xac, 0xd1, 0x35, 0x9f, 0xfe, 0xe0, 0xc7, 0xd4, 0xb5, 0x41, 0x8a, 0xa4, 0x43, 0x59,
	0xbc, 0x51, 0xa8, 0x2e, 0xfd, 0x47, 0x6f, 0x4a, 0x5d, 0x72, 0x21, 0x42, 0x89, 0xcb, 0xde, 0xa7,
	0xeb, 0x2c, 0x04, 0x53, 0xb5, 0x71, 0x14, 0x03, 0x7d, 0x9c, 0xbe, 0x44, 0x1f, 0xa1, 0x6f, 0x55,
	0x66, 0x92, 0x49, 0x66, 0x5c, 0x91, 0xbd, 0xe8, 0x45, 0xef, 0xe6, 0x3b, 0x67, 0xbe, 0x93, 0xef,
	0xfc, 0x9b, 0xa0, 0x95, 0xac, 0xd2, 0xc5, 0xae, 0xff, 0xb0, 0x4a, 0xfa, 0x3c, 0x5a, 0xec, 0x06,
	0xeb, 0x64, 0xc5, 0x57, 0xc4, 0x91, 0x56, 0x61, 0xa0, 0xd7, 0xc0, 0x8c, 0xf1, 0x80, 0x7d, 0xdf,
	0xb2, 0x0d, 0x27, 0xff, 0xa3, 0xb2, 0x60, 0xa9, 0x6b, 0x79, 0x56, 0xcf, 0x09, 0xc4, 0x91, 0xb4,
	0x50, 0xdd, 0x85, 0xf1, 0x96, 0xb9, 0xe7, 0xd2, 0x96, 0x01, 0x4a, 0x51, 0x97, 0xac, 0x75, 0x9c,
	0x92, 0x36, 0x6a, 0x09, 0xdb, 0x6c, 0x63, 0x2e, 0x69, 0xf5, 0x20, 0x47, 0xb4, 0x0b, 0xf8, 0x27,
	0x22, 0x53, 0x0f, 0x75, 0x5f, 0xc5, 0x28, 0xbe, 0x62, 0xe9, 0x5f, 0xf1, 0xd0, 0xbc, 0xd9, 0x47,
	0x9b, 0x13, 0x31, 0xae, 0x80, 0xfc, 0x86, 0x88, 0xe2, 0xe2, 0x82, 0x09, 0xc4, 0xe6, 0xb9, 0x14,
	0x05, 0xe9, 0x2f, 0x0b, 0x8d, 0xd9, 0x7d, 0xb8, 0x54, 0x91, 0x3a, 0xa8, 0x6f, 0x78, 0x98, 0xf0,
	0x49, 0x11, 0xae, 0xc0, 0x84, 0xa2, 0xa9, 0xce, 0xb7, 0xe9, 0x5a, 0x25, 0x6e, 0xd8, 0x44, 0xce,
	0x6c, 0x39, 0x17, 0xec, 0x8a, 0xf4, 0xe6, 0x88, 0x74, 0x81, 0xec, 0x24, 0x99, 0xb6, 0xf4, 0x69,
	0x16, 0xf2, 0x02, 0xce, 0x82, 0xa5, 0x5f, 0x12, 0xf6, 0x10, 0xed, 0xdd, 0xaa, 0x74, 0x97, 0x06,
	0x51, 0x85, 0x38, 0xfa, 0x16, 0x71, 0xb7, 0xe6, 0x59, 0x3d, 0x3b, 0xc8, 0x00, 0x7d, 0x8d, 0x8b,
	0xc9, 0xdd, 0xcd, 0x92, 0x27, 0xe9, 0x93, 0xdb, 0x33, 0x82, 0x93, 0x65, 0x2b, 0xaa, 0x72, 0x05,
	0x7b, 0x1e, 0xf2, 0xd0, 0xb5, 0xbc, 0x4a, 0xaf, 0x31, 0x24, 0x83, 0xa2, 0xf7, 0x83, 0x3c, 0x6c,
	0x20, 0xfd, 0xf4, 0x25, 0x1a, 0x53, 0xad, 0x61, 0x04, 0xf6, 0x82, 0xa5, 0x1b, 0x49, 0x73, 0x02,
	0x79, 0xa6, 0x7b, 0x38, 0xd3, 0xa2, 0x67, 0x43, 0x23, 0x6e, 0x57, 0x8b, 0x5b, 0xdc, 0x19, 0x8c,
	0x43, 0x1e, 0x6a, 0xdf, 0xe8, 0xbc, 0x85, 0x53, 0x98, 0x9e, 0x9a, 0xcd, 0x87, 0xf3, 0x77, 0x16,
	0xed, 0xe3, 0x3f, 0x9f, 0xf1, 0x4f, 0x71, 0xac, 0xe4, 0x19, 0x95, 0xb4, 0x0e, 0x2a, 0x49, 0x7f,
	0xa0, 0xa1, 0xae, 0x0b, 0xa9, 0xd7, 0x86, 0x54, 0x4f, 0x93, 0xaa, 0xdd, 0xfa, 0x7b, 0x62, 0x7f,
	0x5a, 0x68, 0x7e, 0x5e, 0x6d, 0x97, 0xfc, 0x9f, 0x1e, 0x37, 0x4a, 0x81, 0x5c, 0x65, 0xbe, 0x82,
	0xf7, 0x02, 0x49, 0x81, 0x76, 0x90, 0x01, 0xea, 0x01, 0x63, 0x16, 0x9f, 0x9a, 0x89, 0x57, 0xa8,
	0xcb, 0x1b, 0xf9, 0x02, 0xce, 0x59, 0xcc, 0xd4, 0x02, 0xda, 0x81, 0x82, 0xc3, 0xdf, 0x15, 0xd8,
	0xb7, 0xd1, 0xe4, 0x8e, 0x8c, 0x50, 0x99, 0x31, 0x4e, 0x2e, 0xb5, 0x1e, 0x94, 0xef, 0x4f, 0xe7,
	0xd9, 0xa1, 0x79, 0x1d, 0xa7, 0xf4, 0x4c, 0x90, 0xfc, 0x03, 0x92, 0x7f, 0x9c, 0xe4, 0x97, 0xa4,
	0xf7, 0xa8, 0xca, 0xb7, 0x81, 0x3c, 0xd7, 0xfc, 0xfa, 0x7b, 0xd2, 0xb9, 0x7c, 0xec, 0xc8, 0xa8,
	0x6f, 0x60, 0x8b, 0xfd, 0x21, 0x6d, 0x5d, 0x4e, 0xf9, 0x7c, 0x74, 0x5a, 0x8f, 0xec, 0x05, 0x4f,
	0xcc, 0xbe, 0xc1, 0xd3, 0x76, 0xca, 0xe0, 0x4d, 0x35, 0xa9, 0x1f, 0x51, 0xcb, 0x06, 0x91, 0xb8,
	0x47, 0x66, 0x33, 0xe3, 0xb6, 0x8f, 0x4f, 0x6d, 0x96, 0xa8, 0xec, 0xa3, 0x91, 0xa8, 0x3e, 0x7f,
	0x46, 0xa2, 0x65, 0xcb, 0xb3, 0xc2, 0x8e, 0x59, 0x6c, 0x14, 0xb6, 0x6c, 0xb7, 0x51, 0x58, 0xd5,
	0x63, 0x7a, 0xf6, 0xb5, 0x26, 0x7f, 0x22, 0xa3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x96, 0x47,
	0xfe, 0x1c, 0x5c, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TiKVClient is the client API for TiKV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TiKVClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Exist(ctx context.Context, in *ExistRequest, opts ...grpc.CallOption) (*ExistReply, error)
	Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (*ScanReply, error)
	MGet(ctx context.Context, in *MGetRequest, opts ...grpc.CallOption) (*MGetReply, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllReply, error)
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountReply, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelReply, error)
}

type tiKVClient struct {
	cc *grpc.ClientConn
}

func NewTiKVClient(cc *grpc.ClientConn) TiKVClient {
	return &tiKVClient{cc}
}

func (c *tiKVClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetReply, error) {
	out := new(SetReply)
	err := c.cc.Invoke(ctx, "/roykvtikv.TiKV/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/roykvtikv.TiKV/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) Exist(ctx context.Context, in *ExistRequest, opts ...grpc.CallOption) (*ExistReply, error) {
	out := new(ExistReply)
	err := c.cc.Invoke(ctx, "/roykvtikv.TiKV/Exist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (*ScanReply, error) {
	out := new(ScanReply)
	err := c.cc.Invoke(ctx, "/roykvtikv.TiKV/Scan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) MGet(ctx context.Context, in *MGetRequest, opts ...grpc.CallOption) (*MGetReply, error) {
	out := new(MGetReply)
	err := c.cc.Invoke(ctx, "/roykvtikv.TiKV/MGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllReply, error) {
	out := new(GetAllReply)
	err := c.cc.Invoke(ctx, "/roykvtikv.TiKV/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountReply, error) {
	out := new(CountReply)
	err := c.cc.Invoke(ctx, "/roykvtikv.TiKV/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelReply, error) {
	out := new(DelReply)
	err := c.cc.Invoke(ctx, "/roykvtikv.TiKV/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TiKVServer is the server API for TiKV service.
type TiKVServer interface {
	Set(context.Context, *SetRequest) (*SetReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
	Exist(context.Context, *ExistRequest) (*ExistReply, error)
	Scan(context.Context, *ScanRequest) (*ScanReply, error)
	MGet(context.Context, *MGetRequest) (*MGetReply, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllReply, error)
	Count(context.Context, *CountRequest) (*CountReply, error)
	Del(context.Context, *DelRequest) (*DelReply, error)
}

// UnimplementedTiKVServer can be embedded to have forward compatible implementations.
type UnimplementedTiKVServer struct {
}

func (*UnimplementedTiKVServer) Set(ctx context.Context, req *SetRequest) (*SetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedTiKVServer) Get(ctx context.Context, req *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTiKVServer) Exist(ctx context.Context, req *ExistRequest) (*ExistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exist not implemented")
}
func (*UnimplementedTiKVServer) Scan(ctx context.Context, req *ScanRequest) (*ScanReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (*UnimplementedTiKVServer) MGet(ctx context.Context, req *MGetRequest) (*MGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGet not implemented")
}
func (*UnimplementedTiKVServer) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedTiKVServer) Count(ctx context.Context, req *CountRequest) (*CountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (*UnimplementedTiKVServer) Del(ctx context.Context, req *DelRequest) (*DelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}

func RegisterTiKVServer(s *grpc.Server, srv TiKVServer) {
	s.RegisterService(&_TiKV_serviceDesc, srv)
}

func _TiKV_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roykvtikv.TiKV/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roykvtikv.TiKV/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_Exist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).Exist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roykvtikv.TiKV/Exist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).Exist(ctx, req.(*ExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roykvtikv.TiKV/Scan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).Scan(ctx, req.(*ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_MGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).MGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roykvtikv.TiKV/MGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).MGet(ctx, req.(*MGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roykvtikv.TiKV/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roykvtikv.TiKV/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roykvtikv.TiKV/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TiKV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "roykvtikv.TiKV",
	HandlerType: (*TiKVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _TiKV_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TiKV_Get_Handler,
		},
		{
			MethodName: "Exist",
			Handler:    _TiKV_Exist_Handler,
		},
		{
			MethodName: "Scan",
			Handler:    _TiKV_Scan_Handler,
		},
		{
			MethodName: "MGet",
			Handler:    _TiKV_MGet_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _TiKV_GetAll_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _TiKV_Count_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _TiKV_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "roykv-for-tikv.proto",
}
