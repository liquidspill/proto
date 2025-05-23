// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: nexus/metadata/v1/metadata.proto

package metadatav1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/liquidspill/proto/go/nexus/metadata/v1"
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
	// MetadataServiceName is the fully-qualified name of the MetadataService service.
	MetadataServiceName = "metadata.v1.MetadataService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MetadataServiceCreateTenantProcedure is the fully-qualified name of the MetadataService's
	// CreateTenant RPC.
	MetadataServiceCreateTenantProcedure = "/metadata.v1.MetadataService/CreateTenant"
	// MetadataServiceDeleteTenantProcedure is the fully-qualified name of the MetadataService's
	// DeleteTenant RPC.
	MetadataServiceDeleteTenantProcedure = "/metadata.v1.MetadataService/DeleteTenant"
	// MetadataServiceCreatePartitionProcedure is the fully-qualified name of the MetadataService's
	// CreatePartition RPC.
	MetadataServiceCreatePartitionProcedure = "/metadata.v1.MetadataService/CreatePartition"
	// MetadataServiceListPartitionsProcedure is the fully-qualified name of the MetadataService's
	// ListPartitions RPC.
	MetadataServiceListPartitionsProcedure = "/metadata.v1.MetadataService/ListPartitions"
	// MetadataServiceValidateMetadataVersionProcedure is the fully-qualified name of the
	// MetadataService's ValidateMetadataVersion RPC.
	MetadataServiceValidateMetadataVersionProcedure = "/metadata.v1.MetadataService/ValidateMetadataVersion"
	// MetadataServiceGetMetadataVersionProcedure is the fully-qualified name of the MetadataService's
	// GetMetadataVersion RPC.
	MetadataServiceGetMetadataVersionProcedure = "/metadata.v1.MetadataService/GetMetadataVersion"
	// MetadataServiceListMetadataVersionsProcedure is the fully-qualified name of the MetadataService's
	// ListMetadataVersions RPC.
	MetadataServiceListMetadataVersionsProcedure = "/metadata.v1.MetadataService/ListMetadataVersions"
)

// MetadataServiceClient is a client for the metadata.v1.MetadataService service.
type MetadataServiceClient interface {
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

// NewMetadataServiceClient constructs a client for the metadata.v1.MetadataService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMetadataServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MetadataServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	metadataServiceMethods := v1.File_nexus_metadata_v1_metadata_proto.Services().ByName("MetadataService").Methods()
	return &metadataServiceClient{
		createTenant: connect.NewClient[v1.CreateTenantRequest, v1.CreateTenantResponse](
			httpClient,
			baseURL+MetadataServiceCreateTenantProcedure,
			connect.WithSchema(metadataServiceMethods.ByName("CreateTenant")),
			connect.WithClientOptions(opts...),
		),
		deleteTenant: connect.NewClient[v1.DeleteTenantRequest, v1.DeleteTenantResponse](
			httpClient,
			baseURL+MetadataServiceDeleteTenantProcedure,
			connect.WithSchema(metadataServiceMethods.ByName("DeleteTenant")),
			connect.WithClientOptions(opts...),
		),
		createPartition: connect.NewClient[v1.CreatePartitionRequest, v1.CreatePartitionResponse](
			httpClient,
			baseURL+MetadataServiceCreatePartitionProcedure,
			connect.WithSchema(metadataServiceMethods.ByName("CreatePartition")),
			connect.WithClientOptions(opts...),
		),
		listPartitions: connect.NewClient[v1.ListPartitionsRequest, v1.ListPartitionsResponse](
			httpClient,
			baseURL+MetadataServiceListPartitionsProcedure,
			connect.WithSchema(metadataServiceMethods.ByName("ListPartitions")),
			connect.WithClientOptions(opts...),
		),
		validateMetadataVersion: connect.NewClient[v1.ValidateMetadataVersionRequest, v1.ValidateMetadataVersionResponse](
			httpClient,
			baseURL+MetadataServiceValidateMetadataVersionProcedure,
			connect.WithSchema(metadataServiceMethods.ByName("ValidateMetadataVersion")),
			connect.WithClientOptions(opts...),
		),
		getMetadataVersion: connect.NewClient[v1.GetMetadataVersionRequest, v1.GetMetadataVersionResponse](
			httpClient,
			baseURL+MetadataServiceGetMetadataVersionProcedure,
			connect.WithSchema(metadataServiceMethods.ByName("GetMetadataVersion")),
			connect.WithClientOptions(opts...),
		),
		listMetadataVersions: connect.NewClient[v1.ListMetadataVersionsRequest, v1.ListMetadataVersionsResponse](
			httpClient,
			baseURL+MetadataServiceListMetadataVersionsProcedure,
			connect.WithSchema(metadataServiceMethods.ByName("ListMetadataVersions")),
			connect.WithClientOptions(opts...),
		),
	}
}

