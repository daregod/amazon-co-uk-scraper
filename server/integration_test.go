package server_test

import (
	"bytes"
	"encoding/json"
	"time"

	"net/http"
	"net/http/httptest"

	"github.com/daregod/amazon-co-uk-scraper/server"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integration", func() {
	It("Works end-to-end", func() {
		By("Setup")
		srv := server.NewServer()
		gin.SetMode(gin.ReleaseMode) // Suppress stdout logging
		r := gin.New()
		srv.MountRoutes(r)

		//		log.AddHandler(console.New(true), log.AllLevels...) // More logging

		By("No jobID in request will return 404")
		w := httptest.NewRecorder()
		errReq, _ := http.NewRequest("GET", "/api/job", nil)
		r.ServeHTTP(w, errReq)
		Expect(w.Code).To(Equal(http.StatusNotFound))

		By("Wrong jobID requested will return 400")
		w = httptest.NewRecorder()
		errReq, _ = http.NewRequest("GET", "/api/job/0", nil)
		r.ServeHTTP(w, errReq)
		Expect(w.Code).To(Equal(http.StatusBadRequest))
		jobData := mustUnMarshal(w.Body.Bytes())
		Expect(*jobData.Status).To(Equal(server.StatusError))

		By("Correct POST request will enqueue scraping")
		w = httptest.NewRecorder()
		postReq, _ := http.NewRequest("POST", "/api/enqueue", bytes.NewBuffer([]byte(`["https://www.amazon.co.uk/gp/product/1509836071","http://amazon.co.uk/gp/product/1787125645"]`)))
		r.ServeHTTP(w, postReq)
		Expect(w.Code).To(Equal(http.StatusOK))
		jobData = mustUnMarshal(w.Body.Bytes())
		jobID := *jobData.ID

		By("Try to immidiate get job result will return STILL SCRAPING error, then after download content OK returned")
	LOOP:
		for {
			w = httptest.NewRecorder()
			okReq, _ := http.NewRequest("GET", "/api/job/"+jobID, nil)
			r.ServeHTTP(w, okReq)
			jobData = mustUnMarshal(w.Body.Bytes())
			switch *jobData.Status {
			case server.StatusOK:
				Expect(w.Code).To(Equal(http.StatusOK))
				break LOOP
			case server.StatusError: // STILL SCRAPING
				Expect(w.Code).To(Equal(http.StatusBadRequest))
				time.Sleep(500 * time.Millisecond)
				continue
			}
		}
		Expect(len(jobData.Data)).To(Equal(2))

		By("Correct DELETE request will return DELETED OK")
		w = httptest.NewRecorder()
		delReq, _ := http.NewRequest("DELETE", "/api/deletejob/"+jobID, nil)
		r.ServeHTTP(w, delReq)
		Expect(w.Code).To(Equal(http.StatusOK))
		jobData = mustUnMarshal(w.Body.Bytes())
		Expect(*jobData.Status).To(Equal(server.StatusDeleted))

		By("Already deleted jobID request will return 400")
		w = httptest.NewRecorder()
		errReq, _ = http.NewRequest("GET", "/api/job/"+jobID, nil)
		r.ServeHTTP(w, errReq)
		Expect(w.Code).To(Equal(http.StatusBadRequest))
		jobData = mustUnMarshal(w.Body.Bytes())
		Expect(*jobData.Status).To(Equal(server.StatusError))
	})
})

func mustUnMarshal(data []byte) server.JobData {
	jobData := server.JobData{}
	json.Unmarshal(data, &jobData)
	return jobData
}
