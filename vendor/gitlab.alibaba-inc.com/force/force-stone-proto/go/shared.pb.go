// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared.proto

package gitaly

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type SignatureType int32

const (
	SignatureType_NONE SignatureType = 0
	SignatureType_PGP  SignatureType = 1
	SignatureType_X509 SignatureType = 2
)

var SignatureType_name = map[int32]string{
	0: "NONE",
	1: "PGP",
	2: "X509",
}

var SignatureType_value = map[string]int32{
	"NONE": 0,
	"PGP":  1,
	"X509": 2,
}

func (x SignatureType) String() string {
	return proto.EnumName(SignatureType_name, int32(x))
}

func (SignatureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8a4e87e678c5ced, []int{0}
}

type Repository struct {
	StorageName  string `protobuf:"bytes,2,opt,name=storage_name,json=storageName,proto3" json:"storage_name,omitempty"`
	RelativePath string `protobuf:"bytes,3,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
	// Sets the GIT_OBJECT_DIRECTORY envvar on git commands to the value of this field.
	// It influences the object storage directory the SHA1 directories are created underneath.
	GitObjectDirectory string `protobuf:"bytes,4,opt,name=git_object_directory,json=gitObjectDirectory,proto3" json:"git_object_directory,omitempty"`
	// Sets the GIT_ALTERNATE_OBJECT_DIRECTORIES envvar on git commands to the values of this field.
	// It influences the list of Git object directories which can be used to search for Git objects.
	GitAlternateObjectDirectories []string `protobuf:"bytes,5,rep,name=git_alternate_object_directories,json=gitAlternateObjectDirectories,proto3" json:"git_alternate_object_directories,omitempty"`
	// Used in callbacks to GitLab so that it knows what repository the event is
	// associated with. May be left empty on RPC's that do not perform callbacks.
	GlRepository string `protobuf:"bytes,6,opt,name=gl_repository,json=glRepository,proto3" json:"gl_repository,omitempty"`
	// Repository id
	GlRepositoryId       int64    `protobuf:"varint,7,opt,name=gl_repository_id,json=glRepositoryId,proto3" json:"gl_repository_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a4e87e678c5ced, []int{0}
}

func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

func (m *Repository) GetRelativePath() string {
	if m != nil {
		return m.RelativePath
	}
	return ""
}

func (m *Repository) GetGitObjectDirectory() string {
	if m != nil {
		return m.GitObjectDirectory
	}
	return ""
}

func (m *Repository) GetGitAlternateObjectDirectories() []string {
	if m != nil {
		return m.GitAlternateObjectDirectories
	}
	return nil
}

func (m *Repository) GetGlRepository() string {
	if m != nil {
		return m.GlRepository
	}
	return ""
}

func (m *Repository) GetGlRepositoryId() int64 {
	if m != nil {
		return m.GlRepositoryId
	}
	return 0
}

// Corresponds to Gitlab::Git::Commit
type GitCommit struct {
	Id        string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Subject   []byte        `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body      []byte        `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Author    *CommitAuthor `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Committer *CommitAuthor `protobuf:"bytes,5,opt,name=committer,proto3" json:"committer,omitempty"`
	ParentIds []string      `protobuf:"bytes,6,rep,name=parent_ids,json=parentIds,proto3" json:"parent_ids,omitempty"`
	// If body exceeds a certain threshold, it will be nullified,
	// but its size will be set in body_size so we can know if
	// a commit had a body in the first place.
	BodySize             int64    `protobuf:"varint,7,opt,name=body_size,json=bodySize,proto3" json:"body_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitCommit) Reset()         { *m = GitCommit{} }
func (m *GitCommit) String() string { return proto.CompactTextString(m) }
func (*GitCommit) ProtoMessage()    {}
func (*GitCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a4e87e678c5ced, []int{1}
}

func (m *GitCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitCommit.Unmarshal(m, b)
}
func (m *GitCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitCommit.Marshal(b, m, deterministic)
}
func (m *GitCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitCommit.Merge(m, src)
}
func (m *GitCommit) XXX_Size() int {
	return xxx_messageInfo_GitCommit.Size(m)
}
func (m *GitCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_GitCommit.DiscardUnknown(m)
}

var xxx_messageInfo_GitCommit proto.InternalMessageInfo

func (m *GitCommit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GitCommit) GetSubject() []byte {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *GitCommit) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *GitCommit) GetAuthor() *CommitAuthor {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *GitCommit) GetCommitter() *CommitAuthor {
	if m != nil {
		return m.Committer
	}
	return nil
}

