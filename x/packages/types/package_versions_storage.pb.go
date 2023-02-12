// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: packages/package_versions_storage.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/lavanet/lava/common/types"
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

type PackageVersionsStorage struct {
	PackageIndex string       `protobuf:"bytes,1,opt,name=packageIndex,proto3" json:"packageIndex,omitempty"`
	FixatedEntry *types.Entry `protobuf:"bytes,2,opt,name=fixatedEntry,proto3" json:"fixatedEntry,omitempty"`
}

func (m *PackageVersionsStorage) Reset()         { *m = PackageVersionsStorage{} }
func (m *PackageVersionsStorage) String() string { return proto.CompactTextString(m) }
func (*PackageVersionsStorage) ProtoMessage()    {}
func (*PackageVersionsStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e6b251271e68534, []int{0}
}
func (m *PackageVersionsStorage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PackageVersionsStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PackageVersionsStorage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PackageVersionsStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageVersionsStorage.Merge(m, src)
}
func (m *PackageVersionsStorage) XXX_Size() int {
	return m.Size()
}
func (m *PackageVersionsStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageVersionsStorage.DiscardUnknown(m)
}

var xxx_messageInfo_PackageVersionsStorage proto.InternalMessageInfo

func (m *PackageVersionsStorage) GetPackageIndex() string {
	if m != nil {
		return m.PackageIndex
	}
	return ""
}

func (m *PackageVersionsStorage) GetFixatedEntry() *types.Entry {
	if m != nil {
		return m.FixatedEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*PackageVersionsStorage)(nil), "lavanet.lava.packages.PackageVersionsStorage")
}

func init() {
	proto.RegisterFile("packages/package_versions_storage.proto", fileDescriptor_3e6b251271e68534)
}

var fileDescriptor_3e6b251271e68534 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x48, 0x4c, 0xce,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x87, 0x32, 0xe2, 0xcb, 0x52, 0x8b, 0x8a, 0x33, 0xf3, 0xf3, 0x8a,
	0xe3, 0x8b, 0x4b, 0xf2, 0x8b, 0x12, 0xd3, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44,
	0x73, 0x12, 0xcb, 0x12, 0xf3, 0x52, 0x4b, 0xf4, 0x40, 0xb4, 0x1e, 0x4c, 0x97, 0x94, 0x54, 0x72,
	0x7e, 0x6e, 0x6e, 0x7e, 0x9e, 0x7e, 0x6a, 0x5e, 0x49, 0x51, 0xa5, 0x5b, 0x66, 0x45, 0x62, 0x49,
	0x66, 0x7e, 0x1e, 0x44, 0x8b, 0x52, 0x0d, 0x97, 0x58, 0x00, 0x44, 0x5d, 0x18, 0xd4, 0xcc, 0x60,
	0x88, 0x91, 0x42, 0x4a, 0x5c, 0x3c, 0x50, 0x13, 0x3c, 0xf3, 0x52, 0x52, 0x2b, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x50, 0xc4, 0x84, 0xec, 0xb8, 0x78, 0xd2, 0x40, 0xe6, 0xa5, 0xa6, 0xb8,
	0x82, 0xcc, 0x96, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd2, 0x43, 0x71, 0x07, 0xc4, 0x76,
	0x3d, 0xb0, 0x8a, 0x20, 0x14, 0xf5, 0x4e, 0x4e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c,
	0xc7, 0x10, 0xa5, 0x91, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0x04, 0xd2, 0xac, 0x0f, 0x35, 0x0d, 0x4c,
	0xeb, 0x57, 0xe8, 0xc3, 0x43, 0xa3, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x11, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x26, 0x4d, 0xeb, 0x26, 0x01, 0x00, 0x00,
}

func (m *PackageVersionsStorage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PackageVersionsStorage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PackageVersionsStorage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FixatedEntry != nil {
		{
			size, err := m.FixatedEntry.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPackageVersionsStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.PackageIndex) > 0 {
		i -= len(m.PackageIndex)
		copy(dAtA[i:], m.PackageIndex)
		i = encodeVarintPackageVersionsStorage(dAtA, i, uint64(len(m.PackageIndex)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPackageVersionsStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovPackageVersionsStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PackageVersionsStorage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PackageIndex)
	if l > 0 {
		n += 1 + l + sovPackageVersionsStorage(uint64(l))
	}
	if m.FixatedEntry != nil {
		l = m.FixatedEntry.Size()
		n += 1 + l + sovPackageVersionsStorage(uint64(l))
	}
	return n
}

func sovPackageVersionsStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPackageVersionsStorage(x uint64) (n int) {
	return sovPackageVersionsStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PackageVersionsStorage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackageVersionsStorage
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
			return fmt.Errorf("proto: PackageVersionsStorage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PackageVersionsStorage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackageIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackageVersionsStorage
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
				return ErrInvalidLengthPackageVersionsStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPackageVersionsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PackageIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixatedEntry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackageVersionsStorage
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
				return ErrInvalidLengthPackageVersionsStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPackageVersionsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FixatedEntry == nil {
				m.FixatedEntry = &types.Entry{}
			}
			if err := m.FixatedEntry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPackageVersionsStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPackageVersionsStorage
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
func skipPackageVersionsStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPackageVersionsStorage
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
					return 0, ErrIntOverflowPackageVersionsStorage
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
					return 0, ErrIntOverflowPackageVersionsStorage
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
				return 0, ErrInvalidLengthPackageVersionsStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPackageVersionsStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPackageVersionsStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPackageVersionsStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPackageVersionsStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPackageVersionsStorage = fmt.Errorf("proto: unexpected end of group")
)
