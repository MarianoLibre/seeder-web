package main

import (
	"database/sql"
	"time"

	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	user := os.Getenv("USERNAME")
	pass := os.Getenv("PASS")
	dbName := os.Getenv("DATABASE")
	dataSource := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, pass, dbName)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	engine.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		}))

	router := NewRouter(engine, db)
	router.engine.Run()
}