func (m *GitCommit) GetParentIds() []string {
	if m != nil {
		return m.ParentIds
	}
	return nil
}

func (m *GitCommit) GetBodySize() int64 {
	if m != nil {
		return m.BodySize
	}
	return 0
}

type CommitAuthor struct {
	Name                 []byte               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                []byte               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Timezone             []byte               `protobuf:"bytes,4,opt,name=timezone,proto3" json:"timezone,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CommitAuthor) Reset()         { *m = CommitAuthor{} }
func (m *CommitAuthor) String() string { return proto.CompactTextString(m) }
func (*CommitAuthor) ProtoMessage()    {}
func (*CommitAuthor) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a4e87e678c5ced, []int{2}
}

func (m *CommitAuthor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitAuthor.Unmarshal(m, b)
}
func (m *CommitAuthor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitAuthor.Marshal(b, m, deterministic)
}
func (m *CommitAuthor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitAuthor.Merge(m, src)
}
func (m *CommitAuthor) XXX_Size() int {
	return xxx_messageInfo_CommitAuthor.Size(m)
}
func (m *CommitAuthor) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitAuthor.DiscardUnknown(m)
}

var xxx_messageInfo_CommitAuthor proto.InternalMessageInfo

func (m *CommitAuthor) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CommitAuthor) GetEmail() []byte {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *CommitAuthor) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *CommitAuthor) GetTimezone() []byte {
	if m != nil {
		return m.Timezone
	}
	return nil
}

type ExitStatus struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitStatus) Reset()         { *m = ExitStatus{} }
func (m *ExitStatus) String() string { return proto.CompactTextString(m) }
func (*ExitStatus) ProtoMessage()    {}
func (*ExitStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a4e87e678c5ced, []int{3}
}

func (m *ExitStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitStatus.Unmarshal(m, b)
}
func (m *ExitStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitStatus.Marshal(b, m, deterministic)
}
func (m *ExitStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitStatus.Merge(m, src)
}
func (m *ExitStatus) XXX_Size() int {
	return xxx_messageInfo_ExitStatus.Size(m)
}
func (m *ExitStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ExitStatus proto.InternalMessageInfo

func (m *ExitStatus) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Namespace struct {
	StorageName          string   `protobuf:"bytes,2,opt,name=storage_name,json=storageName,proto3" json:"storage_name,omitempty"`
	RelativePath         string   `protobuf:"bytes,3,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a4e87e678c5ced, []int{4}
}

func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

func (m *Namespace) GetRelativePath() string {
	if m != nil {
		return m.RelativePath
	}
	return ""
}

