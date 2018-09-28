package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/problems-store/domain"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

var (
	tmpl = `import java.util.*;
import com.jalgoarena.type.*;

public class Solution {
    /**
{{ range .Function.Parameters }}     * @param {{ .Name }} {{ .Comment }}{{ end }}
     * @return{{ .Function.Return.Comment }}
     */
    public {{ javaTypeDeclaration .Function.Return.Type .Function.Return.Generic }} {{ .Function.Name }}({{ methodParameters .Function.Parameters }}) {
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
	t.Funcs(template.FuncMap{
		"javaTypeDeclaration": javaTypeDeclaration,
		"methodParameters":    methodParameters,
	})
	t, _ = t.Parse(tmpl)

	buf := new(bytes.Buffer)
	err := t.Execute(buf, *problem)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
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

func loadProblem(problemsJson []byte) (*domain.Problem, error) {
	var problem domain.Problem

	if err := json.Unmarshal(problemsJson, &problem); err != nil {
		return nil, err
	}

	return &problem, nil
}
