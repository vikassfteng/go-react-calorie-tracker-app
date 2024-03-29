package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/vikassfteng/go-react-calorie-tracker-app/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
}

