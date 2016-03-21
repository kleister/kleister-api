package controller

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"

	. "github.com/franela/goblin"
)

func TestBuild(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := *model.Test()

	g := Goblin(t)
	g.Describe("GetBuilds", func() {
		var builds model.Builds

		g.BeforeEach(func() {
			builds = model.Builds{
				&model.Build{
					Name: "Build 1",
				},
				&model.Build{
					Name: "Build 3",
				},
				&model.Build{
					Name: "Build 2",
				},
			}

			for _, record := range builds {
				store.Create(record)
			}
		})

		g.AfterEach(func() {
			store.Delete(&model.Build{})
		})

		g.It("should respond with json content type", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetBuild(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
		})

		g.It("should serve a collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetBuild(ctx)

			out := model.Builds{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(len(out)).Equal(len(builds))
			g.Assert(out[0]).Equal(builds[2])
		})
	})
}
