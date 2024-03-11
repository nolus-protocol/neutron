// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/feeburner/total_burned_neutrons_amount.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// TotalBurnedNeutronsAmount defines total amount of burned neutron fees
type TotalBurnedNeutronsAmount struct {
	Coin types.Coin `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin" yaml:"coin"`
}

func (m *TotalBurnedNeutronsAmount) Reset()         { *m = TotalBurnedNeutronsAmount{} }
func (m *TotalBurnedNeutronsAmount) String() string { return proto.CompactTextString(m) }
func (*TotalBurnedNeutronsAmount) ProtoMessage()    {}
func (*TotalBurnedNeutronsAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfe0eb8ae8e764c8, []int{0}
}
func (m *TotalBurnedNeutronsAmount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TotalBurnedNeutronsAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TotalBurnedNeutronsAmount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TotalBurnedNeutronsAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalBurnedNeutronsAmount.Merge(m, src)
}
func (m *TotalBurnedNeutronsAmount) XXX_Size() int {
	return m.Size()
}
func (m *TotalBurnedNeutronsAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalBurnedNeutronsAmount.DiscardUnknown(m)
}

var xxx_messageInfo_TotalBurnedNeutronsAmount proto.InternalMessageInfo

func (m *TotalBurnedNeutronsAmount) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*TotalBurnedNeutronsAmount)(nil), "neutron.feeburner.TotalBurnedNeutronsAmount")
}

func init() {
	proto.RegisterFile("neutron/feeburner/total_burned_neutrons_amount.proto", fileDescriptor_dfe0eb8ae8e764c8)
}

var fileDescriptor_dfe0eb8ae8e764c8 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc9, 0x4b, 0x2d, 0x2d,
	0x29, 0xca, 0xcf, 0xd3, 0x4f, 0x4b, 0x4d, 0x4d, 0x2a, 0x2d, 0xca, 0x4b, 0x2d, 0xd2, 0x2f, 0xc9,
	0x2f, 0x49, 0xcc, 0x89, 0x07, 0x73, 0x52, 0xe2, 0xa1, 0xd2, 0xc5, 0xf1, 0x89, 0xb9, 0xf9, 0xa5,
	0x79, 0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x82, 0x50, 0x61, 0x3d, 0xb8, 0x2e, 0x29,
	0xb9, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0x62, 0xfd, 0xa4, 0xc4, 0xe2, 0x54, 0xfd, 0x32, 0xc3, 0xa4,
	0xd4, 0x92, 0x44, 0x43, 0xfd, 0xe4, 0xfc, 0xcc, 0x3c, 0x88, 0x16, 0x29, 0x91, 0xf4, 0xfc, 0xf4,
	0x7c, 0x30, 0x53, 0x1f, 0xc4, 0x82, 0x88, 0x2a, 0xc5, 0x73, 0x49, 0x86, 0x80, 0xac, 0x73, 0x02,
	0xdb, 0xe6, 0x07, 0xb5, 0xcc, 0x11, 0x6c, 0x97, 0x90, 0x13, 0x17, 0x0b, 0xc8, 0x00, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x6e, 0x23, 0x49, 0x3d, 0x88, 0x0d, 0x7a, 0x20, 0x1b, 0xf4, 0xa0, 0x36, 0xe8,
	0x39, 0xe7, 0x67, 0xe6, 0x39, 0x09, 0x9f, 0xb8, 0x27, 0xcf, 0xf0, 0xe9, 0x9e, 0x3c, 0x77, 0x65,
	0x62, 0x6e, 0x8e, 0x95, 0x12, 0x48, 0x93, 0x52, 0x10, 0x58, 0xaf, 0x93, 0xef, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0x43, 0xbd, 0xa3, 0x9b, 0x5f, 0x94, 0x0e, 0x63, 0xeb, 0x97, 0x19, 0xeb, 0x57,
	0x20, 0x87, 0x4a, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0xd9, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x00, 0x76, 0x52, 0xde, 0x37, 0x01, 0x00, 0x00,
}

func (m *TotalBurnedNeutronsAmount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TotalBurnedNeutronsAmount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TotalBurnedNeutronsAmount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTotalBurnedNeutronsAmount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTotalBurnedNeutronsAmount(dAtA []byte, offset int, v uint64) int {
	offset -= sovTotalBurnedNeutronsAmount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TotalBurnedNeutronsAmount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Coin.Size()
	n += 1 + l + sovTotalBurnedNeutronsAmount(uint64(l))
	return n
}

func sovTotalBurnedNeutronsAmount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTotalBurnedNeutronsAmount(x uint64) (n int) {
	return sovTotalBurnedNeutronsAmount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TotalBurnedNeutronsAmount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTotalBurnedNeutronsAmount
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
			return fmt.Errorf("proto: TotalBurnedNeutronsAmount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TotalBurnedNeutronsAmount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTotalBurnedNeutronsAmount
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
				return ErrInvalidLengthTotalBurnedNeutronsAmount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTotalBurnedNeutronsAmount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTotalBurnedNeutronsAmount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTotalBurnedNeutronsAmount
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
func skipTotalBurnedNeutronsAmount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTotalBurnedNeutronsAmount
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
					return 0, ErrIntOverflowTotalBurnedNeutronsAmount
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
					return 0, ErrIntOverflowTotalBurnedNeutronsAmount
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
				return 0, ErrInvalidLengthTotalBurnedNeutronsAmount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTotalBurnedNeutronsAmount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTotalBurnedNeutronsAmount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTotalBurnedNeutronsAmount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTotalBurnedNeutronsAmount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTotalBurnedNeutronsAmount = fmt.Errorf("proto: unexpected end of group")
)
