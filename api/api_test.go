package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoStuffWithTestServer(t *testing.T) {

	//start the local HTTP server
	server := httptest.NewServer(http.HandleFunc(func(rw http.ResponseWriter, req *http.Request) {
		//Test request parameters
		equals(t, req.URL.String(), "some/path")
		//Send response to be tested
		rw.Write([]byte(`OK`))
	}))
	//close the server when test finished
	defer server.Close()
	//Use Client and URL from our local test server
	api := API{server.Client(), server.URL}
	body, err := api.DoStuff()
	ok(t, err)
	equals(t, []byte("OK"), body)
}
