package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/go-htmx/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))

	r.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		fmt.Println(err) // TODO: report err
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	}))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("'%s' path not found", c.Request.URL.Path),
		})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": fmt.Sprintf("'%s' method not allowed", c.Request.Method),
		})
	})

	routes.ApiRouter(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Fatal(r.Run(":" + port))
}
