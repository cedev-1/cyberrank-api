package models

type Profile struct {
    Username   string             `json:"username"`
    Rank       string             `json:"rank"`
    Platform   string             `json:"platform"`
    Categories []CategoryProgress `json:"categories,omitempty"`
    Overall    *OverallProgress   `json:"overall,omitempty"`
}

type CategoryProgress struct {
    Name       string `json:"name"`
    Percentage int    `json:"percentage"`
}

type OverallProgress struct {
    Percentage int    `json:"percentage"`
    Solved     int    `json:"solved"`
    Total      int    `json:"total"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

type DetailedProfile struct {
	Username           string `json:"username"`
	Rank               string `json:"rank"`
	Platform           string `json:"platform"`
	ID                 string `json:"id"`
	Avatar             string `json:"avatar"`
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
	CompletedRoomsNumber int `json:"completedRoomsNumber"`
	Streak             int    `json:"streak"`
	TopPercentage      int    `json:"topPercentage"`
	IsInTopTenPercent  bool   `json:"isInTopTenPercent"`
	BadgeImageURL      string `json:"badgeImageURL"`
}
