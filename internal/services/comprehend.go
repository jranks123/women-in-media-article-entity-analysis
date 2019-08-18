package services

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/pkg/errors"
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

func GetEntitiesForArticle(article models.Content) ([]*comprehend.Entity, error) {

	var bodyTextArray = utils.SplitSubN(article.Fields.BodyText, 3000)

	var allEntities []*comprehend.Entity

	// hack to stop it failing on long articles
	for _, bodyTextSegment := range bodyTextArray {
		entities, err := GetEntitiesFromBodyText(bodyTextSegment)
		if err != nil {
			return nil, errors.Wrap(err, "Error retrieving entities from body text")
		}

		for _, entity := range entities {
			allEntities = append(allEntities, entity)
		}
	}

	return allEntities, nil
}
