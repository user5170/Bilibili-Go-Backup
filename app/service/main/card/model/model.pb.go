// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/main/card/model/model.proto

/*
	Package model is a generated protocol buffer package.

	It is generated from these files:
		app/service/main/card/model/model.proto

	It has these top-level messages:
		UserEquip
*/
package model

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UserEquip struct {
	Mid        int64 `protobuf:"varint,1,opt,name=Mid,proto3" json:"mid"`
	CardID     int64 `protobuf:"varint,2,opt,name=CardID,proto3" json:"card_id"`
	ExpireTime int64 `protobuf:"varint,3,opt,name=ExpireTime,proto3" json:"expire_time"`
}

func (m *UserEquip) Reset()                    { *m = UserEquip{} }
func (m *UserEquip) String() string            { return proto.CompactTextString(m) }
func (*UserEquip) ProtoMessage()               {}
func (*UserEquip) Descriptor() ([]byte, []int) { return fileDescriptorModel, []int{0} }

func init() {
	proto.RegisterType((*UserEquip)(nil), "account.service.card.UserEquip")
}
func (m *UserEquip) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserEquip) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintModel(dAtA, i, uint64(m.Mid))
	}
	if m.CardID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintModel(dAtA, i, uint64(m.CardID))
	}
	if m.ExpireTime != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintModel(dAtA, i, uint64(m.ExpireTime))
	}
	return i, nil
}

func encodeVarintModel(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserEquip) Size() (n int) {
	var l int
	_ = l
	if m.Mid != 0 {
		n += 1 + sovModel(uint64(m.Mid))
	}
	if m.CardID != 0 {
		n += 1 + sovModel(uint64(m.CardID))
	}
	if m.ExpireTime != 0 {
		n += 1 + sovModel(uint64(m.ExpireTime))
	}
	return n
}

func sovModel(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozModel(x uint64) (n int) {
	return sovModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserEquip) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserEquip: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserEquip: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardID", wireType)
			}
			m.CardID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CardID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireTime", wireType)
			}
			m.ExpireTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModel
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
func skipModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowModel
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthModel
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowModel
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipModel(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthModel = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModel   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("app/service/main/card/model/model.proto", fileDescriptorModel) }

var fileDescriptorModel = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0x2c, 0x28, 0xd0,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x4f, 0x4e, 0x2c,
	0x4a, 0xd1, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0x81, 0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x22, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x7a, 0x50, 0xc5, 0x7a, 0x20, 0x75, 0x52, 0xba,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa,
	0x60, 0xc5, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x0c, 0x51, 0xaa, 0xe1, 0xe2,
	0x0c, 0x2d, 0x4e, 0x2d, 0x72, 0x2d, 0x2c, 0xcd, 0x2c, 0x10, 0x92, 0xe4, 0x62, 0xf6, 0xcd, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x76, 0x62, 0x7f, 0x75, 0x4f, 0x9e, 0x39, 0x37, 0x33, 0x25,
	0x08, 0x24, 0x26, 0xa4, 0xcc, 0xc5, 0xe6, 0x9c, 0x58, 0x94, 0xe2, 0xe9, 0x22, 0xc1, 0x04, 0x96,
	0xe5, 0x7e, 0x75, 0x4f, 0x9e, 0x1d, 0x64, 0x61, 0x7c, 0x66, 0x4a, 0x10, 0x54, 0x4a, 0x48, 0x9f,
	0x8b, 0xcb, 0xb5, 0xa2, 0x20, 0xb3, 0x28, 0x35, 0x24, 0x33, 0x37, 0x55, 0x82, 0x19, 0xac, 0x90,
	0xff, 0xd5, 0x3d, 0x79, 0xee, 0x54, 0xb0, 0x68, 0x7c, 0x49, 0x66, 0x6e, 0x6a, 0x10, 0x92, 0x12,
	0x27, 0xe9, 0x13, 0x0f, 0xe5, 0x18, 0x2e, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0xa3, 0x58, 0xc1, 0xbe, 0x4c, 0x62, 0x03, 0xbb, 0xd0,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xac, 0x3c, 0x2e, 0x11, 0x01, 0x00, 0x00,
}