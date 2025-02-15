// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/ab.proto

package orderer

import (
	context "context"
	fmt "fmt"
	common "github.com/SmartBFT-Go/fabric-protos-go/v2/common"
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

// If BLOCK_UNTIL_READY is specified, the reply will block until the requested blocks are available,
// if FAIL_IF_NOT_READY is specified, the reply will return an error indicating that the block is not
// found.  To request that all blocks be returned indefinitely as they are created, behavior should be
// set to BLOCK_UNTIL_READY and the stop should be set to specified with a number of MAX_UINT64
type SeekInfo_SeekBehavior int32

const (
	SeekInfo_BLOCK_UNTIL_READY SeekInfo_SeekBehavior = 0
	SeekInfo_FAIL_IF_NOT_READY SeekInfo_SeekBehavior = 1
)

var SeekInfo_SeekBehavior_name = map[int32]string{
	0: "BLOCK_UNTIL_READY",
	1: "FAIL_IF_NOT_READY",
}

var SeekInfo_SeekBehavior_value = map[string]int32{
	"BLOCK_UNTIL_READY": 0,
	"FAIL_IF_NOT_READY": 1,
}

func (x SeekInfo_SeekBehavior) String() string {
	return proto.EnumName(SeekInfo_SeekBehavior_name, int32(x))
}

func (SeekInfo_SeekBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{6, 0}
}

// SeekErrorTolerance indicates to the server how block provider errors should be tolerated.  By default,
// if the deliver service detects a problem in the underlying block source (typically, in the orderer,
// a consenter error), it will begin to reject deliver requests.  This is to prevent a client from waiting
// for blocks from an orderer which is stuck in an errored state.  This is almost always the desired behavior
// and clients should stick with the default STRICT checking behavior.  However, in some scenarios, particularly
// when attempting to recover from a crash or other corruption, it's desirable to force an orderer to respond
// with blocks on a best effort basis, even if the backing consensus implementation is in an errored state.
// In this case, set the SeekErrorResponse to BEST_EFFORT to ignore the consenter errors.
type SeekInfo_SeekErrorResponse int32

const (
	SeekInfo_STRICT      SeekInfo_SeekErrorResponse = 0
	SeekInfo_BEST_EFFORT SeekInfo_SeekErrorResponse = 1
)

var SeekInfo_SeekErrorResponse_name = map[int32]string{
	0: "STRICT",
	1: "BEST_EFFORT",
}

var SeekInfo_SeekErrorResponse_value = map[string]int32{
	"STRICT":      0,
	"BEST_EFFORT": 1,
}

func (x SeekInfo_SeekErrorResponse) String() string {
	return proto.EnumName(SeekInfo_SeekErrorResponse_name, int32(x))
}

func (SeekInfo_SeekErrorResponse) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{6, 1}
}

// SeekContentType indicates what type of content to deliver in response to a request. If BLOCK is specified,
// the orderer will stream blocks back to the peer. This is the default behavior. If HEADER is  specified, the
// orderer will stream only a the header and the signature, and the payload field will be set to nil. This allows
// the requester to ascertain that the respective signed block exists in the orderer (or cluster of orderers).
type SeekInfo_SeekContentType int32

const (
	SeekInfo_BLOCK           SeekInfo_SeekContentType = 0
	SeekInfo_HEADER_WITH_SIG SeekInfo_SeekContentType = 1
)

var SeekInfo_SeekContentType_name = map[int32]string{
	0: "BLOCK",
	1: "HEADER_WITH_SIG",
}

var SeekInfo_SeekContentType_value = map[string]int32{
	"BLOCK":           0,
	"HEADER_WITH_SIG": 1,
}

func (x SeekInfo_SeekContentType) String() string {
	return proto.EnumName(SeekInfo_SeekContentType_name, int32(x))
}

func (SeekInfo_SeekContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{6, 2}
}

