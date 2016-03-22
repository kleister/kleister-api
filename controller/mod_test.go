package controller

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"

	. "github.com/franela/goblin"
)

func TestMod(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := *model.Test()

	g := Goblin(t)
	g.Describe("GetMods", func() {
		var mods model.Mods

		g.BeforeEach(func() {
			mods = model.Mods{
				&model.Mod{
					Name: "Mod 1",
				},
				&model.Mod{
					Name: "Mod 2",
				},
				&model.Mod{
					Name: "Mod 3",
				},
			}

			for _, record := range mods {
				store.Create(record)
			}
		})

		g.AfterEach(func() {
			store.Delete(&model.Mod{})
		})

		g.It("should respond with json content type", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetMods(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
		})

		g.It("should serve a collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetMods(ctx)

			out := model.Mods{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(len(out)).Equal(len(mods))
		})

		g.It("should sort the collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetMods(ctx)

			out := model.Mods{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(out[0].Name).Equal(mods[0].Name)
			g.Assert(out[1].Name).Equal(mods[1].Name)
			g.Assert(out[2].Name).Equal(mods[2].Name)
		})
	})
}
