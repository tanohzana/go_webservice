package main

import "testing"

func testGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articlesList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.ID != articlesList[i].ID ||
			v.Title != articlesList[i].Content ||
			v.Content != articlesList[i].Content {
			t.Fail()
			break
		}
	}
}
