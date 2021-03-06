// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared.proto

package satellite

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
	GlRepositoryId int64 `protobuf:"varint,7,opt,name=gl_repository_id,json=glRepositoryId,proto3" json:"gl_repository_id,omitempty"`
	// The human-readable GitLab project path (e.g. gitlab-org/gitlab-ce).
	// When hashed storage is use, this associates a project path with its
	// path on disk. The name can change over time (e.g. when a project is
	// renamed). This is primarily used for logging/debugging at the
	// moment.
	GlProjectPath string `protobuf:"bytes,8,opt,name=gl_project_path,json=glProjectPath,proto3" json:"gl_project_path,omitempty"`
	// configs which will be set into GIT_CONFIG_PARAMETERS
	Configs              []string `protobuf:"bytes,9,rep,name=configs,proto3" json:"configs,omitempty"`
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

func (m *Repository) GetGlProjectPath() string {
	if m != nil {
		return m.GlProjectPath
	}
	return ""
}

func (m *Repository) GetConfigs() []string {
	if m != nil {
		return m.Configs
	}
	return nil
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a4e87e678c5ced, []int{8}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("satellite.SignatureType", SignatureType_name, SignatureType_value)
	proto.RegisterType((*Repository)(nil), "satellite.Repository")
	proto.RegisterType((*GitCommit)(nil), "satellite.GitCommit")
	proto.RegisterType((*CommitAuthor)(nil), "satellite.CommitAuthor")
	proto.RegisterType((*ExitStatus)(nil), "satellite.ExitStatus")
	proto.RegisterType((*Namespace)(nil), "satellite.Namespace")
	proto.RegisterType((*Branch)(nil), "satellite.Branch")
	proto.RegisterType((*Tag)(nil), "satellite.Tag")
	proto.RegisterType((*User)(nil), "satellite.User")
	proto.RegisterType((*Empty)(nil), "satellite.Empty")
}

func init() { proto.RegisterFile("shared.proto", fileDescriptor_d8a4e87e678c5ced) }

