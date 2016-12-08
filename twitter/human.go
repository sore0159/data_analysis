package twitter

import (
	"math"
)

func Human(td TweetData) bool {
	days := int64(math.Floor(td.TweetDate.Sub(td.UserSinceDate).Hours() / 24.0))
	if days < 30 {
		return false // human, maybe, but not enough time for style to affect followers?
	}
	if days < 10 {
		if td.TweetCount > 300 {
			return false
		}
	} else if int64(td.TweetCount)/days > 10 {
		return false
	}
	return true
}
