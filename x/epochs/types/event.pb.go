// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/epochs/event.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventEndEpoch struct {
	CurrentEpoch int64 `protobuf:"varint,1,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
}

func (m *EventEndEpoch) Reset()         { *m = EventEndEpoch{} }
func (m *EventEndEpoch) String() string { return proto.CompactTextString(m) }
func (*EventEndEpoch) ProtoMessage()    {}
func (*EventEndEpoch) Descriptor() ([]byte, []int) {
	return fileDescriptor_35437b9f5f16af7d, []int{0}
}
func (m *EventEndEpoch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEndEpoch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEndEpoch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEndEpoch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEndEpoch.Merge(m, src)
}
func (m *EventEndEpoch) XXX_Size() int {
	return m.Size()
}
func (m *EventEndEpoch) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEndEpoch.DiscardUnknown(m)
}

var xxx_messageInfo_EventEndEpoch proto.InternalMessageInfo

func (m *EventEndEpoch) GetCurrentEpoch() int64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

type EventBeginEpoch struct {
	CurrentEpoch int64     `protobuf:"varint,1,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
	StartTime    time.Time `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
}

func (m *EventBeginEpoch) Reset()         { *m = EventBeginEpoch{} }
func (m *EventBeginEpoch) String() string { return proto.CompactTextString(m) }
func (*EventBeginEpoch) ProtoMessage()    {}
func (*EventBeginEpoch) Descriptor() ([]byte, []int) {
	return fileDescriptor_35437b9f5f16af7d, []int{1}
}
func (m *EventBeginEpoch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBeginEpoch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBeginEpoch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBeginEpoch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBeginEpoch.Merge(m, src)
}
func (m *EventBeginEpoch) XXX_Size() int {
	return m.Size()
}
func (m *EventBeginEpoch) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBeginEpoch.DiscardUnknown(m)
}

var xxx_messageInfo_EventBeginEpoch proto.InternalMessageInfo

func (m *EventBeginEpoch) GetCurrentEpoch() int64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func (m *EventBeginEpoch) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*EventEndEpoch)(nil), "pylons.epochs.EventEndEpoch")
	proto.RegisterType((*EventBeginEpoch)(nil), "pylons.epochs.EventBeginEpoch")
}

func init() { proto.RegisterFile("pylons/epochs/event.proto", fileDescriptor_35437b9f5f16af7d) }

var fileDescriptor_35437b9f5f16af7d = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xa8, 0xcc, 0xc9,
	0xcf, 0x2b, 0xd6, 0x4f, 0x2d, 0xc8, 0x4f, 0xce, 0x28, 0xd6, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x85, 0x48, 0xe9, 0x41, 0xa4, 0xa4, 0x44, 0xd2, 0xf3,
	0xd3, 0xf3, 0xc1, 0x32, 0xfa, 0x20, 0x16, 0x44, 0x91, 0x94, 0x7c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e,
	0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x5f, 0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b,
	0x00, 0x51, 0xa0, 0x64, 0xc2, 0xc5, 0xeb, 0x0a, 0x32, 0xd4, 0x35, 0x2f, 0xc5, 0x15, 0x64, 0x90,
	0x90, 0x32, 0x17, 0x6f, 0x72, 0x69, 0x51, 0x51, 0x6a, 0x5e, 0x49, 0x3c, 0xd8, 0x64, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xe6, 0x20, 0x1e, 0xa8, 0x20, 0x58, 0x91, 0xd2, 0x04, 0x46, 0x2e, 0x7e, 0xb0,
	0x36, 0xa7, 0xd4, 0xf4, 0xcc, 0x3c, 0xe2, 0x35, 0x0a, 0x45, 0x70, 0x71, 0x15, 0x97, 0x24, 0x16,
	0x95, 0xc4, 0x83, 0xdc, 0x21, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa5, 0x07, 0x71, 0xa4,
	0x1e, 0xcc, 0x91, 0x7a, 0x21, 0x30, 0x47, 0x3a, 0xc9, 0x9e, 0xb8, 0x27, 0xcf, 0xf0, 0xe9, 0x9e,
	0xbc, 0x60, 0x65, 0x62, 0x6e, 0x8e, 0x95, 0x12, 0x42, 0xaf, 0xd2, 0x84, 0xfb, 0xf2, 0x8c, 0x41,
	0x9c, 0x60, 0x01, 0x90, 0x72, 0x27, 0xb7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0x88, 0xd2, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0f, 0x00, 0x87,
	0x99, 0x6e, 0x49, 0x6a, 0x72, 0x86, 0x3e, 0x34, 0x68, 0x2b, 0x60, 0x81, 0x5b, 0x52, 0x59, 0x90,
	0x5a, 0x9c, 0xc4, 0x06, 0x76, 0x85, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x3c, 0xc0, 0x05,
	0x7a, 0x01, 0x00, 0x00,
}

func (m *EventEndEpoch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEndEpoch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEndEpoch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpoch != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventBeginEpoch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBeginEpoch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBeginEpoch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEvent(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.CurrentEpoch != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventEndEpoch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrentEpoch != 0 {
		n += 1 + sovEvent(uint64(m.CurrentEpoch))
	}
	return n
}

func (m *EventBeginEpoch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrentEpoch != 0 {
		n += 1 + sovEvent(uint64(m.CurrentEpoch))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventEndEpoch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventEndEpoch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEndEpoch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventBeginEpoch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventBeginEpoch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBeginEpoch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)