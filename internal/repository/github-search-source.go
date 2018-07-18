package repository

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tushar9989/github-search/internal/models"
)

type GitHubSearchSource struct {
}

func getSearchURL(search string) (string, error) {
	if search == "" {
		return "", errors.New("Search string cannot be blank")
	}
	return "https://github.com/search?q=" + search, nil
}

func getRepositoryURLPrefix() string {
	return "https://github.com/"
}

func getHTML(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", errors.New("HTTP request returned non 200 result")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body[:]), nil
}

//GetResultFromHTML Given HTML parses the html and extracts required results
func GetResultFromHTML(html string) ([]models.SearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		return nil, err
	}

	result := make([]models.SearchResult, 0)

	doc.Find(".repo-list-item").Each(func(i int, s *goquery.Selection) {

		urlSuffix := s.Find("div h3 a").Text()
		repositoryURL := getRepositoryURLPrefix() + urlSuffix
		split := strings.Split(urlSuffix, "/")
		owner, projectName := split[0], split[1]
		lastUpdated, _ := s.Find("relative-time").Attr("datetime")
		description := s.Find("div:first-child > p.text-gray").Text()
		description = strings.TrimSpace(description)

		searchResultItem := models.SearchResult{ProjectName: projectName, Owner: owner, Description: description, LastUpdated: lastUpdated, RepositoryURL: repositoryURL}
		result = append(result, searchResultItem)
	})

	return result, nil
}

//Search generates the url makes the call and parses the result
func (current *GitHubSearchSource) Search(search string) ([]models.SearchResult, error) {
	url, err := getSearchURL(search)

	if err != nil {
		return nil, err
	}

	html, err := getHTML(url)

	if err != nil {
		return nil, err
	}

	result, err := GetResultFromHTML(html)

	if err != nil {
		return result, err
	}

	return result, nil
}
