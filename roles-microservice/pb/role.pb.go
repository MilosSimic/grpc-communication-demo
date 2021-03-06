// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role.proto

/*
Package role is a generated protocol buffer package.

It is generated from these files:
	role.proto

It has these top-level messages:
	EmptyRequest
	GetUserRoleRequest
	RolesReply
	Role
	UserRoleReply
*/
package role

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

// Requests
type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetUserRoleRequest struct {
	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetUserRoleRequest) Reset()                    { *m = GetUserRoleRequest{} }
func (m *GetUserRoleRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRoleRequest) ProtoMessage()               {}
func (*GetUserRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetUserRoleRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

// Replys
type RolesReply struct {
	Roles []*Role `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
}

func (m *RolesReply) Reset()                    { *m = RolesReply{} }
func (m *RolesReply) String() string            { return proto.CompactTextString(m) }
func (*RolesReply) ProtoMessage()               {}
func (*RolesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RolesReply) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Role struct {
	Id   int32  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Role) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserRoleReply struct {
	UserId int32   `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Roles  []*Role `protobuf:"bytes,2,rep,name=roles" json:"roles,omitempty"`
}

func (m *UserRoleReply) Reset()                    { *m = UserRoleReply{} }
func (m *UserRoleReply) String() string            { return proto.CompactTextString(m) }
func (*UserRoleReply) ProtoMessage()               {}
func (*UserRoleReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserRoleReply) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserRoleReply) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "role.EmptyRequest")
	proto.RegisterType((*GetUserRoleRequest)(nil), "role.GetUserRoleRequest")
	proto.RegisterType((*RolesReply)(nil), "role.RolesReply")
	proto.RegisterType((*Role)(nil), "role.Role")
	proto.RegisterType((*UserRoleReply)(nil), "role.UserRoleReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Roles service

type RolesClient interface {
	GetRoles(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RolesReply, error)
	GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*UserRoleReply, error)
}

type rolesClient struct {
	cc *grpc.ClientConn
}

func NewRolesClient(cc *grpc.ClientConn) RolesClient {
	return &rolesClient{cc}
}

func (c *rolesClient) GetRoles(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RolesReply, error) {
	out := new(RolesReply)
	err := grpc.Invoke(ctx, "/role.Roles/GetRoles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*UserRoleReply, error) {
	out := new(UserRoleReply)
	err := grpc.Invoke(ctx, "/role.Roles/GetUserRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Roles service

type RolesServer interface {
	GetRoles(context.Context, *EmptyRequest) (*RolesReply, error)
	GetUserRole(context.Context, *GetUserRoleRequest) (*UserRoleReply, error)
}

func RegisterRolesServer(s *grpc.Server, srv RolesServer) {
	s.RegisterService(&_Roles_serviceDesc, srv)
}

func _Roles_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.Roles/GetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GetRoles(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_GetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.Roles/GetUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GetUserRole(ctx, req.(*GetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Roles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "role.Roles",
	HandlerType: (*RolesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoles",
			Handler:    _Roles_GetRoles_Handler,
		},
		{
			MethodName: "GetUserRole",
			Handler:    _Roles_GetUserRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}

func init() { proto.RegisterFile("role.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x6c, 0x4c, 0x52, 0xe0, 0x0a, 0x15, 0x7a, 0x0c, 0x58, 0x9d, 0x22, 0x4f, 0x15, 0x12, 0x19,
	0xc2, 0xce, 0x86, 0x2a, 0x18, 0x2d, 0x31, 0x23, 0x90, 0xdf, 0x10, 0x29, 0xc5, 0xc1, 0x76, 0x87,
	0x8a, 0x9f, 0x47, 0xb6, 0x91, 0x12, 0x54, 0x65, 0xbb, 0xe7, 0x77, 0xbe, 0xbb, 0x77, 0x80, 0xb3,
	0x3d, 0x37, 0x83, 0xb3, 0xc1, 0x52, 0x19, 0xb1, 0x5a, 0xe3, 0xea, 0x79, 0x3f, 0x84, 0xa3, 0xe6,
	0xef, 0x03, 0xfb, 0xa0, 0x1e, 0x40, 0x3b, 0x0e, 0x6f, 0x9e, 0x9d, 0xb6, 0x3d, 0xff, 0xbd, 0xd2,
	0x1d, 0xce, 0x0f, 0x9e, 0xdd, 0x7b, 0x67, 0x64, 0x51, 0x17, 0xdb, 0x4a, 0x2f, 0xe3, 0xf8, 0x62,
	0x54, 0x03, 0x44, 0x9e, 0xd7, 0x3c, 0xf4, 0x47, 0xaa, 0x51, 0x45, 0x51, 0x2f, 0x8b, 0xfa, 0x6c,
	0xbb, 0x6a, 0xd1, 0x24, 0xbb, 0x24, 0x94, 0x17, 0xea, 0x1e, 0x65, 0x1c, 0x69, 0x0d, 0xd1, 0x19,
	0x29, 0x92, 0x96, 0xe8, 0x0c, 0x11, 0xca, 0xaf, 0x8f, 0x3d, 0x27, 0xf5, 0x4b, 0x9d, 0xb0, 0x7a,
	0xc5, 0xf5, 0x98, 0x23, 0xca, 0xcf, 0xa5, 0x18, 0x7d, 0xc5, 0x8c, 0x6f, 0xfb, 0x83, 0x2a, 0xe5,
	0xa4, 0x16, 0x17, 0x3b, 0x0e, 0x19, 0x53, 0xe6, 0x4d, 0xef, 0xdf, 0xdc, 0x8c, 0x7f, 0xf3, 0x51,
	0x6a, 0x41, 0x4f, 0x58, 0x4d, 0x3a, 0x21, 0x99, 0x29, 0xa7, 0x35, 0x6d, 0x6e, 0xf3, 0xe6, 0x5f,
	0x6a, 0xb5, 0xf8, 0x5c, 0xa6, 0xc2, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xee, 0xf7, 0x3f,
	0x56, 0x7e, 0x01, 0x00, 0x00,
}
