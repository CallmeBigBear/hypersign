// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hypersign/ssi/v1/credential_schema.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CredentialSchemaDocument struct {
	Context      []string                  `protobuf:"bytes,1,rep,name=context,json=@context,proto3" json:"@context"`
	Type         string                    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ModelVersion string                    `protobuf:"bytes,3,opt,name=modelVersion,proto3" json:"modelVersion,omitempty"`
	Id           string                    `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                    `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Author       string                    `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
	Authored     string                    `protobuf:"bytes,7,opt,name=authored,proto3" json:"authored,omitempty"`
	Schema       *CredentialSchemaProperty `protobuf:"bytes,8,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (m *CredentialSchemaDocument) Reset()         { *m = CredentialSchemaDocument{} }
func (m *CredentialSchemaDocument) String() string { return proto.CompactTextString(m) }
func (*CredentialSchemaDocument) ProtoMessage()    {}
func (*CredentialSchemaDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0546f877a0a994, []int{0}
}
func (m *CredentialSchemaDocument) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialSchemaDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialSchemaDocument.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialSchemaDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialSchemaDocument.Merge(m, src)
}
func (m *CredentialSchemaDocument) XXX_Size() int {
	return m.Size()
}
func (m *CredentialSchemaDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialSchemaDocument.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialSchemaDocument proto.InternalMessageInfo

func (m *CredentialSchemaDocument) GetContext() []string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *CredentialSchemaDocument) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CredentialSchemaDocument) GetModelVersion() string {
	if m != nil {
		return m.ModelVersion
	}
	return ""
}

func (m *CredentialSchemaDocument) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CredentialSchemaDocument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CredentialSchemaDocument) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *CredentialSchemaDocument) GetAuthored() string {
	if m != nil {
		return m.Authored
	}
	return ""
}

func (m *CredentialSchemaDocument) GetSchema() *CredentialSchemaProperty {
	if m != nil {
		return m.Schema
	}
	return nil
}

type CredentialSchemaProperty struct {
	Schema               string   `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Properties           string   `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
	Required             []string `protobuf:"bytes,5,rep,name=required,proto3" json:"required,omitempty"`
	AdditionalProperties bool     `protobuf:"varint,6,opt,name=additionalProperties,proto3" json:"additionalProperties,omitempty"`
}

func (m *CredentialSchemaProperty) Reset()         { *m = CredentialSchemaProperty{} }
func (m *CredentialSchemaProperty) String() string { return proto.CompactTextString(m) }
func (*CredentialSchemaProperty) ProtoMessage()    {}
func (*CredentialSchemaProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0546f877a0a994, []int{1}
}
func (m *CredentialSchemaProperty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialSchemaProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialSchemaProperty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialSchemaProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialSchemaProperty.Merge(m, src)
}
func (m *CredentialSchemaProperty) XXX_Size() int {
	return m.Size()
}
func (m *CredentialSchemaProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialSchemaProperty.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialSchemaProperty proto.InternalMessageInfo

func (m *CredentialSchemaProperty) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *CredentialSchemaProperty) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CredentialSchemaProperty) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CredentialSchemaProperty) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

func (m *CredentialSchemaProperty) GetRequired() []string {
	if m != nil {
		return m.Required
	}
	return nil
}

func (m *CredentialSchemaProperty) GetAdditionalProperties() bool {
	if m != nil {
		return m.AdditionalProperties
	}
	return false
}

type CredentialSchemaState struct {
	CredentialSchemaDocument *CredentialSchemaDocument `protobuf:"bytes,1,opt,name=credentialSchemaDocument,proto3" json:"credentialSchemaDocument,omitempty"`
	CredentialSchemaProof    *DocumentProof            `protobuf:"bytes,2,opt,name=credentialSchemaProof,proto3" json:"credentialSchemaProof,omitempty"`
}

