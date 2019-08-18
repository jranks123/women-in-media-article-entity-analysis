package main

import (
	"fmt"
	"women-in-media-article-entity-analysis/internal"
)

func main() {
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article left join author on article.id  = author.id WHERE article.canonical_url in ('https://www.theguardian.com/film/2019/may/09/emma-thompsons-best-films-ranked') ORDER BY published::date ASC"

	err := internal.RedoGenderAnalysis(query, true)

	if err != nil {
		fmt.Println("There was an error")
	} else {
		fmt.Println("Successfully redid gender analysis")
	}
}
