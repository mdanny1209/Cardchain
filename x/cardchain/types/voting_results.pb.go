// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/voting_results.proto

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

type VotingResults struct {
	TotalVotes              uint64          `protobuf:"varint,1,opt,name=totalVotes,proto3" json:"totalVotes,omitempty"`
	TotalFairEnoughVotes    uint64          `protobuf:"varint,2,opt,name=totalFairEnoughVotes,proto3" json:"totalFairEnoughVotes,omitempty"`
	TotalOverpoweredVotes   uint64          `protobuf:"varint,3,opt,name=totalOverpoweredVotes,proto3" json:"totalOverpoweredVotes,omitempty"`
	TotalUnderpoweredVotes  uint64          `protobuf:"varint,4,opt,name=totalUnderpoweredVotes,proto3" json:"totalUnderpoweredVotes,omitempty"`
	TotalInappropriateVotes uint64          `protobuf:"varint,5,opt,name=totalInappropriateVotes,proto3" json:"totalInappropriateVotes,omitempty"`
	CardResults             []*VotingResult `protobuf:"bytes,6,rep,name=cardResults,proto3" json:"cardResults,omitempty"`
	Notes                   string          `protobuf:"bytes,7,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (m *VotingResults) Reset()         { *m = VotingResults{} }
func (m *VotingResults) String() string { return proto.CompactTextString(m) }
func (*VotingResults) ProtoMessage()    {}
func (*VotingResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_bde1fc3e79f719ea, []int{0}
}
func (m *VotingResults) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VotingResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VotingResults.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VotingResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VotingResults.Merge(m, src)
}
func (m *VotingResults) XXX_Size() int {
	return m.Size()
}
func (m *VotingResults) XXX_DiscardUnknown() {
	xxx_messageInfo_VotingResults.DiscardUnknown(m)
}

var xxx_messageInfo_VotingResults proto.InternalMessageInfo

func (m *VotingResults) GetTotalVotes() uint64 {
	if m != nil {
		return m.TotalVotes
	}
	return 0
}

func (m *VotingResults) GetTotalFairEnoughVotes() uint64 {
	if m != nil {
		return m.TotalFairEnoughVotes
	}
	return 0
}

func (m *VotingResults) GetTotalOverpoweredVotes() uint64 {
	if m != nil {
		return m.TotalOverpoweredVotes
	}
	return 0
}

func (m *VotingResults) GetTotalUnderpoweredVotes() uint64 {
	if m != nil {
		return m.TotalUnderpoweredVotes
	}
	return 0
}

func (m *VotingResults) GetTotalInappropriateVotes() uint64 {
	if m != nil {
		return m.TotalInappropriateVotes
	}
	return 0
}

func (m *VotingResults) GetCardResults() []*VotingResult {
	if m != nil {
		return m.CardResults
	}
	return nil
}

func (m *VotingResults) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func init() {
	proto.RegisterType((*VotingResults)(nil), "DecentralCardGame.cardchain.cardchain.VotingResults")
}

func init() { proto.RegisterFile("cardchain/voting_results.proto", fileDescriptor_bde1fc3e79f719ea) }

var fileDescriptor_bde1fc3e79f719ea = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x3b, 0xfd, 0xfa, 0xf3, 0x9f, 0xe2, 0x66, 0xa8, 0x5a, 0x04, 0x87, 0x22, 0x08, 0x59,
	0x25, 0xd0, 0x8a, 0x74, 0x6d, 0xfd, 0xc0, 0x95, 0x10, 0x68, 0x17, 0x6e, 0x64, 0x9a, 0x0c, 0x49,
	0x20, 0x9d, 0x19, 0x26, 0x93, 0xaa, 0x6f, 0xe1, 0x63, 0xb9, 0xec, 0xd2, 0xa5, 0x24, 0x2f, 0xe1,
	0x52, 0xbc, 0x23, 0x6d, 0xd0, 0x06, 0xdc, 0xcd, 0xdc, 0x73, 0x7f, 0xe7, 0xc2, 0x39, 0x98, 0x06,
	0x4c, 0x87, 0x41, 0xcc, 0x12, 0xe1, 0xad, 0xa4, 0x49, 0x44, 0xf4, 0xa0, 0x79, 0x96, 0xa7, 0x26,
	0x73, 0x95, 0x96, 0x46, 0x92, 0xd3, 0x4b, 0x1e, 0x70, 0x61, 0x34, 0x4b, 0xa7, 0x4c, 0x87, 0x37,
	0x6c, 0xc9, 0xdd, 0x0d, 0xb1, 0x7d, 0x1d, 0x1d, 0xd7, 0xd8, 0x58, 0x97, 0x93, 0x8f, 0x26, 0xde,
	0x9b, 0xc3, 0xdc, 0xb7, 0xee, 0x84, 0x62, 0x6c, 0xa4, 0x61, 0xe9, 0x5c, 0x1a, 0x9e, 0x0d, 0xd0,
	0x10, 0x39, 0x6d, 0xbf, 0x32, 0x21, 0x23, 0xdc, 0x87, 0xdf, 0x35, 0x4b, 0xf4, 0x95, 0x90, 0x79,
	0x14, 0xdb, 0xcd, 0x26, 0x6c, 0xee, 0xd4, 0xc8, 0x19, 0xde, 0x87, 0xf9, 0xdd, 0x8a, 0x6b, 0x25,
	0x1f, 0xb9, 0xe6, 0xa1, 0x85, 0x5a, 0x00, 0xed, 0x16, 0xc9, 0x39, 0x3e, 0x00, 0x61, 0x26, 0xc2,
	0x1f, 0x58, 0x1b, 0xb0, 0x1a, 0x95, 0x4c, 0xf0, 0x21, 0x28, 0xb7, 0x82, 0x29, 0xa5, 0xa5, 0xd2,
	0x09, 0x33, 0xdc, 0x82, 0x1d, 0x00, 0xeb, 0x64, 0x32, 0xc3, 0xbd, 0xaf, 0xb8, 0xbe, 0xa3, 0x18,
	0x74, 0x87, 0x2d, 0xa7, 0x37, 0x1a, 0xbb, 0x7f, 0x4a, 0xda, 0xad, 0xc6, 0xe8, 0x57, 0x7d, 0x48,
	0x1f, 0x77, 0x04, 0x9c, 0xff, 0x37, 0x44, 0xce, 0x7f, 0xdf, 0x7e, 0x2e, 0xfc, 0xd7, 0x82, 0xa2,
	0x75, 0x41, 0xd1, 0x7b, 0x41, 0xd1, 0x4b, 0x49, 0x1b, 0xeb, 0x92, 0x36, 0xde, 0x4a, 0xda, 0xb8,
	0x9f, 0x44, 0x89, 0x89, 0xf3, 0x85, 0x1b, 0xc8, 0xa5, 0xf7, 0xeb, 0xb6, 0x37, 0xdd, 0x14, 0xfa,
	0xe4, 0x6d, 0xcb, 0x35, 0xcf, 0x8a, 0x67, 0x8b, 0x2e, 0xb4, 0x3a, 0xfe, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0x81, 0x8a, 0xf2, 0x77, 0x3d, 0x02, 0x00, 0x00,
}

func (m *VotingResults) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VotingResults) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VotingResults) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Notes) > 0 {
		i -= len(m.Notes)
		copy(dAtA[i:], m.Notes)
		i = encodeVarintVotingResults(dAtA, i, uint64(len(m.Notes)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.CardResults) > 0 {
		for iNdEx := len(m.CardResults) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CardResults[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVotingResults(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.TotalInappropriateVotes != 0 {
		i = encodeVarintVotingResults(dAtA, i, uint64(m.TotalInappropriateVotes))
		i--
		dAtA[i] = 0x28
	}
	if m.TotalUnderpoweredVotes != 0 {
		i = encodeVarintVotingResults(dAtA, i, uint64(m.TotalUnderpoweredVotes))
		i--
		dAtA[i] = 0x20
	}
	if m.TotalOverpoweredVotes != 0 {
		i = encodeVarintVotingResults(dAtA, i, uint64(m.TotalOverpoweredVotes))
		i--
		dAtA[i] = 0x18
	}
	if m.TotalFairEnoughVotes != 0 {
		i = encodeVarintVotingResults(dAtA, i, uint64(m.TotalFairEnoughVotes))
		i--
		dAtA[i] = 0x10
	}
	if m.TotalVotes != 0 {
		i = encodeVarintVotingResults(dAtA, i, uint64(m.TotalVotes))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVotingResults(dAtA []byte, offset int, v uint64) int {
	offset -= sovVotingResults(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VotingResults) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalVotes != 0 {
		n += 1 + sovVotingResults(uint64(m.TotalVotes))
	}
	if m.TotalFairEnoughVotes != 0 {
		n += 1 + sovVotingResults(uint64(m.TotalFairEnoughVotes))
	}
	if m.TotalOverpoweredVotes != 0 {
		n += 1 + sovVotingResults(uint64(m.TotalOverpoweredVotes))
	}
	if m.TotalUnderpoweredVotes != 0 {
		n += 1 + sovVotingResults(uint64(m.TotalUnderpoweredVotes))
	}
	if m.TotalInappropriateVotes != 0 {
		n += 1 + sovVotingResults(uint64(m.TotalInappropriateVotes))
	}
	if len(m.CardResults) > 0 {
		for _, e := range m.CardResults {
			l = e.Size()
			n += 1 + l + sovVotingResults(uint64(l))
		}
	}
	l = len(m.Notes)
	if l > 0 {
		n += 1 + l + sovVotingResults(uint64(l))
	}
	return n
}

func sovVotingResults(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVotingResults(x uint64) (n int) {
	return sovVotingResults(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VotingResults) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVotingResults
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
			return fmt.Errorf("proto: VotingResults: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VotingResults: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotes", wireType)
			}
			m.TotalVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVotingResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalFairEnoughVotes", wireType)
			}
			m.TotalFairEnoughVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVotingResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalFairEnoughVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalOverpoweredVotes", wireType)
			}
			m.TotalOverpoweredVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVotingResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalOverpoweredVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalUnderpoweredVotes", wireType)
			}
			m.TotalUnderpoweredVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVotingResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalUnderpoweredVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalInappropriateVotes", wireType)
			}
			m.TotalInappropriateVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVotingResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalInappropriateVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardResults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVotingResults
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
				return ErrInvalidLengthVotingResults
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVotingResults
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CardResults = append(m.CardResults, &VotingResult{})
			if err := m.CardResults[len(m.CardResults)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVotingResults
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
				return ErrInvalidLengthVotingResults
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVotingResults
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVotingResults(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVotingResults
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
func skipVotingResults(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVotingResults
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
					return 0, ErrIntOverflowVotingResults
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
					return 0, ErrIntOverflowVotingResults
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
				return 0, ErrInvalidLengthVotingResults
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVotingResults
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVotingResults
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVotingResults        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVotingResults          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVotingResults = fmt.Errorf("proto: unexpected end of group")
)
