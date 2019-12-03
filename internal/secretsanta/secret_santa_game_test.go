package secretsanta

import (
	"testing"
)

func TestAssignGivingV2(t *testing.T) {
	friendList := ReadJsonFile("relationships.json")
	sortedFriends := SortFriends(friendList)
	friendsWithGifts := AssignGiving(sortedFriends)
	//for _, friend := range friendsWithGifts.FriendList {
		//println(friend.Name + " gives to " + *friend.GivesTo)
	//}
	WriteJson(friendsWithGifts, "output.json")
}





