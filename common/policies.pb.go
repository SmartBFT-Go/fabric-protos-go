// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/policies.proto

package common

import (
	fmt "fmt"
	msp "github.com/SmartBFT-Go/fabric-protos-go/msp"
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
	Policy_UNKNOWN       Policy_PolicyType = 0
	Policy_SIGNATURE     Policy_PolicyType = 1
	Policy_MSP           Policy_PolicyType = 2
	Policy_IMPLICIT_META Policy_PolicyType = 3
)

var Policy_PolicyType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SIGNATURE",
	2: "MSP",
	3: "IMPLICIT_META",
}

var Policy_PolicyType_value = map[string]int32{
	"UNKNOWN":       0,
	"SIGNATURE":     1,
	"MSP":           2,
	"IMPLICIT_META": 3,
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
	return fileDescriptor_0d02cf0d453425a3, []int{4}
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
	proto.RegisterType((*Policy)(nil), "common.Policy")
	proto.RegisterType((*SignaturePolicyEnvelope)(nil), "common.SignaturePolicyEnvelope")
	proto.RegisterType((*SignaturePolicy)(nil), "common.SignaturePolicy")
	proto.RegisterType((*SignaturePolicy_NOutOf)(nil), "common.SignaturePolicy.NOutOf")
	proto.RegisterType((*ImplicitMetaPolicy)(nil), "common.ImplicitMetaPolicy")
	proto.RegisterType((*ApplicationPolicy)(nil), "common.ApplicationPolicy")
}

func init() { proto.RegisterFile("common/policies.proto", fileDescriptor_0d02cf0d453425a3) }

