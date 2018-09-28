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
      "comment": "N'th term of Fibonacci sequence",
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

func TestGetSkeletonCodeForTwoSum(t *testing.T) {
	var (
		problemJson = `{
  "id": "2-sum",
  "func": {
    "name": "twoSum",
    "returnStatement": {
      "type": "[I",
      "comment": "Indices of the two numbers",
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

func TestGetSkeletonCodeForWordLadder(t *testing.T) {
	var (
		problemJson = `{
  "id": "word-ladder",
  "func": {
    "name": "ladderLength",
    "returnStatement": {
      "type": "java.lang.Integer",
      "comment": "The shortest length",
      "generic": ""
    },
    "parameters": [
      {
        "name": "begin",
        "type": "java.lang.String",
        "comment": "the begin word",
        "generic": ""
      },
      {
        "name": "end",
        "type": "java.lang.String",
        "comment": "the end word",
        "generic": ""
      },
      {
        "name": "dict",
        "type": "java.util.HashSet",
        "comment": "the dictionary",
        "generic": "String"
      }
    ]
  }
}`
		javaSourceCode = `import java.util.*;
import com.jalgoarena.type.*;

public class Solution {
    /**
     * @param begin the begin word
     * @param end the end word
     * @param dict the dictionary
     * @return The shortest length
     */
    public int ladderLength(String begin, String end, HashSet<String> dict) {
        // Write your code here
    }
}`
	)

	httpClient := &MockHttpClient{problemJson: problemJson}
	app.SetHttpClient(httpClient)

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/code/java/word-ladder", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, javaSourceCode, w.Body.String())
}

func TestGetSkeletonCodeForInsertRange(t *testing.T) {
	var (
		problemJson = `{
  "id": "insert-range",
  "func": {
    "name": "insertRange",
    "returnStatement": {
      "type": "java.util.ArrayList",
      "comment": "Array with inserted ranges",
      "generic": "Interval"
    },
    "parameters": [
      {
        "name": "intervalsList",
        "type": "java.util.ArrayList",
        "comment": "sorted, non-overlapping list of Intervals",
        "generic": "Interval"
      },
      {
        "name": "insert",
        "type": "com.jalgoarena.type.Interval",
        "comment": "interval to insert",
        "generic": ""
      }
    ]
  }
}`
		javaSourceCode = `import java.util.*;
import com.jalgoarena.type.*;

public class Solution {
    /**
     * @param intervalsList sorted, non-overlapping list of Intervals
     * @param insert interval to insert
     * @return Array with inserted ranges
     */
    public ArrayList<Interval> insertRange(ArrayList<Interval> intervalsList, Interval insert) {
        // Write your code here
    }
}`
	)

	httpClient := &MockHttpClient{problemJson: problemJson}
	app.SetHttpClient(httpClient)

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/code/java/insert-range", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, javaSourceCode, w.Body.String())
}
