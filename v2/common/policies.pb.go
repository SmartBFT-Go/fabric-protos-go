// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/policies.proto

package common

import (
	fmt "fmt"
	msp "github.com/SmartBFT-Go/fabric-protos-go/v2/msp"
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

type Policy_PolicyType int32

const (
	Policy_UNKNOWN          Policy_PolicyType = 0
	Policy_SIGNATURE        Policy_PolicyType = 1
	Policy_MSP              Policy_PolicyType = 2
	Policy_IMPLICIT_META    Policy_PolicyType = 3
	Policy_IMPLICIT_ORDERER Policy_PolicyType = 4
)

var Policy_PolicyType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SIGNATURE",
	2: "MSP",
	3: "IMPLICIT_META",
	4: "IMPLICIT_ORDERER",
}

var Policy_PolicyType_value = map[string]int32{
	"UNKNOWN":          0,
	"SIGNATURE":        1,
	"MSP":              2,
	"IMPLICIT_META":    3,
	"IMPLICIT_ORDERER": 4,
}

func (x Policy_PolicyType) String() string {
	return proto.EnumName(Policy_PolicyType_name, int32(x))
}

func (Policy_PolicyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{0, 0}
}

type ImplicitMetaPolicy_Rule int32

const (
	ImplicitMetaPolicy_ANY      ImplicitMetaPolicy_Rule = 0
	ImplicitMetaPolicy_ALL      ImplicitMetaPolicy_Rule = 1
	ImplicitMetaPolicy_MAJORITY ImplicitMetaPolicy_Rule = 2
)

var ImplicitMetaPolicy_Rule_name = map[int32]string{
	0: "ANY",
	1: "ALL",
	2: "MAJORITY",
}

var ImplicitMetaPolicy_Rule_value = map[string]int32{
	"ANY":      0,
	"ALL":      1,
	"MAJORITY": 2,
}

func (x ImplicitMetaPolicy_Rule) String() string {
	return proto.EnumName(ImplicitMetaPolicy_Rule_name, int32(x))
}

func (ImplicitMetaPolicy_Rule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{3, 0}
}

type ImplicitOrdererPolicy_Rule int32

const (
	ImplicitOrdererPolicy_SMARTBFT ImplicitOrdererPolicy_Rule = 0
)

var ImplicitOrdererPolicy_Rule_name = map[int32]string{
	0: "SMARTBFT",
}

var ImplicitOrdererPolicy_Rule_value = map[string]int32{
	"SMARTBFT": 0,
}

func (x ImplicitOrdererPolicy_Rule) String() string {
	return proto.EnumName(ImplicitOrdererPolicy_Rule_name, int32(x))
}

func (ImplicitOrdererPolicy_Rule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{4, 0}
}

