// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/monitoring/metricsscope/v1/metrics_scopes.proto

package metricsscopepb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetricsScopesClient is the client API for MetricsScopes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsScopesClient interface {
	// Returns a specific `Metrics Scope`.
	GetMetricsScope(ctx context.Context, in *GetMetricsScopeRequest, opts ...grpc.CallOption) (*MetricsScope, error)
	// Returns a list of every `Metrics Scope` that a specific `MonitoredProject`
	// has been added to. The metrics scope representing the specified monitored
	// project will always be the first entry in the response.
	ListMetricsScopesByMonitoredProject(ctx context.Context, in *ListMetricsScopesByMonitoredProjectRequest, opts ...grpc.CallOption) (*ListMetricsScopesByMonitoredProjectResponse, error)
	// Adds a `MonitoredProject` with the given project ID
	// to the specified `Metrics Scope`.
	CreateMonitoredProject(ctx context.Context, in *CreateMonitoredProjectRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a `MonitoredProject` from the specified `Metrics Scope`.
	DeleteMonitoredProject(ctx context.Context, in *DeleteMonitoredProjectRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type metricsScopesClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsScopesClient(cc grpc.ClientConnInterface) MetricsScopesClient {
	return &metricsScopesClient{cc}
}

func (c *metricsScopesClient) GetMetricsScope(ctx context.Context, in *GetMetricsScopeRequest, opts ...grpc.CallOption) (*MetricsScope, error) {
	out := new(MetricsScope)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.metricsscope.v1.MetricsScopes/GetMetricsScope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsScopesClient) ListMetricsScopesByMonitoredProject(ctx context.Context, in *ListMetricsScopesByMonitoredProjectRequest, opts ...grpc.CallOption) (*ListMetricsScopesByMonitoredProjectResponse, error) {
	out := new(ListMetricsScopesByMonitoredProjectResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.metricsscope.v1.MetricsScopes/ListMetricsScopesByMonitoredProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsScopesClient) CreateMonitoredProject(ctx context.Context, in *CreateMonitoredProjectRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.metricsscope.v1.MetricsScopes/CreateMonitoredProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsScopesClient) DeleteMonitoredProject(ctx context.Context, in *DeleteMonitoredProjectRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.metricsscope.v1.MetricsScopes/DeleteMonitoredProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsScopesServer is the server API for MetricsScopes service.
// All implementations must embed UnimplementedMetricsScopesServer
// for forward compatibility
type MetricsScopesServer interface {
	// Returns a specific `Metrics Scope`.
	GetMetricsScope(context.Context, *GetMetricsScopeRequest) (*MetricsScope, error)
	// Returns a list of every `Metrics Scope` that a specific `MonitoredProject`
	// has been added to. The metrics scope representing the specified monitored
	// project will always be the first entry in the response.
	ListMetricsScopesByMonitoredProject(context.Context, *ListMetricsScopesByMonitoredProjectRequest) (*ListMetricsScopesByMonitoredProjectResponse, error)
	// Adds a `MonitoredProject` with the given project ID
	// to the specified `Metrics Scope`.
	CreateMonitoredProject(context.Context, *CreateMonitoredProjectRequest) (*longrunningpb.Operation, error)
	// Deletes a `MonitoredProject` from the specified `Metrics Scope`.
	DeleteMonitoredProject(context.Context, *DeleteMonitoredProjectRequest) (*longrunningpb.Operation, error)
	mustEmbedUnimplementedMetricsScopesServer()
}

// UnimplementedMetricsScopesServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsScopesServer struct {
}

func (UnimplementedMetricsScopesServer) GetMetricsScope(context.Context, *GetMetricsScopeRequest) (*MetricsScope, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricsScope not implemented")
}
func (UnimplementedMetricsScopesServer) ListMetricsScopesByMonitoredProject(context.Context, *ListMetricsScopesByMonitoredProjectRequest) (*ListMetricsScopesByMonitoredProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetricsScopesByMonitoredProject not implemented")
}
func (UnimplementedMetricsScopesServer) CreateMonitoredProject(context.Context, *CreateMonitoredProjectRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMonitoredProject not implemented")
}
func (UnimplementedMetricsScopesServer) DeleteMonitoredProject(context.Context, *DeleteMonitoredProjectRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMonitoredProject not implemented")
}
func (UnimplementedMetricsScopesServer) mustEmbedUnimplementedMetricsScopesServer() {}

// UnsafeMetricsScopesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsScopesServer will
// result in compilation errors.
type UnsafeMetricsScopesServer interface {
	mustEmbedUnimplementedMetricsScopesServer()
}

func RegisterMetricsScopesServer(s grpc.ServiceRegistrar, srv MetricsScopesServer) {
	s.RegisterService(&MetricsScopes_ServiceDesc, srv)
}

func _MetricsScopes_GetMetricsScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsScopesServer).GetMetricsScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.metricsscope.v1.MetricsScopes/GetMetricsScope",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsScopesServer).GetMetricsScope(ctx, req.(*GetMetricsScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsScopes_ListMetricsScopesByMonitoredProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetricsScopesByMonitoredProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsScopesServer).ListMetricsScopesByMonitoredProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.metricsscope.v1.MetricsScopes/ListMetricsScopesByMonitoredProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsScopesServer).ListMetricsScopesByMonitoredProject(ctx, req.(*ListMetricsScopesByMonitoredProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsScopes_CreateMonitoredProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMonitoredProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsScopesServer).CreateMonitoredProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.metricsscope.v1.MetricsScopes/CreateMonitoredProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsScopesServer).CreateMonitoredProject(ctx, req.(*CreateMonitoredProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsScopes_DeleteMonitoredProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMonitoredProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsScopesServer).DeleteMonitoredProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.metricsscope.v1.MetricsScopes/DeleteMonitoredProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsScopesServer).DeleteMonitoredProject(ctx, req.(*DeleteMonitoredProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsScopes_ServiceDesc is the grpc.ServiceDesc for MetricsScopes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsScopes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.monitoring.metricsscope.v1.MetricsScopes",
	HandlerType: (*MetricsScopesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetricsScope",
			Handler:    _MetricsScopes_GetMetricsScope_Handler,
		},
		{
			MethodName: "ListMetricsScopesByMonitoredProject",
			Handler:    _MetricsScopes_ListMetricsScopesByMonitoredProject_Handler,
		},
		{
			MethodName: "CreateMonitoredProject",
			Handler:    _MetricsScopes_CreateMonitoredProject_Handler,
		},
		{
			MethodName: "DeleteMonitoredProject",
			Handler:    _MetricsScopes_DeleteMonitoredProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/monitoring/metricsscope/v1/metrics_scopes.proto",
}
