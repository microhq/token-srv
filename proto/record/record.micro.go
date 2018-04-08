// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/token-srv/proto/record/record.proto

/*
Package record is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/token-srv/proto/record/record.proto

It has these top-level messages:
	Token
	CreateRequest
	CreateResponse
	DeleteRequest
	DeleteResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	SearchRequest
	SearchResponse
	GenerateRequest
	GenerateResponse
*/
package record

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Record service

type RecordService interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type recordService struct {
	c           client.Client
	serviceName string
}

func RecordServiceClient(serviceName string, c client.Client) RecordService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "record"
	}
	return &recordService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *recordService) Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Record.Generate", in)
	out := new(GenerateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Record.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Record.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Record.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Record.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Record.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Record service

type RecordHandler interface {
	Generate(context.Context, *GenerateRequest, *GenerateResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterRecordHandler(s server.Server, hdlr RecordHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Record{hdlr}, opts...))
}

type Record struct {
	RecordHandler
}

func (h *Record) Generate(ctx context.Context, in *GenerateRequest, out *GenerateResponse) error {
	return h.RecordHandler.Generate(ctx, in, out)
}

func (h *Record) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.RecordHandler.Create(ctx, in, out)
}

func (h *Record) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.RecordHandler.Read(ctx, in, out)
}

func (h *Record) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.RecordHandler.Delete(ctx, in, out)
}

func (h *Record) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.RecordHandler.Update(ctx, in, out)
}

func (h *Record) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.RecordHandler.Search(ctx, in, out)
}