// Policy expresses a policy which the orderer can evaluate, because there has been some desire expressed to support
// multiple policy engines, this is typed as a oneof for now
type Policy struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{0}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Policy) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// SignaturePolicyEnvelope wraps a SignaturePolicy and includes a version for future enhancements
type SignaturePolicyEnvelope struct {
	Version              int32               `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Rule                 *SignaturePolicy    `protobuf:"bytes,2,opt,name=rule,proto3" json:"rule,omitempty"`
	Identities           []*msp.MSPPrincipal `protobuf:"bytes,3,rep,name=identities,proto3" json:"identities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SignaturePolicyEnvelope) Reset()         { *m = SignaturePolicyEnvelope{} }
func (m *SignaturePolicyEnvelope) String() string { return proto.CompactTextString(m) }
func (*SignaturePolicyEnvelope) ProtoMessage()    {}
func (*SignaturePolicyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{1}
}

func (m *SignaturePolicyEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignaturePolicyEnvelope.Unmarshal(m, b)
}
func (m *SignaturePolicyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignaturePolicyEnvelope.Marshal(b, m, deterministic)
}
func (m *SignaturePolicyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignaturePolicyEnvelope.Merge(m, src)
}
func (m *SignaturePolicyEnvelope) XXX_Size() int {
	return xxx_messageInfo_SignaturePolicyEnvelope.Size(m)
}
func (m *SignaturePolicyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_SignaturePolicyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_SignaturePolicyEnvelope proto.InternalMessageInfo

func (m *SignaturePolicyEnvelope) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *SignaturePolicyEnvelope) GetRule() *SignaturePolicy {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *SignaturePolicyEnvelope) GetIdentities() []*msp.MSPPrincipal {
	if m != nil {
		return m.Identities
	}
	return nil
}

// SignaturePolicy is a recursive message structure which defines a featherweight DSL for describing
// policies which are more complicated than 'exactly this signature'.  The NOutOf operator is sufficent
// to express AND as well as OR, as well as of course N out of the following M policies
// SignedBy implies that the signature is from a valid certificate which is signed by the trusted
// authority specified in the bytes.  This will be the certificate itself for a self-signed certificate
// and will be the CA for more traditional certificates
type SignaturePolicy struct {
	// Types that are valid to be assigned to Type:
	//	*SignaturePolicy_SignedBy
	//	*SignaturePolicy_NOutOf_
	Type                 isSignaturePolicy_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SignaturePolicy) Reset()         { *m = SignaturePolicy{} }
func (m *SignaturePolicy) String() string { return proto.CompactTextString(m) }
func (*SignaturePolicy) ProtoMessage()    {}
func (*SignaturePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{2}
}

func (m *SignaturePolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignaturePolicy.Unmarshal(m, b)
}
func (m *SignaturePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignaturePolicy.Marshal(b, m, deterministic)
}
func (m *SignaturePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignaturePolicy.Merge(m, src)
}
func (m *SignaturePolicy) XXX_Size() int {
	return xxx_messageInfo_SignaturePolicy.Size(m)
}
func (m *SignaturePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_SignaturePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_SignaturePolicy proto.InternalMessageInfo

type isSignaturePolicy_Type interface {
	isSignaturePolicy_Type()
}

type SignaturePolicy_SignedBy struct {
	SignedBy int32 `protobuf:"varint,1,opt,name=signed_by,json=signedBy,proto3,oneof"`
}

type SignaturePolicy_NOutOf_ struct {
	NOutOf *SignaturePolicy_NOutOf `protobuf:"bytes,2,opt,name=n_out_of,json=nOutOf,proto3,oneof"`
}

func (*SignaturePolicy_SignedBy) isSignaturePolicy_Type() {}

func (*SignaturePolicy_NOutOf_) isSignaturePolicy_Type() {}

func (m *SignaturePolicy) GetType() isSignaturePolicy_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SignaturePolicy) GetSignedBy() int32 {
	if x, ok := m.GetType().(*SignaturePolicy_SignedBy); ok {
		return x.SignedBy
	}
	return 0
}

func (m *SignaturePolicy) GetNOutOf() *SignaturePolicy_NOutOf {
	if x, ok := m.GetType().(*SignaturePolicy_NOutOf_); ok {
		return x.NOutOf
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignaturePolicy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignaturePolicy_SignedBy)(nil),
		(*SignaturePolicy_NOutOf_)(nil),
	}
}

type SignaturePolicy_NOutOf struct {
	N                    int32              `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	Rules                []*SignaturePolicy `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SignaturePolicy_NOutOf) Reset()         { *m = SignaturePolicy_NOutOf{} }
func (m *SignaturePolicy_NOutOf) String() string { return proto.CompactTextString(m) }
func (*SignaturePolicy_NOutOf) ProtoMessage()    {}
func (*SignaturePolicy_NOutOf) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{2, 0}
}

func (m *SignaturePolicy_NOutOf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignaturePolicy_NOutOf.Unmarshal(m, b)
}
func (m *SignaturePolicy_NOutOf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignaturePolicy_NOutOf.Marshal(b, m, deterministic)
}
func (m *SignaturePolicy_NOutOf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignaturePolicy_NOutOf.Merge(m, src)
}
func (m *SignaturePolicy_NOutOf) XXX_Size() int {
	return xxx_messageInfo_SignaturePolicy_NOutOf.Size(m)
}
func (m *SignaturePolicy_NOutOf) XXX_DiscardUnknown() {
	xxx_messageInfo_SignaturePolicy_NOutOf.DiscardUnknown(m)
}

