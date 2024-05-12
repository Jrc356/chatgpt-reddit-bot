package reddit

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

const (
	redditBaseURL = "https://www.reddit.com"
)

type Post struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"selftext"`
}

type RedditResponse struct {
	Data struct {
		Children []struct {
			Data Post `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type NewPostsParams struct {
	Subreddit string
	N         int
	Random    bool
}

func GetNewPosts(params NewPostsParams) ([]Post, error) {
	url := fmt.Sprintf("%s/r/%s/new.json", redditBaseURL, params.Subreddit)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var redditResponse RedditResponse
	if err := json.NewDecoder(resp.Body).Decode(&redditResponse); err != nil {
		return nil, err
	}

	var allPosts, filteredPosts []Post
	for _, p := range redditResponse.Data.Children {
		allPosts = append(allPosts, p.Data)
	}
	for len(filteredPosts) < params.N {
		if params.Random {
			filteredPosts = append(filteredPosts, allPosts[rand.Intn(len(allPosts))])
		} else {
			filteredPosts = append(filteredPosts, allPosts[len(filteredPosts)])
		}
	}
	return filteredPosts, nil
}
