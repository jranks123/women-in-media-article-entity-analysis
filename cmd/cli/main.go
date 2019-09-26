package main

import (
	"fmt"
	"women-in-media-article-entity-analysis/internal"
)

func main() {
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article left join author on article.id  = author.id WHERE article.id = '4020' ORDER BY published::date ASC"

	err := internal.RedoGenderAnalysis(query, true)

	if err != nil {
		fmt.Println(err, "There was an error")
	} else {
		fmt.Println("Successfully redid gender analysis")
	}
}
