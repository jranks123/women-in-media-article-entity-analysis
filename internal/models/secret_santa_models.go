package models


type FriendList struct {
	Friends []Friend
}

type Friend struct {
	Id          int
	Name        string
	Email 	    string
	DoesNotKnow []*Friend
	Knows       []*Friend
	GivesTo 	*Friend
}