// Corresponds to Gitlab::Git::Branch
type Branch struct {
	Name                 []byte     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TargetCommit         *GitCommit `protobuf:"bytes,2,opt,name=target_commit,json=targetCommit,proto3" json:"target_commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Branch) Reset()         { *m = Branch{} }
func (m *Branch) String() string { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()    {}
func (*Branch) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a4e87e678c5ced, []int{5}
}

func (m *Branch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Branch.Unmarshal(m, b)
}
func (m *Branch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Branch.Marshal(b, m, deterministic)
}
func (m *Branch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Branch.Merge(m, src)
}
func (m *Branch) XXX_Size() int {
	return xxx_messageInfo_Branch.Size(m)
}
func (m *Branch) XXX_DiscardUnknown() {
	xxx_messageInfo_Branch.DiscardUnknown(m)
}

var xxx_messageInfo_Branch proto.InternalMessageInfo

func (m *Branch) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Branch) GetTargetCommit() *GitCommit {
	if m != nil {
		return m.TargetCommit
	}
	return nil
}

type Tag struct {
	Name         []byte     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id           string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	TargetCommit *GitCommit `protobuf:"bytes,3,opt,name=target_commit,json=targetCommit,proto3" json:"target_commit,omitempty"`
	// If message exceeds a certain threshold, it will be nullified,
	// but its size will be set in message_size so we can know if
	// a tag had a message in the first place.
	Message              []byte        `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	MessageSize          int64         `protobuf:"varint,5,opt,name=message_size,json=messageSize,proto3" json:"message_size,omitempty"`
	Tagger               *CommitAuthor `protobuf:"bytes,6,opt,name=tagger,proto3" json:"tagger,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a4e87e678c5ced, []int{6}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Tag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tag) GetTargetCommit() *GitCommit {
	if m != nil {
		return m.TargetCommit
	}
	return nil
}

func (m *Tag) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Tag) GetMessageSize() int64 {
	if m != nil {
		return m.MessageSize
	}
	return 0
}

func (m *Tag) GetTagger() *CommitAuthor {
	if m != nil {
		return m.Tagger
	}
	return nil
}

type User struct {
	GlId                 string   `protobuf:"bytes,1,opt,name=gl_id,json=glId,proto3" json:"gl_id,omitempty"`
	Name                 []byte   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                []byte   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	GlUsername           string   `protobuf:"bytes,4,opt,name=gl_username,json=glUsername,proto3" json:"gl_username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a4e87e678c5ced, []int{7}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetGlId() string {
	if m != nil {
		return m.GlId
	}
	return ""
}

func (m *User) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *User) GetEmail() []byte {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *User) GetGlUsername() string {
	if m != nil {
		return m.GlUsername
	}
	return ""
}

func init() {
	proto.RegisterEnum("gitaly.SignatureType", SignatureType_name, SignatureType_value)
	proto.RegisterType((*Repository)(nil), "gitaly.Repository")
	proto.RegisterType((*GitCommit)(nil), "gitaly.GitCommit")
	proto.RegisterType((*CommitAuthor)(nil), "gitaly.CommitAuthor")
	proto.RegisterType((*ExitStatus)(nil), "gitaly.ExitStatus")
	proto.RegisterType((*Namespace)(nil), "gitaly.Namespace")
	proto.RegisterType((*Branch)(nil), "gitaly.Branch")
	proto.RegisterType((*Tag)(nil), "gitaly.Tag")
	proto.RegisterType((*User)(nil), "gitaly.User")
}

func init() { proto.RegisterFile("shared.proto", fileDescriptor_d8a4e87e678c5ced) }

