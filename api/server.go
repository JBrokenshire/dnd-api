package api

import (
	"dnd-api/db"
	"dnd-api/db/repositories"
	"dnd-api/pkg/dependencies"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"time"
)

type Server struct {
	Echo         *echo.Echo
	Db           *gorm.DB
	Repos        *repositories.Repos
	Dependencies *dependencies.DependencyService
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

	s.Echo.HideBanner = true

	s.Repos = repositories.NewRepos(s.Db)
	s.Dependencies = dependencies.NewDependencyService(s.Db)

	return s
}

// Start runs anything needed for the application before starting the server
func (s *Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}
