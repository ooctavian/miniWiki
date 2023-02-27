package integrationtests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testClient struct {
	test   *testing.T
	client *http.Client
	url    string
}

func newTestClient(url string, t *testing.T) *testClient {
	return &testClient{
		client: &http.Client{Timeout: 10 * time.Second},
		url:    url,
		test:   t,
	}
}

func (t testClient) newRequest(method string, path string, body interface{}) *http.Response {
	var bodyReader io.Reader
	if body != nil {
		bodyJson, err := json.Marshal(body)
		assert.NoError(t.test, err)
		bodyReader = bytes.NewReader(bodyJson)
	}
	req, err := http.NewRequest(method, t.url+path, bodyReader)
	req.Header.Set("content-type", "application/json")
	assert.NoError(t.test, err)
	res, err := t.client.Do(req)
	assert.NoError(t.test, err)
	return res
}

func (t testClient) Get(path string) *http.Response {
	return t.newRequest("GET", path, nil)
}

func (t testClient) Post(path string, body interface{}) *http.Response {
	return t.newRequest("POST", path, body)
}

func (t testClient) Put(path string, body interface{}) *http.Response {
	return t.newRequest("PUT", path, body)
}

func (t testClient) Patch(path string, body interface{}) *http.Response {
	return t.newRequest("PATCH", path, body)
}

func (t testClient) Delete(path string, body interface{}) *http.Response {
	return t.newRequest("PATCH", path, body)
}

func (t testClient) GetBody(response *http.Response) string {
	body, err := io.ReadAll(response.Body)
	assert.NoError(t.test, err)
	return string(body)
}
