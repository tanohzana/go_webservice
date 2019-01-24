package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articlesList = []article{
	article{ID: 1, Title: "This is the first title", Content: "This content is awesome"},
	article{ID: 2, Title: "This is the second title", Content: "This content is even more awesome"},
}

func getAllArticles() []article {
	return articlesList
}

func getArticleByID(articleId int) (*article, error) {
	for _, a := range articlesList {
		if a.ID == articleId {
			return &a, nil
		}
	}

	return nil, errors.New("article not found")
}
