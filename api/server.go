package api

import (
	"github.com/JBrokenshire/dnd-api/db"
	"github.com/JBrokenshire/dnd-api/db/repositories"
	"github.com/JBrokenshire/dnd-api/pkg/dependencies"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

type Server struct {
	Echo         *echo.Echo
	Db           *gorm.DB
	Dependencies *dependencies.DependencyService
	Repos        *repositories.Repos
}

func NewServer() *Server {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = utc

	s := &Server{
		Echo: echo.New(),
		Db:   db.Init(),
	}

	s.Repos = repositories.NewRepos(s.Db)
	s.Dependencies = dependencies.NewDependencyService(s.Db)

	// Make sure production dependencies can start. Otherwise, we don't start the service
	if os.Getenv("ENVIRONMENT") == "production" {
		s.Dependencies.PreWarmServices()
	}

	s.Echo.HideBanner = true

	return s
}

// Start runs anything needed for the application before starting the server
func (s *Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}
