package controller

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"

	. "github.com/franela/goblin"
)

func TestForge(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := *model.Test()

	g := Goblin(t)
	g.Describe("GetForge", func() {
		var forges model.Forges

		g.BeforeEach(func() {
			forges = model.Forges{
				&model.Forge{
					Name:      "10.13.4.1614",
					Minecraft: "1.7.10",
				},
				&model.Forge{
					Name:      "11.15.1.1765",
					Minecraft: "1.8.9",
				},
				&model.Forge{
					Name:      "10.12.2.1147",
					Minecraft: "1.7.2",
				},
			}

			for _, record := range forges {
				store.Create(record)
			}
		})

		g.AfterEach(func() {
			store.Delete(&model.Forge{})
		})

		g.It("should respond with json content type", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetForge(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
		})

		g.It("should serve a collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetForge(ctx)

			out := model.Forges{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(len(out)).Equal(len(forges))
		})

		g.It("should sort the collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetForge(ctx)

			out := model.Forges{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(out[0].Name).Equal(forges[1].Name)
			g.Assert(out[1].Name).Equal(forges[2].Name)
			g.Assert(out[2].Name).Equal(forges[0].Name)
		})
	})
}
