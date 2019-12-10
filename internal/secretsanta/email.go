package secretsanta

import (
	"log"
	"net/smtp"
)

func SendEmail(body string, emailTo string) {
	from := "jonnyrankin@googlemail.com"
	pass := ""
	to := emailTo

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: [SECRET MESSAGE DO NOT SHARE DELETE ON READ FUCKKKKK]\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

func SendEmails(fileLocation string) {
	friendList := ReadJsonFile(fileLocation)

	for _, friend := range friendList.FriendList {
		emailCopy := "Good morning  " + friend.Name + ". This is secret santa. Before continuing, please make sure no one else can see your screen" +
			" \n \n \n \n \n ... \n \n \n \n \n \n \n \n ...  \n \n \n \n SCROLL DOWN TO SEE FULL MESSAGE \n \n \n \n ...  \n \n \n \n \n \n \n \n ... \n \n \n \n " +
			"Excellent. I hope you've been a good human this year. Even if you haven't, there is someone who has. " +
			"They deserve a good present. A real good present. That person is..." +
			" \n \n \n \n \n ... \n \n \n \n \n \n \n \n ...  \n \n \n \n SCROLL DOWN TO SEE FULL MESSAGE \n \n \n \n ...  \n \n \n \n \n \n \n \n ... \n \n \n \n " +
			"\n" + *friend.GivesTo + ". Yes, you read that right, it's " + *friend.GivesTo + "!!! They've been bloody good this year and they need the best goddamn present they've ever had. You have 11 days. Make them count.\n \n Loves and kisses, \n S \n\n\n " +
			"(PS: although this email has come from Jonny's account, he hasn't seen who you've got. NOONE HAS BUT YOU. keep it secret. keep it safe."

		SendEmail(emailCopy, friend.Email)
	}

}