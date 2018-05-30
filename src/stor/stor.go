package stor

import (
	"fmt"
	"sync"

	"github.com/daregod/amazon-co-uk-scraper/src/scraper"
	"github.com/twinj/uuid"
)

// jobData collect bulk parsed data with job ID
type jobData struct {
	// GUID
	id string
	// Parsed data
	data []scraper.AmazonCoUkBulkData
}

// amazon.co.uk parsed data job store interface
type BulkStor interface {
	// SaveBulk will store bulk parsed data
	SaveBulk(bulkData []scraper.AmazonCoUkBulkData) (jobID string)
	// GetBulk will fetch bulk parsed data by JobID
	GetBulk(jobID string) ([]scraper.AmazonCoUkBulkData, error)
	// DeleteBulk will delete stor entry by JobID
	DeleteBulk(jobID string) error
}

type memStor struct {
	sync.RWMutex
	jobsById map[string]jobData
}

func NewStor() BulkStor {
	return &memStor{
		jobsById: make(map[string]jobData),
	}
}

func (mem *memStor) SaveBulk(bulkData []scraper.AmazonCoUkBulkData) (jobID string) {
	job := jobData{
		id:   genGUID(),
		data: bulkData,
	}
	mem.Lock()
	defer mem.Unlock()
	mem.jobsById[job.id] = job
	return job.id
}

func (mem *memStor) GetBulk(jobID string) ([]scraper.AmazonCoUkBulkData, error) {
	mem.RLock()
	defer mem.RUnlock()
	if jd, ok := mem.jobsById[jobID]; ok {
		return jd.data, nil
	}
	return nil, fmt.Errorf("GET. Job not found: %s", jobID)
}

func (mem *memStor) DeleteBulk(jobID string) error {
	mem.Lock()
	defer mem.Unlock()
	if _, ok := mem.jobsById[jobID]; ok {
		delete(mem.jobsById, jobID)
	} else {
		return fmt.Errorf("DELETE. Job not found: %s", jobID)
	}
	return nil
}

// genGUID will generate unique GUID string
func genGUID() string {
	return uuid.NewV4().String()
}
