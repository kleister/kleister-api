// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: teams/v1/teams.proto

package teamsv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/kleister/kleister-api/pkg/service/teams/v1"
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
	// TeamsServiceName is the fully-qualified name of the TeamsService service.
	TeamsServiceName = "teams.v1.TeamsService"
)

// TeamsServiceClient is a client for the teams.v1.TeamsService service.
type TeamsServiceClient interface {
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
	Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error)
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	Show(context.Context, *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error)
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
	ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error)
	AttachUser(context.Context, *connect_go.Request[v1.AttachUserRequest]) (*connect_go.Response[v1.AttachUserResponse], error)
	DropUser(context.Context, *connect_go.Request[v1.DropUserRequest]) (*connect_go.Response[v1.DropUserResponse], error)
	ListPacks(context.Context, *connect_go.Request[v1.ListPacksRequest]) (*connect_go.Response[v1.ListPacksResponse], error)
	AttachPack(context.Context, *connect_go.Request[v1.AttachPackRequest]) (*connect_go.Response[v1.AttachPackResponse], error)
	DropPack(context.Context, *connect_go.Request[v1.DropPackRequest]) (*connect_go.Response[v1.DropPackResponse], error)
	ListMods(context.Context, *connect_go.Request[v1.ListModsRequest]) (*connect_go.Response[v1.ListModsResponse], error)
	AttachMod(context.Context, *connect_go.Request[v1.AttachModRequest]) (*connect_go.Response[v1.AttachModResponse], error)
	DropMod(context.Context, *connect_go.Request[v1.DropModRequest]) (*connect_go.Response[v1.DropModResponse], error)
}

// NewTeamsServiceClient constructs a client for the teams.v1.TeamsService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTeamsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TeamsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &teamsServiceClient{
		list: connect_go.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/List",
			opts...,
		),
		create: connect_go.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/Create",
			opts...,
		),
		update: connect_go.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/Update",
			opts...,
		),
		show: connect_go.NewClient[v1.ShowRequest, v1.ShowResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/Show",
			opts...,
		),
		delete: connect_go.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/Delete",
			opts...,
		),
		listUsers: connect_go.NewClient[v1.ListUsersRequest, v1.ListUsersResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/ListUsers",
			opts...,
		),
		attachUser: connect_go.NewClient[v1.AttachUserRequest, v1.AttachUserResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/AttachUser",
			opts...,
		),
		dropUser: connect_go.NewClient[v1.DropUserRequest, v1.DropUserResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/DropUser",
			opts...,
		),
		listPacks: connect_go.NewClient[v1.ListPacksRequest, v1.ListPacksResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/ListPacks",
			opts...,
		),
		attachPack: connect_go.NewClient[v1.AttachPackRequest, v1.AttachPackResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/AttachPack",
			opts...,
		),
		dropPack: connect_go.NewClient[v1.DropPackRequest, v1.DropPackResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/DropPack",
			opts...,
		),
		listMods: connect_go.NewClient[v1.ListModsRequest, v1.ListModsResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/ListMods",
			opts...,
		),
		attachMod: connect_go.NewClient[v1.AttachModRequest, v1.AttachModResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/AttachMod",
			opts...,
		),
		dropMod: connect_go.NewClient[v1.DropModRequest, v1.DropModResponse](
			httpClient,
			baseURL+"/teams.v1.TeamsService/DropMod",
			opts...,
		),
	}
}

// teamsServiceClient implements TeamsServiceClient.
type teamsServiceClient struct {
	list       *connect_go.Client[v1.ListRequest, v1.ListResponse]
	create     *connect_go.Client[v1.CreateRequest, v1.CreateResponse]
	update     *connect_go.Client[v1.UpdateRequest, v1.UpdateResponse]
	show       *connect_go.Client[v1.ShowRequest, v1.ShowResponse]
	delete     *connect_go.Client[v1.DeleteRequest, v1.DeleteResponse]
	listUsers  *connect_go.Client[v1.ListUsersRequest, v1.ListUsersResponse]
	attachUser *connect_go.Client[v1.AttachUserRequest, v1.AttachUserResponse]
	dropUser   *connect_go.Client[v1.DropUserRequest, v1.DropUserResponse]
	listPacks  *connect_go.Client[v1.ListPacksRequest, v1.ListPacksResponse]
	attachPack *connect_go.Client[v1.AttachPackRequest, v1.AttachPackResponse]
	dropPack   *connect_go.Client[v1.DropPackRequest, v1.DropPackResponse]
	listMods   *connect_go.Client[v1.ListModsRequest, v1.ListModsResponse]
	attachMod  *connect_go.Client[v1.AttachModRequest, v1.AttachModResponse]
	dropMod    *connect_go.Client[v1.DropModRequest, v1.DropModResponse]
}

