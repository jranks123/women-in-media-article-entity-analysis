package internal

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"women-in-media-article-entity-analysis/internal/models"
	"women-in-media-article-entity-analysis/internal/services"
	"women-in-media-article-entity-analysis/internal/utils"
)

func ConstructContentAnalysis(content models.Content, entities [] models.EntityWithNextWord, cacheHit bool) *models.ContentAnalysis {
	var bylines []*models.Byline = nil

	bylinesArray := strings.Split(strings.Replace(content.Fields.Byline, " and ", ",", -1), ",")

	for _, byline := range bylinesArray {
		bylines = append(bylines, &models.Byline{byline, ""})
	}

	var people []*models.Person = nil
	//var locations []*comprehend.Entity = nil
	//var organisations []*comprehend.Entity = nil
	//var creativeWorkTitles []*comprehend.Entity = nil
	//var commercialItems []*comprehend.Entity = nil
	//var events []*comprehend.Entity = nil

	for _, entity := range entities {
		if *entity.Entity.Type == "PERSON" {
			people = append(people, &models.Person{EntityWithNextWord: entity})
		}
		//if *entity.Type == "LOCATION" {
		//	locations = append(locations, entity)
		//}
		//if *entity.Type == "ORGANIZATION" {
		//	organisations = append(organisations, entity)
		//}
		//if *entity.Type == "COMMERCIAL_ITEM" {
		//	commercialItems = append(commercialItems, entity)
		//}
		//if *entity.Type == "TITLE" {
		//	creativeWorkTitles = append(creativeWorkTitles, entity)
		//}
		//if *entity.Type == "EVENT" {
		//	events = append(events, entity)
		//}

	}

	contentAnalysis := models.ContentAnalysis{
		Path:               content.Url,
		Headline:           content.Fields.Headline,
		BodyText:           content.Fields.BodyText,
		Bylines:            bylines,
		People:             people,
		//Locations:          locations,
		//Organisations:      organisations,
		//CreativeWorkTitles: creativeWorkTitles,
		//CommercialItems:    commercialItems,
		//Events:             events,
		CacheHit:           cacheHit,
		Section:            content.Section,
		WebPublicationDate: content.WebPublicationDate,
		Id:                 content.Id,
	}

	return &contentAnalysis
}

func AddGenderToContentAnalysisSlice(contentAnalysisSlice []*models.ContentAnalysis) ([]*models.ContentAnalysis, error) {
	var contentAnalysisWithGenderSlice []*models.ContentAnalysis
	for _, contentAnalysis := range contentAnalysisSlice {
		contentAnalysisWithGender, err := AddGenderToContentAnalysis(contentAnalysis)
		if err != nil {
			return nil, errors.Wrap(err, "Error appending gender for slice for "+contentAnalysis.Path)
		}
		contentAnalysisWithGenderSlice = append(contentAnalysisWithGenderSlice, contentAnalysisWithGender)
		return contentAnalysisWithGenderSlice, nil
	}
	return contentAnalysisWithGenderSlice, nil
}

func GetGenderAnalysisForName(name string) (*models.Gender, error) {
	genderAnalysis, err := services.GetGenderAnalysis(name)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting gender analysis for "+name)
	}

	if len(genderAnalysis.People) > 0 {
		if genderAnalysis.People[0].GenderGuess == "Female" {
			gender := models.Gender("Female")
			return &gender, nil
		}
		if genderAnalysis.People[0].GenderGuess == "Male" {
			gender := models.Gender("Male")
			return &gender, nil
		}
	}

	return nil, nil
}

func AddGenderToContentAnalysis(contentAnalysis *models.ContentAnalysis) (*models.ContentAnalysis, error) {
	for _, person := range contentAnalysis.People {
		gender, err := GetGenderAnalysisForName(*person.Entity.Text)
		if err != nil {
			return nil, errors.Wrap(err, "Error adding gender analysis for person "+*person.Entity.Text)
		}
		if gender != nil {
			person.Gender = *gender
		}
	}

	for _, byline := range contentAnalysis.Bylines {
		gender, err := GetGenderAnalysisForName(byline.Name)

		if err != nil {
			return nil, errors.Wrap(err, "Error getting gender analysis for byline "+byline.Name)
		}
		if gender != nil {
			byline.Gender = *gender
		}
	}

	return contentAnalysis, nil

}

