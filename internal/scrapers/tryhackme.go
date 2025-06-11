package scrapers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cedev-1/cyberrank-api/internal/models"
	"github.com/cedev-1/cyberrank-api/pkg/httpclient"
)

type TryHackMeResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID                 string `json:"_id"`
		IDNum              int    `json:"id"`
		Avatar             string `json:"avatar"`
		Username           string `json:"username"`
		Level              int    `json:"level"`
		Country            string `json:"country"`
		About              string `json:"about"`
		LinkedInUsername   string `json:"linkedInUsername"`
		GithubUsername     string `json:"githubUsername"`
		TwitterUsername    string `json:"twitterUsername"`
		InstagramUsername  string `json:"instagramUsername"`
		PersonalWebsite    string `json:"personalWebsite"`
		Subscribed         int    `json:"subscribed"`
		BadgesNumber       int    `json:"badgesNumber"`
		DateSignUp         string `json:"dateSignUp"`
		CompletedRoomsNumber int    `json:"completedRoomsNumber"`
		Streak             int    `json:"streak"`
		Rank               int    `json:"rank"`
		TopPercentage      int    `json:"topPercentage"`
		IsInTopTenPercent  bool   `json:"isInTopTenPercent"`
		BadgeImageURL      string `json:"badgeImageURL"`
	} `json:"data"`
}

func GetTryHackMeRank(username string, detailed bool) (interface{}, error) {
	url := fmt.Sprintf("https://tryhackme.com/api/v2/public-profile?username=%s", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", httpclient.GetUserAgent())

	resp, err := httpclient.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("user not found")
	}

	var thm TryHackMeResponse
	if err := json.NewDecoder(resp.Body).Decode(&thm); err != nil {
		return nil, err
	}

	if thm.Status != "success" {
		return nil, fmt.Errorf("user not found")
	}

	if detailed {
		return models.DetailedProfile{
			Username:           thm.Data.Username,
			Rank:               fmt.Sprintf("%d", thm.Data.Rank),
			Platform:           "tryhackme",
			ID:                 thm.Data.ID,
			Avatar:             thm.Data.Avatar,
			Level:              thm.Data.Level,
			Country:            thm.Data.Country,
			About:              thm.Data.About,
			LinkedInUsername:   thm.Data.LinkedInUsername,
			GithubUsername:     thm.Data.GithubUsername,
			TwitterUsername:    thm.Data.TwitterUsername,
			InstagramUsername:  thm.Data.InstagramUsername,
			PersonalWebsite:    thm.Data.PersonalWebsite,
			Subscribed:         thm.Data.Subscribed,
			BadgesNumber:       thm.Data.BadgesNumber,
			DateSignUp:         thm.Data.DateSignUp,
			CompletedRoomsNumber: thm.Data.CompletedRoomsNumber,
			Streak:             thm.Data.Streak,
			TopPercentage:      thm.Data.TopPercentage,
			IsInTopTenPercent:  thm.Data.IsInTopTenPercent,
			BadgeImageURL:      thm.Data.BadgeImageURL,
		}, nil
	}

	return models.Profile{
		Username: thm.Data.Username,
		Rank:     fmt.Sprintf("%d", thm.Data.Rank),
		Platform: "tryhackme",
	}, nil
}
