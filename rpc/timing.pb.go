// Code generated by protoc-gen-go. DO NOT EDIT.
// source: timing.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Timing struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Dir                  string   `protobuf:"bytes,3,opt,name=dir,proto3" json:"dir,omitempty"`
	Program              string   `protobuf:"bytes,4,opt,name=program,proto3" json:"program,omitempty"`
	Args                 string   `protobuf:"bytes,5,opt,name=args,proto3" json:"args,omitempty"`
	StdOut               string   `protobuf:"bytes,6,opt,name=std_out,json=stdOut,proto3" json:"std_out,omitempty"`
	StdErr               string   `protobuf:"bytes,7,opt,name=std_err,json=stdErr,proto3" json:"std_err,omitempty"`
	Time                 int64    `protobuf:"varint,8,opt,name=time,proto3" json:"time,omitempty"`
	Timeout              int64    `protobuf:"varint,9,opt,name=timeout,proto3" json:"timeout,omitempty"`
	IsMonitor            bool     `protobuf:"varint,10,opt,name=is_monitor,json=isMonitor,proto3" json:"is_monitor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timing) Reset()         { *m = Timing{} }
func (m *Timing) String() string { return proto.CompactTextString(m) }
func (*Timing) ProtoMessage()    {}
func (*Timing) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae56e3e700389cde, []int{0}
}

func (m *Timing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timing.Unmarshal(m, b)
}
func (m *Timing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timing.Marshal(b, m, deterministic)
}
func (m *Timing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timing.Merge(m, src)
}
func (m *Timing) XXX_Size() int {
	return xxx_messageInfo_Timing.Size(m)
}
func (m *Timing) XXX_DiscardUnknown() {
	xxx_messageInfo_Timing.DiscardUnknown(m)
}

var xxx_messageInfo_Timing proto.InternalMessageInfo

func (m *Timing) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Timing) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Timing) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

func (m *Timing) GetProgram() string {
	if m != nil {
		return m.Program
	}
	return ""
}

func (m *Timing) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *Timing) GetStdOut() string {
	if m != nil {
		return m.StdOut
	}
	return ""
}

func (m *Timing) GetStdErr() string {
	if m != nil {
		return m.StdErr
	}
	return ""
}

func (m *Timing) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Timing) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Timing) GetIsMonitor() bool {
	if m != nil {
		return m.IsMonitor
	}
	return false
}

type TimingResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Timing               *Timing  `protobuf:"bytes,2,opt,name=timing,proto3" json:"timing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimingResponse) Reset()         { *m = TimingResponse{} }
func (m *TimingResponse) String() string { return proto.CompactTextString(m) }
func (*TimingResponse) ProtoMessage()    {}
func (*TimingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae56e3e700389cde, []int{1}
}

