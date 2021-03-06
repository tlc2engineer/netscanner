// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Empty
	Tag
	Tags
	Names
	Succesfull
	Ident
	ServName
	Live
	ChangeTag
	ChangeTags
	CreateTag
	TextMessage
*/
package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Пустое сообщение
type Empty struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Empty) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// тег
type Tag struct {
	Data  float64 `protobuf:"fixed64,1,opt,name=data" json:"data,omitempty"`
	Idata int64   `protobuf:"varint,2,opt,name=idata" json:"idata,omitempty"`
	Type  uint32  `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	Name  string  `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Tag) GetData() float64 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *Tag) GetIdata() int64 {
	if m != nil {
		return m.Idata
	}
	return 0
}

func (m *Tag) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Теги
type Tags struct {
	Tags []*Tag `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty"`
}

func (m *Tags) Reset()                    { *m = Tags{} }
func (m *Tags) String() string            { return proto.CompactTextString(m) }
func (*Tags) ProtoMessage()               {}
func (*Tags) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Tags) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Имена
type Names struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *Names) Reset()                    { *m = Names{} }
func (m *Names) String() string            { return proto.CompactTextString(m) }
func (*Names) ProtoMessage()               {}
func (*Names) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Names) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// Удачно
type Succesfull struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *Succesfull) Reset()                    { *m = Succesfull{} }
func (m *Succesfull) String() string            { return proto.CompactTextString(m) }
func (*Succesfull) ProtoMessage()               {}
func (*Succesfull) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Succesfull) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

// Идентификатор сервиса
type Ident struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Number int32  `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
}

func (m *Ident) Reset()                    { *m = Ident{} }
func (m *Ident) String() string            { return proto.CompactTextString(m) }
func (*Ident) ProtoMessage()               {}
func (*Ident) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Ident) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ident) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

// Имя сервиса
type ServName struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Descr string `protobuf:"bytes,2,opt,name=descr" json:"descr,omitempty"`
}

func (m *ServName) Reset()                    { *m = ServName{} }
func (m *ServName) String() string            { return proto.CompactTextString(m) }
func (*ServName) ProtoMessage()               {}
func (*ServName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServName) GetDescr() string {
	if m != nil {
		return m.Descr
	}
	return ""
}

// Live
type Live struct {
	Number int32  `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Id     *Ident `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Stop   bool   `protobuf:"varint,3,opt,name=stop" json:"stop,omitempty"`
}

func (m *Live) Reset()                    { *m = Live{} }
func (m *Live) String() string            { return proto.CompactTextString(m) }
func (*Live) ProtoMessage()               {}
func (*Live) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Live) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Live) GetId() *Ident {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Live) GetStop() bool {
	if m != nil {
		return m.Stop
	}
	return false
}

// Изменить тег
type ChangeTag struct {
	Tag   *Tag   `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	Ident *Ident `protobuf:"bytes,2,opt,name=ident" json:"ident,omitempty"`
}

func (m *ChangeTag) Reset()                    { *m = ChangeTag{} }
func (m *ChangeTag) String() string            { return proto.CompactTextString(m) }
func (*ChangeTag) ProtoMessage()               {}
func (*ChangeTag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ChangeTag) GetTag() *Tag {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *ChangeTag) GetIdent() *Ident {
	if m != nil {
		return m.Ident
	}
	return nil
}

// Изменить теги
type ChangeTags struct {
	Tags  *Tags  `protobuf:"bytes,1,opt,name=tags" json:"tags,omitempty"`
	Ident *Ident `protobuf:"bytes,2,opt,name=ident" json:"ident,omitempty"`
}

func (m *ChangeTags) Reset()                    { *m = ChangeTags{} }
func (m *ChangeTags) String() string            { return proto.CompactTextString(m) }
func (*ChangeTags) ProtoMessage()               {}
func (*ChangeTags) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ChangeTags) GetTags() *Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ChangeTags) GetIdent() *Ident {
	if m != nil {
		return m.Ident
	}
	return nil
}

// Создать тег
type CreateTag struct {
	Tag   *Tag   `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	Ident *Ident `protobuf:"bytes,2,opt,name=ident" json:"ident,omitempty"`
}

func (m *CreateTag) Reset()                    { *m = CreateTag{} }
func (m *CreateTag) String() string            { return proto.CompactTextString(m) }
func (*CreateTag) ProtoMessage()               {}
func (*CreateTag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CreateTag) GetTag() *Tag {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *CreateTag) GetIdent() *Ident {
	if m != nil {
		return m.Ident
	}
	return nil
}

// Текстовое сообщение
type TextMessage struct {
	Id   *Ident `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *TextMessage) Reset()                    { *m = TextMessage{} }
