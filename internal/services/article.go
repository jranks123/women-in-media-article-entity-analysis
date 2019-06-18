package services

import (
	"database/sql"
	"fmt"
	"women-in-media-article-entity-analysis/internal/models"
)

type ArticleQueryParams struct {
	From    string
	To      string
	Section string
}

type Article struct {
	published     string
	content       string
	canonical_url string
}

type Articles struct {
	rows *sql.Rows
}

func (i *Articles) Next() bool {
	hasNext := i.rows.Next()
	if !hasNext {
		i.rows.Close()
	}
	return hasNext
}

func (i *Articles) Article() (models.Content, error) {

	var (
		id            string
		published     string
		content       string
		canonical_url string
		headline      string
		name          string
		section       string
	)

	err := i.rows.Scan(&id, &published, &content, &canonical_url, &headline, &name, &section)
	if err != nil {
		i.rows.Close()
		return models.Content{}, err
	}

	return models.Content{
		WebPublicationDate: published,
		Section:            section,
		Fields: models.ContentFields{
			Headline: headline,
			Byline:   name,
			BodyText: content,
		},
		Id: id,
	}, nil
}

func GetArticles(db *sql.DB, params ArticleQueryParams) ([]models.Content, error) {
	rows, err := db.Query(
		"SELECT article.id as id, published, content, canonical_url, headline, name, section FROM article join author on article.id  = author.id  ORDER BY published::date ASC limit 10",
	)
	if err != nil {
		return nil, err
	}

	articles := Articles{rows: rows}

	var contentArrray []models.Content

	for articles.Next() {
		article, err := articles.Article()
		if err == nil {
			fmt.Print(article.Section)
			contentArrray = append(contentArrray, article)
		}
	}
	return contentArrray, nil
}
