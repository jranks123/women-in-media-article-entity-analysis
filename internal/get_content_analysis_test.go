package internal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetContentAnalysis(t *testing.T) {
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article left join author on article.id  = author.id WHERE article.id ='5377' ORDER BY published::date ASC"

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
