package models

import "fmt"

// SearchResult Represents a single result from the github page
type SearchResult struct {
	ProjectName   string
	Owner         string
	Description   string
	LastUpdated   string
	RepositoryURL string
}

func (item SearchResult) String() string {
	return fmt.Sprintf("ProjectName: %v\nOwner: %v\nDescription: %v\nLastUpdated: %v\nRepositoryURL: %v\n", item.ProjectName, item.Owner, item.Description, item.LastUpdated, item.RepositoryURL)
}