type BroadcastResponse struct {
	// Status code, which may be used to programatically respond to success/failure
	Status common.Status `protobuf:"varint,1,opt,name=status,proto3,enum=common.Status" json:"status,omitempty"`
	// Info string which may contain additional information about the status returned
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastResponse) Reset()         { *m = BroadcastResponse{} }
func (m *BroadcastResponse) String() string { return proto.CompactTextString(m) }
func (*BroadcastResponse) ProtoMessage()    {}
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{0}
}

func (m *BroadcastResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastResponse.Unmarshal(m, b)
}
func (m *BroadcastResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastResponse.Marshal(b, m, deterministic)
}
func (m *BroadcastResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastResponse.Merge(m, src)
}
func (m *BroadcastResponse) XXX_Size() int {
	return xxx_messageInfo_BroadcastResponse.Size(m)
}
func (m *BroadcastResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastResponse proto.InternalMessageInfo

func (m *BroadcastResponse) GetStatus() common.Status {
	if m != nil {
		return m.Status
	}
	return common.Status_UNKNOWN
}

func (m *BroadcastResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type SeekNewest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SeekNewest) Reset()         { *m = SeekNewest{} }
func (m *SeekNewest) String() string { return proto.CompactTextString(m) }
func (*SeekNewest) ProtoMessage()    {}
func (*SeekNewest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{1}
}

func (m *SeekNewest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SeekNewest.Unmarshal(m, b)
}
func (m *SeekNewest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SeekNewest.Marshal(b, m, deterministic)
}
func (m *SeekNewest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeekNewest.Merge(m, src)
}
func (m *SeekNewest) XXX_Size() int {
	return xxx_messageInfo_SeekNewest.Size(m)
}
func (m *SeekNewest) XXX_DiscardUnknown() {
	xxx_messageInfo_SeekNewest.DiscardUnknown(m)
}

var xxx_messageInfo_SeekNewest proto.InternalMessageInfo

type SeekOldest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SeekOldest) Reset()         { *m = SeekOldest{} }
func (m *SeekOldest) String() string { return proto.CompactTextString(m) }
func (*SeekOldest) ProtoMessage()    {}
func (*SeekOldest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{2}
}

func (m *SeekOldest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SeekOldest.Unmarshal(m, b)
}
func (m *SeekOldest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SeekOldest.Marshal(b, m, deterministic)
}
func (m *SeekOldest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeekOldest.Merge(m, src)
}
func (m *SeekOldest) XXX_Size() int {
	return xxx_messageInfo_SeekOldest.Size(m)
}
func (m *SeekOldest) XXX_DiscardUnknown() {
	xxx_messageInfo_SeekOldest.DiscardUnknown(m)
}

var xxx_messageInfo_SeekOldest proto.InternalMessageInfo

type SeekSpecified struct {
	Number               uint64   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SeekSpecified) Reset()         { *m = SeekSpecified{} }
func (m *SeekSpecified) String() string { return proto.CompactTextString(m) }
func (*SeekSpecified) ProtoMessage()    {}
func (*SeekSpecified) Descriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{3}
}

func (m *SeekSpecified) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SeekSpecified.Unmarshal(m, b)
}
func (m *SeekSpecified) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SeekSpecified.Marshal(b, m, deterministic)
}
func (m *SeekSpecified) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeekSpecified.Merge(m, src)
}
func (m *SeekSpecified) XXX_Size() int {
	return xxx_messageInfo_SeekSpecified.Size(m)
}
func (m *SeekSpecified) XXX_DiscardUnknown() {
	xxx_messageInfo_SeekSpecified.DiscardUnknown(m)
}

var xxx_messageInfo_SeekSpecified proto.InternalMessageInfo

func (m *SeekSpecified) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