func (m *TextMessage) String() string            { return proto.CompactTextString(m) }
func (*TextMessage) ProtoMessage()               {}
func (*TextMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TextMessage) GetId() *Ident {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TextMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*Tag)(nil), "Tag")
	proto.RegisterType((*Tags)(nil), "Tags")
	proto.RegisterType((*Names)(nil), "Names")
	proto.RegisterType((*Succesfull)(nil), "Succesfull")
	proto.RegisterType((*Ident)(nil), "Ident")
	proto.RegisterType((*ServName)(nil), "ServName")
	proto.RegisterType((*Live)(nil), "Live")
	proto.RegisterType((*ChangeTag)(nil), "ChangeTag")
	proto.RegisterType((*ChangeTags)(nil), "ChangeTags")
	proto.RegisterType((*CreateTag)(nil), "CreateTag")
	proto.RegisterType((*TextMessage)(nil), "TextMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GetTags service

type GetTagsClient interface {
	// Чтение тегов без регистрации
	GetNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Names, error)
	Get(ctx context.Context, in *Names, opts ...grpc.CallOption) (*Tags, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Tags, error)
	SetTag(ctx context.Context, in *ChangeTag, opts ...grpc.CallOption) (*Succesfull, error)
	SetTags(ctx context.Context, in *ChangeTags, opts ...grpc.CallOption) (*Succesfull, error)
	Register(ctx context.Context, in *ServName, opts ...grpc.CallOption) (*Ident, error)
	UnRegister(ctx context.Context, in *Ident, opts ...grpc.CallOption) (*Succesfull, error)
	Create(ctx context.Context, in *CreateTag, opts ...grpc.CallOption) (*Succesfull, error)
	IsAlive(ctx context.Context, in *Live, opts ...grpc.CallOption) (*Live, error)
	Ask(ctx context.Context, in *TextMessage, opts ...grpc.CallOption) (*TextMessage, error)
}

type getTagsClient struct {
	cc *grpc.ClientConn
}

func NewGetTagsClient(cc *grpc.ClientConn) GetTagsClient {
	return &getTagsClient{cc}
}

