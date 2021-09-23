// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operations.proto

package rpc

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
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

type RemoteStartTransactionReq struct {
	CpId                 string   `protobuf:"bytes,1,opt,name=cp_id,json=cpId,proto3" json:"cp_id,omitempty"`
	ConnectorId          int32    `protobuf:"varint,2,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteStartTransactionReq) Reset()         { *m = RemoteStartTransactionReq{} }
func (m *RemoteStartTransactionReq) String() string { return proto.CompactTextString(m) }
func (*RemoteStartTransactionReq) ProtoMessage()    {}
func (*RemoteStartTransactionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a5877375e491e, []int{0}
}
func (m *RemoteStartTransactionReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoteStartTransactionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoteStartTransactionReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoteStartTransactionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteStartTransactionReq.Merge(m, src)
}
func (m *RemoteStartTransactionReq) XXX_Size() int {
	return m.Size()
}
func (m *RemoteStartTransactionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteStartTransactionReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteStartTransactionReq proto.InternalMessageInfo

type RemoteStartTransactionResp struct {
	TransactionId        int32    `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteStartTransactionResp) Reset()         { *m = RemoteStartTransactionResp{} }
func (m *RemoteStartTransactionResp) String() string { return proto.CompactTextString(m) }
func (*RemoteStartTransactionResp) ProtoMessage()    {}
func (*RemoteStartTransactionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a5877375e491e, []int{1}
}
func (m *RemoteStartTransactionResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoteStartTransactionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoteStartTransactionResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoteStartTransactionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteStartTransactionResp.Merge(m, src)
}
func (m *RemoteStartTransactionResp) XXX_Size() int {
	return m.Size()
}
func (m *RemoteStartTransactionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteStartTransactionResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteStartTransactionResp proto.InternalMessageInfo

type RemoteStopTransactionReq struct {
	CpId                 string   `protobuf:"bytes,1,opt,name=cp_id,json=cpId,proto3" json:"cp_id,omitempty"`
	TransactionId        int32    `protobuf:"varint,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteStopTransactionReq) Reset()         { *m = RemoteStopTransactionReq{} }
func (m *RemoteStopTransactionReq) String() string { return proto.CompactTextString(m) }
func (*RemoteStopTransactionReq) ProtoMessage()    {}
func (*RemoteStopTransactionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a5877375e491e, []int{2}
}
func (m *RemoteStopTransactionReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoteStopTransactionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoteStopTransactionReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoteStopTransactionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteStopTransactionReq.Merge(m, src)
}
func (m *RemoteStopTransactionReq) XXX_Size() int {
	return m.Size()
}
func (m *RemoteStopTransactionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteStopTransactionReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteStopTransactionReq proto.InternalMessageInfo

type RemoteStopTransactionResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteStopTransactionResp) Reset()         { *m = RemoteStopTransactionResp{} }
func (m *RemoteStopTransactionResp) String() string { return proto.CompactTextString(m) }
func (*RemoteStopTransactionResp) ProtoMessage()    {}
func (*RemoteStopTransactionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a5877375e491e, []int{3}
}
func (m *RemoteStopTransactionResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoteStopTransactionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoteStopTransactionResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoteStopTransactionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteStopTransactionResp.Merge(m, src)
}
func (m *RemoteStopTransactionResp) XXX_Size() int {
	return m.Size()
}
func (m *RemoteStopTransactionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteStopTransactionResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteStopTransactionResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RemoteStartTransactionReq)(nil), "ocpp.RemoteStartTransactionReq")
	proto.RegisterType((*RemoteStartTransactionResp)(nil), "ocpp.RemoteStartTransactionResp")
	proto.RegisterType((*RemoteStopTransactionReq)(nil), "ocpp.RemoteStopTransactionReq")
	proto.RegisterType((*RemoteStopTransactionResp)(nil), "ocpp.RemoteStopTransactionResp")
}

func init() { proto.RegisterFile("operations.proto", fileDescriptor_1b4a5877375e491e) }

var fileDescriptor_1b4a5877375e491e = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x2f, 0x48, 0x2d,
	0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4f,
	0x2e, 0x28, 0x90, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1,
	0xa4, 0x14, 0xcc, 0x25, 0x19, 0x94, 0x9a, 0x9b, 0x5f, 0x92, 0x1a, 0x5c, 0x92, 0x58, 0x54, 0x12,
	0x52, 0x94, 0x98, 0x57, 0x9c, 0x98, 0x0c, 0x32, 0x35, 0x28, 0xb5, 0x50, 0x48, 0x98, 0x8b, 0x35,
	0xb9, 0x20, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x25, 0xb9, 0xc0, 0x33,
	0x45, 0x48, 0x91, 0x8b, 0x27, 0x39, 0x3f, 0x2f, 0x2f, 0x35, 0xb9, 0x24, 0xbf, 0x08, 0x24, 0xc7,
	0xa4, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x0d, 0x17, 0xf3, 0x4c, 0x51, 0x72, 0xe6, 0x92, 0xc2, 0x65,
	0x68, 0x71, 0x81, 0x90, 0x2a, 0x17, 0x5f, 0x09, 0x42, 0x08, 0x66, 0x3c, 0x6b, 0x10, 0x2f, 0x92,
	0xa8, 0x67, 0x8a, 0x52, 0x18, 0x97, 0x04, 0xcc, 0x90, 0xfc, 0x02, 0x62, 0x1c, 0x86, 0x69, 0x2e,
	0x13, 0x36, 0x73, 0xa5, 0x11, 0x3e, 0x46, 0x33, 0xb7, 0xb8, 0xc0, 0x49, 0xf2, 0xc4, 0x43, 0x39,
	0x86, 0x17, 0x0f, 0xe5, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23,
	0x39, 0xc6, 0x28, 0xe6, 0xa2, 0x82, 0xe4, 0x24, 0x36, 0x70, 0x80, 0x19, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x31, 0x78, 0x95, 0x39, 0x79, 0x01, 0x00, 0x00,
}

type RemoteStartTransactionReqFace interface {
	Proto() github_com_gogo_protobuf_proto.Message
	GetCpId() string
	GetConnectorId() int32
}

func (this *RemoteStartTransactionReq) Proto() github_com_gogo_protobuf_proto.Message {
	return this
}

func (this *RemoteStartTransactionReq) TestProto() github_com_gogo_protobuf_proto.Message {
	return NewRemoteStartTransactionReqFromFace(this)
}

func (this *RemoteStartTransactionReq) GetCpId() string {
	return this.CpId
}

func (this *RemoteStartTransactionReq) GetConnectorId() int32 {
	return this.ConnectorId
}

func NewRemoteStartTransactionReqFromFace(that RemoteStartTransactionReqFace) *RemoteStartTransactionReq {
	this := &RemoteStartTransactionReq{}
	this.CpId = that.GetCpId()
	this.ConnectorId = that.GetConnectorId()
	return this
}

type RemoteStartTransactionRespFace interface {
	Proto() github_com_gogo_protobuf_proto.Message
	GetTransactionId() int32
}

func (this *RemoteStartTransactionResp) Proto() github_com_gogo_protobuf_proto.Message {
	return this
}

func (this *RemoteStartTransactionResp) TestProto() github_com_gogo_protobuf_proto.Message {
	return NewRemoteStartTransactionRespFromFace(this)
}

func (this *RemoteStartTransactionResp) GetTransactionId() int32 {
	return this.TransactionId
}

func NewRemoteStartTransactionRespFromFace(that RemoteStartTransactionRespFace) *RemoteStartTransactionResp {
	this := &RemoteStartTransactionResp{}
	this.TransactionId = that.GetTransactionId()
	return this
}

type RemoteStopTransactionReqFace interface {
	Proto() github_com_gogo_protobuf_proto.Message
	GetCpId() string
	GetTransactionId() int32
}

func (this *RemoteStopTransactionReq) Proto() github_com_gogo_protobuf_proto.Message {
	return this
}

func (this *RemoteStopTransactionReq) TestProto() github_com_gogo_protobuf_proto.Message {
	return NewRemoteStopTransactionReqFromFace(this)
}

func (this *RemoteStopTransactionReq) GetCpId() string {
	return this.CpId
}

func (this *RemoteStopTransactionReq) GetTransactionId() int32 {
	return this.TransactionId
}

func NewRemoteStopTransactionReqFromFace(that RemoteStopTransactionReqFace) *RemoteStopTransactionReq {
	this := &RemoteStopTransactionReq{}
	this.CpId = that.GetCpId()
	this.TransactionId = that.GetTransactionId()
	return this
}

type RemoteStopTransactionRespFace interface {
	Proto() github_com_gogo_protobuf_proto.Message
}

func (this *RemoteStopTransactionResp) Proto() github_com_gogo_protobuf_proto.Message {
	return this
}

func (this *RemoteStopTransactionResp) TestProto() github_com_gogo_protobuf_proto.Message {
	return NewRemoteStopTransactionRespFromFace(this)
}

func NewRemoteStopTransactionRespFromFace(that RemoteStopTransactionRespFace) *RemoteStopTransactionResp {
	this := &RemoteStopTransactionResp{}
	return this
}

func (m *RemoteStartTransactionReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteStartTransactionReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoteStartTransactionReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ConnectorId != 0 {
		i = encodeVarintOperations(dAtA, i, uint64(m.ConnectorId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CpId) > 0 {
		i -= len(m.CpId)
		copy(dAtA[i:], m.CpId)
		i = encodeVarintOperations(dAtA, i, uint64(len(m.CpId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RemoteStartTransactionResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteStartTransactionResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoteStartTransactionResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TransactionId != 0 {
		i = encodeVarintOperations(dAtA, i, uint64(m.TransactionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RemoteStopTransactionReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteStopTransactionReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoteStopTransactionReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TransactionId != 0 {
		i = encodeVarintOperations(dAtA, i, uint64(m.TransactionId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CpId) > 0 {
		i -= len(m.CpId)
		copy(dAtA[i:], m.CpId)
		i = encodeVarintOperations(dAtA, i, uint64(len(m.CpId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RemoteStopTransactionResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteStopTransactionResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoteStopTransactionResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintOperations(dAtA []byte, offset int, v uint64) int {
	offset -= sovOperations(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RemoteStartTransactionReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CpId)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.ConnectorId != 0 {
		n += 1 + sovOperations(uint64(m.ConnectorId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RemoteStartTransactionResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TransactionId != 0 {
		n += 1 + sovOperations(uint64(m.TransactionId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RemoteStopTransactionReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CpId)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.TransactionId != 0 {
		n += 1 + sovOperations(uint64(m.TransactionId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RemoteStopTransactionResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOperations(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOperations(x uint64) (n int) {
	return sovOperations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RemoteStartTransactionReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: RemoteStartTransactionReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteStartTransactionReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CpId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CpId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectorId", wireType)
			}
			m.ConnectorId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnectorId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOperations
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
func (m *RemoteStartTransactionResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: RemoteStartTransactionResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteStartTransactionResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionId", wireType)
			}
			m.TransactionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOperations
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
func (m *RemoteStopTransactionReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: RemoteStopTransactionReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteStopTransactionReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CpId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CpId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionId", wireType)
			}
			m.TransactionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOperations
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
func (m *RemoteStopTransactionResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: RemoteStopTransactionResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteStopTransactionResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOperations
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
func skipOperations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOperations
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
					return 0, ErrIntOverflowOperations
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
					return 0, ErrIntOverflowOperations
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
				return 0, ErrInvalidLengthOperations
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOperations
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOperations
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOperations        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOperations          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOperations = fmt.Errorf("proto: unexpected end of group")
)