// SeekNextCommit refers to the next block that will be committed
type SeekNextCommit struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SeekNextCommit) Reset()         { *m = SeekNextCommit{} }
func (m *SeekNextCommit) String() string { return proto.CompactTextString(m) }
func (*SeekNextCommit) ProtoMessage()    {}
func (*SeekNextCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{4}
}

func (m *SeekNextCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SeekNextCommit.Unmarshal(m, b)
}
func (m *SeekNextCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SeekNextCommit.Marshal(b, m, deterministic)
}
func (m *SeekNextCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeekNextCommit.Merge(m, src)
}
func (m *SeekNextCommit) XXX_Size() int {
	return xxx_messageInfo_SeekNextCommit.Size(m)
}
func (m *SeekNextCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_SeekNextCommit.DiscardUnknown(m)
}

var xxx_messageInfo_SeekNextCommit proto.InternalMessageInfo

type SeekPosition struct {
	// Types that are valid to be assigned to Type:
	//	*SeekPosition_Newest
	//	*SeekPosition_Oldest
	//	*SeekPosition_Specified
	//	*SeekPosition_NextCommit
	Type                 isSeekPosition_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SeekPosition) Reset()         { *m = SeekPosition{} }
func (m *SeekPosition) String() string { return proto.CompactTextString(m) }
func (*SeekPosition) ProtoMessage()    {}
func (*SeekPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{5}
}

func (m *SeekPosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SeekPosition.Unmarshal(m, b)
}
func (m *SeekPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SeekPosition.Marshal(b, m, deterministic)
}
func (m *SeekPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeekPosition.Merge(m, src)
}
func (m *SeekPosition) XXX_Size() int {
	return xxx_messageInfo_SeekPosition.Size(m)
}
func (m *SeekPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_SeekPosition.DiscardUnknown(m)
}

var xxx_messageInfo_SeekPosition proto.InternalMessageInfo

type isSeekPosition_Type interface {
	isSeekPosition_Type()
}

type SeekPosition_Newest struct {
	Newest *SeekNewest `protobuf:"bytes,1,opt,name=newest,proto3,oneof"`
}

type SeekPosition_Oldest struct {
	Oldest *SeekOldest `protobuf:"bytes,2,opt,name=oldest,proto3,oneof"`
}

type SeekPosition_Specified struct {
	Specified *SeekSpecified `protobuf:"bytes,3,opt,name=specified,proto3,oneof"`
}

type SeekPosition_NextCommit struct {
	NextCommit *SeekNextCommit `protobuf:"bytes,4,opt,name=next_commit,json=nextCommit,proto3,oneof"`
}

func (*SeekPosition_Newest) isSeekPosition_Type() {}

func (*SeekPosition_Oldest) isSeekPosition_Type() {}

func (*SeekPosition_Specified) isSeekPosition_Type() {}

func (*SeekPosition_NextCommit) isSeekPosition_Type() {}

func (m *SeekPosition) GetType() isSeekPosition_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SeekPosition) GetNewest() *SeekNewest {
	if x, ok := m.GetType().(*SeekPosition_Newest); ok {
		return x.Newest
	}
	return nil
}

func (m *SeekPosition) GetOldest() *SeekOldest {
	if x, ok := m.GetType().(*SeekPosition_Oldest); ok {
		return x.Oldest
	}
	return nil
}

func (m *SeekPosition) GetSpecified() *SeekSpecified {
	if x, ok := m.GetType().(*SeekPosition_Specified); ok {
		return x.Specified
	}
	return nil
}

func (m *SeekPosition) GetNextCommit() *SeekNextCommit {
	if x, ok := m.GetType().(*SeekPosition_NextCommit); ok {
		return x.NextCommit
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SeekPosition) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SeekPosition_Newest)(nil),
		(*SeekPosition_Oldest)(nil),
		(*SeekPosition_Specified)(nil),
		(*SeekPosition_NextCommit)(nil),
	}
}

