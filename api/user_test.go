package api

// import (
// 	"encoding/json"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/model"
// 	"github.com/kleister/kleister-api/store/data"

// 	. "github.com/franela/goblin"
// )

// func TestUser(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	store := data.Test()

// 	g := Goblin(t)
// 	g.Describe("GetUsers", func() {
// 		var users model.Users

// 		g.BeforeEach(func() {
// 			users = model.Users{
// 				&model.User{
// 					Username: "thomas",
// 					Email:    "thomas@webhippie.de",
// 					Password: "test12345",
// 				},
// 				&model.User{
// 					Username: "brad",
// 					Email:    "brad@webhippie.de",
// 					Password: "test12345",
// 				},
// 				&model.User{
// 					Username: "felix",
// 					Email:    "felix@webhippie.de",
// 					Password: "test12345",
// 				},
// 			}

// 			for _, record := range users {
// 				store.Create(record)
// 			}
// 		})

// 		g.AfterEach(func() {
// 			store.Delete(&model.User{})
// 		})

// 		g.It("should respond with json content type", func() {
// 			ctx, rw, _ := gin.CreateTestContext()
// 			ctx.Set("store", store)

// 			GetUsers(ctx)

// 			g.Assert(rw.Code).Equal(200)
// 			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
// 		})

// 		g.It("should serve a collection", func() {
// 			ctx, rw, _ := gin.CreateTestContext()
// 			ctx.Set("store", store)

// 			GetUsers(ctx)

// 			out := model.Users{}
// 			json.NewDecoder(rw.Body).Decode(&out)

// 			g.Assert(len(out)).Equal(len(users))
// 		})

// 		g.It("should sort the collection", func() {
// 			ctx, rw, _ := gin.CreateTestContext()
// 			ctx.Set("store", store)

// 			GetUsers(ctx)

// 			out := model.Users{}
// 			json.NewDecoder(rw.Body).Decode(&out)

// 			g.Assert(out[0].Username).Equal(users[1].Username)
// 			g.Assert(out[1].Username).Equal(users[2].Username)
// 			g.Assert(out[2].Username).Equal(users[0].Username)
// 		})
// 	})
// }
