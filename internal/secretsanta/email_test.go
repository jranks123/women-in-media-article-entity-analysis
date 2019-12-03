package secretsanta

import "testing"

func TestSendEmail(t *testing.T) {

	emailCopy := "Good morning  " + "Jonny" + ". This is secret santa. Before continuing, please make sure no one else can see your screen" +
		" \n \n \n \n \n ... \n \n \n \n \n \n \n \n ...  \n \n \n \n SCROLL DOWN TO SEE FULL MESSAGE \n \n \n \n ...  \n \n \n \n \n \n \n \n ... \n \n \n \n " +
		"Excellent. I hope you've been a good human this year. Even if you haven't, there is someone who has. " +
		"They deserve a good present. A real good present. That person is..." +
		" \n \n \n \n \n ... \n \n \n \n \n \n \n \n ...  \n \n \n \n SCROLL DOWN TO SEE FULL MESSAGE \n \n \n \n ...  \n \n \n \n \n \n \n \n ... \n \n \n \n " +
		"\n" + "Kitty" + ". Yes, you read that right, it's" + "Kitty" + "!!! They've been bloody good this year and they need the best goddamn present they've ever had. You have 11 days. Make them count.\n \n Loves and kisses, \n S \n\n\n " +
		"(PS: although this email has come from Jonny's account, he hasn't seen who you've got. NOONE HAS BUT YOU. keep it secret. keep it safe."

	SendEmail(emailCopy, "jonnyrankin@googlemail.com")
}


func TestSendEmails(t *testing.T) {
	SendEmails("output.json")
}