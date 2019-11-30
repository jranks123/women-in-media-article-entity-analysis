package secretsanta

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestAssignGivingV2(t *testing.T) {
	friendList := ReadJsonFile()
	sortedFriends := SortFriends(friendList)
	friendsWithGifts := AssignGiving(sortedFriends)
	for _, friend := range friendsWithGifts.FriendList {
		println(friend.Name + " gives to " + *friend.GivesTo)
	}

	file, _ := json.MarshalIndent(friendList, "", " ")

	_ = ioutil.WriteFile("output.json", file, 0644)
}