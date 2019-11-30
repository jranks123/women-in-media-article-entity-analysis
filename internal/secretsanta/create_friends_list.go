package secretsanta

import (
	"encoding/json"
	"io/ioutil"
	"women-in-media-article-entity-analysis/internal/models"
)

func ReadJsonFile() models.FriendList {
	file, _ := ioutil.ReadFile("relationships.json")
	friendList := models.FriendList{}

	_ = json.Unmarshal([]byte(file), &friendList)

	return friendList

}
