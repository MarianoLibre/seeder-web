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
	_, err = sql.Open("mysql", "%s:%s@/%s?parseTime=true")
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

	user := os.Getenv("USERNAME")
	pass := os.Getenv("PASS")
	database := os.Getenv("DATABASE")
	fmt.Printf("%s:%s@/%s?parseTime=true\n", user, pass, database)
}
