package services

import (
	"fmt"
	"testing"
	"women-in-media-article-entity-analysis/internal/models"
)

func TestGetEntitiesForArticle(t *testing.T) {
	content := models.Content{
		WebPublicationDate: "2019-06-17T04:30:09Z",
		Section:            "environment",
		Fields: models.ContentFields{
			Headline: "Country diary: the biting hordes of clegs are late this year",
			Byline:   "Lucy Hughes-Hallett",
			BodyText: "\u003cdiv\u003e\u003cp\u003eShortly before his death, the independent journalist Jagendra Singh wrote a social media post about the threats he believed he faced for investigating sand mining in northern \u003ca href=\"https://www.theguardian.com/world/india\"\u003eIndia\u003c/a\u003e.\u003c/p\u003e\n\u003cp\u003e“Politicians, thugs, and police, all are after me. Writing the truth is weighing heavily on my life.”\u003c/p\u003e\n\u003cp\u003eTwo weeks later, on 1 June 2015, the veteran reporter was set upon by a gang that entered his home, doused him with petrol, and set him alight.\u003c/p\u003e\n\u003cp\u003eHe later died from his injuries.\u003c/p\u003e\n\u003cp\u003eSingh had been investigating land grabs and alleged illegal extraction of sand from the Garra River. \u003c/p\u003e\n\u003cp\u003eSand is mined both legally and illegally, to extract minerals and for construction and reclamation projects.\u003c/p\u003e\n\u003cp\u003eThis multibillion-dollar business – which supplies up to 50bn tonnes to the construction industry each year – is prohibited in an increasing number of states around the world as erosion, floods and other environmental consequences become harder to ignore.\u003c/p\u003e\n\u003cp\u003eSingh’s son said his father had previously been threatened. Police concluded he had killed himself. The family said they were given 3m rupees (£34,000) as a “wedding gift” for Singh’s daughter. They took it to mean that they should accept the police verdict.\u003c/p\u003e\n\u003cfigure\u003e\u003ca href=\"#img-2\"\u003e\u003cspan\u003e\u003c/span\u003e\u003c/a\u003e\u003cfigcaption\u003e\u003cspan\u003e\u003c/span\u003e Diksha Singh. Photograph: Forbidden Stories\n\u003c/figcaption\u003e\u003c/figure\u003e\u003cp\u003eSingh’s daughter, Diksha, has refused to touch the money. She said her father was fighting for justice. “My father was one such rare people who exposed the truth.”\u003c/p\u003e\n\u003cp\u003eSince Singh, two other journalists have been killed while investigating sand mining in India. Karun Misra died in February 2016. Sandeep Sharma was killed in March 2018.\u003c/p\u003e\n\u003c/div\u003e\n\n",
		},
		Id: "4137",
	}

	res, err := GetEntitiesForArticle(content)
	if err != nil {
		t.Error(err)
	} else {
		for _, entity := range res {
			if *entity.Entity.Type == "PERSON" && *entity.Entity.Score > 0.90 {
				fmt.Println(entity.Entity.GoString())
			}
		}
	}
}

func TestGetNextWordAfterEntities(t *testing.T) {
	copy := "hello Jonny Rankin happy Ben sad then my mate Benji ran"
	res, err := GetEntitiesFromBodyText(copy)
	if err != nil {
		t.Error(err)
	} else {
		nextWords := GetNextWordAfterEntities(res, copy)
		for _, entityWithNextWord := range nextWords {
			fmt.Println(entityWithNextWord.NextWord)
		}
	}
}

func TestGetEntitiesFromBodyText(t *testing.T) {
	res, err := GetEntitiesFromBodyText("hello Jonny happy")
	if err != nil {
		t.Error(err)
	} else {
		for _, entity := range res {
			fmt.Println(entity.GoString())
		}
	}
}
