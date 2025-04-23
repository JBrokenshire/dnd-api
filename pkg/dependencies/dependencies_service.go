package dependencies

import (
	"dnd-api/pkg/file_storage"
	"dnd-api/services/file_service"
	"github.com/jinzhu/gorm"
)

type DependencyService struct {
	db *gorm.DB

	fileStore   file_storage.Store
	fileService file_service.Service
}

func NewDependencyService(db *gorm.DB) *DependencyService {
	ds := &DependencyService{db: db}
	return ds
}

func (s *DependencyService) GetDB() *gorm.DB {
	return s.db
}

func (s *DependencyService) GetFileStore() file_storage.Store {
	if s.fileStore == nil {
		s.fileStore = file_storage.NewLocalStorage("assets/files/")
	}
	return s.fileStore
}

func (s *DependencyService) SetFileStore(store file_storage.Store) {
	s.fileStore = store
}

func (s *DependencyService) GetFileService() file_service.Service {
	if s.fileService == nil {
		s.fileService = file_service.NewFileService(s.GetDB())
	}
	return s.fileService
}

func (s *DependencyService) SetFileService(service file_service.Service) {
	s.fileService = service
}
