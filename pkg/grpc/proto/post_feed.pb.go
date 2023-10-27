// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: pkg/grpc/proto/post_feed.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Post_PrivacyType int32

const (
	Post_UNKNOWN Post_PrivacyType = 0
	Post_PUBLIC  Post_PrivacyType = 1
	Post_PRIVATE Post_PrivacyType = 2 // Only me
)

// Enum value maps for Post_PrivacyType.
var (
	Post_PrivacyType_name = map[int32]string{
		0: "UNKNOWN",
		1: "PUBLIC",
		2: "PRIVATE",
	}
	Post_PrivacyType_value = map[string]int32{
		"UNKNOWN": 0,
		"PUBLIC":  1,
		"PRIVATE": 2,
	}
)

func (x Post_PrivacyType) Enum() *Post_PrivacyType {
	p := new(Post_PrivacyType)
	*p = x
	return p
}

func (x Post_PrivacyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Post_PrivacyType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_grpc_proto_post_feed_proto_enumTypes[0].Descriptor()
}

func (Post_PrivacyType) Type() protoreflect.EnumType {
	return &file_pkg_grpc_proto_post_feed_proto_enumTypes[0]
}

func (x Post_PrivacyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Post_PrivacyType.Descriptor instead.
func (Post_PrivacyType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{3, 0}
}

type GetPostFeedsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPostFeedsRequest) Reset() {
	*x = GetPostFeedsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostFeedsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostFeedsRequest) ProtoMessage() {}

func (x *GetPostFeedsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostFeedsRequest.ProtoReflect.Descriptor instead.
func (*GetPostFeedsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{0}
}

type GetPostFeedsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetPostFeedsResponse) Reset() {
	*x = GetPostFeedsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostFeedsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostFeedsResponse) ProtoMessage() {}

func (x *GetPostFeedsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostFeedsResponse.ProtoReflect.Descriptor instead.
func (*GetPostFeedsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{1}
}

func (x *GetPostFeedsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type RefreshPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refresh bool `protobuf:"varint,1,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

func (x *RefreshPostRequest) Reset() {
	*x = RefreshPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshPostRequest) ProtoMessage() {}

func (x *RefreshPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshPostRequest.ProtoReflect.Descriptor instead.
func (*RefreshPostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshPostRequest) GetRefresh() bool {
	if x != nil {
		return x.Refresh
	}
	return false
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content      string           `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	CreatedAt    string           `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    string           `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	LikeCount    int64            `protobuf:"varint,5,opt,name=LikeCount,proto3" json:"LikeCount,omitempty"`
	CommentCount int64            `protobuf:"varint,6,opt,name=CommentCount,proto3" json:"CommentCount,omitempty"`
	LikedByMe    bool             `protobuf:"varint,7,opt,name=LikedByMe,proto3" json:"LikedByMe,omitempty"`
	IsRepost     bool             `protobuf:"varint,8,opt,name=isRepost,proto3" json:"isRepost,omitempty"`
	RepostCount  int64            `protobuf:"varint,9,opt,name=RepostCount,proto3" json:"RepostCount,omitempty"`
	Privacy      Post_PrivacyType `protobuf:"varint,10,opt,name=Privacy,proto3,enum=feeds.Post_PrivacyType" json:"Privacy,omitempty"`
	Edges        *Post_Edge       `protobuf:"bytes,11,opt,name=edges,proto3" json:"edges,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{3}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Post) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Post) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *Post) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Post) GetLikedByMe() bool {
	if x != nil {
		return x.LikedByMe
	}
	return false
}

func (x *Post) GetIsRepost() bool {
	if x != nil {
		return x.IsRepost
	}
	return false
}

func (x *Post) GetRepostCount() int64 {
	if x != nil {
		return x.RepostCount
	}
	return 0
}

func (x *Post) GetPrivacy() Post_PrivacyType {
	if x != nil {
		return x.Privacy
	}
	return Post_UNKNOWN
}

func (x *Post) GetEdges() *Post_Edge {
	if x != nil {
		return x.Edges
	}
	return nil
}

type Post_Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *Post_User      `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Comments []*Post_Comment `protobuf:"bytes,2,rep,name=comments,proto3" json:"comments,omitempty"`
	Medias   []*Post_Media   `protobuf:"bytes,3,rep,name=medias,proto3" json:"medias,omitempty"`
	Business *Post_Business  `protobuf:"bytes,4,opt,name=business,proto3" json:"business,omitempty"`
	Parent   *Post           `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *Post_Edge) Reset() {
	*x = Post_Edge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post_Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post_Edge) ProtoMessage() {}

func (x *Post_Edge) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post_Edge.ProtoReflect.Descriptor instead.
func (*Post_Edge) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Post_Edge) GetUser() *Post_User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Post_Edge) GetComments() []*Post_Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *Post_Edge) GetMedias() []*Post_Media {
	if x != nil {
		return x.Medias
	}
	return nil
}

func (x *Post_Edge) GetBusiness() *Post_Business {
	if x != nil {
		return x.Business
	}
	return nil
}

func (x *Post_Edge) GetParent() *Post {
	if x != nil {
		return x.Parent
	}
	return nil
}

type Post_CommentEdge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *Post_User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Post_CommentEdge) Reset() {
	*x = Post_CommentEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post_CommentEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post_CommentEdge) ProtoMessage() {}

func (x *Post_CommentEdge) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post_CommentEdge.ProtoReflect.Descriptor instead.
func (*Post_CommentEdge) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Post_CommentEdge) GetUser() *Post_User {
	if x != nil {
		return x.User
	}
	return nil
}

type Post_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Auth0Id    string `protobuf:"bytes,2,opt,name=auth0_id,json=auth0Id,proto3" json:"auth0_id,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Picture    string `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	CoverImage string `protobuf:"bytes,5,opt,name=cover_image,json=coverImage,proto3" json:"cover_image,omitempty"`
	Username   string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	Bio        string `protobuf:"bytes,7,opt,name=bio,proto3" json:"bio,omitempty"`
}

func (x *Post_User) Reset() {
	*x = Post_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post_User) ProtoMessage() {}

func (x *Post_User) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post_User.ProtoReflect.Descriptor instead.
func (*Post_User) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{3, 2}
}

