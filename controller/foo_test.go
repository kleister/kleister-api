package controller

// import (
//   "encoding/json"
//   "testing"

//   "github.com/Pallinder/go-randomdata"
//   "github.com/bluele/factory-go/factory"
//   "github.com/gin-gonic/gin"
//   "github.com/solderapp/solder-api/model"

//   . "github.com/franela/goblin"
// )

// var ForgeFactory = factory.NewFactory(
//   &model.Forge{},
// ).SeqInt("ID", func(n int) (interface{}, error) {
//   return n, nil
// }).Attr("Name", func(args factory.Args) (interface{}, error) {
//   return randomdata.StringNumberExt(3, ".", 1), nil
// }).Attr("Minecraft", func(args factory.Args) (interface{}, error) {
//   return randomdata.StringNumberExt(3, ".", 1), nil
// })

// func TestForge(t *testing.T) {
//   gin.SetMode(gin.TestMode)
//   store := *model.Test()

//   g := Goblin(t)
//   g.Describe("GetForge", func() {
//     var forges model.Forges

//     g.BeforeEach(func() {
//       forges = model.Forges{
//         ForgeFactory.MustCreate().(*model.Forge),
//         ForgeFactory.MustCreate().(*model.Forge),
//         ForgeFactory.MustCreate().(*model.Forge),
//       }

//       for _, record := range forges {
//         store.Create(record)
//       }
//     })

//     g.AfterEach(func() {
//       store.Delete(&model.Forge{})
//     })

//     g.It("should respond with json content type", func() {
//       ctx, rw, _ := gin.CreateTestContext()
//       ctx.Set("store", store)

//       GetForge(ctx)

//       g.Assert(rw.Code).Equal(200)
//       g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
//     })

//     g.It("should serve a collection", func() {
//       ctx, rw, _ := gin.CreateTestContext()
//       ctx.Set("store", store)

//       GetForge(ctx)

//       out := model.Forges{}
//       json.NewDecoder(rw.Body).Decode(&out)

//       g.Assert(len(out)).Equal(len(forges))
//       g.Assert(out).Equal(forges)
//     })
//   })
// }
