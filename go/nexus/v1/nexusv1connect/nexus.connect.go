// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: nexus/v1/nexus.proto

package nexusv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/liquidspill/proto/go/nexus/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// NexusServiceName is the fully-qualified name of the NexusService service.
	NexusServiceName = "nexus.v1.NexusService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// NexusServiceCreateTenantProcedure is the fully-qualified name of the NexusService's CreateTenant
	// RPC.
	NexusServiceCreateTenantProcedure = "/nexus.v1.NexusService/CreateTenant"
	// NexusServiceDeleteTenantProcedure is the fully-qualified name of the NexusService's DeleteTenant
	// RPC.
	NexusServiceDeleteTenantProcedure = "/nexus.v1.NexusService/DeleteTenant"
	// NexusServiceCreatePartitionProcedure is the fully-qualified name of the NexusService's
	// CreatePartition RPC.
	NexusServiceCreatePartitionProcedure = "/nexus.v1.NexusService/CreatePartition"
	// NexusServiceListPartitionsProcedure is the fully-qualified name of the NexusService's
	// ListPartitions RPC.
	NexusServiceListPartitionsProcedure = "/nexus.v1.NexusService/ListPartitions"
	// NexusServiceValidateMetadataVersionProcedure is the fully-qualified name of the NexusService's
	// ValidateMetadataVersion RPC.
	NexusServiceValidateMetadataVersionProcedure = "/nexus.v1.NexusService/ValidateMetadataVersion"
	// NexusServiceGetMetadataVersionProcedure is the fully-qualified name of the NexusService's
	// GetMetadataVersion RPC.
	NexusServiceGetMetadataVersionProcedure = "/nexus.v1.NexusService/GetMetadataVersion"
	// NexusServiceListMetadataVersionsProcedure is the fully-qualified name of the NexusService's
	// ListMetadataVersions RPC.
	NexusServiceListMetadataVersionsProcedure = "/nexus.v1.NexusService/ListMetadataVersions"
)

// NexusServiceClient is a client for the nexus.v1.NexusService service.
type NexusServiceClient interface {
	// ===== Management operations =====
	CreateTenant(context.Context, *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error)
	DeleteTenant(context.Context, *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error)
	// ===== Data operations =====
	CreatePartition(context.Context, *connect.Request[v1.CreatePartitionRequest]) (*connect.Response[v1.CreatePartitionResponse], error)
	ListPartitions(context.Context, *connect.Request[v1.ListPartitionsRequest]) (*connect.Response[v1.ListPartitionsResponse], error)
	ValidateMetadataVersion(context.Context, *connect.Request[v1.ValidateMetadataVersionRequest]) (*connect.Response[v1.ValidateMetadataVersionResponse], error)
	GetMetadataVersion(context.Context, *connect.Request[v1.GetMetadataVersionRequest]) (*connect.Response[v1.GetMetadataVersionResponse], error)
	ListMetadataVersions(context.Context, *connect.Request[v1.ListMetadataVersionsRequest]) (*connect.Response[v1.ListMetadataVersionsResponse], error)
}

// NewNexusServiceClient constructs a client for the nexus.v1.NexusService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNexusServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) NexusServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	nexusServiceMethods := v1.File_nexus_v1_nexus_proto.Services().ByName("NexusService").Methods()
	return &nexusServiceClient{
		createTenant: connect.NewClient[v1.CreateTenantRequest, v1.CreateTenantResponse](
			httpClient,
			baseURL+NexusServiceCreateTenantProcedure,
			connect.WithSchema(nexusServiceMethods.ByName("CreateTenant")),
			connect.WithClientOptions(opts...),
		),
		deleteTenant: connect.NewClient[v1.DeleteTenantRequest, v1.DeleteTenantResponse](
			httpClient,
			baseURL+NexusServiceDeleteTenantProcedure,
			connect.WithSchema(nexusServiceMethods.ByName("DeleteTenant")),
			connect.WithClientOptions(opts...),
		),
		createPartition: connect.NewClient[v1.CreatePartitionRequest, v1.CreatePartitionResponse](
			httpClient,
			baseURL+NexusServiceCreatePartitionProcedure,
			connect.WithSchema(nexusServiceMethods.ByName("CreatePartition")),
			connect.WithClientOptions(opts...),
		),
		listPartitions: connect.NewClient[v1.ListPartitionsRequest, v1.ListPartitionsResponse](
			httpClient,
			baseURL+NexusServiceListPartitionsProcedure,
			connect.WithSchema(nexusServiceMethods.ByName("ListPartitions")),
			connect.WithClientOptions(opts...),
		),
		validateMetadataVersion: connect.NewClient[v1.ValidateMetadataVersionRequest, v1.ValidateMetadataVersionResponse](
			httpClient,
			baseURL+NexusServiceValidateMetadataVersionProcedure,
			connect.WithSchema(nexusServiceMethods.ByName("ValidateMetadataVersion")),
			connect.WithClientOptions(opts...),
		),
		getMetadataVersion: connect.NewClient[v1.GetMetadataVersionRequest, v1.GetMetadataVersionResponse](
			httpClient,
			baseURL+NexusServiceGetMetadataVersionProcedure,
			connect.WithSchema(nexusServiceMethods.ByName("GetMetadataVersion")),
			connect.WithClientOptions(opts...),
		),
		listMetadataVersions: connect.NewClient[v1.ListMetadataVersionsRequest, v1.ListMetadataVersionsResponse](
			httpClient,
			baseURL+NexusServiceListMetadataVersionsProcedure,
			connect.WithSchema(nexusServiceMethods.ByName("ListMetadataVersions")),
			connect.WithClientOptions(opts...),
		),
	}
}

