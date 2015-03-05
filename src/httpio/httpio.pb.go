// Code generated by protoc-gen-go.
// source: httpio.proto
// DO NOT EDIT!

/*
Package httpio is a generated protocol buffer package.

It is generated from these files:
	httpio.proto

It has these top-level messages:
	HttpGetRequest
	HttpGetReply
*/
package httpio

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type HttpGetRequest struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *HttpGetRequest) Reset()         { *m = HttpGetRequest{} }
func (m *HttpGetRequest) String() string { return proto.CompactTextString(m) }
func (*HttpGetRequest) ProtoMessage()    {}

type HttpGetReply struct {
	Duration uint64 `protobuf:"varint,1,opt,name=duration" json:"duration,omitempty"`
}

func (m *HttpGetReply) Reset()         { *m = HttpGetReply{} }
func (m *HttpGetReply) String() string { return proto.CompactTextString(m) }
func (*HttpGetReply) ProtoMessage()    {}

func init() {
}

// Client API for Requester service

type RequesterClient interface {
	GetRequest(ctx context.Context, in *HttpGetRequest, opts ...grpc.CallOption) (*HttpGetReply, error)
}

type requesterClient struct {
	cc *grpc.ClientConn
}

func NewRequesterClient(cc *grpc.ClientConn) RequesterClient {
	return &requesterClient{cc}
}

func (c *requesterClient) GetRequest(ctx context.Context, in *HttpGetRequest, opts ...grpc.CallOption) (*HttpGetReply, error) {
	out := new(HttpGetReply)
	err := grpc.Invoke(ctx, "/httpio.Requester/GetRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Requester service

type RequesterServer interface {
	GetRequest(context.Context, *HttpGetRequest) (*HttpGetReply, error)
}

func RegisterRequesterServer(s *grpc.Server, srv RequesterServer) {
	s.RegisterService(&_Requester_serviceDesc, srv)
}

func _Requester_GetRequest_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(HttpGetRequest)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RequesterServer).GetRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Requester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "httpio.Requester",
	HandlerType: (*RequesterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRequest",
			Handler:    _Requester_GetRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