func (x *Post_User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post_User) GetAuth0Id() string {
	if x != nil {
		return x.Auth0Id
	}
	return ""
}

func (x *Post_User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Post_User) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Post_User) GetCoverImage() string {
	if x != nil {
		return x.CoverImage
	}
	return ""
}

func (x *Post_User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Post_User) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

type Post_Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	URL       string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	MediaType string `protobuf:"bytes,3,opt,name=MediaType,proto3" json:"MediaType,omitempty"`
}

func (x *Post_Media) Reset() {
	*x = Post_Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post_Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post_Media) ProtoMessage() {}

func (x *Post_Media) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post_Media.ProtoReflect.Descriptor instead.
func (*Post_Media) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{3, 3}
}

func (x *Post_Media) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post_Media) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *Post_Media) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

type Post_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content   string            `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	CreatedAt string            `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt string            `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Edges     *Post_CommentEdge `protobuf:"bytes,5,opt,name=edges,proto3" json:"edges,omitempty"`
}

func (x *Post_Comment) Reset() {
	*x = Post_Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post_Comment) ProtoMessage() {}

func (x *Post_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post_Comment.ProtoReflect.Descriptor instead.
func (*Post_Comment) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{3, 4}
}

func (x *Post_Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post_Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post_Comment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Post_Comment) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Post_Comment) GetEdges() *Post_CommentEdge {
	if x != nil {
		return x.Edges
	}
	return nil
}

type Post_Business struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Picture    string `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty"`
	CoverImage string `protobuf:"bytes,4,opt,name=cover_image,json=coverImage,proto3" json:"cover_image,omitempty"`
}

func (x *Post_Business) Reset() {
	*x = Post_Business{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post_Business) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post_Business) ProtoMessage() {}

func (x *Post_Business) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_post_feed_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post_Business.ProtoReflect.Descriptor instead.
func (*Post_Business) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_post_feed_proto_rawDescGZIP(), []int{3, 5}
}

func (x *Post_Business) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post_Business) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Post_Business) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Post_Business) GetCoverImage() string {
	if x != nil {
		return x.CoverImage
	}
	return ""
}

var File_pkg_grpc_proto_post_feed_proto protoreflect.FileDescriptor

var file_pkg_grpc_proto_post_feed_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x22, 0xb7, 0x09, 0x0a, 0x04, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x69, 0x6b, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4c, 0x69, 0x6b,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x69,
	0x6b, 0x65, 0x64, 0x42, 0x79, 0x4d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x4c,
	0x69, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x4d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x64, 0x67,
	0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x52, 0x05, 0x65, 0x64, 0x67, 0x65,
	0x73, 0x1a, 0xdf, 0x01, 0x0a, 0x04, 0x45, 0x64, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x2f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x29, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x12, 0x30, 0x0a, 0x08,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x23,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x1a, 0x33, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x64,
	0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0xae, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x30, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x30, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x1a, 0x47, 0x0a, 0x05, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x52, 0x4c, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x1a, 0x9e, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x64, 0x67, 0x65, 0x52, 0x05, 0x65, 0x64,
	0x67, 0x65, 0x73, 0x1a, 0x69, 0x0a, 0x08, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x33,
	0x0a, 0x0b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55,
	0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54,
	0x45, 0x10, 0x02, 0x32, 0xe8, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x73, 0x12, 0x1a, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x46,
	0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0b,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x1a, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x65, 0x65,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x7a,
	0x69, 0x65, 0x73, 0x74, 0x61, 0x72, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x69, 0x6f, 0x2f, 0x70, 0x6c,
	0x61, 0x63, 0x69, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_proto_post_feed_proto_rawDescOnce sync.Once
	file_pkg_grpc_proto_post_feed_proto_rawDescData = file_pkg_grpc_proto_post_feed_proto_rawDesc
)

