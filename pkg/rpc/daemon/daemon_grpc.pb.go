// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package daemon

import (
	context "context"
	iptables "github.com/datawire/telepresence2/pkg/rpc/iptables"
	version "github.com/datawire/telepresence2/pkg/rpc/version"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonClient interface {
	// Returns version information from the Daemon
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*version.VersionInfo, error)
	// Returns the current connectivity status
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DaemonStatus, error)
	// Turns network overrides off.
	Pause(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PauseInfo, error)
	// Turns network overrides back on (after using Pause)
	Resume(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResumeInfo, error)
	// Logger accepts a stream that can be used for posting messages to
	// the daemon logger.
	Logger(ctx context.Context, opts ...grpc.CallOption) (Daemon_LoggerClient, error)
	// Quits (terminates) the service.
	Quit(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// DeleteIPTable deletes the IP-table for the given name
	DeleteIPTable(ctx context.Context, in *TableName, opts ...grpc.CallOption) (*empty.Empty, error)
	// IPTable returns the IP-table for the given name
	IPTable(ctx context.Context, in *TableName, opts ...grpc.CallOption) (*iptables.Table, error)
	// IPTable returns all IP-tables
	AllIPTables(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Tables, error)
	// Update assigns adds or updates an IP-table.
	Update(ctx context.Context, in *iptables.Table, opts ...grpc.CallOption) (*empty.Empty, error)
	// DnsSearchPath returns the DNS search path
	DnsSearchPath(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Paths, error)
	// SetSearch sets a new search path.
	SetDnsSearchPath(ctx context.Context, in *Paths, opts ...grpc.CallOption) (*empty.Empty, error)
}

type daemonClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonClient(cc grpc.ClientConnInterface) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*version.VersionInfo, error) {
	out := new(version.VersionInfo)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DaemonStatus, error) {
	out := new(DaemonStatus)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Pause(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PauseInfo, error) {
	out := new(PauseInfo)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Resume(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResumeInfo, error) {
	out := new(ResumeInfo)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/Resume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Logger(ctx context.Context, opts ...grpc.CallOption) (Daemon_LoggerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Daemon_serviceDesc.Streams[0], "/telepresence.Daemon/Logger", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonLoggerClient{stream}
	return x, nil
}

type Daemon_LoggerClient interface {
	Send(*LogMessage) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type daemonLoggerClient struct {
	grpc.ClientStream
}

func (x *daemonLoggerClient) Send(m *LogMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *daemonLoggerClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) Quit(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/Quit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) DeleteIPTable(ctx context.Context, in *TableName, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/DeleteIPTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) IPTable(ctx context.Context, in *TableName, opts ...grpc.CallOption) (*iptables.Table, error) {
	out := new(iptables.Table)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/IPTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) AllIPTables(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Tables, error) {
	out := new(Tables)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/AllIPTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Update(ctx context.Context, in *iptables.Table, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) DnsSearchPath(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Paths, error) {
	out := new(Paths)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/DnsSearchPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) SetDnsSearchPath(ctx context.Context, in *Paths, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.Daemon/SetDnsSearchPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServer is the server API for Daemon service.
// All implementations must embed UnimplementedDaemonServer
// for forward compatibility
type DaemonServer interface {
	// Returns version information from the Daemon
	Version(context.Context, *empty.Empty) (*version.VersionInfo, error)
	// Returns the current connectivity status
	Status(context.Context, *empty.Empty) (*DaemonStatus, error)
	// Turns network overrides off.
	Pause(context.Context, *empty.Empty) (*PauseInfo, error)
	// Turns network overrides back on (after using Pause)
	Resume(context.Context, *empty.Empty) (*ResumeInfo, error)
	// Logger accepts a stream that can be used for posting messages to
	// the daemon logger.
	Logger(Daemon_LoggerServer) error
	// Quits (terminates) the service.
	Quit(context.Context, *empty.Empty) (*empty.Empty, error)
	// DeleteIPTable deletes the IP-table for the given name
	DeleteIPTable(context.Context, *TableName) (*empty.Empty, error)
	// IPTable returns the IP-table for the given name
	IPTable(context.Context, *TableName) (*iptables.Table, error)
	// IPTable returns all IP-tables
	AllIPTables(context.Context, *empty.Empty) (*Tables, error)
	// Update assigns adds or updates an IP-table.
	Update(context.Context, *iptables.Table) (*empty.Empty, error)
	// DnsSearchPath returns the DNS search path
	DnsSearchPath(context.Context, *empty.Empty) (*Paths, error)
	// SetSearch sets a new search path.
	SetDnsSearchPath(context.Context, *Paths) (*empty.Empty, error)
	mustEmbedUnimplementedDaemonServer()
}

// UnimplementedDaemonServer must be embedded to have forward compatible implementations.
type UnimplementedDaemonServer struct {
}

func (UnimplementedDaemonServer) Version(context.Context, *empty.Empty) (*version.VersionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedDaemonServer) Status(context.Context, *empty.Empty) (*DaemonStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedDaemonServer) Pause(context.Context, *empty.Empty) (*PauseInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedDaemonServer) Resume(context.Context, *empty.Empty) (*ResumeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (UnimplementedDaemonServer) Logger(Daemon_LoggerServer) error {
	return status.Errorf(codes.Unimplemented, "method Logger not implemented")
}
func (UnimplementedDaemonServer) Quit(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Quit not implemented")
}
func (UnimplementedDaemonServer) DeleteIPTable(context.Context, *TableName) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIPTable not implemented")
}
func (UnimplementedDaemonServer) IPTable(context.Context, *TableName) (*iptables.Table, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IPTable not implemented")
}
func (UnimplementedDaemonServer) AllIPTables(context.Context, *empty.Empty) (*Tables, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllIPTables not implemented")
}
func (UnimplementedDaemonServer) Update(context.Context, *iptables.Table) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDaemonServer) DnsSearchPath(context.Context, *empty.Empty) (*Paths, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DnsSearchPath not implemented")
}
func (UnimplementedDaemonServer) SetDnsSearchPath(context.Context, *Paths) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDnsSearchPath not implemented")
}
func (UnimplementedDaemonServer) mustEmbedUnimplementedDaemonServer() {}

// UnsafeDaemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServer will
// result in compilation errors.
type UnsafeDaemonServer interface {
	mustEmbedUnimplementedDaemonServer()
}

func RegisterDaemonServer(s grpc.ServiceRegistrar, srv DaemonServer) {
	s.RegisterService(&_Daemon_serviceDesc, srv)
}

func _Daemon_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Pause(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Resume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Resume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/Resume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Resume(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Logger_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DaemonServer).Logger(&daemonLoggerServer{stream})
}

type Daemon_LoggerServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*LogMessage, error)
	grpc.ServerStream
}

type daemonLoggerServer struct {
	grpc.ServerStream
}

func (x *daemonLoggerServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *daemonLoggerServer) Recv() (*LogMessage, error) {
	m := new(LogMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Daemon_Quit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Quit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/Quit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Quit(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_DeleteIPTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).DeleteIPTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/DeleteIPTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).DeleteIPTable(ctx, req.(*TableName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_IPTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).IPTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/IPTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).IPTable(ctx, req.(*TableName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_AllIPTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).AllIPTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/AllIPTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).AllIPTables(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iptables.Table)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Update(ctx, req.(*iptables.Table))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_DnsSearchPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).DnsSearchPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/DnsSearchPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).DnsSearchPath(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_SetDnsSearchPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paths)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).SetDnsSearchPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Daemon/SetDnsSearchPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).SetDnsSearchPath(ctx, req.(*Paths))
	}
	return interceptor(ctx, in, info, handler)
}

var _Daemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telepresence.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Daemon_Version_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Daemon_Status_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _Daemon_Pause_Handler,
		},
		{
			MethodName: "Resume",
			Handler:    _Daemon_Resume_Handler,
		},
		{
			MethodName: "Quit",
			Handler:    _Daemon_Quit_Handler,
		},
		{
			MethodName: "DeleteIPTable",
			Handler:    _Daemon_DeleteIPTable_Handler,
		},
		{
			MethodName: "IPTable",
			Handler:    _Daemon_IPTable_Handler,
		},
		{
			MethodName: "AllIPTables",
			Handler:    _Daemon_AllIPTables_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Daemon_Update_Handler,
		},
		{
			MethodName: "DnsSearchPath",
			Handler:    _Daemon_DnsSearchPath_Handler,
		},
		{
			MethodName: "SetDnsSearchPath",
			Handler:    _Daemon_SetDnsSearchPath_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logger",
			Handler:       _Daemon_Logger_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "rpc/daemon/daemon.proto",
}
