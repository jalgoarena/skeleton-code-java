package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/skeleton-code-java/app"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("health", app.HealthCheck)
	v1 := router.Group("api/v1")
	{
		v1.GET("/code/java/:problemId", app.GetSkeletonCode)
	}

	return router
}

var (
	problemsUrl string
	port        string
)

func init() {
	flag.StringVar(&problemsUrl, "problems-url", "http://localhost:8080", "Problems store url")
	flag.StringVar(&port, "port", "8081", "Port to listen on")
	flag.Parse()
}

func main() {
	app.SetProblemsHost(problemsUrl)
	app.SetHttpClient(&http.Client{})
	router := SetupRouter()
	router.Run(":" + port)
}
