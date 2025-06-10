package scrapers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/cedev-1/cyberrank-api/internal/models"
	"github.com/cedev-1/cyberrank-api/pkg/httpclient"

	"github.com/PuerkitoBio/goquery"
)

type RootMeProfile struct {
	Rank       string
	Categories []models.CategoryProgress
	Overall    *models.OverallProgress
}

func GetRootMeProfile(username string) (*RootMeProfile, error) {
	url := fmt.Sprintf("https://www.root-me.org/%s", username)

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

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	rank := doc.Find("div.small-6.medium-3.columns.text-center h3").First().Text()
	rank = strings.TrimSpace(rank)

	if rank == "" {
		return nil, fmt.Errorf("rank not found")
	}

	categories := extractCategories(doc)

	overall := extractOverallProgress(doc)

	return &RootMeProfile{
		Rank:       rank,
		Categories: categories,
		Overall:    overall,
	}, nil
}

func extractCategories(doc *goquery.Document) []models.CategoryProgress {
	var categories []models.CategoryProgress

	doc.Find("div.t-body.tb-padding").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Find("h3").Text(), "Validations") {
			s.Find("div.c100").Each(func(j int, category *goquery.Selection) {
				class, exists := category.Attr("class")
				if !exists {
					return
				}

				percentage := extractPercentageFromClass(class)

				link := category.Find("a")
				if link.Length() == 0 {
					return
				}

				title, _ := link.Attr("title")

				categories = append(categories, models.CategoryProgress{
					Name:       title,
					Percentage: percentage,
				})
			})
		}
	})

	return categories
}

func extractOverallProgress(doc *goquery.Document) *models.OverallProgress {
	var overall *models.OverallProgress

	doc.Find("div.t-body.tb-padding").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Find("h3").Text(), "Validations") {
			progressText := s.Find("h3.text-center").Text()

			re := regexp.MustCompile(`(\d+)%.*?(\d+)/(\d+)`)
			matches := re.FindStringSubmatch(progressText)

			if len(matches) == 4 {
				percentage, _ := strconv.Atoi(matches[1])
				solved, _ := strconv.Atoi(matches[2])
				total, _ := strconv.Atoi(matches[3])

				overall = &models.OverallProgress{
					Percentage: percentage,
					Solved:     solved,
					Total:      total,
				}
			}
		}
	})

	return overall
}

func extractPercentageFromClass(class string) int {
	re := regexp.MustCompile(`p(\d+)`)
	matches := re.FindStringSubmatch(class)

	if len(matches) > 1 {
		percentage, err := strconv.Atoi(matches[1])
		if err == nil {
			return percentage
		}
	}

	return 0
}

func GetRootMeRank(username string) (string, error) {
	profile, err := GetRootMeProfile(username)
	if err != nil {
		return "", err
	}
	return profile.Rank, nil
}
