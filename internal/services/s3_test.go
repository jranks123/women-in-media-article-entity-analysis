package services_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"testing"
	"women-in-media-article-entity-analysis/internal"
	"women-in-media-article-entity-analysis/internal/models"
	"women-in-media-article-entity-analysis/internal/services"
)

func TestGetContentAnalysisFromS3(t *testing.T) {
	res, err := services.GetContentAnalysisFromS3("/commentisfree/2019/apr/08/wall-street-socialism-jpmorgan-jamie-dimon-bailout")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res.Path)
	}
}

func TestStoreContentAnalysisInS3(t *testing.T) {
	contentFields := models.ContentFields{"test_headline", "test_byline", "test_body"}
	content := models.Content{"2019-04-11T13:35:13Z", "https://phescreening.blog.gov.uk/2019/05/21/make-sure-your-leaflets-do-not-get-lost-in-the-post/", "football", contentFields, "test"}
	var events []*comprehend.Entity = nil
	contentAnalysis := internal.ConstructContentAnalysis(
		content,
		events,
		false,
	)

	err := services.StoreContentAnalysisInS3(contentAnalysis)
	if err != nil {
		t.Error(err)
	}
}
