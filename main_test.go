package main

import (
	"bytes"
	"github.com/jalgoarena/skeleton-code-java/app"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type MockHttpClient struct {
	problemJson string
}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
	response := &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(m.problemJson))),
		StatusCode: http.StatusOK,
	}

	return response, nil
}

func TestMain(m *testing.M) {
	app.SetProblemsHost(problemsUrl)
	os.Exit(m.Run())
}

func TestGetSkeletonCodeForFib(t *testing.T) {
	var (
		problemJson = `{
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
		javaSourceCode = `import java.util.*;
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
	)

	httpClient := &MockHttpClient{problemJson: problemJson}
	app.SetHttpClient(httpClient)

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/code/java/fib", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, javaSourceCode, w.Body.String())
}

func TestGetSkeletonCodeTwoSum(t *testing.T) {
	var (
		problemJson = `{
  "id": "2-sum",
  "func": {
    "name": "twoSum",
    "returnStatement": {
      "type": "[I",
      "comment": " Indices of the two numbers",
      "generic": ""
    },
    "parameters": [
      {
        "name": "numbers",
        "type": "[I",
        "comment": "An array of Integer",
        "generic": ""
      },
      {
        "name": "target",
        "type": "java.lang.Integer",
        "comment": "target = numbers[index1] + numbers[index2]",
        "generic": ""
      }
    ]
  }
}`
		javaSourceCode = `import java.util.*;
import com.jalgoarena.type.*;

public class Solution {
    /**
     * @param numbers An array of Integer
     * @param target target = numbers[index1] + numbers[index2]
     * @return Indices of the two numbers
     */
    public int[] twoSum(int[] numbers, int target) {
        // Write your code here
    }
}`
	)

	httpClient := &MockHttpClient{problemJson: problemJson}
	app.SetHttpClient(httpClient)

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/code/java/2-sum", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, javaSourceCode, w.Body.String())
}
