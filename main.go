package main

import (
	"log"
	"os"

	"example.com/go-htmx/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATH", "DELETE", "OPTIONS"},
	}))

	routes.ApiRouter(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Fatal(r.Run(":" + port))
}
