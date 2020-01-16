package services

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/pkg/errors"
	"strings"
	"women-in-media-article-entity-analysis/internal/models"
	"women-in-media-article-entity-analysis/internal/utils"
)

func GetComprehendClient(profile string) (*comprehend.Comprehend, error) {
	sess, err := GetAwsSession(profile, "us-east-1")

	if err != nil {
		return nil, errors.Wrap(err, "unable to create new sessions")
	}

	return comprehend.New(sess), nil
}

func GetEntitiesFromBodyText(bodyText string) ([]*comprehend.Entity, error) {
	client, err := GetComprehendClient("bechdel")

	if err != nil {
		return nil, errors.Wrap(err, "couldn't create client")
	}

	input := &comprehend.DetectEntitiesInput{}
	input.SetText(bodyText)
	input.SetLanguageCode("en")
	result, err := client.DetectEntities(input)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get entities")
	}
	return result.Entities, nil
}

func GetNextWordAfterEntities(entities []*comprehend.Entity, bodyTextSegment string) [] models.EntityWithNextWord {

	var entitiesWithNextWord []models.EntityWithNextWord
	for _, entity := range entities {
		bodyTextSegmentFromEntity :=  bodyTextSegment[int(*entity.BeginOffset) + len(*entity.Text)-1:]
		wordsFromEntity := strings.Fields(bodyTextSegmentFromEntity)
		if len(wordsFromEntity) >= 2 {
			nextWord := wordsFromEntity[1]
			entityWithNextWord := models.EntityWithNextWord{Entity: entity, NextWord: nextWord}
			entitiesWithNextWord = append(entitiesWithNextWord, entityWithNextWord)
		}
	}
	return entitiesWithNextWord
}

func GetEntitiesForArticle(article models.Content) ([] models.EntityWithNextWord, error) {

	var bodyTextArray = utils.SplitSubN(article.Fields.BodyText, 4000)

	var allEntities [] models.EntityWithNextWord

	// hack to stop it failing on long articles
	for _, bodyTextSegment := range bodyTextArray {
		entities, err := GetEntitiesFromBodyText(bodyTextSegment)
		entitiesWithNextWord := GetNextWordAfterEntities(entities, bodyTextSegment)
		if err != nil {
			return nil, errors.Wrap(err, "Error retrieving entities from body text")
		}

		for _, entity := range entitiesWithNextWord {
			allEntities = append(allEntities, entity)
		}
	}

	return allEntities, nil
}
