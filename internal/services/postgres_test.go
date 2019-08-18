package services

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetArticleFields(t *testing.T) {
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article join author on article.id  = author.id ORDER BY published::date ASC"

	p := JobParameters{
		Db: DbParameters{
			DbName:   "public",
			Host:     "article-data.ckelnxbp6kie.us-east-2.rds.amazonaws.com ",
			Port:     5432,
			User:     "article_data_master",
			Password: "AimangeiL2PhahNah5eXooB9quaiLoo7xi",
		},
		Query: query,
	}

	db, err := ConnectToPostgres(p.Db)

	res, err := GetArticleFields(db, p)

	if err != nil {
		t.Error(err)
	} else {
		res, err := json.Marshal(res)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(string(res))
		}
	}
}

func TestCheckIfArticleHasEntities(t *testing.T) {
	res, err := GetEntitiesFromPostgres("https://phescreening.blog.gov.uk/2019/05/21/make-sure-your-leaflets-do-not-get-lost-in-the-post/")
	if err != nil {
		t.Error(err)
	} else {
		res, err := json.Marshal(res)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(string(res))
		}
	}
}
