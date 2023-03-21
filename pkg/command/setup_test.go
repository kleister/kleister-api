package command

import (
	"testing"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestSetupLogger(t *testing.T) {
	assert.NoError(t, setupLogger(config.Load()))
}
