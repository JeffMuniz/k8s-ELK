// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/docker/swarmkit/api/resource.proto

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/docker/swarmkit/protobuf/plugin"

import deepcopy "github.com/docker/swarmkit/api/deepcopy"

import context "golang.org/x/net v0.7.0x/net/context"
import grpc "google.golang.org/x/net v0.7.0grpc"

import raftselector "github.com/docker/swarmkit/manager/raftselector"
import codes "google.golang.org/x/net v0.7.0grpc/codes"
import status "google.golang.org/x/net v0.7.0grpc/status"
import metadata "google.golang.org/x/net v0.7.0grpc/metadata"
import peer "google.golang.org/x/net v0.7.0grpc/peer"
import rafttime "time"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AttachNetworkRequest struct {
	Config      *NetworkAttachmentConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	ContainerID string                   `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
}

func (m *AttachNetworkRequest) Reset()                    { *m = AttachNetworkRequest{} }
func (*AttachNetworkRequest) ProtoMessage()               {}
func (*AttachNetworkRequest) Descriptor() ([]byte, []int) { return fileDescriptorResource, []int{0} }

type AttachNetworkResponse struct {
	AttachmentID string `protobuf:"bytes,1,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
}

func (m *AttachNetworkResponse) Reset()                    { *m = AttachNetworkResponse{} }
func (*AttachNetworkResponse) ProtoMessage()               {}
func (*AttachNetworkResponse) Descriptor() ([]byte, []int) { return fileDescriptorResource, []int{1} }

