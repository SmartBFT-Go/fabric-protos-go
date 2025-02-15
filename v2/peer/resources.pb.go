// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/resources.proto

package peer

import (
	fmt "fmt"
	common "github.com/SmartBFT-Go/fabric-protos-go/v2/common"
	proto "github.com/golang/protobuf/proto"
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

// ChaincodeIdentifier identifies a piece of chaincode.  For a peer to accept invocations of
// this chaincode, the hash of the installed code must match, as must the version string
// included with the install command.
type ChaincodeIdentifier struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeIdentifier) Reset()         { *m = ChaincodeIdentifier{} }
func (m *ChaincodeIdentifier) String() string { return proto.CompactTextString(m) }
func (*ChaincodeIdentifier) ProtoMessage()    {}
func (*ChaincodeIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_4991d8496920b696, []int{0}
}

func (m *ChaincodeIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeIdentifier.Unmarshal(m, b)
}
func (m *ChaincodeIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeIdentifier.Marshal(b, m, deterministic)
}
func (m *ChaincodeIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeIdentifier.Merge(m, src)
}
func (m *ChaincodeIdentifier) XXX_Size() int {
	return xxx_messageInfo_ChaincodeIdentifier.Size(m)
}
func (m *ChaincodeIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeIdentifier proto.InternalMessageInfo

func (m *ChaincodeIdentifier) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ChaincodeIdentifier) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// ChaincodeValidation instructs the peer how transactions for this chaincode should be
// validated.  The only validation mechanism which ships with fabric today is the standard
// 'vscc' validation mechanism.  This built in validation method utilizes an endorsement policy
// which checks that a sufficient number of signatures have been included.  The 'arguement'
// field encodes any parameters required by the validation implementation.
type ChaincodeValidation struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Argument             []byte   `protobuf:"bytes,2,opt,name=argument,proto3" json:"argument,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeValidation) Reset()         { *m = ChaincodeValidation{} }
func (m *ChaincodeValidation) String() string { return proto.CompactTextString(m) }
func (*ChaincodeValidation) ProtoMessage()    {}
func (*ChaincodeValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4991d8496920b696, []int{1}
}

func (m *ChaincodeValidation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeValidation.Unmarshal(m, b)
}
func (m *ChaincodeValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeValidation.Marshal(b, m, deterministic)
}
func (m *ChaincodeValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeValidation.Merge(m, src)
}
func (m *ChaincodeValidation) XXX_Size() int {
	return xxx_messageInfo_ChaincodeValidation.Size(m)
}
func (m *ChaincodeValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeValidation.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeValidation proto.InternalMessageInfo

func (m *ChaincodeValidation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeValidation) GetArgument() []byte {
	if m != nil {
		return m.Argument
	}
	return nil
}

// VSCCArgs is passed (marshaled) as a parameter to the VSCC imlementation via the
// argument field of the ChaincodeValidation message.
type VSCCArgs struct {
	EndorsementPolicyRef string   `protobuf:"bytes,1,opt,name=endorsement_policy_ref,json=endorsementPolicyRef,proto3" json:"endorsement_policy_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VSCCArgs) Reset()         { *m = VSCCArgs{} }
func (m *VSCCArgs) String() string { return proto.CompactTextString(m) }
func (*VSCCArgs) ProtoMessage()    {}
func (*VSCCArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_4991d8496920b696, []int{2}
}

func (m *VSCCArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VSCCArgs.Unmarshal(m, b)
}
func (m *VSCCArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VSCCArgs.Marshal(b, m, deterministic)
}
func (m *VSCCArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VSCCArgs.Merge(m, src)
}
func (m *VSCCArgs) XXX_Size() int {
	return xxx_messageInfo_VSCCArgs.Size(m)
}
func (m *VSCCArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_VSCCArgs.DiscardUnknown(m)
}

var xxx_messageInfo_VSCCArgs proto.InternalMessageInfo

func (m *VSCCArgs) GetEndorsementPolicyRef() string {
	if m != nil {
		return m.EndorsementPolicyRef
	}
	return ""
}

// ChaincodeEndorsement instructs the peer how transactions should be endorsed.  The only
// endorsement mechanism which ships with the fabric today is the standard 'escc' mechanism.
// This code simply simulates the proposal to generate a RW set, then signs the result
// using the peer's local signing identity.
type ChaincodeEndorsement struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeEndorsement) Reset()         { *m = ChaincodeEndorsement{} }
func (m *ChaincodeEndorsement) String() string { return proto.CompactTextString(m) }
func (*ChaincodeEndorsement) ProtoMessage()    {}
func (*ChaincodeEndorsement) Descriptor() ([]byte, []int) {
	return fileDescriptor_4991d8496920b696, []int{3}
}

func (m *ChaincodeEndorsement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeEndorsement.Unmarshal(m, b)
}
func (m *ChaincodeEndorsement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeEndorsement.Marshal(b, m, deterministic)
}
func (m *ChaincodeEndorsement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeEndorsement.Merge(m, src)
}
func (m *ChaincodeEndorsement) XXX_Size() int {
	return xxx_messageInfo_ChaincodeEndorsement.Size(m)
}
func (m *ChaincodeEndorsement) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeEndorsement.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeEndorsement proto.InternalMessageInfo

