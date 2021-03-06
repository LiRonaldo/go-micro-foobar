// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/foobar/foobar.proto

package go_micro_service_foobar

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

// Client API for Foobar service

type FoobarService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Foobar_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Foobar_PingPongService, error)
}

type foobarService struct {
	c    client.Client
	name string
}

func NewFoobarService(name string, c client.Client) FoobarService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.service.foobar"
	}
	return &foobarService{
		c:    c,
		name: name,
	}
}

func (c *foobarService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Foobar.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Foobar_StreamService, error) {
	req := c.c.NewRequest(c.name, "Foobar.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &foobarServiceStream{stream}, nil
}

type Foobar_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type foobarServiceStream struct {
	stream client.Stream
}

func (x *foobarServiceStream) Close() error {
	return x.stream.Close()
}

func (x *foobarServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *foobarServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *foobarServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *foobarService) PingPong(ctx context.Context, opts ...client.CallOption) (Foobar_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Foobar.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &foobarServicePingPong{stream}, nil
}

type Foobar_PingPongService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type foobarServicePingPong struct {
	stream client.Stream
}

func (x *foobarServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *foobarServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *foobarServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *foobarServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *foobarServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Foobar service

type FoobarHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Foobar_StreamStream) error
	PingPong(context.Context, Foobar_PingPongStream) error
}

func RegisterFoobarHandler(s server.Server, hdlr FoobarHandler, opts ...server.HandlerOption) error {
	type foobar interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Foobar struct {
		foobar
	}
	h := &foobarHandler{hdlr}
	return s.Handle(s.NewHandler(&Foobar{h}, opts...))
}

type foobarHandler struct {
	FoobarHandler
}

func (h *foobarHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.FoobarHandler.Call(ctx, in, out)
}

func (h *foobarHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.FoobarHandler.Stream(ctx, m, &foobarStreamStream{stream})
}

type Foobar_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type foobarStreamStream struct {
	stream server.Stream
}

func (x *foobarStreamStream) Close() error {
	return x.stream.Close()
}

func (x *foobarStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *foobarStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *foobarStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *foobarHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.FoobarHandler.PingPong(ctx, &foobarPingPongStream{stream})
}

type Foobar_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type foobarPingPongStream struct {
	stream server.Stream
}

func (x *foobarPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *foobarPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *foobarPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *foobarPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *foobarPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
