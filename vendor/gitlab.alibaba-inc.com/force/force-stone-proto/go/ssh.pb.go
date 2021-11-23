// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssh.proto

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

type SSHUploadPackRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-pack' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
	// Parameters to use with git -c (key=value pairs)
	GitConfigOptions []string `protobuf:"bytes,4,rep,name=git_config_options,json=gitConfigOptions,proto3" json:"git_config_options,omitempty"`
	// Git protocol version
	GitProtocol          string   `protobuf:"bytes,5,opt,name=git_protocol,json=gitProtocol,proto3" json:"git_protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSHUploadPackRequest) Reset()         { *m = SSHUploadPackRequest{} }
func (m *SSHUploadPackRequest) String() string { return proto.CompactTextString(m) }
func (*SSHUploadPackRequest) ProtoMessage()    {}
func (*SSHUploadPackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0eae71e2e883eb, []int{0}
}

func (m *SSHUploadPackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSHUploadPackRequest.Unmarshal(m, b)
}
func (m *SSHUploadPackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSHUploadPackRequest.Marshal(b, m, deterministic)
}
func (m *SSHUploadPackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSHUploadPackRequest.Merge(m, src)
}
func (m *SSHUploadPackRequest) XXX_Size() int {
	return xxx_messageInfo_SSHUploadPackRequest.Size(m)
}
func (m *SSHUploadPackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SSHUploadPackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SSHUploadPackRequest proto.InternalMessageInfo

func (m *SSHUploadPackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHUploadPackRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

func (m *SSHUploadPackRequest) GetGitConfigOptions() []string {
	if m != nil {
		return m.GitConfigOptions
	}
	return nil
}

func (m *SSHUploadPackRequest) GetGitProtocol() string {
	if m != nil {
		return m.GitProtocol
	}
	return ""
}

type SSHUploadPackResponse struct {
	// A chunk of raw data from 'git upload-pack' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git upload-pack' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This field may be nil. This is intentional: only when the remote
	// command has finished can we return its exit status.
	ExitStatus           *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SSHUploadPackResponse) Reset()         { *m = SSHUploadPackResponse{} }
func (m *SSHUploadPackResponse) String() string { return proto.CompactTextString(m) }
func (*SSHUploadPackResponse) ProtoMessage()    {}
func (*SSHUploadPackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0eae71e2e883eb, []int{1}
}

func (m *SSHUploadPackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSHUploadPackResponse.Unmarshal(m, b)
}
func (m *SSHUploadPackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSHUploadPackResponse.Marshal(b, m, deterministic)
}
func (m *SSHUploadPackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSHUploadPackResponse.Merge(m, src)
}
func (m *SSHUploadPackResponse) XXX_Size() int {
	return xxx_messageInfo_SSHUploadPackResponse.Size(m)
}
func (m *SSHUploadPackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SSHUploadPackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SSHUploadPackResponse proto.InternalMessageInfo

func (m *SSHUploadPackResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHUploadPackResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHUploadPackResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

type SSHReceivePackRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-pack' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
	// Contents of GL_ID, GL_REPOSITORY, and GL_USERNAME environment variables
	// for 'git receive-pack'
	GlId         string `protobuf:"bytes,3,opt,name=gl_id,json=glId,proto3" json:"gl_id,omitempty"`
	GlRepository string `protobuf:"bytes,4,opt,name=gl_repository,json=glRepository,proto3" json:"gl_repository,omitempty"`
	GlUsername   string `protobuf:"bytes,5,opt,name=gl_username,json=glUsername,proto3" json:"gl_username,omitempty"`
	// Git protocol version
	GitProtocol string `protobuf:"bytes,6,opt,name=git_protocol,json=gitProtocol,proto3" json:"git_protocol,omitempty"`
	// Parameters to use with git -c (key=value pairs)
	GitConfigOptions     []string `protobuf:"bytes,7,rep,name=git_config_options,json=gitConfigOptions,proto3" json:"git_config_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSHReceivePackRequest) Reset()         { *m = SSHReceivePackRequest{} }
func (m *SSHReceivePackRequest) String() string { return proto.CompactTextString(m) }
func (*SSHReceivePackRequest) ProtoMessage()    {}
func (*SSHReceivePackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0eae71e2e883eb, []int{2}
}

