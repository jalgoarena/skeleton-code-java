package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/problems-store/domain"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tmpl = `import java.util.*;
import com.jalgoarena.type.*;
public class Solution {
	{{ .Title }} {
        // Write your code here
    }
}
`

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

	sourceCode := buildSourceCode(problem)

	c.String(200, "%v", sourceCode)
}

func buildSourceCode(problem *domain.Problem) string {
	t := template.New("sourceCode")
	t, _ = t.Parse(tmpl)

	buf := new(bytes.Buffer)
	t.Execute(buf, *problem)

	return buf.String()
}

func loadProblem(problemsJson []byte) (*domain.Problem, error) {
	var problem domain.Problem

	if err := json.Unmarshal(problemsJson, &problem); err != nil {
		return nil, err
	}

	return &problem, nil
}
