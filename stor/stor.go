package stor

import (
	"fmt"
	"sync"

	"github.com/daregod/amazon-co-uk-scraper/scraper"
)

// BulkStor - amazon.co.uk parsed data job store interface
type BulkStor interface {
	// SaveBulk will store bulk parsed data
	SaveBulk(jobID string, bulkData []scraper.AmazonCoUkBulkData)
	// GetBulk will fetch bulk parsed data by JobID
	GetBulk(jobID string) ([]scraper.AmazonCoUkBulkData, error)
	// DeleteBulk will delete stor entry by JobID
	DeleteBulk(jobID string) error
}

// memStor in-memory BulkStor storage
type memStor struct {
	sync.RWMutex
	bulkByID map[string][]scraper.AmazonCoUkBulkData
}

// NewStor construct BulkStor storage
func NewStor() BulkStor {
	return &memStor{
		bulkByID: make(map[string][]scraper.AmazonCoUkBulkData),
	}
}

func (mem *memStor) SaveBulk(jobID string, bulkData []scraper.AmazonCoUkBulkData) {
	mem.Lock()
	defer mem.Unlock()
	mem.bulkByID[jobID] = bulkData
}

func (mem *memStor) GetBulk(jobID string) ([]scraper.AmazonCoUkBulkData, error) {
	mem.RLock()
	defer mem.RUnlock()
	if jd, ok := mem.bulkByID[jobID]; ok {
		return jd, nil
	}
	return nil, fmt.Errorf("GET. Job not found: %s", jobID)
}

func (mem *memStor) DeleteBulk(jobID string) error {
	mem.Lock()
	defer mem.Unlock()
	if _, ok := mem.bulkByID[jobID]; ok {
		delete(mem.bulkByID, jobID)
	} else {
		return fmt.Errorf("DELETE. Job not found: %s", jobID)
	}
	return nil
}
