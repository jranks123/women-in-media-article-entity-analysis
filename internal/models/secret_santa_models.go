package models


type Friend struct {
	Id          int
	Name        string
	Email 	    string
	DoesNotKnow []*Friend
	Knows       []*Friend
	GivesTo 	*Friend
}

type FriendList struct {
	FriendList []FriendV2 `json:"friendList"`
}
type FriendV2 struct {
	Id          int  `json:"id"`
	Name        string `json:"name"`
	Email 	    string `json:"email"`
	Knows       []string `json:"knows"`
	GivesTo 	*string `json:"givesTo"`
}
