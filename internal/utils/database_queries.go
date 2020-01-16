package utils



func ConstructContentAnalysisQuery(queryCondition string) string {
	return "SELECT article.id, published, content, canonical_url, headline, name, section " +
			"FROM article " +
			"left join author " +
			"on article.id  = author.id " +
		   queryCondition + " ORDER BY published::date ASC"
}


func ConstructPrintResultsQuery(queryCondition string) string {
	return "SELECT coalesce(text, ''), gender, nextWord, coalesce(score, 0), article.id" +
		" FROM article article" +
		" LEFT join article_entities ae" +
		" ON article.id = ae.article_id " +
		" LEFT join names n " +
		" ON ae.text = n.name " +
		" LEFT join author " +
		" ON aa.author_id = author.id " +
		" LEFT join names n2 " +
		" ON author.name = n2.name " + queryCondition
}