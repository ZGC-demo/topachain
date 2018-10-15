// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consensus.proto

package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TxStatus int32

const (
	TX_OK TxStatus = 0
)

var TxStatus_name = map[int32]string{
	0: "TX_OK",
}
var TxStatus_value = map[string]int32{
	"TX_OK": 0,
}

func (TxStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptorConsensus, []int{0} }

// consensus block config
type ConsensusBlockConfig struct {
	Timeout      string `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	MaxTxCount   uint64 `protobuf:"varint,2,opt,name=maxTxCount,proto3" json:"maxTxCount,omitempty"`
	MaxTxSize    uint64 `protobuf:"varint,3,opt,name=maxTxSize,proto3" json:"maxTxSize,omitempty"`
	MaxBlockSize uint64 `protobuf:"varint,4,opt,name=maxBlockSize,proto3" json:"maxBlockSize,omitempty"`
}

func (m *ConsensusBlockConfig) Reset()                    { *m = ConsensusBlockConfig{} }
func (*ConsensusBlockConfig) ProtoMessage()               {}
func (*ConsensusBlockConfig) Descriptor() ([]byte, []int) { return fileDescriptorConsensus, []int{0} }

func (m *ConsensusBlockConfig) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

func (m *ConsensusBlockConfig) GetMaxTxCount() uint64 {
	if m != nil {
		return m.MaxTxCount
	}
	return 0
}

func (m *ConsensusBlockConfig) GetMaxTxSize() uint64 {
	if m != nil {
		return m.MaxTxSize
	}
	return 0
}

func (m *ConsensusBlockConfig) GetMaxBlockSize() uint64 {
	if m != nil {
		return m.MaxBlockSize
	}
	return 0
}

//
type TxResponseSync struct {
	Id     string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status TxStatus `protobuf:"varint,2,opt,name=status,proto3,enum=types.TxStatus" json:"status,omitempty"`
}

func (m *TxResponseSync) Reset()                    { *m = TxResponseSync{} }
func (*TxResponseSync) ProtoMessage()               {}
func (*TxResponseSync) Descriptor() ([]byte, []int) { return fileDescriptorConsensus, []int{1} }

func (m *TxResponseSync) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TxResponseSync) GetStatus() TxStatus {
	if m != nil {
		return m.Status
	}
	return TX_OK
}

func init() {
	proto.RegisterType((*ConsensusBlockConfig)(nil), "types.consensusBlockConfig")
	proto.RegisterType((*TxResponseSync)(nil), "types.txResponseSync")
	proto.RegisterEnum("types.TxStatus", TxStatus_name, TxStatus_value)
}
func (x TxStatus) String() string {
	s, ok := TxStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *ConsensusBlockConfig) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ConsensusBlockConfig)
	if !ok {
		that2, ok := that.(ConsensusBlockConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Timeout != that1.Timeout {
		return false
	}
	if this.MaxTxCount != that1.MaxTxCount {
		return false
	}
	if this.MaxTxSize != that1.MaxTxSize {
		return false
	}
	if this.MaxBlockSize != that1.MaxBlockSize {
		return false
	}
	return true
}
func (this *TxResponseSync) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TxResponseSync)
	if !ok {
		that2, ok := that.(TxResponseSync)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (this *ConsensusBlockConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&types.ConsensusBlockConfig{")
	s = append(s, "Timeout: "+fmt.Sprintf("%#v", this.Timeout)+",\n")
	s = append(s, "MaxTxCount: "+fmt.Sprintf("%#v", this.MaxTxCount)+",\n")
	s = append(s, "MaxTxSize: "+fmt.Sprintf("%#v", this.MaxTxSize)+",\n")
	s = append(s, "MaxBlockSize: "+fmt.Sprintf("%#v", this.MaxBlockSize)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TxResponseSync) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&types.TxResponseSync{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringConsensus(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ConsensusBlockConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusBlockConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Timeout) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConsensus(dAtA, i, uint64(len(m.Timeout)))
		i += copy(dAtA[i:], m.Timeout)
	}
	if m.MaxTxCount != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConsensus(dAtA, i, uint64(m.MaxTxCount))
	}
	if m.MaxTxSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintConsensus(dAtA, i, uint64(m.MaxTxSize))
	}
	if m.MaxBlockSize != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintConsensus(dAtA, i, uint64(m.MaxBlockSize))
	}
	return i, nil
}

func (m *TxResponseSync) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxResponseSync) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConsensus(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Status != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConsensus(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func encodeFixed64Consensus(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Consensus(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConsensus(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ConsensusBlockConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Timeout)
	if l > 0 {
		n += 1 + l + sovConsensus(uint64(l))
	}
	if m.MaxTxCount != 0 {
		n += 1 + sovConsensus(uint64(m.MaxTxCount))
	}
	if m.MaxTxSize != 0 {
		n += 1 + sovConsensus(uint64(m.MaxTxSize))
	}
	if m.MaxBlockSize != 0 {
		n += 1 + sovConsensus(uint64(m.MaxBlockSize))
	}
	return n
}

func (m *TxResponseSync) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovConsensus(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovConsensus(uint64(m.Status))
	}
	return n
}

func sovConsensus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConsensus(x uint64) (n int) {
	return sovConsensus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ConsensusBlockConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConsensusBlockConfig{`,
		`Timeout:` + fmt.Sprintf("%v", this.Timeout) + `,`,
		`MaxTxCount:` + fmt.Sprintf("%v", this.MaxTxCount) + `,`,
		`MaxTxSize:` + fmt.Sprintf("%v", this.MaxTxSize) + `,`,
		`MaxBlockSize:` + fmt.Sprintf("%v", this.MaxBlockSize) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TxResponseSync) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TxResponseSync{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConsensus(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ConsensusBlockConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensus
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
			return fmt.Errorf("proto: consensusBlockConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: consensusBlockConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConsensus
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timeout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTxCount", wireType)
			}
			m.MaxTxCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTxCount |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTxSize", wireType)
			}
			m.MaxTxSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTxSize |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBlockSize", wireType)
			}
			m.MaxBlockSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBlockSize |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConsensus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConsensus
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
func (m *TxResponseSync) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensus
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
			return fmt.Errorf("proto: txResponseSync: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: txResponseSync: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConsensus
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (TxStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConsensus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConsensus
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
func skipConsensus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConsensus
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
					return 0, ErrIntOverflowConsensus
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
					return 0, ErrIntOverflowConsensus
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
				return 0, ErrInvalidLengthConsensus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConsensus
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
				next, err := skipConsensus(dAtA[start:])
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
	ErrInvalidLengthConsensus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConsensus   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("consensus.proto", fileDescriptorConsensus) }

var fileDescriptorConsensus = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4a, 0x33, 0x41,
	0x14, 0xc5, 0x77, 0xf2, 0x25, 0xf9, 0xcc, 0x45, 0x12, 0x19, 0x14, 0xb6, 0x90, 0x4b, 0x88, 0x85,
	0xc1, 0x62, 0x03, 0xfa, 0x06, 0x49, 0x25, 0x16, 0xe2, 0x6e, 0x0a, 0xb1, 0x91, 0xc9, 0x66, 0xcd,
	0x0e, 0xba, 0x33, 0x8b, 0x73, 0x17, 0x36, 0xa9, 0x7c, 0x04, 0xf1, 0x29, 0x7c, 0x14, 0xcb, 0x94,
	0x96, 0xee, 0xd8, 0x58, 0xe6, 0x11, 0xc4, 0x21, 0xf1, 0x4f, 0x79, 0x7e, 0xe7, 0xc0, 0x39, 0xf7,
	0x42, 0x27, 0xd6, 0xca, 0x24, 0xca, 0x14, 0x26, 0xc8, 0xef, 0x35, 0x69, 0xde, 0xa0, 0x79, 0x9e,
	0x98, 0xde, 0x13, 0x83, 0xdd, 0x6f, 0x6b, 0x78, 0xa7, 0xe3, 0xdb, 0x91, 0x56, 0x37, 0x72, 0xc6,
	0x7d, 0xf8, 0x4f, 0x32, 0x4b, 0x74, 0x41, 0x3e, 0xeb, 0xb2, 0x7e, 0x2b, 0xdc, 0x48, 0x8e, 0x00,
	0x99, 0x28, 0xc7, 0xe5, 0x48, 0x17, 0x8a, 0xfc, 0x5a, 0x97, 0xf5, 0xeb, 0xe1, 0x2f, 0xc2, 0xf7,
	0xa1, 0xe5, 0x54, 0x24, 0x17, 0x89, 0xff, 0xcf, 0xd9, 0x3f, 0x80, 0xf7, 0x60, 0x3b, 0x13, 0xa5,
	0x6b, 0x72, 0x81, 0xba, 0x0b, 0xfc, 0x61, 0xbd, 0x53, 0x68, 0x53, 0x19, 0x26, 0x26, 0xff, 0x1a,
	0x16, 0xcd, 0x55, 0xcc, 0xdb, 0x50, 0x93, 0xd3, 0xf5, 0x90, 0x9a, 0x9c, 0xf2, 0x43, 0x68, 0x1a,
	0x12, 0x54, 0x18, 0xd7, 0xdf, 0x3e, 0xee, 0x04, 0xee, 0x9c, 0x80, 0xca, 0xc8, 0xe1, 0x70, 0x6d,
	0x1f, 0xed, 0xc1, 0xd6, 0x86, 0xf1, 0x16, 0x34, 0xc6, 0x97, 0xd7, 0xe7, 0x67, 0x3b, 0xde, 0xf0,
	0x62, 0x59, 0xa1, 0xf7, 0x5a, 0xa1, 0xb7, 0xaa, 0x90, 0x3d, 0x58, 0x64, 0xcf, 0x16, 0xd9, 0x8b,
	0x45, 0xb6, 0xb4, 0xc8, 0xde, 0x2c, 0xb2, 0x0f, 0x8b, 0xde, 0xca, 0x22, 0x7b, 0x7c, 0x47, 0xef,
	0xea, 0x60, 0x26, 0x29, 0x2d, 0x26, 0x41, 0xac, 0xb3, 0x41, 0x26, 0x15, 0x2d, 0x52, 0xa1, 0x07,
	0xa4, 0x73, 0x11, 0xa7, 0x42, 0xaa, 0x81, 0xab, 0x9e, 0x34, 0xdd, 0x5f, 0x4f, 0x3e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xf2, 0xe1, 0x0d, 0xaa, 0x6a, 0x01, 0x00, 0x00,
}
