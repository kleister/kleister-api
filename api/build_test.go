package api

// import (
// 	"encoding/json"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/model"
// 	"github.com/kleister/kleister-api/store/data"

// 	. "github.com/franela/goblin"
// )

// func TestBuild(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	store := data.Test()

// 	g := Goblin(t)
// 	g.Describe("GetBuilds", func() {
// 		var pack model.Pack
// 		var builds model.Builds

// 		g.BeforeEach(func() {
// 			pack = model.Pack{
// 				Name: "Pack",
// 			}

// 			store.Create(&pack)

// 			builds = model.Builds{
// 				&model.Build{
// 					PackID: pack.ID,
// 					Name:   "Build 1",
// 				},
// 				&model.Build{
// 					PackID: pack.ID,
// 					Name:   "Build 3",
// 				},
// 				&model.Build{
// 					PackID: pack.ID,
// 					Name:   "Build 2",
// 				},
// 			}

// 			for _, record := range builds {
// 				store.Create(record)
// 			}
// 		})

// 		g.AfterEach(func() {
// 			store.Delete(&model.Build{})
// 			store.Delete(&model.Pack{})
// 		})

// 		g.It("should respond with json content type", func() {
// 			ctx, rw, _ := gin.CreateTestContext()
// 			ctx.Set("store", store)
// 			ctx.Set("pack", &pack)

// 			GetBuilds(ctx)

// 			g.Assert(rw.Code).Equal(200)
// 			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
// 		})

// 		g.It("should serve a collection", func() {
// 			ctx, rw, _ := gin.CreateTestContext()
// 			ctx.Set("store", store)
// 			ctx.Set("pack", &pack)

// 			GetBuilds(ctx)

// 			out := model.Builds{}
// 			json.NewDecoder(rw.Body).Decode(&out)

// 			g.Assert(len(out)).Equal(len(builds))
// 		})

// 		g.It("should sort the collection", func() {
// 			ctx, rw, _ := gin.CreateTestContext()
// 			ctx.Set("store", store)
// 			ctx.Set("pack", &pack)

// 			GetBuilds(ctx)

// 			out := model.Builds{}
// 			json.NewDecoder(rw.Body).Decode(&out)

// 			g.Assert(out[0].Name).Equal(builds[0].Name)
// 			g.Assert(out[1].Name).Equal(builds[2].Name)
// 			g.Assert(out[2].Name).Equal(builds[1].Name)
// 		})
// 	})
// }
