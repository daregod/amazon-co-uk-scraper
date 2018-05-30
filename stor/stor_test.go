package stor_test

import (
	"testing"

	"github.com/daregod/amazon-co-uk-scraper/scraper"

	"github.com/daregod/amazon-co-uk-scraper/stor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "stor")
}

var _ = Describe("Stor", func() {
	var st stor.BulkStor
	BeforeEach(func() {
		st = stor.NewStor()
		Expect(st).ToNot(BeNil())
	})
	It("Compiles", func() {
		Expect(true).To(BeTrue())
	})
	It("Stor works", func() {
		errSt := "Parse trouble fields: Price"
		data1 := []scraper.AmazonCoUkBulkData{
			{
				URL:   "https://www.amazon.co.uk/gp/product/059652692X",
				Error: &errSt,
			},
		}
		data2 := []scraper.AmazonCoUkBulkData{
			{
				URL:  "https://www.amazon.co.uk/gp/product/1509836071",
				Meta: &scraper.AmazonCoUkParsedData{Title: "The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts", Price: "8.49", Image: "https://images-eu.ssl-images-amazon.com/images/I/61modEZimPL._SX218_BO1,204,203,200_QL40_.jpg", Available: true},
			},
			{
				URL:  "https://www.amazon.co.uk/gp/product/1787125645",
				Meta: &scraper.AmazonCoUkParsedData{Title: "Go Systems Programming: Master Linux and Unix system level programming with Go", Price: "41.99", Image: "https://images-eu.ssl-images-amazon.com/images/I/41y7-qWywtL._SX218_BO1,204,203,200_QL40_.jpg", Available: true},
			},
		}

		job1Id := "12345"
		job2Id := "11111"
		By("Save first")
		st.SaveBulk(job1Id, data1)
		job, err := st.GetBulk(job1Id)
		Expect(err).To(Succeed())
		Expect(job).To(ConsistOf(data1))
		job, err = st.GetBulk(job2Id)
		Expect(err).ToNot(Succeed())

		By("Save second")
		st.SaveBulk(job2Id, data2)
		job, err = st.GetBulk(job2Id)
		Expect(err).To(Succeed())
		Expect(job).To(Equal(data2))
		job, err = st.GetBulk(job1Id)
		Expect(err).To(Succeed())
		Expect(job).To(Equal(data1))

		By("Delete first")
		st.DeleteBulk(job1Id)
		job, err = st.GetBulk(job1Id)
		Expect(err).ToNot(Succeed())
		job, err = st.GetBulk(job2Id)
		Expect(err).To(Succeed())
		Expect(job).To(Equal(data2))

		By("Delete second")
		st.DeleteBulk(job2Id)
		job, err = st.GetBulk(job1Id)
		Expect(err).ToNot(Succeed())
		job, err = st.GetBulk(job2Id)
		Expect(err).ToNot(Succeed())
	})
})
