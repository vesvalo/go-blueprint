// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: resources/proto/ms/ms.proto

package ms

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/protobuf/field_mask"
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

// Client API for Ms service

type MsService interface {
	Search(ctx context.Context, in *SearchIn, opts ...client.CallOption) (*SearchOut, error)
	New(ctx context.Context, in *NewIn, opts ...client.CallOption) (*NewOut, error)
}

type msService struct {
	c    client.Client
	name string
}

func NewMsService(name string, c client.Client) MsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "ms"
	}
	return &msService{
		c:    c,
		name: name,
	}
}

func (c *msService) Search(ctx context.Context, in *SearchIn, opts ...client.CallOption) (*SearchOut, error) {
	req := c.c.NewRequest(c.name, "Ms.Search", in)
	out := new(SearchOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msService) New(ctx context.Context, in *NewIn, opts ...client.CallOption) (*NewOut, error) {
	req := c.c.NewRequest(c.name, "Ms.New", in)
	out := new(NewOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ms service

type MsHandler interface {
	Search(context.Context, *SearchIn, *SearchOut) error
	New(context.Context, *NewIn, *NewOut) error
}

func RegisterMsHandler(s server.Server, hdlr MsHandler, opts ...server.HandlerOption) error {
	type ms interface {
		Search(ctx context.Context, in *SearchIn, out *SearchOut) error
		New(ctx context.Context, in *NewIn, out *NewOut) error
	}
	type Ms struct {
		ms
	}
	h := &msHandler{hdlr}
	return s.Handle(s.NewHandler(&Ms{h}, opts...))
}

type msHandler struct {
	MsHandler
}

func (h *msHandler) Search(ctx context.Context, in *SearchIn, out *SearchOut) error {
	return h.MsHandler.Search(ctx, in, out)
}

func (h *msHandler) New(ctx context.Context, in *NewIn, out *NewOut) error {
	return h.MsHandler.New(ctx, in, out)
}