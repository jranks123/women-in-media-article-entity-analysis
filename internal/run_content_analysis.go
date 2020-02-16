package internal

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/service/comprehend"
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

	contentAnalysis := models.ContentAnalysis{
		Path:               content.Url,
		Headline:           content.Fields.Headline,
		BodyText:           content.Fields.BodyText,
		Bylines:            bylines,
		Entities:           entities,
		CacheHit:           cacheHit,
		Section:            content.Section,
		WebPublicationDate: content.WebPublicationDate,
		Id:                 content.Id,
	}

	return &contentAnalysis
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


func StoreEntity(dbs *sql.DB, entity *models.EntityWithNextWord, element *models.ContentAnalysis) error {
	sqlStatement := "INSERT INTO article_entities (article_id, beginoffset, endoffset, score, text, type, nextWord) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := dbs.Exec(sqlStatement, element.Id, entity.Entity.BeginOffset, entity.Entity.EndOffset, entity.Entity.Score, entity.Entity.Text, entity.Entity.Type, entity.NextWord)
	if err != nil {
		return errors.Wrap(err, "Could not store article in article db")
	}
	return nil
}

func StorePersonGender(dbs *sql.DB, p services.JobParameters, name string, gender models.Gender) error {
	sqlStatement := "INSERT INTO names (name, gender) VALUES ($1, $2) ON conflict (name) do update set gender = $2"

	println("About to store " + name + " with gender " + string(gender) )
	_, err := dbs.Exec(sqlStatement, name, gender)
	if err != nil {
		return errors.Wrap(err, "Could not store name in article db")
	}
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

func RedoGenderAnalysis(query string, manual bool) error {
	db, p, err := GetDbAndParameters(query)

	if err != nil {
		return errors.Wrap(err, "Couldn't get db and job parameters")

	}

	defer db.Close()

	contentSlice, err := services.GetArticlesArray(db, *p)

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



		error := ComputeAndStoreGenderOfEntities(entitiesFromPostgres, manual, corrections, db, p)
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

			if gender == nil && manual && utils.WordCount(byline.Name) > 1 {
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

		if len(corrections) > 0 {
			correctionsErr := services.StoreCorrections(corrections)
			if correctionsErr != nil {
				return errors.Wrap(correctionsErr, "Could not store corrections")
			}
		}


	}
	return nil
}

func ComputeAndStoreGenderOfEntities(entities []comprehend.Entity, manual bool, corrections map[string]string, db *sql.DB, p *services.JobParameters) error {
	corrections["a"] = "hello"

	for _, entity := range entities {
		if *entity.Type == "PERSON" && utils.EntityPassesConfidenceChecks(*entity.Text, *entity.Score) {
			var gender *models.Gender
			gender, err := GetGenderAnalysisForName(*entity.Text)
			if err != nil {
				return errors.Wrap(err, "Error getting gender analysis for "+*entity.Text)
			}

			_, valueFound := corrections[*entity.Text]

			if gender == nil && manual && !valueFound {
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

func GetAndStoreArticleEntities(query string) ([]*models.ContentAnalysis, error) {
	fmt.Println("About to get and store entites")
	db, p, err := GetDbAndParameters(query)

	if err != nil {
		return nil, errors.Wrap(err, "Couldn't get db and job parameters")
	}

	contentSlice, err := services.GetArticlesArray(db, *p)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting articles")
	}

	var contentAnalysisSlice []*models.ContentAnalysis
	fmt.Println("Number of articles: ")
	fmt.Println(len(contentSlice))

	for _, article := range contentSlice {

		entitiesFromPostgres, err := services.GetEntitiesFromPostgres(article.Url)

		if err != nil {
			return nil, errors.Wrap(err, "error checking if article has entities")
		}

		if len(entitiesFromPostgres) == 0 {

			fmt.Println("About to get entities for article", article.Url)

			entities, err := services.GetEntitiesForArticle(article)
			// fine here
			if err != nil {
				return nil, errors.Wrap(err, "Error getting entities for article "+article.Url)
			}
			contentAnalysis := ConstructContentAnalysis(article, entities, false)
			for _, entity := range contentAnalysis.Entities {
				if *entity.Entity.Type != "DATE" && *entity.Entity.Type != "QUANTITY" &&  *entity.Entity.Type != "OTHER" {
					err := StoreEntity(db, &entity, contentAnalysis)
					if err != nil {
						return nil, errors.Wrap(err, "Could not store article in article db")
					}
				}
			}
			fmt.Println("Finished storing entities for article " + article.Url)
			contentAnalysisSlice = append(contentAnalysisSlice, contentAnalysis)
		} else {
			fmt.Println("already run analysis for ", article.Url)
		}
	}

	closeError := db.Close()
	if closeError != nil {
		return nil, errors.Wrap(closeError, "Error closing db")
	}
	return contentAnalysisSlice, nil
}
