// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/dice.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DiceService_RollDice_FullMethodName = "/dice.DiceService/RollDice"
)

// DiceServiceClient is the client API for DiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiceServiceClient interface {
	RollDice(ctx context.Context, in *RollDiceRequest, opts ...grpc.CallOption) (*RollDiceResponse, error)
}

type diceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiceServiceClient(cc grpc.ClientConnInterface) DiceServiceClient {
	return &diceServiceClient{cc}
}

func (c *diceServiceClient) RollDice(ctx context.Context, in *RollDiceRequest, opts ...grpc.CallOption) (*RollDiceResponse, error) {
	out := new(RollDiceResponse)
	err := c.cc.Invoke(ctx, DiceService_RollDice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiceServiceServer is the server API for DiceService service.
// All implementations must embed UnimplementedDiceServiceServer
// for forward compatibility
type DiceServiceServer interface {
	RollDice(context.Context, *RollDiceRequest) (*RollDiceResponse, error)
	mustEmbedUnimplementedDiceServiceServer()
}

// UnimplementedDiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiceServiceServer struct {
}

func (UnimplementedDiceServiceServer) RollDice(context.Context, *RollDiceRequest) (*RollDiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollDice not implemented")
}
func (UnimplementedDiceServiceServer) mustEmbedUnimplementedDiceServiceServer() {}

// UnsafeDiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiceServiceServer will
// result in compilation errors.
type UnsafeDiceServiceServer interface {
	mustEmbedUnimplementedDiceServiceServer()
}

func RegisterDiceServiceServer(s grpc.ServiceRegistrar, srv DiceServiceServer) {
	s.RegisterService(&DiceService_ServiceDesc, srv)
}

func _DiceService_RollDice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollDiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiceServiceServer).RollDice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiceService_RollDice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiceServiceServer).RollDice(ctx, req.(*RollDiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiceService_ServiceDesc is the grpc.ServiceDesc for DiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dice.DiceService",
	HandlerType: (*DiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RollDice",
			Handler:    _DiceService_RollDice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dice.proto",
}
