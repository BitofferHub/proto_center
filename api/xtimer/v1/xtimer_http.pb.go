// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v4.25.1
// source: api/xtimer/v1/xtimer.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationXTimerCreateTimer = "/api.xtimer.v1.XTimer/CreateTimer"
const OperationXTimerEnableTimer = "/api.xtimer.v1.XTimer/EnableTimer"

type XTimerHTTPServer interface {
	CreateTimer(context.Context, *CreateTimerRequest) (*CreateTimerReply, error)
	EnableTimer(context.Context, *EnableTimerRequest) (*EnableTimerReply, error)
}

func RegisterXTimerHTTPServer(s *http.Server, srv XTimerHTTPServer) {
	r := s.Route("/")
	r.POST("/xtimer/create_timer", _XTimer_CreateTimer0_HTTP_Handler(srv))
	r.POST("/xtimer/enable_timer", _XTimer_EnableTimer0_HTTP_Handler(srv))
}

func _XTimer_CreateTimer0_HTTP_Handler(srv XTimerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTimerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationXTimerCreateTimer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTimer(ctx, req.(*CreateTimerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTimerReply)
		return ctx.Result(200, reply)
	}
}

func _XTimer_EnableTimer0_HTTP_Handler(srv XTimerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EnableTimerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationXTimerEnableTimer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EnableTimer(ctx, req.(*EnableTimerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EnableTimerReply)
		return ctx.Result(200, reply)
	}
}

type XTimerHTTPClient interface {
	CreateTimer(ctx context.Context, req *CreateTimerRequest, opts ...http.CallOption) (rsp *CreateTimerReply, err error)
	EnableTimer(ctx context.Context, req *EnableTimerRequest, opts ...http.CallOption) (rsp *EnableTimerReply, err error)
}

type XTimerHTTPClientImpl struct {
	cc *http.Client
}

func NewXTimerHTTPClient(client *http.Client) XTimerHTTPClient {
	return &XTimerHTTPClientImpl{client}
}

func (c *XTimerHTTPClientImpl) CreateTimer(ctx context.Context, in *CreateTimerRequest, opts ...http.CallOption) (*CreateTimerReply, error) {
	var out CreateTimerReply
	pattern := "/xtimer/create_timer"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationXTimerCreateTimer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *XTimerHTTPClientImpl) EnableTimer(ctx context.Context, in *EnableTimerRequest, opts ...http.CallOption) (*EnableTimerReply, error) {
	var out EnableTimerReply
	pattern := "/xtimer/enable_timer"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationXTimerEnableTimer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
