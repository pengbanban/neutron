// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/dex/pool_reserves.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_neutron_org_neutron_v4_utils_math "github.com/neutron-org/neutron/v4/utils/math"
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

type PoolReservesKey struct {
	TradePairId           *TradePairID `protobuf:"bytes,1,opt,name=trade_pair_id,json=tradePairId,proto3" json:"trade_pair_id,omitempty"`
	TickIndexTakerToMaker int64        `protobuf:"varint,2,opt,name=tick_index_taker_to_maker,json=tickIndexTakerToMaker,proto3" json:"tick_index_taker_to_maker,omitempty"`
	Fee                   uint64       `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *PoolReservesKey) Reset()         { *m = PoolReservesKey{} }
func (m *PoolReservesKey) String() string { return proto.CompactTextString(m) }
func (*PoolReservesKey) ProtoMessage()    {}
func (*PoolReservesKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fe9f734c7ad538, []int{0}
}
func (m *PoolReservesKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolReservesKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolReservesKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolReservesKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolReservesKey.Merge(m, src)
}
func (m *PoolReservesKey) XXX_Size() int {
	return m.Size()
}
func (m *PoolReservesKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolReservesKey.DiscardUnknown(m)
}

var xxx_messageInfo_PoolReservesKey proto.InternalMessageInfo

func (m *PoolReservesKey) GetTradePairId() *TradePairID {
	if m != nil {
		return m.TradePairId
	}
	return nil
}

func (m *PoolReservesKey) GetTickIndexTakerToMaker() int64 {
	if m != nil {
		return m.TickIndexTakerToMaker
	}
	return 0
}

func (m *PoolReservesKey) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type PoolReserves struct {
	Key                *PoolReservesKey                                     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ReservesMakerDenom cosmossdk_io_math.Int                                `protobuf:"bytes,2,opt,name=reserves_maker_denom,json=reservesMakerDenom,proto3,customtype=cosmossdk.io/math.Int" json:"reserves_maker_denom" yaml:"reserves_maker_denom"`
	PriceTakerToMaker  github_com_neutron_org_neutron_v4_utils_math.PrecDec `protobuf:"bytes,3,opt,name=price_taker_to_maker,json=priceTakerToMaker,proto3,customtype=github.com/neutron-org/neutron/v4/utils/math.PrecDec" json:"price_taker_to_maker" yaml:"price_taker_to_maker"`
	MakerPrice         github_com_neutron_org_neutron_v4_utils_math.PrecDec `protobuf:"bytes,5,opt,name=maker_price,json=makerPrice,proto3,customtype=github.com/neutron-org/neutron/v4/utils/math.PrecDec" json:"maker_price" yaml:"maker_price"`
}

func (m *PoolReserves) Reset()         { *m = PoolReserves{} }
func (m *PoolReserves) String() string { return proto.CompactTextString(m) }
func (*PoolReserves) ProtoMessage()    {}
func (*PoolReserves) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fe9f734c7ad538, []int{1}
}
func (m *PoolReserves) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolReserves) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolReserves.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolReserves) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolReserves.Merge(m, src)
}
func (m *PoolReserves) XXX_Size() int {
	return m.Size()
}
func (m *PoolReserves) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolReserves.DiscardUnknown(m)
}

var xxx_messageInfo_PoolReserves proto.InternalMessageInfo

func (m *PoolReserves) GetKey() *PoolReservesKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*PoolReservesKey)(nil), "neutron.dex.PoolReservesKey")
	proto.RegisterType((*PoolReserves)(nil), "neutron.dex.PoolReserves")
}

func init() { proto.RegisterFile("neutron/dex/pool_reserves.proto", fileDescriptor_f0fe9f734c7ad538) }

var fileDescriptor_f0fe9f734c7ad538 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x8b, 0xd3, 0x40,
	0x14, 0xc6, 0x3b, 0x46, 0x05, 0xa7, 0x8a, 0x3a, 0x74, 0x21, 0xae, 0x92, 0x94, 0x9c, 0x7a, 0xd9,
	0x04, 0x74, 0x0f, 0x22, 0x9e, 0x96, 0x82, 0x14, 0x11, 0x4a, 0xe8, 0xc9, 0x4b, 0xc8, 0x66, 0x9e,
	0xdd, 0xa1, 0x4d, 0x5e, 0x98, 0x99, 0x2e, 0xed, 0xd5, 0x83, 0x67, 0xef, 0x5e, 0xfd, 0x63, 0xf6,
	0xb8, 0x47, 0xf1, 0x10, 0xa4, 0xbd, 0x79, 0x5c, 0xff, 0x01, 0x99, 0x49, 0xaa, 0x6d, 0x09, 0x78,
	0xf0, 0x94, 0xf7, 0xbe, 0xf7, 0xe5, 0xe5, 0xf7, 0x4d, 0x86, 0xfa, 0x05, 0x2c, 0xb4, 0xc4, 0x22,
	0xe2, 0xb0, 0x8c, 0x4a, 0xc4, 0x79, 0x22, 0x41, 0x81, 0xbc, 0x04, 0x15, 0x96, 0x12, 0x35, 0xb2,
	0x6e, 0x63, 0x08, 0x39, 0x2c, 0x8f, 0x7b, 0x53, 0x9c, 0xa2, 0xd5, 0x23, 0x53, 0xd5, 0x96, 0xe3,
	0xbd, 0x1d, 0x5a, 0xa6, 0x1c, 0x92, 0x32, 0x15, 0x32, 0x11, 0xbc, 0x36, 0x04, 0x5f, 0x08, 0x7d,
	0x38, 0x46, 0x9c, 0xc7, 0xcd, 0xea, 0xb7, 0xb0, 0x62, 0xaf, 0xe9, 0x83, 0x3d, 0xab, 0x4b, 0xfa,
	0x64, 0xd0, 0x7d, 0xee, 0x86, 0x3b, 0xdf, 0x0b, 0x27, 0xc6, 0x31, 0x4e, 0x85, 0x1c, 0x0d, 0xe3,
	0xae, 0xfe, 0xd3, 0x70, 0xf6, 0x92, 0x3e, 0xd1, 0x22, 0x9b, 0x25, 0xa2, 0xe0, 0xb0, 0x4c, 0x74,
	0x3a, 0x03, 0x99, 0x68, 0x4c, 0x72, 0x53, 0xb8, 0xb7, 0xfa, 0x64, 0xe0, 0xc4, 0x47, 0xc6, 0x30,
	0x32, 0xf3, 0x89, 0x51, 0x27, 0xf8, 0xce, 0x3c, 0xd8, 0x23, 0xea, 0x7c, 0x00, 0x70, 0x9d, 0x3e,
	0x19, 0xdc, 0x8e, 0x4d, 0x19, 0xfc, 0x72, 0xe8, 0xfd, 0x5d, 0x3a, 0x16, 0x52, 0x67, 0x06, 0xab,
	0x06, 0xe8, 0xd9, 0x1e, 0xd0, 0x41, 0x8a, 0xd8, 0x18, 0xd9, 0x27, 0x42, 0x7b, 0xdb, 0x53, 0xab,
	0x11, 0x12, 0x0e, 0x05, 0xe6, 0x16, 0xe4, 0xde, 0xd9, 0xe4, 0xaa, 0xf2, 0x3b, 0xdf, 0x2b, 0xff,
	0x28, 0x43, 0x95, 0xa3, 0x52, 0x7c, 0x16, 0x0a, 0x8c, 0xf2, 0x54, 0x5f, 0x84, 0xa3, 0x42, 0xff,
	0xac, 0xfc, 0xd6, 0x97, 0x6f, 0x2a, 0xff, 0xe9, 0x2a, 0xcd, 0xe7, 0xaf, 0x82, 0xb6, 0x69, 0x10,
	0xb3, 0xad, 0x6c, 0x63, 0x0d, 0x8d, 0xc8, 0xbe, 0x12, 0xda, 0x2b, 0xa5, 0xc8, 0xe0, 0xf0, 0x44,
	0x1c, 0x0b, 0xa2, 0x1a, 0x90, 0xd3, 0xa9, 0xd0, 0x17, 0x8b, 0xf3, 0x30, 0xc3, 0x3c, 0x6a, 0xc2,
	0x9d, 0xa0, 0x9c, 0x6e, 0xeb, 0xe8, 0xf2, 0x34, 0x5a, 0x68, 0x31, 0x57, 0x35, 0xe3, 0x58, 0x42,
	0x36, 0x84, 0xcc, 0x70, 0xb6, 0xed, 0xfe, 0xcb, 0xd9, 0x36, 0x0d, 0xe2, 0xc7, 0x56, 0xde, 0xfb,
	0x05, 0x1f, 0x09, 0xed, 0xd6, 0x59, 0xec, 0xcc, 0xbd, 0x63, 0xe9, 0xd2, 0xff, 0xa4, 0xdb, 0x5d,
	0x79, 0x53, 0xf9, 0xac, 0x86, 0xda, 0x11, 0x83, 0x98, 0xda, 0x6e, 0x6c, 0x9a, 0xb3, 0x37, 0x57,
	0x6b, 0x8f, 0x5c, 0xaf, 0x3d, 0xf2, 0x63, 0xed, 0x91, 0xcf, 0x1b, 0xaf, 0x73, 0xbd, 0xf1, 0x3a,
	0xdf, 0x36, 0x5e, 0xe7, 0xfd, 0xc9, 0xbf, 0x01, 0x96, 0xf5, 0x55, 0x5f, 0x95, 0xa0, 0xce, 0xef,
	0xda, 0x3b, 0xfe, 0xe2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x4e, 0x36, 0xa8, 0x4a, 0x03,
	0x00, 0x00,
}

func (m *PoolReservesKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolReservesKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolReservesKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != 0 {
		i = encodeVarintPoolReserves(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x18
	}
	if m.TickIndexTakerToMaker != 0 {
		i = encodeVarintPoolReserves(dAtA, i, uint64(m.TickIndexTakerToMaker))
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
			i = encodeVarintPoolReserves(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PoolReserves) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolReserves) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolReserves) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MakerPrice.Size()
		i -= size
		if _, err := m.MakerPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolReserves(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.PriceTakerToMaker.Size()
		i -= size
		if _, err := m.PriceTakerToMaker.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolReserves(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.ReservesMakerDenom.Size()
		i -= size
		if _, err := m.ReservesMakerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolReserves(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPoolReserves(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoolReserves(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolReserves(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolReservesKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TradePairId != nil {
		l = m.TradePairId.Size()
		n += 1 + l + sovPoolReserves(uint64(l))
	}
	if m.TickIndexTakerToMaker != 0 {
		n += 1 + sovPoolReserves(uint64(m.TickIndexTakerToMaker))
	}
	if m.Fee != 0 {
		n += 1 + sovPoolReserves(uint64(m.Fee))
	}
	return n
}

func (m *PoolReserves) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovPoolReserves(uint64(l))
	}
	l = m.ReservesMakerDenom.Size()
	n += 1 + l + sovPoolReserves(uint64(l))
	l = m.PriceTakerToMaker.Size()
	n += 1 + l + sovPoolReserves(uint64(l))
	l = m.MakerPrice.Size()
	n += 1 + l + sovPoolReserves(uint64(l))
	return n
}

func sovPoolReserves(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolReserves(x uint64) (n int) {
	return sovPoolReserves(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolReservesKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolReserves
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
			return fmt.Errorf("proto: PoolReservesKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolReservesKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradePairId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
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
					return ErrIntOverflowPoolReserves
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPoolReserves(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolReserves
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
func (m *PoolReserves) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolReserves
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
			return fmt.Errorf("proto: PoolReserves: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolReserves: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Key == nil {
				m.Key = &PoolReservesKey{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesMakerDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesMakerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceTakerToMaker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PriceTakerToMaker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MakerPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolReserves(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolReserves
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
func skipPoolReserves(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolReserves
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
					return 0, ErrIntOverflowPoolReserves
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
					return 0, ErrIntOverflowPoolReserves
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
				return 0, ErrInvalidLengthPoolReserves
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolReserves
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolReserves
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolReserves        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolReserves          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolReserves = fmt.Errorf("proto: unexpected end of group")
)