var fileDescriptor_d8a4e87e678c5ced = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6e, 0xdb, 0x38,
	0x10, 0x86, 0x57, 0xb2, 0xec, 0x58, 0x63, 0x25, 0xf0, 0x72, 0x83, 0x85, 0x90, 0x45, 0x10, 0xaf,
	0xf6, 0x62, 0x2c, 0x02, 0x25, 0x70, 0xd1, 0x02, 0x3d, 0x26, 0x6d, 0x10, 0xa4, 0x87, 0x24, 0x50,
	0x1c, 0xa0, 0x37, 0x81, 0x96, 0x26, 0x34, 0x0b, 0xc9, 0x12, 0x48, 0x3a, 0xa8, 0x73, 0xee, 0xd3,
	0xf4, 0x59, 0xfa, 0x1e, 0x7d, 0x8d, 0x82, 0xa4, 0xe4, 0x3a, 0x6d, 0x5a, 0xf4, 0xd0, 0x9b, 0x66,
	0xf8, 0x73, 0x38, 0xf3, 0xf3, 0xa3, 0x20, 0x90, 0x73, 0x2a, 0x30, 0x8f, 0x6b, 0x51, 0xa9, 0x8a,
	0xf4, 0x18, 0x57, 0xb4, 0x58, 0xed, 0x1d, 0xb0, 0xaa, 0x62, 0x05, 0x1e, 0x99, 0xec, 0x6c, 0x79,
	0x77, 0xa4, 0x78, 0x89, 0x52, 0xd1, 0xb2, 0xb6, 0xc2, 0xe8, 0xa3, 0x0b, 0x90, 0x60, 0x5d, 0x49,
	0xae, 0x2a, 0xb1, 0x22, 0xff, 0x42, 0x20, 0x55, 0x25, 0x28, 0xc3, 0x74, 0x41, 0x4b, 0x0c, 0xdd,
	0x91, 0x33, 0xf6, 0x93, 0x41, 0x93, 0xbb, 0xa4, 0x25, 0x92, 0xff, 0x60, 0x5b, 0x60, 0x41, 0x15,
	0xbf, 0xc7, 0xb4, 0xa6, 0x6a, 0x1e, 0x76, 0x8c, 0x26, 0x68, 0x93, 0xd7, 0x54, 0xcd, 0xc9, 0x31,
	0xec, 0x32, 0xae, 0xd2, 0x6a, 0xf6, 0x0e, 0x33, 0x95, 0xe6, 0x5c, 0x60, 0xa6, 0xeb, 0x87, 0x9e,
	0xd1, 0x12, 0xc6, 0xd5, 0x95, 0x59, 0x7a, 0xdd, 0xae, 0x90, 0x73, 0x18, 0xe9, 0x1d, 0xb4, 0x50,
	0x28, 0x16, 0x54, 0xe1, 0xb7, 0x7b, 0x39, 0xca, 0xb0, 0x3b, 0xea, 0x8c, 0xfd, 0x64, 0x9f, 0x71,
	0x75, 0xd2, 0xca, 0x1e, 0x97, 0xe1, 0x28, 0x75, 0x7f, 0xac, 0x48, 0xc5, 0x7a, 0xa6, 0xb0, 0x67,
	0xfb, 0x63, 0xc5, 0xc6, 0x9c, 0x63, 0x18, 0x3e, 0x12, 0xa5, 0x3c, 0x0f, 0xb7, 0x46, 0xce, 0xb8,
	0x93, 0xec, 0x6c, 0xea, 0x2e, 0xf2, 0x37, 0x5e, 0xdf, 0x19, 0xba, 0x89, 0xa7, 0x27, 0x8d, 0x3e,
	0x3b, 0xe0, 0x9f, 0x73, 0xf5, 0xaa, 0x2a, 0x4b, 0xae, 0xc8, 0x0e, 0xb8, 0x3c, 0x0f, 0x1d, 0x53,
	0xdd, 0xe5, 0x39, 0x09, 0x61, 0x4b, 0x2e, 0x4d, 0x3b, 0xc6, 0xb6, 0x20, 0x69, 0x43, 0x42, 0xc0,
	0x9b, 0x55, 0xf9, 0xca, 0x38, 0x15, 0x24, 0xe6, 0x9b, 0x1c, 0x42, 0x8f, 0x2e, 0xd5, 0xbc, 0x12,
	0xc6, 0x93, 0xc1, 0x64, 0x37, 0xb6, 0x57, 0x16, 0xdb, 0xea, 0x27, 0x66, 0x2d, 0x69, 0x34, 0x64,
	0x02, 0x7e, 0x66, 0xf2, 0x0a, 0x45, 0xd8, 0xfd, 0xc9, 0x86, 0xaf, 0x32, 0xb2, 0x0f, 0x50, 0x53,
	0x81, 0x0b, 0x95, 0xf2, 0x5c, 0x86, 0x3d, 0xe3, 0x9d, 0x6f, 0x33, 0x17, 0xb9, 0x24, 0xff, 0x80,
	0xaf, 0x1b, 0x49, 0x25, 0x7f, 0xc0, 0x66, 0xf6, 0xbe, 0x4e, 0xdc, 0xf0, 0x07, 0x8c, 0x3e, 0x38,
	0x10, 0x6c, 0xd6, 0xd5, 0x23, 0x18, 0x20, 0x1c, 0x3b, 0x82, 0xfe, 0x26, 0xbb, 0xd0, 0xc5, 0x92,
	0xf2, 0xa2, 0x19, 0xd7, 0x06, 0x24, 0x06, 0x2f, 0xa7, 0x0a, 0xcd, 0xb0, 0x83, 0xc9, 0x5e, 0x6c,
	0x09, 0x8c, 0x5b, 0x02, 0xe3, 0x69, 0x4b, 0x60, 0x62, 0x74, 0x64, 0x0f, 0xfa, 0x1a, 0xca, 0x87,
	0x6a, 0x81, 0xc6, 0x8a, 0x20, 0x59, 0xc7, 0x51, 0x04, 0x70, 0xf6, 0x9e, 0xab, 0x1b, 0x45, 0xd5,
	0x52, 0xea, 0xf3, 0xee, 0x69, 0xb1, 0xb4, 0x4d, 0x74, 0x13, 0x1b, 0x44, 0x39, 0xf8, 0x9a, 0x4b,
	0x59, 0xd3, 0x0c, 0x7f, 0x17, 0xbf, 0xcd, 0xad, 0xfb, 0x8b, 0xb6, 0x70, 0x34, 0x85, 0xde, 0xa9,
	0xa0, 0x8b, 0x6c, 0xfe, 0xa4, 0x13, 0x2f, 0x60, 0x5b, 0x51, 0xc1, 0x50, 0xa5, 0xd6, 0x7e, 0x73,
	0xee, 0x60, 0xf2, 0x67, 0x7b, 0x45, 0x6b, 0x68, 0x92, 0xc0, 0xea, 0x6c, 0x14, 0x7d, 0x72, 0xa0,
	0x33, 0xa5, 0xec, 0xc9, 0x9a, 0x16, 0x2f, 0x77, 0x8d, 0xd7, 0x77, 0x67, 0x74, 0x7e, 0xe9, 0x0c,
	0x8d, 0x65, 0x89, 0x52, 0x52, 0xd6, 0xda, 0xdb, 0x86, 0xda, 0xac, 0xe6, 0xd3, 0x42, 0xd0, 0x35,
	0x10, 0x0c, 0x9a, 0x9c, 0xe6, 0x40, 0x53, 0xaa, 0x28, 0x63, 0x28, 0xcc, 0x2b, 0xfa, 0x21, 0xa5,
	0x56, 0x13, 0xdd, 0x81, 0x77, 0x2b, 0x51, 0x90, 0xbf, 0xa0, 0xcb, 0x8a, 0x74, 0xfd, 0x38, 0x3c,
	0x56, 0x5c, 0xe4, 0xeb, 0x19, 0xdd, 0xa7, 0x08, 0xea, 0x6c, 0x12, 0x74, 0x00, 0x03, 0x56, 0xa4,
	0x4b, 0xa9, 0x5f, 0x78, 0x89, 0xcd, 0x3f, 0x03, 0x58, 0x71, 0xdb, 0x64, 0xfe, 0x3f, 0x84, 0xed,
	0x1b, 0xce, 0x16, 0x54, 0x2d, 0x05, 0x4e, 0x57, 0x35, 0x92, 0x3e, 0x78, 0x97, 0x57, 0x97, 0x67,
	0xc3, 0x3f, 0xc8, 0x16, 0x74, 0xae, 0xcf, 0xaf, 0x87, 0x8e, 0x4e, 0xbd, 0x7d, 0x7e, 0xfc, 0x72,
	0xe8, 0x9e, 0x86, 0xf0, 0x77, 0x56, 0x95, 0x31, 0x2d, 0xf8, 0x8c, 0xce, 0x68, 0x7c, 0x57, 0x89,
	0x0c, 0x63, 0x26, 0xea, 0x6c, 0xd6, 0x33, 0x50, 0x3e, 0xfb, 0x12, 0x00, 0x00, 0xff, 0xff, 0x34,
	0x2c, 0x38, 0x08, 0x3c, 0x05, 0x00, 0x00,
}
