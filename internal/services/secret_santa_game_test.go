package services

import (
	"testing"
	"women-in-media-article-entity-analysis/internal/models"
)

func TestCreateDoesKnow(t *testing.T) {
	friendsList := InitialiseFriendsWithDoesNotKnow()
	friendsWithFullKnows := CreateDoesKnow(friendsList)

	for _, friend := range friendsWithFullKnows {
		println(friend.Name)
		println("knows")
		for _, knownFriend := range friend.Knows {
			println(knownFriend.Name)
		}
		println("and")

	}
}

func TestRemoveFriendFromKnowsList(t *testing.T) {
	Jessie := models.Friend{1, "Jessie", "", []*models.Friend{}, []*models.Friend{}, nil}
	JakeG := models.Friend{2, "JakeG", "", []*models.Friend{}, []*models.Friend{}, nil}
	Micha := models.Friend{3, "Micha", "", []*models.Friend{}, []*models.Friend{}, nil}
	Dylan := models.Friend{4, "Dylan", "", []*models.Friend{}, []*models.Friend{}, nil}

	Jessie.Knows = append(Jessie.Knows, &JakeG)
	Jessie.Knows = append(Jessie.Knows, &Micha)
	Jessie.Knows = append(Jessie.Knows, &Dylan)

	Jessie.DoesNotKnow = append(Jessie.DoesNotKnow, &Dylan)

	res := RemoveFriendFromKnowsList(Jessie, Dylan)
	for _, friend := range res.Knows {
		println(friend.Name)
	}

}

func TestFilterDoesNotKnow(t *testing.T) {
	friendsList := InitialiseFriendsWithDoesNotKnow()
	friendsWithFullKnows := CreateDoesKnow(friendsList)
	friendsWithDoesNotKnowRemoved := FilterDoesKnow(friendsWithFullKnows)

	for _, friend := range friendsWithDoesNotKnowRemoved {
		println(friend.Name)
		println("knows")
		for _, knownFriend := range friend.Knows {
			println(knownFriend.Name)
		}
		println("and")

	}
}

func TestSortFriends(t *testing.T) {
	friendsList := InitialiseFriendsWithDoesNotKnow()
	friendsWithFullKnows := CreateDoesKnow(friendsList)
	friendsWithDoesNotKnowRemoved := FilterDoesKnow(friendsWithFullKnows)
	sortedFriends := SortFriends(friendsWithDoesNotKnowRemoved)
	for _, friend := range sortedFriends {
		println(friend.Name)
	}
}

func TestAssignGiving(t *testing.T) {
	friendsList := InitialiseFriendsWithDoesNotKnow()
	friendsWithFullKnows := CreateDoesKnow(friendsList)
	friendsWithDoesNotKnowRemoved := FilterDoesKnow(friendsWithFullKnows)
	sortedFriends := SortFriends(friendsWithDoesNotKnowRemoved)
	friendsWithGifts := AssignGiving(sortedFriends)
	for _, friend := range friendsWithGifts {
		println(friend.Name + " gives to " + friend.GivesTo.Name)
	}
}