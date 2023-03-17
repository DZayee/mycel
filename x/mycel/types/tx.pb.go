// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/mycel/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgRegisterDomain struct {
	Creator                  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Parent                   string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	RegistrationPeriodInYear uint64 `protobuf:"varint,4,opt,name=registrationPeriodInYear,proto3" json:"registrationPeriodInYear,omitempty"`
}

func (m *MsgRegisterDomain) Reset()         { *m = MsgRegisterDomain{} }
func (m *MsgRegisterDomain) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterDomain) ProtoMessage()    {}
func (*MsgRegisterDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddc129799a3b0dc2, []int{0}
}
func (m *MsgRegisterDomain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterDomain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterDomain.Merge(m, src)
}
func (m *MsgRegisterDomain) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterDomain.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterDomain proto.InternalMessageInfo

func (m *MsgRegisterDomain) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRegisterDomain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgRegisterDomain) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *MsgRegisterDomain) GetRegistrationPeriodInYear() uint64 {
	if m != nil {
		return m.RegistrationPeriodInYear
	}
	return 0
}

type MsgRegisterDomainResponse struct {
}

func (m *MsgRegisterDomainResponse) Reset()         { *m = MsgRegisterDomainResponse{} }
func (m *MsgRegisterDomainResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterDomainResponse) ProtoMessage()    {}
func (*MsgRegisterDomainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddc129799a3b0dc2, []int{1}
}
func (m *MsgRegisterDomainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterDomainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterDomainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterDomainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterDomainResponse.Merge(m, src)
}
func (m *MsgRegisterDomainResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterDomainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterDomainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterDomainResponse proto.InternalMessageInfo

type MsgUpdateWalletRecord struct {
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Parent           string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	WalletRecordType string `protobuf:"bytes,4,opt,name=walletRecordType,proto3" json:"walletRecordType,omitempty"`
}

func (m *MsgUpdateWalletRecord) Reset()         { *m = MsgUpdateWalletRecord{} }
func (m *MsgUpdateWalletRecord) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateWalletRecord) ProtoMessage()    {}
func (*MsgUpdateWalletRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddc129799a3b0dc2, []int{2}
}
func (m *MsgUpdateWalletRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateWalletRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateWalletRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateWalletRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateWalletRecord.Merge(m, src)
}
func (m *MsgUpdateWalletRecord) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateWalletRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateWalletRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateWalletRecord proto.InternalMessageInfo

func (m *MsgUpdateWalletRecord) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateWalletRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgUpdateWalletRecord) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *MsgUpdateWalletRecord) GetWalletRecordType() string {
	if m != nil {
		return m.WalletRecordType
	}
	return ""
}

type MsgUpdateWalletRecordResponse struct {
}

func (m *MsgUpdateWalletRecordResponse) Reset()         { *m = MsgUpdateWalletRecordResponse{} }
func (m *MsgUpdateWalletRecordResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateWalletRecordResponse) ProtoMessage()    {}
func (*MsgUpdateWalletRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddc129799a3b0dc2, []int{3}
}
func (m *MsgUpdateWalletRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateWalletRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateWalletRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateWalletRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateWalletRecordResponse.Merge(m, src)
}
func (m *MsgUpdateWalletRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateWalletRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateWalletRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateWalletRecordResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterDomain)(nil), "mycel.mycel.MsgRegisterDomain")
	proto.RegisterType((*MsgRegisterDomainResponse)(nil), "mycel.mycel.MsgRegisterDomainResponse")
	proto.RegisterType((*MsgUpdateWalletRecord)(nil), "mycel.mycel.MsgUpdateWalletRecord")
	proto.RegisterType((*MsgUpdateWalletRecordResponse)(nil), "mycel.mycel.MsgUpdateWalletRecordResponse")
}

func init() { proto.RegisterFile("mycel/mycel/tx.proto", fileDescriptor_ddc129799a3b0dc2) }