// SeekInfo specifies the range of requested blocks to return
// If the start position is not found, an error is immediately returned
// Otherwise, blocks are returned until a missing block is encountered, then behavior is dictated
// by the SeekBehavior specified.
type SeekInfo struct {
	Start                *SeekPosition              `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	Stop                 *SeekPosition              `protobuf:"bytes,2,opt,name=stop,proto3" json:"stop,omitempty"`
	Behavior             SeekInfo_SeekBehavior      `protobuf:"varint,3,opt,name=behavior,proto3,enum=orderer.SeekInfo_SeekBehavior" json:"behavior,omitempty"`
	ErrorResponse        SeekInfo_SeekErrorResponse `protobuf:"varint,4,opt,name=error_response,json=errorResponse,proto3,enum=orderer.SeekInfo_SeekErrorResponse" json:"error_response,omitempty"`
	ContentType          SeekInfo_SeekContentType   `protobuf:"varint,5,opt,name=content_type,json=contentType,proto3,enum=orderer.SeekInfo_SeekContentType" json:"content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SeekInfo) Reset()         { *m = SeekInfo{} }
func (m *SeekInfo) String() string { return proto.CompactTextString(m) }
func (*SeekInfo) ProtoMessage()    {}
func (*SeekInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{6}
}

func (m *SeekInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SeekInfo.Unmarshal(m, b)
}
func (m *SeekInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SeekInfo.Marshal(b, m, deterministic)
}
func (m *SeekInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeekInfo.Merge(m, src)
}
func (m *SeekInfo) XXX_Size() int {
	return xxx_messageInfo_SeekInfo.Size(m)
}
func (m *SeekInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SeekInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SeekInfo proto.InternalMessageInfo

func (m *SeekInfo) GetStart() *SeekPosition {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *SeekInfo) GetStop() *SeekPosition {
	if m != nil {
		return m.Stop
	}
	return nil
}

func (m *SeekInfo) GetBehavior() SeekInfo_SeekBehavior {
	if m != nil {
		return m.Behavior
	}
	return SeekInfo_BLOCK_UNTIL_READY
}

func (m *SeekInfo) GetErrorResponse() SeekInfo_SeekErrorResponse {
	if m != nil {
		return m.ErrorResponse
	}
	return SeekInfo_STRICT
}

func (m *SeekInfo) GetContentType() SeekInfo_SeekContentType {
	if m != nil {
		return m.ContentType
	}
	return SeekInfo_BLOCK
}

type DeliverResponse struct {
	// Types that are valid to be assigned to Type:
	//	*DeliverResponse_Status
	//	*DeliverResponse_Block
	Type                 isDeliverResponse_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DeliverResponse) Reset()         { *m = DeliverResponse{} }
func (m *DeliverResponse) String() string { return proto.CompactTextString(m) }
func (*DeliverResponse) ProtoMessage()    {}
func (*DeliverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79fce58dd8d86d62, []int{7}
}

func (m *DeliverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliverResponse.Unmarshal(m, b)
}
func (m *DeliverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliverResponse.Marshal(b, m, deterministic)
}
func (m *DeliverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliverResponse.Merge(m, src)
}
func (m *DeliverResponse) XXX_Size() int {
	return xxx_messageInfo_DeliverResponse.Size(m)
}
func (m *DeliverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeliverResponse proto.InternalMessageInfo

type isDeliverResponse_Type interface {
	isDeliverResponse_Type()
}

type DeliverResponse_Status struct {
	Status common.Status `protobuf:"varint,1,opt,name=status,proto3,enum=common.Status,oneof"`
}

type DeliverResponse_Block struct {
	Block *common.Block `protobuf:"bytes,2,opt,name=block,proto3,oneof"`
}

func (*DeliverResponse_Status) isDeliverResponse_Type() {}

func (*DeliverResponse_Block) isDeliverResponse_Type() {}

func (m *DeliverResponse) GetType() isDeliverResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DeliverResponse) GetStatus() common.Status {
	if x, ok := m.GetType().(*DeliverResponse_Status); ok {
		return x.Status
	}
	return common.Status_UNKNOWN
}

func (m *DeliverResponse) GetBlock() *common.Block {
	if x, ok := m.GetType().(*DeliverResponse_Block); ok {
		return x.Block
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DeliverResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DeliverResponse_Status)(nil),
		(*DeliverResponse_Block)(nil),
	}
}

func init() {
	proto.RegisterEnum("orderer.SeekInfo_SeekBehavior", SeekInfo_SeekBehavior_name, SeekInfo_SeekBehavior_value)
	proto.RegisterEnum("orderer.SeekInfo_SeekErrorResponse", SeekInfo_SeekErrorResponse_name, SeekInfo_SeekErrorResponse_value)
	proto.RegisterEnum("orderer.SeekInfo_SeekContentType", SeekInfo_SeekContentType_name, SeekInfo_SeekContentType_value)
	proto.RegisterType((*BroadcastResponse)(nil), "orderer.BroadcastResponse")
	proto.RegisterType((*SeekNewest)(nil), "orderer.SeekNewest")
	proto.RegisterType((*SeekOldest)(nil), "orderer.SeekOldest")
	proto.RegisterType((*SeekSpecified)(nil), "orderer.SeekSpecified")
	proto.RegisterType((*SeekNextCommit)(nil), "orderer.SeekNextCommit")
	proto.RegisterType((*SeekPosition)(nil), "orderer.SeekPosition")
	proto.RegisterType((*SeekInfo)(nil), "orderer.SeekInfo")
	proto.RegisterType((*DeliverResponse)(nil), "orderer.DeliverResponse")
}

func init() {
	proto.RegisterFile("orderer/ab.proto", fileDescriptor_79fce58dd8d86d62)
}

var fileDescriptor_79fce58dd8d86d62 = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdf, 0x6e, 0xda, 0x4a,
	0x10, 0xc6, 0xed, 0x13, 0x20, 0x61, 0x48, 0xc0, 0xd9, 0x28, 0x39, 0x28, 0x17, 0x47, 0x39, 0xae,
	0xd2, 0x52, 0x55, 0x81, 0x94, 0x4a, 0x95, 0x1a, 0xb5, 0x17, 0x18, 0x4c, 0x71, 0x1b, 0x85, 0x6a,
	0x71, 0xd5, 0x3f, 0x37, 0x96, 0x6d, 0x16, 0xe2, 0x06, 0xbc, 0xd6, 0x7a, 0x93, 0x26, 0xcf, 0xd0,
	0xc7, 0xeb, 0x4b, 0xf4, 0x31, 0xaa, 0x5d, 0xaf, 0x21, 0x21, 0x69, 0xaf, 0xf0, 0x8c, 0x7f, 0xdf,
	0xcc, 0x37, 0xcb, 0xac, 0xc1, 0xa0, 0x6c, 0x4c, 0x18, 0x61, 0x2d, 0x3f, 0x68, 0x26, 0x8c, 0x72,
	0x8a, 0xd6, 0x55, 0x66, 0x7f, 0x27, 0xa4, 0xf3, 0x39, 0x8d, 0x5b, 0xd9, 0x4f, 0xf6, 0xd6, 0x1c,
	0xc2, 0xb6, 0xc5, 0xa8, 0x3f, 0x0e, 0xfd, 0x94, 0x63, 0x92, 0x26, 0x34, 0x4e, 0x09, 0x7a, 0x0c,
	0xa5, 0x94, 0xfb, 0xfc, 0x32, 0xad, 0xeb, 0x07, 0x7a, 0xa3, 0xda, 0xae, 0x36, 0x95, 0x66, 0x24,
	0xb3, 0x58, 0xbd, 0x45, 0x08, 0x0a, 0x51, 0x3c, 0xa1, 0xf5, 0x7f, 0x0e, 0xf4, 0x46, 0x19, 0xcb,
	0x67, 0x73, 0x13, 0x60, 0x44, 0xc8, 0xc5, 0x19, 0xf9, 0x4e, 0x52, 0x9e, 0x47, 0xc3, 0xd9, 0x58,
	0x44, 0x4f, 0x60, 0x4b, 0x44, 0xa3, 0x84, 0x84, 0xd1, 0x24, 0x22, 0x63, 0xb4, 0x07, 0xa5, 0xf8,
	0x72, 0x1e, 0x10, 0x26, 0x1b, 0x15, 0xb0, 0x8a, 0x4c, 0x03, 0xaa, 0x59, 0x91, 0x6b, 0xde, 0xa5,
	0xf3, 0x79, 0xc4, 0xcd, 0x5f, 0x3a, 0x6c, 0x8a, 0xd4, 0x07, 0x9a, 0x46, 0x3c, 0xa2, 0x31, 0x3a,
	0x82, 0x52, 0x2c, 0x7b, 0x48, 0x69, 0xa5, 0xbd, 0xd3, 0x54, 0x73, 0x36, 0x97, 0xed, 0x07, 0x1a,
	0x56, 0x90, 0xc0, 0xa9, 0x34, 0x21, 0xcd, 0xae, 0xe2, 0x99, 0x3f, 0x81, 0x67, 0x10, 0x7a, 0x09,
	0xe5, 0x34, 0x77, 0x59, 0x5f, 0x93, 0x8a, 0xbd, 0x3b, 0x8a, 0xc5, 0x0c, 0x03, 0x0d, 0x2f, 0x51,
	0x74, 0x02, 0x95, 0x98, 0x5c, 0x73, 0x2f, 0x94, 0xae, 0xeb, 0x05, 0xa9, 0xfc, 0x77, 0xc5, 0x5a,
	0x3e, 0xd4, 0x40, 0xc3, 0x10, 0x2f, 0x22, 0xab, 0x04, 0x05, 0xf7, 0x26, 0x21, 0xe6, 0xcf, 0x35,
	0xd8, 0x10, 0xa0, 0x13, 0x4f, 0x28, 0x7a, 0x06, 0xc5, 0x94, 0xfb, 0x2c, 0x9f, 0x72, 0xf7, 0x4e,
	0xa9, 0xfc, 0x30, 0x70, 0xc6, 0xa0, 0xa7, 0x50, 0x48, 0x39, 0x4d, 0xd4, 0x88, 0x7f, 0x60, 0x25,
	0x82, 0x4e, 0x60, 0x23, 0x20, 0xe7, 0xfe, 0x55, 0x44, 0x99, 0x9c, 0xaf, 0xda, 0xfe, 0xef, 0x0e,
	0x2e, 0x9a, 0xcb, 0x07, 0x4b, 0x51, 0x78, 0xc1, 0xa3, 0x77, 0x50, 0x25, 0x8c, 0x51, 0xe6, 0x31,
	0xb5, 0x30, 0x72, 0xce, 0x6a, 0xfb, 0xd1, 0xc3, 0x15, 0x6c, 0xc1, 0xe6, 0xbb, 0x85, 0xb7, 0xc8,
	0xed, 0x10, 0xf5, 0x60, 0x33, 0xa4, 0x31, 0x27, 0x31, 0xf7, 0xf8, 0x4d, 0x42, 0xea, 0x45, 0x59,
	0xe9, 0xff, 0x87, 0x2b, 0x75, 0x33, 0x52, 0x9c, 0x12, 0xae, 0x84, 0xcb, 0xc0, 0x7c, 0x9d, 0x2d,
	0x47, 0xee, 0x15, 0xed, 0xc2, 0xb6, 0x75, 0x3a, 0xec, 0xbe, 0xf7, 0x3e, 0x9e, 0xb9, 0xce, 0xa9,
	0x87, 0xed, 0x4e, 0xef, 0x8b, 0xa1, 0x89, 0x74, 0xbf, 0xe3, 0x9c, 0x7a, 0x4e, 0xdf, 0x3b, 0x1b,
	0xba, 0x2a, 0xad, 0x9b, 0xc7, 0xb0, 0x7d, 0xcf, 0x27, 0x02, 0x28, 0x8d, 0x5c, 0xec, 0x74, 0x5d,
	0x43, 0x43, 0x35, 0xa8, 0x58, 0xf6, 0xc8, 0xf5, 0xec, 0x7e, 0x7f, 0x88, 0x5d, 0x43, 0x37, 0x9f,
	0x43, 0x6d, 0xc5, 0x0f, 0x2a, 0x43, 0x51, 0xb6, 0x34, 0x34, 0xb4, 0x03, 0xb5, 0x81, 0xdd, 0xe9,
	0xd9, 0xd8, 0xfb, 0xe4, 0xb8, 0x03, 0x6f, 0xe4, 0xbc, 0x35, 0x74, 0xf3, 0x1b, 0xd4, 0x7a, 0x64,
	0x16, 0x5d, 0x91, 0x65, 0x8b, 0xc6, 0xdf, 0xaf, 0x99, 0x58, 0x47, 0x75, 0xd1, 0x0e, 0xa1, 0x18,
	0xcc, 0x68, 0x78, 0xa1, 0xfe, 0xd9, 0xad, 0x1c, 0xb4, 0x44, 0x72, 0xa0, 0xe1, 0xec, 0x6d, 0xbe,
	0x41, 0xed, 0x1f, 0x3a, 0xd4, 0x3a, 0x9c, 0xce, 0xa3, 0x70, 0x71, 0xb7, 0xd1, 0x1b, 0x28, 0x2f,
	0x03, 0x23, 0x2f, 0x60, 0xc7, 0x57, 0x64, 0x46, 0x13, 0xb2, 0xbf, 0xbf, 0x38, 0xf1, 0x7b, 0x9f,
	0x83, 0x86, 0x7e, 0xac, 0xa3, 0x57, 0xb0, 0xae, 0xec, 0x3f, 0x20, 0xae, 0x2f, 0xc4, 0x2b, 0x23,
	0x0a, 0xa9, 0xf5, 0x19, 0x0e, 0x29, 0x9b, 0x36, 0xcf, 0x6f, 0x12, 0xc2, 0x66, 0x64, 0x3c, 0x25,
	0xac, 0x39, 0xf1, 0x03, 0x16, 0x85, 0xd9, 0x27, 0x28, 0xcd, 0xc5, 0x5f, 0x5b, 0xd3, 0x88, 0x9f,
	0x5f, 0x06, 0xa2, 0x7c, 0xeb, 0x16, 0xdd, 0xca, 0xe8, 0xa3, 0x8c, 0x3e, 0x9a, 0xd2, 0x96, 0x12,
	0x04, 0x25, 0x99, 0x7a, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x4e, 0x38, 0xe9, 0xf5, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AtomicBroadcastClient is the client API for AtomicBroadcast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AtomicBroadcastClient interface {
	// broadcast receives a reply of Acknowledgement for each common.Envelope in order, indicating success or type of failure
	Broadcast(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_BroadcastClient, error)
	// deliver first requires an Envelope of type DELIVER_SEEK_INFO with Payload data as a mashaled SeekInfo message, then a stream of block replies is received.
	Deliver(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_DeliverClient, error)
}

type atomicBroadcastClient struct {
	cc grpc.ClientConnInterface
}

func NewAtomicBroadcastClient(cc grpc.ClientConnInterface) AtomicBroadcastClient {
	return &atomicBroadcastClient{cc}
}

func (c *atomicBroadcastClient) Broadcast(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AtomicBroadcast_serviceDesc.Streams[0], "/orderer.AtomicBroadcast/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &atomicBroadcastBroadcastClient{stream}
	return x, nil
}

type AtomicBroadcast_BroadcastClient interface {
	Send(*common.Envelope) error
	Recv() (*BroadcastResponse, error)
	grpc.ClientStream
}

type atomicBroadcastBroadcastClient struct {
	grpc.ClientStream
}

func (x *atomicBroadcastBroadcastClient) Send(m *common.Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *atomicBroadcastBroadcastClient) Recv() (*BroadcastResponse, error) {
	m := new(BroadcastResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *atomicBroadcastClient) Deliver(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_DeliverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AtomicBroadcast_serviceDesc.Streams[1], "/orderer.AtomicBroadcast/Deliver", opts...)
	if err != nil {
		return nil, err
	}
	x := &atomicBroadcastDeliverClient{stream}
	return x, nil
}

type AtomicBroadcast_DeliverClient interface {
	Send(*common.Envelope) error
	Recv() (*DeliverResponse, error)
	grpc.ClientStream
}

type atomicBroadcastDeliverClient struct {
	grpc.ClientStream
}

func (x *atomicBroadcastDeliverClient) Send(m *common.Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *atomicBroadcastDeliverClient) Recv() (*DeliverResponse, error) {
	m := new(DeliverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AtomicBroadcastServer is the server API for AtomicBroadcast service.
type AtomicBroadcastServer interface {
	// broadcast receives a reply of Acknowledgement for each common.Envelope in order, indicating success or type of failure
	Broadcast(AtomicBroadcast_BroadcastServer) error
	// deliver first requires an Envelope of type DELIVER_SEEK_INFO with Payload data as a mashaled SeekInfo message, then a stream of block replies is received.
	Deliver(AtomicBroadcast_DeliverServer) error
}

// UnimplementedAtomicBroadcastServer can be embedded to have forward compatible implementations.
type UnimplementedAtomicBroadcastServer struct {
}

func (*UnimplementedAtomicBroadcastServer) Broadcast(srv AtomicBroadcast_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (*UnimplementedAtomicBroadcastServer) Deliver(srv AtomicBroadcast_DeliverServer) error {
	return status.Errorf(codes.Unimplemented, "method Deliver not implemented")
}

func RegisterAtomicBroadcastServer(s *grpc.Server, srv AtomicBroadcastServer) {
	s.RegisterService(&_AtomicBroadcast_serviceDesc, srv)
}

func _AtomicBroadcast_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AtomicBroadcastServer).Broadcast(&atomicBroadcastBroadcastServer{stream})
}

type AtomicBroadcast_BroadcastServer interface {
	Send(*BroadcastResponse) error
	Recv() (*common.Envelope, error)
	grpc.ServerStream
}

type atomicBroadcastBroadcastServer struct {
	grpc.ServerStream
}

func (x *atomicBroadcastBroadcastServer) Send(m *BroadcastResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *atomicBroadcastBroadcastServer) Recv() (*common.Envelope, error) {
	m := new(common.Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AtomicBroadcast_Deliver_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AtomicBroadcastServer).Deliver(&atomicBroadcastDeliverServer{stream})
}

type AtomicBroadcast_DeliverServer interface {
	Send(*DeliverResponse) error
	Recv() (*common.Envelope, error)
	grpc.ServerStream
}

type atomicBroadcastDeliverServer struct {
	grpc.ServerStream
}

func (x *atomicBroadcastDeliverServer) Send(m *DeliverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *atomicBroadcastDeliverServer) Recv() (*common.Envelope, error) {
	m := new(common.Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AtomicBroadcast_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderer.AtomicBroadcast",
	HandlerType: (*AtomicBroadcastServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Broadcast",
			Handler:       _AtomicBroadcast_Broadcast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Deliver",
			Handler:       _AtomicBroadcast_Deliver_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orderer/ab.proto",
}