var fileDescriptor_0d02cf0d453425a3 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xd1, 0x4e, 0xdb, 0x30,
	0x14, 0x86, 0xe3, 0xb6, 0x04, 0x7a, 0x0a, 0x23, 0x58, 0x4c, 0x44, 0x68, 0x1b, 0x28, 0x9a, 0x26,
	0xa4, 0x89, 0x54, 0x82, 0x5d, 0x71, 0x57, 0x58, 0x45, 0xb3, 0x35, 0x69, 0xe5, 0x96, 0x4d, 0xec,
	0x26, 0x4a, 0x82, 0x1b, 0x2c, 0xa5, 0x76, 0x14, 0x27, 0x68, 0x7d, 0x8b, 0x5d, 0xed, 0x51, 0x76,
	0xb3, 0x97, 0x9b, 0x12, 0x27, 0x5b, 0xc5, 0xd4, 0x3b, 0x9f, 0x93, 0xcf, 0xbf, 0xff, 0xff, 0xd8,
	0x81, 0x97, 0x91, 0x58, 0x2e, 0x05, 0xef, 0xa7, 0x22, 0x61, 0x11, 0xa3, 0xd2, 0x4e, 0x33, 0x91,
	0x0b, 0xac, 0xab, 0xf6, 0xf1, 0xd1, 0x52, 0xa6, 0xfd, 0xa5, 0x4c, 0xfd, 0x34, 0x63, 0x3c, 0x62,
	0x69, 0x90, 0x28, 0xc0, 0xfa, 0x0e, 0xfa, 0xb4, 0xdc, 0xb2, 0xc2, 0x18, 0x3a, 0xf9, 0x2a, 0xa5,
	0x26, 0x3a, 0x45, 0x67, 0x5b, 0xa4, 0x5a, 0xe3, 0x43, 0xd8, 0x7a, 0x0a, 0x92, 0x82, 0x9a, 0xad,
	0x53, 0x74, 0xb6, 0x4b, 0x54, 0x61, 0x7d, 0x04, 0x50, 0x7b, 0xe6, 0x25, 0xd3, 0x83, 0xed, 0x3b,
	0xef, 0xb3, 0x37, 0xf9, 0xea, 0x19, 0x1a, 0xde, 0x83, 0xee, 0xcc, 0xb9, 0xf5, 0x06, 0xf3, 0x3b,
	0x32, 0x34, 0x10, 0xde, 0x86, 0xb6, 0x3b, 0x9b, 0x1a, 0x2d, 0x7c, 0x00, 0x7b, 0x8e, 0x3b, 0x1d,
	0x3b, 0x37, 0xce, 0xdc, 0x77, 0x87, 0xf3, 0x81, 0xd1, 0xb6, 0x7e, 0x22, 0x38, 0x9a, 0xb1, 0x98,
	0x07, 0x79, 0x91, 0x51, 0xa5, 0x37, 0xe4, 0x4f, 0x34, 0x11, 0x29, 0xc5, 0x26, 0x6c, 0x3f, 0xd1,
	0x4c, 0x32, 0xc1, 0x6b, 0x3b, 0x4d, 0x89, 0xdf, 0x43, 0x27, 0x2b, 0x12, 0x65, 0xa8, 0x77, 0x71,
	0x64, 0xab, 0x7c, 0xf6, 0x33, 0x21, 0x52, 0x41, 0xf8, 0x03, 0x00, 0x7b, 0xa0, 0x3c, 0x67, 0x39,
	0xa3, 0xd2, 0x6c, 0x9f, 0xb6, 0xcf, 0x7a, 0x17, 0x87, 0xcd, 0x16, 0x77, 0x36, 0x9d, 0x36, 0xc3,
	0x20, 0x6b, 0x9c, 0xf5, 0x1b, 0xc1, 0xfe, 0x33, 0x3d, 0xfc, 0x1a, 0xba, 0x92, 0xc5, 0x9c, 0x3e,
	0xf8, 0xe1, 0x4a, 0x59, 0x1a, 0x69, 0x64, 0x47, 0xb5, 0xae, 0x57, 0xf8, 0x0a, 0x76, 0xb8, 0x2f,
	0x8a, 0xdc, 0x17, 0x8b, 0xda, 0xd9, 0x9b, 0x0d, 0xce, 0x6c, 0x6f, 0x52, 0xe4, 0x93, 0xc5, 0x48,
	0x23, 0x3a, 0xaf, 0x56, 0xc7, 0x43, 0xd0, 0x55, 0x0f, 0xef, 0x02, 0x6a, 0xf2, 0x22, 0x8e, 0xcf,
	0x61, 0xab, 0x0c, 0x21, 0xcd, 0x56, 0xe5, 0x7b, 0x63, 0x54, 0x45, 0x5d, 0xeb, 0xd0, 0x29, 0xaf,
	0xc3, 0xfa, 0x81, 0x00, 0x3b, 0xcb, 0xb4, 0x7c, 0x05, 0xb9, 0x4b, 0xf3, 0xe0, 0x6f, 0x00, 0x90,
	0x45, 0xe8, 0x57, 0xcf, 0x43, 0x25, 0xe8, 0x92, 0xae, 0x2c, 0xc2, 0xfa, 0xf3, 0xe5, 0xda, 0x58,
	0x5f, 0x5c, 0x9c, 0x34, 0x67, 0xfd, 0x2f, 0x64, 0x93, 0x22, 0xa1, 0x6a, 0xbc, 0xd6, 0x3b, 0xe8,
	0x94, 0x55, 0x79, 0xcb, 0x03, 0xef, 0xde, 0xd0, 0xaa, 0xc5, 0x78, 0x6c, 0x20, 0xbc, 0x0b, 0x3b,
	0xee, 0xe0, 0xd3, 0x84, 0x38, 0xf3, 0x7b, 0xa3, 0x65, 0xfd, 0x42, 0x70, 0x30, 0x48, 0x4b, 0xa5,
	0x20, 0x67, 0x82, 0xd7, 0x47, 0x8e, 0xc1, 0x90, 0x4d, 0x94, 0x75, 0x5f, 0xbd, 0x7f, 0xc7, 0x6f,
	0x78, 0x1e, 0x23, 0x8d, 0xec, 0xcb, 0x67, 0x17, 0x74, 0x0b, 0x27, 0xd1, 0x63, 0xc0, 0x39, 0x4d,
	0xfc, 0x48, 0xf0, 0x05, 0x8b, 0x6b, 0x49, 0x3f, 0xa3, 0x0b, 0x9a, 0x51, 0x1e, 0xa9, 0x6c, 0xdd,
	0x91, 0x46, 0x5e, 0xd5, 0xe0, 0x4d, 0xc5, 0xd5, 0x53, 0x6c, 0xa8, 0xab, 0x96, 0x89, 0x9a, 0x59,
	0x5e, 0x7f, 0x81, 0xb7, 0x22, 0x8b, 0xed, 0xc7, 0x55, 0x4a, 0xb3, 0x84, 0x3e, 0xc4, 0x34, 0xb3,
	0x17, 0x41, 0x98, 0xb1, 0x48, 0xfd, 0x3c, 0xb2, 0xf6, 0xf9, 0xcd, 0x8e, 0x59, 0xfe, 0x58, 0x84,
	0x65, 0xd9, 0x5f, 0x83, 0xfb, 0x0a, 0x3e, 0x57, 0xf0, 0x79, 0x2c, 0xfa, 0x8a, 0x0f, 0xf5, 0xaa,
	0x73, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x77, 0x84, 0xc8, 0xc7, 0xb5, 0x03, 0x00, 0x00,
}
