package utils

import "women-in-media-article-entity-analysis/internal/models"

func EntityPassesConfidenceChecks(name string, score float64) bool {
	return score > 0.9 && WordCount(name) > 1
}

func EntityCounts(entities []*models.EntityResult) (int, int, int, int){
	articleCount := make( map[string]int )
	menCount := 0
	womenCount := 0
	noGenderCount := 0
	for _, entity := range entities {
		articleCount[entity.ArticleId]++

		if entity.Gender.String == "Male" {
			menCount++
		} else if entity.Gender.String == "Female" {
			womenCount++
		} else {
			noGenderCount++
		}


	}
	return len(articleCount), menCount, womenCount, noGenderCount
}