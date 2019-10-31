package internal

import (
	"encoding/json"
	"fmt"
	"testing"
	"women-in-media-article-entity-analysis/internal/services"
)

func TestGetContentAnalysis(t *testing.T) {
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article left join author on article.id  = author.id WHERE article.canonical_url in ('https://www.dailymail.co.uk/tvshowbiz/article-7474785/Ellen-DeGeneres-adopts-new-poodle-puppy-introduces-studio-audience.html') ORDER BY published::date ASC"

	res, err := GetContentAnalysis(query)

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

func TestRedoGenderAnalysis(t *testing.T) {
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article left join author on article.id  = author.id WHERE article.canonical_url in ('https://www.dailymail.co.uk/tvshowbiz/article-7474785/Ellen-DeGeneres-adopts-new-poodle-puppy-introduces-studio-audience.html') ORDER BY published::date ASC"

	err := RedoGenderAnalysis(query, true)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("Successfully redid gender analysis")
	}
}

func TestComputeAndStoreGenderOfEntities(t *testing.T) {
	db, p, err := GetDbAndParameters("")

	if err != nil {
		t.Error(err)

	}

	defer db.Close()

	entites, err := services.GetEntitiesFromPostgres("https://www.dailymail.co.uk/tvshowbiz/article-7474785/Ellen-DeGeneres-adopts-new-poodle-puppy-introduces-studio-audience.html")
	if err != nil {
		t.Error(err)
	}

	corrections := make(map[string]string)
	error := ComputeAndStoreGenderOfEntities(entites, false, corrections, db, p)
	if error != nil {
		t.Error(err)
	}

	println(corrections["a"])
}
