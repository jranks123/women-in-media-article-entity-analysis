package services

import (
	"math/rand"
	"sort"
	"time"
	"women-in-media-article-entity-analysis/internal/models"
)


func InitialiseFriendsWithDoesNotKnow() []models.Friend {

	friendsList := []models.Friend{}
	JakeG := models.Friend{1, "JakeG", "", []*models.Friend{}, []*models.Friend{}, nil}
	Jessie := models.Friend{1, "Jessie", "", []*models.Friend{}, []*models.Friend{}, nil}
	Micha := models.Friend{3, "Micha", "", []*models.Friend{}, []*models.Friend{}, nil}
	Dylan := models.Friend{4, "Dylan", "", []*models.Friend{}, []*models.Friend{}, nil }
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Micha)

	friendsList = append(friendsList, Jessie)
	friendsList = append(friendsList, JakeG)
	friendsList = append(friendsList, Micha)
	friendsList = append(friendsList, Dylan)


	return friendsList
}


func RemoveFriendFromKnowsList(friend models.Friend, unknownFriend models.Friend) models.Friend {
	n := 0
	for _, knownFriend := range friend.Knows {
		if knownFriend.Id != unknownFriend.Id {
			friend.Knows[n] = knownFriend
			n++
		}
	}
	friend.Knows = friend.Knows[:n]
	return friend
}

func CreateDoesKnow(friendsList []models.Friend)  []models.Friend {
	// add all friends to Does Know
	for i := 0;  i< len(friendsList); i++ {
		for j := 0;  j< len(friendsList); j++ {
			if friendsList[i].Id != friendsList[j].Id {
				friendsList[i].Knows = append(friendsList[i].Knows, &friendsList[j])
			}
		}
	}

	return friendsList
}


func FilterDoesKnow(friendsList []models.Friend)  []models.Friend{
	for i := 0; i < len(friendsList); i++ {
		for _, unknownFriend := range friendsList[i].DoesNotKnow {
			//remove unknown friend from Does Know
			friendsList[i] = RemoveFriendFromKnowsList(friendsList[i], *unknownFriend)
		}
	}
	return friendsList
}

func SortFriends(friendsList []models.Friend) []models.Friend {
	sort.Slice(friendsList, func(i, j int) bool {
		return len(friendsList[i].Knows) < len(friendsList[j].Knows)
	})
	return friendsList
}

func PickRandomFriend(friends []*models.Friend) models.Friend {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	return *friends[r.Intn(len(friends))]
}

func FriendIsAlreadyReceivingGift(friendToCheck models.Friend, friendsList []models.Friend ) bool {
	for _, friend := range friendsList {
		if friend.GivesTo != nil && friend.GivesTo.Name == friendToCheck.Name {
			return true
		}
	}
	return false
}

func SetGivesToForPersonrReceivingFrom(givingFriend models.Friend, receivingFriend models.Friend, friendsList []models.Friend ) []models.Friend {
	for i := 0; i < len(friendsList); i++ {
		if friendsList[i].Name == givingFriend.Name {
			friendsList[i].GivesTo = &receivingFriend
		}
	}
	return friendsList
}

func AssignGiving(friendsList []models.Friend) []models.Friend {

	for i := 0; i < len(friendsList); i++ {

		// set gives to
		for friendsList[i].GivesTo == nil {
			potentialFriendToGiveTo := PickRandomFriend(friendsList[i].Knows)
			if !FriendIsAlreadyReceivingGift(potentialFriendToGiveTo, friendsList) {
				friendsList[i].GivesTo = &potentialFriendToGiveTo
			}
		}

		// set receives from
		for FriendIsAlreadyReceivingGift(friendsList[i], friendsList) == false {
			potentialFriendToRecieveFrom := PickRandomFriend(friendsList[i].Knows)
			if potentialFriendToRecieveFrom.GivesTo == nil {
				friendsList = SetGivesToForPersonrReceivingFrom(potentialFriendToRecieveFrom, friendsList[i], friendsList)
			}
		}
	}

	return friendsList
}