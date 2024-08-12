// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/dex/limit_order_tranche_user.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type LimitOrderTrancheUser struct {
	TradePairId           *TradePairID          `protobuf:"bytes,1,opt,name=trade_pair_id,json=tradePairId,proto3" json:"trade_pair_id,omitempty"`
	TickIndexTakerToMaker int64                 `protobuf:"varint,2,opt,name=tick_index_taker_to_maker,json=tickIndexTakerToMaker,proto3" json:"tick_index_taker_to_maker,omitempty"`
	TrancheKey            string                `protobuf:"bytes,3,opt,name=tranche_key,json=trancheKey,proto3" json:"tranche_key,omitempty"`
	Address               string                `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	SharesOwned           cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=shares_owned,json=sharesOwned,proto3,customtype=cosmossdk.io/math.Int" json:"shares_owned" yaml:"shares_owned"`
	SharesWithdrawn       cosmossdk_io_math.Int `protobuf:"bytes,6,opt,name=shares_withdrawn,json=sharesWithdrawn,proto3,customtype=cosmossdk.io/math.Int" json:"shares_withdrawn" yaml:"shares_withdrawn"`
	// TODO: remove this in next release. It is no longer used
	SharesCancelled cosmossdk_io_math.Int `protobuf:"bytes,7,opt,name=shares_cancelled,json=sharesCancelled,proto3,customtype=cosmossdk.io/math.Int" json:"shares_cancelled" yaml:"shares_cancelled"`
	OrderType       LimitOrderType        `protobuf:"varint,8,opt,name=order_type,json=orderType,proto3,enum=neutron.dex.LimitOrderType" json:"order_type,omitempty"`
}

func (m *LimitOrderTrancheUser) Reset()         { *m = LimitOrderTrancheUser{} }
func (m *LimitOrderTrancheUser) String() string { return proto.CompactTextString(m) }
func (*LimitOrderTrancheUser) ProtoMessage()    {}
func (*LimitOrderTrancheUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e5ffbd487ea05f, []int{0}
}
func (m *LimitOrderTrancheUser) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderTrancheUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderTrancheUser.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderTrancheUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderTrancheUser.Merge(m, src)
}
func (m *LimitOrderTrancheUser) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderTrancheUser) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderTrancheUser.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderTrancheUser proto.InternalMessageInfo

func (m *LimitOrderTrancheUser) GetTradePairId() *TradePairID {
	if m != nil {
		return m.TradePairId
	}
	return nil
}

func (m *LimitOrderTrancheUser) GetTickIndexTakerToMaker() int64 {
	if m != nil {
		return m.TickIndexTakerToMaker
	}
	return 0
}

func (m *LimitOrderTrancheUser) GetTrancheKey() string {
	if m != nil {
		return m.TrancheKey
	}
	return ""
}

func (m *LimitOrderTrancheUser) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *LimitOrderTrancheUser) GetOrderType() LimitOrderType {
	if m != nil {
		return m.OrderType
	}
	return LimitOrderType_GOOD_TIL_CANCELLED
}

func init() {
	proto.RegisterType((*LimitOrderTrancheUser)(nil), "neutron.dex.LimitOrderTrancheUser")
}

func init() {
	proto.RegisterFile("neutron/dex/limit_order_tranche_user.proto", fileDescriptor_67e5ffbd487ea05f)
}

