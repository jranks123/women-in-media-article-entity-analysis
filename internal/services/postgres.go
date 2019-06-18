package services

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"women-in-media-article-entity-analysis/internal/models"
)

type contributionsDatabase struct {
	underlying *sql.DB
	stmts      map[string]*sql.Stmt
}

func connectToPostgres(p DbParameters) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d",
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

func GetArticleFields() ([]models.Content, error) {

	p := JobParameters{
		Db: DbParameters{
			DbName:   "public",
			Host:     "article-data.ckelnxbp6kie.us-east-2.rds.amazonaws.com ",
			Port:     5432,
			User:     "article_data_master",
			Password: "AimangeiL2PhahNah5eXooB9quaiLoo7xi",
		},
		From:    "2019-06-17",
		To:      "2019-06-18",
		Section: "environment",
	}

	db, err := connectToPostgres(p.Db)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	defer db.Close()

	articles, err := GetArticles(db, ArticleQueryParams{From: p.From, To: p.To, Section: p.Section})
	if err != nil {
		return nil, errors.Wrap(err, "unable to get contributions")
	}

	return articles, nil
}
