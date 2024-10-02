// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.11.4
// source: errdefs.proto

package errdefs

import (
	pb "github.com/moby/buildkit/solver/pb"
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

type Vertex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digest string `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *Vertex) Reset() {
	*x = Vertex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdefs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vertex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vertex) ProtoMessage() {}

func (x *Vertex) ProtoReflect() protoreflect.Message {
	mi := &file_errdefs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vertex.ProtoReflect.Descriptor instead.
func (*Vertex) Descriptor() ([]byte, []int) {
	return file_errdefs_proto_rawDescGZIP(), []int{0}
}

func (x *Vertex) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info   *pb.SourceInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Ranges []*pb.Range    `protobuf:"bytes,2,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdefs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_errdefs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_errdefs_proto_rawDescGZIP(), []int{1}
}

func (x *Source) GetInfo() *pb.SourceInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Source) GetRanges() []*pb.Range {
	if x != nil {
		return x.Ranges
	}
	return nil
}

type FrontendCap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FrontendCap) Reset() {
	*x = FrontendCap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdefs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontendCap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendCap) ProtoMessage() {}

func (x *FrontendCap) ProtoReflect() protoreflect.Message {
	mi := &file_errdefs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendCap.ProtoReflect.Descriptor instead.
func (*FrontendCap) Descriptor() ([]byte, []int) {
	return file_errdefs_proto_rawDescGZIP(), []int{2}
}

func (x *FrontendCap) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Subrequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Subrequest) Reset() {
	*x = Subrequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdefs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subrequest) ProtoMessage() {}

