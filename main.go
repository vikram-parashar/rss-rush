package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vikram-parashar/rss-rush/api"
	"github.com/vikram-parashar/rss-rush/database"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("cant get PORT from .env")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalf("cant get DB_URL from .env")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbConn := database.New(db)

	r := gin.Default()
	api.SetupRoutes(r, dbConn)

	go startScrapping(5, 1*time.Minute, dbConn)

	r.Run(fmt.Sprintf(":%v", port))
	log.Println("hi")
}
