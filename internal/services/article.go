package services

import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/pkg/errors"
	"women-in-media-article-entity-analysis/internal/models"
)

type ArticleQueryParams struct {
	From    string
	To      string
	Section string
}

func (i *QueryResult) Person() (models.Person, error) {

	var (
		beginoffset int64
		endoffset   int64
		score       float64
		text        string
		entityType  string
		gender      string
	)

	err := i.rows.Scan(&beginoffset, &endoffset, &score, &text, &entityType, &gender)
	if err != nil {
		fmt.Println("Trouble", err)
	}

	var entity = comprehend.Entity{
		BeginOffset: &beginoffset,
		EndOffset:   &endoffset,
		Score:       &score,
		Text:        &text,
		Type:        &entityType,
	}

	return models.Person{Entity: entity, Gender: models.Gender(gender)}, nil
}

func (i *QueryResult) Article() (models.Content, error) {

	var (
		id            string
		published     string
		content       string
		canonical_url string
		headline      string
		name          sql.NullString
		section       string
	)

	err := i.rows.Scan(&id, &published, &content, &canonical_url, &headline, &name, &section)
	if err != nil {
		fmt.Println("Trouble", err)
	}

	var byline string

	if name.Valid {
		byline = name.String
	} else {
		byline = ""
	}

	return models.Content{
		WebPublicationDate: published,
		Url:                canonical_url,
		Section:            section,
		Fields: models.ContentFields{
			Headline: headline,
			Byline:   byline,
			BodyText: content,
		},
		Id: id,
	}, nil
}

func GetPeople(db *sql.DB, query string) ([]models.Person, error) {
	people, err := queryDb(db, query)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to run query")
	}

	var peopleArray []models.Person

	for people.Next() {

		person, err := people.Person()
		if err == nil {
			peopleArray = append(peopleArray, person)
		}
	}
	return peopleArray, nil

}

func GetArticles(db *sql.DB, query string) ([]models.Content, error) {
	articles, err := queryDb(db, query)

	if err != nil {
		return nil, errors.Wrap(err, "Could not run query")
	}

	var contentArrray []models.Content

	for articles.Next() {
		println("artce")
		article, err := articles.Article()
		if err == nil {
			contentArrray = append(contentArrray, article)
		} else {
			fmt.Println("error")
		}
	}
	return contentArrray, nil
}

func GetArticle(db *sql.DB, query string) ([]models.Content, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	articles := QueryResult{rows: rows}

	var contentArrray []models.Content

	for articles.Next() {
		article, err := articles.Article()
		if err == nil {
			contentArrray = append(contentArrray, article)
		}
	}
	return contentArrray, nil
}
