package internal

import (
	"encoding/json"
	"fmt"
	"testing"
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
