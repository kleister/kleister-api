// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: forge/v1/forge.proto

package forgev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/kleister/kleister-api/pkg/service/forge/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ForgeServiceName is the fully-qualified name of the ForgeService service.
	ForgeServiceName = "forge.v1.ForgeService"
)

// ForgeServiceClient is a client for the forge.v1.ForgeService service.
type ForgeServiceClient interface {
	Search(context.Context, *connect_go.Request[v1.SearchRequest]) (*connect_go.Response[v1.SearchResponse], error)
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	ListBuilds(context.Context, *connect_go.Request[v1.ListBuildsRequest]) (*connect_go.Response[v1.ListBuildsResponse], error)
}

// NewForgeServiceClient constructs a client for the forge.v1.ForgeService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewForgeServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ForgeServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &forgeServiceClient{
		search: connect_go.NewClient[v1.SearchRequest, v1.SearchResponse](
			httpClient,
			baseURL+"/forge.v1.ForgeService/Search",
			opts...,
		),
		update: connect_go.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+"/forge.v1.ForgeService/Update",
			opts...,
		),
		listBuilds: connect_go.NewClient[v1.ListBuildsRequest, v1.ListBuildsResponse](
			httpClient,
			baseURL+"/forge.v1.ForgeService/ListBuilds",
			opts...,
		),
	}
}

// forgeServiceClient implements ForgeServiceClient.
type forgeServiceClient struct {
	search     *connect_go.Client[v1.SearchRequest, v1.SearchResponse]
	update     *connect_go.Client[v1.UpdateRequest, v1.UpdateResponse]
	listBuilds *connect_go.Client[v1.ListBuildsRequest, v1.ListBuildsResponse]
}

// Search calls forge.v1.ForgeService.Search.
func (c *forgeServiceClient) Search(ctx context.Context, req *connect_go.Request[v1.SearchRequest]) (*connect_go.Response[v1.SearchResponse], error) {
	return c.search.CallUnary(ctx, req)
}

// Update calls forge.v1.ForgeService.Update.
func (c *forgeServiceClient) Update(ctx context.Context, req *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// ListBuilds calls forge.v1.ForgeService.ListBuilds.
func (c *forgeServiceClient) ListBuilds(ctx context.Context, req *connect_go.Request[v1.ListBuildsRequest]) (*connect_go.Response[v1.ListBuildsResponse], error) {
	return c.listBuilds.CallUnary(ctx, req)
}

// ForgeServiceHandler is an implementation of the forge.v1.ForgeService service.
type ForgeServiceHandler interface {
	Search(context.Context, *connect_go.Request[v1.SearchRequest]) (*connect_go.Response[v1.SearchResponse], error)
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	ListBuilds(context.Context, *connect_go.Request[v1.ListBuildsRequest]) (*connect_go.Response[v1.ListBuildsResponse], error)
}

// NewForgeServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewForgeServiceHandler(svc ForgeServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/forge.v1.ForgeService/Search", connect_go.NewUnaryHandler(
		"/forge.v1.ForgeService/Search",
		svc.Search,
		opts...,
	))
	mux.Handle("/forge.v1.ForgeService/Update", connect_go.NewUnaryHandler(
		"/forge.v1.ForgeService/Update",
		svc.Update,
		opts...,
	))
	mux.Handle("/forge.v1.ForgeService/ListBuilds", connect_go.NewUnaryHandler(
		"/forge.v1.ForgeService/ListBuilds",
		svc.ListBuilds,
		opts...,
	))
	return "/forge.v1.ForgeService/", mux
}

// UnimplementedForgeServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedForgeServiceHandler struct{}

func (UnimplementedForgeServiceHandler) Search(context.Context, *connect_go.Request[v1.SearchRequest]) (*connect_go.Response[v1.SearchResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("forge.v1.ForgeService.Search is not implemented"))
}

func (UnimplementedForgeServiceHandler) Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("forge.v1.ForgeService.Update is not implemented"))
}

func (UnimplementedForgeServiceHandler) ListBuilds(context.Context, *connect_go.Request[v1.ListBuildsRequest]) (*connect_go.Response[v1.ListBuildsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("forge.v1.ForgeService.ListBuilds is not implemented"))
}
