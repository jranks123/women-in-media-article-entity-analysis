package internal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetContentAnalysis(t *testing.T) {
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article left join author on article.id  = author.id WHERE article.canonical_url in ('https://www.spectator.co.uk/2019/08/right-from-wrong-a-guide-to-the-new-european-politics/') ORDER BY published::date ASC"

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
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article left join author on article.id  = author.id WHERE article.canonical_url in ('https://www.spectator.co.uk/2019/08/right-from-wrong-a-guide-to-the-new-european-politics/') ORDER BY published::date ASC"

	err := RedoGenderAnalysis(query, true)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("Successfully redid gender analysis")
	}
}
