// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: versions/v1/versions.proto

package versionsv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/kleister/kleister-api/pkg/service/versions/v1"
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
	// VersionsServiceName is the fully-qualified name of the VersionsService service.
	VersionsServiceName = "versions.v1.VersionsService"
)

// VersionsServiceClient is a client for the versions.v1.VersionsService service.
type VersionsServiceClient interface {
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
	Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error)
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	Show(context.Context, *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error)
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
	ListBuilds(context.Context, *connect_go.Request[v1.ListBuildsRequest]) (*connect_go.Response[v1.ListBuildsResponse], error)
	AttachBuild(context.Context, *connect_go.Request[v1.AttachBuildRequest]) (*connect_go.Response[v1.AttachBuildResponse], error)
	DropBuild(context.Context, *connect_go.Request[v1.DropBuildRequest]) (*connect_go.Response[v1.DropBuildResponse], error)
}

// NewVersionsServiceClient constructs a client for the versions.v1.VersionsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewVersionsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) VersionsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &versionsServiceClient{
		list: connect_go.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+"/versions.v1.VersionsService/List",
			opts...,
		),
		create: connect_go.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+"/versions.v1.VersionsService/Create",
			opts...,
		),
		update: connect_go.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+"/versions.v1.VersionsService/Update",
			opts...,
		),
		show: connect_go.NewClient[v1.ShowRequest, v1.ShowResponse](
			httpClient,
			baseURL+"/versions.v1.VersionsService/Show",
			opts...,
		),
		delete: connect_go.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+"/versions.v1.VersionsService/Delete",
			opts...,
		),
		listBuilds: connect_go.NewClient[v1.ListBuildsRequest, v1.ListBuildsResponse](
			httpClient,
			baseURL+"/versions.v1.VersionsService/ListBuilds",
			opts...,
		),
		attachBuild: connect_go.NewClient[v1.AttachBuildRequest, v1.AttachBuildResponse](
			httpClient,
			baseURL+"/versions.v1.VersionsService/AttachBuild",
			opts...,
		),
		dropBuild: connect_go.NewClient[v1.DropBuildRequest, v1.DropBuildResponse](
			httpClient,
			baseURL+"/versions.v1.VersionsService/DropBuild",
			opts...,
		),
	}
}

// versionsServiceClient implements VersionsServiceClient.
type versionsServiceClient struct {
	list        *connect_go.Client[v1.ListRequest, v1.ListResponse]
	create      *connect_go.Client[v1.CreateRequest, v1.CreateResponse]
	update      *connect_go.Client[v1.UpdateRequest, v1.UpdateResponse]
	show        *connect_go.Client[v1.ShowRequest, v1.ShowResponse]
	delete      *connect_go.Client[v1.DeleteRequest, v1.DeleteResponse]
	listBuilds  *connect_go.Client[v1.ListBuildsRequest, v1.ListBuildsResponse]
	attachBuild *connect_go.Client[v1.AttachBuildRequest, v1.AttachBuildResponse]
	dropBuild   *connect_go.Client[v1.DropBuildRequest, v1.DropBuildResponse]
}

