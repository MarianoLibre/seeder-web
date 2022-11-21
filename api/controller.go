package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Seeder struct {
	repository Repository
}

func NewSeeder(r Repository) *Seeder {
	return &Seeder{
		repository: r,
	}
}


func (s *Seeder) SayHi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html" , nil)
	}
}

func (s *Seeder) Drop() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := s.repository.DropTables()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "All tables dropped")
	}
}

func (s *Seeder) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := s.repository.CreateTables()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return 
		}
		ctx.JSON(http.StatusOK, "All tables created")
		return 
	}
}