// metadataServiceClient implements MetadataServiceClient.
type metadataServiceClient struct {
	createTenant            *connect.Client[v1.CreateTenantRequest, v1.CreateTenantResponse]
	deleteTenant            *connect.Client[v1.DeleteTenantRequest, v1.DeleteTenantResponse]
	createPartition         *connect.Client[v1.CreatePartitionRequest, v1.CreatePartitionResponse]
	listPartitions          *connect.Client[v1.ListPartitionsRequest, v1.ListPartitionsResponse]
	validateMetadataVersion *connect.Client[v1.ValidateMetadataVersionRequest, v1.ValidateMetadataVersionResponse]
	getMetadataVersion      *connect.Client[v1.GetMetadataVersionRequest, v1.GetMetadataVersionResponse]
	listMetadataVersions    *connect.Client[v1.ListMetadataVersionsRequest, v1.ListMetadataVersionsResponse]
}

// CreateTenant calls metadata.v1.MetadataService.CreateTenant.
func (c *metadataServiceClient) CreateTenant(ctx context.Context, req *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error) {
	return c.createTenant.CallUnary(ctx, req)
}

// DeleteTenant calls metadata.v1.MetadataService.DeleteTenant.
func (c *metadataServiceClient) DeleteTenant(ctx context.Context, req *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error) {
	return c.deleteTenant.CallUnary(ctx, req)
}

// CreatePartition calls metadata.v1.MetadataService.CreatePartition.
func (c *metadataServiceClient) CreatePartition(ctx context.Context, req *connect.Request[v1.CreatePartitionRequest]) (*connect.Response[v1.CreatePartitionResponse], error) {
	return c.createPartition.CallUnary(ctx, req)
}

// ListPartitions calls metadata.v1.MetadataService.ListPartitions.
func (c *metadataServiceClient) ListPartitions(ctx context.Context, req *connect.Request[v1.ListPartitionsRequest]) (*connect.Response[v1.ListPartitionsResponse], error) {
	return c.listPartitions.CallUnary(ctx, req)
}

// ValidateMetadataVersion calls metadata.v1.MetadataService.ValidateMetadataVersion.
func (c *metadataServiceClient) ValidateMetadataVersion(ctx context.Context, req *connect.Request[v1.ValidateMetadataVersionRequest]) (*connect.Response[v1.ValidateMetadataVersionResponse], error) {
	return c.validateMetadataVersion.CallUnary(ctx, req)
}

// GetMetadataVersion calls metadata.v1.MetadataService.GetMetadataVersion.
func (c *metadataServiceClient) GetMetadataVersion(ctx context.Context, req *connect.Request[v1.GetMetadataVersionRequest]) (*connect.Response[v1.GetMetadataVersionResponse], error) {
	return c.getMetadataVersion.CallUnary(ctx, req)
}

// ListMetadataVersions calls metadata.v1.MetadataService.ListMetadataVersions.
func (c *metadataServiceClient) ListMetadataVersions(ctx context.Context, req *connect.Request[v1.ListMetadataVersionsRequest]) (*connect.Response[v1.ListMetadataVersionsResponse], error) {
	return c.listMetadataVersions.CallUnary(ctx, req)
}

