// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: cloudevent.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CloudEventService_Publish_FullMethodName   = "/io.cloudevents.v1.CloudEventService/Publish"
	CloudEventService_Subscribe_FullMethodName = "/io.cloudevents.v1.CloudEventService/Subscribe"
)

// CloudEventServiceClient is the client API for CloudEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudEventServiceClient interface {
	Publish(ctx context.Context, in *CloudEvent, opts ...grpc.CallOption) (*empty.Empty, error)
	Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (CloudEventService_SubscribeClient, error)
}

type cloudEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudEventServiceClient(cc grpc.ClientConnInterface) CloudEventServiceClient {
	return &cloudEventServiceClient{cc}
}

func (c *cloudEventServiceClient) Publish(ctx context.Context, in *CloudEvent, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, CloudEventService_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudEventServiceClient) Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (CloudEventService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &CloudEventService_ServiceDesc.Streams[0], CloudEventService_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &cloudEventServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CloudEventService_SubscribeClient interface {
	Recv() (*CloudEvent, error)
	grpc.ClientStream
}

type cloudEventServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *cloudEventServiceSubscribeClient) Recv() (*CloudEvent, error) {
	m := new(CloudEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CloudEventServiceServer is the server API for CloudEventService service.
// All implementations must embed UnimplementedCloudEventServiceServer
// for forward compatibility
type CloudEventServiceServer interface {
	Publish(context.Context, *CloudEvent) (*empty.Empty, error)
	Subscribe(*Subscription, CloudEventService_SubscribeServer) error
	mustEmbedUnimplementedCloudEventServiceServer()
}

// UnimplementedCloudEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCloudEventServiceServer struct {
}

func (UnimplementedCloudEventServiceServer) Publish(context.Context, *CloudEvent) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedCloudEventServiceServer) Subscribe(*Subscription, CloudEventService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedCloudEventServiceServer) mustEmbedUnimplementedCloudEventServiceServer() {}

// UnsafeCloudEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudEventServiceServer will
// result in compilation errors.
type UnsafeCloudEventServiceServer interface {
	mustEmbedUnimplementedCloudEventServiceServer()
}

func RegisterCloudEventServiceServer(s grpc.ServiceRegistrar, srv CloudEventServiceServer) {
	s.RegisterService(&CloudEventService_ServiceDesc, srv)
}

func _CloudEventService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudEventServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudEventService_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudEventServiceServer).Publish(ctx, req.(*CloudEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudEventService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Subscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CloudEventServiceServer).Subscribe(m, &cloudEventServiceSubscribeServer{stream})
}

type CloudEventService_SubscribeServer interface {
	Send(*CloudEvent) error
	grpc.ServerStream
}

type cloudEventServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *cloudEventServiceSubscribeServer) Send(m *CloudEvent) error {
	return x.ServerStream.SendMsg(m)
}

// CloudEventService_ServiceDesc is the grpc.ServiceDesc for CloudEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.cloudevents.v1.CloudEventService",
	HandlerType: (*CloudEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _CloudEventService_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _CloudEventService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloudevent.proto",
}
