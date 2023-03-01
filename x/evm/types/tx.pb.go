// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargazer/evm/v1alpha1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// EthTransactionRequest encapsulates an Ethereum transaction as an SDK message.
type EthTransactionRequest struct {
	// data is inner transaction data of the Ethereum transaction
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// from is the ethereum signer address in hex format. This address value is checked
	// against the address derived from the signature (V, R, S) using the
	// secp256k1 elliptic curve
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
}

func (m *EthTransactionRequest) Reset()         { *m = EthTransactionRequest{} }
func (m *EthTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*EthTransactionRequest) ProtoMessage()    {}
func (*EthTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72f89b557d3ac34e, []int{0}
}
func (m *EthTransactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthTransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthTransactionRequest.Merge(m, src)
}
func (m *EthTransactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *EthTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EthTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EthTransactionRequest proto.InternalMessageInfo

// EthTransactionResponse defines the Msg/EthereumTx response type.
type EthTransactionResponse struct {
	// `receipt` represents the json-encoded bytes of the ethereum receipt.
	Receipt []byte `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
}

func (m *EthTransactionResponse) Reset()         { *m = EthTransactionResponse{} }
func (m *EthTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*EthTransactionResponse) ProtoMessage()    {}
func (*EthTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72f89b557d3ac34e, []int{1}
}
func (m *EthTransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthTransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthTransactionResponse.Merge(m, src)
}
func (m *EthTransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *EthTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EthTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EthTransactionResponse proto.InternalMessageInfo

// ExtensionOptionsEthTransaction is an extension option for ethereum transactions
type ExtensionOptionsEthTransaction struct {
}

func (m *ExtensionOptionsEthTransaction) Reset()         { *m = ExtensionOptionsEthTransaction{} }
func (m *ExtensionOptionsEthTransaction) String() string { return proto.CompactTextString(m) }
func (*ExtensionOptionsEthTransaction) ProtoMessage()    {}
func (*ExtensionOptionsEthTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_72f89b557d3ac34e, []int{2}
}
func (m *ExtensionOptionsEthTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtensionOptionsEthTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtensionOptionsEthTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtensionOptionsEthTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionOptionsEthTransaction.Merge(m, src)
}
func (m *ExtensionOptionsEthTransaction) XXX_Size() int {
	return m.Size()
}
func (m *ExtensionOptionsEthTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionOptionsEthTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionOptionsEthTransaction proto.InternalMessageInfo

// `UpdateParamsRequest` is the Msg/UpdateParams request type.
//
// Since: cosmos-sdk 0.47
type UpdateParamsRequest struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the x/evm parameters to update.
	//
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *UpdateParamsRequest) Reset()         { *m = UpdateParamsRequest{} }
func (m *UpdateParamsRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateParamsRequest) ProtoMessage()    {}
func (*UpdateParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72f89b557d3ac34e, []int{3}
}
func (m *UpdateParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateParamsRequest.Merge(m, src)
}
func (m *UpdateParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpdateParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateParamsRequest proto.InternalMessageInfo

func (m *UpdateParamsRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *UpdateParamsRequest) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// `UpdateParamsResponse` defines the response structure for executing a
// UpdateParamsResponse message.
//
// Since: cosmos-sdk 0.47
type UpdateParamsResponse struct {
}

func (m *UpdateParamsResponse) Reset()         { *m = UpdateParamsResponse{} }
func (m *UpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateParamsResponse) ProtoMessage()    {}
func (*UpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72f89b557d3ac34e, []int{4}
}
func (m *UpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateParamsResponse.Merge(m, src)
}
func (m *UpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EthTransactionRequest)(nil), "stargazer.evm.v1alpha1.EthTransactionRequest")
	proto.RegisterType((*EthTransactionResponse)(nil), "stargazer.evm.v1alpha1.EthTransactionResponse")
	proto.RegisterType((*ExtensionOptionsEthTransaction)(nil), "stargazer.evm.v1alpha1.ExtensionOptionsEthTransaction")
	proto.RegisterType((*UpdateParamsRequest)(nil), "stargazer.evm.v1alpha1.UpdateParamsRequest")
	proto.RegisterType((*UpdateParamsResponse)(nil), "stargazer.evm.v1alpha1.UpdateParamsResponse")
}

func init() { proto.RegisterFile("stargazer/evm/v1alpha1/tx.proto", fileDescriptor_72f89b557d3ac34e) }

var fileDescriptor_72f89b557d3ac34e = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3f, 0x6f, 0x13, 0x31,
	0x14, 0x3f, 0xa3, 0x50, 0x14, 0x53, 0x75, 0x38, 0x42, 0x08, 0x37, 0x38, 0x55, 0x90, 0xa0, 0x2a,
	0xd4, 0xa7, 0x14, 0x09, 0xa1, 0x8a, 0xa5, 0x91, 0x32, 0x22, 0xd0, 0x15, 0x16, 0x16, 0xe4, 0xde,
	0x3d, 0x2e, 0x16, 0xdc, 0xd9, 0xd8, 0xee, 0x29, 0x65, 0x42, 0x4c, 0x8c, 0xec, 0x2c, 0x7c, 0x84,
	0x0e, 0x7c, 0x88, 0x8e, 0x15, 0x13, 0x13, 0x42, 0xc9, 0xd0, 0x91, 0xaf, 0x80, 0xce, 0x76, 0x08,
	0xa9, 0x12, 0xa9, 0x93, 0xdf, 0x7b, 0xfe, 0xfd, 0xde, 0xfb, 0xbd, 0x3f, 0xb8, 0xab, 0x0d, 0x53,
	0x39, 0xfb, 0x00, 0x2a, 0x86, 0xaa, 0x88, 0xab, 0x3e, 0x7b, 0x27, 0x47, 0xac, 0x1f, 0x9b, 0x31,
	0x95, 0x4a, 0x18, 0x11, 0xb6, 0xff, 0x01, 0x28, 0x54, 0x05, 0x9d, 0x01, 0xa2, 0x5b, 0xa9, 0xd0,
	0x85, 0xd0, 0x71, 0xa1, 0xf3, 0xb8, 0xea, 0xd7, 0x8f, 0x23, 0x44, 0xb7, 0xdd, 0xc7, 0x6b, 0xeb,
	0xc5, 0xce, 0xf1, 0x5f, 0xad, 0x5c, 0xe4, 0xc2, 0xc5, 0x6b, 0xcb, 0x47, 0xef, 0xac, 0x90, 0x20,
	0x99, 0x62, 0x85, 0xa7, 0xf6, 0x86, 0xf8, 0xe6, 0xd0, 0x8c, 0x5e, 0x28, 0x56, 0x6a, 0x96, 0x1a,
	0x2e, 0xca, 0x04, 0xde, 0x1f, 0x81, 0x36, 0x61, 0x88, 0x1b, 0x19, 0x33, 0xac, 0x83, 0x36, 0xd1,
	0xd6, 0x7a, 0x62, 0xed, 0x3a, 0xf6, 0x46, 0x89, 0xa2, 0x73, 0x65, 0x13, 0x6d, 0x35, 0x13, 0x6b,
	0xef, 0x35, 0x3e, 0x7f, 0xeb, 0x06, 0xbd, 0xc7, 0xb8, 0x7d, 0x31, 0x8d, 0x96, 0xa2, 0xd4, 0x10,
	0x76, 0xf0, 0x35, 0x05, 0x29, 0x70, 0x69, 0x7c, 0xaa, 0x99, 0xeb, 0x99, 0x77, 0x31, 0x19, 0x8e,
	0x0d, 0x94, 0x9a, 0x8b, 0xf2, 0x99, 0xac, 0xa9, 0x7a, 0x31, 0x93, 0xc7, 0x7d, 0x45, 0xf8, 0xc6,
	0x4b, 0x99, 0x31, 0x03, 0xcf, 0xad, 0xfe, 0x99, 0xce, 0x47, 0xb8, 0xc9, 0x8e, 0xcc, 0x48, 0x28,
	0x6e, 0x8e, 0x6d, 0x85, 0xe6, 0xa0, 0xf3, 0xe3, 0xfb, 0x4e, 0xcb, 0x0f, 0x68, 0x3f, 0xcb, 0x14,
	0x68, 0x7d, 0x60, 0x14, 0x2f, 0xf3, 0x64, 0x0e, 0x0d, 0x9f, 0xe0, 0x35, 0x37, 0x08, 0xdb, 0xcd,
	0xf5, 0x5d, 0x42, 0x97, 0x2f, 0x84, 0xba, 0x72, 0x83, 0xc6, 0xe9, 0xaf, 0x6e, 0x90, 0x78, 0xce,
	0xde, 0xc6, 0xa7, 0xf3, 0x93, 0xed, 0x79, 0xb6, 0x5e, 0x1b, 0xb7, 0x16, 0xc5, 0xb9, 0xee, 0x77,
	0xff, 0x20, 0x8c, 0x9f, 0xea, 0xfc, 0x00, 0x54, 0xc5, 0x53, 0x08, 0x05, 0xde, 0x58, 0x6c, 0x2e,
	0xdc, 0x59, 0x55, 0x76, 0xe9, 0x56, 0x22, 0x7a, 0x59, 0xb8, 0x9f, 0x3e, 0xc7, 0xeb, 0xff, 0xeb,
	0x0a, 0xef, 0xaf, 0xe2, 0x2f, 0x19, 0x6d, 0xf4, 0xe0, 0x72, 0x60, 0x57, 0x2a, 0xba, 0xfa, 0xf1,
	0xfc, 0x64, 0x1b, 0x0d, 0xf6, 0x4f, 0x27, 0x04, 0x9d, 0x4d, 0x08, 0xfa, 0x3d, 0x21, 0xe8, 0xcb,
	0x94, 0x04, 0x67, 0x53, 0x12, 0xfc, 0x9c, 0x92, 0xe0, 0xd5, 0x3d, 0xf9, 0x36, 0xa7, 0x87, 0xa0,
	0x58, 0x3a, 0x62, 0xbc, 0xa4, 0x19, 0x54, 0xf1, 0xfc, 0x42, 0xc7, 0xf6, 0x46, 0xcd, 0xb1, 0x04,
	0x7d, 0xb8, 0x66, 0x4f, 0xf3, 0xe1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0xb7, 0xac, 0xb4,
	0x44, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	// EthTransaction defines a method submitting Ethereum transactions.
	EthTransaction(ctx context.Context, in *EthTransactionRequest, opts ...grpc.CallOption) (*EthTransactionResponse, error)
	// `UpdateParams` defines a governance operation for updating the x/evm module
	// parameters. The authority is defaults to the x/gov module account.
	//
	// Since: cosmos-sdk 0.47
	UpdateParams(ctx context.Context, in *UpdateParamsRequest, opts ...grpc.CallOption) (*UpdateParamsResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) EthTransaction(ctx context.Context, in *EthTransactionRequest, opts ...grpc.CallOption) (*EthTransactionResponse, error) {
	out := new(EthTransactionResponse)
	err := c.cc.Invoke(ctx, "/stargazer.evm.v1alpha1.MsgService/EthTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) UpdateParams(ctx context.Context, in *UpdateParamsRequest, opts ...grpc.CallOption) (*UpdateParamsResponse, error) {
	out := new(UpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/stargazer.evm.v1alpha1.MsgService/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	// EthTransaction defines a method submitting Ethereum transactions.
	EthTransaction(context.Context, *EthTransactionRequest) (*EthTransactionResponse, error)
	// `UpdateParams` defines a governance operation for updating the x/evm module
	// parameters. The authority is defaults to the x/gov module account.
	//
	// Since: cosmos-sdk 0.47
	UpdateParams(context.Context, *UpdateParamsRequest) (*UpdateParamsResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) EthTransaction(ctx context.Context, req *EthTransactionRequest) (*EthTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthTransaction not implemented")
}
func (*UnimplementedMsgServiceServer) UpdateParams(ctx context.Context, req *UpdateParamsRequest) (*UpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_EthTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EthTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).EthTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stargazer.evm.v1alpha1.MsgService/EthTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).EthTransaction(ctx, req.(*EthTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stargazer.evm.v1alpha1.MsgService/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).UpdateParams(ctx, req.(*UpdateParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stargazer.evm.v1alpha1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EthTransaction",
			Handler:    _MsgService_EthTransaction_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _MsgService_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stargazer/evm/v1alpha1/tx.proto",
}

func (m *EthTransactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthTransactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthTransactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EthTransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthTransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthTransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receipt) > 0 {
		i -= len(m.Receipt)
		copy(dAtA[i:], m.Receipt)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receipt)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExtensionOptionsEthTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtensionOptionsEthTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtensionOptionsEthTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *UpdateParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *EthTransactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *EthTransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Receipt)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *ExtensionOptionsEthTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UpdateParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *UpdateParamsResponse) Size() (n int) {
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
func (m *EthTransactionRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EthTransactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthTransactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
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
func (m *EthTransactionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EthTransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthTransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receipt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receipt = append(m.Receipt[:0], dAtA[iNdEx:postIndex]...)
			if m.Receipt == nil {
				m.Receipt = []byte{}
			}
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
func (m *ExtensionOptionsEthTransaction) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ExtensionOptionsEthTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtensionOptionsEthTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *UpdateParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UpdateParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *UpdateParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
