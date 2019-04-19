// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/service/inventory.proto

package mu_micro_book_srv_inventory

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Service service

type Service interface {
	Sell(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Confirm(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type service struct {
	c    client.Client
	name string
}

func NewService(name string, c client.Client) Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mu.micro.book.srv.inventory"
	}
	return &service{
		c:    c,
		name: name,
	}
}

func (c *service) Sell(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Service.Sell", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) Confirm(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Service.Confirm", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceHandler interface {
	Sell(context.Context, *Request, *Response) error
	Confirm(context.Context, *Request, *Response) error
}

func RegisterServiceHandler(s server.Server, hdlr ServiceHandler, opts ...server.HandlerOption) error {
	type service interface {
		Sell(ctx context.Context, in *Request, out *Response) error
		Confirm(ctx context.Context, in *Request, out *Response) error
	}
	type Service struct {
		service
	}
	h := &serviceHandler{hdlr}
	return s.Handle(s.NewHandler(&Service{h}, opts...))
}

type serviceHandler struct {
	ServiceHandler
}

func (h *serviceHandler) Sell(ctx context.Context, in *Request, out *Response) error {
	return h.ServiceHandler.Sell(ctx, in, out)
}

func (h *serviceHandler) Confirm(ctx context.Context, in *Request, out *Response) error {
	return h.ServiceHandler.Confirm(ctx, in, out)
}