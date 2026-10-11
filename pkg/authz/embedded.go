package authz

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/kleister/kleister-api/pkg/config"
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/openfga/openfga/pkg/server"
	serverconfig "github.com/openfga/openfga/pkg/server/config"
	"github.com/openfga/openfga/pkg/storage"
	"github.com/openfga/openfga/pkg/storage/migrate"
	"github.com/openfga/openfga/pkg/storage/mysql"
	"github.com/openfga/openfga/pkg/storage/postgres"
	"github.com/openfga/openfga/pkg/storage/sqlcommon"
	"github.com/openfga/openfga/pkg/storage/sqlite"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// storeName is the deterministic OpenFGA store name used for the embedded
// driver, looked up by name on boot rather than persisting a generated id.
const storeName = "kleister"

// defaultAuthzName is used when no name is configured for the embedded
// sqlite3 driver.
const defaultAuthzName = "authz.sqlite3"

// resolveCredentials resolves the username/password for a networked
// embedded driver, supporting the same file:// and base64:// secret
// prefixes config.Database's credentials do.
func resolveCredentials(cfg config.Authz) (username, password string, err error) {
	username, err = config.Value(cfg.Username)

	if err != nil {
		return "", "", fmt.Errorf("failed to parse authz username secret: %w", err)
	}

	password, err = config.Value(cfg.Password)

	if err != nil {
		return "", "", fmt.Errorf("failed to parse authz password secret: %w", err)
	}

	return username, password, nil
}

// embeddedDSN builds the connection string and the engine name expected by
// OpenFGA's migrate/storage packages ("sqlite", "mysql" or "postgres") from
// the same shape of fields config.Database uses (driver/address/port/
// username/password/name/options), rather than a single DSN string.
func embeddedDSN(cfg config.Authz) (engine, uri string, err error) {
	switch cfg.Driver {
	case "", "sqlite", "sqlite3":
		name := cfg.Name

		if name == "" {
			name = defaultAuthzName
		}

		if len(cfg.Options) == 0 {
			return "sqlite", name, nil
		}

		meta := url.Values{}

		for key, val := range cfg.Options {
			meta.Set(key, val)
		}

		return "sqlite", fmt.Sprintf("%s?%s", name, meta.Encode()), nil

	case "mysql", "mariadb":
		username, password, err := resolveCredentials(cfg)

		if err != nil {
			return "", "", err
		}

		port := cfg.Port

		if port == "" {
			port = "3306"
		}

		meta := url.Values{}

		for key, val := range cfg.Options {
			meta.Set(key, val)
		}

		if !meta.Has("charset") {
			meta.Set("charset", "utf8")
		}

		if !meta.Has("parseTime") {
			meta.Set("parseTime", "True")
		}

		if !meta.Has("loc") {
			meta.Set("loc", "Local")
		}

		auth := username

		if password != "" {
			auth = fmt.Sprintf("%s:%s", username, password)
		}

		return "mysql", fmt.Sprintf(
			"%s@tcp(%s:%s)/%s?%s",
			auth,
			cfg.Address,
			port,
			cfg.Name,
			meta.Encode(),
		), nil

	case "postgres", "postgresql":
		username, password, err := resolveCredentials(cfg)

		if err != nil {
			return "", "", err
		}

		port := cfg.Port

		if port == "" {
			port = "5432"
		}

		u := &url.URL{
			Scheme: "postgres",
			Host:   fmt.Sprintf("%s:%s", cfg.Address, port),
			Path:   "/" + cfg.Name,
		}

		if username != "" {
			if password != "" {
				u.User = url.UserPassword(username, password)
			} else {
				u.User = url.User(username)
			}
		}

		meta := url.Values{}

		for key, val := range cfg.Options {
			meta.Set(key, val)
		}

		if !meta.Has("sslmode") {
			meta.Set("sslmode", "disable")
		}

		u.RawQuery = meta.Encode()

		return "postgres", u.String(), nil

	default:
		return "", "", fmt.Errorf("authz: unknown database driver: %s", cfg.Driver)
	}
}

func openEmbeddedDatastore(engine, uri string, cfg *sqlcommon.Config) (storage.OpenFGADatastore, error) {
	switch engine {
	case "sqlite":
		return sqlite.New(uri, cfg)
	case "mysql":
		return mysql.New(uri, cfg)
	case "postgres":
		return postgres.New(uri, cfg)
	default:
		return nil, fmt.Errorf("authz: unknown database driver: %s", engine)
	}
}

