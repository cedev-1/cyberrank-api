package scrapers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cedev-1/cyberrank-api/pkg/httpclient"
)

type TryHackMeResponse struct {
	Status string `json:"status"`
	Data   struct {
		Rank int `json:"rank"`
	} `json:"data"`
}

func GetTryHackMeRank(username string) (int, error) {
	url := fmt.Sprintf("https://tryhackme.com/api/v2/public-profile?username=%s", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("User-Agent", httpclient.GetUserAgent())

	resp, err := httpclient.Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("user not found")
	}

	var thm TryHackMeResponse
	if err := json.NewDecoder(resp.Body).Decode(&thm); err != nil {
		return 0, err
	}

	if thm.Status != "success" {
		return 0, fmt.Errorf("user not found")
	}

	return thm.Data.Rank, nil
}
