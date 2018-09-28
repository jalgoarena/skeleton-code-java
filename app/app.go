package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/problems-store/domain"
	"io/ioutil"
	"net/http"
	"text/template"
)

var (
	tmpl = `import java.util.*;
import com.jalgoarena.type.*;

public class Solution {
    /**
{{ range .Function.Parameters }}    * @param {{ .Name }} {{ .Comment }}{{ end }}
    * @return {{ .Function.Return.Comment }}
    */
    public void {{ .Function.Name }}() {
        // Write your code here
    }
}`
	problemsHost string
	client       HttpClient
)

func SetProblemsHost(host string) {
	problemsHost = host
}

type HttpClient interface {
	Get(string) (*http.Response, error)
}

func SetHttpClient(httpClient HttpClient) {
	client = httpClient
}

// curl -i http://localhost:8080/api/v1/code/java/fib
func GetSkeletonCode(c *gin.Context) {
	problemId := c.Param("problemId")
	problemUrl := fmt.Sprintf("%s/api/v1/problems/%s", problemsHost, problemId)

	resp, err := client.Get(problemUrl)

	if err != nil {
		c.String(http.StatusInternalServerError, "[err] %s: %v", c.Request.RequestURI, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		c.String(resp.StatusCode, "[err] %s: %v", problemUrl, string(body))
		return
	}

	problem, err := loadProblem(body)

	if err != nil {
		c.String(http.StatusInternalServerError, "Cannot parse problem %s: %v", problemId, err)
		return
	}

	sourceCode, err := buildSourceCode(problem)
	if err != nil {
		c.String(http.StatusInternalServerError, "Cannot process template %s: %v", problemId, err)
	}

	c.String(http.StatusOK, "%v", sourceCode)
}

func buildSourceCode(problem *domain.Problem) (string, error) {
	t := template.New("sourceCode")
	t, _ = t.Parse(tmpl)

	buf := new(bytes.Buffer)
	err := t.Execute(buf, *problem)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func loadProblem(problemsJson []byte) (*domain.Problem, error) {
	var problem domain.Problem

	if err := json.Unmarshal(problemsJson, &problem); err != nil {
		return nil, err
	}

	return &problem, nil
}
