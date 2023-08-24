// Tencent is pleased to support the open source community by making tRPC available.
// Copyright (C) 2023 THL A29 Limited, a Tencent company. All rights reserved.
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License that can be found in the LICENSE file.

// Code generated by trpc-go/trpc-go-cmdline v2.0.17. DO NOT EDIT.
// source: echo.proto

package echo

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// EchoService defines service
type EchoService interface {
	// UnaryEcho UnaryEcho is unary echo.
	UnaryEcho(ctx context.Context, req *EchoRequest) (*EchoResponse, error)
}

func EchoService_UnaryEcho_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &EchoRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(EchoService).UnaryEcho(ctx, reqbody.(*EchoRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// EchoServer_ServiceDesc descriptor for server.RegisterService
var EchoServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.examples.echo.Echo",
	HandlerType: ((*EchoService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.examples.echo.Echo/UnaryEcho",
			Func: EchoService_UnaryEcho_Handler,
		},
	},
}

// RegisterEchoService register service
func RegisterEchoService(s server.Service, svr EchoService) {
	if err := s.Register(&EchoServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("Echo register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedEcho struct{}

// UnaryEcho UnaryEcho is unary echo.
func (s *UnimplementedEcho) UnaryEcho(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, errors.New("rpc UnaryEcho of service Echo is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// EchoClientProxy defines service client proxy
type EchoClientProxy interface {
	// UnaryEcho UnaryEcho is unary echo.
	UnaryEcho(ctx context.Context, req *EchoRequest, opts ...client.Option) (rsp *EchoResponse, err error)
}

type EchoClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewEchoClientProxy = func(opts ...client.Option) EchoClientProxy {
	return &EchoClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *EchoClientProxyImpl) UnaryEcho(ctx context.Context, req *EchoRequest, opts ...client.Option) (*EchoResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.examples.echo.Echo/UnaryEcho")
	msg.WithCalleeServiceName(EchoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("examples")
	msg.WithCalleeServer("echo")
	msg.WithCalleeService("Echo")
	msg.WithCalleeMethod("UnaryEcho")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &EchoResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
