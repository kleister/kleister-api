package controller

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"

	. "github.com/franela/goblin"
)

func TestUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := *model.Test()

	g := Goblin(t)
	g.Describe("GetUsers", func() {
		var users model.Users

		g.BeforeEach(func() {
			users = model.Users{
				&model.User{
					Username: "thomas",
				},
				&model.User{
					Username: "brad",
				},
				&model.User{
					Username: "felix",
				},
			}

			for _, record := range users {
				store.Create(record)
			}
		})

		g.AfterEach(func() {
			store.Delete(&model.User{})
		})

		g.It("should respond with json content type", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetUser(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
		})

		g.It("should serve a collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetUser(ctx)

			out := model.Users{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(len(out)).Equal(len(users))
			g.Assert(out[0]).Equal(users[2])
		})
	})
}
