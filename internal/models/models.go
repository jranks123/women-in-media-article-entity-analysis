package models

import (
	"database/sql"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

type ContentFields struct {
	Headline string `json:"headline"`
	Byline   string `json:"byline"`
	BodyText string `json:"bodyText"`
}

type GenderAnalysis struct {
	People []struct {
		Text        string `json:"text"`
		Normal      string `json:"normal	"`
		FirstName   string `json:"firstName"`
		MiddleName  string `json:"middleName"`
		NickName    string `json:"nickName"`
		LastName    string `json:"lastName"`
		GenderGuess string `json:"genderGuess"`
		Pronoun     string `json:"pronoun"`
	} `json:"people"`
	Names map[string]string `json:"names"`
}
type Content struct {
	WebPublicationDate string        `json:"webPublicationDate"`
	Url                string        `json:"url"`
	Section            string        `json:"sectionId"`
	Fields             ContentFields `json:"fields"`
	Id                 string        `json:"id"`
}

type CapiSearchResponse struct {
	Status  string
	Results []Content
}

type Gender string

type Byline struct {
	Name   string `json:"name"`
	Gender Gender `json:"gender"`
}

type Person struct {
	EntityWithNextWord
	Gender Gender
}

type EntityWithNextWord struct {
	Entity *comprehend.Entity;
	NextWord string
}

type ContentAnalysis struct {
	Path               string               `json:"path"`
	Headline           string               `json:"headline"`
	BodyText           string               `json:"bodyText"`
	Bylines            []*Byline            `json:"bylines"`
	People             []*Person            `json:"people"`
	Locations          []*comprehend.Entity `json:"locations"`
	Organisations      []*comprehend.Entity `json:"organisations"`
	CreativeWorkTitles []*comprehend.Entity `json:"creativeWorkTitles"`
	CommercialItems    []*comprehend.Entity `json:"commercialItems"`
	Events             []*comprehend.Entity `json:"events"`
	CacheHit           bool                 `json:"cacheHit"`
	WebPublicationDate string               `json:"webPublicationDate"`
	Section            string               `json:"section"`
	Id                 string               `json:"id"`
}



/* Query structs */

type EntityResult struct {
	Name string
	Gender sql.NullString
	NextWord sql.NullString
	Score float64
	ArticleId string
}

type GenderCountResult struct {
	Entities []*EntityResult
	NumberOfWomen int
	NumberofMen int
	NumberOfGenderless int
	NumberOfFemaleJournalists int
	NumberOfMaleJournalists int
	NumberOfArticles int
}
