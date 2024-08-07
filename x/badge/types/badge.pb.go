// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: concord/badge/badge.proto

package types

import (
	fmt "fmt"
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

type Badge struct {
	ClassId      string `protobuf:"bytes,1,opt,name=classId,proto3" json:"classId,omitempty"`
	BadgeId      string `protobuf:"bytes,2,opt,name=badgeId,proto3" json:"badgeId,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Uri          string `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	Creator      string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Owner        string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	Transferable bool   `protobuf:"varint,8,opt,name=transferable,proto3" json:"transferable,omitempty"`
}

func (m *Badge) Reset()         { *m = Badge{} }
func (m *Badge) String() string { return proto.CompactTextString(m) }
func (*Badge) ProtoMessage()    {}
func (*Badge) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a6387d6354d73e0, []int{0}
}
func (m *Badge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Badge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Badge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Badge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Badge.Merge(m, src)
}
func (m *Badge) XXX_Size() int {
	return m.Size()
}
func (m *Badge) XXX_DiscardUnknown() {
	xxx_messageInfo_Badge.DiscardUnknown(m)
}

var xxx_messageInfo_Badge proto.InternalMessageInfo

func (m *Badge) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *Badge) GetBadgeId() string {
	if m != nil {
		return m.BadgeId
	}
	return ""
}

func (m *Badge) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Badge) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Badge) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Badge) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Badge) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Badge) GetTransferable() bool {
	if m != nil {
		return m.Transferable
	}
	return false
}

func init() {
	proto.RegisterType((*Badge)(nil), "concord.badge.Badge")
}

func init() { proto.RegisterFile("concord/badge/badge.proto", fileDescriptor_4a6387d6354d73e0) }

var fileDescriptor_4a6387d6354d73e0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x63, 0xda, 0xb4, 0xc5, 0x80, 0x84, 0x2c, 0x86, 0xc7, 0x62, 0x45, 0x9d, 0x32, 0x25,
	0x03, 0x7f, 0x90, 0xad, 0x6b, 0x47, 0x36, 0xc7, 0x36, 0x25, 0x52, 0x6b, 0x47, 0xcf, 0xae, 0x80,
	0xbf, 0xe0, 0xb3, 0x18, 0xbb, 0xc1, 0x88, 0x92, 0x1f, 0x41, 0x79, 0x49, 0x25, 0xba, 0x58, 0xf7,
	0x9e, 0x63, 0xe9, 0x49, 0x97, 0x3f, 0x6a, 0xef, 0xb4, 0x47, 0x53, 0xd6, 0xca, 0xec, 0xec, 0xf8,
	0x16, 0x2d, 0xfa, 0xe8, 0xc5, 0xdd, 0xa4, 0x0a, 0x82, 0xeb, 0x6f, 0xc6, 0xd3, 0x6a, 0x48, 0x02,
	0xf8, 0x52, 0xef, 0x55, 0x08, 0x1b, 0x03, 0x2c, 0x63, 0xf9, 0xf5, 0xf6, 0x5c, 0x07, 0x43, 0x9f,
	0x37, 0x06, 0xae, 0x46, 0x33, 0x55, 0x21, 0xf8, 0xdc, 0xa9, 0x83, 0x85, 0x19, 0x61, 0xca, 0x22,
	0xe3, 0x37, 0xc6, 0x06, 0x8d, 0x4d, 0x1b, 0x1b, 0xef, 0x60, 0x4e, 0xea, 0x3f, 0x12, 0xf7, 0x7c,
	0x76, 0xc4, 0x06, 0x52, 0x32, 0x43, 0xa4, 0xdb, 0x68, 0x55, 0xf4, 0x08, 0x8b, 0xe9, 0xf6, 0x58,
	0xc5, 0x03, 0x4f, 0xfd, 0x9b, 0xb3, 0x08, 0x4b, 0xe2, 0x63, 0x11, 0x6b, 0x7e, 0x1b, 0x51, 0xb9,
	0xf0, 0x62, 0x51, 0xd5, 0x7b, 0x0b, 0xab, 0x8c, 0xe5, 0xab, 0xed, 0x05, 0xab, 0xaa, 0xaf, 0x4e,
	0xb2, 0x53, 0x27, 0xd9, 0x6f, 0x27, 0xd9, 0x67, 0x2f, 0x93, 0x53, 0x2f, 0x93, 0x9f, 0x5e, 0x26,
	0xcf, 0xf9, 0xae, 0x89, 0xaf, 0xc7, 0xba, 0xd0, 0xfe, 0x50, 0x0e, 0x6b, 0xa0, 0x8b, 0xe5, 0x79,
	0xb0, 0xf7, 0x69, 0xb2, 0xf8, 0xd1, 0xda, 0x50, 0x2f, 0x68, 0xb3, 0xa7, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x7c, 0x13, 0x2f, 0x3d, 0x50, 0x01, 0x00, 0x00,
}

func (m *Badge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Badge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Badge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Transferable {
		i--
		if m.Transferable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintBadge(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBadge(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintBadge(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintBadge(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintBadge(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BadgeId) > 0 {
		i -= len(m.BadgeId)
		copy(dAtA[i:], m.BadgeId)
		i = encodeVarintBadge(dAtA, i, uint64(len(m.BadgeId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintBadge(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBadge(dAtA []byte, offset int, v uint64) int {
	offset -= sovBadge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Badge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovBadge(uint64(l))
	}
	l = len(m.BadgeId)
	if l > 0 {
		n += 1 + l + sovBadge(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovBadge(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovBadge(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovBadge(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBadge(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovBadge(uint64(l))
	}
	if m.Transferable {
		n += 2
	}
	return n
}

func sovBadge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBadge(x uint64) (n int) {
	return sovBadge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Badge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBadge
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
			return fmt.Errorf("proto: Badge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Badge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadge
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
				return ErrInvalidLengthBadge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BadgeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadge
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
				return ErrInvalidLengthBadge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BadgeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadge
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
				return ErrInvalidLengthBadge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadge
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
				return ErrInvalidLengthBadge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadge
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
				return ErrInvalidLengthBadge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadge
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
				return ErrInvalidLengthBadge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadge
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
				return ErrInvalidLengthBadge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transferable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadge
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
			m.Transferable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBadge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBadge
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
func skipBadge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBadge
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
					return 0, ErrIntOverflowBadge
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
					return 0, ErrIntOverflowBadge
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
				return 0, ErrInvalidLengthBadge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBadge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBadge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBadge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBadge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBadge = fmt.Errorf("proto: unexpected end of group")
)
