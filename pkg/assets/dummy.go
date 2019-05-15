package assets

import (
	// Fake the import to make the dep tree happy.
	_ "golang.org/x/net/context"

	// Fake the import to make the dep tree happy.
	_ "golang.org/x/net/webdav"
)

//go:generate gorunpkg github.com/UnnoTed/fileb0x ab0x.yaml