var fileDescriptor_d8a4e87e678c5ced = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5d, 0x6f, 0xe4, 0x34,
	0x14, 0x25, 0x99, 0xcc, 0x47, 0xee, 0x64, 0x96, 0x91, 0x29, 0xda, 0xa8, 0x68, 0xb5, 0x43, 0x90,
	0xd0, 0x08, 0xa1, 0x74, 0x55, 0xb4, 0x0f, 0xfb, 0xb8, 0x85, 0xaa, 0x2a, 0x0f, 0xed, 0x28, 0x6d,
	0x05, 0x6f, 0x91, 0x27, 0xb9, 0xf5, 0x18, 0x39, 0x1f, 0xb2, 0x3d, 0x15, 0xd3, 0x67, 0x7e, 0x21,
	0xbf, 0x03, 0xfe, 0x03, 0xb2, 0x9d, 0x0c, 0x53, 0xa8, 0x2a, 0x1e, 0x78, 0xf3, 0xbd, 0x3e, 0xbe,
	0xf7, 0x9e, 0xe3, 0x63, 0x43, 0xa4, 0x36, 0x54, 0x62, 0x99, 0xb6, 0xb2, 0xd1, 0x0d, 0x09, 0x15,
	0xd5, 0x28, 0x04, 0xd7, 0x78, 0xfc, 0x96, 0x35, 0x0d, 0x13, 0x78, 0x62, 0x37, 0xd6, 0xdb, 0xfb,
	0x13, 0xcd, 0x2b, 0x54, 0x9a, 0x56, 0xad, 0xc3, 0x26, 0x7f, 0xf8, 0x00, 0x19, 0xb6, 0x8d, 0xe2,
	0xba, 0x91, 0x3b, 0xf2, 0x25, 0x44, 0x4a, 0x37, 0x92, 0x32, 0xcc, 0x6b, 0x5a, 0x61, 0xec, 0x2f,
	0xbc, 0x65, 0x98, 0x4d, 0xbb, 0xdc, 0x15, 0xad, 0x90, 0x7c, 0x05, 0x33, 0x89, 0x82, 0x6a, 0xfe,
	0x80, 0x79, 0x4b, 0xf5, 0x26, 0x1e, 0x58, 0x4c, 0xd4, 0x27, 0x57, 0x54, 0x6f, 0xc8, 0x3b, 0x38,
	0x62, 0x5c, 0xe7, 0xcd, 0xfa, 0x17, 0x2c, 0x74, 0x5e, 0x72, 0x89, 0x85, 0xa9, 0x1f, 0x07, 0x16,
	0x4b, 0x18, 0xd7, 0xd7, 0x76, 0xeb, 0x87, 0x7e, 0x87, 0x5c, 0xc0, 0xc2, 0x9c, 0xa0, 0x42, 0xa3,
	0xac, 0xa9, 0xc6, 0x7f, 0x9e, 0xe5, 0xa8, 0xe2, 0xe1, 0x62, 0xb0, 0x0c, 0xb3, 0x37, 0x8c, 0xeb,
	0x8f, 0x3d, 0xec, 0x69, 0x19, 0x8e, 0xca, 0xcc, 0xc7, 0x44, 0x2e, 0xf7, 0x9c, 0xe2, 0x91, 0x9b,
	0x8f, 0x89, 0x03, 0x9e, 0x4b, 0x98, 0x3f, 0x01, 0xe5, 0xbc, 0x8c, 0xc7, 0x0b, 0x6f, 0x39, 0xc8,
	0x5e, 0x1d, 0xe2, 0x2e, 0x4b, 0xf2, 0x35, 0x7c, 0xca, 0x44, 0xde, 0xca, 0xc6, 0x4e, 0x63, 0x09,
	0x4f, 0x6c, 0xc1, 0x19, 0x13, 0x2b, 0x97, 0xb5, 0x8c, 0x63, 0x18, 0x17, 0x4d, 0x7d, 0xcf, 0x99,
	0x8a, 0x43, 0x3b, 0x66, 0x1f, 0xfe, 0x18, 0x4c, 0xbc, 0xb9, 0x9f, 0x05, 0xe6, 0x68, 0xf2, 0xa7,
	0x07, 0xe1, 0x05, 0xd7, 0xdf, 0x37, 0x55, 0xc5, 0x35, 0x79, 0x05, 0x3e, 0x2f, 0x63, 0xcf, 0x96,
	0xf3, 0x79, 0x69, 0x6a, 0xa8, 0xad, 0x25, 0x64, 0x85, 0x8f, 0xb2, 0x3e, 0x24, 0x04, 0x82, 0x75,
	0x53, 0xee, 0xac, 0xd6, 0x51, 0x66, 0xd7, 0xe4, 0x04, 0x46, 0x74, 0xab, 0x37, 0x8d, 0xb4, 0xaa,
	0x4e, 0x4f, 0x5f, 0xa7, 0xfb, 0x7b, 0x4f, 0x5d, 0x83, 0x8f, 0x76, 0x3b, 0xeb, 0x60, 0xe4, 0x3d,
	0x84, 0x85, 0xcd, 0x6b, 0x94, 0xf1, 0xf0, 0xe5, 0x33, 0x7f, 0x23, 0xc9, 0x1b, 0x80, 0x96, 0x4a,
	0xac, 0x75, 0xce, 0x4b, 0x15, 0x8f, 0x2c, 0xb9, 0xd0, 0x65, 0x2e, 0x4b, 0x45, 0xbe, 0x80, 0xd0,
	0x8c, 0x93, 0x2b, 0xfe, 0x88, 0x9d, 0x86, 0x13, 0x93, 0xb8, 0xe1, 0x8f, 0x98, 0xfc, 0xe6, 0x41,
	0x74, 0x58, 0xd7, 0x10, 0xb1, 0xc6, 0xf2, 0x1c, 0x11, 0xb3, 0x26, 0x47, 0x30, 0xc4, 0x8a, 0x72,
	0xd1, 0x91, 0x76, 0x01, 0x49, 0x21, 0x28, 0xa9, 0x46, 0x4b, 0x79, 0x7a, 0x7a, 0x9c, 0x3a, 0x27,
	0xa7, 0xbd, 0x93, 0xd3, 0xdb, 0xde, 0xc9, 0x99, 0xc5, 0x91, 0x63, 0x98, 0x18, 0x73, 0x3f, 0x36,
	0x35, 0x5a, 0x41, 0xa2, 0x6c, 0x1f, 0x27, 0x09, 0xc0, 0xf9, 0xaf, 0x5c, 0xdf, 0x68, 0xaa, 0xb7,
	0xca, 0xf4, 0x7b, 0xa0, 0x62, 0xeb, 0x86, 0x18, 0x66, 0x2e, 0x48, 0x4a, 0x08, 0x8d, 0xbf, 0x55,
	0x4b, 0x0b, 0xfc, 0xbf, 0xde, 0x41, 0x77, 0xf7, 0x61, 0xdd, 0x17, 0x4e, 0x7e, 0x82, 0xd1, 0x99,
	0xa4, 0x75, 0xb1, 0x79, 0x56, 0x89, 0x0f, 0x30, 0xd3, 0x54, 0x32, 0xd4, 0xb9, 0x93, 0xdf, 0xf6,
	0x9d, 0x9e, 0x1e, 0x1d, 0xdc, 0xd2, 0xde, 0x3d, 0x59, 0xe4, 0xa0, 0x2e, 0x4a, 0x7e, 0xf7, 0x60,
	0x70, 0x4b, 0xd9, 0xb3, 0x65, 0x9d, 0xcf, 0xfc, 0xbd, 0xcf, 0xfe, 0xd5, 0x66, 0xf0, 0x5f, 0xdb,
	0x18, 0x8b, 0x56, 0xa8, 0x14, 0x65, 0xbd, 0xc8, 0x7d, 0x68, 0x24, 0xeb, 0x96, 0xce, 0x0a, 0x43,
	0x6b, 0x85, 0x69, 0x97, 0x33, 0x6e, 0x30, 0x8e, 0xd5, 0x94, 0x31, 0x94, 0xf6, 0x4d, 0xbe, 0xe4,
	0x58, 0x07, 0x4b, 0xee, 0x21, 0xb8, 0x53, 0x28, 0xc9, 0x67, 0x30, 0x64, 0x22, 0xdf, 0xbf, 0x95,
	0x80, 0x89, 0xcb, 0x72, 0xcf, 0xd4, 0x7f, 0xce, 0x4a, 0x83, 0x43, 0x2b, 0xbd, 0x85, 0x29, 0x13,
	0xf9, 0x56, 0x99, 0x2f, 0xa3, 0xc2, 0xee, 0x13, 0x02, 0x26, 0xee, 0xba, 0x4c, 0x32, 0x86, 0xe1,
	0x79, 0xd5, 0xea, 0xdd, 0x37, 0xdf, 0xc2, 0xec, 0x86, 0xb3, 0x9a, 0xea, 0xad, 0xc4, 0xdb, 0x5d,
	0x8b, 0x64, 0x02, 0xc1, 0xd5, 0xf5, 0xd5, 0xf9, 0xfc, 0x13, 0x32, 0x86, 0xc1, 0xea, 0x62, 0x35,
	0xf7, 0x4c, 0xea, 0xe7, 0xf7, 0xef, 0x3e, 0xcc, 0xfd, 0xb3, 0xd7, 0xf0, 0x79, 0xd1, 0x54, 0x29,
	0x15, 0x7c, 0x4d, 0xd7, 0x34, 0xa5, 0x8c, 0xeb, 0x94, 0xc9, 0xb6, 0x58, 0x8f, 0xac, 0x4b, 0xbf,
	0xfb, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xff, 0x0a, 0x9c, 0x11, 0x98, 0x05, 0x00, 0x00,
}
