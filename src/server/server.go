package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/daregod/amazon-co-uk-scraper/src/scraper"
	"github.com/daregod/amazon-co-uk-scraper/src/stor"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/log"
	"github.com/twinj/uuid"
)

// JobData is a Output API struct
type JobData struct {
	ID     *string                      `json:"id,omitempty"`
	Data   []scraper.AmazonCoUkBulkData `json:"data,omitempty"`
	Status *string                      `json:"status,omitempty"`
	Error  *string                      `json:"error,omitempty"`
}

var (
	// StatusOK = OK
	StatusOK = "OK"
	// StatusDeleted = Deleted OK
	StatusDeleted = "DELETED OK"
	// StatusProceed = Proceed OK
	StatusProceed = "PROCEED OK"
	// StatusError = Error
	StatusError = "Error"
)

// Server is base server implementation
type Server struct {
	// Storage for amazon.co.uk bulk parsed data
	Storage stor.BulkStor
	// Mutex cover proceed cache
	sync.RWMutex
	// cache store processed queue
	proceed map[string]struct{}
}

// NewServer constructor
func NewServer() *Server {
	return &Server{
		Storage: stor.NewStor(),
		proceed: make(map[string]struct{}),
	}
}

// MountRoutes assign route handlers
func (srv *Server) MountRoutes(r gin.IRouter) {
	r.POST("/api/enqueue", srv.routePUT)
	r.GET("/api/job/:id", srv.routeGET)
	r.DELETE("/api/deletejob/:id", srv.routeDELETE)
}

// routePUT is a PUT handler
func (srv *Server) routePUT(c *gin.Context) {
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.WithError(err).WithFields(log.F("RequestURI", c.Request.RequestURI)).Alert("CANNOT READ REQUEST BODY")
		sendError(c, "", err)
		return
	}

	var urls = []string{}
	err = json.Unmarshal(reqBody, &urls)
	if err != nil {
		log.WithError(err).WithFields(log.F("Request Body", string(reqBody))).Alert("CANNOT UNMARSHAL REQUEST BODY")
		sendError(c, "", err)
		return
	}
	id := genGUID()
	srv.addCacheEntry(id)
	log.WithFields(log.F("ID", id), log.F("urls", urls)).Info("ENQUEUE")
	go func() {
		pd := scraper.NewProcessor().ProcessUrls(urls)
		srv.Storage.SaveBulk(id, pd)
	}()
	sendResponse(c, JobData{ID: &id, Status: &StatusProceed})
}

// routeGET is a GET handler
func (srv *Server) routeGET(c *gin.Context) {
	id := c.Param("id")
	bulk, err := srv.Storage.GetBulk(id)
	if err != nil {
		if srv.isInProgres(id) {
			log.WithFields(log.F("id", id)).Info("STILL SCRAPING")
			sendError(c, id, fmt.Errorf("STILL SCRAPING. Come back later"))
			return
		}
		log.WithError(err).WithFields(log.F("id", id)).Alert("NO STORAGE RECORD")
		sendError(c, id, err)
		return
	}
	srv.dropCacheEntry(id)
	log.WithFields(log.F("ID", id), log.F("bulk size", len(bulk))).Info("GET JOB")
	sendResponse(c, JobData{ID: &id, Status: &StatusOK, Data: bulk})
}

// routeDELETE is a DELETE handler
func (srv *Server) routeDELETE(c *gin.Context) {
	id := c.Param("id")
	srv.dropCacheEntry(id)
	err := srv.Storage.DeleteBulk(id)
	if err != nil {
		log.WithFields(log.F("id", id)).WithError(err).Alert("NO STORAGE RECORD")
		sendError(c, id, err)
		return
	}
	log.WithFields(log.F("ID", id)).Info("DELETE JOB")
	sendResponse(c, JobData{ID: &id, Status: &StatusDeleted})
}

func (srv *Server) addCacheEntry(id string) {
	srv.Lock()
	defer srv.Unlock()
	srv.proceed[id] = struct{}{}
}

func (srv *Server) dropCacheEntry(id string) {
	srv.Lock()
	defer srv.Unlock()
	if _, ok := srv.proceed[id]; ok {
		delete(srv.proceed, id)
	}
}

func (srv *Server) isInProgres(id string) bool {
	srv.RLock()
	defer srv.RUnlock()
	_, ok := srv.proceed[id]
	return ok
}

// sendResponse will send normal message to client
func sendResponse(c *gin.Context, data interface{}) {
	c.Data(http.StatusOK, "application/json", mustMarshall(data))
}

// sendError will send error message to client
func sendError(c *gin.Context, id string, err error) {
	output := JobData{Status: &StatusError}
	if id != "" {
		output.ID = &id
	}
	if err != nil {
		errSt := err.Error()
		output.Error = &errSt
	}
	c.Data(http.StatusBadRequest, "application/json", mustMarshall(output))
}

// mustMarshall suppress error, because marshalled only fixed structs
func mustMarshall(data interface{}) []byte {
	marshalled, _ := json.Marshal(data)
	return marshalled
}

// genGUID will generate unique GUID string
func genGUID() string {
	return uuid.NewV4().String()
}
