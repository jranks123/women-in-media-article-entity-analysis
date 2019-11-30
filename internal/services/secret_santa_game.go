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
	Jessie := models.Friend{2, "Jessie", "", []*models.Friend{}, []*models.Friend{}, nil}
	Jonny := models.Friend{4, "Jonny", "", []*models.Friend{}, []*models.Friend{}, nil}
	Kitty := models.Friend{5, "Kitty", "", []*models.Friend{}, []*models.Friend{}, nil}
	Oscar := models.Friend{6, "Oscar", "", []*models.Friend{}, []*models.Friend{}, nil}
	Tom := models.Friend{7, "Tom", "", []*models.Friend{}, []*models.Friend{}, nil}
	Oriel := models.Friend{8, "Oriel", "", []*models.Friend{}, []*models.Friend{}, nil}
	Max := models.Friend{9, "Max", "", []*models.Friend{}, []*models.Friend{}, nil}
	Doug := models.Friend{10, "Doug", "", []*models.Friend{}, []*models.Friend{}, nil}
	Esther := models.Friend{11, "Esther", "", []*models.Friend{}, []*models.Friend{}, nil}
	Freddie := models.Friend{12, "Freddie", "", []*models.Friend{}, []*models.Friend{}, nil}
	Hannah := models.Friend{13, "Hannah", "", []*models.Friend{}, []*models.Friend{}, nil}
	Giorgi := models.Friend{14, "Giorgi", "", []*models.Friend{}, []*models.Friend{}, nil}
	Xave := models.Friend{15, "Xave", "", []*models.Friend{}, []*models.Friend{}, nil}
	JakeM := models.Friend{16, "JakeM", "", []*models.Friend{}, []*models.Friend{}, nil}
	Jen := models.Friend{17, "Jen", "", []*models.Friend{}, []*models.Friend{}, nil}
	Hen := models.Friend{18, "Hen", "", []*models.Friend{}, []*models.Friend{}, nil}
	Josh := models.Friend{19, "Josh", "", []*models.Friend{}, []*models.Friend{}, nil}
	Misha := models.Friend{20, "Misha", "", []*models.Friend{}, []*models.Friend{}, nil}
	Rafa := models.Friend{21, "Rafa", "", []*models.Friend{}, []*models.Friend{}, nil}
	Juby := models.Friend{22, "Juby", "", []*models.Friend{}, []*models.Friend{}, nil}
	George := models.Friend{23, "George", "", []*models.Friend{}, []*models.Friend{}, nil}
	Oli := models.Friend{24, "Oli", "", []*models.Friend{}, []*models.Friend{}, nil}
	Micha := models.Friend{3, "Micha", "", []*models.Friend{}, []*models.Friend{}, nil}
	JoePS := models.Friend{26, "JoePS", "", []*models.Friend{}, []*models.Friend{}, nil}
	Martha := models.Friend{27, "Martha", "", []*models.Friend{}, []*models.Friend{}, nil}
	Amy := models.Friend{28, "Amy", "", []*models.Friend{}, []*models.Friend{}, nil}
	Dylan := models.Friend{29, "Dylan", "", []*models.Friend{}, []*models.Friend{}, nil}


	// JAKE G
	JakeG.DoesNotKnow = append(JakeG.DoesNotKnow, &George)
	George.DoesNotKnow = append(George.DoesNotKnow, &JakeG)

	// ORIEL
	Oriel.DoesNotKnow = append(Oriel.DoesNotKnow, &George)
	George.DoesNotKnow = append(George.DoesNotKnow, &Oriel)

	Oriel.DoesNotKnow = append(Oriel.DoesNotKnow, &Micha)
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Oriel)

	Oriel.DoesNotKnow = append(Oriel.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Oriel)

	Oriel.DoesNotKnow = append(Oriel.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Oriel)


	//DOUG
	Doug.DoesNotKnow = append(Doug.DoesNotKnow, &Jen)
	Jen.DoesNotKnow = append(Jen.DoesNotKnow, &Doug)

	Doug.DoesNotKnow = append(Doug.DoesNotKnow, &Hen)
	Hen.DoesNotKnow = append(Hen.DoesNotKnow, &Doug)

	Doug.DoesNotKnow = append(Doug.DoesNotKnow, &George)
	George.DoesNotKnow = append(George.DoesNotKnow, &Doug)

	Doug.DoesNotKnow = append(Doug.DoesNotKnow, &Oli)
	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Doug)

	Doug.DoesNotKnow = append(Doug.DoesNotKnow, &Micha)
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Doug)

	Doug.DoesNotKnow = append(Doug.DoesNotKnow, &JoePS)
	JoePS.DoesNotKnow = append(JoePS.DoesNotKnow, &Doug)

	Doug.DoesNotKnow = append(Doug.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Doug)

	Doug.DoesNotKnow = append(Doug.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Doug)

	// MAX

	Max.DoesNotKnow = append(Max.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Max)

	Max.DoesNotKnow = append(Max.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Max)


	// ESTHER

	Esther.DoesNotKnow = append(Esther.DoesNotKnow, &Oli)
	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Esther)

	Esther.DoesNotKnow = append(Esther.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Esther)

	Esther.DoesNotKnow = append(Esther.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Esther)

	Esther.DoesNotKnow = append(Esther.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Esther)

	// FREDDIE

	Freddie.DoesNotKnow = append(Freddie.DoesNotKnow, &Micha)
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Freddie)

	Freddie.DoesNotKnow = append(Freddie.DoesNotKnow, &Micha)
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Freddie)

	Freddie.DoesNotKnow = append(Freddie.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Freddie)

	Freddie.DoesNotKnow = append(Freddie.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Freddie)

	// HANNAH

	Hannah.DoesNotKnow = append(Hannah.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Hannah)

	Hannah.DoesNotKnow = append(Hannah.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Hannah)

	Hannah.DoesNotKnow = append(Hannah.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Hannah)

	// GEORGIE

	Giorgi.DoesNotKnow = append(Giorgi.DoesNotKnow, &Oli)
	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Giorgi)

	Giorgi.DoesNotKnow = append(Giorgi.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Giorgi)

	Giorgi.DoesNotKnow = append(Giorgi.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Giorgi)


	// XAVE

	Xave.DoesNotKnow = append(Xave.DoesNotKnow, &Oli)
	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Xave)

	Xave.DoesNotKnow = append(Xave.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Xave)

	Xave.DoesNotKnow = append(Xave.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Xave)


	// Jen

	Jen.DoesNotKnow = append(Jen.DoesNotKnow, &Oli)
	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Jen)

	Jen.DoesNotKnow = append(Jen.DoesNotKnow, &Micha)
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Jen)

	Jen.DoesNotKnow = append(Jen.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Jen)

	Jen.DoesNotKnow = append(Jen.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Jen)

	Jen.DoesNotKnow = append(Jen.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Jen)

	// Hen

	Hen.DoesNotKnow = append(Hen.DoesNotKnow, &Oli)
	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Hen)

	Hen.DoesNotKnow = append(Hen.DoesNotKnow, &Micha)
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Hen)

	Hen.DoesNotKnow = append(Hen.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Hen)

	Hen.DoesNotKnow = append(Hen.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Hen)

	Hen.DoesNotKnow = append(Hen.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Hen)

	// Rafa

	Rafa.DoesNotKnow = append(Rafa.DoesNotKnow, &Oli)
	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Rafa)

	Rafa.DoesNotKnow = append(Rafa.DoesNotKnow, &Micha)
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Rafa)

	Rafa.DoesNotKnow = append(Rafa.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Rafa)

	Rafa.DoesNotKnow = append(Rafa.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Rafa)

	Rafa.DoesNotKnow = append(Rafa.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Rafa)

	// Juby

	Juby.DoesNotKnow = append(Juby.DoesNotKnow, &Oli)
	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Juby)

	Juby.DoesNotKnow = append(Juby.DoesNotKnow, &Micha)
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Juby)

	Juby.DoesNotKnow = append(Juby.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Juby)

	Juby.DoesNotKnow = append(Juby.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Juby)

	Juby.DoesNotKnow = append(Juby.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Juby)



	// George

	George.DoesNotKnow = append(George.DoesNotKnow, &Oli)
	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &George)

	George.DoesNotKnow = append(George.DoesNotKnow, &Micha)
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &George)

	George.DoesNotKnow = append(George.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &George)

	George.DoesNotKnow = append(George.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &George)

	George.DoesNotKnow = append(George.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &George)

	// Jake M

	JakeM.DoesNotKnow = append(JakeM.DoesNotKnow, &Oli)
	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &JakeM)

	JakeM.DoesNotKnow = append(JakeM.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &JakeM)

	JakeM.DoesNotKnow = append(JakeM.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &JakeM)

	JakeM.DoesNotKnow = append(JakeM.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &JakeM)



	// Oli

	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Micha)
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Oli)

	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &JoePS)
	JoePS.DoesNotKnow = append(JoePS.DoesNotKnow, &Oli)

	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Oli)

	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Oli)

	Oli.DoesNotKnow = append(Oli.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Oli)


	// Micha

	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &JoePS)
	JoePS.DoesNotKnow = append(JoePS.DoesNotKnow, &Micha)

	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Micha)

	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Micha)

	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Micha)


	// Joe PS

	JoePS.DoesNotKnow = append(JoePS.DoesNotKnow, &Martha)
	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &JoePS)

	JoePS.DoesNotKnow = append(JoePS.DoesNotKnow, &Amy)
	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &JoePS)

	JoePS.DoesNotKnow = append(JoePS.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &JoePS)

	// Martha

	Martha.DoesNotKnow = append(Martha.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Martha)

	// Amy

	Amy.DoesNotKnow = append(Amy.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Amy)













	// MICHAZ
	Micha.DoesNotKnow = append(Micha.DoesNotKnow, &Dylan)
	Dylan.DoesNotKnow = append(Dylan.DoesNotKnow, &Micha)


	friendsList = append(friendsList,JakeG )
	friendsList = append(friendsList,Jessie )
	friendsList = append(friendsList,Jonny )
	friendsList = append(friendsList,Kitty )
	friendsList = append(friendsList,Oscar )
	friendsList = append(friendsList,Tom )
	friendsList = append(friendsList,Oriel )
	friendsList = append(friendsList,Max )
	friendsList = append(friendsList,Doug )
	friendsList = append(friendsList,Esther )
	friendsList = append(friendsList,Freddie )
	friendsList = append(friendsList,Hannah )
	friendsList = append(friendsList,Giorgi )
	friendsList = append(friendsList,Xave )
	friendsList = append(friendsList,JakeM )
	friendsList = append(friendsList,Jen )
	friendsList = append(friendsList,Hen )
	friendsList = append(friendsList,Josh )
	friendsList = append(friendsList,Misha )
	friendsList = append(friendsList,Rafa )
	friendsList = append(friendsList,Juby )
	friendsList = append(friendsList,George )
	friendsList = append(friendsList,Oli )
	friendsList = append(friendsList,Micha )
	friendsList = append(friendsList,JoePS )
	friendsList = append(friendsList,Martha )
	friendsList = append(friendsList,Amy )
	friendsList = append(friendsList,Dylan )

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
	rand.Seed(time.Now().UnixNano())
	return *friends[rand.Intn(len(friends))]
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
		println(friendsList[i].Name)
		// set gives to
		for friendsList[i].GivesTo == nil {
			potentialFriendToGiveTo := PickRandomFriend(friendsList[i].Knows)
			if !FriendIsAlreadyReceivingGift(potentialFriendToGiveTo, friendsList) &&  friendsList[i].Name != potentialFriendToGiveTo.Name {
				friendsList[i].GivesTo = &potentialFriendToGiveTo
				if friendsList[i].Name == potentialFriendToGiveTo.Name {
					println("duplicate giving for " + friendsList[i].GivesTo.Name)
				}
			}
		}

		// set receives from
		for FriendIsAlreadyReceivingGift(friendsList[i], friendsList) == false {
			potentialFriendToRecieveFrom := PickRandomFriend(friendsList[i].Knows)
			if potentialFriendToRecieveFrom.GivesTo == nil && friendsList[i].Name != potentialFriendToRecieveFrom.Name{
				friendsList = SetGivesToForPersonrReceivingFrom(potentialFriendToRecieveFrom, friendsList[i], friendsList)
				if friendsList[i].GivesTo.Name == potentialFriendToRecieveFrom.Name {
					println("duplicate receiving for " + friendsList[i].GivesTo.Name)
				}
			}
		}
	}

	return friendsList
}