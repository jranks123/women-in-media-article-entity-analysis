package secretsanta

import (
	"encoding/json"
	"io/ioutil"
	"women-in-media-article-entity-analysis/internal/models"
)

func ReadJsonFile(fileName string) models.FriendList {
	file, _ := ioutil.ReadFile(fileName)
	friendList := models.FriendList{}

	_ = json.Unmarshal([]byte(file), &friendList)

	return friendList

}

func WriteJson(friendList models.FriendList, fileName string) {
	file, _ := json.MarshalIndent(friendList, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)
}
