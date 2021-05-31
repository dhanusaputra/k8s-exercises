package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	httpClient = &http.Client{Timeout: time.Second * 10}

	backendURL = os.Getenv("BACKEND_URL")
)

const defaultBackendURL = "http://backend-svc:8080"

// GraphqlResponse ...
type GraphqlResponse struct {
	Data   interface{}    `json:"data"`
	Errors []GraphqlError `json:"errors,omitempty"`
}

// Todos ...
type Todos struct {
	Todos []Todo `json:"todos"`
}

// Todo ...
type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// GraphqlError ...
type GraphqlError struct {
	Message string        `json:"message"`
	Path    []interface{} `json:"path,omitempty"`
}

// ReqBackend ...
func ReqBackend(query string) (*GraphqlResponse, int, error) {
	if len(backendURL) == 0 {
		backendURL = defaultBackendURL
	}

  req, err := http.NewRequest(http.MethodPost, backendURL+"/query", bytes.NewBufferString(query))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if !isStatusCode2xx(resp.StatusCode) {
		return nil, resp.StatusCode, fmt.Errorf("failed with statusCode: %d, body: %s", resp.StatusCode, string(body))
	}

	graphqlResp := &GraphqlResponse{Data: &Todos{}}

	if err := json.Unmarshal(body, graphqlResp); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(graphqlResp.Errors) > 0 {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed with graphql, err: %s", graphqlResp.Errors[0].Message)
	}

	return graphqlResp, resp.StatusCode, nil
}

func isStatusCode2xx(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
