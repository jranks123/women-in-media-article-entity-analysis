package utils



func ConstructContentAnalysisQuery(queryCondition string) string {
	return "SELECT article.id, published, content, canonical_url, headline, name, section " +
			"FROM article " +
			"left join author " +
			"on article.id  = author.id " +
		   queryCondition + " ORDER BY published::date ASC"
}


func ConstructEntitiesQuery(queryCondition string) string {
	return "SELECT text, gender, nextWord, score, article_id" +
		" FROM article article" +
		" LEFT join article_entities ae" +
		" ON article.id = ae.article_id " +
		" LEFT join names n " +
		" ON ae.text = n.name " + queryCondition + "AND text is not null"
}