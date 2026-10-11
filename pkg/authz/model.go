package authz

import (
	"crypto/sha256"
	_ "embed"
	"fmt"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/openfga/language/pkg/go/transformer"
	"google.golang.org/protobuf/proto"
)

//go:embed model.fga
var modelDSL string

// compileModel compiles the embedded DSL authorization model (model.fga)
// into its proto representation. The DSL file is the single source of truth
// for the authorization model and is checked into git like a migration; the
// compiled JSON/proto form is never hand-authored or committed.
func compileModel() (*openfgav1.AuthorizationModel, error) {
	model, err := transformer.TransformDSLToProto(modelDSL)

	if err != nil {
		return nil, fmt.Errorf("failed to compile authorization model: %w", err)
	}

	return model, nil
}

// modelDigest returns a stable hash of a compiled model's relevant fields,
// used to decide whether an existing stored authorization model already
// matches the compiled one, so bootstrap can skip writing a new version.
func modelDigest(model *openfgav1.AuthorizationModel) (string, error) {
	data, err := proto.MarshalOptions{Deterministic: true}.Marshal(&openfgav1.AuthorizationModel{
		SchemaVersion:   model.GetSchemaVersion(),
		TypeDefinitions: model.GetTypeDefinitions(),
		Conditions:      model.GetConditions(),
	})

	if err != nil {
		return "", fmt.Errorf("failed to marshal authorization model: %w", err)
	}

	sum := sha256.Sum256(data)

	return fmt.Sprintf("%x", sum[:]), nil
}
