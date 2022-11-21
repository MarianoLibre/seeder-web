package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type router struct {
	engine *gin.Engine
	database *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) router {
	r := router{engine: eng, database: db}

	r.engine.LoadHTMLGlob("index.html")
	repo := NewRepository(db)
	handler := NewSeeder(repo)
	r.engine.GET("", handler.SayHi())

	group := r.engine.Group("/seeder")
	group.GET("/drop-tables", handler.Drop())
	group.GET("/create-tables", handler.Create())
	return r
}