type DetachNetworkRequest struct {
	AttachmentID string `protobuf:"bytes,1,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
}

func (m *DetachNetworkRequest) Reset()                    { *m = DetachNetworkRequest{} }
func (*DetachNetworkRequest) ProtoMessage()               {}
func (*DetachNetworkRequest) Descriptor() ([]byte, []int) { return fileDescriptorResource, []int{2} }

type DetachNetworkResponse struct {
}

func (m *DetachNetworkResponse) Reset()                    { *m = DetachNetworkResponse{} }
func (*DetachNetworkResponse) ProtoMessage()               {}
func (*DetachNetworkResponse) Descriptor() ([]byte, []int) { return fileDescriptorResource, []int{3} }

func init() {
	proto.RegisterType((*AttachNetworkRequest)(nil), "docker.swarmkit.v1.AttachNetworkRequest")
	proto.RegisterType((*AttachNetworkResponse)(nil), "docker.swarmkit.v1.AttachNetworkResponse")
	proto.RegisterType((*DetachNetworkRequest)(nil), "docker.swarmkit.v1.DetachNetworkRequest")
	proto.RegisterType((*DetachNetworkResponse)(nil), "docker.swarmkit.v1.DetachNetworkResponse")
}

type authenticatedWrapperResourceAllocatorServer struct {
	local     ResourceAllocatorServer
	authorize func(context.Context, []string) error
}

func NewAuthenticatedWrapperResourceAllocatorServer(local ResourceAllocatorServer, authorize func(context.Context, []string) error) ResourceAllocatorServer {
	return &authenticatedWrapperResourceAllocatorServer{
		local:     local,
		authorize: authorize,
	}
}

func (p *authenticatedWrapperResourceAllocatorServer) AttachNetwork(ctx context.Context, r *AttachNetworkRequest) (*AttachNetworkResponse, error) {

	if err := p.authorize(ctx, []string{"swarm-worker", "swarm-manager"}); err != nil {
		return nil, err
	}
	return p.local.AttachNetwork(ctx, r)
}

func (p *authenticatedWrapperResourceAllocatorServer) DetachNetwork(ctx context.Context, r *DetachNetworkRequest) (*DetachNetworkResponse, error) {

	if err := p.authorize(ctx, []string{"swarm-worker", "swarm-manager"}); err != nil {
		return nil, err
	}
	return p.local.DetachNetwork(ctx, r)
}

func (m *AttachNetworkRequest) Copy() *AttachNetworkRequest {
	if m == nil {
		return nil
	}
	o := &AttachNetworkRequest{}
	o.CopyFrom(m)
	return o
}

func (m *AttachNetworkRequest) CopyFrom(src interface{}) {

	o := src.(*AttachNetworkRequest)
	*m = *o
	if o.Config != nil {
		m.Config = &NetworkAttachmentConfig{}
		deepcopy.Copy(m.Config, o.Config)
	}
}

func (m *AttachNetworkResponse) Copy() *AttachNetworkResponse {
	if m == nil {
		return nil
	}
	o := &AttachNetworkResponse{}
	o.CopyFrom(m)
	return o
}

func (m *AttachNetworkResponse) CopyFrom(src interface{}) {

	o := src.(*AttachNetworkResponse)
	*m = *o
}

func (m *DetachNetworkRequest) Copy() *DetachNetworkRequest {
	if m == nil {
		return nil
	}
	o := &DetachNetworkRequest{}
	o.CopyFrom(m)
	return o
}

func (m *DetachNetworkRequest) CopyFrom(src interface{}) {

	o := src.(*DetachNetworkRequest)
	*m = *o
}

func (m *DetachNetworkResponse) Copy() *DetachNetworkResponse {
	if m == nil {
		return nil
	}
	o := &DetachNetworkResponse{}
	o.CopyFrom(m)
	return o
}

func (m *DetachNetworkResponse) CopyFrom(src interface{}) {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceAllocator service

type ResourceAllocatorClient interface {
	AttachNetwork(ctx context.Context, in *AttachNetworkRequest, opts ...grpc.CallOption) (*AttachNetworkResponse, error)
	DetachNetwork(ctx context.Context, in *DetachNetworkRequest, opts ...grpc.CallOption) (*DetachNetworkResponse, error)
}

type resourceAllocatorClient struct {
	cc *grpc.ClientConn
}

func NewResourceAllocatorClient(cc *grpc.ClientConn) ResourceAllocatorClient {
	return &resourceAllocatorClient{cc}
}

func (c *resourceAllocatorClient) AttachNetwork(ctx context.Context, in *AttachNetworkRequest, opts ...grpc.CallOption) (*AttachNetworkResponse, error) {
	out := new(AttachNetworkResponse)
	err := grpc.Invoke(ctx, "/docker.swarmkit.v1.ResourceAllocator/AttachNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAllocatorClient) DetachNetwork(ctx context.Context, in *DetachNetworkRequest, opts ...grpc.CallOption) (*DetachNetworkResponse, error) {
	out := new(DetachNetworkResponse)
	err := grpc.Invoke(ctx, "/docker.swarmkit.v1.ResourceAllocator/DetachNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceAllocator service

type ResourceAllocatorServer interface {
	AttachNetwork(context.Context, *AttachNetworkRequest) (*AttachNetworkResponse, error)
	DetachNetwork(context.Context, *DetachNetworkRequest) (*DetachNetworkResponse, error)
}

func RegisterResourceAllocatorServer(s *grpc.Server, srv ResourceAllocatorServer) {
	s.RegisterService(&_ResourceAllocator_serviceDesc, srv)
}

func _ResourceAllocator_AttachNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAllocatorServer).AttachNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.swarmkit.v1.ResourceAllocator/AttachNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAllocatorServer).AttachNetwork(ctx, req.(*AttachNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAllocator_DetachNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetachNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAllocatorServer).DetachNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.swarmkit.v1.ResourceAllocator/DetachNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAllocatorServer).DetachNetwork(ctx, req.(*DetachNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceAllocator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "docker.swarmkit.v1.ResourceAllocator",
	HandlerType: (*ResourceAllocatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AttachNetwork",
			Handler:    _ResourceAllocator_AttachNetwork_Handler,
		},
		{
			MethodName: "DetachNetwork",
			Handler:    _ResourceAllocator_DetachNetwork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/docker/swarmkit/api/resource.proto",
}

func (m *AttachNetworkRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttachNetworkRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResource(dAtA, i, uint64(m.Config.Size()))
		n1, err := m.Config.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.ContainerID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintResource(dAtA, i, uint64(len(m.ContainerID)))
		i += copy(dAtA[i:], m.ContainerID)
	}
	return i, nil
}

func (m *AttachNetworkResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttachNetworkResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AttachmentID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResource(dAtA, i, uint64(len(m.AttachmentID)))
		i += copy(dAtA[i:], m.AttachmentID)
	}
	return i, nil
}

func (m *DetachNetworkRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DetachNetworkRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AttachmentID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResource(dAtA, i, uint64(len(m.AttachmentID)))
		i += copy(dAtA[i:], m.AttachmentID)
	}
	return i, nil
}

func (m *DetachNetworkResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DetachNetworkResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

type raftProxyResourceAllocatorServer struct {
	local                       ResourceAllocatorServer
	connSelector                raftselector.ConnProvider
	localCtxMods, remoteCtxMods []func(context.Context) (context.Context, error)
}

func NewRaftProxyResourceAllocatorServer(local ResourceAllocatorServer, connSelector raftselector.ConnProvider, localCtxMod, remoteCtxMod func(context.Context) (context.Context, error)) ResourceAllocatorServer {
	redirectChecker := func(ctx context.Context) (context.Context, error) {
		p, ok := peer.FromContext(ctx)
		if !ok {
			return ctx, status.Errorf(codes.InvalidArgument, "remote addr is not found in context")
		}
		addr := p.Addr.String()
		md, ok := metadata.FromIncomingContext(ctx)
		if ok && len(md["redirect"]) != 0 {
			return ctx, status.Errorf(codes.ResourceExhausted, "more than one redirect to leader from: %s", md["redirect"])
		}
		if !ok {
			md = metadata.New(map[string]string{})
		}
		md["redirect"] = append(md["redirect"], addr)
		return metadata.NewOutgoingContext(ctx, md), nil
	}
	remoteMods := []func(context.Context) (context.Context, error){redirectChecker}
	remoteMods = append(remoteMods, remoteCtxMod)

	var localMods []func(context.Context) (context.Context, error)
	if localCtxMod != nil {
		localMods = []func(context.Context) (context.Context, error){localCtxMod}
	}

	return &raftProxyResourceAllocatorServer{
		local:         local,
		connSelector:  connSelector,
		localCtxMods:  localMods,
		remoteCtxMods: remoteMods,
	}
}
func (p *raftProxyResourceAllocatorServer) runCtxMods(ctx context.Context, ctxMods []func(context.Context) (context.Context, error)) (context.Context, error) {
	var err error
	for _, mod := range ctxMods {
		ctx, err = mod(ctx)
		if err != nil {
			return ctx, err
		}
	}
	return ctx, nil
}
func (p *raftProxyResourceAllocatorServer) pollNewLeaderConn(ctx context.Context) (*grpc.ClientConn, error) {
	ticker := rafttime.NewTicker(500 * rafttime.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			conn, err := p.connSelector.LeaderConn(ctx)
			if err != nil {
				return nil, err
			}

			client := NewHealthClient(conn)

			resp, err := client.Check(ctx, &HealthCheckRequest{Service: "Raft"})
			if err != nil || resp.Status != HealthCheckResponse_SERVING {
				continue
			}
			return conn, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (p *raftProxyResourceAllocatorServer) AttachNetwork(ctx context.Context, r *AttachNetworkRequest) (*AttachNetworkResponse, error) {

	conn, err := p.connSelector.LeaderConn(ctx)
	if err != nil {
		if err == raftselector.ErrIsLeader {
			ctx, err = p.runCtxMods(ctx, p.localCtxMods)
			if err != nil {
				return nil, err
			}
			return p.local.AttachNetwork(ctx, r)
		}
		return nil, err
	}
	modCtx, err := p.runCtxMods(ctx, p.remoteCtxMods)
	if err != nil {
		return nil, err
	}

	resp, err := NewResourceAllocatorClient(conn).AttachNetwork(modCtx, r)
	if err != nil {
		if !strings.Contains(err.Error(), "is closing") && !strings.Contains(err.Error(), "the connection is unavailable") && !strings.Contains(err.Error(), "connection error") {
			return resp, err
		}
		conn, err := p.pollNewLeaderConn(ctx)
		if err != nil {
			if err == raftselector.ErrIsLeader {
				return p.local.AttachNetwork(ctx, r)
			}
			return nil, err
		}
		return NewResourceAllocatorClient(conn).AttachNetwork(modCtx, r)
	}
	return resp, err
}

func (p *raftProxyResourceAllocatorServer) DetachNetwork(ctx context.Context, r *DetachNetworkRequest) (*DetachNetworkResponse, error) {

	conn, err := p.connSelector.LeaderConn(ctx)
	if err != nil {
		if err == raftselector.ErrIsLeader {
			ctx, err = p.runCtxMods(ctx, p.localCtxMods)
			if err != nil {
				return nil, err
			}
			return p.local.DetachNetwork(ctx, r)
		}
		return nil, err
	}
	modCtx, err := p.runCtxMods(ctx, p.remoteCtxMods)
	if err != nil {
		return nil, err
	}

	resp, err := NewResourceAllocatorClient(conn).DetachNetwork(modCtx, r)
	if err != nil {
		if !strings.Contains(err.Error(), "is closing") && !strings.Contains(err.Error(), "the connection is unavailable") && !strings.Contains(err.Error(), "connection error") {
			return resp, err
		}
		conn, err := p.pollNewLeaderConn(ctx)
		if err != nil {
			if err == raftselector.ErrIsLeader {
				return p.local.DetachNetwork(ctx, r)
			}
			return nil, err
		}
		return NewResourceAllocatorClient(conn).DetachNetwork(modCtx, r)
	}
	return resp, err
}

func (m *AttachNetworkRequest) Size() (n int) {
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func (m *AttachNetworkResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.AttachmentID)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func (m *DetachNetworkRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.AttachmentID)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func (m *DetachNetworkResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovResource(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AttachNetworkRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AttachNetworkRequest{`,
		`Config:` + strings.Replace(fmt.Sprintf("%v", this.Config), "NetworkAttachmentConfig", "NetworkAttachmentConfig", 1) + `,`,
		`ContainerID:` + fmt.Sprintf("%v", this.ContainerID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AttachNetworkResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AttachNetworkResponse{`,
		`AttachmentID:` + fmt.Sprintf("%v", this.AttachmentID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DetachNetworkRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DetachNetworkRequest{`,
		`AttachmentID:` + fmt.Sprintf("%v", this.AttachmentID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DetachNetworkResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DetachNetworkResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringResource(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AttachNetworkRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AttachNetworkRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AttachNetworkRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &NetworkAttachmentConfig{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AttachNetworkResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AttachNetworkResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AttachNetworkResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttachmentID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DetachNetworkRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DetachNetworkRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DetachNetworkRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttachmentID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DetachNetworkResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DetachNetworkResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DetachNetworkResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResource
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResource
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthResource
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowResource
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipResource(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthResource = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/docker/swarmkit/api/resource.proto", fileDescriptorResource)
}

var fileDescriptorResource = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x19, 0x16, 0x24, 0xdf, 0x50, 0xf2, 0x69, 0x03, 0x91, 0x90, 0x58, 0x48, 0xdd, 0xa0,
	0x86, 0x36, 0x62, 0x8c, 0x6b, 0xfe, 0x6c, 0xba, 0x90, 0x45, 0x5f, 0xc0, 0x0c, 0xed, 0x50, 0x1a,
	0x68, 0xa7, 0x4e, 0xa7, 0x12, 0x77, 0x6e, 0x5d, 0xb9, 0xf5, 0x1d, 0x4c, 0x7c, 0x0e, 0xe2, 0xca,
	0xa5, 0x2b, 0x22, 0x7d, 0x00, 0x9f, 0xc1, 0xd0, 0x29, 0x10, 0x70, 0xa2, 0xc4, 0x55, 0xa7, 0xd3,
	0x73, 0xce, 0xfd, 0xdd, 0x7b, 0x0b, 0x1b, 0x8e, 0xcb, 0x86, 0x51, 0x5f, 0xb3, 0x88, 0xa7, 0xdb,
	0xc4, 0x1a, 0x61, 0xaa, 0x87, 0x13, 0x44, 0xbd, 0x91, 0xcb, 0x74, 0x14, 0xb8, 0x3a, 0xc5, 0x21,
	0x89, 0xa8, 0x85, 0xb5, 0x80, 0x12, 0x46, 0x64, 0x99, 0x6b, 0xb4, 0xa5, 0x46, 0xbb, 0x3d, 0xab,
	0x9c, 0xfc, 0x12, 0xc1, 0xee, 0x02, 0x1c, 0x72, 0x7f, 0xa5, 0xe8, 0x10, 0x87, 0x24, 0x47, 0x7d,
	0x71, 0x4a, 0x6f, 0x2f, 0x7f, 0x48, 0x48, 0x14, 0xfd, 0x68, 0xa0, 0x07, 0xe3, 0xc8, 0x71, 0xfd,
	0xf4, 0xc1, 0x8d, 0xea, 0x23, 0x80, 0xc5, 0x16, 0x63, 0xc8, 0x1a, 0xf6, 0x30, 0x9b, 0x10, 0x3a,
	0x32, 0xf1, 0x4d, 0x84, 0x43, 0x26, 0x77, 0x60, 0xce, 0x22, 0xfe, 0xc0, 0x75, 0xca, 0xa0, 0x06,
	0xea, 0xf9, 0xe6, 0xa9, 0xf6, 0x1d, 0x5c, 0x4b, 0x3d, 0x3c, 0xc0, 0xc3, 0x3e, 0xeb, 0x24, 0x16,
	0x33, 0xb5, 0xca, 0x4d, 0x28, 0x59, 0xc4, 0x67, 0xc8, 0xf5, 0x31, 0xbd, 0x76, 0xed, 0x72, 0xb6,
	0x06, 0xea, 0xff, 0xda, 0xff, 0xe3, 0x59, 0x35, 0xdf, 0x59, 0xde, 0x1b, 0x5d, 0x33, 0xbf, 0x12,
	0x19, 0xb6, 0xda, 0x83, 0xa5, 0x2d, 0xa0, 0x30, 0x20, 0x7e, 0x88, 0xe5, 0x0b, 0x58, 0x40, 0xab,
	0x42, 0x8b, 0x34, 0x90, 0xa4, 0xed, 0xc5, 0xb3, 0xaa, 0xb4, 0x26, 0x30, 0xba, 0xa6, 0xb4, 0x96,
	0x19, 0xb6, 0x7a, 0x05, 0x8b, 0x5d, 0x2c, 0x68, 0xf0, 0x8f, 0x71, 0x07, 0xb0, 0xb4, 0x15, 0xc7,
	0xf1, 0x9a, 0xcf, 0x59, 0xb8, 0x6f, 0xa6, 0xbb, 0x6e, 0x8d, 0xc7, 0xc4, 0x42, 0x8c, 0x50, 0xf9,
	0x01, 0xc0, 0xc2, 0x46, 0x3b, 0x72, 0x5d, 0x34, 0x48, 0xd1, 0x0a, 0x2a, 0xc7, 0x3b, 0x28, 0x79,
	0x71, 0xf5, 0xe8, 0xf5, 0xe5, 0xf3, 0x29, 0x7b, 0x08, 0xa5, 0x44, 0xda, 0x58, 0x7c, 0xc3, 0x14,
	0x16, 0xf8, 0x9b, 0x87, 0x7c, 0xe4, 0x60, 0xce, 0xb2, 0xc1, 0x2e, 0x66, 0x11, 0x4d, 0x4b, 0xcc,
	0x22, 0x1c, 0xc4, 0x4e, 0x2c, 0xed, 0xf2, 0x74, 0xae, 0x64, 0xde, 0xe7, 0x4a, 0xe6, 0x3e, 0x56,
	0xc0, 0x34, 0x56, 0xc0, 0x5b, 0xac, 0x80, 0x8f, 0x58, 0x01, 0xfd, 0x5c, 0xf2, 0x63, 0x9e, 0x7f,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x7a, 0x29, 0xfc, 0x58, 0x03, 0x00, 0x00,
}
