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

type QueryResult struct {
	rows *sql.Rows
}

func (i *QueryResult) Next() bool {
	hasNext := i.rows.Next()
	if !hasNext {
		i.rows.Close()
	}
	return hasNext
}

func ConnectToPostgres(p DbParameters) (*sql.DB, error) {
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
	db, err := ConnectToPostgres(p)
	if err != nil {
		return nil, err
	}

	return &contributionsDatabase{
		underlying: db,
		stmts:      make(map[string]*sql.Stmt),
	}, nil
}

func queryDb(db *sql.DB, query string) (*QueryResult, error) {
	rows, err := db.Query(
		query,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to run query")
	}

	articles := QueryResult{rows: rows}
	return &articles, nil
}

func GetArticleFields(db *sql.DB, p JobParameters) ([]models.Content, error) {

	articles, err := GetArticles(db, p.Query)

	if err != nil {
		return nil, errors.Wrap(err, "unable to get contributions")
	}

	return articles, nil
}

func GetEntitiesFromPostgres(url string) ([]models.Person, error) {

	p := JobParameters{
		Db: DbParameters{
			DbName:   "public",
			Host:     "article-data.ckelnxbp6kie.us-east-2.rds.amazonaws.com ",
			Port:     5432,
			User:     "article_data_master",
			Password: "AimangeiL2PhahNah5eXooB9quaiLoo7xi",
		},
		Query: fmt.Sprintf("SELECT beginoffset, endoffset, score, text, type as entityType FROM article join article_entities ON article.id = article_entities.article_id WHERE canonical_url = '%s' ORDER BY published ASC", url),
	}

	db, err := ConnectToPostgres(p.Db)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	defer db.Close()

	entities, err := GetPeople(db, p.Query)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting people from Postgres")
	}

	return entities, nil
}

func GetBylinesFromPostgres(url string) ([]models.Byline, error) {
	p := JobParameters{
		Db: DbParameters{
			DbName:   "public",
			Host:     "article-data.ckelnxbp6kie.us-east-2.rds.amazonaws.com",
			Port:     5432,
			User:     "article_data_master",
			Password: "AimangeiL2PhahNah5eXooB9quaiLoo7xi",
		},
		Query: fmt.Sprintf("SELECT name FROM article join author_attr ON article.id = author_attr.article_id join author on author_attr.author_id = author.id WHERE canonical_url = '%s' ORDER BY published ASC", url),
	}

	db, err := ConnectToPostgres(p.Db)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	defer db.Close()

	entities, err := GetBylines(db, p.Query)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting people from Postgres")
	}

	return entities, nil
}