func StoreArticleAnalysis(dbs *sql.DB, p services.JobParameters, entity *models.Person, element *models.ContentAnalysis) error {
	sqlStatement := "INSERT INTO article_entities (article_id, beginoffset, endoffset, score, text, type, nextWord) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := dbs.Exec(sqlStatement, element.Id, entity.Entity.BeginOffset, entity.Entity.EndOffset, entity.Entity.Score, entity.Entity.Text, entity.Entity.Type, entity.NextWord)
	if err != nil {
		return errors.Wrap(err, "Could not store article in article db")
	}
	return nil
}

func StorePersonGender(dbs *sql.DB, p services.JobParameters, name string, gender models.Gender) error {
	sqlStatement := "INSERT INTO names (name, gender) VALUES ($1, $2) ON conflict (name) do update set gender = $2"

	println("About to do " + name + " with gender " + string(gender))
	_, err := dbs.Exec(sqlStatement, name, gender)
	if err != nil {
		return errors.Wrap(err, "Could not store name in article db")
	}
	return nil

}

func StoreAllContentAnalysis(dbs *sql.DB, p services.JobParameters, contentAnalysisSlice []*models.ContentAnalysis) error {
	for _, element := range contentAnalysisSlice {
		for _, entity := range element.People {
			err := StoreArticleAnalysis(dbs, p, entity, element)
			if err != nil {
				return errors.Wrap(err, "Could not store article in article db")
			}
		}

		for _, byline := range element.Bylines {
			err := StorePersonGender(dbs, p, byline.Name, byline.Gender)
			if err != nil {
				return errors.Wrap(err, "Could not store byline gender")
			}
		}

	}

	println("Successfully stored entities")

	return nil
}

func GetDbAndParameters(query string) (*sql.DB, *services.JobParameters, error) {
	p := services.JobParameters{
		Db: services.DbParameters{
			DbName:   "public",
			Host:     "article-data.ckelnxbp6kie.us-east-2.rds.amazonaws.com",
			Port:     5432,
			User:     "article_data_master",
			Password: "AimangeiL2PhahNah5eXooB9quaiLoo7xi",
		},
		Query: query,
	}

	db, err := services.ConnectToPostgres(p.Db)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to connect to database")

	}

	return db, &p, nil

}

func MapGenderFromInputToGender(genderFromInput string) *models.Gender {
	var gender models.Gender
	if genderFromInput == "m" {
		gender = models.Gender("Male")
	} else if genderFromInput == "f" {
		gender = models.Gender("Female")
	} else if genderFromInput == "n" {
		gender = models.Gender("NotName")
	}
	return &gender
}

func GetGenderFromUserInput(name string) *models.Gender {
	//reading a string
	reader := bufio.NewReader(os.Stdin)
	var genderFromInput string
	var inputIsValid = false
	fmt.Println("Please enter f for Female, m for Male, n for Not a Name, u for Unknown")
	for inputIsValid == false {
		fmt.Println(" Enter gender for:")
		fmt.Println(name + " (m/f/u)")
		genderFromInput, _ = reader.ReadString('\n')
		genderFromInput = strings.TrimSpace(genderFromInput)

		if genderFromInput == "m" || genderFromInput == "f" || genderFromInput == "u" || genderFromInput == "n" {
			inputIsValid = true
		} else {
			fmt.Println("Please enter f for Female, m for Male, u for Unknown")
		}
	}

	gender := MapGenderFromInputToGender(genderFromInput)
	return gender

}

