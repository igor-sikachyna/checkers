// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/checkers/system_info.proto

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

type SystemInfo struct {
	NextId        uint64 `protobuf:"varint,1,opt,name=nextId,proto3" json:"nextId,omitempty"`
	FifoHeadIndex string `protobuf:"bytes,2,opt,name=fifoHeadIndex,proto3" json:"fifoHeadIndex,omitempty"`
	FifoTailIndex string `protobuf:"bytes,3,opt,name=fifoTailIndex,proto3" json:"fifoTailIndex,omitempty"`
}

func (m *SystemInfo) Reset()         { *m = SystemInfo{} }
func (m *SystemInfo) String() string { return proto.CompactTextString(m) }
func (*SystemInfo) ProtoMessage()    {}
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fddf76acd3e854e, []int{0}
}
func (m *SystemInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SystemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SystemInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SystemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemInfo.Merge(m, src)
}
func (m *SystemInfo) XXX_Size() int {
	return m.Size()
}
func (m *SystemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SystemInfo proto.InternalMessageInfo

func (m *SystemInfo) GetNextId() uint64 {
	if m != nil {
		return m.NextId
	}
	return 0
}

func (m *SystemInfo) GetFifoHeadIndex() string {
	if m != nil {
		return m.FifoHeadIndex
	}
	return ""
}

func (m *SystemInfo) GetFifoTailIndex() string {
	if m != nil {
		return m.FifoTailIndex
	}
	return ""
}

func init() {
	proto.RegisterType((*SystemInfo)(nil), "checkers.checkers.SystemInfo")
}

func init() {
	proto.RegisterFile("checkers/checkers/system_info.proto", fileDescriptor_4fddf76acd3e854e)
}

var fileDescriptor_4fddf76acd3e854e = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xce, 0x48, 0x4d,
	0xce, 0x4e, 0x2d, 0x2a, 0xd6, 0x87, 0x33, 0x8a, 0x2b, 0x8b, 0x4b, 0x52, 0x73, 0xe3, 0x33, 0xf3,
	0xd2, 0xf2, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x61, 0x72, 0x7a, 0x30, 0x86, 0x52,
	0x01, 0x17, 0x57, 0x30, 0x58, 0x9d, 0x67, 0x5e, 0x5a, 0xbe, 0x90, 0x18, 0x17, 0x5b, 0x5e, 0x6a,
	0x45, 0x89, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x94, 0x27, 0xa4, 0xc2, 0xc5,
	0x9b, 0x96, 0x99, 0x96, 0xef, 0x91, 0x9a, 0x98, 0xe2, 0x99, 0x97, 0x92, 0x5a, 0x21, 0xc1, 0xa4,
	0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2a, 0x08, 0x53, 0x15, 0x92, 0x98, 0x99, 0x03, 0x51, 0xc5, 0x8c,
	0x50, 0x05, 0x17, 0x74, 0xf2, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0xe3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xcc, 0xf4, 0xfc, 0x22,
	0xdd, 0xe2, 0xcc, 0xec, 0xc4, 0xe4, 0x8c, 0xca, 0xbc, 0x44, 0x84, 0xa7, 0x2a, 0x10, 0xcc, 0x92,
	0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xd7, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc6,
	0x49, 0x0a, 0x8f, 0x01, 0x01, 0x00, 0x00,
}

func (m *SystemInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SystemInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SystemInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FifoTailIndex) > 0 {
		i -= len(m.FifoTailIndex)
		copy(dAtA[i:], m.FifoTailIndex)
		i = encodeVarintSystemInfo(dAtA, i, uint64(len(m.FifoTailIndex)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FifoHeadIndex) > 0 {
		i -= len(m.FifoHeadIndex)
		copy(dAtA[i:], m.FifoHeadIndex)
		i = encodeVarintSystemInfo(dAtA, i, uint64(len(m.FifoHeadIndex)))
		i--
		dAtA[i] = 0x12
	}
	if m.NextId != 0 {
		i = encodeVarintSystemInfo(dAtA, i, uint64(m.NextId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSystemInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovSystemInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SystemInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextId != 0 {
		n += 1 + sovSystemInfo(uint64(m.NextId))
	}
	l = len(m.FifoHeadIndex)
	if l > 0 {
		n += 1 + l + sovSystemInfo(uint64(l))
	}
	l = len(m.FifoTailIndex)
	if l > 0 {
		n += 1 + l + sovSystemInfo(uint64(l))
	}
	return n
}

func sovSystemInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSystemInfo(x uint64) (n int) {
	return sovSystemInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SystemInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSystemInfo
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
			return fmt.Errorf("proto: SystemInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SystemInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextId", wireType)
			}
			m.NextId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystemInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FifoHeadIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystemInfo
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
				return ErrInvalidLengthSystemInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSystemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FifoHeadIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FifoTailIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystemInfo
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
				return ErrInvalidLengthSystemInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSystemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FifoTailIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSystemInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSystemInfo
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
func skipSystemInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSystemInfo
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
					return 0, ErrIntOverflowSystemInfo
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
					return 0, ErrIntOverflowSystemInfo
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
				return 0, ErrInvalidLengthSystemInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSystemInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSystemInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSystemInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSystemInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSystemInfo = fmt.Errorf("proto: unexpected end of group")
)
