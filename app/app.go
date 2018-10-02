package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/problems-store/app"
	"github.com/jalgoarena/problems-store/domain"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var (
	tmpl = `import java.util.*;
import com.jalgoarena.type.*;

public class Solution {
    /**
{{ $last := (indexOfLastElement .Function.Parameters )}}{{ range $index, $param := .Function.Parameters }}     * @param {{ $param.Name }} {{ $param.Comment }}{{if ne $index $last}}{{print "\n"}}{{ end }}{{ end }}
     * @return {{ .Function.Return.Comment }}
     */
    public {{ javaTypeDeclaration .Function.Return.Type .Function.Return.Generic }} {{ .Function.Name }}({{ methodParameters .Function.Parameters }}) {
        // Write your code here
    }
}`
	problemsHost string
	client       HttpClient
	problems     []domain.Problem
)

func SetupProblems(host string, httpClient HttpClient) {
	log.Printf("[INFO] Using problems host: %s\n", host)

	client = httpClient
	problemsHost = host

	downloadProblems()
}

func downloadProblems() {
	problemsUrl := fmt.Sprintf("%s/api/v1/problems", problemsHost)

	resp, err := client.Get(problemsUrl)

	if err != nil {
		log.Printf("[ERROR] could not download problems: %v\n", err)
		return
	}

	defer resp.Body.Close()
	err = loadProblems(resp.Body)

	if err != nil {
		log.Printf("[ERROR] could not read problems: %v\n", err)
		return
	}

	log.Printf("[INFO] problems loaded, count = %d", len(problems))
}

func loadProblems(problemsJson io.Reader) error {
	jsonParser := json.NewDecoder(problemsJson)

	if err := jsonParser.Decode(&problems); err != nil {
		return err
	}

	return nil
}

type HttpClient interface {
	Get(string) (*http.Response, error)
}

func HealthCheck(c *gin.Context) {
	_, err := client.Get(fmt.Sprintf("%s/health", problemsHost))

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "fail",
			"reason": fmt.Sprintf("problems service unavailable: %s", problemsHost),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "lang": "java"})
}

// curl -i http://localhost:8080/api/v1/code/java/fib
func GetSkeletonCode(c *gin.Context) {
	problemId := c.Param("problemId")

	if len(problems) == 0 {
		c.String(http.StatusInternalServerError, "[err] %s: problems could not be downloaded", c.Request.RequestURI)
		return
	}

	problem := app.First(problems, func(problem domain.Problem) bool {
		return problem.Id == problemId
	})

	sourceCode, err := buildSourceCode(&problem)
	if err != nil {
		c.String(http.StatusInternalServerError, "Cannot process template %s: %v", problemId, err)
	}

	c.String(http.StatusOK, "%v", sourceCode)
}

func buildSourceCode(problem *domain.Problem) (string, error) {
	t := template.New("sourceCode")
	t.Funcs(template.FuncMap{
		"javaTypeDeclaration": javaTypeDeclaration,
		"methodParameters":    methodParameters,
		"indexOfLastElement":  indexOfLastElement,
	})
	t, _ = t.Parse(tmpl)

	buf := new(bytes.Buffer)
	err := t.Execute(buf, *problem)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func indexOfLastElement(params []domain.Parameter) int {
	return len(params) - 1
}

func methodParameters(parameters []domain.Parameter) string {
	parametersAsStrings := mapParameters(&parameters, func(parameter domain.Parameter) string {
		return fmt.Sprintf("%s %s", javaTypeDeclaration(parameter.Type, parameter.Generic), parameter.Name)
	})

	return strings.Join(parametersAsStrings, ", ")
}

func mapParameters(parameters *[]domain.Parameter, f func(domain.Parameter) string) []string {
	result := make([]string, len(*parameters))
	for i, v := range *parameters {
		result[i] = f(v)
	}
	return result
}

func javaTypeDeclaration(returnType string, generic string) string {
	switch returnType {
	case "void":
		return returnType
	case "[I":
		return "int[]"
	case "[D":
		return "double[]"
	case "[[I":
		return "int[][]"
	case "[[C":
		return "char[][]"
	}

	simpleType := simpleType(returnType)
	return simpleOrGenericType(simpleType, generic)
}

func simpleOrGenericType(simpleType string, generic string) string {
	switch simpleType {
	case "Boolean":
		return "bool"
	case "Long":
		return "long"
	case "Integer":
		return "int"
	case "Double":
		return "double"
	default:
		return withGeneric(simpleType, generic)
	}
}

func simpleType(typeWithPackage string) string {
	typeParts := strings.Split(typeWithPackage, ".")
	return typeParts[len(typeParts)-1]
}

func withGeneric(baseType string, generic string) string {
	switch generic {
	case "":
		return baseType
	}

	return fmt.Sprintf("%s<%s>", baseType, generic)
}
