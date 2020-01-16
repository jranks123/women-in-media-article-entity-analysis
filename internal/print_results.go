package internal

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strconv"
	"women-in-media-article-entity-analysis/internal/models"
	"women-in-media-article-entity-analysis/internal/services"
	"women-in-media-article-entity-analysis/internal/utils"
)

func PrintResults(queryCondition string) error {
	query := utils.ConstructPrintResultsQuery(queryCondition)
	db, _, err := GetDbAndParameters(query)

	entities, err := services.QueryDb(db, query)

	if err != nil {
		return errors.Wrap(err, "Could not run query")
	}

	var entitiesArray []*models.EntityResult

	for entities.Next() {
		entity, err := entities.EntityResult()
		if err == nil{
			if utils.EntityPassesConfidenceChecks(entity.Name, entity.Score) {
				entitiesArray = append(entitiesArray, entity)
			}
		} else {
			fmt.Println("error parsing entity ")
		}
	}


	articleCount, menCount, womenCount, noGenderCount := utils.EntityCounts(entitiesArray)

	results := models.GenderCountResult{
		Entities: entitiesArray,
		NumberOfWomen: womenCount,
		NumberofMen: menCount,
		NumberOfGenderless: noGenderCount,
		NumberOfFemaleJournalists: 0,
		NumberOfMaleJournalists: 0,
		NumberOfArticles: articleCount,
	}

	resultString := ""
	resultString += ("Total number of articles: " + strconv.Itoa(results.NumberOfArticles) + "\n")
	resultString += ("Total number of women: " + strconv.Itoa(results.NumberOfWomen) + "\n")
	resultString += ("Total number of men: " + strconv.Itoa(results.NumberofMen)+ "\n")
	resultString += ("Named entities:" + "\n")
	for _, entity := range results.Entities {
		resultString += (entity.Name + " (" + entity.Gender.String + ")" + "\n" )
	}

	d1 := []byte(resultString)
	error := ioutil.WriteFile("../results.html", d1, 0644)
	if error != nil {
		println("problem writing file")
	}

	return nil
}
