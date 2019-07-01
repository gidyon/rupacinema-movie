// Code generated by protoc-gen-go. DO NOT EDIT.
// source: movie.proto

package movie

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Appropriate audience for a movie
type Audience int32

const (
	// Not rated for a kind of audience
	// NR
	Audience_NOT_RATED Audience = 0
	// All ages admitted
	// GE
	Audience_GENERAL_AUDIENCE Audience = 1
	// Some material may not be suitable for children
	// PG
	Audience_PARENTAL_GUIDANCE Audience = 2
	// Some material may be unappropriate for children under 13
	// PG-13
	Audience_PARENTAL_GUIDANCE_13 Audience = 3
	// No one 17 and uUnder Admitted
	// NC-17
	Audience_NO_ONE_17 Audience = 4
)

var Audience_name = map[int32]string{
	0: "NOT_RATED",
	1: "GENERAL_AUDIENCE",
	2: "PARENTAL_GUIDANCE",
	3: "PARENTAL_GUIDANCE_13",
	4: "NO_ONE_17",
}

var Audience_value = map[string]int32{
	"NOT_RATED":            0,
	"GENERAL_AUDIENCE":     1,
	"PARENTAL_GUIDANCE":    2,
	"PARENTAL_GUIDANCE_13": 3,
	"NO_ONE_17":            4,
}

func (x Audience) String() string {
	return proto.EnumName(Audience_name, int32(x))
}

func (Audience) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{0}
}

// Movie resource represents a video content played at the cinema's
type Movie struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price                string   `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	TrailerUrl           string   `protobuf:"bytes,5,opt,name=trailer_url,json=trailerUrl,proto3" json:"trailer_url,omitempty"`
	ReleaseDate          string   `protobuf:"bytes,6,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Ratings              float32  `protobuf:"fixed32,7,opt,name=ratings,proto3" json:"ratings,omitempty"`
	MovieDurationMins    int64    `protobuf:"varint,8,opt,name=movie_duration_mins,json=movieDurationMins,proto3" json:"movie_duration_mins,omitempty"`
	Photos               []string `protobuf:"bytes,9,rep,name=photos,proto3" json:"photos,omitempty"`
	Category             []string `protobuf:"bytes,10,rep,name=category,proto3" json:"category,omitempty"`
	AudienceLabel        Audience `protobuf:"varint,11,opt,name=audience_label,json=audienceLabel,proto3,enum=rupacinema.movie.Audience" json:"audience_label,omitempty"`
	AllVotes             int32    `protobuf:"varint,12,opt,name=all_votes,json=allVotes,proto3" json:"all_votes,omitempty"`
	CurrentVotes         int32    `protobuf:"varint,13,opt,name=current_votes,json=currentVotes,proto3" json:"current_votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{0}
}

func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (m *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(m, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Movie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Movie) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Movie) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Movie) GetTrailerUrl() string {
	if m != nil {
		return m.TrailerUrl
	}
	return ""
}

func (m *Movie) GetReleaseDate() string {
	if m != nil {
		return m.ReleaseDate
	}
	return ""
}

func (m *Movie) GetRatings() float32 {
	if m != nil {
		return m.Ratings
	}
	return 0
}

func (m *Movie) GetMovieDurationMins() int64 {
	if m != nil {
		return m.MovieDurationMins
	}
	return 0
}

func (m *Movie) GetPhotos() []string {
	if m != nil {
		return m.Photos
	}
	return nil
}

func (m *Movie) GetCategory() []string {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *Movie) GetAudienceLabel() Audience {
	if m != nil {
		return m.AudienceLabel
	}
	return Audience_NOT_RATED
}

func (m *Movie) GetAllVotes() int32 {
	if m != nil {
		return m.AllVotes
	}
	return 0
}

func (m *Movie) GetCurrentVotes() int32 {
	if m != nil {
		return m.CurrentVotes
	}
	return 0
}

// Request to create a movie resource
type CreateMovieRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Movie                *Movie   `protobuf:"bytes,2,opt,name=movie,proto3" json:"movie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMovieRequest) Reset()         { *m = CreateMovieRequest{} }
func (m *CreateMovieRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMovieRequest) ProtoMessage()    {}
func (*CreateMovieRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{1}
}

