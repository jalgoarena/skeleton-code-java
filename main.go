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
	SetupProblemsHost()
	router := SetupRouter()

	const defaultPort = "8081"
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = defaultPort
	}

	router.Run(":" + port)
}

func SetupProblemsHost() {
	const defaultProblemsHost = "http://localhost:8080"

	problemsHost := os.Getenv("PROBLEMS_HOST")

	if len(problemsHost) == 0 {
		problemsHost = defaultProblemsHost
	}

	app.SetProblemsHost(problemsHost)
}