var xxx_messageInfo_SignaturePolicy_NOutOf proto.InternalMessageInfo

func (m *SignaturePolicy_NOutOf) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *SignaturePolicy_NOutOf) GetRules() []*SignaturePolicy {
	if m != nil {
		return m.Rules
	}
	return nil
}

// ImplicitMetaPolicy is a policy type which depends on the hierarchical nature of the configuration
// It is implicit because the rule is generate implicitly based on the number of sub policies
// It is meta because it depends only on the result of other policies
// When evaluated, this policy iterates over all immediate child sub-groups, retrieves the policy
// of name sub_policy, evaluates the collection and applies the rule.
// For example, with 4 sub-groups, and a policy name of "foo", ImplicitMetaPolicy retrieves
// each sub-group, retrieves policy "foo" for each subgroup, evaluates it, and, in the case of ANY
// 1 satisfied is sufficient, ALL would require 4 signatures, and MAJORITY would require 3 signatures.
type ImplicitMetaPolicy struct {
	SubPolicy            string                  `protobuf:"bytes,1,opt,name=sub_policy,json=subPolicy,proto3" json:"sub_policy,omitempty"`
	Rule                 ImplicitMetaPolicy_Rule `protobuf:"varint,2,opt,name=rule,proto3,enum=common.ImplicitMetaPolicy_Rule" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ImplicitMetaPolicy) Reset()         { *m = ImplicitMetaPolicy{} }
func (m *ImplicitMetaPolicy) String() string { return proto.CompactTextString(m) }
func (*ImplicitMetaPolicy) ProtoMessage()    {}
func (*ImplicitMetaPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{3}
}

func (m *ImplicitMetaPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImplicitMetaPolicy.Unmarshal(m, b)
}
func (m *ImplicitMetaPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImplicitMetaPolicy.Marshal(b, m, deterministic)
}
func (m *ImplicitMetaPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImplicitMetaPolicy.Merge(m, src)
}
func (m *ImplicitMetaPolicy) XXX_Size() int {
	return xxx_messageInfo_ImplicitMetaPolicy.Size(m)
}
func (m *ImplicitMetaPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_ImplicitMetaPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_ImplicitMetaPolicy proto.InternalMessageInfo

func (m *ImplicitMetaPolicy) GetSubPolicy() string {
	if m != nil {
		return m.SubPolicy
	}
	return ""
}

func (m *ImplicitMetaPolicy) GetRule() ImplicitMetaPolicy_Rule {
	if m != nil {
		return m.Rule
	}
	return ImplicitMetaPolicy_ANY
}

// ImplicitOrdererPolicy is a policy for block validation when the ordering service consists of a cluster of orderers
// running a distributed consensus algorithm.
//
// In the SmartBFT algorithm it is required that (at least) Q out of N orderers sign the block.
// (e.g. When N=3F+1, Q=2F+1). The policy checks the validity of at least Q signatures, and in addition, makes sure
// that each signer's certificate is in the last config block Consenters set.
//
// This makes sure that the signature comes from the actual node that participated in the consensus
// algorithm, not only that it has a valid role in the organization that issued the certificate.
type ImplicitOrdererPolicy struct {
	Rule                 ImplicitOrdererPolicy_Rule `protobuf:"varint,1,opt,name=rule,proto3,enum=common.ImplicitOrdererPolicy_Rule" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ImplicitOrdererPolicy) Reset()         { *m = ImplicitOrdererPolicy{} }
func (m *ImplicitOrdererPolicy) String() string { return proto.CompactTextString(m) }
func (*ImplicitOrdererPolicy) ProtoMessage()    {}
func (*ImplicitOrdererPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{4}
}

