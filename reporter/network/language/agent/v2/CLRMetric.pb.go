// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent-v2/CLRMetric.proto

package v2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	v2 "github.com/SkyAPM/SkyAPM-php-sdk/reporter/network/common/v2"
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

type CLRMetricCollectionV2 struct {
	Metrics              []*CLRMetricV2 `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	ServiceInstanceId    int32          `protobuf:"varint,2,opt,name=serviceInstanceId,proto3" json:"serviceInstanceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CLRMetricCollectionV2) Reset()         { *m = CLRMetricCollectionV2{} }
func (m *CLRMetricCollectionV2) String() string { return proto.CompactTextString(m) }
func (*CLRMetricCollectionV2) ProtoMessage()    {}
func (*CLRMetricCollectionV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f032372dea03c68, []int{0}
}

func (m *CLRMetricCollectionV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CLRMetricCollectionV2.Unmarshal(m, b)
}
func (m *CLRMetricCollectionV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CLRMetricCollectionV2.Marshal(b, m, deterministic)
}
func (m *CLRMetricCollectionV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CLRMetricCollectionV2.Merge(m, src)
}
func (m *CLRMetricCollectionV2) XXX_Size() int {
	return xxx_messageInfo_CLRMetricCollectionV2.Size(m)
}
func (m *CLRMetricCollectionV2) XXX_DiscardUnknown() {
	xxx_messageInfo_CLRMetricCollectionV2.DiscardUnknown(m)
}

var xxx_messageInfo_CLRMetricCollectionV2 proto.InternalMessageInfo

func (m *CLRMetricCollectionV2) GetMetrics() []*CLRMetricV2 {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *CLRMetricCollectionV2) GetServiceInstanceId() int32 {
	if m != nil {
		return m.ServiceInstanceId
	}
	return 0
}

func init() {
	proto.RegisterType((*CLRMetricCollectionV2)(nil), "CLRMetricCollectionV2")
}

func init() { proto.RegisterFile("language-agent-v2/CLRMetric.proto", fileDescriptor_1f032372dea03c68) }

var fileDescriptor_1f032372dea03c68 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xdf, 0x4a, 0xfb, 0x30,
	0x14, 0x80, 0x7f, 0xdd, 0x0f, 0x1d, 0x64, 0x5e, 0x68, 0xc5, 0x31, 0x7a, 0x35, 0x87, 0xc8, 0x2e,
	0xdc, 0x89, 0x64, 0x6f, 0x60, 0xaf, 0x26, 0x53, 0x46, 0x07, 0x15, 0xbc, 0x8b, 0xd9, 0xa1, 0x96,
	0x36, 0x39, 0x25, 0x8d, 0x1d, 0xf3, 0x91, 0x7c, 0x4a, 0x59, 0xff, 0x79, 0xa1, 0x57, 0x81, 0xef,
	0x7c, 0x39, 0x09, 0x1f, 0xbb, 0xce, 0xa5, 0x49, 0x3e, 0x64, 0x82, 0x0b, 0x99, 0xa0, 0x71, 0x8b,
	0x4a, 0xf0, 0x70, 0x1d, 0x3d, 0xa1, 0xb3, 0xa9, 0x82, 0xc2, 0x92, 0xa3, 0xe0, 0x52, 0x91, 0xd6,
	0x64, 0x78, 0x73, 0xb4, 0xf0, 0xbc, 0x85, 0xe1, 0x3a, 0x6a, 0xc8, 0x4c, 0xb3, 0xab, 0xfe, 0x66,
	0x48, 0x79, 0x8e, 0xca, 0xa5, 0x64, 0x62, 0xe1, 0xdf, 0xb2, 0xa1, 0xae, 0x69, 0x39, 0xf1, 0xa6,
	0xff, 0xe7, 0x23, 0x71, 0x06, 0xbd, 0x18, 0x8b, 0xa8, 0x1b, 0xfa, 0x77, 0xec, 0xa2, 0x44, 0x5b,
	0xa5, 0x0a, 0x57, 0xa6, 0x74, 0xd2, 0x28, 0x5c, 0xed, 0x26, 0x83, 0xa9, 0x37, 0x3f, 0x89, 0x7e,
	0x0f, 0xc4, 0x23, 0x1b, 0xf7, 0x5b, 0x22, 0x2c, 0xc8, 0xba, 0x6d, 0xe3, 0xf8, 0xf7, 0x6c, 0xa8,
	0x9a, 0xf7, 0xfd, 0x31, 0xfc, 0xf9, 0xa5, 0x60, 0x04, 0x21, 0x69, 0x2d, 0xcd, 0xae, 0x8c, 0xc5,
	0xec, 0xdf, 0xc3, 0x27, 0x5b, 0x92, 0x4d, 0x40, 0x16, 0x52, 0xbd, 0x23, 0x94, 0xd9, 0x61, 0x2f,
	0xf3, 0x2c, 0x35, 0x47, 0xa2, 0xc1, 0xa0, 0xdb, 0x93, 0xcd, 0xa0, 0x0b, 0x05, 0x75, 0x28, 0xa8,
	0xc4, 0xc6, 0x7b, 0xbd, 0xf9, 0x71, 0x79, 0xeb, 0xf1, 0xce, 0xe3, 0xb5, 0xc7, 0x2b, 0xf1, 0x35,
	0x08, 0xb6, 0xd9, 0xe1, 0xa5, 0x5d, 0xf9, 0xdc, 0x68, 0x9b, 0x63, 0x33, 0x45, 0xf9, 0xdb, 0x69,
	0x5d, 0x6f, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xc8, 0x68, 0x00, 0x89, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CLRMetricReportServiceClient is the client API for CLRMetricReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CLRMetricReportServiceClient interface {
	Collect(ctx context.Context, in *CLRMetricCollectionV2, opts ...grpc.CallOption) (*v2.CommandsV2, error)
}

type cLRMetricReportServiceClient struct {
	cc *grpc.ClientConn
}

func NewCLRMetricReportServiceClient(cc *grpc.ClientConn) CLRMetricReportServiceClient {
	return &cLRMetricReportServiceClient{cc}
}

func (c *cLRMetricReportServiceClient) Collect(ctx context.Context, in *CLRMetricCollectionV2, opts ...grpc.CallOption) (*v2.CommandsV2, error) {
	out := new(v2.CommandsV2)
	err := c.cc.Invoke(ctx, "/CLRMetricReportService/collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CLRMetricReportServiceServer is the server API for CLRMetricReportService service.
type CLRMetricReportServiceServer interface {
	Collect(context.Context, *CLRMetricCollectionV2) (*v2.CommandsV2, error)
}

// UnimplementedCLRMetricReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCLRMetricReportServiceServer struct {
}

func (*UnimplementedCLRMetricReportServiceServer) Collect(ctx context.Context, req *CLRMetricCollectionV2) (*v2.CommandsV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterCLRMetricReportServiceServer(s *grpc.Server, srv CLRMetricReportServiceServer) {
	s.RegisterService(&_CLRMetricReportService_serviceDesc, srv)
}

func _CLRMetricReportService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CLRMetricCollectionV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CLRMetricReportServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CLRMetricReportService/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CLRMetricReportServiceServer).Collect(ctx, req.(*CLRMetricCollectionV2))
	}
	return interceptor(ctx, in, info, handler)
}

var _CLRMetricReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CLRMetricReportService",
	HandlerType: (*CLRMetricReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "collect",
			Handler:    _CLRMetricReportService_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent-v2/CLRMetric.proto",
}