var fileDescriptor_67e5ffbd487ea05f = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4d, 0x6b, 0xdb, 0x4e,
	0x10, 0xc6, 0xad, 0x7f, 0xf2, 0x4f, 0x9a, 0x55, 0xdf, 0x50, 0x63, 0xba, 0x4d, 0x41, 0x32, 0x3e,
	0x99, 0x42, 0x24, 0x48, 0x7b, 0x28, 0xa1, 0xa7, 0x34, 0x50, 0x4c, 0x53, 0x52, 0x84, 0x4b, 0xa1,
	0x17, 0xb1, 0xd1, 0x0e, 0xd2, 0x62, 0x4b, 0x2b, 0x76, 0xd7, 0xb5, 0xf4, 0x2d, 0xfa, 0xb1, 0x72,
	0xcc, 0xb1, 0xf4, 0x20, 0x5a, 0xfb, 0xd6, 0x63, 0x3e, 0x41, 0x59, 0xeb, 0x05, 0x89, 0x1c, 0x4a,
	0x4f, 0x9a, 0x79, 0x9e, 0x67, 0xe6, 0x27, 0x89, 0x41, 0x2f, 0x52, 0x58, 0x2a, 0xc1, 0x53, 0x8f,
	0x42, 0xee, 0x2d, 0x58, 0xc2, 0x54, 0xc0, 0x05, 0x05, 0x11, 0x28, 0x41, 0xd2, 0x30, 0x86, 0x60,
	0x29, 0x41, 0xb8, 0x99, 0xe0, 0x8a, 0x5b, 0x66, 0x9d, 0x75, 0x29, 0xe4, 0x47, 0x87, 0x11, 0x8f,
	0xf8, 0x56, 0xf7, 0x74, 0x55, 0x45, 0x8e, 0x9c, 0xee, 0x3a, 0x25, 0x08, 0x85, 0x20, 0x23, 0x4c,
	0x04, 0x8c, 0xd6, 0x81, 0xc3, 0x5e, 0x20, 0xaf, 0xd4, 0xf1, 0xaf, 0x5d, 0x34, 0xbc, 0xd0, 0xf0,
	0x4b, 0xcd, 0x9e, 0x55, 0xe8, 0x4f, 0x12, 0x84, 0xf5, 0x06, 0x3d, 0xe8, 0xad, 0xc1, 0xc6, 0xc8,
	0x98, 0x98, 0x27, 0xd8, 0xed, 0xbc, 0x8b, 0x3b, 0xd3, 0x89, 0x8f, 0x84, 0x89, 0xe9, 0xb9, 0x6f,
	0xaa, 0xb6, 0xa1, 0xd6, 0x6b, 0xf4, 0x4c, 0xb1, 0x70, 0x1e, 0xb0, 0x94, 0x42, 0x1e, 0x28, 0x32,
	0xd7, 0x1f, 0xc6, 0x83, 0x44, 0x17, 0xf8, 0xbf, 0x91, 0x31, 0xd9, 0xf1, 0x87, 0x3a, 0x30, 0xd5,
	0xfe, 0x4c, 0xab, 0x33, 0xfe, 0x41, 0x3f, 0x2c, 0x07, 0x99, 0xcd, 0x1f, 0x98, 0x43, 0x81, 0x77,
	0x46, 0xc6, 0xe4, 0xc0, 0x47, 0xb5, 0xf4, 0x1e, 0x0a, 0x0b, 0xa3, 0x7d, 0x42, 0xa9, 0x00, 0x29,
	0xf1, 0xee, 0xd6, 0x6c, 0x5a, 0x2b, 0x42, 0xf7, 0x65, 0x4c, 0x04, 0xc8, 0x80, 0xaf, 0x52, 0xa0,
	0xf8, 0x7f, 0x6d, 0x9f, 0x9d, 0x5f, 0x97, 0xce, 0xe0, 0x47, 0xe9, 0x0c, 0x43, 0x2e, 0x13, 0x2e,
	0x25, 0x9d, 0xbb, 0x8c, 0x7b, 0x09, 0x51, 0xb1, 0x3b, 0x4d, 0xd5, 0xef, 0xd2, 0xe9, 0x0d, 0xdd,
	0x96, 0xce, 0x93, 0x82, 0x24, 0x8b, 0xd3, 0x71, 0x57, 0x1d, 0xfb, 0x66, 0xd5, 0x5e, 0xea, 0xce,
	0x5a, 0xa1, 0xc7, 0xb5, 0xbb, 0x62, 0x2a, 0xa6, 0x82, 0xac, 0x52, 0xbc, 0xb7, 0x85, 0x5d, 0xfc,
	0x0d, 0x76, 0x67, 0xf0, 0xb6, 0x74, 0x9e, 0xf6, 0x80, 0xad, 0x33, 0xf6, 0x1f, 0x55, 0xd2, 0xe7,
	0x46, 0xe9, 0x80, 0x43, 0x92, 0x86, 0xb0, 0x58, 0x00, 0xc5, 0xfb, 0xff, 0x06, 0x6e, 0x07, 0xef,
	0x80, 0x5b, 0xa7, 0x05, 0xbf, 0x6d, 0x14, 0xeb, 0x14, 0xa1, 0xfa, 0x3a, 0x8b, 0x0c, 0xf0, 0xbd,
	0x91, 0x31, 0x79, 0x78, 0xf2, 0xbc, 0x77, 0x0a, 0x9d, 0x2b, 0x2a, 0x32, 0xf0, 0x0f, 0x78, 0x53,
	0x9e, 0xbd, 0xbb, 0x5e, 0xdb, 0xc6, 0xcd, 0xda, 0x36, 0x7e, 0xae, 0x6d, 0xe3, 0xdb, 0xc6, 0x1e,
	0xdc, 0x6c, 0xec, 0xc1, 0xf7, 0x8d, 0x3d, 0xf8, 0x72, 0x1c, 0x31, 0x15, 0x2f, 0xaf, 0xdc, 0x90,
	0x27, 0x5e, 0xbd, 0xeb, 0x98, 0x8b, 0xa8, 0xa9, 0xbd, 0xaf, 0xaf, 0xbc, 0xbc, 0xba, 0xd7, 0x22,
	0x03, 0x79, 0xb5, 0xb7, 0xbd, 0xd9, 0x97, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x41, 0xf7,
	0x93, 0x3b, 0x03, 0x00, 0x00,
}