// embedded wraps an in-process OpenFGA server for single-binary deployments.
type embedded struct {
	server  *server.Server
	storeID string
	modelID string
}

var _ Checker = (*embedded)(nil)

func newEmbedded(cfg config.Authz) (Checker, error) {
	engine, uri, err := embeddedDSN(cfg)

	if err != nil {
		return nil, err
	}

	if err := migrate.RunMigrations(migrate.MigrationConfig{
		Engine:      engine,
		URI:         uri,
		Timeout:     30 * time.Second,
		PingTimeout: 5 * time.Second,
	}); err != nil {
		return nil, fmt.Errorf("failed to migrate authz datastore: %w", err)
	}

	ds, err := openEmbeddedDatastore(engine, uri, &sqlcommon.Config{
		MaxTuplesPerWriteField: serverconfig.DefaultMaxTuplesPerWrite,
		MaxTypesPerModelField:  serverconfig.DefaultMaxTypesPerAuthorizationModel,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to open authz datastore: %w", err)
	}

	srv, err := server.NewServerWithOpts(
		server.WithDatastore(ds),
		server.WithMaxAuthorizationModelSizeInBytes(serverconfig.DefaultMaxAuthorizationModelSizeInBytes),
		server.WithResolveNodeLimit(serverconfig.DefaultResolveNodeLimit),
		server.WithResolveNodeBreadthLimit(serverconfig.DefaultResolveNodeBreadthLimit),
	)

	if err != nil {
		ds.Close()

		return nil, fmt.Errorf("failed to start embedded authz server: %w", err)
	}

	e := &embedded{
		server: srv,
	}

	if err := e.bootstrap(context.Background()); err != nil {
		srv.Close()

		return nil, err
	}

	return e, nil
}

func (e *embedded) bootstrap(ctx context.Context) error {
	storeID, err := e.resolveStore(ctx)

	if err != nil {
		return err
	}

	e.storeID = storeID

	model, err := compileModel()

	if err != nil {
		return err
	}

	digest, err := modelDigest(model)

	if err != nil {
		return err
	}

	existing, err := e.server.ReadAuthorizationModels(ctx, &openfgav1.ReadAuthorizationModelsRequest{
		StoreId:  storeID,
		PageSize: wrapperspb.Int32(1),
	})

	if err != nil {
		return fmt.Errorf("failed to read existing authorization models: %w", err)
	}

	if models := existing.GetAuthorizationModels(); len(models) > 0 {
		latest := models[0]
		latestDigest, err := modelDigest(latest)

		if err != nil {
			return err
		}

		if latestDigest == digest {
			e.modelID = latest.GetId()

			return nil
		}
	}

	written, err := e.server.WriteAuthorizationModel(ctx, &openfgav1.WriteAuthorizationModelRequest{
		StoreId:         storeID,
		SchemaVersion:   model.GetSchemaVersion(),
		TypeDefinitions: model.GetTypeDefinitions(),
		Conditions:      model.GetConditions(),
	})

	if err != nil {
		return fmt.Errorf("failed to write authorization model: %w", err)
	}

	e.modelID = written.GetAuthorizationModelId()

	return nil
}

func (e *embedded) resolveStore(ctx context.Context) (string, error) {
	existing, err := e.server.ListStores(ctx, &openfgav1.ListStoresRequest{
		Name: storeName,
	})

	if err != nil {
		return "", fmt.Errorf("failed to list authz stores: %w", err)
	}

	for _, s := range existing.GetStores() {
		if s.GetName() == storeName {
			return s.GetId(), nil
		}
	}

	created, err := e.server.CreateStore(ctx, &openfgav1.CreateStoreRequest{
		Name: storeName,
	})

	if err != nil {
		return "", fmt.Errorf("failed to create authz store: %w", err)
	}

	return created.GetId(), nil
}

func (e *embedded) Check(ctx context.Context, subject, relation, object string) (bool, error) {
	resp, err := e.server.Check(ctx, &openfgav1.CheckRequest{
		StoreId:              e.storeID,
		AuthorizationModelId: e.modelID,
		TupleKey: &openfgav1.CheckRequestTupleKey{
			User:     subject,
			Relation: relation,
			Object:   object,
		},
	})

	if err != nil {
		return false, err
	}

	return resp.GetAllowed(), nil
}

func (e *embedded) Write(ctx context.Context, tuples ...TupleKey) error {
	if len(tuples) == 0 {
		return nil
	}

	keys := make([]*openfgav1.TupleKey, 0, len(tuples))

	for _, t := range tuples {
		keys = append(keys, &openfgav1.TupleKey{
			User:     t.Subject,
			Relation: t.Relation,
			Object:   t.Object,
		})
	}

	_, err := e.server.Write(ctx, &openfgav1.WriteRequest{
		StoreId:              e.storeID,
		AuthorizationModelId: e.modelID,
		Writes: &openfgav1.WriteRequestWrites{
			TupleKeys:   keys,
			OnDuplicate: "ignore",
		},
	})

	return err
}

func (e *embedded) Delete(ctx context.Context, tuples ...TupleKey) error {
	if len(tuples) == 0 {
		return nil
	}

	keys := make([]*openfgav1.TupleKeyWithoutCondition, 0, len(tuples))

	for _, t := range tuples {
		keys = append(keys, &openfgav1.TupleKeyWithoutCondition{
			User:     t.Subject,
			Relation: t.Relation,
			Object:   t.Object,
		})
	}

	_, err := e.server.Write(ctx, &openfgav1.WriteRequest{
		StoreId:              e.storeID,
		AuthorizationModelId: e.modelID,
		Deletes: &openfgav1.WriteRequestDeletes{
			TupleKeys: keys,
			OnMissing: "ignore",
		},
	})

	return err
}

func (e *embedded) Read(ctx context.Context, object string, relations ...string) ([]TupleKey, error) {
	return e.read(ctx, &openfgav1.ReadRequestTupleKey{Object: object}, relations...)
}

func (e *embedded) ReadByUser(ctx context.Context, subject, objectType string, relations ...string) ([]TupleKey, error) {
	return e.read(ctx, &openfgav1.ReadRequestTupleKey{Object: objectType + ":", User: subject}, relations...)
}

func (e *embedded) ReadAll(ctx context.Context) ([]TupleKey, error) {
	return e.read(ctx, nil)
}

func (e *embedded) read(ctx context.Context, tupleKey *openfgav1.ReadRequestTupleKey, relations ...string) ([]TupleKey, error) {
	byRelation := make(map[string][]TupleKey)
	wanted := make(map[string]bool, len(relations))

	for _, r := range relations {
		wanted[r] = true
	}

	var (
		token string
		out   []TupleKey
	)

	for {
		resp, err := e.server.Read(ctx, &openfgav1.ReadRequest{
			StoreId:           e.storeID,
			TupleKey:          tupleKey,
			ContinuationToken: token,
		})

		if err != nil {
			return nil, err
		}

		for _, tuple := range resp.GetTuples() {
			key := tuple.GetKey()

			if len(wanted) > 0 && !wanted[key.GetRelation()] {
				continue
			}

			byRelation[key.GetRelation()] = append(byRelation[key.GetRelation()], TupleKey{
				Object:    key.GetObject(),
				Relation:  key.GetRelation(),
				Subject:   key.GetUser(),
				CreatedAt: tuple.GetTimestamp().AsTime(),
			})
		}

		token = resp.GetContinuationToken()

		if token == "" {
			break
		}
	}

	for _, tuples := range byRelation {
		out = append(out, tuples...)
	}

	return out, nil
}

func (e *embedded) ListObjects(ctx context.Context, subject, relation, objectType string) ([]string, error) {
	resp, err := e.server.ListObjects(ctx, &openfgav1.ListObjectsRequest{
		StoreId:              e.storeID,
		AuthorizationModelId: e.modelID,
		Type:                 objectType,
		Relation:             relation,
		User:                 subject,
	})

	if err != nil {
		return nil, err
	}

	return resp.GetObjects(), nil
}

func (e *embedded) Ready(ctx context.Context) error {
	ready, err := e.server.IsReady(ctx)

	if err != nil {
		return err
	}

	if !ready {
		return fmt.Errorf("authz server is not ready")
	}

	return nil
}

func (e *embedded) Close() error {
	e.server.Close()
	return nil
}
