package integrationtests_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testClient struct {
	test   *testing.T
	client *http.Client
	url    string
	ctx    context.Context
}

func newTestClient(url string, t *testing.T, ctx context.Context) *testClient {
	return &testClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		url:  url,
		test: t,
		ctx:  ctx,
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
	req = req.WithContext(t.ctx)
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

func (t testClient) WithCookies(cookies []*http.Cookie) (testClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return testClient{}, err
	}
	u, err := url.Parse(t.url)
	if err != nil {
		return testClient{}, err
	}
	jar.SetCookies(u, cookies)
	t.client.Jar = jar
	return t, nil
}
