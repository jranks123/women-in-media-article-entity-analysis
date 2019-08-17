package internal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetContentAnalysis(t *testing.T) {
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article left join author on article.id  = author.id WHERE article.canonical_url in ('https://www.theguardian.com/world/2019/jun/03/canada-endangered-killer-whales-orcas-calf', 'https://www.theguardian.com/work-smarter/2019/may/29/from-flying-taxis-to-ai-apps-tech-trends-that-show-the-future-of-business-travel', 'https://www.theguardian.com/film/2019/may/09/emma-thompsons-best-films-ranked', 'https://www.theguardian.com/fashion/2019/may/29/rory-stewarts-nerdiness-means-he-is-just-the-sort-of-oddball-the-british-love', 'https://www.theguardian.com/sport/rugby-union', 'https://www.theguardian.com/cities/2019/jun/03/hidden-tokyo-share-your-videos-of-your-favourite-quiet-place-', 'https://www.theguardian.com/lifeandstyle/2019/apr/21/playful-yet-powerful-time-to-think-pink-pat-mcgrath-eye-shadow', 'https://www.theguardian.com/sport/tennis', 'https://www.theguardian.com/books/2019/may/20/no-win-race-a-story-of-belonging-britishness-and-sport-derek-bardowell-review', 'https://www.theguardian.com/stage/2019/jun/02/king-hedley-ii-review-rutherford-and-son-the-starry-messenger', 'https://www.theguardian.com/guardian-masterclasses/guardian-masterclass-blog/2019/feb/05/i-pride-myself-on-empathy-and-honesty-tim-lotts-mentoring-for-writers-blog', 'https://www.theguardian.com/food/2019/may/31/fiona-beckett-rose-drink-summer-wine', 'https://www.theguardian.com/education/2019/may/31/why-are-students-at-university-so-stressed', 'https://www.theguardian.com/sport/99-94-cricket-blog/2019/may/31/county-cricket-talking-points-hampshire-yorkshire-somerset-alastair-cook', 'https://www.theguardian.com/uk-news/2019/jun/03/man-carrying-knives-at-gatwick-airport-shot-with-taser-say-police', 'https://www.theguardian.com/commentisfree/2019/jun/03/love-island-itv-contestants-social-media', 'https://www.theguardian.com/stage/gallery/2015/mar/30/arthur-miller-death-of-a-salesman-in-pictures', 'https://www.theguardian.com/commentisfree/video/2017/may/17/macron-was-a-victory-for-hope-that-makes-the-price-of-failure-even-higher-video', 'https://www.theguardian.com/books/2019/jun/03/young-people-are-full-of-rage-and-terror-and-that-gives-us-power-meet-the-activists', 'https://www.theguardian.com/education/2019/may/28/fear-lgbt-inclusive-lessons-harks-back-to-80s-peter-tatchell', 'https://www.theguardian.com/sport/golf', 'https://www.theguardian.com/artanddesign/gallery/2019/may/27/myanmars-ruby-gems-mining-in-pictures', 'https://www.theguardian.com/games/2019/apr/29/why-i-love-notebooks-in-video-games-red-dead-redemption-2-discworld-noir', 'https://www.theguardian.com/science/2019/may/31/sexist-research-means-drugs-more-tailored-to-men-says-scientist', 'https://www.theguardian.com/artanddesign/2019/may/11/observer-archive-students-demonstrate-in-paris-10-may-1968', 'https://www.theguardian.com/global/video/2019/jan/25/owen-jones-meets-tim-martin-wetherspoons-no-deal-brexit-poverty-wages-dont-ask-childish-questions-video', 'https://www.theguardian.com/travel/2019/may/03/locanda-on-weir-porlock-weir-somerset-restaurant-bed-breakfast-review', 'https://www.theguardian.com/tv-and-radio/2019/jun/03/best-tv-of-2019-so-far', 'https://www.theguardian.com/culture/2019/may/06/dont-go-near-the-hog-roast-loyle-carner-festival-food-tips', 'https://www.theguardian.com/travel/2019/apr/09/10-best-small-family-friendly-festivals-2019-uk', 'https://www.theguardian.com/lifeandstyle/2017/oct/23/should-i-avoid-drinking-in-front-of-my-children', 'https://www.theguardian.com/technology/2019/apr/24/huawei-p30-pro-review-leica-quad-camera-zoom', 'https://www.theguardian.com/sport/blog/2019/may/29/talking-horses-dettori-on-anapurna-as-obrien-sends-four-to-oaks-derby-news-latest', 'https://www.theguardian.com/music/2019/may/26/flamagra-flying-lotus-album-review-damien-morris', 'https://www.theguardian.com/tv-and-radio/2019/may/28/the-planets-review-bbc-two-brian-cox-solar-system', 'https://www.theguardian.com/world/2019/jun/02/ten-killed-in-israeli-airstrikes-in-syria-in-reply-to-rocket-attack', 'https://www.theguardian.com/guardian-masterclasses/2016/feb/12/how-to-get-your-novel-published-a-class-with-literary-agent-ed-wilson-and-publisher-suzie-doore', 'https://www.theguardian.com/uk-news/2019/jun/03/britains-got-talent-performer-stabbed-court-hears-desmond-sylva-simonne-kerr', 'https://www.theguardian.com/commentisfree/2019/may/27/ken-loach-socially-conscious-cinema-cannes-drama', 'https://www.theguardian.com/commentisfree/series/comment-is-free-weekly', 'https://www.theguardian.com/fashion/2019/apr/30/feminism-marrakech-and-diana-ross-the-second-coming-of-dior', 'https://www.theguardian.com/environment/2019/may/30/renewable-energy-jobs-in-uk-plunge-by-a-third', 'https://www.theguardian.com/guardian-masterclasses/2018/mar/19/how-to-build-your-perfect-pitch-deck-and-win-andy-pemberton-andy-cowles-business-course', 'https://www.theguardian.com/crosswords/crossword-blog/2018/jul/15/old-setters-and-new-conventions', 'https://www.theguardian.com/technology/2019/jan/15/amazon-echo-show-2nd-gen-review-alexa-speaker', 'https://www.theguardian.com/world/gallery/2019/may/09/washington-spy-museum-in-pictures', 'https://www.theguardian.com/sport/live/2019/jun/01/exeter-saracens-premiership-rugby-union-final-2019-live', 'https://www.theguardian.com/food/2019/may/12/nigel-slater-fennel-recipes-vegetarian-soup-ice-cream', 'https://www.theguardian.com/travel/2019/jun/02/greece-greek-islands-10-best-family-friendly-hotels-villas-places-to-stay') ORDER BY published::date ASC"

	res, err := GetContentAnalysis(query)

	if err != nil {
		t.Error(err)
	} else {
		res, err := json.Marshal(res)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(string(res))
		}
	}
}

func TestRedoGenderAnalysis(t *testing.T) {
	query := "SELECT article.id, published, content, canonical_url, headline, name, section FROM article left join author on article.id  = author.id WHERE article.canonical_url in ('https://www.theguardian.com/world/2019/jun/03/canada-endangered-killer-whales-orcas-calf', 'https://www.theguardian.com/work-smarter/2019/may/29/from-flying-taxis-to-ai-apps-tech-trends-that-show-the-future-of-business-travel', 'https://www.theguardian.com/film/2019/may/09/emma-thompsons-best-films-ranked') ORDER BY published::date ASC"

	err := RedoGenderAnalysis(query)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("Successfully redid gender analysis")
	}
}
