package api

// import (
// 	"encoding/json"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/model"
// 	"github.com/kleister/kleister-api/store/data"

// 	. "github.com/franela/goblin"
// )

// func TestKey(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	store := data.Test()

// 	g := Goblin(t)
// 	g.Describe("GetKeys", func() {
// 		var keys model.Keys

// 		g.BeforeEach(func() {
// 			keys = model.Keys{
// 				&model.Key{
// 					Name:  "Key 2",
// 					Value: "KEY2",
// 				},
// 				&model.Key{
// 					Name:  "Key 1",
// 					Value: "KEY1",
// 				},
// 				&model.Key{
// 					Name:  "Key 3",
// 					Value: "KEY3",
// 				},
// 			}

// 			for _, record := range keys {
// 				store.Create(record)
// 			}
// 		})

// 		g.AfterEach(func() {
// 			store.Delete(&model.Key{})
// 		})

// 		g.It("should respond with json content type", func() {
// 			ctx, rw, _ := gin.CreateTestContext()
// 			ctx.Set("store", store)

// 			GetKeys(ctx)

// 			g.Assert(rw.Code).Equal(200)
// 			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
// 		})

// 		g.It("should serve a collection", func() {
// 			ctx, rw, _ := gin.CreateTestContext()
// 			ctx.Set("store", store)

// 			GetKeys(ctx)

// 			out := model.Keys{}
// 			json.NewDecoder(rw.Body).Decode(&out)

// 			g.Assert(len(out)).Equal(len(keys))
// 		})

// 		g.It("should sort the collection", func() {
// 			ctx, rw, _ := gin.CreateTestContext()
// 			ctx.Set("store", store)

// 			GetKeys(ctx)

// 			out := model.Keys{}
// 			json.NewDecoder(rw.Body).Decode(&out)

// 			g.Assert(out[0].Name).Equal(keys[1].Name)
// 			g.Assert(out[1].Name).Equal(keys[0].Name)
// 			g.Assert(out[2].Name).Equal(keys[2].Name)
// 		})
// 	})
// }
