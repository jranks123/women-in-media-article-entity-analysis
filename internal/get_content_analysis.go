package internal

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/pkg/errors"
	"strings"
	"women-in-media-article-entity-analysis/internal/models"
	"women-in-media-article-entity-analysis/internal/services"
)

func ConstructContentAnalysis(content models.Content, entities []*comprehend.Entity, cacheHit bool) *models.ContentAnalysis {
	var bylines []*models.Byline = nil

	bylinesArray := strings.Split(strings.Replace(content.Fields.Byline, " and ", ",", -1), ",")

	for _, byline := range bylinesArray {
		bylines = append(bylines, &models.Byline{byline, ""})
	}

	var people []*models.Person = nil
	var locations []*comprehend.Entity = nil
	var organisations []*comprehend.Entity = nil
	var creativeWorkTitles []*comprehend.Entity = nil
	var commercialItems []*comprehend.Entity = nil
	var events []*comprehend.Entity = nil

	for _, entity := range entities {
		if *entity.Type == "PERSON" {
			people = append(people, &models.Person{Entity: *entity})
		}
		if *entity.Type == "LOCATION" {
			locations = append(locations, entity)
		}
		if *entity.Type == "ORGANIZATION" {
			organisations = append(organisations, entity)
		}
		if *entity.Type == "COMMERCIAL_ITEM" {
			commercialItems = append(commercialItems, entity)
		}
		if *entity.Type == "TITLE" {
			creativeWorkTitles = append(creativeWorkTitles, entity)
		}
		if *entity.Type == "EVENT" {
			events = append(events, entity)
		}

	}

	contentAnalysis := models.ContentAnalysis{
		Path:               content.Url,
		Headline:           content.Fields.Headline,
		BodyText:           content.Fields.BodyText,
		Bylines:            bylines,
		People:             people,
		Locations:          locations,
		Organisations:      organisations,
		CreativeWorkTitles: creativeWorkTitles,
		CommercialItems:    commercialItems,
		Events:             events,
		CacheHit:           cacheHit,
		Section:            content.Section,
		WebPublicationDate: content.WebPublicationDate,
	}

	return &contentAnalysis
}

func AddGenderToContentAnalysis(contentAnalysis *models.ContentAnalysis) (*models.ContentAnalysis, error) {
	for _, person := range contentAnalysis.People {
		genderAnalysis, err := services.GetGenderAnalysis(*person.Text)

		if err != nil {
			return nil, errors.Wrap(err, "Error getting gender analysis for "+*person.Text)
		}

		if len(genderAnalysis.People) > 0 {
			if genderAnalysis.People[0].GenderGuess == "Female" {
				person.Gender = "Female"
			}
			if genderAnalysis.People[0].GenderGuess == "Male" {
				person.Gender = "Male"
			}
		}
	}

	for _, person := range contentAnalysis.Bylines {
		genderAnalysis, err := services.GetGenderAnalysis(person.Name)

		if err != nil {
			return nil, errors.Wrap(err, "Error getting gender analysis for byline "+person.Name)
		}

		if len(genderAnalysis.People) > 0 {
			if genderAnalysis.People[0].GenderGuess == "Female" {
				person.Gender = "Female"
			}
			if genderAnalysis.People[0].GenderGuess == "Male" {
				person.Gender = "Male"
			}
		}
	}

	return contentAnalysis, nil

}

func GetContentAnalysis() ([]*models.ContentAnalysis, error) {
	contentSlice, err := services.GetArticleFields() //will return error if object is not in s3
	if err != nil {
		fmt.Println("Failed to get articles")
	}

	var contentAnalysisSlice []*models.ContentAnalysis

	for _, element := range contentSlice {

		// TODO: check to see if we've already run the analysis for this article

		entities, err := services.GetEntitiesForArticle(element)
		if err != nil {
			return nil, errors.Wrap(err, "Error getting entities for article "+element.Url)
		}
		contentAnalysis := ConstructContentAnalysis(element, entities, false)
		contentAnalysisWithGender, err := AddGenderToContentAnalysis(contentAnalysis)
		contentAnalysisSlice = append(contentAnalysisSlice, contentAnalysisWithGender)
	}

	return contentAnalysisSlice, nil

	// TODO: write a function that stores the content analysis
}