func RedoGenderAnalysis(query string, maunal bool) error {
	db, p, err := GetDbAndParameters(query)

	if err != nil {
		return errors.Wrap(err, "Couldn't get db and job parameters")

	}

	defer db.Close()

	contentSlice, err := services.GetArticleFields(db, *p)

	println("Content slice size:", len(contentSlice))

	if err != nil {
		fmt.Println(err, "Failed to get articles")
	}

	for _, element := range contentSlice {

		corrections := make(map[string]string)

		entitiesFromPostgres, err := services.GetEntitiesFromPostgres(element.Url)

		if err != nil {
			return errors.Wrap(err, "error checking if article has entities")
		}

		error := ComputeAndStoreGenderOfEntities(entitiesFromPostgres, maunal, corrections, db, p)
		if error != nil {
			return errors.Wrap(error, "error in compute and store gender of entities")
		}

		bylines, err := services.GetBylinesFromPostgres(element.Url)
		for _, byline := range bylines {
			println(byline.Name)
			gender, err := GetGenderAnalysisForName(byline.Name)

			if err != nil {
				return errors.Wrap(err, "Error getting gender analysis for "+byline.Name)
			}

			if gender == nil && maunal && utils.WordCount(byline.Name) > 1 {
				gender = GetGenderFromUserInput(byline.Name)
				corrections[byline.Name] = string(*gender)
			}

			if gender != nil {
				storeErr := StorePersonGender(db, *p, byline.Name, *gender)
				if storeErr != nil {
					return errors.Wrap(storeErr, "Error storing content analysis")
				}
			}
		}

		correctionsErr := services.StoreCorrections(corrections)
		if correctionsErr != nil {
			return errors.Wrap(correctionsErr, "Could not store corrections")
		}

	}
	return nil
}

func ComputeAndStoreGenderOfEntities(entities []models.Person, maunal bool, corrections map[string]string, db *sql.DB, p *services.JobParameters) error {
	corrections["a"] = "hello"

	for _, entitityWithNextWord := range entities {
		entity := entitityWithNextWord.Entity
		if *entity.Type == "PERSON" && utils.EntityPassesConfidenceChecks(*entity.Text, *entity.Score) {
			var gender *models.Gender
			gender, err := GetGenderAnalysisForName(*entity.Text)
			if err != nil {
				return errors.Wrap(err, "Error getting gender analysis for "+*entity.Text)
			}

			if gender == nil && maunal {
				gender = GetGenderFromUserInput(*entity.Text)
				corrections[*entity.Text] = string(*gender)
			}

			if gender != nil {
				storeErr := StorePersonGender(db, *p, *entity.Text, *gender)
				if storeErr != nil {
					return errors.Wrap(storeErr, "Error storing content analysis")
				}
			}
		}
	}
	return nil
}

func GetContentAnalysis(query string) ([]*models.ContentAnalysis, error) {

	db, p, err := GetDbAndParameters(query)

	if err != nil {
		return nil, errors.Wrap(err, "Couldn't get db and job parameters")
	}

	contentSlice, err := services.GetArticleFields(db, *p)

	var contentAnalysisSlice []*models.ContentAnalysis

	for _, element := range contentSlice {

		entitiesFromPostgres, err := services.GetEntitiesFromPostgres(element.Url)

		if err != nil {
			return nil, errors.Wrap(err, "error checking if article has entities")
		}

		if len(entitiesFromPostgres) == 0 {

			fmt.Println("about to get entities for article", element.Url)

			entities, err := services.GetEntitiesForArticle(element)
			if err != nil {
				return nil, errors.Wrap(err, "Error getting entities for article "+element.Url)
			}
			contentAnalysis := ConstructContentAnalysis(element, entities, false)
			contentAnalysisSlice = append(contentAnalysisSlice, contentAnalysis)
		} else {
			fmt.Println("already run analysis for ", element.Url)
		}
	}

	contentAnalysisSlice, err = AddGenderToContentAnalysisSlice(contentAnalysisSlice)

	storeErr := StoreAllContentAnalysis(db, *p, contentAnalysisSlice)

	if storeErr != nil {
		return nil, errors.Wrap(storeErr, "Error storing content analysis")
	}

	return contentAnalysisSlice, nil
}
