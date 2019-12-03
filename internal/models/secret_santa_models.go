package models

type FriendList struct {
	FriendList []Friend `json:"friendList"`
}
type Friend struct {
	Id          int  `json:"id"`
	Name        string `json:"name"`
	Email 	    string `json:"email"`
	Knows       []string `json:"knows"`
	GivesTo 	*string `json:"givesTo"`
}
