package api

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"

	. "github.com/franela/goblin"
)

func TestVersion(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := *model.Test()

	g := Goblin(t)
	g.Describe("GetVersions", func() {
		var mod model.Mod
		var versions model.Versions

		g.BeforeEach(func() {
			mod = model.Mod{
				Name: "Mod",
			}

			store.Create(&mod)

			versions = model.Versions{
				&model.Version{
					ModID: mod.ID,
					Name:  "Version 3",
				},
				&model.Version{
					ModID: mod.ID,
					Name:  "Version 1",
				},
				&model.Version{
					ModID: mod.ID,
					Name:  "Version 2",
				},
			}

			for _, record := range versions {
				store.Create(record)
			}
		})

		g.AfterEach(func() {
			store.Delete(&model.Version{})
			store.Delete(&model.Mod{})
		})

		g.It("should respond with json content type", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)
			ctx.Set("mod", &mod)

			GetVersions(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
		})

		g.It("should serve a collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)
			ctx.Set("mod", &mod)

			GetVersions(ctx)

			out := model.Versions{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(len(out)).Equal(len(versions))
		})

		g.It("should sort the collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)
			ctx.Set("mod", &mod)

			GetVersions(ctx)

			out := model.Versions{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(out[0].Name).Equal(versions[1].Name)
			g.Assert(out[1].Name).Equal(versions[2].Name)
			g.Assert(out[2].Name).Equal(versions[0].Name)
		})
	})
}
