package controller

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"

	. "github.com/franela/goblin"
)

func TestKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := *model.Test()

	g := Goblin(t)
	g.Describe("GetKeys", func() {
		var keys model.Keys

		g.BeforeEach(func() {
			keys = model.Keys{
				&model.Key{
					Name: "Key 2",
				},
				&model.Key{
					Name: "Key 1",
				},
				&model.Key{
					Name: "Key 3",
				},
			}

			for _, record := range keys {
				store.Create(record)
			}
		})

		g.AfterEach(func() {
			store.Delete(&model.Key{})
		})

		g.It("should respond with json content type", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetKey(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
		})

		g.It("should serve a collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetKey(ctx)

			out := model.Keys{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(len(out)).Equal(len(keys))
			g.Assert(out[0]).Equal(keys[2])
		})
	})
}
