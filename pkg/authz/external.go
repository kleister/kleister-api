package authz

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	fgaSdk "github.com/openfga/go-sdk"
	fgaClient "github.com/openfga/go-sdk/client"
	"github.com/openfga/go-sdk/credentials"
	"github.com/openfga/language/pkg/go/transformer"

	"github.com/kleister/kleister-api/pkg/config"
)

// external wraps a client talking to a standalone OpenFGA service, used for
// larger deployments that run OpenFGA as its own service.
type external struct {
	client  *fgaClient.OpenFgaClient
	storeID string
	modelID string
}

var _ Checker = (*external)(nil)

func newExternal(cfg config.Authz) (Checker, error) {
	if cfg.API == "" {
		return nil, fmt.Errorf("authz: api is required for the external driver")
	}

	clientCfg := &fgaClient.ClientConfiguration{
		ApiUrl:  cfg.API,
		StoreId: cfg.Store,
	}

	if cfg.Token != "" {
		token, err := config.Value(cfg.Token)

		if err != nil {
			return nil, fmt.Errorf("failed to parse authz token secret: %w", err)
		}

		clientCfg.Credentials = &credentials.Credentials{
			Method: credentials.CredentialsMethodApiToken,
			Config: &credentials.Config{
				ApiToken: token,
			},
		}
	}

	client, err := fgaClient.NewSdkClient(clientCfg)

	if err != nil {
		return nil, fmt.Errorf("failed to create authz client: %w", err)
	}

	e := &external{
		client:  client,
		storeID: cfg.Store,
	}

	if err := e.bootstrap(context.Background()); err != nil {
		return nil, err
	}

	return e, nil
}

type compiledModel struct {
	SchemaVersion   string                       `json:"schema_version"`
	TypeDefinitions []fgaSdk.TypeDefinition      `json:"type_definitions"`
	Conditions      *map[string]fgaSdk.Condition `json:"conditions,omitempty"`
}

func (e *external) bootstrap(ctx context.Context) error {
	if e.storeID == "" {
		storeID, err := e.resolveStore(ctx)

		if err != nil {
			return err
		}

		e.storeID = storeID
	}

	data, err := transformer.TransformDSLToJSON(modelDSL)

	if err != nil {
		return fmt.Errorf("failed to compile authorization model: %w", err)
	}

	var compiled compiledModel

	if err := json.Unmarshal([]byte(data), &compiled); err != nil {
		return fmt.Errorf("failed to parse compiled authorization model: %w", err)
	}

	digest, err := canonicalDigest(compiled.SchemaVersion, compiled.TypeDefinitions, compiled.Conditions)

	if err != nil {
		return err
	}

	one := int32(1)

	existing, err := e.client.ReadAuthorizationModels(ctx).
		Options(fgaClient.ClientReadAuthorizationModelsOptions{
			PageSize: &one,
			StoreId:  &e.storeID,
		}).
		Execute()

	if err != nil {
		return fmt.Errorf("failed to read existing authorization models: %w", err)
	}

	if models := existing.GetAuthorizationModels(); len(models) > 0 {
		latest := models[0]
		latestDigest, err := canonicalDigest(latest.SchemaVersion, latest.TypeDefinitions, latest.Conditions)

		if err != nil {
			return err
		}

		if latestDigest == digest {
			e.modelID = latest.Id

			return nil
		}
	}

	written, err := e.client.WriteAuthorizationModel(ctx).
		Body(fgaClient.ClientWriteAuthorizationModelRequest{
			SchemaVersion:   compiled.SchemaVersion,
			TypeDefinitions: compiled.TypeDefinitions,
			Conditions:      compiled.Conditions,
		}).
		Options(fgaClient.ClientWriteAuthorizationModelOptions{
			StoreId: &e.storeID,
		}).
		Execute()

	if err != nil {
		return fmt.Errorf("failed to write authorization model: %w", err)
	}

	e.modelID = written.AuthorizationModelId

	return nil
}

func canonicalDigest(schemaVersion string, typeDefinitions []fgaSdk.TypeDefinition, conditions *map[string]fgaSdk.Condition) (string, error) {
	data, err := json.Marshal(compiledModel{
		SchemaVersion:   schemaVersion,
		TypeDefinitions: typeDefinitions,
		Conditions:      conditions,
	})

	if err != nil {
		return "", fmt.Errorf("failed to marshal authorization model: %w", err)
	}

	sum := sha256.Sum256(data)

	return fmt.Sprintf("%x", sum[:]), nil
}

func (e *external) resolveStore(ctx context.Context) (string, error) {
	name := storeName

	existing, err := e.client.ListStores(ctx).
		Options(fgaClient.ClientListStoresOptions{
			Name: &name,
		}).
		Execute()

	if err != nil {
		return "", fmt.Errorf("failed to list authz stores: %w", err)
	}

	for _, s := range existing.GetStores() {
		if s.Name == storeName {
			return s.Id, nil
		}
	}

	created, err := e.client.CreateStore(ctx).
		Body(fgaClient.ClientCreateStoreRequest{
			Name: storeName,
		}).
		Execute()

	if err != nil {
		return "", fmt.Errorf("failed to create authz store: %w", err)
	}

	return created.Id, nil
}

