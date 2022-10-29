// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hpoc/shared/options/event/v1/register.proto

package eventv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

type MessageRegister struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRegister) Reset()         { *m = MessageRegister{} }
func (m *MessageRegister) String() string { return proto.CompactTextString(m) }
func (*MessageRegister) ProtoMessage()    {}
func (*MessageRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_341f440650aae2cb, []int{0}
}
func (m *MessageRegister) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageRegister.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRegister.Merge(m, src)
}
func (m *MessageRegister) XXX_Size() int {
	return m.Size()
}
func (m *MessageRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRegister.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRegister proto.InternalMessageInfo

func (m *MessageRegister) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

var E_RegisterEvent = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MessageOptions)(nil),
	ExtensionType: (*MessageRegister)(nil),
	Field:         82000,
	Name:          "hpoc.shared.options.event.v1.register_event",
	Tag:           "bytes,82000,opt,name=register_event",
	Filename:      "hpoc/shared/options/event/v1/register.proto",
}

func init() {
	proto.RegisterType((*MessageRegister)(nil), "hpoc.shared.options.event.v1.MessageRegister")
	proto.RegisterExtension(E_RegisterEvent)
}

func init() {
	proto.RegisterFile("hpoc/shared/options/event/v1/register.proto", fileDescriptor_341f440650aae2cb)
}

var fileDescriptor_341f440650aae2cb = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x6a, 0xeb, 0x30,
	0x14, 0x86, 0x23, 0x73, 0x73, 0xe1, 0xfa, 0x92, 0x7b, 0x21, 0x74, 0x08, 0x25, 0xb8, 0xa6, 0x4b,
	0x03, 0xc5, 0x47, 0xa8, 0xdd, 0xd4, 0x2d, 0x10, 0x9a, 0x25, 0x4d, 0x48, 0x20, 0x43, 0x31, 0x14,
	0xc7, 0x51, 0x15, 0x41, 0x13, 0x19, 0x5b, 0x71, 0xd6, 0x3e, 0x46, 0xe7, 0x8e, 0xa5, 0x0f, 0x52,
	0x3a, 0x65, 0xec, 0x58, 0xdc, 0xad, 0x4f, 0xd0, 0xb1, 0x58, 0x92, 0x97, 0x0e, 0xde, 0x6c, 0xfe,
	0xef, 0x7c, 0xe7, 0x3f, 0x72, 0x4f, 0x57, 0x89, 0x8c, 0x71, 0xb6, 0x8a, 0x52, 0xb6, 0xc4, 0x32,
	0x51, 0x42, 0x6e, 0x32, 0xcc, 0x72, 0xb6, 0x51, 0x38, 0x27, 0x38, 0x65, 0x5c, 0x64, 0x8a, 0xa5,
	0x90, 0xa4, 0x52, 0xc9, 0x76, 0xb7, 0x84, 0xc1, 0xc0, 0x60, 0x61, 0xd0, 0x30, 0xe4, 0xe4, 0xd0,
	0xe7, 0x52, 0xf2, 0x3b, 0x86, 0x35, 0xbb, 0xd8, 0xde, 0xe2, 0x25, 0xcb, 0xe2, 0x54, 0x24, 0x4a,
	0xda, 0xf9, 0xe3, 0x13, 0xf7, 0xff, 0x88, 0x65, 0x59, 0xc4, 0xd9, 0xd4, 0x8a, 0xdb, 0x07, 0x6e,
	0x53, 0xc9, 0x44, 0xc4, 0x1d, 0xe4, 0xa3, 0xde, 0x9f, 0xa9, 0xf9, 0xa1, 0x3b, 0xf7, 0x5f, 0xb5,
	0xfa, 0x46, 0xfb, 0xdb, 0x47, 0x60, 0xec, 0x50, 0xd9, 0xc1, 0x9a, 0xc6, 0xa6, 0x46, 0x67, 0x7f,
	0xdf, 0xf4, 0x51, 0xef, 0xef, 0x59, 0x00, 0x75, 0x25, 0xe1, 0xc7, 0xfe, 0x69, 0xab, 0xda, 0x33,
	0x28, 0x89, 0xfe, 0xb3, 0xf3, 0x52, 0x78, 0x68, 0x5f, 0x78, 0xe8, 0xbd, 0xf0, 0xd0, 0xc3, 0x87,
	0xd7, 0x70, 0xfd, 0x58, 0xae, 0x6b, 0x9d, 0xfd, 0x56, 0x65, 0x9b, 0x94, 0xdd, 0x26, 0xe8, 0xfa,
	0x8a, 0x0b, 0xb5, 0xda, 0x2e, 0x20, 0x96, 0x6b, 0xbc, 0x16, 0x64, 0xc7, 0x04, 0x8e, 0x12, 0x11,
	0x98, 0xf9, 0x40, 0x5f, 0x10, 0x70, 0x89, 0x39, 0xdb, 0xe0, 0xba, 0xe7, 0xbf, 0xd0, 0x1f, 0x39,
	0x79, 0x43, 0xe8, 0x0b, 0x35, 0x1e, 0x9d, 0x5f, 0xc3, 0xd9, 0x78, 0xf0, 0xe4, 0x74, 0x87, 0x65,
	0x99, 0x99, 0x29, 0x63, 0xcf, 0x07, 0x5d, 0x1f, 0xe6, 0xe4, 0xd5, 0xc4, 0xa1, 0x89, 0x43, 0x1b,
	0x87, 0x3a, 0x0e, 0xe7, 0xa4, 0x70, 0x7a, 0x75, 0x71, 0x78, 0x39, 0xe9, 0x8f, 0x98, 0x8a, 0x96,
	0x91, 0x8a, 0x3e, 0x1d, 0xbf, 0x44, 0x29, 0x35, 0x2c, 0xa5, 0x16, 0xa6, 0x54, 0xd3, 0x94, 0xce,
	0xc9, 0xe2, 0xb7, 0xbe, 0xe5, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xf7, 0xc6, 0x23, 0x46,
	0x02, 0x00, 0x00,
}

func (m *MessageRegister) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageRegister) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MessageRegister) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Topic) > 0 {
		i -= len(m.Topic)
		copy(dAtA[i:], m.Topic)
		i = encodeVarintRegister(dAtA, i, uint64(len(m.Topic)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRegister(dAtA []byte, offset int, v uint64) int {
	offset -= sovRegister(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MessageRegister) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovRegister(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRegister(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRegister(x uint64) (n int) {
	return sovRegister(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MessageRegister) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegister
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
			return fmt.Errorf("proto: MessageRegister: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageRegister: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegister
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
				return ErrInvalidLengthRegister
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegister
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegister(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRegister
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
func skipRegister(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRegister
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
					return 0, ErrIntOverflowRegister
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
					return 0, ErrIntOverflowRegister
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
				return 0, ErrInvalidLengthRegister
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRegister
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRegister
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRegister        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRegister          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRegister = fmt.Errorf("proto: unexpected end of group")
)