func (m *CreateMovieRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMovieRequest.Unmarshal(m, b)
}
func (m *CreateMovieRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMovieRequest.Marshal(b, m, deterministic)
}
func (m *CreateMovieRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMovieRequest.Merge(m, src)
}
func (m *CreateMovieRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMovieRequest.Size(m)
}
func (m *CreateMovieRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMovieRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMovieRequest proto.InternalMessageInfo

func (m *CreateMovieRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *CreateMovieRequest) GetMovie() *Movie {
	if m != nil {
		return m.Movie
	}
	return nil
}

// Request to update a movie resource. Performs a full update
type UpdateMovieRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Movie                *Movie   `protobuf:"bytes,3,opt,name=movie,proto3" json:"movie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMovieRequest) Reset()         { *m = UpdateMovieRequest{} }
func (m *UpdateMovieRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMovieRequest) ProtoMessage()    {}
func (*UpdateMovieRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{2}
}

func (m *UpdateMovieRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMovieRequest.Unmarshal(m, b)
}
func (m *UpdateMovieRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMovieRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMovieRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMovieRequest.Merge(m, src)
}
func (m *UpdateMovieRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMovieRequest.Size(m)
}
func (m *UpdateMovieRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMovieRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMovieRequest proto.InternalMessageInfo

func (m *UpdateMovieRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *UpdateMovieRequest) GetMovie() *Movie {
	if m != nil {
		return m.Movie
	}
	return nil
}

// Request to remove a movie resource
type DeleteMovieRequest struct {
	MovieId              string   `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	AdminId              string   `protobuf:"bytes,2,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMovieRequest) Reset()         { *m = DeleteMovieRequest{} }
func (m *DeleteMovieRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMovieRequest) ProtoMessage()    {}
func (*DeleteMovieRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{3}
}

func (m *DeleteMovieRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMovieRequest.Unmarshal(m, b)
}
func (m *DeleteMovieRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMovieRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMovieRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMovieRequest.Merge(m, src)
}
func (m *DeleteMovieRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMovieRequest.Size(m)
}
func (m *DeleteMovieRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMovieRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMovieRequest proto.InternalMessageInfo

func (m *DeleteMovieRequest) GetMovieId() string {
	if m != nil {
		return m.MovieId
	}
	return ""
}

func (m *DeleteMovieRequest) GetAdminId() string {
	if m != nil {
		return m.AdminId
	}
	return ""
}

// Request to retrieve a collection of movies
type ListMoviesRequest struct {
	// Page token
	PageToken            int32    `protobuf:"varint,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMoviesRequest) Reset()         { *m = ListMoviesRequest{} }
func (m *ListMoviesRequest) String() string { return proto.CompactTextString(m) }
func (*ListMoviesRequest) ProtoMessage()    {}
func (*ListMoviesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{4}
}

func (m *ListMoviesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMoviesRequest.Unmarshal(m, b)
}
func (m *ListMoviesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMoviesRequest.Marshal(b, m, deterministic)
}
func (m *ListMoviesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMoviesRequest.Merge(m, src)
}
func (m *ListMoviesRequest) XXX_Size() int {
	return xxx_messageInfo_ListMoviesRequest.Size(m)
}
func (m *ListMoviesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMoviesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMoviesRequest proto.InternalMessageInfo

func (m *ListMoviesRequest) GetPageToken() int32 {
	if m != nil {
		return m.PageToken
	}
	return 0
}

// Response from ListMoviesRequest
type ListMoviesResponse struct {
	// To be used as page_token in the next call
	NextPageToken int32 `protobuf:"varint,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Collection of movie resource
	Movies               []*Movie `protobuf:"bytes,2,rep,name=movies,proto3" json:"movies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMoviesResponse) Reset()         { *m = ListMoviesResponse{} }
func (m *ListMoviesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMoviesResponse) ProtoMessage()    {}
func (*ListMoviesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{5}
}

func (m *ListMoviesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMoviesResponse.Unmarshal(m, b)
}
func (m *ListMoviesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMoviesResponse.Marshal(b, m, deterministic)
}
func (m *ListMoviesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMoviesResponse.Merge(m, src)
}
func (m *ListMoviesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMoviesResponse.Size(m)
}
func (m *ListMoviesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMoviesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMoviesResponse proto.InternalMessageInfo

func (m *ListMoviesResponse) GetNextPageToken() int32 {
	if m != nil {
		return m.NextPageToken
	}
	return 0
}

func (m *ListMoviesResponse) GetMovies() []*Movie {
	if m != nil {
		return m.Movies
	}
	return nil
}

// Request to retrieve a single movie resource
type GetMovieRequest struct {
	MovieId              string   `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMovieRequest) Reset()         { *m = GetMovieRequest{} }
func (m *GetMovieRequest) String() string { return proto.CompactTextString(m) }
func (*GetMovieRequest) ProtoMessage()    {}
func (*GetMovieRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{6}
}

func (m *GetMovieRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMovieRequest.Unmarshal(m, b)
}
func (m *GetMovieRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMovieRequest.Marshal(b, m, deterministic)
}
func (m *GetMovieRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMovieRequest.Merge(m, src)
}
func (m *GetMovieRequest) XXX_Size() int {
	return xxx_messageInfo_GetMovieRequest.Size(m)
}
func (m *GetMovieRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMovieRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMovieRequest proto.InternalMessageInfo

func (m *GetMovieRequest) GetMovieId() string {
	if m != nil {
		return m.MovieId
	}
	return ""
}

// Request to request for playback for a particular movie
type RequestMovieReplayRequest struct {
	MovieId              string   `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMovieReplayRequest) Reset()         { *m = RequestMovieReplayRequest{} }
func (m *RequestMovieReplayRequest) String() string { return proto.CompactTextString(m) }
func (*RequestMovieReplayRequest) ProtoMessage()    {}
func (*RequestMovieReplayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{7}
}

func (m *RequestMovieReplayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMovieReplayRequest.Unmarshal(m, b)
}
func (m *RequestMovieReplayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMovieReplayRequest.Marshal(b, m, deterministic)
}
func (m *RequestMovieReplayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMovieReplayRequest.Merge(m, src)
}
func (m *RequestMovieReplayRequest) XXX_Size() int {
	return xxx_messageInfo_RequestMovieReplayRequest.Size(m)
}
func (m *RequestMovieReplayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMovieReplayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMovieReplayRequest proto.InternalMessageInfo

func (m *RequestMovieReplayRequest) GetMovieId() string {
	if m != nil {
		return m.MovieId
	}
	return ""
}

func (m *RequestMovieReplayRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterEnum("rupacinema.movie.Audience", Audience_name, Audience_value)
	proto.RegisterType((*Movie)(nil), "rupacinema.movie.Movie")
	proto.RegisterType((*CreateMovieRequest)(nil), "rupacinema.movie.CreateMovieRequest")
	proto.RegisterType((*UpdateMovieRequest)(nil), "rupacinema.movie.UpdateMovieRequest")
	proto.RegisterType((*DeleteMovieRequest)(nil), "rupacinema.movie.DeleteMovieRequest")
	proto.RegisterType((*ListMoviesRequest)(nil), "rupacinema.movie.ListMoviesRequest")
	proto.RegisterType((*ListMoviesResponse)(nil), "rupacinema.movie.ListMoviesResponse")
	proto.RegisterType((*GetMovieRequest)(nil), "rupacinema.movie.GetMovieRequest")
	proto.RegisterType((*RequestMovieReplayRequest)(nil), "rupacinema.movie.RequestMovieReplayRequest")
}

func init() { proto.RegisterFile("movie.proto", fileDescriptor_fde087a4194eda75) }

var fileDescriptor_fde087a4194eda75 = []byte{
	// 822 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0xae, 0xa4, 0xe8, 0x6f, 0x64, 0x39, 0xf2, 0xc4, 0x55, 0x36, 0x4a, 0xda, 0x2a, 0x4c, 0x50,
	0x08, 0x6e, 0x2c, 0x21, 0xce, 0xa1, 0x40, 0x6f, 0x6c, 0x24, 0x18, 0x2a, 0x1c, 0xd9, 0x20, 0xec,
	0x1e, 0x7a, 0x61, 0x57, 0xe4, 0x44, 0x5d, 0x84, 0x22, 0xd9, 0xdd, 0x55, 0x52, 0xa3, 0xe8, 0xa5,
	0x40, 0x9f, 0xa0, 0xaf, 0xd1, 0xb7, 0xe9, 0x2b, 0xf4, 0x41, 0x0a, 0x2e, 0x49, 0x87, 0x35, 0x23,
	0xd7, 0xc8, 0x49, 0x9a, 0x6f, 0xbe, 0x99, 0x6f, 0x39, 0xf3, 0x61, 0xa0, 0xb3, 0x8e, 0xde, 0x0a,
	0x1a, 0xc7, 0x32, 0xd2, 0x11, 0xf6, 0xe4, 0x26, 0xe6, 0x9e, 0x08, 0x69, 0xcd, 0xc7, 0x06, 0x1f,
	0x3c, 0x5c, 0x45, 0xd1, 0x2a, 0xa0, 0x89, 0xc9, 0x2f, 0x37, 0xaf, 0x27, 0xb4, 0x8e, 0xf5, 0x65,
	0x4a, 0x1f, 0x3c, 0xca, 0x92, 0x3c, 0x16, 0x13, 0x1e, 0x86, 0x91, 0xe6, 0x5a, 0x44, 0xa1, 0xca,
	0xb2, 0xcf, 0xcc, 0x8f, 0x77, 0xb8, 0xa2, 0xf0, 0x50, 0xbd, 0xe3, 0xab, 0x15, 0xc9, 0x49, 0x14,
	0x1b, 0x46, 0x99, 0x6d, 0xfd, 0x55, 0x83, 0xfa, 0xab, 0x44, 0x12, 0x77, 0xa1, 0x2a, 0x7c, 0x56,
	0x19, 0x56, 0x46, 0x6d, 0xa7, 0x2a, 0x7c, 0xdc, 0x87, 0xba, 0x16, 0x3a, 0x20, 0x56, 0x35, 0x50,
	0x1a, 0x24, 0x68, 0x2c, 0x85, 0x47, 0xac, 0x96, 0xa2, 0x26, 0xc0, 0x21, 0x74, 0x7c, 0x52, 0x9e,
	0x14, 0x46, 0x87, 0xdd, 0x31, 0xb9, 0x22, 0x84, 0x5f, 0x40, 0x47, 0x4b, 0x2e, 0x02, 0x92, 0xee,
	0x46, 0x06, 0xac, 0x6e, 0x18, 0x90, 0x41, 0x17, 0x32, 0xc0, 0xc7, 0xb0, 0x23, 0x29, 0x20, 0xae,
	0xc8, 0xf5, 0xb9, 0x26, 0xd6, 0x48, 0x7b, 0x64, 0xd8, 0x94, 0x6b, 0x42, 0x06, 0x4d, 0xc9, 0xb5,
	0x08, 0x57, 0x8a, 0x35, 0x87, 0x95, 0x51, 0xd5, 0xc9, 0x43, 0x1c, 0xc3, 0x3d, 0x33, 0x37, 0xd7,
	0xdf, 0x48, 0xf3, 0x79, 0xee, 0x5a, 0x84, 0x8a, 0xb5, 0x86, 0x95, 0x51, 0xcd, 0xd9, 0x33, 0xa9,
	0x69, 0x96, 0x79, 0x25, 0x42, 0x85, 0x7d, 0x68, 0xc4, 0x3f, 0x45, 0x3a, 0x52, 0xac, 0x3d, 0xac,
	0x8d, 0xda, 0x4e, 0x16, 0xe1, 0x00, 0x5a, 0x1e, 0xd7, 0xb4, 0x8a, 0xe4, 0x25, 0x03, 0x93, 0xb9,
	0x8a, 0xd1, 0x86, 0x5d, 0xbe, 0xf1, 0x05, 0x85, 0x1e, 0xb9, 0x01, 0x5f, 0x52, 0xc0, 0x3a, 0xc3,
	0xca, 0x68, 0xf7, 0x68, 0x30, 0xbe, 0xbe, 0xbd, 0xb1, 0x9d, 0xf1, 0x9c, 0x6e, 0x5e, 0x71, 0x92,
	0x14, 0xe0, 0x43, 0x68, 0xf3, 0x20, 0x70, 0xdf, 0x46, 0x9a, 0x14, 0xdb, 0x19, 0x56, 0x46, 0x75,
	0xa7, 0xc5, 0x83, 0xe0, 0xfb, 0x24, 0xc6, 0x27, 0xd0, 0xf5, 0x36, 0x52, 0x52, 0xa8, 0x33, 0x42,
	0xd7, 0x10, 0x76, 0x32, 0xd0, 0x90, 0xac, 0x25, 0xe0, 0x4b, 0x49, 0x5c, 0x93, 0xd9, 0x99, 0x43,
	0x3f, 0x6f, 0x48, 0x69, 0xfc, 0x0c, 0x40, 0xa6, 0x7f, 0xdd, 0xab, 0x15, 0xb6, 0x33, 0x64, 0xee,
	0xe3, 0x21, 0xd4, 0xcd, 0xbb, 0xcc, 0x26, 0x3b, 0x47, 0xf7, 0xcb, 0x0f, 0x4e, 0xbb, 0xa5, 0xac,
	0x44, 0xe3, 0x22, 0xf6, 0x3f, 0x56, 0xa3, 0x76, 0x2b, 0x8d, 0xef, 0x00, 0xa7, 0x14, 0xd0, 0x35,
	0x8d, 0x07, 0xd0, 0x4a, 0xd7, 0x78, 0xa5, 0xd0, 0x34, 0xf1, 0xdc, 0x4f, 0x52, 0xdc, 0x5f, 0x8b,
	0x30, 0x49, 0xa5, 0x86, 0x6c, 0x9a, 0x78, 0xee, 0x5b, 0x47, 0xb0, 0x77, 0x22, 0x94, 0x36, 0x9d,
	0x54, 0xe1, 0xb9, 0x31, 0x5f, 0x91, 0xab, 0xa3, 0x37, 0x14, 0x9a, 0x66, 0x75, 0xa7, 0x9d, 0x20,
	0xe7, 0x09, 0x60, 0xad, 0x01, 0x8b, 0x35, 0x2a, 0x8e, 0x42, 0x45, 0xf8, 0x25, 0xdc, 0x0d, 0xe9,
	0x17, 0xed, 0x96, 0x2a, 0xbb, 0x09, 0x7c, 0x96, 0x57, 0xe3, 0x04, 0x1a, 0xe6, 0x5d, 0x8a, 0x55,
	0x87, 0xb5, 0x9b, 0xbe, 0x36, 0xa3, 0x59, 0xcf, 0xe0, 0xee, 0x31, 0xe9, 0x5b, 0x7e, 0xab, 0x75,
	0x0a, 0x0f, 0x32, 0x56, 0x56, 0x11, 0x07, 0xfc, 0xf2, 0x16, 0x33, 0xba, 0x0f, 0xcd, 0x8d, 0x22,
	0xf9, 0x7e, 0x44, 0x8d, 0x24, 0x9c, 0xfb, 0x07, 0x11, 0xb4, 0x72, 0x4b, 0x62, 0x17, 0xda, 0x8b,
	0xd3, 0x73, 0xd7, 0xb1, 0xcf, 0x67, 0xd3, 0xde, 0x27, 0xb8, 0x0f, 0xbd, 0xe3, 0xd9, 0x62, 0xe6,
	0xd8, 0x27, 0xae, 0x7d, 0x31, 0x9d, 0xcf, 0x16, 0x2f, 0x67, 0xbd, 0x0a, 0x7e, 0x0a, 0x7b, 0x67,
	0xb6, 0x33, 0x5b, 0x9c, 0xdb, 0x27, 0xee, 0xf1, 0xc5, 0x7c, 0x6a, 0x27, 0x70, 0x15, 0x19, 0xec,
	0x97, 0x60, 0xf7, 0xf9, 0x8b, 0x5e, 0x2d, 0xed, 0xea, 0x9e, 0x2e, 0x66, 0xee, 0xf3, 0xaf, 0x7b,
	0x77, 0x8e, 0xfe, 0xa8, 0x43, 0xcb, 0xbc, 0xdd, 0x3e, 0x9b, 0xa3, 0x07, 0x9d, 0x82, 0x67, 0xf1,
	0x69, 0x79, 0x58, 0x65, 0x4b, 0x0f, 0xfa, 0xe3, 0xf4, 0xc8, 0x8d, 0xf3, 0x0b, 0x38, 0x9e, 0x25,
	0x17, 0xd0, 0xea, 0xff, 0xfe, 0xf7, 0x3f, 0x7f, 0x56, 0x7b, 0x56, 0xc7, 0x5c, 0xbf, 0x74, 0xbc,
	0xdf, 0x54, 0x0e, 0x12, 0x91, 0x82, 0x69, 0x3f, 0x24, 0x52, 0xf6, 0xf4, 0xff, 0x89, 0x0c, 0xae,
	0x8b, 0xfc, 0x08, 0x9d, 0x82, 0x6b, 0x3f, 0x24, 0x52, 0x36, 0xf5, 0x56, 0x91, 0x7b, 0x46, 0xa4,
	0x7b, 0x50, 0x14, 0xc1, 0x00, 0xe0, 0xbd, 0x2f, 0xf1, 0x49, 0x59, 0xa0, 0xe4, 0xf4, 0xc1, 0xd3,
	0x9b, 0x49, 0xa9, 0xb5, 0x73, 0x35, 0xfc, 0x8f, 0xda, 0x6b, 0x68, 0xe5, 0xb6, 0xc4, 0xc7, 0xe5,
	0x36, 0xd7, 0x2c, 0x3b, 0xd8, 0x66, 0x73, 0xeb, 0x73, 0xd3, 0x9c, 0x61, 0xbf, 0xd0, 0x7c, 0xf2,
	0x6b, 0x6e, 0xd3, 0xdf, 0xf0, 0x1d, 0x60, 0xd9, 0xd0, 0xf8, 0x55, 0xb9, 0xdd, 0x56, 0xdb, 0x6f,
	0x9d, 0xe2, 0x23, 0x23, 0xdd, 0xc7, 0xfd, 0xa2, 0x74, 0x52, 0xb8, 0xe4, 0xde, 0x9b, 0x6f, 0x9b,
	0x3f, 0xa4, 0xf7, 0x66, 0xd9, 0x30, 0x65, 0x2f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x08, 0x07,
	0x95, 0xa4, 0x77, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MovieAPIClient is the client API for MovieAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieAPIClient interface {
	// Creates a new movie resource. Admins only. Requires authentication
	CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Updates information for a given movie resource. Admins only. Requires authentication.
	UpdateMovie(ctx context.Context, in *UpdateMovieRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Removes permanently a movie resource. Admins only. Requires authentication.
	DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// ListMovies retrieves a collection of movie resource
	ListMovies(ctx context.Context, in *ListMoviesRequest, opts ...grpc.CallOption) (*ListMoviesResponse, error)
	// Retrieves a movie resource. Any User.
	GetMovie(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*Movie, error)
	// Requests a movie to be replayed again. By Users. Requires authentication
	RequestMovieReplay(ctx context.Context, in *RequestMovieReplayRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type movieAPIClient struct {
	cc *grpc.ClientConn
}

func NewMovieAPIClient(cc *grpc.ClientConn) MovieAPIClient {
	return &movieAPIClient{cc}
}

func (c *movieAPIClient) CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rupacinema.movie.MovieAPI/CreateMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieAPIClient) UpdateMovie(ctx context.Context, in *UpdateMovieRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rupacinema.movie.MovieAPI/UpdateMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieAPIClient) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rupacinema.movie.MovieAPI/DeleteMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieAPIClient) ListMovies(ctx context.Context, in *ListMoviesRequest, opts ...grpc.CallOption) (*ListMoviesResponse, error) {
	out := new(ListMoviesResponse)
	err := c.cc.Invoke(ctx, "/rupacinema.movie.MovieAPI/ListMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieAPIClient) GetMovie(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/rupacinema.movie.MovieAPI/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieAPIClient) RequestMovieReplay(ctx context.Context, in *RequestMovieReplayRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rupacinema.movie.MovieAPI/RequestMovieReplay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieAPIServer is the server API for MovieAPI service.
type MovieAPIServer interface {
	// Creates a new movie resource. Admins only. Requires authentication
	CreateMovie(context.Context, *CreateMovieRequest) (*empty.Empty, error)
	// Updates information for a given movie resource. Admins only. Requires authentication.
	UpdateMovie(context.Context, *UpdateMovieRequest) (*empty.Empty, error)
	// Removes permanently a movie resource. Admins only. Requires authentication.
	DeleteMovie(context.Context, *DeleteMovieRequest) (*empty.Empty, error)
	// ListMovies retrieves a collection of movie resource
	ListMovies(context.Context, *ListMoviesRequest) (*ListMoviesResponse, error)
	// Retrieves a movie resource. Any User.
	GetMovie(context.Context, *GetMovieRequest) (*Movie, error)
	// Requests a movie to be replayed again. By Users. Requires authentication
	RequestMovieReplay(context.Context, *RequestMovieReplayRequest) (*empty.Empty, error)
}

func RegisterMovieAPIServer(s *grpc.Server, srv MovieAPIServer) {
	s.RegisterService(&_MovieAPI_serviceDesc, srv)
}

func _MovieAPI_CreateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieAPIServer).CreateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rupacinema.movie.MovieAPI/CreateMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieAPIServer).CreateMovie(ctx, req.(*CreateMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieAPI_UpdateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieAPIServer).UpdateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rupacinema.movie.MovieAPI/UpdateMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieAPIServer).UpdateMovie(ctx, req.(*UpdateMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieAPI_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieAPIServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rupacinema.movie.MovieAPI/DeleteMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieAPIServer).DeleteMovie(ctx, req.(*DeleteMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieAPI_ListMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieAPIServer).ListMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rupacinema.movie.MovieAPI/ListMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieAPIServer).ListMovies(ctx, req.(*ListMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieAPI_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieAPIServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rupacinema.movie.MovieAPI/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieAPIServer).GetMovie(ctx, req.(*GetMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieAPI_RequestMovieReplay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMovieReplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieAPIServer).RequestMovieReplay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rupacinema.movie.MovieAPI/RequestMovieReplay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieAPIServer).RequestMovieReplay(ctx, req.(*RequestMovieReplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rupacinema.movie.MovieAPI",
	HandlerType: (*MovieAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMovie",
			Handler:    _MovieAPI_CreateMovie_Handler,
		},
		{
			MethodName: "UpdateMovie",
			Handler:    _MovieAPI_UpdateMovie_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _MovieAPI_DeleteMovie_Handler,
		},
		{
			MethodName: "ListMovies",
			Handler:    _MovieAPI_ListMovies_Handler,
		},
		{
			MethodName: "GetMovie",
			Handler:    _MovieAPI_GetMovie_Handler,
		},
		{
			MethodName: "RequestMovieReplay",
			Handler:    _MovieAPI_RequestMovieReplay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}