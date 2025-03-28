// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: housework/v1/housework.proto

package housework

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Chore struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Complete      bool                   `protobuf:"varint,1,opt,name=complete,proto3" json:"complete,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Chore) Reset() {
	*x = Chore{}
	mi := &file_housework_v1_housework_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chore) ProtoMessage() {}

func (x *Chore) ProtoReflect() protoreflect.Message {
	mi := &file_housework_v1_housework_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chore.ProtoReflect.Descriptor instead.
func (*Chore) Descriptor() ([]byte, []int) {
	return file_housework_v1_housework_proto_rawDescGZIP(), []int{0}
}

func (x *Chore) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

func (x *Chore) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Chores struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chores        []*Chore               `protobuf:"bytes,1,rep,name=chores,proto3" json:"chores,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Chores) Reset() {
	*x = Chores{}
	mi := &file_housework_v1_housework_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chores) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chores) ProtoMessage() {}

func (x *Chores) ProtoReflect() protoreflect.Message {
	mi := &file_housework_v1_housework_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chores.ProtoReflect.Descriptor instead.
func (*Chores) Descriptor() ([]byte, []int) {
	return file_housework_v1_housework_proto_rawDescGZIP(), []int{1}
}

func (x *Chores) GetChores() []*Chore {
	if x != nil {
		return x.Chores
	}
	return nil
}

type CompleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChoreNumber   int32                  `protobuf:"varint,1,opt,name=chore_number,json=choreNumber,proto3" json:"chore_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompleteRequest) Reset() {
	*x = CompleteRequest{}
	mi := &file_housework_v1_housework_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteRequest) ProtoMessage() {}

func (x *CompleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_housework_v1_housework_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteRequest.ProtoReflect.Descriptor instead.
func (*CompleteRequest) Descriptor() ([]byte, []int) {
	return file_housework_v1_housework_proto_rawDescGZIP(), []int{2}
}

func (x *CompleteRequest) GetChoreNumber() int32 {
	if x != nil {
		return x.ChoreNumber
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_housework_v1_housework_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_housework_v1_housework_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_housework_v1_housework_proto_rawDescGZIP(), []int{3}
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_housework_v1_housework_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_housework_v1_housework_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_housework_v1_housework_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_housework_v1_housework_proto protoreflect.FileDescriptor

const file_housework_v1_housework_proto_rawDesc = "" +
	"\n" +
	"\x1chousework/v1/housework.proto\x12\thousework\"E\n" +
	"\x05Chore\x12\x1a\n" +
	"\bcomplete\x18\x01 \x01(\bR\bcomplete\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\"2\n" +
	"\x06Chores\x12(\n" +
	"\x06chores\x18\x01 \x03(\v2\x10.housework.ChoreR\x06chores\"4\n" +
	"\x0fCompleteRequest\x12!\n" +
	"\fchore_number\x18\x01 \x01(\x05R\vchoreNumber\"\a\n" +
	"\x05Empty\"$\n" +
	"\bResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2\xa4\x01\n" +
	"\tRobotMaid\x12-\n" +
	"\x03Add\x12\x11.housework.Chores\x1a\x13.housework.Response\x12;\n" +
	"\bComplete\x12\x1a.housework.CompleteRequest\x1a\x13.housework.Response\x12+\n" +
	"\x04List\x12\x10.housework.Empty\x1a\x11.housework.ChoresBMZKgithub.com/testaquatic/NetworkProgrammingWithGo/ch12/housework/v1/houseworkb\x06proto3"

var (
	file_housework_v1_housework_proto_rawDescOnce sync.Once
	file_housework_v1_housework_proto_rawDescData []byte
)

func file_housework_v1_housework_proto_rawDescGZIP() []byte {
	file_housework_v1_housework_proto_rawDescOnce.Do(func() {
		file_housework_v1_housework_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_housework_v1_housework_proto_rawDesc), len(file_housework_v1_housework_proto_rawDesc)))
	})
	return file_housework_v1_housework_proto_rawDescData
}

var file_housework_v1_housework_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_housework_v1_housework_proto_goTypes = []any{
	(*Chore)(nil),           // 0: housework.Chore
	(*Chores)(nil),          // 1: housework.Chores
	(*CompleteRequest)(nil), // 2: housework.CompleteRequest
	(*Empty)(nil),           // 3: housework.Empty
	(*Response)(nil),        // 4: housework.Response
}
var file_housework_v1_housework_proto_depIdxs = []int32{
	0, // 0: housework.Chores.chores:type_name -> housework.Chore
	1, // 1: housework.RobotMaid.Add:input_type -> housework.Chores
	2, // 2: housework.RobotMaid.Complete:input_type -> housework.CompleteRequest
	3, // 3: housework.RobotMaid.List:input_type -> housework.Empty
	4, // 4: housework.RobotMaid.Add:output_type -> housework.Response
	4, // 5: housework.RobotMaid.Complete:output_type -> housework.Response
	1, // 6: housework.RobotMaid.List:output_type -> housework.Chores
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_housework_v1_housework_proto_init() }
func file_housework_v1_housework_proto_init() {
	if File_housework_v1_housework_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_housework_v1_housework_proto_rawDesc), len(file_housework_v1_housework_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_housework_v1_housework_proto_goTypes,
		DependencyIndexes: file_housework_v1_housework_proto_depIdxs,
		MessageInfos:      file_housework_v1_housework_proto_msgTypes,
	}.Build()
	File_housework_v1_housework_proto = out.File
	file_housework_v1_housework_proto_goTypes = nil
	file_housework_v1_housework_proto_depIdxs = nil
}
