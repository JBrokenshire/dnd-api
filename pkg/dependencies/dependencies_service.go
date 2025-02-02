package dependencies

import (
	"github.com/JBrokenshire/dnd-api/pkg/file_storage"
	"github.com/jinzhu/gorm"
	"sync"
)

type DependencyService struct {
	db *gorm.DB
	wg *sync.WaitGroup

	fileStore file_storage.Store
}

func NewDependencyService(db *gorm.DB) *DependencyService {
	ds := &DependencyService{db: db}
	return ds
}

// PreWarmServices will call services when the server starts so it's ready, and not lazy loaded when required. This is
// helpful for any services which may panic, or services which are essential in production. For testing and development
// this is now called.
func (s *DependencyService) PreWarmServices() {
	s.GetFileStore()
}

func (s *DependencyService) GetDB() *gorm.DB {
	return s.db
}

func (s *DependencyService) CreateWg() {
	s.wg = &sync.WaitGroup{}
}

// GetWg returns a
func (s *DependencyService) GetWg() *sync.WaitGroup {
	return s.wg
}

// WgAdd adds 1 to the wait group if hte wait group exists. This is mainly for waiting whilst running tests
func (s *DependencyService) WgAdd() {
	if s.wg != nil {
		s.wg.Add(1)
	}
}

// WgDone calls done on the server WG if it exists.
func (s *DependencyService) WgDone() {
	if s.wg != nil {
		s.wg.Done()
	}
}

func (s *DependencyService) SetFileStore(store file_storage.Store) {
	s.fileStore = store
}

func (s *DependencyService) GetFileStore() file_storage.Store {
	if s.fileStore == nil {
		s.fileStore = file_storage.NewLocalStorage("assets/files/")
	}
	return s.fileStore
}