func (c *getTagsClient) GetNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Names, error) {
	out := new(Names)
	err := grpc.Invoke(ctx, "/GetTags/GetNames", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTagsClient) Get(ctx context.Context, in *Names, opts ...grpc.CallOption) (*Tags, error) {
	out := new(Tags)
	err := grpc.Invoke(ctx, "/GetTags/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTagsClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Tags, error) {
	out := new(Tags)
	err := grpc.Invoke(ctx, "/GetTags/GetAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTagsClient) SetTag(ctx context.Context, in *ChangeTag, opts ...grpc.CallOption) (*Succesfull, error) {
	out := new(Succesfull)
	err := grpc.Invoke(ctx, "/GetTags/SetTag", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTagsClient) SetTags(ctx context.Context, in *ChangeTags, opts ...grpc.CallOption) (*Succesfull, error) {
	out := new(Succesfull)
	err := grpc.Invoke(ctx, "/GetTags/SetTags", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTagsClient) Register(ctx context.Context, in *ServName, opts ...grpc.CallOption) (*Ident, error) {
	out := new(Ident)
	err := grpc.Invoke(ctx, "/GetTags/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTagsClient) UnRegister(ctx context.Context, in *Ident, opts ...grpc.CallOption) (*Succesfull, error) {
	out := new(Succesfull)
	err := grpc.Invoke(ctx, "/GetTags/UnRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTagsClient) Create(ctx context.Context, in *CreateTag, opts ...grpc.CallOption) (*Succesfull, error) {
	out := new(Succesfull)
	err := grpc.Invoke(ctx, "/GetTags/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTagsClient) IsAlive(ctx context.Context, in *Live, opts ...grpc.CallOption) (*Live, error) {
	out := new(Live)
	err := grpc.Invoke(ctx, "/GetTags/IsAlive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTagsClient) Ask(ctx context.Context, in *TextMessage, opts ...grpc.CallOption) (*TextMessage, error) {
	out := new(TextMessage)
	err := grpc.Invoke(ctx, "/GetTags/Ask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetTags service

type GetTagsServer interface {
	// Чтение тегов без регистрации
	GetNames(context.Context, *Empty) (*Names, error)
	Get(context.Context, *Names) (*Tags, error)
	GetAll(context.Context, *Empty) (*Tags, error)
	SetTag(context.Context, *ChangeTag) (*Succesfull, error)
	SetTags(context.Context, *ChangeTags) (*Succesfull, error)
	Register(context.Context, *ServName) (*Ident, error)
	UnRegister(context.Context, *Ident) (*Succesfull, error)
	Create(context.Context, *CreateTag) (*Succesfull, error)
	IsAlive(context.Context, *Live) (*Live, error)
	Ask(context.Context, *TextMessage) (*TextMessage, error)
}

func RegisterGetTagsServer(s *grpc.Server, srv GetTagsServer) {
	s.RegisterService(&_GetTags_serviceDesc, srv)
}

func _GetTags_GetNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTagsServer).GetNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetTags/GetNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTagsServer).GetNames(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetTags_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Names)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTagsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetTags/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTagsServer).Get(ctx, req.(*Names))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetTags_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTagsServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetTags/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTagsServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetTags_SetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTagsServer).SetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetTags/SetTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTagsServer).SetTag(ctx, req.(*ChangeTag))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetTags_SetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeTags)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTagsServer).SetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetTags/SetTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTagsServer).SetTags(ctx, req.(*ChangeTags))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetTags_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTagsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetTags/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTagsServer).Register(ctx, req.(*ServName))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetTags_UnRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ident)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTagsServer).UnRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetTags/UnRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTagsServer).UnRegister(ctx, req.(*Ident))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetTags_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTagsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetTags/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTagsServer).Create(ctx, req.(*CreateTag))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetTags_IsAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Live)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTagsServer).IsAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetTags/IsAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTagsServer).IsAlive(ctx, req.(*Live))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetTags_Ask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTagsServer).Ask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetTags/Ask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTagsServer).Ask(ctx, req.(*TextMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetTags_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GetTags",
	HandlerType: (*GetTagsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNames",
			Handler:    _GetTags_GetNames_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GetTags_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _GetTags_GetAll_Handler,
		},
		{
			MethodName: "SetTag",
			Handler:    _GetTags_SetTag_Handler,
		},
		{
			MethodName: "SetTags",
			Handler:    _GetTags_SetTags_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _GetTags_Register_Handler,
		},
		{
			MethodName: "UnRegister",
			Handler:    _GetTags_UnRegister_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _GetTags_Create_Handler,
		},
		{
			MethodName: "IsAlive",
			Handler:    _GetTags_IsAlive_Handler,
		},
		{
			MethodName: "Ask",
			Handler:    _GetTags_Ask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x8a, 0x1a, 0x41,
	0x10, 0xc6, 0x1d, 0xe7, 0xcf, 0x8e, 0x65, 0x36, 0x84, 0x66, 0x31, 0x3a, 0x6c, 0xc8, 0xd0, 0x4b,
	0xc0, 0x53, 0x1f, 0xdc, 0x5c, 0x72, 0x34, 0x61, 0x91, 0x0d, 0x49, 0x0e, 0xad, 0x21, 0x90, 0x5b,
	0xab, 0xb5, 0x93, 0x41, 0x9d, 0x91, 0xe9, 0x56, 0x76, 0x1f, 0x33, 0x6f, 0x14, 0xba, 0x7a, 0x1c,
	0x8d, 0x2c, 0xe4, 0x90, 0x93, 0x55, 0x35, 0x55, 0x5f, 0x7d, 0xfd, 0xa3, 0x84, 0x4b, 0x8d, 0xd5,
	0x3e, 0x5f, 0xa0, 0xd8, 0x56, 0xa5, 0x29, 0xf9, 0x6b, 0x08, 0xef, 0x36, 0x5b, 0xf3, 0xc4, 0x5e,
	0x42, 0x3b, 0x5f, 0xf6, 0xbd, 0xd4, 0x1b, 0x86, 0xb2, 0x9d, 0x2f, 0xf9, 0x0f, 0xf0, 0x67, 0x2a,
	0x63, 0x0c, 0x82, 0xa5, 0x32, 0x8a, 0x3e, 0x78, 0x92, 0x62, 0x76, 0x05, 0x61, 0x4e, 0xc5, 0x76,
	0xea, 0x0d, 0x7d, 0xe9, 0x12, 0xdb, 0x69, 0x9e, 0xb6, 0xd8, 0xf7, 0x53, 0x6f, 0x78, 0x29, 0x29,
	0xb6, 0xb5, 0x42, 0x6d, 0xb0, 0x1f, 0xa4, 0xde, 0xb0, 0x23, 0x29, 0xe6, 0x29, 0x04, 0x33, 0x95,
	0x69, 0xd6, 0x87, 0xc0, 0xa8, 0x4c, 0xf7, 0xbd, 0xd4, 0x1f, 0x76, 0x47, 0x81, 0x98, 0xa9, 0x4c,
	0x52, 0x85, 0xbf, 0x81, 0xf0, 0x9b, 0xda, 0xa0, 0xb6, 0x8b, 0xec, 0x88, 0xeb, 0xe9, 0x48, 0x97,
	0xf0, 0x6b, 0x80, 0xe9, 0x6e, 0xb1, 0x40, 0xfd, 0xb0, 0x5b, 0xaf, 0xad, 0xef, 0x72, 0x45, 0xf6,
	0x62, 0xd9, 0x2e, 0x57, 0xfc, 0x16, 0xc2, 0xfb, 0x25, 0x16, 0xa6, 0xd9, 0xed, 0x1d, 0x77, 0xb3,
	0x1e, 0x44, 0xc5, 0x6e, 0x33, 0xc7, 0x8a, 0xac, 0x87, 0xb2, 0xce, 0xf8, 0x7b, 0x88, 0xa7, 0x58,
	0xed, 0xed, 0xd6, 0x67, 0xe7, 0xae, 0x20, 0x5c, 0xa2, 0x5e, 0xb8, 0xb1, 0x8e, 0x74, 0x09, 0xff,
	0x0c, 0xc1, 0x97, 0x7c, 0x7f, 0xaa, 0xea, 0x9d, 0xaa, 0xb2, 0x1e, 0x21, 0xb5, 0x23, 0xdd, 0x51,
	0x24, 0xc8, 0x95, 0x45, 0x6b, 0x37, 0x68, 0x53, 0x6e, 0x89, 0x54, 0x2c, 0x29, 0xe6, 0x63, 0xe8,
	0x7c, 0xfa, 0xa5, 0x8a, 0x0c, 0x2d, 0xf4, 0x1e, 0xf8, 0x46, 0x65, 0xa4, 0x76, 0x20, 0x63, 0x0b,
	0xec, 0xda, 0x82, 0xc7, 0xc2, 0x9c, 0x69, 0xba, 0x22, 0xbf, 0x03, 0x68, 0x24, 0x34, 0x1b, 0x34,
	0x78, 0x6d, 0x6b, 0x68, 0x45, 0xb4, 0xe3, 0xfb, 0x0f, 0x19, 0xeb, 0xa4, 0x42, 0x65, 0xfe, 0xc3,
	0xc9, 0x07, 0xe8, 0xce, 0xf0, 0xd1, 0x7c, 0x45, 0xad, 0x55, 0x86, 0x35, 0x07, 0xef, 0x39, 0x0e,
	0x06, 0x1f, 0x4d, 0x0d, 0x95, 0xe2, 0xd1, 0xef, 0x36, 0x5c, 0x4c, 0xd0, 0xcc, 0x9c, 0xcf, 0x78,
	0x82, 0xc6, 0x9d, 0x42, 0x24, 0xe8, 0x4c, 0x93, 0x48, 0x50, 0xce, 0x5b, 0xd6, 0xda, 0x04, 0x0d,
	0xab, 0x0b, 0x89, 0x7b, 0x21, 0x6f, 0xb1, 0x01, 0x44, 0x13, 0x34, 0xe3, 0xf5, 0xba, 0x99, 0x69,
	0x3e, 0xdd, 0x40, 0x34, 0x25, 0x6d, 0x06, 0xa2, 0x41, 0x95, 0x74, 0xc5, 0xf1, 0x9c, 0x78, 0x8b,
	0xbd, 0x83, 0x8b, 0x69, 0x6d, 0xa0, 0x7b, 0xec, 0xd2, 0xe7, 0x6d, 0x6f, 0x21, 0x96, 0x98, 0xe5,
	0xda, 0x60, 0xc5, 0x3a, 0xe2, 0x70, 0x3d, 0x49, 0xfd, 0x3e, 0x5a, 0x06, 0xdf, 0x8b, 0xa6, 0xa5,
	0xae, 0x9f, 0xab, 0xdc, 0x40, 0xe4, 0x60, 0x5b, 0x47, 0x07, 0xea, 0xe7, 0x4d, 0x03, 0xb8, 0xb8,
	0xd7, 0xe3, 0xb5, 0x3d, 0xb5, 0x50, 0xd8, 0x8b, 0x4b, 0xdc, 0x0f, 0xcd, 0xfb, 0x63, 0xbd, 0x62,
	0x2f, 0xc4, 0x09, 0xef, 0xe4, 0xaf, 0x8c, 0xb7, 0x3e, 0xbe, 0x82, 0x98, 0xfe, 0xec, 0xf3, 0xdd,
	0xc3, 0xcf, 0x60, 0xa3, 0xf2, 0x62, 0x1e, 0x51, 0x7e, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0xed,
	0xad, 0x65, 0xe4, 0x0d, 0x04, 0x00, 0x00,
}
