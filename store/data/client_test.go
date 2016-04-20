package data

import (
	"strconv"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/solderapp/solder-api/model"
)

func TestClients(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	store := Test()

	g.Describe("Clients", func() {
		g.BeforeEach(func() {
			store.DeleteClient(&model.Client{})
		})

		g.Describe("GetClient", func() {
			g.It("Should find a client", func() {
				record := model.Client{
					Name:  "Client 1",
					Value: "UUID1",
				}

				store.CreateClient(
					&record,
				)

				result, res := store.GetClient(
					strconv.Itoa(record.ID),
				)

				Expect(res.RecordNotFound()).To(BeFalse())
        Expect(res.Error).NotTo(HaveOccurred())
				Expect(record.Slug).To(Equal("client-1"))
				Expect(record.ID).To(Equal(result.ID))
				Expect(record.Name).To(Equal(result.Name))
				Expect(record.Value).To(Equal(result.Value))
			})

			g.It("Should return RecordNotFound", func() {
				result, res := store.GetClient(
					"foo",
				)

				Expect(res.RecordNotFound()).To(BeTrue())
        Expect(res.Error).To(HaveOccurred())
				Expect(result.ID).To(Equal(0))
			})
		})

		g.Describe("GetClients", func() {
			g.It("Should find clients", func() {

			})
		})

		g.Describe("CreateClient", func() {

		})

		g.Describe("UpdateClient", func() {

		})

		g.Describe("DeleteClient", func() {
			g.It("Should delete a client", func() {
				record := model.Client{
					Name:  "Client 1",
					Value: "UUID1",
				}

				store.CreateClient(
					&record,
				)

				err := store.DeleteClient(
					&record,
				)

				_, res := store.GetClient(
					strconv.Itoa(record.ID),
				)

				Expect(err).NotTo(HaveOccurred())
				Expect(res.RecordNotFound()).To(BeTrue())
			})

			g.It("Should not result in an error", func() {
				err := store.DeleteClient(
					&model.Client{
						ID: 1000,
					},
				)

				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
}