func (e *external) Check(ctx context.Context, subject, relation, object string) (bool, error) {
	resp, err := e.client.Check(ctx).
		Body(fgaClient.ClientCheckRequest{
			User:     subject,
			Relation: relation,
			Object:   object,
		}).
		Options(fgaClient.ClientCheckOptions{
			StoreId:              &e.storeID,
			AuthorizationModelId: &e.modelID,
		}).
		Execute()

	if err != nil {
		return false, err
	}

	return resp.GetAllowed(), nil
}

func (e *external) Write(ctx context.Context, tuples ...TupleKey) error {
	if len(tuples) == 0 {
		return nil
	}

	writes := make([]fgaClient.ClientTupleKey, 0, len(tuples))

	for _, t := range tuples {
		writes = append(writes, fgaClient.ClientTupleKey{
			User:     t.Subject,
			Relation: t.Relation,
			Object:   t.Object,
		})
	}

	_, err := e.client.Write(ctx).
		Body(fgaClient.ClientWriteRequest{
			Writes: writes,
		}).
		Options(fgaClient.ClientWriteOptions{
			StoreId:              &e.storeID,
			AuthorizationModelId: &e.modelID,
			Conflict: fgaClient.ClientWriteConflictOptions{
				OnDuplicateWrites: fgaClient.CLIENT_WRITE_REQUEST_ON_DUPLICATE_WRITES_IGNORE,
			},
		}).
		Execute()

	return err
}

func (e *external) Delete(ctx context.Context, tuples ...TupleKey) error {
	if len(tuples) == 0 {
		return nil
	}

	deletes := make([]fgaClient.ClientTupleKeyWithoutCondition, 0, len(tuples))

	for _, t := range tuples {
		deletes = append(deletes, fgaClient.ClientTupleKeyWithoutCondition{
			User:     t.Subject,
			Relation: t.Relation,
			Object:   t.Object,
		})
	}

	_, err := e.client.Write(ctx).
		Body(fgaClient.ClientWriteRequest{
			Deletes: deletes,
		}).
		Options(fgaClient.ClientWriteOptions{
			StoreId:              &e.storeID,
			AuthorizationModelId: &e.modelID,
			Conflict: fgaClient.ClientWriteConflictOptions{
				OnMissingDeletes: fgaClient.CLIENT_WRITE_REQUEST_ON_MISSING_DELETES_IGNORE,
			},
		}).
		Execute()

	return err
}

func (e *external) Read(ctx context.Context, object string, relations ...string) ([]TupleKey, error) {
	return e.read(ctx, fgaClient.ClientReadRequest{Object: &object}, relations...)
}

func (e *external) ReadByUser(ctx context.Context, subject, objectType string, relations ...string) ([]TupleKey, error) {
	typeOnly := objectType + ":"

	return e.read(ctx, fgaClient.ClientReadRequest{Object: &typeOnly, User: &subject}, relations...)
}

func (e *external) ReadAll(ctx context.Context) ([]TupleKey, error) {
	return e.read(ctx, fgaClient.ClientReadRequest{})
}

func (e *external) read(ctx context.Context, body fgaClient.ClientReadRequest, relations ...string) ([]TupleKey, error) {
	wanted := make(map[string]bool, len(relations))

	for _, r := range relations {
		wanted[r] = true
	}

	var (
		token string
		out   []TupleKey
	)

	for {
		opts := fgaClient.ClientReadOptions{
			StoreId: &e.storeID,
		}

		if token != "" {
			opts.ContinuationToken = &token
		}

		resp, err := e.client.Read(ctx).
			Body(body).
			Options(opts).
			Execute()

		if err != nil {
			return nil, err
		}

		for _, tuple := range resp.Tuples {
			key := tuple.Key

			if len(wanted) > 0 && !wanted[key.Relation] {
				continue
			}

			out = append(out, TupleKey{
				Object:    key.Object,
				Relation:  key.Relation,
				Subject:   key.User,
				CreatedAt: tuple.Timestamp,
			})
		}

		token = resp.ContinuationToken

		if token == "" {
			break
		}
	}

	return out, nil
}

func (e *external) ListObjects(ctx context.Context, subject, relation, objectType string) ([]string, error) {
	resp, err := e.client.ListObjects(ctx).
		Body(fgaClient.ClientListObjectsRequest{
			User:     subject,
			Relation: relation,
			Type:     objectType,
		}).
		Options(fgaClient.ClientListObjectsOptions{
			StoreId:              &e.storeID,
			AuthorizationModelId: &e.modelID,
		}).
		Execute()

	if err != nil {
		return nil, err
	}

	return resp.Objects, nil
}

func (e *external) Ready(ctx context.Context) error {
	_, err := e.client.ListStores(ctx).
		Options(fgaClient.ClientListStoresOptions{}).
		Execute()

	return err
}

func (e *external) Close() error {
	return nil
}
