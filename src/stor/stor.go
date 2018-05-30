package stor

import (
	"fmt"
	"sync"

	"github.com/daregod/amazon-co-uk-scraper/src/scraper"
)

// amazon.co.uk parsed data job store interface
type BulkStor interface {
	// SaveBulk will store bulk parsed data
	SaveBulk(jobID string, bulkData []scraper.AmazonCoUkBulkData)
	// GetBulk will fetch bulk parsed data by JobID
	GetBulk(jobID string) ([]scraper.AmazonCoUkBulkData, error)
	// DeleteBulk will delete stor entry by JobID
	DeleteBulk(jobID string) error
}

type memStor struct {
	sync.RWMutex
	bulkById map[string][]scraper.AmazonCoUkBulkData
}

func NewStor() BulkStor {
	return &memStor{
		bulkById: make(map[string][]scraper.AmazonCoUkBulkData),
	}
}

func (mem *memStor) SaveBulk(jobID string, bulkData []scraper.AmazonCoUkBulkData) {
	mem.Lock()
	defer mem.Unlock()
	mem.bulkById[jobID] = bulkData
}

func (mem *memStor) GetBulk(jobID string) ([]scraper.AmazonCoUkBulkData, error) {
	mem.RLock()
	defer mem.RUnlock()
	if jd, ok := mem.bulkById[jobID]; ok {
		return jd, nil
	}
	return nil, fmt.Errorf("GET. Job not found: %s", jobID)
}

func (mem *memStor) DeleteBulk(jobID string) error {
	mem.Lock()
	defer mem.Unlock()
	if _, ok := mem.bulkById[jobID]; ok {
		delete(mem.bulkById, jobID)
	} else {
		return fmt.Errorf("DELETE. Job not found: %s", jobID)
	}
	return nil
}
