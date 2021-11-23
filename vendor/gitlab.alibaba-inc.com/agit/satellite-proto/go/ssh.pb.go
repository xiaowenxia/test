// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssh.proto

package satellite

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
	GitConfigOptions     []string `protobuf:"bytes,4,rep,name=git_config_options,json=gitConfigOptions,proto3" json:"git_config_options,omitempty"`
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
	GlId                 string   `protobuf:"bytes,3,opt,name=gl_id,json=glId,proto3" json:"gl_id,omitempty"`
	GlRepository         string   `protobuf:"bytes,4,opt,name=gl_repository,json=glRepository,proto3" json:"gl_repository,omitempty"`
	GlUsername           string   `protobuf:"bytes,5,opt,name=gl_username,json=glUsername,proto3" json:"gl_username,omitempty"`
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

type SSHAdvertiseRequest struct {
	Repository           *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	GlId                 string      `protobuf:"bytes,3,opt,name=gl_id,json=glId,proto3" json:"gl_id,omitempty"`
	GlRepository         string      `protobuf:"bytes,4,opt,name=gl_repository,json=glRepository,proto3" json:"gl_repository,omitempty"`
	GlUsername           string      `protobuf:"bytes,5,opt,name=gl_username,json=glUsername,proto3" json:"gl_username,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SSHAdvertiseRequest) Reset()         { *m = SSHAdvertiseRequest{} }
func (m *SSHAdvertiseRequest) String() string { return proto.CompactTextString(m) }
func (*SSHAdvertiseRequest) ProtoMessage()    {}
func (*SSHAdvertiseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0eae71e2e883eb, []int{4}
}

func (m *SSHAdvertiseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSHAdvertiseRequest.Unmarshal(m, b)
}
func (m *SSHAdvertiseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSHAdvertiseRequest.Marshal(b, m, deterministic)
}
func (m *SSHAdvertiseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSHAdvertiseRequest.Merge(m, src)
}
func (m *SSHAdvertiseRequest) XXX_Size() int {
	return xxx_messageInfo_SSHAdvertiseRequest.Size(m)
}
func (m *SSHAdvertiseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SSHAdvertiseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SSHAdvertiseRequest proto.InternalMessageInfo

func (m *SSHAdvertiseRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHAdvertiseRequest) GetGlId() string {
	if m != nil {
		return m.GlId
	}
	return ""
}

func (m *SSHAdvertiseRequest) GetGlRepository() string {
	if m != nil {
		return m.GlRepository
	}
	return ""
}

func (m *SSHAdvertiseRequest) GetGlUsername() string {
	if m != nil {
		return m.GlUsername
	}
	return ""
}

type SSHAdvertiseResponse struct {
	Stdout               []byte      `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr               []byte      `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	ExitStatus           *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SSHAdvertiseResponse) Reset()         { *m = SSHAdvertiseResponse{} }
func (m *SSHAdvertiseResponse) String() string { return proto.CompactTextString(m) }
func (*SSHAdvertiseResponse) ProtoMessage()    {}
func (*SSHAdvertiseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef0eae71e2e883eb, []int{5}
}

func (m *SSHAdvertiseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSHAdvertiseResponse.Unmarshal(m, b)
}
func (m *SSHAdvertiseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSHAdvertiseResponse.Marshal(b, m, deterministic)
}
func (m *SSHAdvertiseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSHAdvertiseResponse.Merge(m, src)
}
func (m *SSHAdvertiseResponse) XXX_Size() int {
	return xxx_messageInfo_SSHAdvertiseResponse.Size(m)
}
func (m *SSHAdvertiseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SSHAdvertiseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SSHAdvertiseResponse proto.InternalMessageInfo

func (m *SSHAdvertiseResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHAdvertiseResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHAdvertiseResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*SSHUploadPackRequest)(nil), "satellite.SSHUploadPackRequest")
	proto.RegisterType((*SSHUploadPackResponse)(nil), "satellite.SSHUploadPackResponse")
	proto.RegisterType((*SSHReceivePackRequest)(nil), "satellite.SSHReceivePackRequest")
	proto.RegisterType((*SSHReceivePackResponse)(nil), "satellite.SSHReceivePackResponse")
	proto.RegisterType((*SSHAdvertiseRequest)(nil), "satellite.SSHAdvertiseRequest")
	proto.RegisterType((*SSHAdvertiseResponse)(nil), "satellite.SSHAdvertiseResponse")
}

func init() { proto.RegisterFile("ssh.proto", fileDescriptor_ef0eae71e2e883eb) }

var fileDescriptor_ef0eae71e2e883eb = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x65, 0x1b, 0xa7, 0x22, 0x93, 0x14, 0x45, 0xdb, 0xa6, 0x58, 0x39, 0x90, 0x60, 0x2e, 0x39,
	0x20, 0xab, 0x2a, 0x82, 0x3b, 0x20, 0xa4, 0x94, 0x0b, 0x68, 0xad, 0x0a, 0x24, 0x0e, 0xd6, 0xc6,
	0x9e, 0x6e, 0x57, 0x6c, 0xbc, 0x66, 0x77, 0x12, 0x95, 0x0b, 0xf0, 0x21, 0xdc, 0xb9, 0xf1, 0x0b,
	0xfc, 0x1a, 0x92, 0x63, 0x82, 0x93, 0x52, 0x2e, 0x55, 0x7b, 0xf3, 0xcc, 0x9b, 0x79, 0x7e, 0xf3,
	0x34, 0xb3, 0xd0, 0xf1, 0xfe, 0x3c, 0x2e, 0x9d, 0x25, 0xcb, 0x3b, 0x5e, 0x12, 0x1a, 0xa3, 0x09,
	0x87, 0x3d, 0x7f, 0x2e, 0x1d, 0xe6, 0x2b, 0x20, 0xfa, 0xc9, 0xe0, 0x20, 0x49, 0xa6, 0xa7, 0xa5,
	0xb1, 0x32, 0x7f, 0x2b, 0xb3, 0x8f, 0x02, 0x3f, 0x2d, 0xd0, 0x13, 0x7f, 0x0a, 0xe0, 0xb0, 0xb4,
	0x5e, 0x93, 0x75, 0x9f, 0x43, 0x36, 0x66, 0x93, 0xee, 0xf1, 0x20, 0x5e, 0xd3, 0xc4, 0x62, 0x0d,
	0x8a, 0x46, 0x21, 0x3f, 0x80, 0xb6, 0xa7, 0x5c, 0x17, 0xe1, 0xce, 0x98, 0x4d, 0x7a, 0x62, 0x15,
	0xf0, 0xc7, 0xc0, 0x95, 0xa6, 0x34, 0xb3, 0xc5, 0x99, 0x56, 0xa9, 0x2d, 0x49, 0xdb, 0xc2, 0x87,
	0xc1, 0xb8, 0x35, 0xe9, 0x88, 0xbe, 0xd2, 0xf4, 0xb2, 0x02, 0xde, 0xac, 0xf2, 0xaf, 0x83, 0xbb,
	0xad, 0x7e, 0x20, 0x06, 0x8d, 0x8e, 0x52, 0x3a, 0x39, 0x47, 0x42, 0xe7, 0xa3, 0xaf, 0x30, 0xd8,
	0xd2, 0xeb, 0x4b, 0x5b, 0x78, 0xe4, 0x87, 0xb0, 0xeb, 0x29, 0xb7, 0x0b, 0xaa, 0xc4, 0xf6, 0x44,
	0x1d, 0xd5, 0x79, 0x74, 0xae, 0x96, 0x54, 0x47, 0xfc, 0x19, 0x74, 0xf1, 0x42, 0x53, 0xea, 0x49,
	0xd2, 0xc2, 0x87, 0xad, 0x4b, 0x13, 0xbe, 0xba, 0xd0, 0x94, 0x54, 0xa0, 0x00, 0x5c, 0x7f, 0x47,
	0xbf, 0x58, 0xa5, 0x40, 0x60, 0x86, 0x7a, 0x89, 0x37, 0x66, 0xd9, 0x3e, 0xb4, 0x95, 0x49, 0x75,
	0x5e, 0x09, 0xeb, 0x88, 0x40, 0x99, 0x93, 0x9c, 0x3f, 0x82, 0x3d, 0x65, 0xd2, 0xc6, 0x4f, 0x82,
	0x0a, 0xec, 0x29, 0xf3, 0x97, 0x9b, 0x8f, 0xa0, 0xab, 0x4c, 0xba, 0xf0, 0xe8, 0x0a, 0x39, 0xc7,
	0xb0, 0x5d, 0x95, 0x80, 0x32, 0xa7, 0x75, 0x26, 0xfa, 0xc6, 0xe0, 0x70, 0x7b, 0x82, 0x5b, 0x36,
	0xf1, 0x07, 0x83, 0xfd, 0x24, 0x99, 0x3e, 0xcf, 0x97, 0xe8, 0x48, 0x7b, 0xbc, 0xa6, 0x85, 0x37,
	0x68, 0xd6, 0x97, 0xea, 0x3e, 0x1a, 0x42, 0x6f, 0xd7, 0xa9, 0xe3, 0xef, 0x3b, 0x00, 0x49, 0x32,
	0x4d, 0xd0, 0x2d, 0x75, 0x86, 0xfc, 0x3d, 0xec, 0x6d, 0xac, 0x3f, 0x1f, 0x35, 0x28, 0xfe, 0x75,
	0xc8, 0xc3, 0xf1, 0xd5, 0x05, 0xab, 0x51, 0xa2, 0x3b, 0x13, 0x76, 0xc4, 0xf8, 0x07, 0xb8, 0xb7,
	0xb9, 0x14, 0x7c, 0xab, 0xf3, 0xf2, 0xc6, 0x0f, 0x1f, 0xfe, 0xa7, 0x62, 0x83, 0xfc, 0x1d, 0xf4,
	0x9b, 0x2e, 0x9e, 0x14, 0x67, 0x96, 0x3f, 0xd8, 0x6c, 0xde, 0xde, 0x85, 0xe1, 0xe8, 0x4a, 0xfc,
	0x0f, 0xf5, 0x11, 0x7b, 0x71, 0x1f, 0x06, 0x99, 0x9d, 0xc7, 0xd2, 0xe8, 0x99, 0x9c, 0xc9, 0x58,
	0x2a, 0x4d, 0xb1, 0x72, 0x65, 0x36, 0xdb, 0xad, 0xde, 0xb7, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x98, 0x21, 0xb4, 0x93, 0x05, 0x05, 0x00, 0x00,
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
	// Get repository advertise reference information
	SSHAdvertiseInfo(ctx context.Context, in *SSHAdvertiseRequest, opts ...grpc.CallOption) (SSHService_SSHAdvertiseInfoClient, error)
}

type sSHServiceClient struct {
	cc *grpc.ClientConn
}

func NewSSHServiceClient(cc *grpc.ClientConn) SSHServiceClient {
	return &sSHServiceClient{cc}
}

func (c *sSHServiceClient) SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadPackClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SSHService_serviceDesc.Streams[0], "/satellite.SSHService/SSHUploadPack", opts...)
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
	stream, err := c.cc.NewStream(ctx, &_SSHService_serviceDesc.Streams[1], "/satellite.SSHService/SSHReceivePack", opts...)
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

func (c *sSHServiceClient) SSHAdvertiseInfo(ctx context.Context, in *SSHAdvertiseRequest, opts ...grpc.CallOption) (SSHService_SSHAdvertiseInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SSHService_serviceDesc.Streams[2], "/satellite.SSHService/SSHAdvertiseInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceSSHAdvertiseInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SSHService_SSHAdvertiseInfoClient interface {
	Recv() (*SSHAdvertiseResponse, error)
	grpc.ClientStream
}

type sSHServiceSSHAdvertiseInfoClient struct {
	grpc.ClientStream
}

func (x *sSHServiceSSHAdvertiseInfoClient) Recv() (*SSHAdvertiseResponse, error) {
	m := new(SSHAdvertiseResponse)
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
	// Get repository advertise reference information
	SSHAdvertiseInfo(*SSHAdvertiseRequest, SSHService_SSHAdvertiseInfoServer) error
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
func (*UnimplementedSSHServiceServer) SSHAdvertiseInfo(req *SSHAdvertiseRequest, srv SSHService_SSHAdvertiseInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method SSHAdvertiseInfo not implemented")
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

func _SSHService_SSHAdvertiseInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SSHAdvertiseRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SSHServiceServer).SSHAdvertiseInfo(m, &sSHServiceSSHAdvertiseInfoServer{stream})
}

type SSHService_SSHAdvertiseInfoServer interface {
	Send(*SSHAdvertiseResponse) error
	grpc.ServerStream
}

type sSHServiceSSHAdvertiseInfoServer struct {
	grpc.ServerStream
}

func (x *sSHServiceSSHAdvertiseInfoServer) Send(m *SSHAdvertiseResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SSHService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "satellite.SSHService",
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
		{
			StreamName:    "SSHAdvertiseInfo",
			Handler:       _SSHService_SSHAdvertiseInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ssh.proto",
}