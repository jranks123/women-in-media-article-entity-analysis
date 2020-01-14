package main

import (
	"fmt"
	"io/ioutil"
	"women-in-media-article-entity-analysis/internal"
	"women-in-media-article-entity-analysis/internal/utils"
)

func main() {
	queryCondition, err := ioutil.ReadFile("/Users/jonathan_rankin/code/women-in-media-article-entity-analysis/cmd/query_condition.sql")
	if err != nil {
		println(err, "Couldn't read file ")
	} else {
		query := utils.ConstructContentAnalysisQuery(string(queryCondition))
		println(query)
		_, err := internal.GetContentAnalysis(string(query))
		if err != nil {
			fmt.Println(err, "There was an error")
		} else {
			err := internal.RedoGenderAnalysis(string(query), true)
			if err != nil {
				fmt.Println(err, "There was an error")
			} else {
				fmt.Println("Successfully did Gender Analysis")
				err := internal.PrintResults(string(queryCondition))
				if err != nil {
					fmt.Println(err, "Wrote results to file")
				}
			}
		}
	}
}
