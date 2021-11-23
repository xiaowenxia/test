// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blob.proto

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

type GetBlobRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// Object ID (SHA1) of the blob we want to get
	Oid string `protobuf:"bytes,2,opt,name=oid,proto3" json:"oid,omitempty"`
	// Maximum number of bytes we want to receive. Use '-1' to get the full blob no matter how big.
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlobRequest) Reset()         { *m = GetBlobRequest{} }
func (m *GetBlobRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlobRequest) ProtoMessage()    {}
func (*GetBlobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6903d1e8a20272e8, []int{0}
}

func (m *GetBlobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlobRequest.Unmarshal(m, b)
}
func (m *GetBlobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlobRequest.Marshal(b, m, deterministic)
}
func (m *GetBlobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlobRequest.Merge(m, src)
}
func (m *GetBlobRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlobRequest.Size(m)
}
func (m *GetBlobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlobRequest proto.InternalMessageInfo

func (m *GetBlobRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetBlobRequest) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func (m *GetBlobRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetBlobResponse struct {
	// Blob size; present only in first response message
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Chunk of blob data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Object ID of the actual blob returned. Empty if no blob was found.
	Oid                  string   `protobuf:"bytes,3,opt,name=oid,proto3" json:"oid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlobResponse) Reset()         { *m = GetBlobResponse{} }
func (m *GetBlobResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlobResponse) ProtoMessage()    {}
func (*GetBlobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6903d1e8a20272e8, []int{1}
}

func (m *GetBlobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlobResponse.Unmarshal(m, b)
}
func (m *GetBlobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlobResponse.Marshal(b, m, deterministic)
}
func (m *GetBlobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlobResponse.Merge(m, src)
}
func (m *GetBlobResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlobResponse.Size(m)
}
func (m *GetBlobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlobResponse proto.InternalMessageInfo

func (m *GetBlobResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetBlobResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetBlobResponse) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

type IsBinaryRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// Object ID (SHA1) of the blob
	Oid                  string   `protobuf:"bytes,2,opt,name=oid,proto3" json:"oid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsBinaryRequest) Reset()         { *m = IsBinaryRequest{} }
func (m *IsBinaryRequest) String() string { return proto.CompactTextString(m) }
func (*IsBinaryRequest) ProtoMessage()    {}
func (*IsBinaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6903d1e8a20272e8, []int{2}
}

func (m *IsBinaryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsBinaryRequest.Unmarshal(m, b)
}
func (m *IsBinaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsBinaryRequest.Marshal(b, m, deterministic)
}
func (m *IsBinaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsBinaryRequest.Merge(m, src)
}
func (m *IsBinaryRequest) XXX_Size() int {
	return xxx_messageInfo_IsBinaryRequest.Size(m)
}
func (m *IsBinaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsBinaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsBinaryRequest proto.InternalMessageInfo

func (m *IsBinaryRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *IsBinaryRequest) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

type IsBinaryResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsBinaryResponse) Reset()         { *m = IsBinaryResponse{} }
func (m *IsBinaryResponse) String() string { return proto.CompactTextString(m) }
func (*IsBinaryResponse) ProtoMessage()    {}
func (*IsBinaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6903d1e8a20272e8, []int{3}
}

func (m *IsBinaryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsBinaryResponse.Unmarshal(m, b)
}
func (m *IsBinaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsBinaryResponse.Marshal(b, m, deterministic)
}
func (m *IsBinaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsBinaryResponse.Merge(m, src)
}
func (m *IsBinaryResponse) XXX_Size() int {
	return xxx_messageInfo_IsBinaryResponse.Size(m)
}
func (m *IsBinaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsBinaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsBinaryResponse proto.InternalMessageInfo

func (m *IsBinaryResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type GetSizeRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// Object ID (SHA1) of the blob
	Oid                  string   `protobuf:"bytes,2,opt,name=oid,proto3" json:"oid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSizeRequest) Reset()         { *m = GetSizeRequest{} }
func (m *GetSizeRequest) String() string { return proto.CompactTextString(m) }
func (*GetSizeRequest) ProtoMessage()    {}
func (*GetSizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6903d1e8a20272e8, []int{4}
}

func (m *GetSizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSizeRequest.Unmarshal(m, b)
}
func (m *GetSizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSizeRequest.Marshal(b, m, deterministic)
}
func (m *GetSizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSizeRequest.Merge(m, src)
}
func (m *GetSizeRequest) XXX_Size() int {
	return xxx_messageInfo_GetSizeRequest.Size(m)
}
func (m *GetSizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSizeRequest proto.InternalMessageInfo

func (m *GetSizeRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetSizeRequest) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

type GetSizeResponse struct {
	Size                 int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSizeResponse) Reset()         { *m = GetSizeResponse{} }
func (m *GetSizeResponse) String() string { return proto.CompactTextString(m) }
func (*GetSizeResponse) ProtoMessage()    {}
func (*GetSizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6903d1e8a20272e8, []int{5}
}

func (m *GetSizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSizeResponse.Unmarshal(m, b)
}
func (m *GetSizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSizeResponse.Marshal(b, m, deterministic)
}
func (m *GetSizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSizeResponse.Merge(m, src)
}
func (m *GetSizeResponse) XXX_Size() int {
	return xxx_messageInfo_GetSizeResponse.Size(m)
}
func (m *GetSizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSizeResponse proto.InternalMessageInfo

func (m *GetSizeResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type GetBlobsRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// Object IDs (SHA1) of the blobs we want to get
	Oids []string `protobuf:"bytes,2,rep,name=oids,proto3" json:"oids,omitempty"`
	// Maximum number of bytes we want to receive. Use '-1' to get the full blobs no matter how big.
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlobsRequest) Reset()         { *m = GetBlobsRequest{} }
func (m *GetBlobsRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlobsRequest) ProtoMessage()    {}
func (*GetBlobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6903d1e8a20272e8, []int{6}
}

func (m *GetBlobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlobsRequest.Unmarshal(m, b)
}
func (m *GetBlobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlobsRequest.Marshal(b, m, deterministic)
}
func (m *GetBlobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlobsRequest.Merge(m, src)
}
func (m *GetBlobsRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlobsRequest.Size(m)
}
func (m *GetBlobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlobsRequest proto.InternalMessageInfo

func (m *GetBlobsRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetBlobsRequest) GetOids() []string {
	if m != nil {
		return m.Oids
	}
	return nil
}

func (m *GetBlobsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetBlobsResponse struct {
	// Blob size; present only on the first message per blob
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Chunk of blob data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Object ID of the current blob. Only present on the first message per blob. Empty if no blob was found.
	Oid                  string   `protobuf:"bytes,3,opt,name=oid,proto3" json:"oid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlobsResponse) Reset()         { *m = GetBlobsResponse{} }
func (m *GetBlobsResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlobsResponse) ProtoMessage()    {}
func (*GetBlobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6903d1e8a20272e8, []int{7}
}

func (m *GetBlobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlobsResponse.Unmarshal(m, b)
}
func (m *GetBlobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlobsResponse.Marshal(b, m, deterministic)
}
func (m *GetBlobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlobsResponse.Merge(m, src)
}
func (m *GetBlobsResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlobsResponse.Size(m)
}
func (m *GetBlobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlobsResponse proto.InternalMessageInfo

func (m *GetBlobsResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetBlobsResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetBlobsResponse) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBlobRequest)(nil), "satellite.GetBlobRequest")
	proto.RegisterType((*GetBlobResponse)(nil), "satellite.GetBlobResponse")
	proto.RegisterType((*IsBinaryRequest)(nil), "satellite.IsBinaryRequest")
	proto.RegisterType((*IsBinaryResponse)(nil), "satellite.IsBinaryResponse")
	proto.RegisterType((*GetSizeRequest)(nil), "satellite.GetSizeRequest")
	proto.RegisterType((*GetSizeResponse)(nil), "satellite.GetSizeResponse")
	proto.RegisterType((*GetBlobsRequest)(nil), "satellite.GetBlobsRequest")
	proto.RegisterType((*GetBlobsResponse)(nil), "satellite.GetBlobsResponse")
}

func init() { proto.RegisterFile("blob.proto", fileDescriptor_6903d1e8a20272e8) }

var fileDescriptor_6903d1e8a20272e8 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0xea, 0x30,
	0x10, 0x24, 0x84, 0xc7, 0x83, 0x05, 0x3d, 0x90, 0xf5, 0x68, 0xd3, 0xf4, 0x12, 0x59, 0xaa, 0x14,
	0xf5, 0x10, 0x55, 0x54, 0xfd, 0x81, 0xa8, 0x15, 0xaa, 0xda, 0x93, 0x39, 0xb5, 0x37, 0x87, 0x58,
	0xd4, 0x52, 0xc0, 0xa9, 0x6d, 0x2a, 0xc1, 0x0f, 0xf5, 0x37, 0xab, 0x38, 0x24, 0xa4, 0x51, 0x38,
	0xc1, 0x6d, 0x9d, 0x75, 0x66, 0x66, 0x67, 0xd6, 0x00, 0x51, 0x22, 0xa2, 0x20, 0x95, 0x42, 0x0b,
	0xd4, 0x57, 0x54, 0xb3, 0x24, 0xe1, 0x9a, 0xb9, 0x43, 0xf5, 0x41, 0x25, 0x8b, 0xf3, 0x06, 0x16,
	0xf0, 0x6f, 0xc6, 0x74, 0x98, 0x88, 0x88, 0xb0, 0xcf, 0x0d, 0x53, 0x1a, 0x3d, 0x00, 0x48, 0x96,
	0x0a, 0xc5, 0xb5, 0x90, 0x5b, 0xc7, 0xf2, 0x2c, 0x7f, 0x30, 0x9d, 0x04, 0xe5, 0xff, 0x01, 0x29,
	0x9b, 0xa4, 0x72, 0x11, 0x8d, 0xc1, 0x16, 0x3c, 0x76, 0xda, 0x9e, 0xe5, 0xf7, 0x49, 0x56, 0xa2,
	0xff, 0xf0, 0x27, 0xe1, 0x2b, 0xae, 0x1d, 0xdb, 0xb3, 0x7c, 0x9b, 0xe4, 0x07, 0xfc, 0x02, 0xa3,
	0x92, 0x50, 0xa5, 0x62, 0xad, 0x18, 0x42, 0xd0, 0x51, 0x7c, 0xc7, 0x0c, 0x97, 0x4d, 0x4c, 0x9d,
	0x7d, 0x8b, 0xa9, 0xa6, 0x06, 0x6f, 0x48, 0x4c, 0x5d, 0x50, 0xd8, 0x25, 0x05, 0x7e, 0x87, 0xd1,
	0xb3, 0x0a, 0xf9, 0x9a, 0xca, 0xed, 0xb9, 0xe5, 0xe3, 0x5b, 0x18, 0x1f, 0xb0, 0xf7, 0x4a, 0x2f,
	0xa0, 0x2b, 0x99, 0xda, 0x24, 0xda, 0x00, 0xf7, 0xc8, 0xfe, 0x84, 0xdf, 0x8c, 0x8b, 0x73, 0xbe,
	0x63, 0x67, 0x97, 0x71, 0x63, 0xfc, 0xca, 0xa1, 0x8f, 0xfb, 0x85, 0x65, 0x69, 0xab, 0x3a, 0x51,
	0x02, 0x82, 0x8e, 0xe0, 0xb1, 0x72, 0xda, 0x9e, 0xed, 0xf7, 0x89, 0xa9, 0x8f, 0x44, 0xf9, 0x0a,
	0xe3, 0x03, 0xe7, 0xa9, 0x59, 0x4e, 0xbf, 0xdb, 0x30, 0xc8, 0xb0, 0xe6, 0x4c, 0x7e, 0xf1, 0x05,
	0x43, 0x8f, 0xf0, 0x77, 0x8f, 0x8e, 0xae, 0x2a, 0xaa, 0x7f, 0x6f, 0xab, 0xeb, 0x36, 0xb5, 0x72,
	0x2d, 0xb8, 0x75, 0x67, 0xa1, 0x27, 0xe8, 0x15, 0x29, 0xa2, 0xea, 0xdd, 0xda, 0xda, 0xb8, 0xd7,
	0x8d, 0xbd, 0x02, 0x08, 0xcd, 0xa0, 0x57, 0x8c, 0x8a, 0x1a, 0x28, 0x55, 0x13, 0x4c, 0xdd, 0x1b,
	0xa3, 0x27, 0x34, 0x53, 0x65, 0x71, 0xd6, 0xa7, 0xaa, 0x6c, 0x4f, 0x7d, 0xaa, 0x6a, 0xfa, 0xb8,
	0x15, 0x5e, 0xc2, 0x64, 0x21, 0x56, 0x01, 0x4d, 0x78, 0x44, 0x23, 0x1a, 0xd0, 0x25, 0xd7, 0xc1,
	0x52, 0xa6, 0x8b, 0xa8, 0x6b, 0xde, 0xf4, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0x2f,
	0xfc, 0xff, 0xfa, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlobServiceClient is the client API for BlobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlobServiceClient interface {
	// GetBlob returns the contents of a blob object referenced by its object
	// ID. We use a stream to return a chunked arbitrarily large binary
	// response
	GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (BlobService_GetBlobClient, error)
	IsBinary(ctx context.Context, in *IsBinaryRequest, opts ...grpc.CallOption) (*IsBinaryResponse, error)
	// GetBlobsBySHA returns the contents of a blob objects referenced by their object
	// ID. We use a stream to return a chunked arbitrarily large binary response.
	// The blobs are sent in a continous stream, the caller is responsible for spliting
	// them up into multiple blobs by their object IDs.
	GetBlobs(ctx context.Context, in *GetBlobsRequest, opts ...grpc.CallOption) (BlobService_GetBlobsClient, error)
	GetSize(ctx context.Context, in *GetSizeRequest, opts ...grpc.CallOption) (*GetSizeResponse, error)
}

type blobServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlobServiceClient(cc *grpc.ClientConn) BlobServiceClient {
	return &blobServiceClient{cc}
}

func (c *blobServiceClient) GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (BlobService_GetBlobClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlobService_serviceDesc.Streams[0], "/satellite.BlobService/GetBlob", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServiceGetBlobClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlobService_GetBlobClient interface {
	Recv() (*GetBlobResponse, error)
	grpc.ClientStream
}

type blobServiceGetBlobClient struct {
	grpc.ClientStream
}

func (x *blobServiceGetBlobClient) Recv() (*GetBlobResponse, error) {
	m := new(GetBlobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blobServiceClient) IsBinary(ctx context.Context, in *IsBinaryRequest, opts ...grpc.CallOption) (*IsBinaryResponse, error) {
	out := new(IsBinaryResponse)
	err := c.cc.Invoke(ctx, "/satellite.BlobService/IsBinary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobServiceClient) GetBlobs(ctx context.Context, in *GetBlobsRequest, opts ...grpc.CallOption) (BlobService_GetBlobsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlobService_serviceDesc.Streams[1], "/satellite.BlobService/GetBlobs", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServiceGetBlobsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlobService_GetBlobsClient interface {
	Recv() (*GetBlobsResponse, error)
	grpc.ClientStream
}

type blobServiceGetBlobsClient struct {
	grpc.ClientStream
}

func (x *blobServiceGetBlobsClient) Recv() (*GetBlobsResponse, error) {
	m := new(GetBlobsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blobServiceClient) GetSize(ctx context.Context, in *GetSizeRequest, opts ...grpc.CallOption) (*GetSizeResponse, error) {
	out := new(GetSizeResponse)
	err := c.cc.Invoke(ctx, "/satellite.BlobService/GetSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlobServiceServer is the server API for BlobService service.
type BlobServiceServer interface {
	// GetBlob returns the contents of a blob object referenced by its object
	// ID. We use a stream to return a chunked arbitrarily large binary
	// response
	GetBlob(*GetBlobRequest, BlobService_GetBlobServer) error
	IsBinary(context.Context, *IsBinaryRequest) (*IsBinaryResponse, error)
	// GetBlobsBySHA returns the contents of a blob objects referenced by their object
	// ID. We use a stream to return a chunked arbitrarily large binary response.
	// The blobs are sent in a continous stream, the caller is responsible for spliting
	// them up into multiple blobs by their object IDs.
	GetBlobs(*GetBlobsRequest, BlobService_GetBlobsServer) error
	GetSize(context.Context, *GetSizeRequest) (*GetSizeResponse, error)
}

// UnimplementedBlobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlobServiceServer struct {
}

func (*UnimplementedBlobServiceServer) GetBlob(req *GetBlobRequest, srv BlobService_GetBlobServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlob not implemented")
}
func (*UnimplementedBlobServiceServer) IsBinary(ctx context.Context, req *IsBinaryRequest) (*IsBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBinary not implemented")
}
func (*UnimplementedBlobServiceServer) GetBlobs(req *GetBlobsRequest, srv BlobService_GetBlobsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlobs not implemented")
}
func (*UnimplementedBlobServiceServer) GetSize(ctx context.Context, req *GetSizeRequest) (*GetSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSize not implemented")
}

func RegisterBlobServiceServer(s *grpc.Server, srv BlobServiceServer) {
	s.RegisterService(&_BlobService_serviceDesc, srv)
}

func _BlobService_GetBlob_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlobRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlobServiceServer).GetBlob(m, &blobServiceGetBlobServer{stream})
}

type BlobService_GetBlobServer interface {
	Send(*GetBlobResponse) error
	grpc.ServerStream
}

type blobServiceGetBlobServer struct {
	grpc.ServerStream
}

func (x *blobServiceGetBlobServer) Send(m *GetBlobResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BlobService_IsBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).IsBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/satellite.BlobService/IsBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).IsBinary(ctx, req.(*IsBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobService_GetBlobs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlobsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlobServiceServer).GetBlobs(m, &blobServiceGetBlobsServer{stream})
}

type BlobService_GetBlobsServer interface {
	Send(*GetBlobsResponse) error
	grpc.ServerStream
}

type blobServiceGetBlobsServer struct {
	grpc.ServerStream
}

func (x *blobServiceGetBlobsServer) Send(m *GetBlobsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BlobService_GetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).GetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/satellite.BlobService/GetSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).GetSize(ctx, req.(*GetSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "satellite.BlobService",
	HandlerType: (*BlobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsBinary",
			Handler:    _BlobService_IsBinary_Handler,
		},
		{
			MethodName: "GetSize",
			Handler:    _BlobService_GetSize_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlob",
			Handler:       _BlobService_GetBlob_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBlobs",
			Handler:       _BlobService_GetBlobs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blob.proto",
}
