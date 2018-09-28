package main

import (
	"bytes"
	"github.com/jalgoarena/skeleton-code-java/app"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var fibProblemJson = `{
  "id": "fib",
  "func": {
    "name": "fib",
    "returnStatement": {
      "type": "java.lang.Long",
      "comment": " N'th term of Fibonacci sequence",
      "generic": ""
    },
    "parameters": [
      {
        "name": "n",
        "type": "java.lang.Integer",
        "comment": "id of fibonacci term to be returned",
        "generic": ""
      }
    ]
  }
}`

var fibSkeletonJavaSourceCode = `import java.util.*;
import com.jalgoarena.type.*;

public class Solution {
    /**
     * @param n id of fibonacci term to be returned
     * @return N'th term of Fibonacci sequence
     */
    public long fib(int n) {
        // Write your code here
    }
}`

type MockHttpClient struct{}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
	response := &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(fibProblemJson))),
		StatusCode: http.StatusOK,
	}

	return response, nil
}

func TestGetSkeletonCode(t *testing.T) {
	httpClient := &MockHttpClient{}
	app.SetProblemsHost(problemsUrl)
	app.SetHttpClient(httpClient)

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/code/java/fib", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, fibSkeletonJavaSourceCode, w.Body.String())
}