func (m *CredentialSchemaState) Reset()         { *m = CredentialSchemaState{} }
func (m *CredentialSchemaState) String() string { return proto.CompactTextString(m) }
func (*CredentialSchemaState) ProtoMessage()    {}
func (*CredentialSchemaState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0546f877a0a994, []int{2}
}
func (m *CredentialSchemaState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialSchemaState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialSchemaState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialSchemaState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialSchemaState.Merge(m, src)
}
func (m *CredentialSchemaState) XXX_Size() int {
	return m.Size()
}
func (m *CredentialSchemaState) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialSchemaState.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialSchemaState proto.InternalMessageInfo

func (m *CredentialSchemaState) GetCredentialSchemaDocument() *CredentialSchemaDocument {
	if m != nil {
		return m.CredentialSchemaDocument
	}
	return nil
}

func (m *CredentialSchemaState) GetCredentialSchemaProof() *DocumentProof {
	if m != nil {
		return m.CredentialSchemaProof
	}
	return nil
}

func init() {
	proto.RegisterType((*CredentialSchemaDocument)(nil), "hypersign.ssi.v1.CredentialSchemaDocument")
	proto.RegisterType((*CredentialSchemaProperty)(nil), "hypersign.ssi.v1.CredentialSchemaProperty")
	proto.RegisterType((*CredentialSchemaState)(nil), "hypersign.ssi.v1.CredentialSchemaState")
}

func init() {
	proto.RegisterFile("hypersign/ssi/v1/credential_schema.proto", fileDescriptor_5c0546f877a0a994)
}

