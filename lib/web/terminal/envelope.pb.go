// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: teleport/lib/web/terminal/envelope.proto

package terminal

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

// Envelope is used to wrap and transend and receive messages between the
// web client and proxy.
type Envelope struct {
	// Version is the version of the protocol.
	Version string `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	// Type is the type of message. For version 1 of the protocol this must
	// not be longer than 1 character.
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	// Payload is the actual data to send.
	Payload              string   `protobuf:"bytes,3,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_016ae8368e6afaa7, []int{0}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return m.Size()
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Envelope) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Envelope) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*Envelope)(nil), "teleport.lib.web.Envelope")
}

func init() {
	proto.RegisterFile("teleport/lib/web/terminal/envelope.proto", fileDescriptor_016ae8368e6afaa7)
}

var fileDescriptor_016ae8368e6afaa7 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x49, 0xcd, 0x49,
	0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0xc9, 0x4c, 0xd2, 0x2f, 0x4f, 0x4d, 0xd2, 0x2f, 0x49, 0x2d,
	0xca, 0xcd, 0xcc, 0x4b, 0xcc, 0xd1, 0x4f, 0xcd, 0x2b, 0x4b, 0xcd, 0xc9, 0x2f, 0x48, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0xa9, 0xd4, 0xcb, 0xc9, 0x4c, 0xd2, 0x2b, 0x4f, 0x4d,
	0x52, 0x0a, 0xe2, 0xe2, 0x70, 0x85, 0xaa, 0x11, 0x92, 0xe0, 0x62, 0x0f, 0x4b, 0x2d, 0x2a, 0xce,
	0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0x84, 0xb8, 0x58, 0x42,
	0x2a, 0x0b, 0x52, 0x25, 0x98, 0xc0, 0xc2, 0x60, 0x36, 0x48, 0x75, 0x40, 0x62, 0x65, 0x4e, 0x7e,
	0x62, 0x8a, 0x04, 0x33, 0x44, 0x35, 0x94, 0xeb, 0xe4, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46, 0x19, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0xa7, 0x17, 0x25, 0x96, 0x65, 0x96, 0x24, 0x96, 0x64, 0xe6, 0x83, 0x1c, 0x87,
	0xd3, 0xd9, 0x49, 0x6c, 0x60, 0xe7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x44, 0xc0, 0x7f,
	0x05, 0xda, 0x00, 0x00, 0x00,
}

func (m *Envelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Envelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Envelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintEnvelope(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintEnvelope(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintEnvelope(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEnvelope(dAtA []byte, offset int, v uint64) int {
	offset -= sovEnvelope(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Envelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovEnvelope(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovEnvelope(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEnvelope(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEnvelope(x uint64) (n int) {
	return sovEnvelope(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Envelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnvelope
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
			return fmt.Errorf("proto: Envelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Envelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
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
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnvelope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
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
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnvelope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
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
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnvelope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnvelope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEnvelope
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
func skipEnvelope(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnvelope
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
					return 0, ErrIntOverflowEnvelope
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
					return 0, ErrIntOverflowEnvelope
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
				return 0, ErrInvalidLengthEnvelope
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEnvelope
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEnvelope
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEnvelope        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnvelope          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEnvelope = fmt.Errorf("proto: unexpected end of group")
)
