// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/tick_liquidity.proto

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

type TickLiquidity struct {
	// Types that are valid to be assigned to Liquidity:
	//	*TickLiquidity_PoolReserves
	//	*TickLiquidity_LimitOrderTranche
	Liquidity isTickLiquidity_Liquidity `protobuf_oneof:"liquidity"`
}

func (m *TickLiquidity) Reset()         { *m = TickLiquidity{} }
func (m *TickLiquidity) String() string { return proto.CompactTextString(m) }
func (*TickLiquidity) ProtoMessage()    {}
func (*TickLiquidity) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bf4777d3c75e20c, []int{0}
}
func (m *TickLiquidity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickLiquidity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickLiquidity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickLiquidity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickLiquidity.Merge(m, src)
}
func (m *TickLiquidity) XXX_Size() int {
	return m.Size()
}
func (m *TickLiquidity) XXX_DiscardUnknown() {
	xxx_messageInfo_TickLiquidity.DiscardUnknown(m)
}

var xxx_messageInfo_TickLiquidity proto.InternalMessageInfo

type isTickLiquidity_Liquidity interface {
	isTickLiquidity_Liquidity()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TickLiquidity_PoolReserves struct {
	PoolReserves *PoolReserves `protobuf:"bytes,1,opt,name=poolReserves,proto3,oneof" json:"poolReserves,omitempty"`
}
type TickLiquidity_LimitOrderTranche struct {
	LimitOrderTranche *LimitOrderTranche `protobuf:"bytes,2,opt,name=limitOrderTranche,proto3,oneof" json:"limitOrderTranche,omitempty"`
}

func (*TickLiquidity_PoolReserves) isTickLiquidity_Liquidity()      {}
func (*TickLiquidity_LimitOrderTranche) isTickLiquidity_Liquidity() {}

func (m *TickLiquidity) GetLiquidity() isTickLiquidity_Liquidity {
	if m != nil {
		return m.Liquidity
	}
	return nil
}

func (m *TickLiquidity) GetPoolReserves() *PoolReserves {
	if x, ok := m.GetLiquidity().(*TickLiquidity_PoolReserves); ok {
		return x.PoolReserves
	}
	return nil
}

func (m *TickLiquidity) GetLimitOrderTranche() *LimitOrderTranche {
	if x, ok := m.GetLiquidity().(*TickLiquidity_LimitOrderTranche); ok {
		return x.LimitOrderTranche
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TickLiquidity) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TickLiquidity_PoolReserves)(nil),
		(*TickLiquidity_LimitOrderTranche)(nil),
	}
}

func init() {
	proto.RegisterType((*TickLiquidity)(nil), "duality.dex.TickLiquidity")
}

func init() { proto.RegisterFile("duality/dex/tick_liquidity.proto", fileDescriptor_1bf4777d3c75e20c) }

var fileDescriptor_1bf4777d3c75e20c = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x29, 0x4d, 0xcc,
	0xc9, 0x2c, 0xa9, 0xd4, 0x4f, 0x49, 0xad, 0xd0, 0x2f, 0xc9, 0x4c, 0xce, 0x8e, 0xcf, 0xc9, 0x2c,
	0x2c, 0xcd, 0x4c, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xaa,
	0xd0, 0x4b, 0x49, 0xad, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x8b, 0xeb, 0x83, 0x58, 0x10,
	0x25, 0x52, 0xaa, 0xc8, 0x86, 0xe4, 0x64, 0xe6, 0x66, 0x96, 0xc4, 0xe7, 0x17, 0xa5, 0xa4, 0x16,
	0xc5, 0x97, 0x14, 0x25, 0xe6, 0x25, 0x67, 0xa4, 0x42, 0x95, 0xc9, 0x23, 0x2b, 0x2b, 0xc8, 0xcf,
	0xcf, 0x89, 0x2f, 0x4a, 0x2d, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x86, 0x28, 0x50, 0x5a, 0xcb, 0xc8,
	0xc5, 0x1b, 0x92, 0x99, 0x9c, 0xed, 0x03, 0x73, 0x82, 0x90, 0x3d, 0x17, 0x0f, 0x48, 0x61, 0x10,
	0x54, 0x9d, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xa4, 0x1e, 0x92, 0x9b, 0xf4, 0x02, 0x90,
	0x14, 0x78, 0x30, 0x04, 0xa1, 0x68, 0x10, 0xf2, 0xe3, 0x12, 0x04, 0x3b, 0xc8, 0x1f, 0xe4, 0x9e,
	0x10, 0x88, 0x73, 0x24, 0x98, 0xc0, 0xa6, 0xc8, 0xa1, 0x98, 0xe2, 0x83, 0xae, 0xca, 0x83, 0x21,
	0x08, 0x53, 0xab, 0x13, 0x37, 0x17, 0x27, 0x3c, 0x80, 0x9c, 0x5c, 0x4f, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x3b, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x57, 0x1f, 0x6a, 0x8b, 0x6e, 0x4e, 0x62, 0x52, 0x31, 0x8c, 0xa3, 0x5f, 0x01, 0x09, 0xf0, 0xca,
	0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xef, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x4a,
	0xd5, 0x5a, 0x8c, 0x01, 0x00, 0x00,
}

func (m *TickLiquidity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickLiquidity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickLiquidity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Liquidity != nil {
		{
			size := m.Liquidity.Size()
			i -= size
			if _, err := m.Liquidity.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *TickLiquidity_PoolReserves) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickLiquidity_PoolReserves) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PoolReserves != nil {
		{
			size, err := m.PoolReserves.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTickLiquidity(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *TickLiquidity_LimitOrderTranche) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickLiquidity_LimitOrderTranche) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.LimitOrderTranche != nil {
		{
			size, err := m.LimitOrderTranche.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTickLiquidity(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintTickLiquidity(dAtA []byte, offset int, v uint64) int {
	offset -= sovTickLiquidity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TickLiquidity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Liquidity != nil {
		n += m.Liquidity.Size()
	}
	return n
}

func (m *TickLiquidity_PoolReserves) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolReserves != nil {
		l = m.PoolReserves.Size()
		n += 1 + l + sovTickLiquidity(uint64(l))
	}
	return n
}
func (m *TickLiquidity_LimitOrderTranche) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LimitOrderTranche != nil {
		l = m.LimitOrderTranche.Size()
		n += 1 + l + sovTickLiquidity(uint64(l))
	}
	return n
}

func sovTickLiquidity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTickLiquidity(x uint64) (n int) {
	return sovTickLiquidity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TickLiquidity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTickLiquidity
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
			return fmt.Errorf("proto: TickLiquidity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickLiquidity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolReserves", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickLiquidity
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
				return ErrInvalidLengthTickLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTickLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PoolReserves{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Liquidity = &TickLiquidity_PoolReserves{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderTranche", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickLiquidity
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
				return ErrInvalidLengthTickLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTickLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &LimitOrderTranche{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Liquidity = &TickLiquidity_LimitOrderTranche{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTickLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTickLiquidity
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
func skipTickLiquidity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTickLiquidity
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
					return 0, ErrIntOverflowTickLiquidity
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
					return 0, ErrIntOverflowTickLiquidity
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
				return 0, ErrInvalidLengthTickLiquidity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTickLiquidity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTickLiquidity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTickLiquidity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTickLiquidity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTickLiquidity = fmt.Errorf("proto: unexpected end of group")
)