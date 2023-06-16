// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/docker/swarmkit/api/health.proto

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/docker/swarmkit/protobuf/plugin"

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorHealth, []int{1, 0}
}

type HealthCheckRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{0} }

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=docker.swarmkit.v1.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{1} }

func init() {
	proto.RegisterType((*HealthCheckRequest)(nil), "docker.swarmkit.v1.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "docker.swarmkit.v1.HealthCheckResponse")
	proto.RegisterEnum("docker.swarmkit.v1.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

type authenticatedWrapperHealthServer struct {
	local     HealthServer
	authorize func(context.Context, []string) error
}

func NewAuthenticatedWrapperHealthServer(local HealthServer, authorize func(context.Context, []string) error) HealthServer {
	return &authenticatedWrapperHealthServer{
		local:     local,
		authorize: authorize,
	}
}

func (p *authenticatedWrapperHealthServer) Check(ctx context.Context, r *HealthCheckRequest) (*HealthCheckResponse, error) {

	if err := p.authorize(ctx, []string{"swarm-manager"}); err != nil {
		return nil, err
	}
	return p.local.Check(ctx, r)
}

func (m *HealthCheckRequest) Copy() *HealthCheckRequest {
	if m == nil {
		return nil
	}
	o := &HealthCheckRequest{}
	o.CopyFrom(m)
	return o
}

func (m *HealthCheckRequest) CopyFrom(src interface{}) {

	o := src.(*HealthCheckRequest)
	*m = *o
}

func (m *HealthCheckResponse) Copy() *HealthCheckResponse {
	if m == nil {
		return nil
	}
	o := &HealthCheckResponse{}
	o.CopyFrom(m)
	return o
}

func (m *HealthCheckResponse) CopyFrom(src interface{}) {

	o := src.(*HealthCheckResponse)
	*m = *o
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Health service

type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/docker.swarmkit.v1.Health/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.swarmkit.v1.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "docker.swarmkit.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/docker/swarmkit/api/health.proto",
}

func (m *HealthCheckRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheckRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Service) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHealth(dAtA, i, uint64(len(m.Service)))
		i += copy(dAtA[i:], m.Service)
	}
	return i, nil
}

func (m *HealthCheckResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheckResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHealth(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func encodeVarintHealth(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

type raftProxyHealthServer struct {
	local                       HealthServer
	connSelector                raftselector.ConnProvider
	localCtxMods, remoteCtxMods []func(context.Context) (context.Context, error)
}

func NewRaftProxyHealthServer(local HealthServer, connSelector raftselector.ConnProvider, localCtxMod, remoteCtxMod func(context.Context) (context.Context, error)) HealthServer {
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

	return &raftProxyHealthServer{
		local:         local,
		connSelector:  connSelector,
		localCtxMods:  localMods,
		remoteCtxMods: remoteMods,
	}
}
func (p *raftProxyHealthServer) runCtxMods(ctx context.Context, ctxMods []func(context.Context) (context.Context, error)) (context.Context, error) {
	var err error
	for _, mod := range ctxMods {
		ctx, err = mod(ctx)
		if err != nil {
			return ctx, err
		}
	}
	return ctx, nil
}
func (p *raftProxyHealthServer) pollNewLeaderConn(ctx context.Context) (*grpc.ClientConn, error) {
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

func (p *raftProxyHealthServer) Check(ctx context.Context, r *HealthCheckRequest) (*HealthCheckResponse, error) {

	conn, err := p.connSelector.LeaderConn(ctx)
	if err != nil {
		if err == raftselector.ErrIsLeader {
			ctx, err = p.runCtxMods(ctx, p.localCtxMods)
			if err != nil {
				return nil, err
			}
			return p.local.Check(ctx, r)
		}
		return nil, err
	}
	modCtx, err := p.runCtxMods(ctx, p.remoteCtxMods)
	if err != nil {
		return nil, err
	}

	resp, err := NewHealthClient(conn).Check(modCtx, r)
	if err != nil {
		if !strings.Contains(err.Error(), "is closing") && !strings.Contains(err.Error(), "the connection is unavailable") && !strings.Contains(err.Error(), "connection error") {
			return resp, err
		}
		conn, err := p.pollNewLeaderConn(ctx)
		if err != nil {
			if err == raftselector.ErrIsLeader {
				return p.local.Check(ctx, r)
			}
			return nil, err
		}
		return NewHealthClient(conn).Check(modCtx, r)
	}
	return resp, err
}

func (m *HealthCheckRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovHealth(uint64(l))
	}
	return n
}

func (m *HealthCheckResponse) Size() (n int) {
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovHealth(uint64(m.Status))
	}
	return n
}

func sovHealth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHealth(x uint64) (n int) {
	return sovHealth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HealthCheckRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HealthCheckRequest{`,
		`Service:` + fmt.Sprintf("%v", this.Service) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HealthCheckResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HealthCheckResponse{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHealth(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HealthCheckRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealth
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
			return fmt.Errorf("proto: HealthCheckRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
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
				return ErrInvalidLengthHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealth
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
func (m *HealthCheckResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealth
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
			return fmt.Errorf("proto: HealthCheckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (HealthCheckResponse_ServingStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHealth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealth
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
func skipHealth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHealth
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
					return 0, ErrIntOverflowHealth
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
					return 0, ErrIntOverflowHealth
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
				return 0, ErrInvalidLengthHealth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHealth
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
				next, err := skipHealth(dAtA[start:])
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
	ErrInvalidLengthHealth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHealth   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("github.com/docker/swarmkit/api/health.proto", fileDescriptorHealth) }

var fileDescriptorHealth = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xc9, 0x4f, 0xce, 0x4e, 0x2d, 0xd2, 0x2f, 0x2e,
	0x4f, 0x2c, 0xca, 0xcd, 0xce, 0x2c, 0xd1, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x48, 0x4d, 0xcc, 0x29,
	0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0xa8, 0xd0, 0x83, 0xa9, 0xd0, 0x2b,
	0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xeb, 0x83, 0x58, 0x10, 0x95, 0x52, 0xe6,
	0x78, 0x8c, 0x05, 0xab, 0x48, 0x2a, 0x4d, 0xd3, 0x2f, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x83, 0x52,
	0x10, 0x8d, 0x4a, 0x7a, 0x5c, 0x42, 0x1e, 0x60, 0x2b, 0x9d, 0x33, 0x52, 0x93, 0xb3, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0xa5, 0x05, 0x8c, 0x5c, 0xc2, 0x28, 0x1a,
	0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x7c, 0xb9, 0xd8, 0x8a, 0x4b, 0x12, 0x4b, 0x4a, 0x8b,
	0xc1, 0x1a, 0xf8, 0x8c, 0x4c, 0xf5, 0x30, 0xdd, 0xae, 0x87, 0x45, 0xa3, 0x5e, 0x30, 0xc8, 0xe0,
	0xbc, 0xf4, 0x60, 0xb0, 0xe6, 0x20, 0xa8, 0x21, 0x4a, 0x56, 0x5c, 0xbc, 0x28, 0x12, 0x42, 0xdc,
	0x5c, 0xec, 0xa1, 0x7e, 0xde, 0x7e, 0xfe, 0xe1, 0x7e, 0x02, 0x0c, 0x20, 0x4e, 0xb0, 0x6b, 0x50,
	0x98, 0xa7, 0x9f, 0xbb, 0x00, 0xa3, 0x10, 0x3f, 0x17, 0xb7, 0x9f, 0x7f, 0x48, 0x3c, 0x4c, 0x80,
	0xc9, 0xa8, 0x92, 0x8b, 0x0d, 0x62, 0x91, 0x50, 0x3e, 0x17, 0x2b, 0xd8, 0x32, 0x21, 0x35, 0x82,
	0xae, 0x01, 0xfb, 0x5b, 0x4a, 0x9d, 0x48, 0x57, 0x2b, 0x89, 0x9e, 0x5a, 0xf7, 0x6e, 0x06, 0x13,
	0x3f, 0x17, 0x2f, 0x58, 0xa1, 0x6e, 0x6e, 0x62, 0x5e, 0x62, 0x7a, 0x6a, 0x91, 0x93, 0xc4, 0x89,
	0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x34, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1,
	0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x93, 0xd8, 0xc0, 0xc1, 0x6d, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x7b, 0xf2, 0xdd, 0x23, 0x00, 0x02, 0x00, 0x00,
}