// List calls versions.v1.VersionsService.List.
func (c *versionsServiceClient) List(ctx context.Context, req *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Create calls versions.v1.VersionsService.Create.
func (c *versionsServiceClient) Create(ctx context.Context, req *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Update calls versions.v1.VersionsService.Update.
func (c *versionsServiceClient) Update(ctx context.Context, req *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Show calls versions.v1.VersionsService.Show.
func (c *versionsServiceClient) Show(ctx context.Context, req *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error) {
	return c.show.CallUnary(ctx, req)
}

// Delete calls versions.v1.VersionsService.Delete.
func (c *versionsServiceClient) Delete(ctx context.Context, req *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// ListBuilds calls versions.v1.VersionsService.ListBuilds.
func (c *versionsServiceClient) ListBuilds(ctx context.Context, req *connect_go.Request[v1.ListBuildsRequest]) (*connect_go.Response[v1.ListBuildsResponse], error) {
	return c.listBuilds.CallUnary(ctx, req)
}

// AttachBuild calls versions.v1.VersionsService.AttachBuild.
func (c *versionsServiceClient) AttachBuild(ctx context.Context, req *connect_go.Request[v1.AttachBuildRequest]) (*connect_go.Response[v1.AttachBuildResponse], error) {
	return c.attachBuild.CallUnary(ctx, req)
}

// DropBuild calls versions.v1.VersionsService.DropBuild.
func (c *versionsServiceClient) DropBuild(ctx context.Context, req *connect_go.Request[v1.DropBuildRequest]) (*connect_go.Response[v1.DropBuildResponse], error) {
	return c.dropBuild.CallUnary(ctx, req)
}

// VersionsServiceHandler is an implementation of the versions.v1.VersionsService service.
type VersionsServiceHandler interface {
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
	Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error)
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	Show(context.Context, *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error)
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
	ListBuilds(context.Context, *connect_go.Request[v1.ListBuildsRequest]) (*connect_go.Response[v1.ListBuildsResponse], error)
	AttachBuild(context.Context, *connect_go.Request[v1.AttachBuildRequest]) (*connect_go.Response[v1.AttachBuildResponse], error)
	DropBuild(context.Context, *connect_go.Request[v1.DropBuildRequest]) (*connect_go.Response[v1.DropBuildResponse], error)
}

// NewVersionsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewVersionsServiceHandler(svc VersionsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/versions.v1.VersionsService/List", connect_go.NewUnaryHandler(
		"/versions.v1.VersionsService/List",
		svc.List,
		opts...,
	))
	mux.Handle("/versions.v1.VersionsService/Create", connect_go.NewUnaryHandler(
		"/versions.v1.VersionsService/Create",
		svc.Create,
		opts...,
	))
	mux.Handle("/versions.v1.VersionsService/Update", connect_go.NewUnaryHandler(
		"/versions.v1.VersionsService/Update",
		svc.Update,
		opts...,
	))
	mux.Handle("/versions.v1.VersionsService/Show", connect_go.NewUnaryHandler(
		"/versions.v1.VersionsService/Show",
		svc.Show,
		opts...,
	))
	mux.Handle("/versions.v1.VersionsService/Delete", connect_go.NewUnaryHandler(
		"/versions.v1.VersionsService/Delete",
		svc.Delete,
		opts...,
	))
	mux.Handle("/versions.v1.VersionsService/ListBuilds", connect_go.NewUnaryHandler(
		"/versions.v1.VersionsService/ListBuilds",
		svc.ListBuilds,
		opts...,
	))
	mux.Handle("/versions.v1.VersionsService/AttachBuild", connect_go.NewUnaryHandler(
		"/versions.v1.VersionsService/AttachBuild",
		svc.AttachBuild,
		opts...,
	))
	mux.Handle("/versions.v1.VersionsService/DropBuild", connect_go.NewUnaryHandler(
		"/versions.v1.VersionsService/DropBuild",
		svc.DropBuild,
		opts...,
	))
	return "/versions.v1.VersionsService/", mux
}

// UnimplementedVersionsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedVersionsServiceHandler struct{}

func (UnimplementedVersionsServiceHandler) List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("versions.v1.VersionsService.List is not implemented"))
}

func (UnimplementedVersionsServiceHandler) Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("versions.v1.VersionsService.Create is not implemented"))
}

func (UnimplementedVersionsServiceHandler) Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("versions.v1.VersionsService.Update is not implemented"))
}

func (UnimplementedVersionsServiceHandler) Show(context.Context, *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("versions.v1.VersionsService.Show is not implemented"))
}

func (UnimplementedVersionsServiceHandler) Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("versions.v1.VersionsService.Delete is not implemented"))
}

func (UnimplementedVersionsServiceHandler) ListBuilds(context.Context, *connect_go.Request[v1.ListBuildsRequest]) (*connect_go.Response[v1.ListBuildsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("versions.v1.VersionsService.ListBuilds is not implemented"))
}

func (UnimplementedVersionsServiceHandler) AttachBuild(context.Context, *connect_go.Request[v1.AttachBuildRequest]) (*connect_go.Response[v1.AttachBuildResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("versions.v1.VersionsService.AttachBuild is not implemented"))
}

func (UnimplementedVersionsServiceHandler) DropBuild(context.Context, *connect_go.Request[v1.DropBuildRequest]) (*connect_go.Response[v1.DropBuildResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("versions.v1.VersionsService.DropBuild is not implemented"))
}