// List calls teams.v1.TeamsService.List.
func (c *teamsServiceClient) List(ctx context.Context, req *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Create calls teams.v1.TeamsService.Create.
func (c *teamsServiceClient) Create(ctx context.Context, req *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Update calls teams.v1.TeamsService.Update.
func (c *teamsServiceClient) Update(ctx context.Context, req *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Show calls teams.v1.TeamsService.Show.
func (c *teamsServiceClient) Show(ctx context.Context, req *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error) {
	return c.show.CallUnary(ctx, req)
}

// Delete calls teams.v1.TeamsService.Delete.
func (c *teamsServiceClient) Delete(ctx context.Context, req *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// ListUsers calls teams.v1.TeamsService.ListUsers.
func (c *teamsServiceClient) ListUsers(ctx context.Context, req *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error) {
	return c.listUsers.CallUnary(ctx, req)
}

// AttachUser calls teams.v1.TeamsService.AttachUser.
func (c *teamsServiceClient) AttachUser(ctx context.Context, req *connect_go.Request[v1.AttachUserRequest]) (*connect_go.Response[v1.AttachUserResponse], error) {
	return c.attachUser.CallUnary(ctx, req)
}

// DropUser calls teams.v1.TeamsService.DropUser.
func (c *teamsServiceClient) DropUser(ctx context.Context, req *connect_go.Request[v1.DropUserRequest]) (*connect_go.Response[v1.DropUserResponse], error) {
	return c.dropUser.CallUnary(ctx, req)
}

// ListPacks calls teams.v1.TeamsService.ListPacks.
func (c *teamsServiceClient) ListPacks(ctx context.Context, req *connect_go.Request[v1.ListPacksRequest]) (*connect_go.Response[v1.ListPacksResponse], error) {
	return c.listPacks.CallUnary(ctx, req)
}

// AttachPack calls teams.v1.TeamsService.AttachPack.
func (c *teamsServiceClient) AttachPack(ctx context.Context, req *connect_go.Request[v1.AttachPackRequest]) (*connect_go.Response[v1.AttachPackResponse], error) {
	return c.attachPack.CallUnary(ctx, req)
}

// DropPack calls teams.v1.TeamsService.DropPack.
func (c *teamsServiceClient) DropPack(ctx context.Context, req *connect_go.Request[v1.DropPackRequest]) (*connect_go.Response[v1.DropPackResponse], error) {
	return c.dropPack.CallUnary(ctx, req)
}

// ListMods calls teams.v1.TeamsService.ListMods.
func (c *teamsServiceClient) ListMods(ctx context.Context, req *connect_go.Request[v1.ListModsRequest]) (*connect_go.Response[v1.ListModsResponse], error) {
	return c.listMods.CallUnary(ctx, req)
}

// AttachMod calls teams.v1.TeamsService.AttachMod.
func (c *teamsServiceClient) AttachMod(ctx context.Context, req *connect_go.Request[v1.AttachModRequest]) (*connect_go.Response[v1.AttachModResponse], error) {
	return c.attachMod.CallUnary(ctx, req)
}

// DropMod calls teams.v1.TeamsService.DropMod.
func (c *teamsServiceClient) DropMod(ctx context.Context, req *connect_go.Request[v1.DropModRequest]) (*connect_go.Response[v1.DropModResponse], error) {
	return c.dropMod.CallUnary(ctx, req)
}

// TeamsServiceHandler is an implementation of the teams.v1.TeamsService service.
type TeamsServiceHandler interface {
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
	Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error)
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	Show(context.Context, *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error)
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
	ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error)
	AttachUser(context.Context, *connect_go.Request[v1.AttachUserRequest]) (*connect_go.Response[v1.AttachUserResponse], error)
	DropUser(context.Context, *connect_go.Request[v1.DropUserRequest]) (*connect_go.Response[v1.DropUserResponse], error)
	ListPacks(context.Context, *connect_go.Request[v1.ListPacksRequest]) (*connect_go.Response[v1.ListPacksResponse], error)
	AttachPack(context.Context, *connect_go.Request[v1.AttachPackRequest]) (*connect_go.Response[v1.AttachPackResponse], error)
	DropPack(context.Context, *connect_go.Request[v1.DropPackRequest]) (*connect_go.Response[v1.DropPackResponse], error)
	ListMods(context.Context, *connect_go.Request[v1.ListModsRequest]) (*connect_go.Response[v1.ListModsResponse], error)
	AttachMod(context.Context, *connect_go.Request[v1.AttachModRequest]) (*connect_go.Response[v1.AttachModResponse], error)
	DropMod(context.Context, *connect_go.Request[v1.DropModRequest]) (*connect_go.Response[v1.DropModResponse], error)
}

// NewTeamsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTeamsServiceHandler(svc TeamsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/teams.v1.TeamsService/List", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/List",
		svc.List,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/Create", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/Create",
		svc.Create,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/Update", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/Update",
		svc.Update,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/Show", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/Show",
		svc.Show,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/Delete", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/Delete",
		svc.Delete,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/ListUsers", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/ListUsers",
		svc.ListUsers,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/AttachUser", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/AttachUser",
		svc.AttachUser,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/DropUser", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/DropUser",
		svc.DropUser,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/ListPacks", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/ListPacks",
		svc.ListPacks,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/AttachPack", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/AttachPack",
		svc.AttachPack,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/DropPack", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/DropPack",
		svc.DropPack,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/ListMods", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/ListMods",
		svc.ListMods,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/AttachMod", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/AttachMod",
		svc.AttachMod,
		opts...,
	))
	mux.Handle("/teams.v1.TeamsService/DropMod", connect_go.NewUnaryHandler(
		"/teams.v1.TeamsService/DropMod",
		svc.DropMod,
		opts...,
	))
	return "/teams.v1.TeamsService/", mux
}

// UnimplementedTeamsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTeamsServiceHandler struct{}

func (UnimplementedTeamsServiceHandler) List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.List is not implemented"))
}

func (UnimplementedTeamsServiceHandler) Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.Create is not implemented"))
}

func (UnimplementedTeamsServiceHandler) Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.Update is not implemented"))
}

func (UnimplementedTeamsServiceHandler) Show(context.Context, *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.Show is not implemented"))
}

func (UnimplementedTeamsServiceHandler) Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.Delete is not implemented"))
}

func (UnimplementedTeamsServiceHandler) ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.ListUsers is not implemented"))
}

func (UnimplementedTeamsServiceHandler) AttachUser(context.Context, *connect_go.Request[v1.AttachUserRequest]) (*connect_go.Response[v1.AttachUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.AttachUser is not implemented"))
}

func (UnimplementedTeamsServiceHandler) DropUser(context.Context, *connect_go.Request[v1.DropUserRequest]) (*connect_go.Response[v1.DropUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.DropUser is not implemented"))
}

func (UnimplementedTeamsServiceHandler) ListPacks(context.Context, *connect_go.Request[v1.ListPacksRequest]) (*connect_go.Response[v1.ListPacksResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.ListPacks is not implemented"))
}

func (UnimplementedTeamsServiceHandler) AttachPack(context.Context, *connect_go.Request[v1.AttachPackRequest]) (*connect_go.Response[v1.AttachPackResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.AttachPack is not implemented"))
}

func (UnimplementedTeamsServiceHandler) DropPack(context.Context, *connect_go.Request[v1.DropPackRequest]) (*connect_go.Response[v1.DropPackResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.DropPack is not implemented"))
}

func (UnimplementedTeamsServiceHandler) ListMods(context.Context, *connect_go.Request[v1.ListModsRequest]) (*connect_go.Response[v1.ListModsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.ListMods is not implemented"))
}

func (UnimplementedTeamsServiceHandler) AttachMod(context.Context, *connect_go.Request[v1.AttachModRequest]) (*connect_go.Response[v1.AttachModResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.AttachMod is not implemented"))
}

func (UnimplementedTeamsServiceHandler) DropMod(context.Context, *connect_go.Request[v1.DropModRequest]) (*connect_go.Response[v1.DropModResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("teams.v1.TeamsService.DropMod is not implemented"))
}
