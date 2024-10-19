// models.article_test.go

package main

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(articles) {
		t.Fail()
	}

	// Check that each member is identical
	for index, article := range alist {
		if article.Content != articles[index].Content ||
			article.ID != articles[index].ID ||
			article.Title != articles[index].Title {

			t.Fail()
			break
		}
	}
}
