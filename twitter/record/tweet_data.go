package record

import (
	"github.com/dghubble/go-twitter/twitter"
	"strconv"
	"strings"
)

type TweetData struct {
	// Poster Data
	UserID     int64
	UserSince  string
	Followers  int
	Friends    int
	TweetCount int
	// Tweet Data
	Time     string
	Location [2]float64
	Links    int
	Words    []string
}

func ParseTweet(tweet *twitter.Tweet) (TweetData, bool) {
	td := TweetData{}
	if tweet.RetweetedStatus != nil {
		return td, false
	}
	if tweet.Coordinates == nil {
		return td, false
	}
	if tweet.User == nil {
		return td, false
	}
	td.UserID = tweet.User.ID
	td.UserSince = tweet.User.CreatedAt
	td.Followers = tweet.User.FollowersCount
	td.Friends = tweet.User.FriendsCount
	td.TweetCount = tweet.User.StatusesCount
	//
	td.Time = tweet.CreatedAt
	td.Location = tweet.Coordinates.Coordinates
	//
	fields := strings.Fields(tweet.Text)
	td.Words = make([]string, 0, len(fields))
	for _, f := range fields {
		if strings.HasPrefix(f, "http://") || strings.HasPrefix(f, "https://") {
			td.Links += 1
			continue
		}
		for _, w := range strings.Split(f, ",") {
			if w != "" {
				td.Words = append(td.Words, w)
			}
		}
	}
	return td, true
}

// AggregateData is channel-sequenced
func (td *TweetData) ToCSV() string {
	return strings.Join([]string{
		"\n__TWEET__",
		strconv.FormatInt(td.UserID, 10),
		td.UserSince,
		strconv.Itoa(td.Followers),
		strconv.Itoa(td.Friends),
		strconv.Itoa(td.TweetCount),
		td.Time,
		strconv.FormatFloat(td.Location[0], 'E', -1, 64),
		strconv.FormatFloat(td.Location[1], 'E', -1, 64),
		strconv.Itoa(td.Links),
		strings.Join(td.Words, " "),
	}, ",") + "\n"
}

func FromCSV(line string) (TweetData, bool) {
	if !strings.HasPrefix(line, "__TWEET__") {
		return TweetData{}, false
	}
	td := TweetData{}
	fields := strings.Split(line, ",")
	if len(fields) != 11 || fields[0] != "__TWEET__" {
		return TweetData{}, false
	}
	var err error
	if td.UserID, err = strconv.ParseInt(fields[1], 10, 64); err != nil {
		return TweetData{}, false
	}
	td.UserSince = fields[2]
	if td.Followers, err = strconv.Atoi(fields[3]); err != nil {
		return TweetData{}, false
	}
	if td.Friends, err = strconv.Atoi(fields[4]); err != nil {
		return TweetData{}, false
	}
	if td.TweetCount, err = strconv.Atoi(fields[5]); err != nil {
		return TweetData{}, false
	}
	td.Time = fields[6]
	if td.Location[0], err = strconv.ParseFloat(fields[7], 64); err != nil {
		return TweetData{}, false
	}
	if td.Location[1], err = strconv.ParseFloat(fields[8], 64); err != nil {
		return TweetData{}, false
	}
	if td.Links, err = strconv.Atoi(fields[9]); err != nil {
		return TweetData{}, false
	}
	td.Words = strings.Split(fields[10], " ")
	return td, true
}
