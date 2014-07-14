package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
)

//
// Get requests.
//

func GetRequest(url, path string) (int, string, error) {
	return getRequest(url, path, map[string]string{})
}

func GetRequestWithHeader(url, path string, header map[string]string) (int, string, error) {
	return getRequest(url, path, header)
}

func getRequest(url, path string, header map[string]string) (int, string, error) {
	req, err := get(url+path, header)
	if err != nil {
		return 0, "", err
	}
	return processRequest(req)
}

func get(url string, header map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Request{}, err
	}

	for headerName, headerValue := range header {
		req.Header.Add(headerName, headerValue)
	}

	return req, nil
}

//
// Post requests.
//

func PostRequest(url, path, data string, header map[string]string) (int, string, error) {
	req, err := post(url+path, data, header)
	if err != nil {
		return 0, "", err
	}
	return processRequest(req)
}

func post(url string, data string, header map[string]string) (*http.Request, error) {
	ioReader := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, ioReader)
	if err != nil {
		return &http.Request{}, err
	}

	for headerName, headerValue := range header {
		req.Header.Add(headerName, headerValue)
	}

	return req, nil
}

//
// Requests.
//

func processRequest(req *http.Request) (int, string, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}

	return res.StatusCode, string(body), nil
}

//
// Server.
//

func CreateServer(router *mux.Router) *httptest.Server {
	return httptest.NewServer(router)
}
