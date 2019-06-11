package services

import (
	"database/sql"
	"fmt"
	"women-in-media-article-entity-analysis/internal/models"
)

type contributionsDatabase struct {
	underlying *sql.DB
	stmts      map[string]*sql.Stmt
}

func connectToPostgres(p DbParameters) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"dbname=contributions user=%s password=%s host=%s port=%d",
		p.User,
		p.Password,
		p.Host,
		p.Port,
	)
	return sql.Open("postgres", connStr)
}

func newContributionsDatabase(p DbParameters) (*contributionsDatabase, error) {
	db, err := connectToPostgres(p)
	if err != nil {
		return nil, err
	}

	return &contributionsDatabase{
		underlying: db,
		stmts:      make(map[string]*sql.Stmt),
	}, nil
}

func getArticleFieldsFromUrl(url string) (*models.Content, error) {
	// call postgres with query
	return nil, nil
}