func (x *Subrequest) ProtoReflect() protoreflect.Message {
	mi := &file_errdefs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subrequest.ProtoReflect.Descriptor instead.
func (*Subrequest) Descriptor() ([]byte, []int) {
	return file_errdefs_proto_rawDescGZIP(), []int{3}
}

func (x *Subrequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Solve struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputIDs []string `protobuf:"bytes,1,rep,name=inputIDs,proto3" json:"inputIDs,omitempty"`
	MountIDs []string `protobuf:"bytes,2,rep,name=mountIDs,proto3" json:"mountIDs,omitempty"`
	Op       *pb.Op   `protobuf:"bytes,3,opt,name=op,proto3" json:"op,omitempty"`
	// Types that are assignable to Subject:
	//
	//	*Solve_File
	//	*Solve_Cache
	Subject     isSolve_Subject   `protobuf_oneof:"subject"`
	Description map[string]string `protobuf:"bytes,6,rep,name=description,proto3" json:"description,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Solve) Reset() {
	*x = Solve{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdefs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Solve) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solve) ProtoMessage() {}

func (x *Solve) ProtoReflect() protoreflect.Message {
	mi := &file_errdefs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solve.ProtoReflect.Descriptor instead.
func (*Solve) Descriptor() ([]byte, []int) {
	return file_errdefs_proto_rawDescGZIP(), []int{4}
}

func (x *Solve) GetInputIDs() []string {
	if x != nil {
		return x.InputIDs
	}
	return nil
}

func (x *Solve) GetMountIDs() []string {
	if x != nil {
		return x.MountIDs
	}
	return nil
}

func (x *Solve) GetOp() *pb.Op {
	if x != nil {
		return x.Op
	}
	return nil
}

func (m *Solve) GetSubject() isSolve_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (x *Solve) GetFile() *FileAction {
	if x, ok := x.GetSubject().(*Solve_File); ok {
		return x.File
	}
	return nil
}

func (x *Solve) GetCache() *ContentCache {
	if x, ok := x.GetSubject().(*Solve_Cache); ok {
		return x.Cache
	}
	return nil
}

func (x *Solve) GetDescription() map[string]string {
	if x != nil {
		return x.Description
	}
	return nil
}

type isSolve_Subject interface {
	isSolve_Subject()
}

type Solve_File struct {
	File *FileAction `protobuf:"bytes,4,opt,name=file,proto3,oneof"`
}

type Solve_Cache struct {
	Cache *ContentCache `protobuf:"bytes,5,opt,name=cache,proto3,oneof"`
}

func (*Solve_File) isSolve_Subject() {}

func (*Solve_Cache) isSolve_Subject() {}

type FileAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index of the file action that failed the exec.
	Index int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *FileAction) Reset() {
	*x = FileAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdefs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileAction) ProtoMessage() {}

func (x *FileAction) ProtoReflect() protoreflect.Message {
	mi := &file_errdefs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileAction.ProtoReflect.Descriptor instead.
func (*FileAction) Descriptor() ([]byte, []int) {
	return file_errdefs_proto_rawDescGZIP(), []int{5}
}

func (x *FileAction) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

type ContentCache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Original index of result that failed the slow cache calculation.
	Index int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *ContentCache) Reset() {
	*x = ContentCache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdefs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentCache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentCache) ProtoMessage() {}

func (x *ContentCache) ProtoReflect() protoreflect.Message {
	mi := &file_errdefs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentCache.ProtoReflect.Descriptor instead.
func (*ContentCache) Descriptor() ([]byte, []int) {
	return file_errdefs_proto_rawDescGZIP(), []int{6}
}

func (x *ContentCache) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_errdefs_proto protoreflect.FileDescriptor

var file_errdefs_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x64, 0x65, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x65, 0x72, 0x72, 0x64, 0x65, 0x66, 0x73, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x62, 0x79, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b,
	0x69, 0x74, 0x2f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x70, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0b, 0x46, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x43, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0a,
	0x53, 0x75, 0x62, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbf,
	0x02, 0x0a, 0x05, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x49, 0x44, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x73,
	0x12, 0x16, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x70,
	0x62, 0x2e, 0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x72, 0x72, 0x64, 0x65, 0x66, 0x73,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x72, 0x72, 0x64, 0x65, 0x66, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x48, 0x00, 0x52, 0x05, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x72, 0x72, 0x64, 0x65, 0x66,
	0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3e, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x22, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x22, 0x24, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x62, 0x79, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x65, 0x72,
	0x72, 0x64, 0x65, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errdefs_proto_rawDescOnce sync.Once
	file_errdefs_proto_rawDescData = file_errdefs_proto_rawDesc
)

func file_errdefs_proto_rawDescGZIP() []byte {
	file_errdefs_proto_rawDescOnce.Do(func() {
		file_errdefs_proto_rawDescData = protoimpl.X.CompressGZIP(file_errdefs_proto_rawDescData)
	})
	return file_errdefs_proto_rawDescData
}

var file_errdefs_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_errdefs_proto_goTypes = []interface{}{
	(*Vertex)(nil),        // 0: errdefs.Vertex
	(*Source)(nil),        // 1: errdefs.Source
	(*FrontendCap)(nil),   // 2: errdefs.FrontendCap
	(*Subrequest)(nil),    // 3: errdefs.Subrequest
	(*Solve)(nil),         // 4: errdefs.Solve
	(*FileAction)(nil),    // 5: errdefs.FileAction
	(*ContentCache)(nil),  // 6: errdefs.ContentCache
	nil,                   // 7: errdefs.Solve.DescriptionEntry
	(*pb.SourceInfo)(nil), // 8: pb.SourceInfo
	(*pb.Range)(nil),      // 9: pb.Range
	(*pb.Op)(nil),         // 10: pb.Op
}
var file_errdefs_proto_depIdxs = []int32{
	8,  // 0: errdefs.Source.info:type_name -> pb.SourceInfo
	9,  // 1: errdefs.Source.ranges:type_name -> pb.Range
	10, // 2: errdefs.Solve.op:type_name -> pb.Op
	5,  // 3: errdefs.Solve.file:type_name -> errdefs.FileAction
	6,  // 4: errdefs.Solve.cache:type_name -> errdefs.ContentCache
	7,  // 5: errdefs.Solve.description:type_name -> errdefs.Solve.DescriptionEntry
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_errdefs_proto_init() }
func file_errdefs_proto_init() {
	if File_errdefs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errdefs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vertex); i {
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
		file_errdefs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
		file_errdefs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontendCap); i {
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
		file_errdefs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subrequest); i {
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
		file_errdefs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Solve); i {
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
		file_errdefs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileAction); i {
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
		file_errdefs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentCache); i {
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
	file_errdefs_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Solve_File)(nil),
		(*Solve_Cache)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errdefs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errdefs_proto_goTypes,
		DependencyIndexes: file_errdefs_proto_depIdxs,
		MessageInfos:      file_errdefs_proto_msgTypes,
	}.Build()
	File_errdefs_proto = out.File
	file_errdefs_proto_rawDesc = nil
	file_errdefs_proto_goTypes = nil
	file_errdefs_proto_depIdxs = nil
}
