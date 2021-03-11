// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/fibo.proto

package fibo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Fibo service

func NewFiboEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Fibo service

type FiboService interface {
	Cal(ctx context.Context, in *FiboRequest, opts ...client.CallOption) (*FiboResponse, error)
}

type fiboService struct {
	c    client.Client
	name string
}

func NewFiboService(name string, c client.Client) FiboService {
	return &fiboService{
		c:    c,
		name: name,
	}
}

func (c *fiboService) Cal(ctx context.Context, in *FiboRequest, opts ...client.CallOption) (*FiboResponse, error) {
	req := c.c.NewRequest(c.name, "Fibo.Cal", in)
	out := new(FiboResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Fibo service

type FiboHandler interface {
	Cal(context.Context, *FiboRequest, *FiboResponse) error
}

func RegisterFiboHandler(s server.Server, hdlr FiboHandler, opts ...server.HandlerOption) error {
	type fibo interface {
		Cal(ctx context.Context, in *FiboRequest, out *FiboResponse) error
	}
	type Fibo struct {
		fibo
	}
	h := &fiboHandler{hdlr}
	return s.Handle(s.NewHandler(&Fibo{h}, opts...))
}

type fiboHandler struct {
	FiboHandler
}

func (h *fiboHandler) Cal(ctx context.Context, in *FiboRequest, out *FiboResponse) error {
	return h.FiboHandler.Cal(ctx, in, out)
}