package service

import (
	"errors"

	"github.com/tushar9989/github-search/internal/models"
	"github.com/tushar9989/github-search/internal/repository"
)

type searchService struct {
	currentSearchTerm string
	resultCache       []models.SearchResult
	searchSource      repository.SearchResultSource
}

func NewSearchService(source repository.SearchResultSource) searchService {
	m := new(searchService)
	m.currentSearchTerm = ""
	m.searchSource = source
	return *m
}

func (current *searchService) GetProjectNames(search string) ([]string, error) {
	if search == "" {
		return nil, errors.New("Search string cannot be blank")
	}

	if search == current.currentSearchTerm && len(current.resultCache) != 0 {
		return getProjectNamesFromList(current.resultCache), nil
	}

	newResult, err := current.searchSource.Search(search)

	if err != nil {
		return nil, err
	}

	current.resultCache = newResult
	current.currentSearchTerm = search

	return getProjectNamesFromList(current.resultCache), nil
}

func (current *searchService) GetItem(index int) (models.SearchResult, error) {
	if index > len(current.resultCache) {
		return models.SearchResult{}, errors.New("Index out of bounds.")
	}

	return current.resultCache[index-1], nil
}

func getProjectNamesFromList(list []models.SearchResult) []string {
	projectNames := make([]string, 0)

	for _, v := range list {
		projectNames = append(projectNames, v.ProjectName)
	}

	return projectNames
}