var fileDescriptor_5c0546f877a0a994 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0x26, 0x6d, 0x9a, 0x6e, 0x2a, 0x84, 0x56, 0x2d, 0x5a, 0x45, 0xc8, 0x8d, 0x72, 0x21,
	0x42, 0xaa, 0xad, 0x86, 0x1f, 0x40, 0x81, 0x23, 0x87, 0xca, 0x15, 0x1c, 0xb8, 0x20, 0x77, 0x77,
	0x13, 0xaf, 0x14, 0x7b, 0xcc, 0xee, 0xba, 0x6a, 0xfe, 0x82, 0x1b, 0xbf, 0xc4, 0xb1, 0x27, 0xd4,
	0x13, 0x42, 0xc9, 0x8d, 0xaf, 0x40, 0xbb, 0xeb, 0x98, 0x10, 0x37, 0x52, 0x6f, 0x33, 0x6f, 0x66,
	0xde, 0xcc, 0x7b, 0xf6, 0xe2, 0x71, 0xba, 0x2c, 0x84, 0xd2, 0x72, 0x9e, 0x47, 0x5a, 0xcb, 0xe8,
	0xf6, 0x32, 0x62, 0x4a, 0x70, 0x91, 0x1b, 0x99, 0x2c, 0xbe, 0x68, 0x96, 0x8a, 0x2c, 0x09, 0x0b,
	0x05, 0x06, 0xc8, 0xf3, 0xba, 0x33, 0xd4, 0x5a, 0x86, 0xb7, 0x97, 0x83, 0x97, 0x8d, 0xd9, 0x42,
	0x01, 0xcc, 0x7c, 0xff, 0xe0, 0x74, 0x0e, 0x73, 0x70, 0x61, 0x64, 0x23, 0x8f, 0x8e, 0xbe, 0xb7,
	0x31, 0x7d, 0x57, 0x6f, 0xb8, 0x76, 0x0b, 0xde, 0x03, 0x2b, 0x33, 0x91, 0x1b, 0xf2, 0x0a, 0x1f,
	0x31, 0xc8, 0x8d, 0xb8, 0x33, 0x14, 0x0d, 0x3b, 0xe3, 0xe3, 0xe9, 0xc9, 0x9f, 0x5f, 0xe7, 0xbd,
	0xb7, 0x15, 0x16, 0xd7, 0x11, 0x21, 0xf8, 0xc0, 0x2c, 0x0b, 0x41, 0xdb, 0x43, 0x34, 0x3e, 0x8e,
	0x5d, 0x4c, 0x46, 0xf8, 0x24, 0x03, 0x2e, 0x16, 0x9f, 0xec, 0x49, 0x90, 0xd3, 0x8e, 0xab, 0xfd,
	0x87, 0x91, 0x67, 0xb8, 0x2d, 0x39, 0x3d, 0x70, 0x95, 0xb6, 0xe4, 0x96, 0x27, 0x4f, 0x32, 0x41,
	0x0f, 0x3d, 0x8f, 0x8d, 0xc9, 0x0b, 0xdc, 0x4d, 0x4a, 0x93, 0x82, 0xa2, 0x5d, 0x87, 0x56, 0x19,
	0x19, 0xe0, 0x9e, 0x8f, 0x04, 0xa7, 0x47, 0xae, 0x52, 0xe7, 0x64, 0x8a, 0xbb, 0xde, 0x2b, 0xda,
	0x1b, 0xa2, 0x71, 0x7f, 0xf2, 0x3a, 0xdc, 0x35, 0x2b, 0xdc, 0x15, 0x7d, 0xa5, 0xa0, 0x10, 0xca,
	0x2c, 0xe3, 0x6a, 0x72, 0xf4, 0x80, 0x9a, 0xce, 0x6c, 0x9a, 0xec, 0x51, 0xd5, 0x02, 0xe4, 0x8f,
	0xf2, 0x19, 0x19, 0xe2, 0x3e, 0x17, 0x9a, 0x29, 0x59, 0x18, 0xab, 0xd9, 0xfb, 0xb1, 0x0d, 0xd5,
	0x56, 0x75, 0xb6, 0xac, 0x0a, 0x30, 0x2e, 0x3c, 0xb3, 0x14, 0xba, 0xb2, 0x63, 0x0b, 0xb1, 0x52,
	0x95, 0xf8, 0x5a, 0x4a, 0x2b, 0xf5, 0xd0, 0x7e, 0x88, 0xb8, 0xce, 0xc9, 0x04, 0x9f, 0x26, 0x9c,
	0x4b, 0xcb, 0x9d, 0x2c, 0xae, 0xfe, 0xb1, 0x58, 0xb3, 0x7a, 0xf1, 0xa3, 0xb5, 0xd1, 0x4f, 0x84,
	0xcf, 0x76, 0xa5, 0x5d, 0x9b, 0xc4, 0x08, 0x32, 0xc3, 0x94, 0xed, 0xf9, 0x1b, 0x9c, 0xd2, 0x27,
	0x59, 0xb9, 0x99, 0x88, 0xf7, 0x72, 0x91, 0x8f, 0xf8, 0x8c, 0x35, 0xbd, 0x85, 0x99, 0x73, 0xac,
	0x3f, 0x39, 0x6f, 0x2e, 0xd9, 0x8c, 0xba, 0xb6, 0xf8, 0xf1, 0xe9, 0xe9, 0x87, 0x1f, 0xab, 0x00,
	0xdd, 0xaf, 0x02, 0xf4, 0x7b, 0x15, 0xa0, 0x6f, 0xeb, 0xa0, 0x75, 0xbf, 0x0e, 0x5a, 0x0f, 0xeb,
	0xa0, 0xf5, 0x79, 0x32, 0x97, 0x26, 0x2d, 0x6f, 0x42, 0x06, 0x59, 0x54, 0x73, 0x5f, 0xb8, 0x37,
	0xc0, 0x60, 0x11, 0xa5, 0x92, 0x5f, 0xe4, 0xc0, 0x45, 0x74, 0xe7, 0x9e, 0x8e, 0xfd, 0x2a, 0xfa,
	0xa6, 0xeb, 0xca, 0x6f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xc3, 0xde, 0xb1, 0x94, 0x03,
	0x00, 0x00,
}

