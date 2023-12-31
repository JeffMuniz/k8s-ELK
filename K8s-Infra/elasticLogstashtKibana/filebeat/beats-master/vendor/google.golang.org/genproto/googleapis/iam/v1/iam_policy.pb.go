// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/iam/v1/iam_policy.proto

package iam

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/x/net v0.7.0genproto/googleapis/api/annotations"
	grpc "google.golang.org/x/net v0.7.0grpc"
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

// Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	// REQUIRED: The resource for which the policy is being specified.
	// See the operation documentation for the appropriate value for this field.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// REQUIRED: The complete policy to be applied to the `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as Projects)
	// might reject them.
	Policy               *Policy  `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetIamPolicyRequest) Reset()         { *m = SetIamPolicyRequest{} }
func (m *SetIamPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*SetIamPolicyRequest) ProtoMessage()    {}
func (*SetIamPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2728eb97d748a32, []int{0}
}

func (m *SetIamPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetIamPolicyRequest.Unmarshal(m, b)
}
func (m *SetIamPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetIamPolicyRequest.Marshal(b, m, deterministic)
}
func (m *SetIamPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetIamPolicyRequest.Merge(m, src)
}
func (m *SetIamPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_SetIamPolicyRequest.Size(m)
}
func (m *SetIamPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetIamPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetIamPolicyRequest proto.InternalMessageInfo

func (m *SetIamPolicyRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *SetIamPolicyRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// Request message for `GetIamPolicy` method.
type GetIamPolicyRequest struct {
	// REQUIRED: The resource for which the policy is being requested.
	// See the operation documentation for the appropriate value for this field.
	Resource             string   `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIamPolicyRequest) Reset()         { *m = GetIamPolicyRequest{} }
func (m *GetIamPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*GetIamPolicyRequest) ProtoMessage()    {}
func (*GetIamPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2728eb97d748a32, []int{1}
}