func file_pkg_grpc_proto_post_feed_proto_rawDescGZIP() []byte {
	file_pkg_grpc_proto_post_feed_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_proto_post_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_proto_post_feed_proto_rawDescData)
	})
	return file_pkg_grpc_proto_post_feed_proto_rawDescData
}

var file_pkg_grpc_proto_post_feed_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_grpc_proto_post_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_grpc_proto_post_feed_proto_goTypes = []interface{}{
	(Post_PrivacyType)(0),        // 0: feeds.Post.PrivacyType
	(*GetPostFeedsRequest)(nil),  // 1: feeds.GetPostFeedsRequest
	(*GetPostFeedsResponse)(nil), // 2: feeds.GetPostFeedsResponse
	(*RefreshPostRequest)(nil),   // 3: feeds.RefreshPostRequest
	(*Post)(nil),                 // 4: feeds.Post
	(*Post_Edge)(nil),            // 5: feeds.Post.Edge
	(*Post_CommentEdge)(nil),     // 6: feeds.Post.CommentEdge
	(*Post_User)(nil),            // 7: feeds.Post.User
	(*Post_Media)(nil),           // 8: feeds.Post.Media
	(*Post_Comment)(nil),         // 9: feeds.Post.Comment
	(*Post_Business)(nil),        // 10: feeds.Post.Business
}
var file_pkg_grpc_proto_post_feed_proto_depIdxs = []int32{
	4,  // 0: feeds.GetPostFeedsResponse.posts:type_name -> feeds.Post
	0,  // 1: feeds.Post.Privacy:type_name -> feeds.Post.PrivacyType
	5,  // 2: feeds.Post.edges:type_name -> feeds.Post.Edge
	7,  // 3: feeds.Post.Edge.user:type_name -> feeds.Post.User
	9,  // 4: feeds.Post.Edge.comments:type_name -> feeds.Post.Comment
	8,  // 5: feeds.Post.Edge.medias:type_name -> feeds.Post.Media
	10, // 6: feeds.Post.Edge.business:type_name -> feeds.Post.Business
	4,  // 7: feeds.Post.Edge.parent:type_name -> feeds.Post
	7,  // 8: feeds.Post.CommentEdge.user:type_name -> feeds.Post.User
	6,  // 9: feeds.Post.Comment.edges:type_name -> feeds.Post.CommentEdge
	1,  // 10: feeds.PostService.GetPostFeeds:input_type -> feeds.GetPostFeedsRequest
	3,  // 11: feeds.PostService.RefreshPost:input_type -> feeds.RefreshPostRequest
	1,  // 12: feeds.PostService.WatchPosts:input_type -> feeds.GetPostFeedsRequest
	2,  // 13: feeds.PostService.GetPostFeeds:output_type -> feeds.GetPostFeedsResponse
	2,  // 14: feeds.PostService.RefreshPost:output_type -> feeds.GetPostFeedsResponse
	2,  // 15: feeds.PostService.WatchPosts:output_type -> feeds.GetPostFeedsResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_grpc_proto_post_feed_proto_init() }
func file_pkg_grpc_proto_post_feed_proto_init() {
	if File_pkg_grpc_proto_post_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_proto_post_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostFeedsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_proto_post_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostFeedsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_proto_post_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_proto_post_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_proto_post_feed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post_Edge); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_proto_post_feed_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post_CommentEdge); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_proto_post_feed_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post_User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_proto_post_feed_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post_Media); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_proto_post_feed_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post_Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_grpc_proto_post_feed_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post_Business); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_grpc_proto_post_feed_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_proto_post_feed_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_proto_post_feed_proto_depIdxs,
		EnumInfos:         file_pkg_grpc_proto_post_feed_proto_enumTypes,
		MessageInfos:      file_pkg_grpc_proto_post_feed_proto_msgTypes,
	}.Build()
	File_pkg_grpc_proto_post_feed_proto = out.File
	file_pkg_grpc_proto_post_feed_proto_rawDesc = nil
	file_pkg_grpc_proto_post_feed_proto_goTypes = nil
	file_pkg_grpc_proto_post_feed_proto_depIdxs = nil
}
