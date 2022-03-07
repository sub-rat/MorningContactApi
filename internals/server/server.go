package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MorningContactApi/internals/features/user"
	"github.com/sub-rat/MorningContactApi/pkg/db/postgres"
	"gorm.io/gorm"
)

type server struct {
	C  *gin.Engine
	DB *gorm.DB
}

func GetServer() *server {
	return &server{
		C:  gin.Default(),
		DB: postgres.ConnectDatabase(),
	}
}

func (s *server) Run() {
	s.initRoutes()
	s.C.Run()
}

func (s *server) initRoutes() {
	s.C.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s.C.GET("/contacts", user.GetContacts)
	s.C.POST("/contacts", user.PostContacts)
}
