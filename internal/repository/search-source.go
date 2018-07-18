package repository

import "github.com/tushar9989/github-search/internal/models"

//SearchResultSource interface that defines the methods needed from a search result source
type SearchResultSource interface {
	Search(search string) ([]models.SearchResult, error)
}
