// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: content.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TitleContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link             string   `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Title            string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	TitleOrig        string   `protobuf:"bytes,3,opt,name=titleOrig,proto3" json:"titleOrig,omitempty"`
	OtherTitle       string   `protobuf:"bytes,4,opt,name=otherTitle,proto3" json:"otherTitle,omitempty"`
	Year             int32    `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
	KinopoiskID      string   `protobuf:"bytes,6,opt,name=kinopoiskID,proto3" json:"kinopoiskID,omitempty"`
	ShikimoriID      string   `protobuf:"bytes,7,opt,name=shikimoriID,proto3" json:"shikimoriID,omitempty"`
	ImdbID           string   `protobuf:"bytes,8,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	AnimeStatus      string   `protobuf:"bytes,9,opt,name=animeStatus,proto3" json:"animeStatus,omitempty"`
	AnimeDescription string   `protobuf:"bytes,10,opt,name=animeDescription,proto3" json:"animeDescription,omitempty"`
	PosterURL        string   `protobuf:"bytes,11,opt,name=posterURL,proto3" json:"posterURL,omitempty"`
	Duration         int32    `protobuf:"varint,12,opt,name=duration,proto3" json:"duration,omitempty"`
	KinopoiskRating  float64  `protobuf:"fixed64,13,opt,name=kinopoiskRating,proto3" json:"kinopoiskRating,omitempty"`
	ImdbRating       float64  `protobuf:"fixed64,14,opt,name=imdbRating,proto3" json:"imdbRating,omitempty"`
	ShikimoriRating  float64  `protobuf:"fixed64,15,opt,name=shikimoriRating,proto3" json:"shikimoriRating,omitempty"`
	PremiereWorld    string   `protobuf:"bytes,16,opt,name=premiereWorld,proto3" json:"premiereWorld,omitempty"`
	EpisodesTotal    int32    `protobuf:"varint,17,opt,name=episodesTotal,proto3" json:"episodesTotal,omitempty"`
	Writers          []string `protobuf:"bytes,18,rep,name=writers,proto3" json:"writers,omitempty"`
	Screenshots      []string `protobuf:"bytes,19,rep,name=screenshots,proto3" json:"screenshots,omitempty"`
}

func (x *TitleContent) Reset() {
	*x = TitleContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TitleContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TitleContent) ProtoMessage() {}

func (x *TitleContent) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TitleContent.ProtoReflect.Descriptor instead.
func (*TitleContent) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{0}
}

func (x *TitleContent) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *TitleContent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TitleContent) GetTitleOrig() string {
	if x != nil {
		return x.TitleOrig
	}
	return ""
}

func (x *TitleContent) GetOtherTitle() string {
	if x != nil {
		return x.OtherTitle
	}
	return ""
}

func (x *TitleContent) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *TitleContent) GetKinopoiskID() string {
	if x != nil {
		return x.KinopoiskID
	}
	return ""
}

func (x *TitleContent) GetShikimoriID() string {
	if x != nil {
		return x.ShikimoriID
	}
	return ""
}

func (x *TitleContent) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *TitleContent) GetAnimeStatus() string {
	if x != nil {
		return x.AnimeStatus
	}
	return ""
}

func (x *TitleContent) GetAnimeDescription() string {
	if x != nil {
		return x.AnimeDescription
	}
	return ""
}

func (x *TitleContent) GetPosterURL() string {
	if x != nil {
		return x.PosterURL
	}
	return ""
}

func (x *TitleContent) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *TitleContent) GetKinopoiskRating() float64 {
	if x != nil {
		return x.KinopoiskRating
	}
	return 0
}

func (x *TitleContent) GetImdbRating() float64 {
	if x != nil {
		return x.ImdbRating
	}
	return 0
}

func (x *TitleContent) GetShikimoriRating() float64 {
	if x != nil {
		return x.ShikimoriRating
	}
	return 0
}

func (x *TitleContent) GetPremiereWorld() string {
	if x != nil {
		return x.PremiereWorld
	}
	return ""
}

func (x *TitleContent) GetEpisodesTotal() int32 {
	if x != nil {
		return x.EpisodesTotal
	}
	return 0
}

func (x *TitleContent) GetWriters() []string {
	if x != nil {
		return x.Writers
	}
	return nil
}

func (x *TitleContent) GetScreenshots() []string {
	if x != nil {
		return x.Screenshots
	}
	return nil
}

type TitleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opt string `protobuf:"bytes,1,opt,name=opt,proto3" json:"opt,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *TitleFilter) Reset() {
	*x = TitleFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TitleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TitleFilter) ProtoMessage() {}

func (x *TitleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TitleFilter.ProtoReflect.Descriptor instead.
func (*TitleFilter) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{1}
}

func (x *TitleFilter) GetOpt() string {
	if x != nil {
		return x.Opt
	}
	return ""
}

func (x *TitleFilter) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type GetTitleContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *TitleFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetTitleContentRequest) Reset() {
	*x = GetTitleContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTitleContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTitleContentRequest) ProtoMessage() {}

func (x *GetTitleContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTitleContentRequest.ProtoReflect.Descriptor instead.
func (*GetTitleContentRequest) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{2}
}

func (x *GetTitleContentRequest) GetFilter() *TitleFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetTitleContentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TitleContent []*TitleContent `protobuf:"bytes,1,rep,name=titleContent,proto3" json:"titleContent,omitempty"`
}

func (x *GetTitleContentReply) Reset() {
	*x = GetTitleContentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTitleContentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTitleContentReply) ProtoMessage() {}

func (x *GetTitleContentReply) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTitleContentReply.ProtoReflect.Descriptor instead.
func (*GetTitleContentReply) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{3}
}

func (x *GetTitleContentReply) GetTitleContent() []*TitleContent {
	if x != nil {
		return x.TitleContent
	}
	return nil
}

var File_content_proto protoreflect.FileDescriptor

var file_content_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x04,
	0x0a, 0x0c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x4f, 0x72, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x6b, 0x69,
	0x6e, 0x6f, 0x70, 0x6f, 0x69, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6b, 0x69, 0x6e, 0x6f, 0x70, 0x6f, 0x69, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x68, 0x69, 0x6b, 0x69, 0x6d, 0x6f, 0x72, 0x69, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x68, 0x69, 0x6b, 0x69, 0x6d, 0x6f, 0x72, 0x69, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x55, 0x52,
	0x4c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x55,
	0x52, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x0f, 0x6b, 0x69, 0x6e, 0x6f, 0x70, 0x6f, 0x69, 0x73, 0x6b, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6b, 0x69, 0x6e, 0x6f, 0x70, 0x6f, 0x69,
	0x73, 0x6b, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x64, 0x62,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x69, 0x6d,
	0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x68, 0x69, 0x6b,
	0x69, 0x6d, 0x6f, 0x72, 0x69, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0f, 0x73, 0x68, 0x69, 0x6b, 0x69, 0x6d, 0x6f, 0x72, 0x69, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x6d, 0x69,
	0x65, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x22, 0x31, 0x0a, 0x0b, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x70, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x70, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x3e, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x49, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x0c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x67, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_content_proto_rawDescOnce sync.Once
	file_content_proto_rawDescData = file_content_proto_rawDesc
)

func file_content_proto_rawDescGZIP() []byte {
	file_content_proto_rawDescOnce.Do(func() {
		file_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_proto_rawDescData)
	})
	return file_content_proto_rawDescData
}

var file_content_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_content_proto_goTypes = []interface{}{
	(*TitleContent)(nil),           // 0: TitleContent
	(*TitleFilter)(nil),            // 1: TitleFilter
	(*GetTitleContentRequest)(nil), // 2: GetTitleContentRequest
	(*GetTitleContentReply)(nil),   // 3: GetTitleContentReply
}
var file_content_proto_depIdxs = []int32{
	1, // 0: GetTitleContentRequest.filter:type_name -> TitleFilter
	0, // 1: GetTitleContentReply.titleContent:type_name -> TitleContent
	2, // 2: ContentService.GetTitleContent:input_type -> GetTitleContentRequest
	3, // 3: ContentService.GetTitleContent:output_type -> GetTitleContentReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_content_proto_init() }
func file_content_proto_init() {
	if File_content_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TitleContent); i {
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
		file_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TitleFilter); i {
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
		file_content_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTitleContentRequest); i {
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
		file_content_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTitleContentReply); i {
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
			RawDescriptor: file_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_content_proto_goTypes,
		DependencyIndexes: file_content_proto_depIdxs,
		MessageInfos:      file_content_proto_msgTypes,
	}.Build()
	File_content_proto = out.File
	file_content_proto_rawDesc = nil
	file_content_proto_goTypes = nil
	file_content_proto_depIdxs = nil
}