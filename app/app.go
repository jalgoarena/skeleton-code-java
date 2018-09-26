package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/problems-store/domain"
	"io/ioutil"
	"net/http"
)

var problemsHost string

func SetProblemsHost(host string) {
	problemsHost = host
}

// curl -i http://localhost:8080/api/v1/code/java/fib
func GetSkeletonCode(c *gin.Context) {
	problemId := c.Param("problemId")

	problemUrl := fmt.Sprintf("%s/api/v1/problems/%s", problemsHost, problemId)
	resp, err := http.Get(problemUrl)

	if err != nil {
		c.String(500, "[err] %s: %v", c.Request.RequestURI, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		c.String(resp.StatusCode, "[err] %s: %v", problemUrl, string(body))
		return
	}

	problem, err := loadProblem(body)

	if err != nil {
		c.String(500, "Cannot parse problem %s: %v", problemId, err)
		return
	}

	// TODO: Step 2 - fill java source code template with data

	sourceCode := fmt.Sprintf("source code for %s", problem.Title)

	c.String(200, "%v for %v", sourceCode, problemId)
}

func loadProblem(problemsJson []byte) (*domain.Problem, error) {
	var problem domain.Problem

	if err := json.Unmarshal(problemsJson, &problem); err != nil {
		return nil, err
	}

	return &problem, nil
}
