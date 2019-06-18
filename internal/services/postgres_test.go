package services

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetArticleFields(t *testing.T) {
	res, err := GetArticleFields()

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
