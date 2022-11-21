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
	seeder := r.engine.Group("/seeder")
	seeder.POST("/drop-tables", handler.Drop())
	seeder.POST("/create-tables", handler.Create())
	seeder.POST("/warehouses", handler.SeedWarehouses())
	seeder.POST("/buyers", handler.SeedBuyers())
	seeder.POST("/sellers", handler.SeedSellers())
	seeder.POST("/employees", handler.SeedEmployees())
	seeder.POST("/products", handler.SeedProducts())
	seeder.POST("/sections", handler.SeedSections())
	seeder.POST("/localities", handler.SeedLocalities())
	seeder.POST("/carries", handler.SeedCarries())
	seeder.POST("/product-records", handler.SeedProductRecords())
	seeder.POST("/product-batches", handler.SeedProductBatches())
	seeder.POST("/purchase-orders", handler.SeedPurchaseOrders())
	seeder.POST("/inbound-orders", handler.SeedInboundOrders())

	return r
}