func (m *CredentialSchemaDocument) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialSchemaDocument) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialSchemaDocument) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Schema != nil {
		{
			size, err := m.Schema.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCredentialSchema(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Authored) > 0 {
		i -= len(m.Authored)
		copy(dAtA[i:], m.Authored)
		i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Authored)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Author) > 0 {
		i -= len(m.Author)
		copy(dAtA[i:], m.Author)
		i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Author)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ModelVersion) > 0 {
		i -= len(m.ModelVersion)
		copy(dAtA[i:], m.ModelVersion)
		i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.ModelVersion)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Context) > 0 {
		for iNdEx := len(m.Context) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Context[iNdEx])
			copy(dAtA[i:], m.Context[iNdEx])
			i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Context[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CredentialSchemaProperty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialSchemaProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialSchemaProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AdditionalProperties {
		i--
		if m.AdditionalProperties {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Required) > 0 {
		for iNdEx := len(m.Required) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Required[iNdEx])
			copy(dAtA[i:], m.Required[iNdEx])
			i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Required[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Properties) > 0 {
		i -= len(m.Properties)
		copy(dAtA[i:], m.Properties)
		i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Properties)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Schema) > 0 {
		i -= len(m.Schema)
		copy(dAtA[i:], m.Schema)
		i = encodeVarintCredentialSchema(dAtA, i, uint64(len(m.Schema)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CredentialSchemaState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialSchemaState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialSchemaState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CredentialSchemaProof != nil {
		{
			size, err := m.CredentialSchemaProof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCredentialSchema(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CredentialSchemaDocument != nil {
		{
			size, err := m.CredentialSchemaDocument.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCredentialSchema(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCredentialSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovCredentialSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CredentialSchemaDocument) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Context) > 0 {
		for _, s := range m.Context {
			l = len(s)
			n += 1 + l + sovCredentialSchema(uint64(l))
		}
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	l = len(m.ModelVersion)
	if l > 0 {
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	l = len(m.Author)
	if l > 0 {
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	l = len(m.Authored)
	if l > 0 {
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	if m.Schema != nil {
		l = m.Schema.Size()
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	return n
}

func (m *CredentialSchemaProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Schema)
	if l > 0 {
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	l = len(m.Properties)
	if l > 0 {
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	if len(m.Required) > 0 {
		for _, s := range m.Required {
			l = len(s)
			n += 1 + l + sovCredentialSchema(uint64(l))
		}
	}
	if m.AdditionalProperties {
		n += 2
	}
	return n
}

func (m *CredentialSchemaState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CredentialSchemaDocument != nil {
		l = m.CredentialSchemaDocument.Size()
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	if m.CredentialSchemaProof != nil {
		l = m.CredentialSchemaProof.Size()
		n += 1 + l + sovCredentialSchema(uint64(l))
	}
	return n
}

func sovCredentialSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCredentialSchema(x uint64) (n int) {
	return sovCredentialSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CredentialSchemaDocument) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredentialSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CredentialSchemaDocument: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialSchemaDocument: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = append(m.Context, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModelVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Author = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authored", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authored = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Schema == nil {
				m.Schema = &CredentialSchemaProperty{}
			}
			if err := m.Schema.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCredentialSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CredentialSchemaProperty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredentialSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CredentialSchemaProperty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialSchemaProperty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Properties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Required", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Required = append(m.Required, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalProperties", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AdditionalProperties = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCredentialSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CredentialSchemaState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredentialSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CredentialSchemaState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialSchemaState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialSchemaDocument", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CredentialSchemaDocument == nil {
				m.CredentialSchemaDocument = &CredentialSchemaDocument{}
			}
			if err := m.CredentialSchemaDocument.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialSchemaProof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CredentialSchemaProof == nil {
				m.CredentialSchemaProof = &DocumentProof{}
			}
			if err := m.CredentialSchemaProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCredentialSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredentialSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCredentialSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCredentialSchema
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCredentialSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCredentialSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCredentialSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCredentialSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCredentialSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCredentialSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCredentialSchema = fmt.Errorf("proto: unexpected end of group")
)
