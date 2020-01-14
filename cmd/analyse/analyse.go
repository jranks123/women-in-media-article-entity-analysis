package main

import (
	"io/ioutil"
	"women-in-media-article-entity-analysis/internal"
)

func main() {
	query, err := ioutil.ReadFile("../query_condition.sql")
	if err != nil {
		println(err, "Couldn't read file ")
	} else {
		internal.PrintResults(string(query))
	}
}
