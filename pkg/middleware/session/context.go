package session

import (
	"fmt"
)

type contextKey struct {
	name string
}

func (c *contextKey) String() string {
	return fmt.Sprintf("middleware-%s", c.name)
}
