package identicon

import (
	"bytes"

	"github.com/rrivera/identicon"
)

// New creates a new identifcon image.
func New(ns, name string) (*bytes.Buffer, error) {
	icon, err := identicon.New(
		ns,
		10,
		7,
	)

	if err != nil {
		return nil, err
	}

	draw, err := icon.Draw(
		name,
	)

	if err != nil {
		return nil, err
	}

	result := bytes.NewBuffer(
		[]byte{},
	)

	if err := draw.Png(
		64,
		result,
	); err != nil {
		return nil, err
	}

	return result, nil
}