func (m *ImplicitOrdererPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImplicitOrdererPolicy.Unmarshal(m, b)
}
func (m *ImplicitOrdererPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImplicitOrdererPolicy.Marshal(b, m, deterministic)
}
func (m *ImplicitOrdererPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImplicitOrdererPolicy.Merge(m, src)
}
func (m *ImplicitOrdererPolicy) XXX_Size() int {
	return xxx_messageInfo_ImplicitOrdererPolicy.Size(m)
}
func (m *ImplicitOrdererPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_ImplicitOrdererPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_ImplicitOrdererPolicy proto.InternalMessageInfo

func (m *ImplicitOrdererPolicy) GetRule() ImplicitOrdererPolicy_Rule {
	if m != nil {
		return m.Rule
	}
	return ImplicitOrdererPolicy_SMARTBFT
}

// ApplicationPolicy captures the diffenrent policy types that
// are set and evaluted at the application level.
//
// Deprecated: Do not use.
type ApplicationPolicy struct {
	// Types that are valid to be assigned to Type:
	//	*ApplicationPolicy_SignaturePolicy
	//	*ApplicationPolicy_ChannelConfigPolicyReference
	Type                 isApplicationPolicy_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ApplicationPolicy) Reset()         { *m = ApplicationPolicy{} }
func (m *ApplicationPolicy) String() string { return proto.CompactTextString(m) }
func (*ApplicationPolicy) ProtoMessage()    {}
func (*ApplicationPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{5}
}

func (m *ApplicationPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationPolicy.Unmarshal(m, b)
}
func (m *ApplicationPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationPolicy.Marshal(b, m, deterministic)
}
func (m *ApplicationPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationPolicy.Merge(m, src)
}
func (m *ApplicationPolicy) XXX_Size() int {
	return xxx_messageInfo_ApplicationPolicy.Size(m)
}
func (m *ApplicationPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationPolicy proto.InternalMessageInfo

type isApplicationPolicy_Type interface {
	isApplicationPolicy_Type()
}

type ApplicationPolicy_SignaturePolicy struct {
	SignaturePolicy *SignaturePolicyEnvelope `protobuf:"bytes,1,opt,name=signature_policy,json=signaturePolicy,proto3,oneof"`
}

type ApplicationPolicy_ChannelConfigPolicyReference struct {
	ChannelConfigPolicyReference string `protobuf:"bytes,2,opt,name=channel_config_policy_reference,json=channelConfigPolicyReference,proto3,oneof"`
}

func (*ApplicationPolicy_SignaturePolicy) isApplicationPolicy_Type() {}

func (*ApplicationPolicy_ChannelConfigPolicyReference) isApplicationPolicy_Type() {}

func (m *ApplicationPolicy) GetType() isApplicationPolicy_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ApplicationPolicy) GetSignaturePolicy() *SignaturePolicyEnvelope {
	if x, ok := m.GetType().(*ApplicationPolicy_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}

func (m *ApplicationPolicy) GetChannelConfigPolicyReference() string {
	if x, ok := m.GetType().(*ApplicationPolicy_ChannelConfigPolicyReference); ok {
		return x.ChannelConfigPolicyReference
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ApplicationPolicy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ApplicationPolicy_SignaturePolicy)(nil),
		(*ApplicationPolicy_ChannelConfigPolicyReference)(nil),
	}
}

func init() {
	proto.RegisterEnum("common.Policy_PolicyType", Policy_PolicyType_name, Policy_PolicyType_value)
	proto.RegisterEnum("common.ImplicitMetaPolicy_Rule", ImplicitMetaPolicy_Rule_name, ImplicitMetaPolicy_Rule_value)
	proto.RegisterEnum("common.ImplicitOrdererPolicy_Rule", ImplicitOrdererPolicy_Rule_name, ImplicitOrdererPolicy_Rule_value)
	proto.RegisterType((*Policy)(nil), "common.Policy")
	proto.RegisterType((*SignaturePolicyEnvelope)(nil), "common.SignaturePolicyEnvelope")
	proto.RegisterType((*SignaturePolicy)(nil), "common.SignaturePolicy")
	proto.RegisterType((*SignaturePolicy_NOutOf)(nil), "common.SignaturePolicy.NOutOf")
	proto.RegisterType((*ImplicitMetaPolicy)(nil), "common.ImplicitMetaPolicy")
	proto.RegisterType((*ImplicitOrdererPolicy)(nil), "common.ImplicitOrdererPolicy")
	proto.RegisterType((*ApplicationPolicy)(nil), "common.ApplicationPolicy")
}

func init() { proto.RegisterFile("common/policies.proto", fileDescriptor_0d02cf0d453425a3) }

var fileDescriptor_0d02cf0d453425a3 = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0xf5, 0x02, 0x21, 0x30, 0x90, 0xc6, 0x59, 0x11, 0x05, 0x45, 0x6d, 0x13, 0x59, 0x55, 0x15,
	0xa9, 0x8a, 0x91, 0x92, 0xaa, 0x87, 0xdc, 0x20, 0xa5, 0x81, 0x16, 0x03, 0x5a, 0x9c, 0x56, 0xc9,
	0xc5, 0x32, 0x66, 0x71, 0x56, 0x32, 0xbb, 0xd6, 0xda, 0x8e, 0xc4, 0x4f, 0x54, 0x3d, 0xf5, 0x53,
	0x7a, 0xe9, 0xcf, 0x55, 0xf6, 0xda, 0x29, 0xa5, 0xca, 0x6d, 0x67, 0xfc, 0x66, 0xde, 0x7b, 0x33,
	0x63, 0x38, 0xf4, 0xc4, 0x6a, 0x25, 0x78, 0x27, 0x14, 0x01, 0xf3, 0x18, 0x8d, 0xcc, 0x50, 0x8a,
	0x58, 0xe0, 0xaa, 0x4a, 0x1f, 0x1f, 0xad, 0xa2, 0xb0, 0xb3, 0x8a, 0x42, 0x27, 0x94, 0x8c, 0x7b,
	0x2c, 0x74, 0x03, 0x05, 0x30, 0xbe, 0x23, 0xa8, 0x4e, 0xd3, 0x9a, 0x35, 0xc6, 0x50, 0x89, 0xd7,
	0x21, 0x6d, 0xa3, 0x53, 0x74, 0xb6, 0x43, 0xb2, 0x37, 0x6e, 0xc1, 0xce, 0xa3, 0x1b, 0x24, 0xb4,
	0x5d, 0x3a, 0x45, 0x67, 0x4d, 0xa2, 0x02, 0xe3, 0x1e, 0x40, 0xd5, 0xd8, 0x29, 0xa6, 0x01, 0xbb,
	0xb7, 0xe3, 0x2f, 0xe3, 0xc9, 0xb7, 0xb1, 0xae, 0xe1, 0x3d, 0xa8, 0xcf, 0x86, 0x37, 0xe3, 0xae,
	0x7d, 0x4b, 0xfa, 0x3a, 0xc2, 0xbb, 0x50, 0xb6, 0x66, 0x53, 0xbd, 0x84, 0x0f, 0x60, 0x6f, 0x68,
	0x4d, 0x47, 0xc3, 0xeb, 0xa1, 0xed, 0x58, 0x7d, 0xbb, 0xab, 0x97, 0x71, 0x0b, 0xf4, 0xa7, 0xd4,
	0x84, 0x7c, 0xec, 0x93, 0x3e, 0xd1, 0x2b, 0xc6, 0x4f, 0x04, 0x47, 0x33, 0xe6, 0x73, 0x37, 0x4e,
	0x24, 0x55, 0x2c, 0x7d, 0xfe, 0x48, 0x03, 0x11, 0x52, 0xdc, 0x86, 0xdd, 0x47, 0x2a, 0x23, 0x26,
	0x78, 0x2e, 0xb2, 0x08, 0xf1, 0x3b, 0xa8, 0xc8, 0x24, 0x50, 0x32, 0x1b, 0x17, 0x47, 0xa6, 0xb2,
	0x6d, 0x6e, 0x35, 0x22, 0x19, 0x08, 0xbf, 0x07, 0x60, 0x0b, 0xca, 0x63, 0x16, 0x33, 0x1a, 0xb5,
	0xcb, 0xa7, 0xe5, 0xb3, 0xc6, 0x45, 0xab, 0x28, 0xb1, 0x66, 0xd3, 0x69, 0x31, 0x23, 0xb2, 0x81,
	0x33, 0x7e, 0x23, 0xd8, 0xdf, 0xea, 0x87, 0x5f, 0x41, 0x3d, 0x62, 0x3e, 0xa7, 0x0b, 0x67, 0xbe,
	0x56, 0x92, 0x06, 0x1a, 0xa9, 0xa9, 0x54, 0x6f, 0x8d, 0xaf, 0xa0, 0xc6, 0x1d, 0x91, 0xc4, 0x8e,
	0x58, 0xe6, 0xca, 0x5e, 0x3f, 0xa3, 0xcc, 0x1c, 0x4f, 0x92, 0x78, 0xb2, 0x1c, 0x68, 0xa4, 0xca,
	0xb3, 0xd7, 0x71, 0x1f, 0xaa, 0x2a, 0x87, 0x9b, 0x80, 0x0a, 0xbf, 0x88, 0xe3, 0x73, 0xd8, 0x49,
	0x4d, 0x44, 0xed, 0x52, 0xa6, 0xfb, 0x59, 0xab, 0x0a, 0xd5, 0xab, 0x42, 0x25, 0x5d, 0x92, 0xf1,
	0x03, 0x01, 0x1e, 0xae, 0xc2, 0xf4, 0x38, 0x62, 0x8b, 0xc6, 0xee, 0x93, 0x01, 0x88, 0x92, 0xb9,
	0x93, 0x5d, 0x8d, 0x72, 0x50, 0x27, 0xf5, 0x28, 0x99, 0xe7, 0x9f, 0x2f, 0x37, 0xc6, 0xfa, 0xe2,
	0xe2, 0xa4, 0xe0, 0xfa, 0xbf, 0x91, 0x49, 0x92, 0x80, 0xaa, 0xf1, 0x1a, 0x6f, 0xa1, 0x92, 0x46,
	0xe9, 0xee, 0xbb, 0xe3, 0x3b, 0x5d, 0xcb, 0x1e, 0xa3, 0x91, 0x8e, 0x70, 0x13, 0x6a, 0x56, 0xf7,
	0xf3, 0x84, 0x0c, 0xed, 0x3b, 0xbd, 0x64, 0x50, 0x38, 0x2c, 0x1a, 0x4d, 0xe4, 0x82, 0x4a, 0x2a,
	0x73, 0xd6, 0x0f, 0x39, 0x2b, 0xca, 0x58, 0x8d, 0x6d, 0xd6, 0x7f, 0xc0, 0x9b, 0xc4, 0xad, 0x9c,
	0xb8, 0x09, 0xb5, 0x99, 0xd5, 0x25, 0x76, 0xef, 0x93, 0xad, 0x6b, 0xc6, 0x2f, 0x04, 0x07, 0xdd,
	0x30, 0x2d, 0x75, 0x63, 0x26, 0x78, 0xce, 0x31, 0x02, 0x3d, 0x2a, 0x26, 0xb6, 0x69, 0xbf, 0xf1,
	0xd7, 0xe5, 0x33, 0x57, 0x38, 0xd0, 0xc8, 0x7e, 0xb4, 0x75, 0x07, 0x37, 0x70, 0xe2, 0x3d, 0xb8,
	0x9c, 0xd3, 0xc0, 0xf1, 0x04, 0x5f, 0x32, 0x3f, 0x6f, 0xe9, 0x48, 0xba, 0xa4, 0x92, 0x72, 0x4f,
	0x8d, 0xb0, 0x3e, 0xd0, 0xc8, 0xcb, 0x1c, 0x78, 0x9d, 0xe1, 0xf2, 0x65, 0x15, 0xa8, 0xab, 0x52,
	0x1b, 0x15, 0x2b, 0xeb, 0x7d, 0x85, 0x37, 0x42, 0xfa, 0xe6, 0xc3, 0x3a, 0xa4, 0x32, 0xa0, 0x0b,
	0x9f, 0x4a, 0x73, 0xe9, 0xce, 0x25, 0xf3, 0xd4, 0xaf, 0x1b, 0xe5, 0x3a, 0xef, 0x4d, 0x9f, 0xc5,
	0x0f, 0xc9, 0x3c, 0x0d, 0x3b, 0x1b, 0xe0, 0x8e, 0x02, 0x9f, 0x2b, 0xf0, 0xb9, 0x2f, 0x3a, 0x0a,
	0x3f, 0xaf, 0x66, 0x99, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x76, 0x42, 0x49, 0xa8, 0x33,
	0x04, 0x00, 0x00,
}
