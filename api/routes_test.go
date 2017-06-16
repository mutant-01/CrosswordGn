package api

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"net/url"
	"path"
)

// TestGenerateRouter tests whether routes and their available HTTP methods exists or not.
func TestGenerateRouter(t *testing.T) {
	testServer := httptest.NewServer(generateRouter())
	defer testServer.Close()
	testUrl, err := url.Parse(testServer.URL)
	if err != nil{
		t.Fatalf("url parse error: %v", err)
	}
	for _, u := range routes {
		testUrl.Path = path.Join(basePath, u.path)
		for method := range u.methods{
			t.Logf("testing HTTP %v on url: %v", method, testUrl)
			req, _ := http.NewRequest(method, testUrl.String(), nil)
			resp, err := http.DefaultClient.Do(req)
			if err != nil{
				t.Fatalf("HTTP protocol error: %v", err)
			}
			t.Logf("HTTP status code: %v", resp.StatusCode)
			if resp.StatusCode == http.StatusNotFound{
				t.Errorf("404 not found for url: %v", testUrl.String())
			}
		}
	}
}
