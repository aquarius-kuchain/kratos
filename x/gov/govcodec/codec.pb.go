// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/gov/govcodec/codec.proto

package govcodec

import (
	fmt "fmt"
	github_com_KuChain_io_kuchain_x_gov_types "github.com/KuChainNetwork/kuchain/x/gov/types"
	types "github.com/KuChainNetwork/kuchain/x/gov/types"
	proposal "github.com/KuChainNetwork/kuchain/x/params/types/proposal"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// MsgSubmitProposal defines the application-level message type for handling
// governance proposals.
type CodecMsgSubmitProposal struct {
	types.MsgSubmitProposalBase `protobuf:"bytes,1,opt,name=base,proto3,embedded=base" json:"base"`
	Content                     *CodecContent `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *CodecMsgSubmitProposal) Reset()         { *m = CodecMsgSubmitProposal{} }
func (m *CodecMsgSubmitProposal) String() string { return proto.CompactTextString(m) }
func (*CodecMsgSubmitProposal) ProtoMessage()    {}
func (*CodecMsgSubmitProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c701bf1d9973c34, []int{0}
}
func (m *CodecMsgSubmitProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodecMsgSubmitProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodecMsgSubmitProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodecMsgSubmitProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodecMsgSubmitProposal.Merge(m, src)
}
func (m *CodecMsgSubmitProposal) XXX_Size() int {
	return m.Size()
}
func (m *CodecMsgSubmitProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CodecMsgSubmitProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CodecMsgSubmitProposal proto.InternalMessageInfo

// Proposal defines the application-level concrete proposal type used in governance
// proposals.
type CodecProposal struct {
	types.ProposalBase `protobuf:"bytes,1,opt,name=base,proto3,embedded=base" json:"base"`
	Content            CodecContent `protobuf:"bytes,2,opt,name=content,proto3" json:"content"`
}

func (m *CodecProposal) Reset()         { *m = CodecProposal{} }
func (m *CodecProposal) String() string { return proto.CompactTextString(m) }
func (*CodecProposal) ProtoMessage()    {}
func (*CodecProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c701bf1d9973c34, []int{1}
}
func (m *CodecProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodecProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodecProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodecProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodecProposal.Merge(m, src)
}
func (m *CodecProposal) XXX_Size() int {
	return m.Size()
}
func (m *CodecProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CodecProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CodecProposal proto.InternalMessageInfo

func (m *CodecProposal) GetContent() CodecContent {
	if m != nil {
		return m.Content
	}
	return CodecContent{}
}

// Content defines the application-level allowed Content to be included in a
// governance proposal.
type CodecContent struct {
	// sum defines a set of all acceptable concrete governance proposal Content types.
	//
	// Types that are valid to be assigned to Sum:
	//	*CodecContent_Text
	//	*CodecContent_ParameterChange
	Sum isCodecContent_Sum `protobuf_oneof:"sum"`
}

func (m *CodecContent) Reset()         { *m = CodecContent{} }
func (m *CodecContent) String() string { return proto.CompactTextString(m) }
func (*CodecContent) ProtoMessage()    {}
func (*CodecContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c701bf1d9973c34, []int{2}
}
func (m *CodecContent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodecContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodecContent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodecContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodecContent.Merge(m, src)
}
func (m *CodecContent) XXX_Size() int {
	return m.Size()
}
func (m *CodecContent) XXX_DiscardUnknown() {
	xxx_messageInfo_CodecContent.DiscardUnknown(m)
}

var xxx_messageInfo_CodecContent proto.InternalMessageInfo

type isCodecContent_Sum interface {
	isCodecContent_Sum()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type CodecContent_Text struct {
	Text *types.TextProposal `protobuf:"bytes,1,opt,name=text,proto3,oneof" json:"text,omitempty"`
}
type CodecContent_ParameterChange struct {
	ParameterChange *proposal.ParameterChangeProposal `protobuf:"bytes,2,opt,name=parameter_change,json=parameterChange,proto3,oneof" json:"parameter_change,omitempty"`
}

func (*CodecContent_Text) isCodecContent_Sum()            {}
func (*CodecContent_ParameterChange) isCodecContent_Sum() {}

func (m *CodecContent) GetSum() isCodecContent_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *CodecContent) GetText() *types.TextProposal {
	if x, ok := m.GetSum().(*CodecContent_Text); ok {
		return x.Text
	}
	return nil
}

func (m *CodecContent) GetParameterChange() *proposal.ParameterChangeProposal {
	if x, ok := m.GetSum().(*CodecContent_ParameterChange); ok {
		return x.ParameterChange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CodecContent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CodecContent_Text)(nil),
		(*CodecContent_ParameterChange)(nil),
	}
}

func init() {
	proto.RegisterType((*CodecMsgSubmitProposal)(nil), "kuchain.x.gov.types.v1.CodecMsgSubmitProposal")
	proto.RegisterType((*CodecProposal)(nil), "kuchain.x.gov.types.v1.CodecProposal")
	proto.RegisterType((*CodecContent)(nil), "kuchain.x.gov.types.v1.CodecContent")
}

func init() { proto.RegisterFile("x/gov/govcodec/codec.proto", fileDescriptor_2c701bf1d9973c34) }

var fileDescriptor_2c701bf1d9973c34 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xaa, 0xd0, 0x4f, 0xcf,
	0x2f, 0x03, 0xe1, 0xe4, 0xfc, 0x94, 0xd4, 0x64, 0x7d, 0x30, 0xa9, 0x57, 0x50, 0x94, 0x5f, 0x92,
	0x2f, 0x24, 0x96, 0x5d, 0x9a, 0x9c, 0x91, 0x98, 0x99, 0xa7, 0x57, 0xa1, 0x97, 0x9e, 0x5f, 0xa6,
	0x57, 0x52, 0x59, 0x90, 0x5a, 0xac, 0x57, 0x66, 0x28, 0xa5, 0x5d, 0x92, 0x91, 0x59, 0x94, 0x12,
	0x5f, 0x90, 0x58, 0x54, 0x52, 0xa9, 0x0f, 0x56, 0xaa, 0x9f, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0xac,
	0x8b, 0xcc, 0x81, 0x18, 0x22, 0xa5, 0x86, 0xa9, 0x38, 0x3d, 0x3f, 0x3d, 0x1f, 0xc1, 0x82, 0xaa,
	0x13, 0x87, 0x38, 0x04, 0x6c, 0x09, 0x84, 0x84, 0x4a, 0x28, 0x57, 0xe8, 0x17, 0x24, 0x16, 0x25,
	0xe6, 0x42, 0x45, 0x41, 0x66, 0x14, 0xe4, 0x17, 0x27, 0xe6, 0x20, 0x2b, 0x52, 0x5a, 0xc9, 0xc8,
	0x25, 0xe6, 0x0c, 0x72, 0xba, 0x6f, 0x71, 0x7a, 0x70, 0x69, 0x52, 0x6e, 0x66, 0x49, 0x00, 0x54,
	0x99, 0x90, 0x2b, 0x17, 0x4b, 0x52, 0x62, 0x71, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91,
	0xba, 0x1e, 0xaa, 0xa7, 0xca, 0x0c, 0xf5, 0x30, 0xb4, 0x38, 0x25, 0x16, 0xa7, 0x3a, 0x71, 0x9c,
	0xb8, 0x27, 0xcf, 0x70, 0xe1, 0x9e, 0x3c, 0x63, 0x10, 0x58, 0xbb, 0x90, 0x1d, 0x17, 0x7b, 0x72,
	0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x89, 0x04, 0x13, 0xd8, 0x24, 0x15, 0x3d, 0xec, 0xc1, 0xa3, 0x07,
	0x76, 0x87, 0x33, 0x44, 0x6d, 0x10, 0x4c, 0x93, 0x15, 0x47, 0xc7, 0x02, 0x79, 0x86, 0x17, 0x0b,
	0xe4, 0x19, 0x95, 0x66, 0x32, 0x72, 0xf1, 0x82, 0xd5, 0xc0, 0x9d, 0x68, 0x83, 0xe2, 0x44, 0x39,
	0x4c, 0x27, 0xe2, 0x75, 0x99, 0x0b, 0x59, 0x2e, 0x73, 0x62, 0x01, 0x19, 0x83, 0x70, 0x1f, 0x0b,
	0xd8, 0x6d, 0x0f, 0x19, 0xb9, 0x78, 0x90, 0x55, 0x09, 0x99, 0x70, 0xb1, 0x94, 0xa4, 0x56, 0x94,
	0xe0, 0x76, 0x5a, 0x48, 0x6a, 0x05, 0x3c, 0xe0, 0x3c, 0x18, 0x82, 0xc0, 0xaa, 0x85, 0x22, 0xb9,
	0x04, 0xc0, 0x71, 0x96, 0x5a, 0x92, 0x5a, 0x14, 0x9f, 0x9c, 0x91, 0x98, 0x97, 0x9e, 0x0a, 0x75,
	0x9b, 0x0e, 0x92, 0x09, 0x90, 0x68, 0x05, 0xfb, 0x0f, 0xa6, 0xd8, 0x19, 0xac, 0x16, 0xc9, 0x3c,
	0xfe, 0x02, 0x54, 0x29, 0x2b, 0x4b, 0x90, 0x3b, 0x4f, 0x6d, 0xd1, 0x35, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xf7, 0x2e, 0x75, 0x06, 0x19, 0xa9, 0x9b, 0x99, 0xaf,
	0x0f, 0x35, 0x5d, 0x1f, 0x29, 0x35, 0xe9, 0xc1, 0x7c, 0xcc, 0xca, 0xc5, 0x5c, 0x5c, 0x9a, 0xeb,
	0xe4, 0x76, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x3a, 0xc4, 0x98, 0x09,
	0xcb, 0x2a, 0x49, 0x6c, 0xe0, 0xa4, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x75, 0xb6,
	0x78, 0x43, 0x03, 0x00, 0x00,
}

func (this *CodecMsgSubmitProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CodecMsgSubmitProposal)
	if !ok {
		that2, ok := that.(CodecMsgSubmitProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.MsgSubmitProposalBase.Equal(&that1.MsgSubmitProposalBase) {
		return false
	}
	if !this.Content.Equal(that1.Content) {
		return false
	}
	return true
}
func (this *CodecProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CodecProposal)
	if !ok {
		that2, ok := that.(CodecProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ProposalBase.Equal(&that1.ProposalBase) {
		return false
	}
	if !this.Content.Equal(&that1.Content) {
		return false
	}
	return true
}
func (this *CodecContent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CodecContent)
	if !ok {
		that2, ok := that.(CodecContent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Sum == nil {
		if this.Sum != nil {
			return false
		}
	} else if this.Sum == nil {
		return false
	} else if !this.Sum.Equal(that1.Sum) {
		return false
	}
	return true
}
func (this *CodecContent_Text) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CodecContent_Text)
	if !ok {
		that2, ok := that.(CodecContent_Text)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Text.Equal(that1.Text) {
		return false
	}
	return true
}
func (this *CodecContent_ParameterChange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CodecContent_ParameterChange)
	if !ok {
		that2, ok := that.(CodecContent_ParameterChange)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ParameterChange.Equal(that1.ParameterChange) {
		return false
	}
	return true
}
func (this *CodecContent) GetContent() github_com_KuChain_io_kuchain_x_gov_types.Content {
	if x := this.GetText(); x != nil {
		return x
	}
	if x := this.GetParameterChange(); x != nil {
		return x
	}
	return nil
}

func (this *CodecContent) SetContent(value github_com_KuChain_io_kuchain_x_gov_types.Content) error {
	if value == nil {
		this.Sum = nil
		return nil
	}
	switch vt := value.(type) {
	case *types.TextProposal:
		this.Sum = &CodecContent_Text{vt}
		return nil
	case types.TextProposal:
		this.Sum = &CodecContent_Text{&vt}
		return nil
	case *proposal.ParameterChangeProposal:
		this.Sum = &CodecContent_ParameterChange{vt}
		return nil
	case proposal.ParameterChangeProposal:
		this.Sum = &CodecContent_ParameterChange{&vt}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message CodecContent", value)
}

func (m *CodecMsgSubmitProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodecMsgSubmitProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodecMsgSubmitProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Content != nil {
		{
			size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.MsgSubmitProposalBase.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCodec(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CodecProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodecProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodecProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCodec(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ProposalBase.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCodec(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CodecContent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodecContent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodecContent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *CodecContent_Text) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodecContent_Text) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Text != nil {
		{
			size, err := m.Text.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *CodecContent_ParameterChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodecContent_ParameterChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ParameterChange != nil {
		{
			size, err := m.ParameterChange.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	offset -= sovCodec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CodecMsgSubmitProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MsgSubmitProposalBase.Size()
	n += 1 + l + sovCodec(uint64(l))
	if m.Content != nil {
		l = m.Content.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *CodecProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ProposalBase.Size()
	n += 1 + l + sovCodec(uint64(l))
	l = m.Content.Size()
	n += 1 + l + sovCodec(uint64(l))
	return n
}

func (m *CodecContent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *CodecContent_Text) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Text != nil {
		l = m.Text.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *CodecContent_ParameterChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ParameterChange != nil {
		l = m.ParameterChange.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CodecMsgSubmitProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: CodecMsgSubmitProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodecMsgSubmitProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgSubmitProposalBase", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MsgSubmitProposalBase.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = &CodecContent{}
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *CodecProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: CodecProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodecProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalBase", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProposalBase.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *CodecContent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: CodecContent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodecContent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.TextProposal{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &CodecContent_Text{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParameterChange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &proposal.ParameterChangeProposal{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &CodecContent_ParameterChange{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCodec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCodec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCodec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCodec = fmt.Errorf("proto: unexpected end of group")
)
