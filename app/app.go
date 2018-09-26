package app

import (
	"github.com/gin-gonic/gin"
)

// curl -i http://localhost:8080/api/v1/code/java/fib
func GetSkeletonCode(c *gin.Context) {
	problemId := c.Param("problemId")

	// TODO: Step 1 - get problem from problem-store service
	// TODO: Step 2 - fill java source code template with data

	sourceCode := "source code"

	c.String(200, "%v for %v", sourceCode, problemId)
}
