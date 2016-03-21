package controller

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"

	. "github.com/franela/goblin"
)

func TestClient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := *model.Test()

	g := Goblin(t)
	g.Describe("GetClients", func() {
		var clients model.Clients

		g.BeforeEach(func() {
			clients = model.Clients{
				&model.Client{
					Name: "Client 3",
				},
				&model.Client{
					Name: "Client 1",
				},
				&model.Client{
					Name: "Client 2",
				},
			}

			for _, record := range clients {
				store.Create(record)
			}
		})

		g.AfterEach(func() {
			store.Delete(&model.Client{})
		})

		g.It("should respond with json content type", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetClient(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
		})

		g.It("should serve a collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetClient(ctx)

			out := model.Clients{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(len(out)).Equal(len(clients))
			g.Assert(out[0]).Equal(clients[2])
		})
	})
}