func (m *TimingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimingResponse.Unmarshal(m, b)
}
func (m *TimingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimingResponse.Marshal(b, m, deterministic)
}
func (m *TimingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimingResponse.Merge(m, src)
}
func (m *TimingResponse) XXX_Size() int {
	return xxx_messageInfo_TimingResponse.Size(m)
}
func (m *TimingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TimingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TimingResponse proto.InternalMessageInfo

func (m *TimingResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TimingResponse) GetTiming() *Timing {
	if m != nil {
		return m.Timing
	}
	return nil
}

type TimingListResponse struct {
	Code                 int32     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Timings              []*Timing `protobuf:"bytes,2,rep,name=timings,proto3" json:"timings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TimingListResponse) Reset()         { *m = TimingListResponse{} }
func (m *TimingListResponse) String() string { return proto.CompactTextString(m) }
func (*TimingListResponse) ProtoMessage()    {}
func (*TimingListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae56e3e700389cde, []int{2}
}

func (m *TimingListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimingListResponse.Unmarshal(m, b)
}
func (m *TimingListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimingListResponse.Marshal(b, m, deterministic)
}
func (m *TimingListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimingListResponse.Merge(m, src)
}
func (m *TimingListResponse) XXX_Size() int {
	return xxx_messageInfo_TimingListResponse.Size(m)
}
func (m *TimingListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TimingListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TimingListResponse proto.InternalMessageInfo

func (m *TimingListResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TimingListResponse) GetTimings() []*Timing {
	if m != nil {
		return m.Timings
	}
	return nil
}

func init() {
	proto.RegisterType((*Timing)(nil), "Timing")
	proto.RegisterType((*TimingResponse)(nil), "TimingResponse")
	proto.RegisterType((*TimingListResponse)(nil), "TimingListResponse")
}

func init() { proto.RegisterFile("timing.proto", fileDescriptor_ae56e3e700389cde) }

var fileDescriptor_ae56e3e700389cde = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x6a, 0xe3, 0x30,
	0x14, 0x85, 0x63, 0x3b, 0x96, 0xe3, 0x3b, 0x43, 0x66, 0xb8, 0xb3, 0x18, 0xe1, 0xfe, 0xc4, 0x35,
	0x14, 0xb2, 0xf2, 0x22, 0x7d, 0x82, 0x16, 0x42, 0x17, 0xfd, 0x03, 0x93, 0x6e, 0xba, 0x09, 0x4e,
	0x24, 0x8c, 0x16, 0xb2, 0x8c, 0xa4, 0x14, 0xfa, 0x4a, 0x7d, 0xb5, 0xbe, 0x44, 0x91, 0x1c, 0x97,
	0x94, 0xd2, 0xae, 0x7c, 0xf4, 0x1d, 0xee, 0xb9, 0x3a, 0x46, 0xf0, 0xdb, 0x0a, 0x29, 0xda, 0xa6,
	0xec, 0xb4, 0xb2, 0x2a, 0x83, 0x4d, 0x6d, 0x78, 0xaf, 0x8b, 0xb7, 0x00, 0xc8, 0xca, 0x9b, 0x38,
	0x85, 0x50, 0x30, 0x1a, 0xe4, 0xc1, 0x3c, 0xaa, 0x42, 0xc1, 0x10, 0x61, 0xdc, 0xd6, 0x92, 0xd3,
	0x30, 0x0f, 0xe6, 0x69, 0xe5, 0x35, 0xfe, 0x85, 0x88, 0x09, 0x4d, 0x23, 0x8f, 0x9c, 0x44, 0x0a,
	0x49, 0xa7, 0x55, 0xa3, 0x6b, 0x49, 0xc7, 0x9e, 0x0e, 0x47, 0x37, 0x5f, 0xeb, 0xc6, 0xd0, 0xb8,
	0x9f, 0x77, 0x1a, 0xff, 0x43, 0x62, 0x2c, 0x5b, 0xab, 0x9d, 0xa5, 0xc4, 0x63, 0x62, 0x2c, 0x7b,
	0xd8, 0xd9, 0xc1, 0xe0, 0x5a, 0xd3, 0xe4, 0xc3, 0x58, 0x6a, 0xed, 0x52, 0xac, 0x90, 0x9c, 0x4e,
	0xfc, 0xbd, 0xbc, 0x76, 0x3b, 0xdd, 0xd7, 0xa5, 0xa4, 0x1e, 0x0f, 0x47, 0x3c, 0x01, 0x10, 0x66,
	0x2d, 0x55, 0x2b, 0xac, 0xd2, 0x14, 0xf2, 0x60, 0x3e, 0xa9, 0x52, 0x61, 0xee, 0x7a, 0x50, 0x2c,
	0x61, 0xda, 0x97, 0xad, 0xb8, 0xe9, 0x54, 0x6b, 0xb8, 0x8b, 0xdf, 0x2a, 0xc6, 0x7d, 0xed, 0xb8,
	0xf2, 0x1a, 0x67, 0x40, 0xfa, 0xff, 0xe5, 0xab, 0xff, 0x5a, 0x24, 0xe5, 0x7e, 0x68, 0x8f, 0x8b,
	0x1b, 0xc0, 0x9e, 0xdc, 0x0a, 0x63, 0x7f, 0x8c, 0x3a, 0xf3, 0x37, 0x15, 0x6d, 0x63, 0x68, 0x98,
	0x47, 0x87, 0x59, 0x03, 0x5f, 0xbc, 0x06, 0x10, 0xaf, 0x84, 0xe4, 0x1a, 0xcf, 0x61, 0xec, 0x02,
	0x91, 0x94, 0x4b, 0xd9, 0xd9, 0x97, 0xec, 0x5f, 0xf9, 0x75, 0x4b, 0x31, 0xc2, 0x19, 0x44, 0xd7,
	0xdc, 0x62, 0x5c, 0xde, 0xd7, 0x92, 0x67, 0x7f, 0xca, 0xcf, 0x8d, 0x8a, 0x11, 0x1e, 0x41, 0x74,
	0xc9, 0x18, 0x0e, 0xab, 0xb2, 0xb4, 0x3c, 0x30, 0x4f, 0x81, 0x3c, 0x76, 0xac, 0xb6, 0xfc, 0x1b,
	0xff, 0x18, 0x48, 0xc5, 0xa5, 0x7a, 0xe6, 0xc3, 0x82, 0x43, 0xf7, 0x2a, 0x7e, 0x8a, 0x74, 0xb7,
	0xdd, 0x10, 0xff, 0x78, 0x2e, 0xde, 0x03, 0x00, 0x00, 0xff, 0xff, 0x98, 0x35, 0x2d, 0x44, 0x58,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimerClient is the client API for Timer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimerClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TimingListResponse, error)
	Get(ctx context.Context, in *Name, opts ...grpc.CallOption) (*TimingResponse, error)
	Add(ctx context.Context, in *Timing, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *Timing, opts ...grpc.CallOption) (*Response, error)
	Remove(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Response, error)
}

type timerClient struct {
	cc *grpc.ClientConn
}

func NewTimerClient(cc *grpc.ClientConn) TimerClient {
	return &timerClient{cc}
}

func (c *timerClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TimingListResponse, error) {
	out := new(TimingListResponse)
	err := c.cc.Invoke(ctx, "/Timer/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timerClient) Get(ctx context.Context, in *Name, opts ...grpc.CallOption) (*TimingResponse, error) {
	out := new(TimingResponse)
	err := c.cc.Invoke(ctx, "/Timer/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timerClient) Add(ctx context.Context, in *Timing, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Timer/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timerClient) Update(ctx context.Context, in *Timing, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Timer/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timerClient) Remove(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Timer/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimerServer is the server API for Timer service.
type TimerServer interface {
	List(context.Context, *Empty) (*TimingListResponse, error)
	Get(context.Context, *Name) (*TimingResponse, error)
	Add(context.Context, *Timing) (*Response, error)
	Update(context.Context, *Timing) (*Response, error)
	Remove(context.Context, *Name) (*Response, error)
}

// UnimplementedTimerServer can be embedded to have forward compatible implementations.
type UnimplementedTimerServer struct {
}

func (*UnimplementedTimerServer) List(ctx context.Context, req *Empty) (*TimingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedTimerServer) Get(ctx context.Context, req *Name) (*TimingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTimerServer) Add(ctx context.Context, req *Timing) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedTimerServer) Update(ctx context.Context, req *Timing) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTimerServer) Remove(ctx context.Context, req *Name) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func RegisterTimerServer(s *grpc.Server, srv TimerServer) {
	s.RegisterService(&_Timer_serviceDesc, srv)
}

func _Timer_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Timer/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimerServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Timer/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimerServer).Get(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timer_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Timer/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimerServer).Add(ctx, req.(*Timing))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timer_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Timer/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimerServer).Update(ctx, req.(*Timing))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timer_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimerServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Timer/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimerServer).Remove(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

var _Timer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Timer",
	HandlerType: (*TimerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Timer_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Timer_Get_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Timer_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Timer_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Timer_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timing.proto",
}
