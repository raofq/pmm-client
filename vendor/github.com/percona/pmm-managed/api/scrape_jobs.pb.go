// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scrape_jobs.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScrapeJob struct {
	// Scrape job name: "example-job" (required)
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Scrape interval: "10s"
	Interval string `protobuf:"bytes,2,opt,name=interval" json:"interval,omitempty"`
	// Scrape timeout: "5s"
	Timeout string `protobuf:"bytes,3,opt,name=timeout" json:"timeout,omitempty"`
	// Metrics path: "/metrics"
	Path string `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	// Scheme for scrapping: "http" or "https"
	Scheme string `protobuf:"bytes,5,opt,name=scheme" json:"scheme,omitempty"`
	// Hosts and ports: "127.0.0.1:9090" (required)
	StatisTargets []string `protobuf:"bytes,6,rep,name=statis_targets,json=statisTargets" json:"statis_targets,omitempty"`
}

func (m *ScrapeJob) Reset()                    { *m = ScrapeJob{} }
func (m *ScrapeJob) String() string            { return proto.CompactTextString(m) }
func (*ScrapeJob) ProtoMessage()               {}
func (*ScrapeJob) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ScrapeJob) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScrapeJob) GetInterval() string {
	if m != nil {
		return m.Interval
	}
	return ""
}

func (m *ScrapeJob) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

func (m *ScrapeJob) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ScrapeJob) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *ScrapeJob) GetStatisTargets() []string {
	if m != nil {
		return m.StatisTargets
	}
	return nil
}

type ScrapeJobsListRequest struct {
}

func (m *ScrapeJobsListRequest) Reset()                    { *m = ScrapeJobsListRequest{} }
func (m *ScrapeJobsListRequest) String() string            { return proto.CompactTextString(m) }
func (*ScrapeJobsListRequest) ProtoMessage()               {}
func (*ScrapeJobsListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type ScrapeJobsListResponse struct {
	Error      Error        `protobuf:"varint,1,opt,name=error,enum=api.Error" json:"error,omitempty"`
	ScrapeJobs []*ScrapeJob `protobuf:"bytes,2,rep,name=scrape_jobs,json=scrapeJobs" json:"scrape_jobs,omitempty"`
}

func (m *ScrapeJobsListResponse) Reset()                    { *m = ScrapeJobsListResponse{} }
func (m *ScrapeJobsListResponse) String() string            { return proto.CompactTextString(m) }
func (*ScrapeJobsListResponse) ProtoMessage()               {}
func (*ScrapeJobsListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ScrapeJobsListResponse) GetError() Error {
	if m != nil {
		return m.Error
	}
	return Error_NO_ERROR
}

func (m *ScrapeJobsListResponse) GetScrapeJobs() []*ScrapeJob {
	if m != nil {
		return m.ScrapeJobs
	}
	return nil
}

type ScrapeJobsGetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ScrapeJobsGetRequest) Reset()                    { *m = ScrapeJobsGetRequest{} }
func (m *ScrapeJobsGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ScrapeJobsGetRequest) ProtoMessage()               {}
func (*ScrapeJobsGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *ScrapeJobsGetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ScrapeJobsGetResponse struct {
	Error     Error      `protobuf:"varint,1,opt,name=error,enum=api.Error" json:"error,omitempty"`
	ScrapeJob *ScrapeJob `protobuf:"bytes,2,opt,name=scrape_job,json=scrapeJob" json:"scrape_job,omitempty"`
}

func (m *ScrapeJobsGetResponse) Reset()                    { *m = ScrapeJobsGetResponse{} }
func (m *ScrapeJobsGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ScrapeJobsGetResponse) ProtoMessage()               {}
func (*ScrapeJobsGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *ScrapeJobsGetResponse) GetError() Error {
	if m != nil {
		return m.Error
	}
	return Error_NO_ERROR
}

func (m *ScrapeJobsGetResponse) GetScrapeJob() *ScrapeJob {
	if m != nil {
		return m.ScrapeJob
	}
	return nil
}

type ScrapeJobsPutRequest struct {
	ScrapeJob *ScrapeJob `protobuf:"bytes,1,opt,name=scrape_job,json=scrapeJob" json:"scrape_job,omitempty"`
}

func (m *ScrapeJobsPutRequest) Reset()                    { *m = ScrapeJobsPutRequest{} }
func (m *ScrapeJobsPutRequest) String() string            { return proto.CompactTextString(m) }
func (*ScrapeJobsPutRequest) ProtoMessage()               {}
func (*ScrapeJobsPutRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *ScrapeJobsPutRequest) GetScrapeJob() *ScrapeJob {
	if m != nil {
		return m.ScrapeJob
	}
	return nil
}

type ScrapeJobsPutResponse struct {
	Error Error `protobuf:"varint,1,opt,name=error,enum=api.Error" json:"error,omitempty"`
}

func (m *ScrapeJobsPutResponse) Reset()                    { *m = ScrapeJobsPutResponse{} }
func (m *ScrapeJobsPutResponse) String() string            { return proto.CompactTextString(m) }
func (*ScrapeJobsPutResponse) ProtoMessage()               {}
func (*ScrapeJobsPutResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *ScrapeJobsPutResponse) GetError() Error {
	if m != nil {
		return m.Error
	}
	return Error_NO_ERROR
}

type ScrapeJobsDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ScrapeJobsDeleteRequest) Reset()                    { *m = ScrapeJobsDeleteRequest{} }
func (m *ScrapeJobsDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ScrapeJobsDeleteRequest) ProtoMessage()               {}
func (*ScrapeJobsDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *ScrapeJobsDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ScrapeJobsDeleteResponse struct {
	Error Error `protobuf:"varint,1,opt,name=error,enum=api.Error" json:"error,omitempty"`
}

func (m *ScrapeJobsDeleteResponse) Reset()                    { *m = ScrapeJobsDeleteResponse{} }
func (m *ScrapeJobsDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*ScrapeJobsDeleteResponse) ProtoMessage()               {}
func (*ScrapeJobsDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *ScrapeJobsDeleteResponse) GetError() Error {
	if m != nil {
		return m.Error
	}
	return Error_NO_ERROR
}

func init() {
	proto.RegisterType((*ScrapeJob)(nil), "api.ScrapeJob")
	proto.RegisterType((*ScrapeJobsListRequest)(nil), "api.ScrapeJobsListRequest")
	proto.RegisterType((*ScrapeJobsListResponse)(nil), "api.ScrapeJobsListResponse")
	proto.RegisterType((*ScrapeJobsGetRequest)(nil), "api.ScrapeJobsGetRequest")
	proto.RegisterType((*ScrapeJobsGetResponse)(nil), "api.ScrapeJobsGetResponse")
	proto.RegisterType((*ScrapeJobsPutRequest)(nil), "api.ScrapeJobsPutRequest")
	proto.RegisterType((*ScrapeJobsPutResponse)(nil), "api.ScrapeJobsPutResponse")
	proto.RegisterType((*ScrapeJobsDeleteRequest)(nil), "api.ScrapeJobsDeleteRequest")
	proto.RegisterType((*ScrapeJobsDeleteResponse)(nil), "api.ScrapeJobsDeleteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ScrapeJobs service

type ScrapeJobsClient interface {
	// List returns all scrape jobs.
	List(ctx context.Context, in *ScrapeJobsListRequest, opts ...grpc.CallOption) (*ScrapeJobsListResponse, error)
	// Get returns a scrape job by name.
	Get(ctx context.Context, in *ScrapeJobsGetRequest, opts ...grpc.CallOption) (*ScrapeJobsGetResponse, error)
	// Put creates or updates a scrape job by name.
	Put(ctx context.Context, in *ScrapeJobsPutRequest, opts ...grpc.CallOption) (*ScrapeJobsPutResponse, error)
	// Delete removes a scrape job by name.
	Delete(ctx context.Context, in *ScrapeJobsDeleteRequest, opts ...grpc.CallOption) (*ScrapeJobsDeleteResponse, error)
}

type scrapeJobsClient struct {
	cc *grpc.ClientConn
}

func NewScrapeJobsClient(cc *grpc.ClientConn) ScrapeJobsClient {
	return &scrapeJobsClient{cc}
}

func (c *scrapeJobsClient) List(ctx context.Context, in *ScrapeJobsListRequest, opts ...grpc.CallOption) (*ScrapeJobsListResponse, error) {
	out := new(ScrapeJobsListResponse)
	err := grpc.Invoke(ctx, "/api.ScrapeJobs/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scrapeJobsClient) Get(ctx context.Context, in *ScrapeJobsGetRequest, opts ...grpc.CallOption) (*ScrapeJobsGetResponse, error) {
	out := new(ScrapeJobsGetResponse)
	err := grpc.Invoke(ctx, "/api.ScrapeJobs/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scrapeJobsClient) Put(ctx context.Context, in *ScrapeJobsPutRequest, opts ...grpc.CallOption) (*ScrapeJobsPutResponse, error) {
	out := new(ScrapeJobsPutResponse)
	err := grpc.Invoke(ctx, "/api.ScrapeJobs/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scrapeJobsClient) Delete(ctx context.Context, in *ScrapeJobsDeleteRequest, opts ...grpc.CallOption) (*ScrapeJobsDeleteResponse, error) {
	out := new(ScrapeJobsDeleteResponse)
	err := grpc.Invoke(ctx, "/api.ScrapeJobs/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ScrapeJobs service

type ScrapeJobsServer interface {
	// List returns all scrape jobs.
	List(context.Context, *ScrapeJobsListRequest) (*ScrapeJobsListResponse, error)
	// Get returns a scrape job by name.
	Get(context.Context, *ScrapeJobsGetRequest) (*ScrapeJobsGetResponse, error)
	// Put creates or updates a scrape job by name.
	Put(context.Context, *ScrapeJobsPutRequest) (*ScrapeJobsPutResponse, error)
	// Delete removes a scrape job by name.
	Delete(context.Context, *ScrapeJobsDeleteRequest) (*ScrapeJobsDeleteResponse, error)
}

func RegisterScrapeJobsServer(s *grpc.Server, srv ScrapeJobsServer) {
	s.RegisterService(&_ScrapeJobs_serviceDesc, srv)
}

func _ScrapeJobs_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrapeJobsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapeJobsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ScrapeJobs/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapeJobsServer).List(ctx, req.(*ScrapeJobsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScrapeJobs_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrapeJobsGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapeJobsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ScrapeJobs/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapeJobsServer).Get(ctx, req.(*ScrapeJobsGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScrapeJobs_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrapeJobsPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapeJobsServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ScrapeJobs/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapeJobsServer).Put(ctx, req.(*ScrapeJobsPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScrapeJobs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrapeJobsDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapeJobsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ScrapeJobs/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapeJobsServer).Delete(ctx, req.(*ScrapeJobsDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScrapeJobs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ScrapeJobs",
	HandlerType: (*ScrapeJobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ScrapeJobs_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ScrapeJobs_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _ScrapeJobs_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ScrapeJobs_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scrape_jobs.proto",
}

func init() { proto.RegisterFile("scrape_jobs.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x38, 0x0d, 0x64, 0x2a, 0x82, 0x3a, 0x82, 0x64, 0x59, 0x0a, 0x0a, 0x2b, 0x81,
	0xa2, 0x88, 0xc4, 0x28, 0x9c, 0x40, 0x1c, 0xa9, 0x2a, 0x21, 0x0e, 0x95, 0xe1, 0xc0, 0x01, 0xa9,
	0xda, 0x54, 0xa3, 0xc4, 0x25, 0xf1, 0x1a, 0xef, 0xa6, 0x17, 0xc4, 0x85, 0x57, 0xe0, 0xc0, 0x2b,
	0xf0, 0x3e, 0xbc, 0x02, 0x0f, 0x82, 0xbc, 0x9b, 0xda, 0x8e, 0x9d, 0xd0, 0xdc, 0x76, 0x67, 0x66,
	0xbf, 0xf9, 0xe7, 0x1f, 0xcb, 0x70, 0xa4, 0x2f, 0x52, 0x99, 0xd0, 0xf9, 0xa5, 0x9a, 0xea, 0x71,
	0x92, 0x2a, 0xa3, 0xd0, 0x97, 0x49, 0xc4, 0x8f, 0x67, 0x4a, 0xcd, 0x16, 0x14, 0xc8, 0x24, 0x0a,
	0x64, 0x1c, 0x2b, 0x23, 0x4d, 0xa4, 0xe2, 0x75, 0x09, 0x87, 0xa9, 0xd4, 0xe4, 0xce, 0xe2, 0xb7,
	0x07, 0xed, 0x0f, 0x16, 0xf2, 0x4e, 0x4d, 0x11, 0xa1, 0x19, 0xcb, 0x25, 0x31, 0xaf, 0xef, 0x0d,
	0xda, 0xa1, 0x3d, 0x23, 0x87, 0xdb, 0x51, 0x6c, 0x28, 0xbd, 0x92, 0x0b, 0xd6, 0xb0, 0xf1, 0xfc,
	0x8e, 0x0c, 0x6e, 0x99, 0x68, 0x49, 0x6a, 0x65, 0x98, 0x6f, 0x53, 0xd7, 0xd7, 0x8c, 0x94, 0x48,
	0x33, 0x67, 0x4d, 0x47, 0xca, 0xce, 0xd8, 0x85, 0x96, 0xbe, 0x98, 0xd3, 0x92, 0xd8, 0x81, 0x8d,
	0xae, 0x6f, 0xf8, 0x14, 0x3a, 0x3a, 0x53, 0xa8, 0xcf, 0x8d, 0x4c, 0x67, 0x64, 0x34, 0x6b, 0xf5,
	0xfd, 0x41, 0x3b, 0xbc, 0xe3, 0xa2, 0x1f, 0x5d, 0x50, 0xf4, 0xe0, 0x7e, 0xae, 0x54, 0xbf, 0x8f,
	0xb4, 0x09, 0xe9, 0xeb, 0x8a, 0xb4, 0x11, 0x5f, 0xa0, 0x5b, 0x4d, 0xe8, 0x44, 0xc5, 0x9a, 0xb0,
	0x0f, 0x07, 0x94, 0xa6, 0x2a, 0xb5, 0x03, 0x75, 0x26, 0x30, 0x96, 0x49, 0x34, 0x3e, 0xc9, 0x22,
	0xa1, 0x4b, 0x60, 0x00, 0x87, 0x25, 0x0f, 0x59, 0xa3, 0xef, 0x0f, 0x0e, 0x27, 0x1d, 0x5b, 0x97,
	0x33, 0x43, 0xd0, 0x39, 0x5e, 0x0c, 0xe1, 0x5e, 0xd1, 0xec, 0x94, 0xae, 0x45, 0x6c, 0xb3, 0x4e,
	0xcc, 0xcb, 0x8a, 0x6d, 0xed, 0xde, 0xba, 0x46, 0x00, 0x85, 0x2e, 0xeb, 0x7b, 0x5d, 0x56, 0x3b,
	0x97, 0x25, 0x4e, 0xca, 0xaa, 0xce, 0x56, 0xb9, 0xaa, 0x4d, 0x8c, 0x77, 0x13, 0xe6, 0x55, 0x59,
	0xb0, 0xc5, 0xec, 0x2b, 0x58, 0x8c, 0xa0, 0x57, 0x3c, 0x7d, 0x4b, 0x0b, 0x32, 0xf4, 0x3f, 0x6b,
	0xde, 0x00, 0xab, 0x97, 0xef, 0xdb, 0x6c, 0xf2, 0xcb, 0x07, 0x28, 0x9e, 0xe3, 0x27, 0x68, 0x66,
	0x6b, 0x47, 0xbe, 0x39, 0x59, 0xf9, 0x23, 0xe1, 0x0f, 0xb7, 0xe6, 0x5c, 0x47, 0xd1, 0xfb, 0xf1,
	0xe7, 0xef, 0xcf, 0xc6, 0x11, 0xde, 0x0d, 0xae, 0x5e, 0x04, 0xce, 0x8e, 0x51, 0xf6, 0x3d, 0xe0,
	0x67, 0xf0, 0x4f, 0xc9, 0xe0, 0x83, 0xca, 0xe3, 0x62, 0xef, 0x9c, 0x6f, 0x4b, 0xad, 0xb1, 0x8f,
	0x2d, 0x96, 0x61, 0xb7, 0x82, 0x0d, 0xbe, 0x65, 0x1e, 0x7c, 0xc7, 0x4b, 0xf0, 0xcf, 0x56, 0x75,
	0x7a, 0xb1, 0xbf, 0x1a, 0xbd, 0xb4, 0x13, 0xf1, 0xdc, 0xd2, 0x9f, 0xf1, 0x27, 0x35, 0x7a, 0xb1,
	0xf2, 0xb1, 0x6d, 0xf4, 0xda, 0x1b, 0x22, 0x41, 0xcb, 0xd9, 0x8c, 0xc7, 0x15, 0xe6, 0xc6, 0xb2,
	0xf8, 0xa3, 0x1d, 0xd9, 0xcd, 0x91, 0x86, 0x3b, 0x46, 0x9a, 0xb6, 0xec, 0x6f, 0xe5, 0xe5, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xa4, 0xdd, 0x5c, 0x9a, 0x04, 0x00, 0x00,
}
