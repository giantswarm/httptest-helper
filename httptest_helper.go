package test

import (
  "io/ioutil"
  "net/http"
  "strings"
  "net/http/httptest"

  "github.com/gorilla/mux"
)

func GetRequest(url string, header map[string]string) (int, string, error) {
  req, err := get(url, header)
  if err != nil {
    return 0, "", err
  }
  return processRequest(req)
}

func PostRequest(url string, data string, header map[string]string) (int, string, error) {
  req, err := post(url, data, header)
  if err != nil {
    return 0, "", err
  }
  return processRequest(req)
}

func CreateServer(router *mux.Router) *httptest.Server {
  return httptest.NewServer(router)
}

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
