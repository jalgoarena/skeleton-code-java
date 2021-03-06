package main

import (
	"bytes"
	"github.com/jalgoarena/skeleton-code-java/api"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockHTTPClient struct {
	problemJSON string
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	response := &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(m.problemJSON))),
		StatusCode: http.StatusOK,
	}

	return response, nil
}

func TestGetSkeletonCodeForFib(t *testing.T) {
	var (
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

	httpClient := &MockHTTPClient{problemJSON: problemsJSON}
	api.SetupProblems(problemsURL, httpClient)

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/code/java/fib", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, javaSourceCode, w.Body.String())
}

func TestGetSkeletonCodeForTwoSum(t *testing.T) {
	var (
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

	httpClient := &MockHTTPClient{problemJSON: problemsJSON}
	api.SetupProblems(problemsURL, httpClient)

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/code/java/2-sum", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, javaSourceCode, w.Body.String())
}

func TestGetSkeletonCodeForWordLadder(t *testing.T) {
	var (
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

	httpClient := &MockHTTPClient{problemJSON: problemsJSON}
	api.SetupProblems(problemsURL, httpClient)

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/code/java/word-ladder", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, javaSourceCode, w.Body.String())
}

func TestGetSkeletonCodeForInsertRange(t *testing.T) {
	var (
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

	httpClient := &MockHTTPClient{problemJSON: problemsJSON}
	api.SetupProblems(problemsURL, httpClient)

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/code/java/insert-range", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, javaSourceCode, w.Body.String())
}

func BenchmarkGetSkeletonCodeForFib(b *testing.B) {
	httpClient := &MockHTTPClient{problemJSON: problemsJSON}
	api.SetupProblems(problemsURL, httpClient)

	testRouter := setupRouter()

	for i := 0; i < b.N; i++ {
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/code/java/fib", nil)
		testRouter.ServeHTTP(resp, req)

		if resp.Code != 200 {
			b.Errorf("GET /api/v1/code/java/fib failed with response code: %d", resp.Code)
		}
	}
}

const problemsJSON = `[{
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
}, {
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
}, {
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
}, {
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
}]`
