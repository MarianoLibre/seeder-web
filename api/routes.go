package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

type router struct {
	engine   *gin.Engine
	database *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) router {
	r := router{engine: eng, database: db}

	// Serve the front
	r.engine.Static("/foo", "../web/dist")

	files, err := ioutil.ReadDir("../web/dist")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		fname := f.Name()
		re := regexp.MustCompile(`.*\.css$|.*\.js$`)
		if re.MatchString(fname) {
			r.engine.StaticFile(
				fmt.Sprintf("/%s", fname),
				fmt.Sprintf("../web/dist/%s", fname),
			)
			msg := fmt.Sprintf("[ File: '%s' loaded ]\n", f.Name())
			color.New(color.BgBlue).Add(color.FgBlack).Print(msg)
		}
	}

	// Connect to the db
	repo := NewRepository(db)
	handler := NewSeeder(repo)

	// Dummy "it's running"
	r.engine.LoadHTMLGlob("index.html")
	r.engine.GET("", handler.SayHi())

	// Routes
	group := r.engine.Group("/seeder")
	group.GET("/drop-tables", handler.Drop())
	group.GET("/create-tables", handler.Create())

	return r
}
