// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/web.proto

package web

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WebClient is the client API for Web service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebClient interface {
	DeleteAllServers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	CreatePlugin(ctx context.Context, in *CreatePluginResponse, opts ...grpc.CallOption) (*Response, error)
	DeletePlugin(ctx context.Context, in *DeletePluginResponse, opts ...grpc.CallOption) (*Response, error)
	GetPlugin(ctx context.Context, in *ResponseById, opts ...grpc.CallOption) (*Plugin, error)
	GetPlugins(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Plugins, error)
	UpdatePlugin(ctx context.Context, in *Plugin, opts ...grpc.CallOption) (*Response, error)
	CreateVersion(ctx context.Context, in *CreateVersionResponse, opts ...grpc.CallOption) (*Response, error)
	DeleteVersion(ctx context.Context, in *DeleteVersionResponse, opts ...grpc.CallOption) (*Response, error)
	GetVersion(ctx context.Context, in *ResponseById, opts ...grpc.CallOption) (*Version, error)
	GetVersions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Versions, error)
	UpdateVersion(ctx context.Context, in *Version, opts ...grpc.CallOption) (*Response, error)
}

type webClient struct {
	cc grpc.ClientConnInterface
}

func NewWebClient(cc grpc.ClientConnInterface) WebClient {
	return &webClient{cc}
}

func (c *webClient) DeleteAllServers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/web.Web/DeleteAllServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) CreatePlugin(ctx context.Context, in *CreatePluginResponse, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/web.Web/CreatePlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) DeletePlugin(ctx context.Context, in *DeletePluginResponse, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/web.Web/DeletePlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) GetPlugin(ctx context.Context, in *ResponseById, opts ...grpc.CallOption) (*Plugin, error) {
	out := new(Plugin)
	err := c.cc.Invoke(ctx, "/web.Web/GetPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) GetPlugins(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Plugins, error) {
	out := new(Plugins)
	err := c.cc.Invoke(ctx, "/web.Web/GetPlugins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) UpdatePlugin(ctx context.Context, in *Plugin, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/web.Web/UpdatePlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) CreateVersion(ctx context.Context, in *CreateVersionResponse, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/web.Web/CreateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) DeleteVersion(ctx context.Context, in *DeleteVersionResponse, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/web.Web/DeleteVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) GetVersion(ctx context.Context, in *ResponseById, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/web.Web/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) GetVersions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Versions, error) {
	out := new(Versions)
	err := c.cc.Invoke(ctx, "/web.Web/GetVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) UpdateVersion(ctx context.Context, in *Version, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/web.Web/UpdateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebServer is the server API for Web service.
// All implementations must embed UnimplementedWebServer
// for forward compatibility
type WebServer interface {
	DeleteAllServers(context.Context, *Empty) (*Response, error)
	CreatePlugin(context.Context, *CreatePluginResponse) (*Response, error)
	DeletePlugin(context.Context, *DeletePluginResponse) (*Response, error)
	GetPlugin(context.Context, *ResponseById) (*Plugin, error)
	GetPlugins(context.Context, *Empty) (*Plugins, error)
	UpdatePlugin(context.Context, *Plugin) (*Response, error)
	CreateVersion(context.Context, *CreateVersionResponse) (*Response, error)
	DeleteVersion(context.Context, *DeleteVersionResponse) (*Response, error)
	GetVersion(context.Context, *ResponseById) (*Version, error)
	GetVersions(context.Context, *Empty) (*Versions, error)
	UpdateVersion(context.Context, *Version) (*Response, error)
	mustEmbedUnimplementedWebServer()
}

// UnimplementedWebServer must be embedded to have forward compatible implementations.
type UnimplementedWebServer struct {
}

func (UnimplementedWebServer) DeleteAllServers(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllServers not implemented")
}
func (UnimplementedWebServer) CreatePlugin(context.Context, *CreatePluginResponse) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlugin not implemented")
}
func (UnimplementedWebServer) DeletePlugin(context.Context, *DeletePluginResponse) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlugin not implemented")
}
func (UnimplementedWebServer) GetPlugin(context.Context, *ResponseById) (*Plugin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlugin not implemented")
}
func (UnimplementedWebServer) GetPlugins(context.Context, *Empty) (*Plugins, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlugins not implemented")
}
func (UnimplementedWebServer) UpdatePlugin(context.Context, *Plugin) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlugin not implemented")
}
func (UnimplementedWebServer) CreateVersion(context.Context, *CreateVersionResponse) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVersion not implemented")
}
func (UnimplementedWebServer) DeleteVersion(context.Context, *DeleteVersionResponse) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVersion not implemented")
}
func (UnimplementedWebServer) GetVersion(context.Context, *ResponseById) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedWebServer) GetVersions(context.Context, *Empty) (*Versions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersions not implemented")
}
func (UnimplementedWebServer) UpdateVersion(context.Context, *Version) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVersion not implemented")
}
func (UnimplementedWebServer) mustEmbedUnimplementedWebServer() {}

// UnsafeWebServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebServer will
// result in compilation errors.
type UnsafeWebServer interface {
	mustEmbedUnimplementedWebServer()
}

func RegisterWebServer(s grpc.ServiceRegistrar, srv WebServer) {
	s.RegisterService(&Web_ServiceDesc, srv)
}

func _Web_DeleteAllServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).DeleteAllServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/DeleteAllServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).DeleteAllServers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_CreatePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePluginResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).CreatePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/CreatePlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).CreatePlugin(ctx, req.(*CreatePluginResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_DeletePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePluginResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).DeletePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/DeletePlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).DeletePlugin(ctx, req.(*DeletePluginResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_GetPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResponseById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).GetPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/GetPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).GetPlugin(ctx, req.(*ResponseById))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_GetPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).GetPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/GetPlugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).GetPlugins(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_UpdatePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Plugin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).UpdatePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/UpdatePlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).UpdatePlugin(ctx, req.(*Plugin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_CreateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVersionResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).CreateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/CreateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).CreateVersion(ctx, req.(*CreateVersionResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_DeleteVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVersionResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).DeleteVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/DeleteVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).DeleteVersion(ctx, req.(*DeleteVersionResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResponseById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).GetVersion(ctx, req.(*ResponseById))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_GetVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).GetVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/GetVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).GetVersions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_UpdateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Version)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).UpdateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/UpdateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).UpdateVersion(ctx, req.(*Version))
	}
	return interceptor(ctx, in, info, handler)
}

// Web_ServiceDesc is the grpc.ServiceDesc for Web service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Web_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "web.Web",
	HandlerType: (*WebServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteAllServers",
			Handler:    _Web_DeleteAllServers_Handler,
		},
		{
			MethodName: "CreatePlugin",
			Handler:    _Web_CreatePlugin_Handler,
		},
		{
			MethodName: "DeletePlugin",
			Handler:    _Web_DeletePlugin_Handler,
		},
		{
			MethodName: "GetPlugin",
			Handler:    _Web_GetPlugin_Handler,
		},
		{
			MethodName: "GetPlugins",
			Handler:    _Web_GetPlugins_Handler,
		},
		{
			MethodName: "UpdatePlugin",
			Handler:    _Web_UpdatePlugin_Handler,
		},
		{
			MethodName: "CreateVersion",
			Handler:    _Web_CreateVersion_Handler,
		},
		{
			MethodName: "DeleteVersion",
			Handler:    _Web_DeleteVersion_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Web_GetVersion_Handler,
		},
		{
			MethodName: "GetVersions",
			Handler:    _Web_GetVersions_Handler,
		},
		{
			MethodName: "UpdateVersion",
			Handler:    _Web_UpdateVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/web.proto",
}
