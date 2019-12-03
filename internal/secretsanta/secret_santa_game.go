package secretsanta

import (
	"math/rand"
	"sort"
	"time"
	"women-in-media-article-entity-analysis/internal/models"
)


func PickRandomFriend(friends []string, friendsList []models.Friend) *models.Friend {
	rand.Seed(time.Now().UnixNano())
	return getFriend(friends[rand.Intn(len(friends))], friendsList)
}

func FriendIsAlreadyReceivingGift(friendToCheck string, friendsList []models.Friend) bool {
	for _, friend := range friendsList {
		if friend.GivesTo != nil && *friend.GivesTo == friendToCheck {
			return true
		}
	}
	return false
}

func SetGivesToForPersonrReceivingFrom(givingFriend *models.Friend, receivingFriend string, friendsList []models.Friend) []models.Friend {
	for i := 0; i < len(friendsList); i++ {
		if friendsList[i].Name == givingFriend.Name {
			friendsList[i].GivesTo = &receivingFriend
		}
	}
	return friendsList
}

func SortFriends(friendsList models.FriendList) models.FriendList {
	sort.Slice(friendsList.FriendList, func(i, j int) bool {
		return len(friendsList.FriendList[i].Knows) < len(friendsList.FriendList[j].Knows)
	})
	return friendsList
}

func getFriend(friendName string, friendsList []models.Friend) *models.Friend {
	for _, friend := range friendsList {
		if friend.Name == friendName {
			return &friend
		}
	}
	println(friendName)
	return nil
}


func AssignGiving(data models.FriendList) models.FriendList {

	friendsList := data.FriendList
	for i := 0; i < len(friendsList); i++ {
		// set gives to
		for friendsList[i].GivesTo == nil {
			potentialFriendToGiveTo := PickRandomFriend(friendsList[i].Knows, friendsList)
			if !FriendIsAlreadyReceivingGift(potentialFriendToGiveTo.Name, friendsList) &&  friendsList[i].Name != potentialFriendToGiveTo.Name {
				friendsList[i].GivesTo = &potentialFriendToGiveTo.Name
				if friendsList[i].Name == potentialFriendToGiveTo.Name {
					println("duplicate giving for " + *friendsList[i].GivesTo)
				}
			}
		}

		// set receives from
		for FriendIsAlreadyReceivingGift(friendsList[i].Name, friendsList) == false {
			potentialFriendToRecieveFrom := PickRandomFriend(friendsList[i].Knows, friendsList)
			if potentialFriendToRecieveFrom != nil && potentialFriendToRecieveFrom.GivesTo == nil && friendsList[i].Name != potentialFriendToRecieveFrom.Name {
				friendsList = SetGivesToForPersonrReceivingFrom(potentialFriendToRecieveFrom, friendsList[i].Name, friendsList)
				if *friendsList[i].GivesTo == potentialFriendToRecieveFrom.Name {
					println("duplicate receiving for " + *friendsList[i].GivesTo)
				}
			}
		}
	}

	return models.FriendList{friendsList}
}