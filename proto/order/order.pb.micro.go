// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/order/order.proto

package order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for Order service

func NewOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Order service

type OrderService interface {
	// 创建出行订单
	CreateOrder(ctx context.Context, in *ReqCreateOrder, opts ...client.CallOption) (*RespCreateOrder, error)
	// 接受订单 / 抢单
	AcceptOrder(ctx context.Context, in *ReqAcceptOrder, opts ...client.CallOption) (*RespAcceptOrder, error)
	// 订单确认上车
	ConfirmGetOn(ctx context.Context, in *ReqConfirmGetOn, opts ...client.CallOption) (*RespConfirmGetOn, error)
	// 订单行程开始
	StartOrder(ctx context.Context, in *ReqStartOrder, opts ...client.CallOption) (*RespStartOrder, error)
	// 订单行程完成
	CancelOrder(ctx context.Context, in *ReqCancelOrder, opts ...client.CallOption) (*RespCancelOrder, error)
	// 订单行程完成
	FinishOrder(ctx context.Context, in *ReqFinishOrder, opts ...client.CallOption) (*RespFinishOrder, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) CreateOrder(ctx context.Context, in *ReqCreateOrder, opts ...client.CallOption) (*RespCreateOrder, error) {
	req := c.c.NewRequest(c.name, "Order.CreateOrder", in)
	out := new(RespCreateOrder)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AcceptOrder(ctx context.Context, in *ReqAcceptOrder, opts ...client.CallOption) (*RespAcceptOrder, error) {
	req := c.c.NewRequest(c.name, "Order.AcceptOrder", in)
	out := new(RespAcceptOrder)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) ConfirmGetOn(ctx context.Context, in *ReqConfirmGetOn, opts ...client.CallOption) (*RespConfirmGetOn, error) {
	req := c.c.NewRequest(c.name, "Order.ConfirmGetOn", in)
	out := new(RespConfirmGetOn)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) StartOrder(ctx context.Context, in *ReqStartOrder, opts ...client.CallOption) (*RespStartOrder, error) {
	req := c.c.NewRequest(c.name, "Order.StartOrder", in)
	out := new(RespStartOrder)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CancelOrder(ctx context.Context, in *ReqCancelOrder, opts ...client.CallOption) (*RespCancelOrder, error) {
	req := c.c.NewRequest(c.name, "Order.CancelOrder", in)
	out := new(RespCancelOrder)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) FinishOrder(ctx context.Context, in *ReqFinishOrder, opts ...client.CallOption) (*RespFinishOrder, error) {
	req := c.c.NewRequest(c.name, "Order.FinishOrder", in)
	out := new(RespFinishOrder)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	// 创建出行订单
	CreateOrder(context.Context, *ReqCreateOrder, *RespCreateOrder) error
	// 接受订单 / 抢单
	AcceptOrder(context.Context, *ReqAcceptOrder, *RespAcceptOrder) error
	// 订单确认上车
	ConfirmGetOn(context.Context, *ReqConfirmGetOn, *RespConfirmGetOn) error
	// 订单行程开始
	StartOrder(context.Context, *ReqStartOrder, *RespStartOrder) error
	// 订单行程完成
	CancelOrder(context.Context, *ReqCancelOrder, *RespCancelOrder) error
	// 订单行程完成
	FinishOrder(context.Context, *ReqFinishOrder, *RespFinishOrder) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		CreateOrder(ctx context.Context, in *ReqCreateOrder, out *RespCreateOrder) error
		AcceptOrder(ctx context.Context, in *ReqAcceptOrder, out *RespAcceptOrder) error
		ConfirmGetOn(ctx context.Context, in *ReqConfirmGetOn, out *RespConfirmGetOn) error
		StartOrder(ctx context.Context, in *ReqStartOrder, out *RespStartOrder) error
		CancelOrder(ctx context.Context, in *ReqCancelOrder, out *RespCancelOrder) error
		FinishOrder(ctx context.Context, in *ReqFinishOrder, out *RespFinishOrder) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) CreateOrder(ctx context.Context, in *ReqCreateOrder, out *RespCreateOrder) error {
	return h.OrderHandler.CreateOrder(ctx, in, out)
}

func (h *orderHandler) AcceptOrder(ctx context.Context, in *ReqAcceptOrder, out *RespAcceptOrder) error {
	return h.OrderHandler.AcceptOrder(ctx, in, out)
}

func (h *orderHandler) ConfirmGetOn(ctx context.Context, in *ReqConfirmGetOn, out *RespConfirmGetOn) error {
	return h.OrderHandler.ConfirmGetOn(ctx, in, out)
}

func (h *orderHandler) StartOrder(ctx context.Context, in *ReqStartOrder, out *RespStartOrder) error {
	return h.OrderHandler.StartOrder(ctx, in, out)
}

func (h *orderHandler) CancelOrder(ctx context.Context, in *ReqCancelOrder, out *RespCancelOrder) error {
	return h.OrderHandler.CancelOrder(ctx, in, out)
}

func (h *orderHandler) FinishOrder(ctx context.Context, in *ReqFinishOrder, out *RespFinishOrder) error {
	return h.OrderHandler.FinishOrder(ctx, in, out)
}
