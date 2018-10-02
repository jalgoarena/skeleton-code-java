package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/skeleton-code-java/api"
	"log"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("health", api.HealthCheck)
	v1 := router.Group("api/v1")
	{
		v1.GET("/code/java/:problemId", api.GetSkeletonCode)
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

	log.SetFlags(log.LstdFlags)
}

func main() {
	api.SetupProblems(problemsUrl, &http.Client{})
	router := SetupRouter()
	router.Run(":" + port)
}