// nexusServiceClient implements NexusServiceClient.
type nexusServiceClient struct {
	createTenant            *connect.Client[v1.CreateTenantRequest, v1.CreateTenantResponse]
	deleteTenant            *connect.Client[v1.DeleteTenantRequest, v1.DeleteTenantResponse]
	createPartition         *connect.Client[v1.CreatePartitionRequest, v1.CreatePartitionResponse]
	listPartitions          *connect.Client[v1.ListPartitionsRequest, v1.ListPartitionsResponse]
	validateMetadataVersion *connect.Client[v1.ValidateMetadataVersionRequest, v1.ValidateMetadataVersionResponse]
	getMetadataVersion      *connect.Client[v1.GetMetadataVersionRequest, v1.GetMetadataVersionResponse]
	listMetadataVersions    *connect.Client[v1.ListMetadataVersionsRequest, v1.ListMetadataVersionsResponse]
}

// CreateTenant calls nexus.v1.NexusService.CreateTenant.
func (c *nexusServiceClient) CreateTenant(ctx context.Context, req *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error) {
	return c.createTenant.CallUnary(ctx, req)
}

// DeleteTenant calls nexus.v1.NexusService.DeleteTenant.
func (c *nexusServiceClient) DeleteTenant(ctx context.Context, req *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error) {
	return c.deleteTenant.CallUnary(ctx, req)
}

// CreatePartition calls nexus.v1.NexusService.CreatePartition.
func (c *nexusServiceClient) CreatePartition(ctx context.Context, req *connect.Request[v1.CreatePartitionRequest]) (*connect.Response[v1.CreatePartitionResponse], error) {
	return c.createPartition.CallUnary(ctx, req)
}

// ListPartitions calls nexus.v1.NexusService.ListPartitions.
func (c *nexusServiceClient) ListPartitions(ctx context.Context, req *connect.Request[v1.ListPartitionsRequest]) (*connect.Response[v1.ListPartitionsResponse], error) {
	return c.listPartitions.CallUnary(ctx, req)
}

// ValidateMetadataVersion calls nexus.v1.NexusService.ValidateMetadataVersion.
func (c *nexusServiceClient) ValidateMetadataVersion(ctx context.Context, req *connect.Request[v1.ValidateMetadataVersionRequest]) (*connect.Response[v1.ValidateMetadataVersionResponse], error) {
	return c.validateMetadataVersion.CallUnary(ctx, req)
}

// GetMetadataVersion calls nexus.v1.NexusService.GetMetadataVersion.
func (c *nexusServiceClient) GetMetadataVersion(ctx context.Context, req *connect.Request[v1.GetMetadataVersionRequest]) (*connect.Response[v1.GetMetadataVersionResponse], error) {
	return c.getMetadataVersion.CallUnary(ctx, req)
}

// ListMetadataVersions calls nexus.v1.NexusService.ListMetadataVersions.
func (c *nexusServiceClient) ListMetadataVersions(ctx context.Context, req *connect.Request[v1.ListMetadataVersionsRequest]) (*connect.Response[v1.ListMetadataVersionsResponse], error) {
	return c.listMetadataVersions.CallUnary(ctx, req)
}

// NexusServiceHandler is an implementation of the nexus.v1.NexusService service.
type NexusServiceHandler interface {
	// ===== Management operations =====
	CreateTenant(context.Context, *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error)
	DeleteTenant(context.Context, *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error)
	// ===== Data operations =====
	CreatePartition(context.Context, *connect.Request[v1.CreatePartitionRequest]) (*connect.Response[v1.CreatePartitionResponse], error)
	ListPartitions(context.Context, *connect.Request[v1.ListPartitionsRequest]) (*connect.Response[v1.ListPartitionsResponse], error)
	ValidateMetadataVersion(context.Context, *connect.Request[v1.ValidateMetadataVersionRequest]) (*connect.Response[v1.ValidateMetadataVersionResponse], error)
	GetMetadataVersion(context.Context, *connect.Request[v1.GetMetadataVersionRequest]) (*connect.Response[v1.GetMetadataVersionResponse], error)
	ListMetadataVersions(context.Context, *connect.Request[v1.ListMetadataVersionsRequest]) (*connect.Response[v1.ListMetadataVersionsResponse], error)
}

// NewNexusServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNexusServiceHandler(svc NexusServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	nexusServiceMethods := v1.File_nexus_v1_nexus_proto.Services().ByName("NexusService").Methods()
	nexusServiceCreateTenantHandler := connect.NewUnaryHandler(
		NexusServiceCreateTenantProcedure,
		svc.CreateTenant,
		connect.WithSchema(nexusServiceMethods.ByName("CreateTenant")),
		connect.WithHandlerOptions(opts...),
	)
	nexusServiceDeleteTenantHandler := connect.NewUnaryHandler(
		NexusServiceDeleteTenantProcedure,
		svc.DeleteTenant,
		connect.WithSchema(nexusServiceMethods.ByName("DeleteTenant")),
		connect.WithHandlerOptions(opts...),
	)
	nexusServiceCreatePartitionHandler := connect.NewUnaryHandler(
		NexusServiceCreatePartitionProcedure,
		svc.CreatePartition,
		connect.WithSchema(nexusServiceMethods.ByName("CreatePartition")),
		connect.WithHandlerOptions(opts...),
	)
	nexusServiceListPartitionsHandler := connect.NewUnaryHandler(
		NexusServiceListPartitionsProcedure,
		svc.ListPartitions,
		connect.WithSchema(nexusServiceMethods.ByName("ListPartitions")),
		connect.WithHandlerOptions(opts...),
	)
	nexusServiceValidateMetadataVersionHandler := connect.NewUnaryHandler(
		NexusServiceValidateMetadataVersionProcedure,
		svc.ValidateMetadataVersion,
		connect.WithSchema(nexusServiceMethods.ByName("ValidateMetadataVersion")),
		connect.WithHandlerOptions(opts...),
	)
	nexusServiceGetMetadataVersionHandler := connect.NewUnaryHandler(
		NexusServiceGetMetadataVersionProcedure,
		svc.GetMetadataVersion,
		connect.WithSchema(nexusServiceMethods.ByName("GetMetadataVersion")),
		connect.WithHandlerOptions(opts...),
	)
	nexusServiceListMetadataVersionsHandler := connect.NewUnaryHandler(
		NexusServiceListMetadataVersionsProcedure,
		svc.ListMetadataVersions,
		connect.WithSchema(nexusServiceMethods.ByName("ListMetadataVersions")),
		connect.WithHandlerOptions(opts...),
	)
	return "/nexus.v1.NexusService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case NexusServiceCreateTenantProcedure:
			nexusServiceCreateTenantHandler.ServeHTTP(w, r)
		case NexusServiceDeleteTenantProcedure:
			nexusServiceDeleteTenantHandler.ServeHTTP(w, r)
		case NexusServiceCreatePartitionProcedure:
			nexusServiceCreatePartitionHandler.ServeHTTP(w, r)
		case NexusServiceListPartitionsProcedure:
			nexusServiceListPartitionsHandler.ServeHTTP(w, r)
		case NexusServiceValidateMetadataVersionProcedure:
			nexusServiceValidateMetadataVersionHandler.ServeHTTP(w, r)
		case NexusServiceGetMetadataVersionProcedure:
			nexusServiceGetMetadataVersionHandler.ServeHTTP(w, r)
		case NexusServiceListMetadataVersionsProcedure:
			nexusServiceListMetadataVersionsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedNexusServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedNexusServiceHandler struct{}

func (UnimplementedNexusServiceHandler) CreateTenant(context.Context, *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nexus.v1.NexusService.CreateTenant is not implemented"))
}

func (UnimplementedNexusServiceHandler) DeleteTenant(context.Context, *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nexus.v1.NexusService.DeleteTenant is not implemented"))
}

func (UnimplementedNexusServiceHandler) CreatePartition(context.Context, *connect.Request[v1.CreatePartitionRequest]) (*connect.Response[v1.CreatePartitionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nexus.v1.NexusService.CreatePartition is not implemented"))
}

func (UnimplementedNexusServiceHandler) ListPartitions(context.Context, *connect.Request[v1.ListPartitionsRequest]) (*connect.Response[v1.ListPartitionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nexus.v1.NexusService.ListPartitions is not implemented"))
}

func (UnimplementedNexusServiceHandler) ValidateMetadataVersion(context.Context, *connect.Request[v1.ValidateMetadataVersionRequest]) (*connect.Response[v1.ValidateMetadataVersionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nexus.v1.NexusService.ValidateMetadataVersion is not implemented"))
}

func (UnimplementedNexusServiceHandler) GetMetadataVersion(context.Context, *connect.Request[v1.GetMetadataVersionRequest]) (*connect.Response[v1.GetMetadataVersionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nexus.v1.NexusService.GetMetadataVersion is not implemented"))
}

func (UnimplementedNexusServiceHandler) ListMetadataVersions(context.Context, *connect.Request[v1.ListMetadataVersionsRequest]) (*connect.Response[v1.ListMetadataVersionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("nexus.v1.NexusService.ListMetadataVersions is not implemented"))
}
