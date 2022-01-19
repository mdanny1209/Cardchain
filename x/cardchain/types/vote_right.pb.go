// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/vote_right.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type VoteRight struct {
	CardId      uint64 `protobuf:"varint,1,opt,name=cardId,proto3" json:"cardId,omitempty"`
	ExpireBlock int32  `protobuf:"varint,2,opt,name=expireBlock,proto3" json:"expireBlock,omitempty"`
}

func (m *VoteRight) Reset()         { *m = VoteRight{} }
func (m *VoteRight) String() string { return proto.CompactTextString(m) }
func (*VoteRight) ProtoMessage()    {}
func (*VoteRight) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2b8ec4267cac2c, []int{0}
}
func (m *VoteRight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteRight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteRight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteRight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRight.Merge(m, src)
}
func (m *VoteRight) XXX_Size() int {
	return m.Size()
}
func (m *VoteRight) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRight.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRight proto.InternalMessageInfo

func (m *VoteRight) GetCardId() uint64 {
	if m != nil {
		return m.CardId
	}
	return 0
}

func (m *VoteRight) GetExpireBlock() int32 {
	if m != nil {
		return m.ExpireBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*VoteRight)(nil), "DecentralCardGame.cardchain.cardchain.VoteRight")
}

func init() { proto.RegisterFile("cardchain/vote_right.proto", fileDescriptor_0d2b8ec4267cac2c) }

var fileDescriptor_0d2b8ec4267cac2c = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x4e, 0x2c, 0x4a,
	0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0xcb, 0x2f, 0x49, 0x8d, 0x2f, 0xca, 0x4c, 0xcf, 0x28,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x75, 0x49, 0x4d, 0x4e, 0xcd, 0x2b, 0x29, 0x4a,
	0xcc, 0x71, 0x4e, 0x2c, 0x4a, 0x71, 0x4f, 0xcc, 0x4d, 0xd5, 0x83, 0xab, 0x46, 0xb0, 0x94, 0x5c,
	0xb9, 0x38, 0xc3, 0xf2, 0x4b, 0x52, 0x83, 0x40, 0x3a, 0x85, 0xc4, 0xb8, 0xd8, 0x40, 0x32, 0x9e,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x50, 0x9e, 0x90, 0x02, 0x17, 0x77, 0x6a, 0x45,
	0x41, 0x66, 0x51, 0xaa, 0x53, 0x4e, 0x7e, 0x72, 0xb6, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10,
	0xb2, 0x90, 0x53, 0xd0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x59, 0xa4,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x63, 0x38, 0x49, 0xdf, 0x19, 0xee,
	0x81, 0x0a, 0x7d, 0x84, 0x67, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x1e, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xdc, 0x89, 0x22, 0xe6, 0x00, 0x00, 0x00,
}

func (m *VoteRight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteRight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteRight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpireBlock != 0 {
		i = encodeVarintVoteRight(dAtA, i, uint64(m.ExpireBlock))
		i--
		dAtA[i] = 0x10
	}
	if m.CardId != 0 {
		i = encodeVarintVoteRight(dAtA, i, uint64(m.CardId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVoteRight(dAtA []byte, offset int, v uint64) int {
	offset -= sovVoteRight(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VoteRight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CardId != 0 {
		n += 1 + sovVoteRight(uint64(m.CardId))
	}
	if m.ExpireBlock != 0 {
		n += 1 + sovVoteRight(uint64(m.ExpireBlock))
	}
	return n
}

func sovVoteRight(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVoteRight(x uint64) (n int) {
	return sovVoteRight(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VoteRight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoteRight
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
			return fmt.Errorf("proto: VoteRight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteRight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardId", wireType)
			}
			m.CardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteRight
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CardId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireBlock", wireType)
			}
			m.ExpireBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteRight
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireBlock |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVoteRight(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteRight
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
func skipVoteRight(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVoteRight
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
					return 0, ErrIntOverflowVoteRight
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
					return 0, ErrIntOverflowVoteRight
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
				return 0, ErrInvalidLengthVoteRight
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVoteRight
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVoteRight
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVoteRight        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVoteRight          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVoteRight = fmt.Errorf("proto: unexpected end of group")
)
