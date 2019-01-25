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

	req.Header.Set("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		mimeType := getMIMETypeFromHeader(w.HeaderMap.Get("content-type"))
		expected := "application/json"
		typeOK := (mimeType == expected)

		return typeOK
	})
}

func TestArticleListXML(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	req.Header.Set("Accept", "application/xml")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		mimeType := getMIMETypeFromHeader(w.HeaderMap.Get("content-type"))
		expected := "application/xml"
		typeOK := (mimeType == expected)

		return typeOK
	})
}

func getMIMETypeFromHeader(contentTypeHeader string) string {
	contentTypeHeader = strings.ToLower(contentTypeHeader)
	contentTypeSplit := strings.Split(contentTypeHeader, ";")
	mimeType := contentTypeSplit[0]

	return mimeType
}
