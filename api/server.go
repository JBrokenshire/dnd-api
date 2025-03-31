package api

import (
	"dnd-api/db"
	"dnd-api/db/repositories"
	"dnd-api/services"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
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

	s.Echo.HideBanner = true

	return s
}

func (s *Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}

func (s *Server) SetupPermissions() {
	log.Println("Setting up permissions in the database")
	permService := services.NewPermissionService(s.Db)
	permService.SyncPermissions()
}