func (m *LimitOrderTrancheUser) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderTrancheUser) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderTrancheUser) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OrderType != 0 {
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(m.OrderType))
		i--
		dAtA[i] = 0x40
	}
	{
		size := m.SharesCancelled.Size()
		i -= size
		if _, err := m.SharesCancelled.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.SharesWithdrawn.Size()
		i -= size
		if _, err := m.SharesWithdrawn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.SharesOwned.Size()
		i -= size
		if _, err := m.SharesOwned.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TrancheKey) > 0 {
		i -= len(m.TrancheKey)
		copy(dAtA[i:], m.TrancheKey)
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(len(m.TrancheKey)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TickIndexTakerToMaker != 0 {
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(m.TickIndexTakerToMaker))
		i--
		dAtA[i] = 0x10
	}
	if m.TradePairId != nil {
		{
			size, err := m.TradePairId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLimitOrderTrancheUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitOrderTrancheUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitOrderTrancheUser) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TradePairId != nil {
		l = m.TradePairId.Size()
		n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	}
	if m.TickIndexTakerToMaker != 0 {
		n += 1 + sovLimitOrderTrancheUser(uint64(m.TickIndexTakerToMaker))
	}
	l = len(m.TrancheKey)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	}
	l = m.SharesOwned.Size()
	n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	l = m.SharesWithdrawn.Size()
	n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	l = m.SharesCancelled.Size()
	n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	if m.OrderType != 0 {
		n += 1 + sovLimitOrderTrancheUser(uint64(m.OrderType))
	}
	return n
}

func sovLimitOrderTrancheUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitOrderTrancheUser(x uint64) (n int) {
	return sovLimitOrderTrancheUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitOrderTrancheUser) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderTrancheUser
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
			return fmt.Errorf("proto: LimitOrderTrancheUser: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderTrancheUser: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradePairId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TradePairId == nil {
				m.TradePairId = &TradePairID{}
			}
			if err := m.TradePairId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndexTakerToMaker", wireType)
			}
			m.TickIndexTakerToMaker = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickIndexTakerToMaker |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrancheKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrancheKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesOwned", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesOwned.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesWithdrawn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesWithdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesCancelled", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesCancelled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderType", wireType)
			}
			m.OrderType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderType |= LimitOrderType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderTrancheUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
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
func skipLimitOrderTrancheUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitOrderTrancheUser
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
					return 0, ErrIntOverflowLimitOrderTrancheUser
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
					return 0, ErrIntOverflowLimitOrderTrancheUser
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
				return 0, ErrInvalidLengthLimitOrderTrancheUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitOrderTrancheUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitOrderTrancheUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitOrderTrancheUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitOrderTrancheUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitOrderTrancheUser = fmt.Errorf("proto: unexpected end of group")
)
