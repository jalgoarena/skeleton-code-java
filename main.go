package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/skeleton-code-java/app"
	"os"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/code/java/:problemId", app.GetSkeletonCode)
	}

	return router
}

func main() {
	const defaultPort = "8080"

	router := SetupRouter()
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = defaultPort
	}

	router.Run(":" + port)
}