func (m *GetIamPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIamPolicyRequest.Unmarshal(m, b)
}
func (m *GetIamPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIamPolicyRequest.Marshal(b, m, deterministic)
}
func (m *GetIamPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIamPolicyRequest.Merge(m, src)
}
func (m *GetIamPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_GetIamPolicyRequest.Size(m)
}
func (m *GetIamPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIamPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIamPolicyRequest proto.InternalMessageInfo

func (m *GetIamPolicyRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

// Request message for `TestIamPermissions` method.
type TestIamPermissionsRequest struct {
	// REQUIRED: The resource for which the policy detail is being requested.
	// See the operation documentation for the appropriate value for this field.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The set of permissions to check for the `resource`. Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For more
	// information see
	// [IAM Overview](https://cloud.google.com/iam/docs/overview#permissions).
	Permissions          []string `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestIamPermissionsRequest) Reset()         { *m = TestIamPermissionsRequest{} }
func (m *TestIamPermissionsRequest) String() string { return proto.CompactTextString(m) }
func (*TestIamPermissionsRequest) ProtoMessage()    {}
func (*TestIamPermissionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2728eb97d748a32, []int{2}
}

func (m *TestIamPermissionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestIamPermissionsRequest.Unmarshal(m, b)
}
func (m *TestIamPermissionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestIamPermissionsRequest.Marshal(b, m, deterministic)
}
func (m *TestIamPermissionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestIamPermissionsRequest.Merge(m, src)
}
func (m *TestIamPermissionsRequest) XXX_Size() int {
	return xxx_messageInfo_TestIamPermissionsRequest.Size(m)
}
func (m *TestIamPermissionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestIamPermissionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestIamPermissionsRequest proto.InternalMessageInfo

func (m *TestIamPermissionsRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *TestIamPermissionsRequest) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// Response message for `TestIamPermissions` method.
type TestIamPermissionsResponse struct {
	// A subset of `TestPermissionsRequest.permissions` that the caller is
	// allowed.
	Permissions          []string `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestIamPermissionsResponse) Reset()         { *m = TestIamPermissionsResponse{} }
func (m *TestIamPermissionsResponse) String() string { return proto.CompactTextString(m) }
func (*TestIamPermissionsResponse) ProtoMessage()    {}
func (*TestIamPermissionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2728eb97d748a32, []int{3}
}

func (m *TestIamPermissionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestIamPermissionsResponse.Unmarshal(m, b)
}
func (m *TestIamPermissionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestIamPermissionsResponse.Marshal(b, m, deterministic)
}
func (m *TestIamPermissionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestIamPermissionsResponse.Merge(m, src)
}
func (m *TestIamPermissionsResponse) XXX_Size() int {
	return xxx_messageInfo_TestIamPermissionsResponse.Size(m)
}
func (m *TestIamPermissionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestIamPermissionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestIamPermissionsResponse proto.InternalMessageInfo

func (m *TestIamPermissionsResponse) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterType((*SetIamPolicyRequest)(nil), "google.iam.v1.SetIamPolicyRequest")
	proto.RegisterType((*GetIamPolicyRequest)(nil), "google.iam.v1.GetIamPolicyRequest")
	proto.RegisterType((*TestIamPermissionsRequest)(nil), "google.iam.v1.TestIamPermissionsRequest")
	proto.RegisterType((*TestIamPermissionsResponse)(nil), "google.iam.v1.TestIamPermissionsResponse")
}

func init() { proto.RegisterFile("google/iam/v1/iam_policy.proto", fileDescriptor_d2728eb97d748a32) }

var fileDescriptor_d2728eb97d748a32 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0x4c, 0xcc, 0xd5, 0x2f, 0x33, 0x04, 0x51, 0xf1, 0x05, 0xf9, 0x39, 0x99,
	0xc9, 0x95, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x10, 0x79, 0xbd, 0xcc, 0xc4, 0x5c,
	0xbd, 0x32, 0x43, 0x29, 0x49, 0xa8, 0xf2, 0xc4, 0x82, 0x4c, 0xfd, 0xa2, 0xd4, 0xe2, 0xfc, 0xd2,
	0xa2, 0xe4, 0x54, 0x88, 0x4a, 0x29, 0x29, 0x54, 0x93, 0x90, 0x4d, 0x91, 0x92, 0x41, 0xd2, 0x96,
	0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x91, 0x55, 0x4a, 0xe0, 0x12,
	0x0e, 0x4e, 0x2d, 0xf1, 0x4c, 0xcc, 0x0d, 0x00, 0xeb, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x92, 0xe2, 0xe2, 0x80, 0x59, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x0b,
	0xe9, 0x72, 0xb1, 0x41, 0x2c, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd5, 0x43, 0x71,
	0xa7, 0x1e, 0xd4, 0x24, 0xa8, 0x22, 0x25, 0x43, 0x2e, 0x61, 0x77, 0xd2, 0x6c, 0x50, 0x8a, 0xe4,
	0x92, 0x0c, 0x49, 0x2d, 0x06, 0xeb, 0x49, 0x2d, 0xca, 0xcd, 0x2c, 0x2e, 0x06, 0x39, 0x98, 0x18,
	0xa7, 0x29, 0x70, 0x71, 0x17, 0x20, 0x74, 0x48, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x06, 0x21, 0x0b,
	0x29, 0xd9, 0x71, 0x49, 0x61, 0x33, 0xba, 0xb8, 0x20, 0x3f, 0xaf, 0x18, 0x43, 0x3f, 0x23, 0x86,
	0x7e, 0xa3, 0x29, 0xcc, 0x5c, 0x9c, 0x9e, 0x8e, 0xbe, 0x10, 0xbf, 0x08, 0x95, 0x70, 0xf1, 0x20,
	0x87, 0x9e, 0x90, 0x12, 0x5a, 0x50, 0x60, 0x09, 0x5a, 0x29, 0xec, 0xc1, 0xa5, 0xa4, 0xd9, 0x74,
	0xf9, 0xc9, 0x64, 0x26, 0x65, 0x25, 0x39, 0x50, 0x04, 0x56, 0xc3, 0x7c, 0x64, 0xab, 0xa5, 0x55,
	0x6b, 0x55, 0x8c, 0x64, 0x8a, 0x15, 0xa3, 0x16, 0xc8, 0x56, 0x77, 0x7c, 0xb6, 0xba, 0x53, 0xc5,
	0xd6, 0x74, 0x34, 0x5b, 0x67, 0x31, 0x72, 0x09, 0x61, 0x06, 0x9d, 0x90, 0x06, 0x9a, 0xc1, 0x38,
	0x23, 0x4e, 0x4a, 0x93, 0x08, 0x95, 0x90, 0x78, 0x50, 0xd2, 0x07, 0x3b, 0x4b, 0x53, 0x49, 0x05,
	0xd3, 0x59, 0x25, 0x18, 0xba, 0xac, 0x18, 0xb5, 0x9c, 0xda, 0x18, 0xb9, 0x04, 0x93, 0xf3, 0x73,
	0x51, 0x6d, 0x70, 0xe2, 0x83, 0x7b, 0x20, 0x00, 0x94, 0xd8, 0x03, 0x18, 0xa3, 0x0c, 0xa0, 0x0a,
	0xd2, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2, 0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0x59,
	0x41, 0x1f, 0x22, 0x95, 0x58, 0x90, 0x59, 0x0c, 0xcd, 0x47, 0xd6, 0x99, 0x89, 0xb9, 0x3f, 0x18,
	0x19, 0x57, 0x31, 0x09, 0xbb, 0x43, 0x74, 0x39, 0xe7, 0xe4, 0x97, 0xa6, 0xe8, 0x79, 0x26, 0xe6,
	0xea, 0x85, 0x19, 0x9e, 0x82, 0x89, 0xc6, 0x80, 0x45, 0x63, 0x3c, 0x13, 0x73, 0x63, 0xc2, 0x0c,
	0x93, 0xd8, 0xc0, 0x66, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xba, 0x05, 0xa3, 0xc3, 0xdc,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IAMPolicyClient is the client API for IAMPolicy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/x/net v0.7.0grpc#ClientConn.NewStream.
type IAMPolicyClient interface {
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	SetIamPolicy(ctx context.Context, in *SetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Gets the access control policy for a resource.
	// Returns an empty policy if the resource exists and does not have a policy
	// set.
	GetIamPolicy(ctx context.Context, in *GetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// If the resource does not exist, this will return an empty set of
	// permissions, not a NOT_FOUND error.
	//
	// Note: This operation is designed to be used for building permission-aware
	// UIs and command-line tools, not for authorization checking. This operation
	// may "fail open" without warning.
	TestIamPermissions(ctx context.Context, in *TestIamPermissionsRequest, opts ...grpc.CallOption) (*TestIamPermissionsResponse, error)
}

type iAMPolicyClient struct {
	cc *grpc.ClientConn
}

func NewIAMPolicyClient(cc *grpc.ClientConn) IAMPolicyClient {
	return &iAMPolicyClient{cc}
}

func (c *iAMPolicyClient) SetIamPolicy(ctx context.Context, in *SetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/google.iam.v1.IAMPolicy/SetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMPolicyClient) GetIamPolicy(ctx context.Context, in *GetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/google.iam.v1.IAMPolicy/GetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMPolicyClient) TestIamPermissions(ctx context.Context, in *TestIamPermissionsRequest, opts ...grpc.CallOption) (*TestIamPermissionsResponse, error) {
	out := new(TestIamPermissionsResponse)
	err := c.cc.Invoke(ctx, "/google.iam.v1.IAMPolicy/TestIamPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMPolicyServer is the server API for IAMPolicy service.
type IAMPolicyServer interface {
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	SetIamPolicy(context.Context, *SetIamPolicyRequest) (*Policy, error)
	// Gets the access control policy for a resource.
	// Returns an empty policy if the resource exists and does not have a policy
	// set.
	GetIamPolicy(context.Context, *GetIamPolicyRequest) (*Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// If the resource does not exist, this will return an empty set of
	// permissions, not a NOT_FOUND error.
	//
	// Note: This operation is designed to be used for building permission-aware
	// UIs and command-line tools, not for authorization checking. This operation
	// may "fail open" without warning.
	TestIamPermissions(context.Context, *TestIamPermissionsRequest) (*TestIamPermissionsResponse, error)
}

func RegisterIAMPolicyServer(s *grpc.Server, srv IAMPolicyServer) {
	s.RegisterService(&_IAMPolicy_serviceDesc, srv)
}

func _IAMPolicy_SetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMPolicyServer).SetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.v1.IAMPolicy/SetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMPolicyServer).SetIamPolicy(ctx, req.(*SetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMPolicy_GetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMPolicyServer).GetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.v1.IAMPolicy/GetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMPolicyServer).GetIamPolicy(ctx, req.(*GetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMPolicy_TestIamPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMPolicyServer).TestIamPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.v1.IAMPolicy/TestIamPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMPolicyServer).TestIamPermissions(ctx, req.(*TestIamPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IAMPolicy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.iam.v1.IAMPolicy",
	HandlerType: (*IAMPolicyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetIamPolicy",
			Handler:    _IAMPolicy_SetIamPolicy_Handler,
		},
		{
			MethodName: "GetIamPolicy",
			Handler:    _IAMPolicy_GetIamPolicy_Handler,
		},
		{
			MethodName: "TestIamPermissions",
			Handler:    _IAMPolicy_TestIamPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/iam/v1/iam_policy.proto",
}
