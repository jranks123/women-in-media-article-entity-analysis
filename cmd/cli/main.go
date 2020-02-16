package main

import (
	"fmt"
	"io/ioutil"
	"women-in-media-article-entity-analysis/internal"
	"women-in-media-article-entity-analysis/internal/utils"
	"flag"
)


func justEntities(query string) {
	_, err := internal.GetAndStoreArticleEntities(string(query))
	if err != nil {
		fmt.Println(err, "There was an error")
	} else {
		fmt.Println("Successfully extracted and stored entities")
	}
}

func justGenderAnalysis(query string, manuallyCorrectGender bool) {
	err := internal.RedoGenderAnalysis(string(query), manuallyCorrectGender)
	if err != nil {
		fmt.Println(err, "There was an error")
	} else {
		fmt.Println("Successfully redid gender analysis")
	}
}




func main() {
	queryCondition, err := ioutil.ReadFile("../query_condition.sql")
	if err != nil {
		println(err, "Couldn't read file ")
	} else {
		query := utils.ConstructContentAnalysisQuery(string(queryCondition))


		runType := flag.String("runType", "ENTITIES_AND_GENDER", "a string")
		manuallyCorrectGender := flag.Bool("manuallyCorrectGender", false, "a string")
		flag.Parse()
		fmt.Println("Possible values of runType: ENTITIES_AND_GENDER, JUST_ENTITIES, JUST_GENDER")
		fmt.Println("runType:", *runType)
		fmt.Println("manuallyCorrectGender:", *manuallyCorrectGender)

		if (*runType == "JUST_ENTITIES") {
			justEntities(query)
		} else if (*runType == "ENTITIES_AND_GENDER") {
			justEntities(query)
			justGenderAnalysis(query, *manuallyCorrectGender)
		} else if (*runType == "JUST_GENDER") {
			justGenderAnalysis(query, *manuallyCorrectGender)
		}
		err := internal.PrintResults(string(queryCondition))
		if err != nil {
			fmt.Println(err, "Wrote results to file")
		}
	}
}
