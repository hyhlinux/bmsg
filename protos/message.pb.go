// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/message.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MessgeStatus int32

const (
	MessgeStatus_UNSEEN MessgeStatus = 0
	MessgeStatus_SEEN   MessgeStatus = 1
	MessgeStatus_ALL    MessgeStatus = 1
)

var MessgeStatus_name = map[int32]string{
	0: "UNSEEN",
	1: "SEEN",
	// Duplicate value: 1: "ALL",
}
var MessgeStatus_value = map[string]int32{
	"UNSEEN": 0,
	"SEEN":   1,
	"ALL":    1,
}

func (x MessgeStatus) String() string {
	return proto.EnumName(MessgeStatus_name, int32(x))
}
func (MessgeStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type Messge struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	FromUserId int64  `protobuf:"varint,2,opt,name=fromUserId" json:"fromUserId,omitempty"`
	ToUserId   int64  `protobuf:"varint,3,opt,name=toUserId" json:"toUserId,omitempty"`
	// 时间戳ms
	CreatedAt int64  `protobuf:"varint,4,opt,name=createdAt" json:"createdAt,omitempty"`
	Title     string `protobuf:"bytes,5,opt,name=title" json:"title,omitempty"`
	Message   string `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	Status    string `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	IsDelete  bool   `protobuf:"varint,8,opt,name=isDelete" json:"isDelete,omitempty"`
}

func (m *Messge) Reset()                    { *m = Messge{} }
func (m *Messge) String() string            { return proto.CompactTextString(m) }
func (*Messge) ProtoMessage()               {}
func (*Messge) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Messge) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Messge) GetFromUserId() int64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func (m *Messge) GetToUserId() int64 {
	if m != nil {
		return m.ToUserId
	}
	return 0
}

func (m *Messge) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Messge) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Messge) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Messge) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Messge) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

type MessgeResponse struct {
	// 创建成功id
	Mid string `protobuf:"bytes,1,opt,name=mid" json:"mid,omitempty"`
}

func (m *MessgeResponse) Reset()                    { *m = MessgeResponse{} }
func (m *MessgeResponse) String() string            { return proto.CompactTextString(m) }
func (*MessgeResponse) ProtoMessage()               {}
func (*MessgeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *MessgeResponse) GetMid() string {
	if m != nil {
		return m.Mid
	}
	return ""
}

