package session

import (
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
)

const (
	// CurrentContextKey defines the context key that stores the user.
	CurrentContextKey = "client"
)

// Current gets the user from the context.
func Current(c *gin.Context) *model.User {
	v, ok := c.Get(CurrentContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.User)

	if !ok {
		return nil
	}

	return r
}

// SetCurrent injects the user into the context.
func SetCurrent() gin.HandlerFunc {
	return func(c *gin.Context) {
		record := &model.User{
			Username: "static",
			Email:    "solder@webhippie.de",
			Permission: &model.Permission{
				DisplayUsers:   true,
				ChangeUsers:    true,
				DeleteUsers:    true,
				DisplayKeys:    true,
				ChangeKeys:     true,
				DeleteKeys:     true,
				DisplayClients: true,
				ChangeClients:  true,
				DeleteClients:  true,
				DisplayPacks:   true,
				ChangePacks:    true,
				DeletePacks:    true,
				DisplayMods:    true,
				ChangeMods:     true,
				DeleteMods:     true,
			},
		}

		c.Set(CurrentContextKey, record)
		c.Next()
	}
}
