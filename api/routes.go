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

	r.engine.Static("/foo", "../web/dist")
	r.engine.StaticFile("/index.b18976e1.css", "../web/dist/index.b18976e1.css")
	r.engine.StaticFile("/index.b38dac7d.js", "../web/dist/index.b38dac7d.js")

	repo := NewRepository(db)
	handler := NewSeeder(repo)

	r.engine.GET("", handler.SayHi())
	group := r.engine.Group("/seeder")
	group.GET("/drop-tables", handler.Drop())
	group.GET("/create-tables", handler.Create())
	return r
}