var fileDescriptor_ddc129799a3b0dc2 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xad, 0x4c, 0x4e,
	0xcd, 0xd1, 0x87, 0x90, 0x25, 0x15, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xdc, 0x60, 0xbe,
	0x1e, 0x98, 0x54, 0x9a, 0xca, 0xc8, 0x25, 0xe8, 0x5b, 0x9c, 0x1e, 0x94, 0x9a, 0x9e, 0x59, 0x5c,
	0x92, 0x5a, 0xe4, 0x92, 0x9f, 0x9b, 0x98, 0x99, 0x27, 0x24, 0xc1, 0xc5, 0x9e, 0x5c, 0x94, 0x9a,
	0x58, 0x92, 0x5f, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0x09, 0x71, 0xb1,
	0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x85, 0xc1, 0x6c, 0x21, 0x31, 0x2e, 0xb6, 0x82, 0xc4,
	0xa2, 0xd4, 0xbc, 0x12, 0x09, 0x66, 0xb0, 0x28, 0x94, 0x27, 0x64, 0xc5, 0x25, 0x51, 0x04, 0x36,
	0xb7, 0x28, 0xb1, 0x24, 0x33, 0x3f, 0x2f, 0x20, 0xb5, 0x28, 0x33, 0x3f, 0xc5, 0x33, 0x2f, 0x32,
	0x35, 0xb1, 0x48, 0x82, 0x45, 0x81, 0x51, 0x83, 0x25, 0x08, 0xa7, 0xbc, 0x92, 0x34, 0x97, 0x24,
	0x86, 0xb3, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0x3a, 0x19, 0xb9, 0x44, 0x7d,
	0x8b, 0xd3, 0x43, 0x0b, 0x52, 0x12, 0x4b, 0x52, 0xc3, 0x13, 0x73, 0x72, 0x52, 0x4b, 0x82, 0x52,
	0x93, 0xf3, 0x8b, 0x52, 0xa8, 0xe4, 0x70, 0x2d, 0x2e, 0x81, 0x72, 0x24, 0x53, 0x43, 0x2a, 0x0b,
	0x52, 0xc1, 0x0e, 0xe6, 0x0c, 0xc2, 0x10, 0x57, 0x92, 0xe7, 0x92, 0xc5, 0xea, 0x14, 0x98, 0x63,
	0x8d, 0x8e, 0x32, 0x72, 0x31, 0xfb, 0x16, 0xa7, 0x0b, 0x45, 0x70, 0xf1, 0xa1, 0x85, 0xb2, 0x9c,
	0x1e, 0x52, 0x4c, 0xe8, 0x61, 0x78, 0x57, 0x4a, 0x0d, 0xbf, 0x3c, 0xcc, 0x06, 0xa1, 0x14, 0x2e,
	0x21, 0x2c, 0x41, 0xa1, 0x84, 0xae, 0x1b, 0x53, 0x8d, 0x94, 0x16, 0x61, 0x35, 0x30, 0x5b, 0x9c,
	0x74, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x18, 0x92, 0xc0, 0x2a,
	0x60, 0x09, 0xad, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x9c, 0xd8, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xa7, 0x94, 0xe5, 0xa3, 0x84, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RegisterDomain(ctx context.Context, in *MsgRegisterDomain, opts ...grpc.CallOption) (*MsgRegisterDomainResponse, error)
	UpdateWalletRecord(ctx context.Context, in *MsgUpdateWalletRecord, opts ...grpc.CallOption) (*MsgUpdateWalletRecordResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterDomain(ctx context.Context, in *MsgRegisterDomain, opts ...grpc.CallOption) (*MsgRegisterDomainResponse, error) {
	out := new(MsgRegisterDomainResponse)
	err := c.cc.Invoke(ctx, "/mycel.mycel.Msg/RegisterDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateWalletRecord(ctx context.Context, in *MsgUpdateWalletRecord, opts ...grpc.CallOption) (*MsgUpdateWalletRecordResponse, error) {
	out := new(MsgUpdateWalletRecordResponse)
	err := c.cc.Invoke(ctx, "/mycel.mycel.Msg/UpdateWalletRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RegisterDomain(context.Context, *MsgRegisterDomain) (*MsgRegisterDomainResponse, error)
	UpdateWalletRecord(context.Context, *MsgUpdateWalletRecord) (*MsgUpdateWalletRecordResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterDomain(ctx context.Context, req *MsgRegisterDomain) (*MsgRegisterDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDomain not implemented")
}
func (*UnimplementedMsgServer) UpdateWalletRecord(ctx context.Context, req *MsgUpdateWalletRecord) (*MsgUpdateWalletRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWalletRecord not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterDomain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mycel.mycel.Msg/RegisterDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterDomain(ctx, req.(*MsgRegisterDomain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateWalletRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateWalletRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateWalletRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mycel.mycel.Msg/UpdateWalletRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateWalletRecord(ctx, req.(*MsgUpdateWalletRecord))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mycel.mycel.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDomain",
			Handler:    _Msg_RegisterDomain_Handler,
		},
		{
			MethodName: "UpdateWalletRecord",
			Handler:    _Msg_UpdateWalletRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mycel/mycel/tx.proto",
}

func (m *MsgRegisterDomain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterDomain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterDomain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RegistrationPeriodInYear != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.RegistrationPeriodInYear))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Parent) > 0 {
		i -= len(m.Parent)
		copy(dAtA[i:], m.Parent)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Parent)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterDomainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterDomainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterDomainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateWalletRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateWalletRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateWalletRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WalletRecordType) > 0 {
		i -= len(m.WalletRecordType)
		copy(dAtA[i:], m.WalletRecordType)
		i = encodeVarintTx(dAtA, i, uint64(len(m.WalletRecordType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Parent) > 0 {
		i -= len(m.Parent)
		copy(dAtA[i:], m.Parent)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Parent)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateWalletRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateWalletRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateWalletRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterDomain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Parent)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.RegistrationPeriodInYear != 0 {
		n += 1 + sovTx(uint64(m.RegistrationPeriodInYear))
	}
	return n
}

func (m *MsgRegisterDomainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateWalletRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Parent)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.WalletRecordType)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateWalletRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterDomain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterDomain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterDomain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistrationPeriodInYear", wireType)
			}
			m.RegistrationPeriodInYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegistrationPeriodInYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegisterDomainResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterDomainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterDomainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateWalletRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateWalletRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateWalletRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletRecordType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WalletRecordType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateWalletRecordResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateWalletRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateWalletRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
