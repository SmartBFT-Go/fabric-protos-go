// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/cluster.proto

package orderer

import (
	context "context"
	fmt "fmt"
	common "github.com/SmartBFT-Go/fabric-protos-go/common"
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

// StepRequest wraps a message that is sent to a cluster member.
type StepRequest struct {
	// Types that are valid to be assigned to Payload:
	//	*StepRequest_ConsensusRequest
	//	*StepRequest_SubmitRequest
	Payload              isStepRequest_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StepRequest) Reset()         { *m = StepRequest{} }
func (m *StepRequest) String() string { return proto.CompactTextString(m) }
func (*StepRequest) ProtoMessage()    {}
func (*StepRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3b50707fd3a71f2, []int{0}
}

func (m *StepRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepRequest.Unmarshal(m, b)
}
func (m *StepRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepRequest.Marshal(b, m, deterministic)
}
func (m *StepRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepRequest.Merge(m, src)
}
func (m *StepRequest) XXX_Size() int {
	return xxx_messageInfo_StepRequest.Size(m)
}
func (m *StepRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StepRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StepRequest proto.InternalMessageInfo

type isStepRequest_Payload interface {
	isStepRequest_Payload()
}

type StepRequest_ConsensusRequest struct {
	ConsensusRequest *ConsensusRequest `protobuf:"bytes,1,opt,name=consensus_request,json=consensusRequest,proto3,oneof"`
}

type StepRequest_SubmitRequest struct {
	SubmitRequest *SubmitRequest `protobuf:"bytes,2,opt,name=submit_request,json=submitRequest,proto3,oneof"`
}

func (*StepRequest_ConsensusRequest) isStepRequest_Payload() {}

func (*StepRequest_SubmitRequest) isStepRequest_Payload() {}

func (m *StepRequest) GetPayload() isStepRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *StepRequest) GetConsensusRequest() *ConsensusRequest {
	if x, ok := m.GetPayload().(*StepRequest_ConsensusRequest); ok {
		return x.ConsensusRequest
	}
	return nil
}

func (m *StepRequest) GetSubmitRequest() *SubmitRequest {
	if x, ok := m.GetPayload().(*StepRequest_SubmitRequest); ok {
		return x.SubmitRequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StepRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StepRequest_ConsensusRequest)(nil),
		(*StepRequest_SubmitRequest)(nil),
	}
}

// StepResponse is a message received from a cluster member.
type StepResponse struct {
	// Types that are valid to be assigned to Payload:
	//	*StepResponse_SubmitRes
	Payload              isStepResponse_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StepResponse) Reset()         { *m = StepResponse{} }
func (m *StepResponse) String() string { return proto.CompactTextString(m) }
func (*StepResponse) ProtoMessage()    {}
func (*StepResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3b50707fd3a71f2, []int{1}
}

func (m *StepResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepResponse.Unmarshal(m, b)
}
func (m *StepResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepResponse.Marshal(b, m, deterministic)
}
func (m *StepResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepResponse.Merge(m, src)
}
func (m *StepResponse) XXX_Size() int {
	return xxx_messageInfo_StepResponse.Size(m)
}
func (m *StepResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StepResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StepResponse proto.InternalMessageInfo

type isStepResponse_Payload interface {
	isStepResponse_Payload()
}

type StepResponse_SubmitRes struct {
	SubmitRes *SubmitResponse `protobuf:"bytes,1,opt,name=submit_res,json=submitRes,proto3,oneof"`
}

func (*StepResponse_SubmitRes) isStepResponse_Payload() {}

func (m *StepResponse) GetPayload() isStepResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *StepResponse) GetSubmitRes() *SubmitResponse {
	if x, ok := m.GetPayload().(*StepResponse_SubmitRes); ok {
		return x.SubmitRes
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StepResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StepResponse_SubmitRes)(nil),
	}
}

// ConsensusRequest is a consensus specific message sent to a cluster member.
type ConsensusRequest struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata             []byte   `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsensusRequest) Reset()         { *m = ConsensusRequest{} }
func (m *ConsensusRequest) String() string { return proto.CompactTextString(m) }
func (*ConsensusRequest) ProtoMessage()    {}
func (*ConsensusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3b50707fd3a71f2, []int{2}
}

func (m *ConsensusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusRequest.Unmarshal(m, b)
}
func (m *ConsensusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusRequest.Marshal(b, m, deterministic)
}
func (m *ConsensusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusRequest.Merge(m, src)
}
func (m *ConsensusRequest) XXX_Size() int {
	return xxx_messageInfo_ConsensusRequest.Size(m)
}
func (m *ConsensusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusRequest proto.InternalMessageInfo

func (m *ConsensusRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *ConsensusRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ConsensusRequest) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// SubmitRequest wraps a transaction to be sent for ordering.
type SubmitRequest struct {
	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// last_validation_seq denotes the last
	// configuration sequence at which the
	// sender validated this message.
	LastValidationSeq uint64 `protobuf:"varint,2,opt,name=last_validation_seq,json=lastValidationSeq,proto3" json:"last_validation_seq,omitempty"`
	// content is the fabric transaction
	// that is forwarded to the cluster member.
	Payload              *common.Envelope `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SubmitRequest) Reset()         { *m = SubmitRequest{} }