func (m *ChaincodeEndorsement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// ConfigTree encapsulates channel and resources configuration of a channel.
// Both configurations are represented as common.Config
type ConfigTree struct {
	ChannelConfig        *common.Config `protobuf:"bytes,1,opt,name=channel_config,json=channelConfig,proto3" json:"channel_config,omitempty"`
	ResourcesConfig      *common.Config `protobuf:"bytes,2,opt,name=resources_config,json=resourcesConfig,proto3" json:"resources_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConfigTree) Reset()         { *m = ConfigTree{} }
func (m *ConfigTree) String() string { return proto.CompactTextString(m) }
func (*ConfigTree) ProtoMessage()    {}
func (*ConfigTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_4991d8496920b696, []int{4}
}

func (m *ConfigTree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigTree.Unmarshal(m, b)
}
func (m *ConfigTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigTree.Marshal(b, m, deterministic)
}
func (m *ConfigTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigTree.Merge(m, src)
}
func (m *ConfigTree) XXX_Size() int {
	return xxx_messageInfo_ConfigTree.Size(m)
}
func (m *ConfigTree) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigTree.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigTree proto.InternalMessageInfo

func (m *ConfigTree) GetChannelConfig() *common.Config {
	if m != nil {
		return m.ChannelConfig
	}
	return nil
}

func (m *ConfigTree) GetResourcesConfig() *common.Config {
	if m != nil {
		return m.ResourcesConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeIdentifier)(nil), "protos.ChaincodeIdentifier")
	proto.RegisterType((*ChaincodeValidation)(nil), "protos.ChaincodeValidation")
	proto.RegisterType((*VSCCArgs)(nil), "protos.VSCCArgs")
	proto.RegisterType((*ChaincodeEndorsement)(nil), "protos.ChaincodeEndorsement")
	proto.RegisterType((*ConfigTree)(nil), "protos.ConfigTree")
}

func init() { proto.RegisterFile("peer/resources.proto", fileDescriptor_4991d8496920b696) }

var fileDescriptor_4991d8496920b696 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xd5, 0x0a, 0x41, 0x6b, 0x4a, 0x41, 0xa6, 0xa0, 0xaa, 0x53, 0x95, 0xa9, 0x42, 0x34,
	0x91, 0xf8, 0x33, 0xb0, 0x01, 0x51, 0x07, 0x36, 0x64, 0x50, 0x07, 0x96, 0xca, 0x75, 0x2e, 0x8e,
	0xa5, 0xc4, 0x8e, 0xce, 0x29, 0xa2, 0x0b, 0x9f, 0x1d, 0xc5, 0x6e, 0x43, 0x86, 0x4e, 0xbe, 0xf3,
	0xfb, 0xdd, 0xd3, 0xe9, 0x1d, 0x19, 0x95, 0x00, 0x18, 0x21, 0x58, 0xb3, 0x41, 0x01, 0x36, 0x2c,
	0xd1, 0x54, 0x86, 0x1e, 0xbb, 0xc7, 0x4e, 0xae, 0x84, 0x29, 0x0a, 0xa3, 0x23, 0x61, 0x74, 0xaa,
	0x64, 0xf5, 0xe3, 0xe5, 0x20, 0x26, 0x97, 0x71, 0xc6, 0x95, 0x16, 0x26, 0x81, 0xb7, 0x04, 0x74,
	0xa5, 0x52, 0x05, 0x48, 0x29, 0x39, 0xca, 0xb8, 0xcd, 0xc6, 0x9d, 0x69, 0x67, 0x36, 0x60, 0xae,
	0xa6, 0x63, 0x72, 0xf2, 0x0d, 0x68, 0x95, 0xd1, 0xe3, 0xee, 0xb4, 0x33, 0xeb, 0xb3, 0x7d, 0x1b,
	0x2c, 0x5a, 0x26, 0x4b, 0x9e, 0xab, 0x84, 0x57, 0xca, 0xe8, 0xda, 0x44, 0xf3, 0x02, 0x9c, 0x49,
	0x9f, 0xb9, 0x9a, 0x4e, 0x48, 0x8f, 0xa3, 0xdc, 0x14, 0xa0, 0x2b, 0xe7, 0x32, 0x60, 0x4d, 0x1f,
	0x3c, 0x93, 0xde, 0xf2, 0x23, 0x8e, 0x5f, 0x50, 0x5a, 0xfa, 0x40, 0xae, 0x41, 0x27, 0x06, 0x2d,
	0xd4, 0xd2, 0xaa, 0x34, 0xb9, 0x12, 0xdb, 0x15, 0x42, 0xba, 0x73, 0x1b, 0xb5, 0xd4, 0x77, 0x27,
	0x32, 0x48, 0x83, 0x1b, 0x32, 0x6a, 0x16, 0x59, 0xfc, 0x03, 0x87, 0x36, 0x09, 0x7e, 0x09, 0x89,
	0x5d, 0x16, 0x9f, 0x08, 0x40, 0x1f, 0xc9, 0x50, 0x64, 0x5c, 0x6b, 0xc8, 0x57, 0x3e, 0x21, 0xc7,
	0x9e, 0xde, 0x0d, 0x43, 0x9f, 0x5b, 0xe8, 0x59, 0x76, 0xb6, 0xa3, 0x7c, 0x4b, 0x9f, 0xc8, 0x45,
	0x13, 0xf8, 0x7e, 0xb0, 0x7b, 0x70, 0xf0, 0xbc, 0xe1, 0xfc, 0xc7, 0x2b, 0x23, 0x81, 0x41, 0x19,
	0x66, 0xdb, 0x12, 0x30, 0x87, 0x44, 0x02, 0x86, 0x29, 0x5f, 0xa3, 0x12, 0xfe, 0x32, 0x36, 0xac,
	0xcf, 0xf9, 0x75, 0x2b, 0x55, 0x95, 0x6d, 0xd6, 0xb5, 0x59, 0xd4, 0x42, 0x23, 0x8f, 0xce, 0x3d,
	0x3a, 0x97, 0x26, 0xaa, 0xe9, 0xb5, 0x3f, 0xf6, 0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae,
	0x32, 0x27, 0x0c, 0x0b, 0x02, 0x00, 0x00,
}
