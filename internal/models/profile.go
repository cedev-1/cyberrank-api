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
