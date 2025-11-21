package main

import (
	"ktfs/config"
	"ktfs/database"
	"ktfs/route"

	"fmt"
	"flag"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	// "github.com/gin-gonic/gin"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	config.ConnectDB()
}

func serveApplication() {
	// gin.SetMode(gin.ReleaseMode)
	for {
		router := route.Api()
		fmt.Println("API is ready to use")
		router.Run(":8080")
		time.Sleep(1 * time.Minute)
	}
}

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	loadEnv()
	loadDatabase()

	var migrate bool
	
	flag.BoolVar(&migrate, "migrate", false, "")
	flag.Parse()

	if migrate {
		database.Migrate()
	} else {
		serveApplication()
	}
}