func (m *SubmitRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitRequest) ProtoMessage()    {}
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3b50707fd3a71f2, []int{3}
}

func (m *SubmitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitRequest.Unmarshal(m, b)
}
func (m *SubmitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitRequest.Marshal(b, m, deterministic)
}
func (m *SubmitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitRequest.Merge(m, src)
}
func (m *SubmitRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitRequest.Size(m)
}
func (m *SubmitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitRequest proto.InternalMessageInfo

func (m *SubmitRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *SubmitRequest) GetLastValidationSeq() uint64 {
	if m != nil {
		return m.LastValidationSeq
	}
	return 0
}

func (m *SubmitRequest) GetPayload() *common.Envelope {
	if m != nil {
		return m.Payload
	}
	return nil
}

// SubmitResponse returns a success
// or failure status to the sender.
type SubmitResponse struct {
	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// Status code, which may be used to programatically respond to success/failure.
	Status common.Status `protobuf:"varint,2,opt,name=status,proto3,enum=common.Status" json:"status,omitempty"`
	// Info string which may contain additional information about the returned status.
	Info                 string   `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitResponse) Reset()         { *m = SubmitResponse{} }
func (m *SubmitResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitResponse) ProtoMessage()    {}
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3b50707fd3a71f2, []int{4}
}

func (m *SubmitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitResponse.Unmarshal(m, b)
}
func (m *SubmitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitResponse.Marshal(b, m, deterministic)
}
func (m *SubmitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitResponse.Merge(m, src)
}
func (m *SubmitResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitResponse.Size(m)
}
func (m *SubmitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitResponse proto.InternalMessageInfo

func (m *SubmitResponse) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *SubmitResponse) GetStatus() common.Status {
	if m != nil {
		return m.Status
	}
	return common.Status_UNKNOWN
}

func (m *SubmitResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*StepRequest)(nil), "orderer.StepRequest")
	proto.RegisterType((*StepResponse)(nil), "orderer.StepResponse")
	proto.RegisterType((*ConsensusRequest)(nil), "orderer.ConsensusRequest")
	proto.RegisterType((*SubmitRequest)(nil), "orderer.SubmitRequest")
	proto.RegisterType((*SubmitResponse)(nil), "orderer.SubmitResponse")
}

func init() { proto.RegisterFile("orderer/cluster.proto", fileDescriptor_e3b50707fd3a71f2) }

var fileDescriptor_e3b50707fd3a71f2 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x8d, 0x1a, 0x13, 0x57, 0x93, 0x44, 0x38, 0x9b, 0xa6, 0x75, 0x7d, 0x2a, 0x82, 0x96, 0x50,
	0x88, 0x54, 0xdc, 0x43, 0x7b, 0x2b, 0x38, 0x14, 0x7c, 0x5e, 0x41, 0x29, 0xbd, 0x98, 0x95, 0x34,
	0x96, 0x05, 0xd2, 0xae, 0xbc, 0xbb, 0x0a, 0xe4, 0x03, 0xfa, 0x25, 0xfd, 0xd1, 0xa2, 0xdd, 0x95,
	0xe4, 0x38, 0x90, 0x93, 0x34, 0xef, 0xcd, 0xbc, 0x79, 0xb3, 0x33, 0x70, 0x23, 0x64, 0x8e, 0x12,
	0x65, 0x9c, 0x55, 0xad, 0xd2, 0x28, 0xa3, 0x46, 0x0a, 0x2d, 0xc8, 0xd4, 0xc1, 0x8b, 0xeb, 0x4c,
	0xd4, 0xb5, 0xe0, 0xb1, 0xfd, 0x58, 0x36, 0xfc, 0xe7, 0xc1, 0x79, 0xa2, 0xb1, 0xa1, 0xb8, 0x6f,
	0x51, 0x69, 0xb2, 0x86, 0xab, 0x4c, 0x70, 0x85, 0x5c, 0xb5, 0x6a, 0x23, 0x2d, 0x38, 0xf7, 0x3e,
	0x78, 0xb7, 0xe7, 0xcb, 0xf7, 0x91, 0x53, 0x8a, 0xee, 0xfb, 0x0c, 0x57, 0xb5, 0x3e, 0xa1, 0xb3,
	0xec, 0x08, 0x23, 0x3f, 0x20, 0x50, 0x6d, 0x5a, 0x97, 0x7a, 0x90, 0x79, 0x65, 0x64, 0xde, 0x0e,
	0x32, 0x89, 0xa1, 0x47, 0x8d, 0x4b, 0x75, 0x08, 0xac, 0x7c, 0x98, 0x36, 0xec, 0xb1, 0x12, 0x2c,
	0x0f, 0x13, 0xb8, 0xb0, 0x26, 0x55, 0xd3, 0xb5, 0x21, 0xdf, 0x01, 0x06, 0x6d, 0xe5, 0xec, 0xbd,
	0x7b, 0xa6, 0x6b, 0x93, 0xd7, 0x27, 0xd4, 0xef, 0x85, 0xd5, 0xa1, 0x68, 0x0a, 0xb3, 0xe3, 0x41,
	0xc8, 0x1c, 0xa6, 0xd9, 0x8e, 0x71, 0x8e, 0x95, 0x51, 0xf5, 0x69, 0x1f, 0x76, 0x8c, 0x2b, 0x34,
	0x73, 0x5c, 0xd0, 0x3e, 0x24, 0x0b, 0x78, 0x5d, 0xa3, 0x66, 0x39, 0xd3, 0x6c, 0x7e, 0x6a, 0xa8,
	0x21, 0x0e, 0xff, 0x7a, 0x70, 0xf9, 0x64, 0xcc, 0x17, 0x3a, 0x44, 0x70, 0x5d, 0x31, 0xa5, 0x37,
	0x0f, 0xac, 0x2a, 0x73, 0xa6, 0x4b, 0xc1, 0x37, 0x0a, 0xf7, 0xa6, 0xdb, 0x84, 0x5e, 0x75, 0xd4,
	0xaf, 0x81, 0x49, 0x70, 0x4f, 0x3e, 0x8f, 0x8e, 0x4e, 0xcd, 0x0b, 0xcc, 0x22, 0xb7, 0xda, 0x9f,
	0xfc, 0x01, 0x2b, 0xd1, 0xe0, 0xe0, 0x31, 0xdc, 0x42, 0xf0, 0xf4, 0x55, 0x5e, 0xf0, 0xf1, 0x09,
	0xce, 0x94, 0x66, 0xba, 0x55, 0xa6, 0x75, 0xb0, 0x0c, 0x7a, 0xd9, 0xc4, 0xa0, 0xd4, 0xb1, 0x84,
	0xc0, 0xa4, 0xe4, 0x5b, 0x61, 0x9a, 0xfb, 0xd4, 0xfc, 0x2f, 0x57, 0x30, 0xbd, 0xb7, 0xd7, 0x47,
	0xbe, 0xc1, 0xa4, 0xdb, 0x19, 0x79, 0x33, 0xee, 0x65, 0xbc, 0xb3, 0xc5, 0xcd, 0x11, 0x6a, 0x5d,
	0xdd, 0x7a, 0x5f, 0xbc, 0xd5, 0x6f, 0xf8, 0x28, 0x64, 0x11, 0xed, 0x1e, 0x1b, 0x94, 0x15, 0xe6,
	0x05, 0xca, 0x68, 0xcb, 0x52, 0x59, 0x66, 0xf6, 0x64, 0x55, 0x5f, 0xf9, 0x27, 0x2e, 0x4a, 0xbd,
	0x6b, 0xd3, 0xce, 0x5e, 0x7c, 0x90, 0x1d, 0xdb, 0xec, 0x3b, 0x9b, 0x7d, 0x57, 0x88, 0xd8, 0x15,
	0xa4, 0x67, 0x06, 0xfa, 0xfa, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xed, 0x23, 0xd0, 0x2a, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterClient interface {
	// Step passes an implementation-specific message to another cluster member.
	Step(ctx context.Context, opts ...grpc.CallOption) (Cluster_StepClient, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) Step(ctx context.Context, opts ...grpc.CallOption) (Cluster_StepClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Cluster_serviceDesc.Streams[0], "/orderer.Cluster/Step", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterStepClient{stream}
	return x, nil
}

type Cluster_StepClient interface {
	Send(*StepRequest) error
	Recv() (*StepResponse, error)
	grpc.ClientStream
}

type clusterStepClient struct {
	grpc.ClientStream
}

func (x *clusterStepClient) Send(m *StepRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterStepClient) Recv() (*StepResponse, error) {
	m := new(StepResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClusterServer is the server API for Cluster service.
type ClusterServer interface {
	// Step passes an implementation-specific message to another cluster member.
	Step(Cluster_StepServer) error
}

// UnimplementedClusterServer can be embedded to have forward compatible implementations.
type UnimplementedClusterServer struct {
}

func (*UnimplementedClusterServer) Step(srv Cluster_StepServer) error {
	return status.Errorf(codes.Unimplemented, "method Step not implemented")
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_Step_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterServer).Step(&clusterStepServer{stream})
}

type Cluster_StepServer interface {
	Send(*StepResponse) error
	Recv() (*StepRequest, error)
	grpc.ServerStream
}

type clusterStepServer struct {
	grpc.ServerStream
}

func (x *clusterStepServer) Send(m *StepResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterStepServer) Recv() (*StepRequest, error) {
	m := new(StepRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderer.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Step",
			Handler:       _Cluster_Step_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orderer/cluster.proto",
}
