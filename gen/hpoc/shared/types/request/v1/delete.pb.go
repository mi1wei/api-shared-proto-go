// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hpoc/shared/types/request/v1/delete.proto

package requestv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Delete struct {
	// if true, server will hard-delete the resource
	HardDelete           bool     `protobuf:"varint,1,opt,name=hard_delete,json=hardDelete,proto3" json:"hard_delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Delete) Reset()         { *m = Delete{} }
func (m *Delete) String() string { return proto.CompactTextString(m) }
func (*Delete) ProtoMessage()    {}
func (*Delete) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c95eef138f85e62, []int{0}
}
func (m *Delete) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Delete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Delete.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Delete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delete.Merge(m, src)
}
func (m *Delete) XXX_Size() int {
	return m.Size()
}
func (m *Delete) XXX_DiscardUnknown() {
	xxx_messageInfo_Delete.DiscardUnknown(m)
}

var xxx_messageInfo_Delete proto.InternalMessageInfo

func (m *Delete) GetHardDelete() bool {
	if m != nil {
		return m.HardDelete
	}
	return false
}

func init() {
	proto.RegisterType((*Delete)(nil), "hpoc.shared.types.request.v1.Delete")
}

func init() {
	proto.RegisterFile("hpoc/shared/types/request/v1/delete.proto", fileDescriptor_4c95eef138f85e62)
}

var fileDescriptor_4c95eef138f85e62 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0x87, 0x7b, 0x41, 0x8a, 0xa4, 0x5b, 0x27, 0x87, 0x12, 0x8b, 0x53, 0x3b, 0xe4, 0x8e, 0xc3,
	0xed, 0xdc, 0x82, 0x60, 0x17, 0x21, 0xa4, 0x25, 0x83, 0x1c, 0xc8, 0x35, 0x79, 0x49, 0x02, 0xc6,
	0x8b, 0xc9, 0x35, 0xe2, 0x37, 0x71, 0x76, 0x14, 0x3f, 0x88, 0x38, 0x39, 0x3a, 0x4a, 0xdc, 0xfc,
	0x04, 0x8e, 0x72, 0x7f, 0xe6, 0x6c, 0x21, 0xcf, 0xf3, 0x3b, 0x9e, 0xd7, 0x5f, 0x97, 0x8d, 0xcc,
	0x48, 0x57, 0x8a, 0x16, 0x72, 0xa2, 0x9e, 0x1a, 0xe8, 0x48, 0x0b, 0x0f, 0x07, 0xe8, 0x14, 0xe9,
	0x29, 0xc9, 0xe1, 0x0e, 0x14, 0xe0, 0xa6, 0x95, 0x4a, 0xce, 0x17, 0x5a, 0xc5, 0x56, 0xc5, 0x46,
	0xc5, 0x4e, 0xc5, 0x3d, 0x3d, 0x5b, 0xfb, 0xd3, 0x4b, 0x63, 0xcf, 0x4f, 0xfd, 0x59, 0x29, 0xda,
	0xfc, 0xd6, 0x8e, 0x4f, 0xd0, 0x12, 0xad, 0x8e, 0x13, 0x5f, 0xff, 0xb2, 0x42, 0xf4, 0xe6, 0xbd,
	0x0f, 0x01, 0xfa, 0x1c, 0x02, 0xf4, 0x3d, 0x04, 0xe8, 0xf9, 0x27, 0x98, 0xf8, 0xcb, 0x4c, 0xd6,
	0x78, 0xec, 0xfd, 0x68, 0x66, 0xc7, 0xb1, 0x4e, 0x89, 0xd1, 0x4d, 0x5c, 0x54, 0xaa, 0x3c, 0xec,
	0x71, 0x26, 0x6b, 0x52, 0x57, 0xf4, 0x11, 0x2a, 0x22, 0x9a, 0x2a, 0xb4, 0xeb, 0xd0, 0x04, 0x87,
	0x85, 0x24, 0x05, 0xdc, 0x93, 0xb1, 0x0b, 0x2f, 0xdc, 0x67, 0x4f, 0xbf, 0x10, 0xfa, 0x43, 0x93,
	0x17, 0xef, 0x68, 0xb3, 0xdd, 0x25, 0xaf, 0xde, 0x62, 0xa3, 0x63, 0xb6, 0x36, 0x66, 0x67, 0x62,
	0x12, 0x17, 0x93, 0xd2, 0x0f, 0x8b, 0xb9, 0xc5, 0xdc, 0x60, 0xee, 0x30, 0x4f, 0xe9, 0xe0, 0xad,
	0xc6, 0x30, 0xbf, 0x8a, 0xa3, 0x6b, 0x50, 0x22, 0x17, 0x4a, 0xfc, 0x7a, 0x4b, 0xad, 0x32, 0x66,
	0x5d, 0xc6, 0x8c, 0xcc, 0x98, 0xb3, 0x19, 0x4b, 0xe9, 0x7e, 0x6a, 0xae, 0x39, 0xff, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0x2a, 0x32, 0xf1, 0xf9, 0xab, 0x01, 0x00, 0x00,
}

func (m *Delete) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Delete) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Delete) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.HardDelete {
		i--
		if m.HardDelete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelete(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelete(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Delete) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HardDelete {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDelete(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelete(x uint64) (n int) {
	return sovDelete(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Delete) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelete
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
			return fmt.Errorf("proto: Delete: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Delete: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HardDelete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelete
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
			m.HardDelete = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDelete(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelete
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDelete(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelete
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
					return 0, ErrIntOverflowDelete
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
					return 0, ErrIntOverflowDelete
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
				return 0, ErrInvalidLengthDelete
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelete
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelete
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelete        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelete          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelete = fmt.Errorf("proto: unexpected end of group")
)
