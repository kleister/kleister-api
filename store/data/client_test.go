package data

import (
  "testing"
  "strconv"

  "github.com/solderapp/solder-api/model"
  . "github.com/franela/goblin"
)

func TestClients(t *testing.T) {
  store := Test()

  g := Goblin(t)
  g.Describe("Clients", func() {
    g.BeforeEach(func() {
      store.DeleteClient(&model.Client{})
    })

    g.It("Should get a Client", func() {
      record := model.Client{
        Name: "Client 1",
        Value: "UUID1",
      }

      store.CreateClient(
        &record,
      )

      result, res := store.GetClient(
        strconv.Itoa(record.ID),
      )

      // TODO(should): Check for existance of Slug
      g.Assert(res.Error == nil).IsTrue()
      g.Assert(record.ID).Equal(result.ID)
      g.Assert(record.Name).Equal(result.Name)
      g.Assert(record.Value).Equal(result.Value)
    })
  })
}
