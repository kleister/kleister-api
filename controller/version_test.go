package controller

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
		var versions model.Versions

		g.BeforeEach(func() {
			versions = model.Versions{
				&model.Version{
					Name: "Version 3",
				},
				&model.Version{
					Name: "Version 1",
				},
				&model.Version{
					Name: "Version 2",
				},
			}

			for _, record := range versions {
				store.Create(record)
			}
		})

		g.AfterEach(func() {
			store.Delete(&model.Version{})
		})

		g.It("should respond with json content type", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetVersion(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
		})

		g.It("should serve a collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetVersion(ctx)

			out := model.Versions{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(len(out)).Equal(len(versions))
			g.Assert(out[0]).Equal(versions[2])
		})
	})
}
