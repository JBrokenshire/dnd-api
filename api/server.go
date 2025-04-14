package api

import (
	"dnd-api/db"
	"dnd-api/db/repositories"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"time"
)

type Server struct {
	Echo  *echo.Echo
	Db    *gorm.DB
	Repos *repositories.Repos
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

	return s
}

// Start runs anything needed for the application before starting the server
func (s *Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}