type MessgeReadRequest struct {
	// 分页
	PageNumber int32 `protobuf:"varint,1,opt,name=pageNumber" json:"pageNumber,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=pageSize" json:"pageSize,omitempty"`
	// 查询相关
	MsgType    string `protobuf:"bytes,3,opt,name=msgType" json:"msgType,omitempty"`
	Id         int64  `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	FromUserId int64  `protobuf:"varint,5,opt,name=FromUserId" json:"FromUserId,omitempty"`
	ToUserId   int64  `protobuf:"varint,6,opt,name=toUserId" json:"toUserId,omitempty"`
	Title      string `protobuf:"bytes,7,opt,name=Title" json:"Title,omitempty"`
	Messge     string `protobuf:"bytes,8,opt,name=Messge" json:"Messge,omitempty"`
	Status     string `protobuf:"bytes,9,opt,name=status" json:"status,omitempty"`
	IsDelete   string `protobuf:"bytes,10,opt,name=isDelete" json:"isDelete,omitempty"`
}

func (m *MessgeReadRequest) Reset()                    { *m = MessgeReadRequest{} }
func (m *MessgeReadRequest) String() string            { return proto.CompactTextString(m) }
func (*MessgeReadRequest) ProtoMessage()               {}
func (*MessgeReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *MessgeReadRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *MessgeReadRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *MessgeReadRequest) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MessgeReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MessgeReadRequest) GetFromUserId() int64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func (m *MessgeReadRequest) GetToUserId() int64 {
	if m != nil {
		return m.ToUserId
	}
	return 0
}

func (m *MessgeReadRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MessgeReadRequest) GetMessge() string {
	if m != nil {
		return m.Messge
	}
	return ""
}

func (m *MessgeReadRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *MessgeReadRequest) GetIsDelete() string {
	if m != nil {
		return m.IsDelete
	}
	return ""
}

type MessgeListResponse struct {
	// 数量总数
	PageNumber int64     `protobuf:"varint,1,opt,name=pageNumber" json:"pageNumber,omitempty"`
	PageSize   int64     `protobuf:"varint,2,opt,name=pageSize" json:"pageSize,omitempty"`
	Nums       int64     `protobuf:"varint,3,opt,name=nums" json:"nums,omitempty"`
	MsgList    []*Messge `protobuf:"bytes,4,rep,name=msgList" json:"msgList,omitempty"`
}

func (m *MessgeListResponse) Reset()                    { *m = MessgeListResponse{} }
func (m *MessgeListResponse) String() string            { return proto.CompactTextString(m) }
func (*MessgeListResponse) ProtoMessage()               {}
func (*MessgeListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *MessgeListResponse) GetPageNumber() int64 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *MessgeListResponse) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *MessgeListResponse) GetNums() int64 {
	if m != nil {
		return m.Nums
	}
	return 0
}

func (m *MessgeListResponse) GetMsgList() []*Messge {
	if m != nil {
		return m.MsgList
	}
	return nil
}

type MessgeDeleteRequest struct {
	// 修改id对应的msg.is_delete true
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *MessgeDeleteRequest) Reset()                    { *m = MessgeDeleteRequest{} }
func (m *MessgeDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*MessgeDeleteRequest) ProtoMessage()               {}
func (*MessgeDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *MessgeDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Messge)(nil), "protos.Messge")
	proto.RegisterType((*MessgeResponse)(nil), "protos.MessgeResponse")
	proto.RegisterType((*MessgeReadRequest)(nil), "protos.MessgeReadRequest")
	proto.RegisterType((*MessgeListResponse)(nil), "protos.MessgeListResponse")
	proto.RegisterType((*MessgeDeleteRequest)(nil), "protos.MessgeDeleteRequest")
	proto.RegisterEnum("protos.MessgeStatus", MessgeStatus_name, MessgeStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MessgeService service

type MessgeServiceClient interface {
	MessgeCreate(ctx context.Context, in *Messge, opts ...grpc.CallOption) (*MessgeResponse, error)
	MessgeUpdate(ctx context.Context, in *Messge, opts ...grpc.CallOption) (*Messge, error)
	MessgeRead(ctx context.Context, in *MessgeReadRequest, opts ...grpc.CallOption) (*MessgeListResponse, error)
	MessgeDelete(ctx context.Context, in *MessgeDeleteRequest, opts ...grpc.CallOption) (*MessgeResponse, error)
}

type messgeServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessgeServiceClient(cc *grpc.ClientConn) MessgeServiceClient {
	return &messgeServiceClient{cc}
}

func (c *messgeServiceClient) MessgeCreate(ctx context.Context, in *Messge, opts ...grpc.CallOption) (*MessgeResponse, error) {
	out := new(MessgeResponse)
	err := grpc.Invoke(ctx, "/protos.MessgeService/MessgeCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messgeServiceClient) MessgeUpdate(ctx context.Context, in *Messge, opts ...grpc.CallOption) (*Messge, error) {
	out := new(Messge)
	err := grpc.Invoke(ctx, "/protos.MessgeService/MessgeUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messgeServiceClient) MessgeRead(ctx context.Context, in *MessgeReadRequest, opts ...grpc.CallOption) (*MessgeListResponse, error) {
	out := new(MessgeListResponse)
	err := grpc.Invoke(ctx, "/protos.MessgeService/MessgeRead", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messgeServiceClient) MessgeDelete(ctx context.Context, in *MessgeDeleteRequest, opts ...grpc.CallOption) (*MessgeResponse, error) {
	out := new(MessgeResponse)
	err := grpc.Invoke(ctx, "/protos.MessgeService/MessgeDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessgeService service

type MessgeServiceServer interface {
	MessgeCreate(context.Context, *Messge) (*MessgeResponse, error)
	MessgeUpdate(context.Context, *Messge) (*Messge, error)
	MessgeRead(context.Context, *MessgeReadRequest) (*MessgeListResponse, error)
	MessgeDelete(context.Context, *MessgeDeleteRequest) (*MessgeResponse, error)
}

func RegisterMessgeServiceServer(s *grpc.Server, srv MessgeServiceServer) {
	s.RegisterService(&_MessgeService_serviceDesc, srv)
}

func _MessgeService_MessgeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Messge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessgeServiceServer).MessgeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MessgeService/MessgeCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessgeServiceServer).MessgeCreate(ctx, req.(*Messge))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessgeService_MessgeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Messge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessgeServiceServer).MessgeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MessgeService/MessgeUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessgeServiceServer).MessgeUpdate(ctx, req.(*Messge))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessgeService_MessgeRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessgeReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessgeServiceServer).MessgeRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MessgeService/MessgeRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessgeServiceServer).MessgeRead(ctx, req.(*MessgeReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessgeService_MessgeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessgeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessgeServiceServer).MessgeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MessgeService/MessgeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessgeServiceServer).MessgeDelete(ctx, req.(*MessgeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessgeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.MessgeService",
	HandlerType: (*MessgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessgeCreate",
			Handler:    _MessgeService_MessgeCreate_Handler,
		},
		{
			MethodName: "MessgeUpdate",
			Handler:    _MessgeService_MessgeUpdate_Handler,
		},
		{
			MethodName: "MessgeRead",
			Handler:    _MessgeService_MessgeRead_Handler,
		},
		{
			MethodName: "MessgeDelete",
			Handler:    _MessgeService_MessgeDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/message.proto",
}

func init() { proto.RegisterFile("protos/message.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0x63, 0x3b, 0x76, 0x92, 0xf9, 0xf5, 0x17, 0xa5, 0x43, 0x54, 0x19, 0x83, 0x50, 0x64,
	0x09, 0xc9, 0xe2, 0x10, 0xa0, 0x5c, 0xb8, 0xb6, 0x90, 0x4a, 0x48, 0x21, 0x07, 0xa7, 0x79, 0x00,
	0x37, 0x1e, 0x22, 0x8b, 0xba, 0x36, 0xde, 0x35, 0x12, 0xbc, 0x01, 0x17, 0xae, 0xbc, 0x19, 0xcf,
	0x83, 0x3c, 0xbb, 0xfe, 0x13, 0xd3, 0x72, 0xdb, 0xef, 0xfc, 0xd9, 0x9d, 0xf9, 0xcc, 0x68, 0x61,
	0x9e, 0x17, 0x99, 0xcc, 0xc4, 0xcb, 0x94, 0x84, 0x88, 0x0e, 0xb4, 0x64, 0x89, 0x8e, 0xb2, 0xfa,
	0xbf, 0x0d, 0x70, 0x3e, 0x92, 0x10, 0x07, 0xc2, 0x29, 0x98, 0x49, 0xec, 0x1a, 0x0b, 0x23, 0x98,
	0x84, 0x66, 0x12, 0xe3, 0x33, 0x80, 0x4f, 0x45, 0x96, 0xee, 0x04, 0x15, 0x1f, 0x62, 0xd7, 0x5c,
	0x18, 0x81, 0x15, 0x76, 0x2c, 0xe8, 0xc1, 0x58, 0x66, 0xda, 0x6b, 0xb1, 0xb7, 0xd1, 0xf8, 0x14,
	0x26, 0xfb, 0x82, 0x22, 0x49, 0xf1, 0x85, 0x74, 0x87, 0xec, 0x6c, 0x0d, 0x38, 0x07, 0x5b, 0x26,
	0xf2, 0x96, 0x5c, 0x9b, 0x1f, 0x53, 0x02, 0x5d, 0x18, 0xe9, 0x1a, 0x5d, 0x87, 0xed, 0xb5, 0xc4,
	0x33, 0x70, 0x84, 0x8c, 0x64, 0x29, 0xdc, 0x11, 0x3b, 0xb4, 0xaa, 0x2a, 0x48, 0xc4, 0x7b, 0xba,
	0x25, 0x49, 0xee, 0x78, 0x61, 0x04, 0xe3, 0xb0, 0xd1, 0xbe, 0x0f, 0x53, 0xd5, 0x57, 0x48, 0x22,
	0xcf, 0xee, 0x04, 0xe1, 0x0c, 0xac, 0xb4, 0x69, 0xb0, 0x3a, 0xfa, 0xbf, 0x4c, 0x38, 0xad, 0x83,
	0xa2, 0x38, 0xa4, 0x2f, 0x25, 0x09, 0x59, 0xf5, 0x9d, 0x47, 0x07, 0xda, 0x94, 0xe9, 0x0d, 0x15,
	0x1c, 0x6e, 0x87, 0x1d, 0x4b, 0xf5, 0x6a, 0xa5, 0xb6, 0xc9, 0x77, 0x62, 0x2a, 0x76, 0xd8, 0x68,
	0xee, 0x41, 0x1c, 0xae, 0xbf, 0xe5, 0xc4, 0x48, 0xaa, 0x1e, 0x94, 0xd4, 0x74, 0x15, 0x0a, 0x4d,
	0xf7, 0xaa, 0xa5, 0x6b, 0x2b, 0xba, 0x57, 0xf7, 0xd3, 0x75, 0x7a, 0x74, 0xe7, 0x60, 0x5f, 0x33,
	0x3f, 0x85, 0x43, 0x89, 0x8a, 0x92, 0x6a, 0x86, 0x59, 0x4c, 0xc2, 0x7a, 0xae, 0x2d, 0xbd, 0xc9,
	0x83, 0xf4, 0x80, 0x3d, 0x2d, 0xbd, 0x9f, 0x06, 0xa0, 0x4a, 0x5f, 0x27, 0x42, 0x36, 0x08, 0xff,
	0x46, 0x63, 0xfd, 0x13, 0x8d, 0xd5, 0x41, 0x83, 0x30, 0xbc, 0x2b, 0x53, 0xa1, 0x57, 0x85, 0xcf,
	0x18, 0x30, 0xae, 0xea, 0x09, 0x77, 0xb8, 0xb0, 0x82, 0xff, 0xce, 0xa7, 0x6a, 0x3d, 0xc5, 0x52,
	0x8f, 0xa5, 0x76, 0xfb, 0xcf, 0xe1, 0x91, 0x32, 0xa9, 0x02, 0xeb, 0x59, 0xf5, 0x76, 0xf6, 0xc5,
	0x6b, 0x38, 0x51, 0x61, 0x5b, 0xd5, 0x23, 0x80, 0xb3, 0xdb, 0x6c, 0x57, 0xab, 0xcd, 0x6c, 0x80,
	0x63, 0x18, 0xf2, 0xc9, 0xc0, 0x11, 0x58, 0x17, 0xeb, 0xf5, 0xcc, 0xf0, 0xcc, 0x99, 0x71, 0xfe,
	0xc3, 0x84, 0xff, 0x75, 0x0e, 0x15, 0x5f, 0x93, 0x3d, 0xe1, 0xdb, 0xfa, 0x92, 0x77, 0xbc, 0xb1,
	0xd8, 0x2b, 0xca, 0x3b, 0xeb, 0x15, 0xa9, 0xe9, 0xf8, 0x03, 0x7c, 0x55, 0x67, 0xee, 0xf2, 0xf8,
	0xbe, 0xcc, 0x9e, 0xf6, 0x07, 0xb8, 0x02, 0x68, 0x37, 0x10, 0x1f, 0xf7, 0x6f, 0x6e, 0xb6, 0xd2,
	0xf3, 0x8e, 0x5d, 0xdd, 0xb1, 0xf0, 0x35, 0x27, 0x5d, 0x3c, 0xf8, 0xe4, 0x38, 0xfa, 0x08, 0xda,
	0xc3, 0xf5, 0x5f, 0xfa, 0x70, 0xba, 0xcf, 0xd2, 0x65, 0x94, 0x7f, 0xce, 0xcb, 0x42, 0x7f, 0x15,
	0x97, 0xc7, 0x74, 0x6e, 0xd4, 0xcf, 0xf1, 0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0xd4,
	0x67, 0x5c, 0x58, 0x04, 0x00, 0x00,
}
