package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := (w.Code == http.StatusOK)

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home page</title>") > 0

		return statusOK && pageOK
	})
}

func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	req.Header.Set("Accept", "application/JSON")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		typeOK := (w.HeaderMap.Get("content-type") == "application/JSON")

		return typeOK
	})
}

func TestArticleListXML(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	req.Header.Set("Accept", "application/XML")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		typeOK := (w.HeaderMap.Get("content-type") == "application/XML")

		return typeOK
	})
}
