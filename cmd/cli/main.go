package main

import (
	"fmt"
	"women-in-media-article-entity-analysis/internal"
)

func main() {
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article left join author on article.id  = author.id WHERE article.id = '1202' ORDER BY published::date ASC"
	println("hello2")
	_,err := internal.GetContentAnalysis(query)
println("hello4")
	if err != nil {
		fmt.Println(err, "There was an error")
	} else {
		fmt.Println("Successfully stored entities in Postgres")
		err := internal.RedoGenderAnalysis(query, true)
		if err != nil {
			fmt.Println(err, "There was an error")
		} else {
			fmt.Println("Successfully stored entities in Postgres")
		}
	}
}
