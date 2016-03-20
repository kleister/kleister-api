package controller

import (
	"testing"

	// "github.com/Pallinder/go-randomdata"
	// "github.com/bluele/factory-go/factory"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"

	. "github.com/franela/goblin"
)

// var MinecraftFactory = factory.NewFactory(
//  &model.Minecraft{},
// ).SeqInt("ID", func(n int) (interface{}, error) {
//  return n, nil
// }).Attr("Name", func(args factory.Args) (interface{}, error) {
//  return randomdata.StringNumberExt(3, "-", 1), nil
// }).Attr("Minecraft", func(args factory.Args) (interface{}, error) {
//  return randomdata.StringNumberExt(3, "-", 1), nil
// })

func TestMinecraft(t *testing.T) {
	store := *model.Test()

	g := Goblin(t)
	g.Describe("GetMinecraft", func() {
		// var minecrafts model.Minecrafts

		// g.BeforeEach(func() {
		//  minecrafts = model.Minecrafts{
		//    MinecraftFactory.MustCreate().(*model.Minecraft),
		//    MinecraftFactory.MustCreate().(*model.Minecraft),
		//    MinecraftFactory.MustCreate().(*model.Minecraft),
		//  }
		// })

		// g.AfterEach(func() {
		// })

		g.It("should serve minecraft versions", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetMinecraft(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
			// g.Assert(rw.Body.Bytes()).Equal()
		})
	})
}
