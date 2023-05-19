// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: polaris/evm/v1alpha1/tx.proto

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
}

func (m *EthTransactionRequest) Reset()         { *m = EthTransactionRequest{} }
func (m *EthTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*EthTransactionRequest) ProtoMessage()    {}
func (*EthTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b33d2a2c64400f, []int{0}
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
	// `gas_used` represents the gas used by the virtual machine execution.
	GasUsed uint64 `protobuf:"varint,1,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	// `vm_error` contains an error message if the virtual machine execution failed.
	VmError string `protobuf:"bytes,2,opt,name=vm_error,json=vmError,proto3" json:"vm_error,omitempty"`
	// `return_data` contains the return data of the virtual machine execution.
	ReturnData []byte `protobuf:"bytes,3,opt,name=return_data,json=returnData,proto3" json:"return_data,omitempty"`
}

func (m *EthTransactionResponse) Reset()         { *m = EthTransactionResponse{} }
func (m *EthTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*EthTransactionResponse) ProtoMessage()    {}
func (*EthTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b33d2a2c64400f, []int{1}
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
	return fileDescriptor_d8b33d2a2c64400f, []int{2}
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
	return fileDescriptor_d8b33d2a2c64400f, []int{3}
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
	proto.RegisterType((*EthTransactionRequest)(nil), "polaris.evm.v1alpha1.EthTransactionRequest")
	proto.RegisterType((*EthTransactionResponse)(nil), "polaris.evm.v1alpha1.EthTransactionResponse")
	proto.RegisterType((*UpdateParamsRequest)(nil), "polaris.evm.v1alpha1.UpdateParamsRequest")
	proto.RegisterType((*UpdateParamsResponse)(nil), "polaris.evm.v1alpha1.UpdateParamsResponse")
}

func init() { proto.RegisterFile("polaris/evm/v1alpha1/tx.proto", fileDescriptor_d8b33d2a2c64400f) }

var fileDescriptor_d8b33d2a2c64400f = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0x41, 0x68, 0xc9, 0xb5, 0xea, 0x70, 0x84, 0x92, 0x5a, 0xe0, 0x94, 0x4c, 0x25, 0x14,
	0x5b, 0x29, 0x12, 0x43, 0x36, 0x22, 0xca, 0x86, 0x84, 0x5c, 0xba, 0xb0, 0x44, 0xd7, 0xf8, 0xc9,
	0xb1, 0x5a, 0xfb, 0xcc, 0xbd, 0xb3, 0xd5, 0x6e, 0x88, 0x89, 0x91, 0x95, 0x8d, 0x9f, 0xd0, 0x81,
	0x1f, 0xd1, 0xb1, 0x62, 0x62, 0x42, 0x28, 0x19, 0xf2, 0x37, 0x90, 0xef, 0x2e, 0x0a, 0xad, 0x3c,
	0x64, 0xf2, 0x7b, 0xef, 0x7b, 0x9f, 0xbf, 0x77, 0xdf, 0x7b, 0xf4, 0x49, 0x2e, 0xce, 0xb8, 0x4c,
	0x30, 0x80, 0x32, 0x0d, 0xca, 0x3e, 0x3f, 0xcb, 0x27, 0xbc, 0x1f, 0xa8, 0x73, 0x3f, 0x97, 0x42,
	0x09, 0xd6, 0xb2, 0xb0, 0x0f, 0x65, 0xea, 0x2f, 0x60, 0xf7, 0xd1, 0x58, 0x60, 0x2a, 0x30, 0x48,
	0x31, 0x0e, 0xca, 0x7e, 0xf5, 0x31, 0xed, 0xee, 0x8e, 0x01, 0x46, 0x3a, 0x0b, 0x4c, 0x62, 0xa1,
	0x56, 0x2c, 0x62, 0x61, 0xea, 0x55, 0x64, 0xab, 0x4f, 0x6b, 0xe5, 0x73, 0x2e, 0x79, 0x6a, 0x89,
	0xdd, 0x3e, 0x7d, 0x78, 0xa8, 0x26, 0x1f, 0x24, 0xcf, 0x90, 0x8f, 0x55, 0x22, 0xb2, 0x10, 0x3e,
	0x15, 0x80, 0x8a, 0x31, 0xda, 0x88, 0xb8, 0xe2, 0x6d, 0xb2, 0x4b, 0xf6, 0x36, 0x43, 0x1d, 0x0f,
	0x1a, 0x5f, 0x7f, 0x74, 0x9c, 0x6e, 0x41, 0xb7, 0x6f, 0x53, 0x30, 0x17, 0x19, 0x02, 0xdb, 0xa1,
	0xf7, 0x63, 0x8e, 0xa3, 0x02, 0x21, 0xd2, 0xbc, 0x46, 0xb8, 0x1e, 0x73, 0x3c, 0x46, 0x88, 0x2a,
	0xa8, 0x4c, 0x47, 0x20, 0xa5, 0x90, 0xed, 0x3b, 0xbb, 0x64, 0xaf, 0x19, 0xae, 0x97, 0xe9, 0x61,
	0x95, 0xb2, 0x0e, 0xdd, 0x90, 0xa0, 0x0a, 0x99, 0x8d, 0xb4, 0xe0, 0x5d, 0x2d, 0x48, 0x4d, 0xe9,
	0xcd, 0x52, 0xf6, 0x3b, 0xa1, 0x0f, 0x8e, 0xf3, 0x88, 0x2b, 0x78, 0xaf, 0x1f, 0xb0, 0x18, 0xf4,
	0x15, 0x6d, 0xf2, 0x42, 0x4d, 0x84, 0x4c, 0xd4, 0x85, 0x56, 0x6d, 0x0e, 0xdb, 0xbf, 0x7e, 0xbe,
	0x68, 0x59, 0x7f, 0x5e, 0x47, 0x91, 0x04, 0xc4, 0x23, 0x25, 0x93, 0x2c, 0x0e, 0x97, 0xad, 0x6c,
	0x40, 0xd7, 0x8c, 0x13, 0x7a, 0x9e, 0x8d, 0x83, 0xc7, 0x7e, 0xdd, 0x36, 0x7c, 0x23, 0x36, 0x6c,
	0x5c, 0xfd, 0xe9, 0x38, 0xa1, 0x65, 0x0c, 0xb6, 0xbe, 0xcc, 0x2f, 0x7b, 0xcb, 0x7f, 0x75, 0xb7,
	0x69, 0xeb, 0xe6, 0x68, 0xc6, 0x90, 0x83, 0x39, 0xa1, 0xf4, 0x1d, 0xc6, 0x47, 0x20, 0xcb, 0x64,
	0x0c, 0xec, 0x94, 0x6e, 0xdd, 0x74, 0x8e, 0x3d, 0xaf, 0x17, 0xad, 0x5d, 0x89, 0xbb, 0xbf, 0x5a,
	0xb3, 0x5d, 0x06, 0xd0, 0xcd, 0xff, 0x67, 0x62, 0xcf, 0xea, 0xd9, 0x35, 0x96, 0xba, 0xbd, 0x55,
	0x5a, 0x8d, 0x8c, 0x7b, 0xef, 0xf3, 0xfc, 0xb2, 0x47, 0x86, 0x6f, 0xaf, 0xa6, 0x1e, 0xb9, 0x9e,
	0x7a, 0xe4, 0xef, 0xd4, 0x23, 0xdf, 0x66, 0x9e, 0x73, 0x3d, 0xf3, 0x9c, 0xdf, 0x33, 0xcf, 0xf9,
	0xb8, 0x9f, 0x9f, 0xc6, 0xfe, 0x09, 0x48, 0x3e, 0x9e, 0xf0, 0x24, 0xf3, 0x23, 0x28, 0x83, 0xc5,
	0x59, 0xda, 0x43, 0x3f, 0xd7, 0xf7, 0xa9, 0x2e, 0x72, 0xc0, 0x93, 0x35, 0x7d, 0x96, 0x2f, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x09, 0x4d, 0x5f, 0x76, 0x3a, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/polaris.evm.v1alpha1.MsgService/EthTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) UpdateParams(ctx context.Context, in *UpdateParamsRequest, opts ...grpc.CallOption) (*UpdateParamsResponse, error) {
	out := new(UpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/polaris.evm.v1alpha1.MsgService/UpdateParams", in, out, opts...)
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
		FullMethod: "/polaris.evm.v1alpha1.MsgService/EthTransaction",
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
		FullMethod: "/polaris.evm.v1alpha1.MsgService/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).UpdateParams(ctx, req.(*UpdateParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "polaris.evm.v1alpha1.MsgService",
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
	Metadata: "polaris/evm/v1alpha1/tx.proto",
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
	if len(m.ReturnData) > 0 {
		i -= len(m.ReturnData)
		copy(dAtA[i:], m.ReturnData)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ReturnData)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VmError) > 0 {
		i -= len(m.VmError)
		copy(dAtA[i:], m.VmError)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VmError)))
		i--
		dAtA[i] = 0x12
	}
	if m.GasUsed != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x8
	}
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
	return n
}

func (m *EthTransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GasUsed != 0 {
		n += 1 + sovTx(uint64(m.GasUsed))
	}
	l = len(m.VmError)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ReturnData)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VmError", wireType)
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
			m.VmError = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnData", wireType)
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
			m.ReturnData = append(m.ReturnData[:0], dAtA[iNdEx:postIndex]...)
			if m.ReturnData == nil {
				m.ReturnData = []byte{}
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