// MetadataServiceHandler is an implementation of the metadata.v1.MetadataService service.
type MetadataServiceHandler interface {
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

// NewMetadataServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMetadataServiceHandler(svc MetadataServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	metadataServiceMethods := v1.File_nexus_metadata_v1_metadata_proto.Services().ByName("MetadataService").Methods()
	metadataServiceCreateTenantHandler := connect.NewUnaryHandler(
		MetadataServiceCreateTenantProcedure,
		svc.CreateTenant,
		connect.WithSchema(metadataServiceMethods.ByName("CreateTenant")),
		connect.WithHandlerOptions(opts...),
	)
	metadataServiceDeleteTenantHandler := connect.NewUnaryHandler(
		MetadataServiceDeleteTenantProcedure,
		svc.DeleteTenant,
		connect.WithSchema(metadataServiceMethods.ByName("DeleteTenant")),
		connect.WithHandlerOptions(opts...),
	)
	metadataServiceCreatePartitionHandler := connect.NewUnaryHandler(
		MetadataServiceCreatePartitionProcedure,
		svc.CreatePartition,
		connect.WithSchema(metadataServiceMethods.ByName("CreatePartition")),
		connect.WithHandlerOptions(opts...),
	)
	metadataServiceListPartitionsHandler := connect.NewUnaryHandler(
		MetadataServiceListPartitionsProcedure,
		svc.ListPartitions,
		connect.WithSchema(metadataServiceMethods.ByName("ListPartitions")),
		connect.WithHandlerOptions(opts...),
	)
	metadataServiceValidateMetadataVersionHandler := connect.NewUnaryHandler(
		MetadataServiceValidateMetadataVersionProcedure,
		svc.ValidateMetadataVersion,
		connect.WithSchema(metadataServiceMethods.ByName("ValidateMetadataVersion")),
		connect.WithHandlerOptions(opts...),
	)
	metadataServiceGetMetadataVersionHandler := connect.NewUnaryHandler(
		MetadataServiceGetMetadataVersionProcedure,
		svc.GetMetadataVersion,
		connect.WithSchema(metadataServiceMethods.ByName("GetMetadataVersion")),
		connect.WithHandlerOptions(opts...),
	)
	metadataServiceListMetadataVersionsHandler := connect.NewUnaryHandler(
		MetadataServiceListMetadataVersionsProcedure,
		svc.ListMetadataVersions,
		connect.WithSchema(metadataServiceMethods.ByName("ListMetadataVersions")),
		connect.WithHandlerOptions(opts...),
	)
	return "/metadata.v1.MetadataService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MetadataServiceCreateTenantProcedure:
			metadataServiceCreateTenantHandler.ServeHTTP(w, r)
		case MetadataServiceDeleteTenantProcedure:
			metadataServiceDeleteTenantHandler.ServeHTTP(w, r)
		case MetadataServiceCreatePartitionProcedure:
			metadataServiceCreatePartitionHandler.ServeHTTP(w, r)
		case MetadataServiceListPartitionsProcedure:
			metadataServiceListPartitionsHandler.ServeHTTP(w, r)
		case MetadataServiceValidateMetadataVersionProcedure:
			metadataServiceValidateMetadataVersionHandler.ServeHTTP(w, r)
		case MetadataServiceGetMetadataVersionProcedure:
			metadataServiceGetMetadataVersionHandler.ServeHTTP(w, r)
		case MetadataServiceListMetadataVersionsProcedure:
			metadataServiceListMetadataVersionsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMetadataServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMetadataServiceHandler struct{}

func (UnimplementedMetadataServiceHandler) CreateTenant(context.Context, *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metadata.v1.MetadataService.CreateTenant is not implemented"))
}

func (UnimplementedMetadataServiceHandler) DeleteTenant(context.Context, *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metadata.v1.MetadataService.DeleteTenant is not implemented"))
}

func (UnimplementedMetadataServiceHandler) CreatePartition(context.Context, *connect.Request[v1.CreatePartitionRequest]) (*connect.Response[v1.CreatePartitionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metadata.v1.MetadataService.CreatePartition is not implemented"))
}

func (UnimplementedMetadataServiceHandler) ListPartitions(context.Context, *connect.Request[v1.ListPartitionsRequest]) (*connect.Response[v1.ListPartitionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metadata.v1.MetadataService.ListPartitions is not implemented"))
}

func (UnimplementedMetadataServiceHandler) ValidateMetadataVersion(context.Context, *connect.Request[v1.ValidateMetadataVersionRequest]) (*connect.Response[v1.ValidateMetadataVersionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metadata.v1.MetadataService.ValidateMetadataVersion is not implemented"))
}

func (UnimplementedMetadataServiceHandler) GetMetadataVersion(context.Context, *connect.Request[v1.GetMetadataVersionRequest]) (*connect.Response[v1.GetMetadataVersionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metadata.v1.MetadataService.GetMetadataVersion is not implemented"))
}

func (UnimplementedMetadataServiceHandler) ListMetadataVersions(context.Context, *connect.Request[v1.ListMetadataVersionsRequest]) (*connect.Response[v1.ListMetadataVersionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metadata.v1.MetadataService.ListMetadataVersions is not implemented"))
}