func (m *SSHReceivePackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSHReceivePackRequest.Unmarshal(m, b)
}
func (m *SSHReceivePackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSHReceivePackRequest.Marshal(b, m, deterministic)
}
func (m *SSHReceivePackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSHReceivePackRequest.Merge(m, src)
}
func (m *SSHReceivePackRequest) XXX_Size() int {
	return xxx_messageInfo_SSHReceivePackRequest.Size(m)
}
func (m *SSHReceivePackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SSHReceivePackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SSHReceivePackRequest proto.InternalMessageInfo

func (m *SSHReceivePackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHReceivePackRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

func (m *SSHReceivePackRequest) GetGlId() string {
	if m != nil {
		return m.GlId
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGlRepository() string {
	if m != nil {
		return m.GlRepository
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGlUsername() string {
	if m != nil {
		return m.GlUsername
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGitProtocol() string {
	if m != nil {
		return m.GitProtocol
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGitConfigOptions() []string {
	if m != nil {
		return m.GitConfigOptions
	}
	return nil
}

type SSHReceivePackResponse struct {
	// A chunk of raw data from 'git receive-pack' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git receive-pack' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This field may be nil. This is intentional: only when the remote
	// command has finished can we return its exit status.
	ExitStatus           *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SSHReceivePackResponse) Reset()         { *m = SSHReceivePackResponse{} }
func (m *SSHReceivePackResponse) String() string { return proto.CompactTextString(m) }
func (*SSHReceivePackResponse) ProtoMessage()    {}
func (*SSHReceivePackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0eae71e2e883eb, []int{3}
}

func (m *SSHReceivePackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSHReceivePackResponse.Unmarshal(m, b)
}
func (m *SSHReceivePackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSHReceivePackResponse.Marshal(b, m, deterministic)
}
func (m *SSHReceivePackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSHReceivePackResponse.Merge(m, src)
}
func (m *SSHReceivePackResponse) XXX_Size() int {
	return xxx_messageInfo_SSHReceivePackResponse.Size(m)
}
func (m *SSHReceivePackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SSHReceivePackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SSHReceivePackResponse proto.InternalMessageInfo

func (m *SSHReceivePackResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHReceivePackResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHReceivePackResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*SSHUploadPackRequest)(nil), "gitaly.SSHUploadPackRequest")
	proto.RegisterType((*SSHUploadPackResponse)(nil), "gitaly.SSHUploadPackResponse")
	proto.RegisterType((*SSHReceivePackRequest)(nil), "gitaly.SSHReceivePackRequest")
	proto.RegisterType((*SSHReceivePackResponse)(nil), "gitaly.SSHReceivePackResponse")
}

func init() { proto.RegisterFile("ssh.proto", fileDescriptor_ef0eae71e2e883eb) }

var fileDescriptor_ef0eae71e2e883eb = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0x6b, 0x5a, 0xe8, 0x69, 0x86, 0x26, 0xb3, 0x55, 0x51, 0xc5, 0x4f, 0x09, 0x37, 0xb9,
	0x40, 0x11, 0xea, 0xde, 0x00, 0x84, 0x34, 0xb8, 0x61, 0x72, 0xd4, 0xeb, 0xc8, 0x4d, 0xce, 0x3c,
	0x0b, 0x37, 0x0e, 0xf6, 0xe9, 0xb4, 0x49, 0xf0, 0x08, 0x3c, 0x0f, 0x6f, 0xc1, 0x33, 0x21, 0x9c,
	0x30, 0xd2, 0x75, 0xbd, 0x84, 0xbb, 0x9c, 0xef, 0xb3, 0x8f, 0xcf, 0xf7, 0x7d, 0x27, 0x30, 0x76,
	0xee, 0x32, 0x6b, 0xac, 0x21, 0xc3, 0x46, 0x52, 0x91, 0xd0, 0x37, 0xb3, 0xc8, 0x5d, 0x0a, 0x8b,
	0x55, 0x8b, 0x26, 0x3f, 0x03, 0x38, 0xce, 0xf3, 0xb3, 0x65, 0xa3, 0x8d, 0xa8, 0xce, 0x45, 0xf9,
	0x99, 0xe3, 0x97, 0x0d, 0x3a, 0x62, 0x0b, 0x00, 0x8b, 0x8d, 0x71, 0x8a, 0x8c, 0xbd, 0x89, 0x83,
	0x79, 0x90, 0x4e, 0x16, 0x2c, 0x6b, 0x7b, 0x64, 0xfc, 0x96, 0xe1, 0xbd, 0x53, 0xec, 0x18, 0x86,
	0x8e, 0x2a, 0x55, 0xc7, 0x07, 0xf3, 0x20, 0x8d, 0x78, 0x5b, 0xb0, 0xd7, 0xc0, 0xa4, 0xa2, 0xa2,
	0x34, 0xf5, 0x85, 0x92, 0x85, 0x69, 0x48, 0x99, 0xda, 0xc5, 0xe1, 0x7c, 0x90, 0x8e, 0xf9, 0x91,
	0x54, 0xf4, 0xce, 0x13, 0x9f, 0x5a, 0x9c, 0xbd, 0x84, 0xe8, 0xf7, 0x69, 0x3f, 0x5d, 0x69, 0x74,
	0x3c, 0x9c, 0x07, 0xe9, 0x98, 0x4f, 0xa4, 0xa2, 0xf3, 0x0e, 0xfa, 0x18, 0x3e, 0x1a, 0x1c, 0x85,
	0xfc, 0xa4, 0xd7, 0xb4, 0x11, 0x56, 0xac, 0x91, 0xd0, 0xba, 0xe4, 0x2b, 0x9c, 0xdc, 0xd1, 0xe3,
	0x1a, 0x53, 0x3b, 0x64, 0x53, 0x18, 0x39, 0xaa, 0xcc, 0x86, 0xbc, 0x98, 0x88, 0x77, 0x55, 0x87,
	0xa3, 0xb5, 0xdd, 0xd4, 0x5d, 0xc5, 0x4e, 0x61, 0x82, 0xd7, 0x8a, 0x0a, 0x47, 0x82, 0x36, 0x2e,
	0x1e, 0x6c, 0x3b, 0xf0, 0xfe, 0x5a, 0x51, 0xee, 0x19, 0x0e, 0x78, 0xfb, 0x9d, 0x7c, 0x3f, 0xf0,
	0xcf, 0x73, 0x2c, 0x51, 0x5d, 0xe1, 0xbf, 0xf1, 0xf3, 0x09, 0x0c, 0xa5, 0x2e, 0x54, 0xe5, 0x47,
	0x1a, 0xf3, 0x50, 0xea, 0x0f, 0x15, 0x7b, 0x05, 0x87, 0x52, 0x17, 0xbd, 0x17, 0x42, 0x4f, 0x46,
	0x52, 0xff, 0xed, 0xcd, 0x5e, 0xc0, 0x44, 0xea, 0x62, 0xe3, 0xd0, 0xd6, 0x62, 0x8d, 0x9d, 0xb5,
	0x20, 0xf5, 0xb2, 0x43, 0x76, 0xcc, 0x1f, 0xed, 0x98, 0xbf, 0x27, 0xcd, 0x87, 0xf7, 0xa7, 0x99,
	0x7c, 0x83, 0xe9, 0x5d, 0x3b, 0xfe, 0x63, 0x1c, 0x8b, 0x1f, 0x01, 0x40, 0x9e, 0x9f, 0xe5, 0x68,
	0xaf, 0x54, 0x89, 0x8c, 0xc3, 0xe1, 0xd6, 0x6e, 0xb0, 0xa7, 0x7f, 0xee, 0xdf, 0xf7, 0x0b, 0xcc,
	0x9e, 0xed, 0x61, 0x5b, 0x05, 0xc9, 0x83, 0x34, 0x78, 0x13, 0xb0, 0x25, 0x3c, 0xde, 0x56, 0xc8,
	0xfa, 0xd7, 0x76, 0x17, 0x61, 0xf6, 0x7c, 0x1f, 0xdd, 0x6f, 0xfb, 0x36, 0x86, 0x69, 0x69, 0xd6,
	0x99, 0xd0, 0x6a, 0x25, 0x56, 0x22, 0xbb, 0x30, 0xb6, 0xc4, 0x4c, 0xda, 0xa6, 0x5c, 0x8d, 0x7c,
	0x3a, 0xa7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x51, 0xcf, 0xad, 0xdb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SSHServiceClient is the client API for SSHService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SSHServiceClient interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadPackClient, error)
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHReceivePackClient, error)
}

type sSHServiceClient struct {
	cc *grpc.ClientConn
}

func NewSSHServiceClient(cc *grpc.ClientConn) SSHServiceClient {
	return &sSHServiceClient{cc}
}

func (c *sSHServiceClient) SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadPackClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SSHService_serviceDesc.Streams[0], "/gitaly.SSHService/SSHUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceSSHUploadPackClient{stream}
	return x, nil
}

type SSHService_SSHUploadPackClient interface {
	Send(*SSHUploadPackRequest) error
	Recv() (*SSHUploadPackResponse, error)
	grpc.ClientStream
}

type sSHServiceSSHUploadPackClient struct {
	grpc.ClientStream
}

func (x *sSHServiceSSHUploadPackClient) Send(m *SSHUploadPackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHServiceSSHUploadPackClient) Recv() (*SSHUploadPackResponse, error) {
	m := new(SSHUploadPackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sSHServiceClient) SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHReceivePackClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SSHService_serviceDesc.Streams[1], "/gitaly.SSHService/SSHReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceSSHReceivePackClient{stream}
	return x, nil
}

type SSHService_SSHReceivePackClient interface {
	Send(*SSHReceivePackRequest) error
	Recv() (*SSHReceivePackResponse, error)
	grpc.ClientStream
}

type sSHServiceSSHReceivePackClient struct {
	grpc.ClientStream
}

func (x *sSHServiceSSHReceivePackClient) Send(m *SSHReceivePackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHServiceSSHReceivePackClient) Recv() (*SSHReceivePackResponse, error) {
	m := new(SSHReceivePackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SSHServiceServer is the server API for SSHService service.
type SSHServiceServer interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(SSHService_SSHUploadPackServer) error
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(SSHService_SSHReceivePackServer) error
}

// UnimplementedSSHServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSSHServiceServer struct {
}

func (*UnimplementedSSHServiceServer) SSHUploadPack(srv SSHService_SSHUploadPackServer) error {
	return status.Errorf(codes.Unimplemented, "method SSHUploadPack not implemented")
}
func (*UnimplementedSSHServiceServer) SSHReceivePack(srv SSHService_SSHReceivePackServer) error {
	return status.Errorf(codes.Unimplemented, "method SSHReceivePack not implemented")
}

func RegisterSSHServiceServer(s *grpc.Server, srv SSHServiceServer) {
	s.RegisterService(&_SSHService_serviceDesc, srv)
}

func _SSHService_SSHUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServiceServer).SSHUploadPack(&sSHServiceSSHUploadPackServer{stream})
}

type SSHService_SSHUploadPackServer interface {
	Send(*SSHUploadPackResponse) error
	Recv() (*SSHUploadPackRequest, error)
	grpc.ServerStream
}

type sSHServiceSSHUploadPackServer struct {
	grpc.ServerStream
}

func (x *sSHServiceSSHUploadPackServer) Send(m *SSHUploadPackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHServiceSSHUploadPackServer) Recv() (*SSHUploadPackRequest, error) {
	m := new(SSHUploadPackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SSHService_SSHReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServiceServer).SSHReceivePack(&sSHServiceSSHReceivePackServer{stream})
}

type SSHService_SSHReceivePackServer interface {
	Send(*SSHReceivePackResponse) error
	Recv() (*SSHReceivePackRequest, error)
	grpc.ServerStream
}

type sSHServiceSSHReceivePackServer struct {
	grpc.ServerStream
}

func (x *sSHServiceSSHReceivePackServer) Send(m *SSHReceivePackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHServiceSSHReceivePackServer) Recv() (*SSHReceivePackRequest, error) {
	m := new(SSHReceivePackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SSHService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.SSHService",
	HandlerType: (*SSHServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SSHUploadPack",
			Handler:       _SSHService_SSHUploadPack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SSHReceivePack",
			Handler:       _SSHService_SSHReceivePack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ssh.proto",
}
