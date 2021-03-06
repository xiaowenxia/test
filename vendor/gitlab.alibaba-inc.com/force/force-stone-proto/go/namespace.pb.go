// Code generated by protoc-gen-go. DO NOT EDIT.
// source: namespace.proto

package gitaly

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

type CreateNamespaceRequest struct {
	Namespace            *Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateNamespaceRequest) Reset()         { *m = CreateNamespaceRequest{} }
func (m *CreateNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNamespaceRequest) ProtoMessage()    {}
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{0}
}

func (m *CreateNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNamespaceRequest.Unmarshal(m, b)
}
func (m *CreateNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *CreateNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNamespaceRequest.Merge(m, src)
}
func (m *CreateNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNamespaceRequest.Size(m)
}
func (m *CreateNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNamespaceRequest proto.InternalMessageInfo

func (m *CreateNamespaceRequest) GetNamespace() *Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

type CreateNamespaceResponse struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNamespaceResponse) Reset()         { *m = CreateNamespaceResponse{} }
func (m *CreateNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateNamespaceResponse) ProtoMessage()    {}
func (*CreateNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{1}
}

func (m *CreateNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNamespaceResponse.Unmarshal(m, b)
}
func (m *CreateNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *CreateNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNamespaceResponse.Merge(m, src)
}
func (m *CreateNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateNamespaceResponse.Size(m)
}
func (m *CreateNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNamespaceResponse proto.InternalMessageInfo

func (m *CreateNamespaceResponse) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type RemoveNamespaceRequest struct {
	Namespace            *Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RemoveNamespaceRequest) Reset()         { *m = RemoveNamespaceRequest{} }
func (m *RemoveNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveNamespaceRequest) ProtoMessage()    {}
func (*RemoveNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{2}
}

func (m *RemoveNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveNamespaceRequest.Unmarshal(m, b)
}
func (m *RemoveNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *RemoveNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveNamespaceRequest.Merge(m, src)
}
func (m *RemoveNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveNamespaceRequest.Size(m)
}
func (m *RemoveNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveNamespaceRequest proto.InternalMessageInfo

func (m *RemoveNamespaceRequest) GetNamespace() *Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

type RemoveNamespaceResponse struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveNamespaceResponse) Reset()         { *m = RemoveNamespaceResponse{} }
func (m *RemoveNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveNamespaceResponse) ProtoMessage()    {}
func (*RemoveNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{3}
}

func (m *RemoveNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveNamespaceResponse.Unmarshal(m, b)
}
func (m *RemoveNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *RemoveNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveNamespaceResponse.Merge(m, src)
}
func (m *RemoveNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveNamespaceResponse.Size(m)
}
func (m *RemoveNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveNamespaceResponse proto.InternalMessageInfo

func (m *RemoveNamespaceResponse) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type MoveNamespaceRequest struct {
	OldNamespace         *Namespace `protobuf:"bytes,1,opt,name=oldNamespace,proto3" json:"oldNamespace,omitempty"`
	NewNamespace         *Namespace `protobuf:"bytes,2,opt,name=newNamespace,proto3" json:"newNamespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MoveNamespaceRequest) Reset()         { *m = MoveNamespaceRequest{} }
func (m *MoveNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*MoveNamespaceRequest) ProtoMessage()    {}
func (*MoveNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{4}
}

func (m *MoveNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveNamespaceRequest.Unmarshal(m, b)
}
func (m *MoveNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *MoveNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveNamespaceRequest.Merge(m, src)
}
func (m *MoveNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_MoveNamespaceRequest.Size(m)
}
func (m *MoveNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MoveNamespaceRequest proto.InternalMessageInfo

func (m *MoveNamespaceRequest) GetOldNamespace() *Namespace {
	if m != nil {
		return m.OldNamespace
	}
	return nil
}

func (m *MoveNamespaceRequest) GetNewNamespace() *Namespace {
	if m != nil {
		return m.NewNamespace
	}
	return nil
}

type MoveNamespaceResponse struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveNamespaceResponse) Reset()         { *m = MoveNamespaceResponse{} }
func (m *MoveNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*MoveNamespaceResponse) ProtoMessage()    {}
func (*MoveNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{5}
}

func (m *MoveNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveNamespaceResponse.Unmarshal(m, b)
}
func (m *MoveNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *MoveNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveNamespaceResponse.Merge(m, src)
}
func (m *MoveNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_MoveNamespaceResponse.Size(m)
}
func (m *MoveNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MoveNamespaceResponse proto.InternalMessageInfo

func (m *MoveNamespaceResponse) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type NamespaceExistsRequest struct {
	Namespace            *Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NamespaceExistsRequest) Reset()         { *m = NamespaceExistsRequest{} }
func (m *NamespaceExistsRequest) String() string { return proto.CompactTextString(m) }
func (*NamespaceExistsRequest) ProtoMessage()    {}
func (*NamespaceExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{6}
}

func (m *NamespaceExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespaceExistsRequest.Unmarshal(m, b)
}
func (m *NamespaceExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespaceExistsRequest.Marshal(b, m, deterministic)
}
func (m *NamespaceExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceExistsRequest.Merge(m, src)
}
func (m *NamespaceExistsRequest) XXX_Size() int {
	return xxx_messageInfo_NamespaceExistsRequest.Size(m)
}
func (m *NamespaceExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceExistsRequest proto.InternalMessageInfo

func (m *NamespaceExistsRequest) GetNamespace() *Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

type NamespaceExistsResponse struct {
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamespaceExistsResponse) Reset()         { *m = NamespaceExistsResponse{} }
func (m *NamespaceExistsResponse) String() string { return proto.CompactTextString(m) }
func (*NamespaceExistsResponse) ProtoMessage()    {}
func (*NamespaceExistsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{7}
}

func (m *NamespaceExistsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespaceExistsResponse.Unmarshal(m, b)
}
func (m *NamespaceExistsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespaceExistsResponse.Marshal(b, m, deterministic)
}
func (m *NamespaceExistsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceExistsResponse.Merge(m, src)
}
func (m *NamespaceExistsResponse) XXX_Size() int {
	return xxx_messageInfo_NamespaceExistsResponse.Size(m)
}
func (m *NamespaceExistsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceExistsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceExistsResponse proto.InternalMessageInfo

func (m *NamespaceExistsResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func init() {
	proto.RegisterType((*CreateNamespaceRequest)(nil), "gitaly.CreateNamespaceRequest")
	proto.RegisterType((*CreateNamespaceResponse)(nil), "gitaly.CreateNamespaceResponse")
	proto.RegisterType((*RemoveNamespaceRequest)(nil), "gitaly.RemoveNamespaceRequest")
	proto.RegisterType((*RemoveNamespaceResponse)(nil), "gitaly.RemoveNamespaceResponse")
	proto.RegisterType((*MoveNamespaceRequest)(nil), "gitaly.MoveNamespaceRequest")
	proto.RegisterType((*MoveNamespaceResponse)(nil), "gitaly.MoveNamespaceResponse")
	proto.RegisterType((*NamespaceExistsRequest)(nil), "gitaly.NamespaceExistsRequest")
	proto.RegisterType((*NamespaceExistsResponse)(nil), "gitaly.NamespaceExistsResponse")
}

func init() { proto.RegisterFile("namespace.proto", fileDescriptor_ecb1e126f615f5dd) }

var fileDescriptor_ecb1e126f615f5dd = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xdb, 0x4a, 0x7f, 0xf5, 0x33, 0x14, 0x15, 0xac, 0x92, 0x56, 0x15, 0x37, 0x79, 0xc5,
	0x06, 0x57, 0x14, 0xf1, 0x02, 0x20, 0x16, 0x5d, 0x34, 0x8b, 0xc0, 0x0b, 0x38, 0xe9, 0x50, 0x22,
	0x25, 0x71, 0xb0, 0xd3, 0x00, 0x7b, 0x1e, 0x95, 0x07, 0x41, 0x69, 0x2e, 0x6e, 0x73, 0x01, 0x81,
	0x58, 0xda, 0x3e, 0xf3, 0x9d, 0xa3, 0x99, 0x31, 0xf4, 0x03, 0xee, 0xa3, 0x0a, 0xb9, 0x83, 0x2c,
	0x94, 0x22, 0x12, 0xa4, 0xbb, 0x74, 0x23, 0xee, 0xbd, 0x8d, 0x7b, 0xea, 0x89, 0x4b, 0x5c, 0xa4,
	0xb7, 0x74, 0x06, 0xc6, 0xad, 0x44, 0x1e, 0xa1, 0x99, 0xcb, 0x2d, 0x7c, 0x5e, 0xa1, 0x8a, 0xc8,
	0x04, 0x76, 0x0a, 0xc4, 0xa8, 0x7d, 0xd6, 0x3e, 0xdf, 0x9d, 0x1e, 0xb0, 0x94, 0xc1, 0xb4, 0x58,
	0x6b, 0xe8, 0x04, 0x86, 0x15, 0x94, 0x0a, 0x45, 0xa0, 0x90, 0x0c, 0xe0, 0x5f, 0xcc, 0xbd, 0x55,
	0xca, 0xf9, 0x6f, 0xa5, 0x87, 0xc4, 0xdb, 0x42, 0x5f, 0xc4, 0x7f, 0xe3, 0x5d, 0x41, 0x7d, 0xe9,
	0xfd, 0xde, 0x86, 0xc1, 0xbc, 0xce, 0xfa, 0x1a, 0x7a, 0xc2, 0x5b, 0x98, 0xdf, 0xbb, 0x6f, 0xc9,
	0x92, 0xb2, 0x00, 0x5f, 0x74, 0x59, 0xa7, 0xb1, 0x6c, 0x53, 0x46, 0x2f, 0xe0, 0x70, 0xfe, 0x83,
	0xd4, 0x33, 0x30, 0x0a, 0xe9, 0xdd, 0xab, 0xab, 0x22, 0xf5, 0xeb, 0x8e, 0x5d, 0xc2, 0xb0, 0x82,
	0xca, 0xbc, 0x0d, 0xe8, 0xe2, 0xfa, 0x26, 0x33, 0xcf, 0x4e, 0xd3, 0x8f, 0x0e, 0xec, 0x17, 0x35,
	0xf7, 0x28, 0x63, 0xd7, 0x41, 0xf2, 0x00, 0xfd, 0xd2, 0xd4, 0xc9, 0x49, 0x6e, 0x5c, 0xbf, 0x59,
	0xe3, 0xd3, 0xc6, 0xf7, 0x34, 0x00, 0x6d, 0x25, 0xd4, 0xd2, 0x3c, 0x35, 0xb5, 0x7e, 0x67, 0x34,
	0xb5, 0x61, 0x11, 0x68, 0x8b, 0x98, 0xb0, 0xb7, 0xd5, 0x6d, 0x72, 0x94, 0xd7, 0xd4, 0xad, 0xc2,
	0xf8, 0xb8, 0xe1, 0x75, 0x33, 0x65, 0xa9, 0x87, 0x3a, 0x65, 0xfd, 0x9c, 0x74, 0xca, 0x86, 0xe6,
	0xd3, 0xd6, 0xcd, 0x08, 0x0c, 0x47, 0xf8, 0x8c, 0x7b, 0xae, 0xcd, 0x6d, 0xce, 0x1e, 0x85, 0x74,
	0x90, 0x2d, 0x65, 0xe8, 0xd8, 0xdd, 0xf5, 0x9f, 0xbd, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0xf8, 0xda, 0xc0, 0xdc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NamespaceServiceClient is the client API for NamespaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NamespaceServiceClient interface {
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error)
	RemoveNamespace(ctx context.Context, in *RemoveNamespaceRequest, opts ...grpc.CallOption) (*RemoveNamespaceResponse, error)
	MoveNamespace(ctx context.Context, in *MoveNamespaceRequest, opts ...grpc.CallOption) (*MoveNamespaceResponse, error)
	NamespaceExists(ctx context.Context, in *NamespaceExistsRequest, opts ...grpc.CallOption) (*NamespaceExistsResponse, error)
}

type namespaceServiceClient struct {
	cc *grpc.ClientConn
}

func NewNamespaceServiceClient(cc *grpc.ClientConn) NamespaceServiceClient {
	return &namespaceServiceClient{cc}
}

func (c *namespaceServiceClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error) {
	out := new(CreateNamespaceResponse)
	err := c.cc.Invoke(ctx, "/gitaly.NamespaceService/CreateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) RemoveNamespace(ctx context.Context, in *RemoveNamespaceRequest, opts ...grpc.CallOption) (*RemoveNamespaceResponse, error) {
	out := new(RemoveNamespaceResponse)
	err := c.cc.Invoke(ctx, "/gitaly.NamespaceService/RemoveNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) MoveNamespace(ctx context.Context, in *MoveNamespaceRequest, opts ...grpc.CallOption) (*MoveNamespaceResponse, error) {
	out := new(MoveNamespaceResponse)
	err := c.cc.Invoke(ctx, "/gitaly.NamespaceService/MoveNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) NamespaceExists(ctx context.Context, in *NamespaceExistsRequest, opts ...grpc.CallOption) (*NamespaceExistsResponse, error) {
	out := new(NamespaceExistsResponse)
	err := c.cc.Invoke(ctx, "/gitaly.NamespaceService/NamespaceExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespaceServiceServer is the server API for NamespaceService service.
type NamespaceServiceServer interface {
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error)
	RemoveNamespace(context.Context, *RemoveNamespaceRequest) (*RemoveNamespaceResponse, error)
	MoveNamespace(context.Context, *MoveNamespaceRequest) (*MoveNamespaceResponse, error)
	NamespaceExists(context.Context, *NamespaceExistsRequest) (*NamespaceExistsResponse, error)
}

// UnimplementedNamespaceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNamespaceServiceServer struct {
}

func (*UnimplementedNamespaceServiceServer) CreateNamespace(ctx context.Context, req *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (*UnimplementedNamespaceServiceServer) RemoveNamespace(ctx context.Context, req *RemoveNamespaceRequest) (*RemoveNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNamespace not implemented")
}
func (*UnimplementedNamespaceServiceServer) MoveNamespace(ctx context.Context, req *MoveNamespaceRequest) (*MoveNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveNamespace not implemented")
}
func (*UnimplementedNamespaceServiceServer) NamespaceExists(ctx context.Context, req *NamespaceExistsRequest) (*NamespaceExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NamespaceExists not implemented")
}

func RegisterNamespaceServiceServer(s *grpc.Server, srv NamespaceServiceServer) {
	s.RegisterService(&_NamespaceService_serviceDesc, srv)
}

func _NamespaceService_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.NamespaceService/CreateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_RemoveNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).RemoveNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.NamespaceService/RemoveNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).RemoveNamespace(ctx, req.(*RemoveNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_MoveNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).MoveNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.NamespaceService/MoveNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).MoveNamespace(ctx, req.(*MoveNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_NamespaceExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespaceExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).NamespaceExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.NamespaceService/NamespaceExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).NamespaceExists(ctx, req.(*NamespaceExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NamespaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.NamespaceService",
	HandlerType: (*NamespaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNamespace",
			Handler:    _NamespaceService_CreateNamespace_Handler,
		},
		{
			MethodName: "RemoveNamespace",
			Handler:    _NamespaceService_RemoveNamespace_Handler,
		},
		{
			MethodName: "MoveNamespace",
			Handler:    _NamespaceService_MoveNamespace_Handler,
		},
		{
			MethodName: "NamespaceExists",
			Handler:    _NamespaceService_NamespaceExists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namespace.proto",
}
