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

func (i *QueryResult) Entity() (*comprehend.Entity, error) {

	var (
		beginoffset int64
		endoffset   int64
		score       float64
		text        string
		entityType  string
	)

	err := i.rows.Scan(&beginoffset, &endoffset, &score, &text, &entityType)
	if err != nil {
		return nil, errors.Wrap(err, "could not read drom postgres")
	}

	var entity = comprehend.Entity{
		BeginOffset: &beginoffset,
		EndOffset:   &endoffset,
		Score:       &score,
		Text:        &text,
		Type:        &entityType,
	}
	return &entity, nil
}

func (i *QueryResult) Byline() (models.Byline, error) {

	var (
		name string
	)

	err := i.rows.Scan(&name)
	if err != nil {
		fmt.Println("Trouble 2", err)
	}

	return models.Byline{Name: name, Gender: models.Gender("")}, nil
}

func (i *QueryResult) EntityFromPostgresResult() (*models.EntityResult, error) {

	var (
		name        string
		gender      sql.NullString
		nextWord    sql.NullString
		score 		float64
		articleId   string
	)

	err := i.rows.Scan(&name, &gender, &nextWord, &score, &articleId)
	if err != nil {
		fmt.Println("Trouble", err)
		return nil, errors.Wrap(err, "Could not scan row")
	}

	return &models.EntityResult{
		Name: name,
		Gender: gender,
		NextWord: nextWord,
		Score: score,
		ArticleId: articleId,
	}, nil
}

func (i *QueryResult) Article() (models.Content, error) {

	var (
		id            string
		published     string
		content       string
		canonical_url string
		headline      string
		name          sql.NullString
		/**/ section string
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

func GetEntities(db *sql.DB, query string) ([]comprehend.Entity, error) {
	entities, err := QueryDb(db, query)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to run query in GetPeople")
	}

	var entityArray []comprehend.Entity

	for entities.Next() {

		entity, err := entities.Entity()

		if err == nil && entities != nil {
			entityArray = append(entityArray, *entity)
		}
	}
	return entityArray, nil

}

func GetBylines(db *sql.DB, query string) ([]models.Byline, error) {
	bylines, err := QueryDb(db, query)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to run query")
	}

	var bylinesArray []models.Byline

	for bylines.Next() {

		person, err := bylines.Byline()
		if err == nil {
			bylinesArray = append(bylinesArray, person)
		}
	}
	return bylinesArray, nil
}

func GetArticles(db *sql.DB, query string) ([]models.Content, error) {
	articles, err := QueryDb(db, query)

	if err != nil {
		return nil, errors.Wrap(err, "Could not run query")
	}

	var contentArrray []models.Content

	for articles.Next() {
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
