package service_test

import (
	"testing"

	"github.com/tushar9989/github-search/internal/models"
	"github.com/tushar9989/github-search/internal/service"
)

type MockSearchSource struct {
	searchCallCount int
}

func (current *MockSearchSource) getSearchCallCount() int {
	return current.searchCallCount
}
func (current *MockSearchSource) Search(search string) ([]models.SearchResult, error) {
	current.searchCallCount++
	return make([]models.SearchResult, 3), nil
}

func TestSearchNotCalledAgain(t *testing.T) {
	mockSearchSource := MockSearchSource{}
	searchService := service.NewSearchService(&mockSearchSource)
	_, _ = searchService.GetProjectNames("one")
	_, _ = searchService.GetProjectNames("one")
	expected := 1
	actual := mockSearchSource.getSearchCallCount()
	if actual != expected {
		t.Error("Search called does not match expected count. Actual: ", actual)
	}
}

func TestErrorOnOutOfBounds(t *testing.T) {
	mockSearchSource := MockSearchSource{}
	service := service.NewSearchService(&mockSearchSource)
	_, _ = service.GetProjectNames("one")
	_, err := service.GetItem(5)
	if err == nil {
		t.Error("Error was not thrown on out of bounds access input. ")
	}
}

func TestSearchNotCalledOnGet(t *testing.T) {
	mockSearchSource := MockSearchSource{}
	service := service.NewSearchService(&mockSearchSource)
	_, _ = service.GetProjectNames("one")
	_, _ = service.GetItem(1)
	_, _ = service.GetItem(2)
	_, _ = service.GetItem(1)
	expected := 1
	actual := mockSearchSource.getSearchCallCount()
	if actual != expected {
		t.Error("Search called for same search term. Actual: ", actual)
	}
}

func TestCacheNotClearedOnBlankInputAndErrorThrown(t *testing.T) {
	mockSearchSource := MockSearchSource{}
	service := service.NewSearchService(&mockSearchSource)
	_, _ = service.GetProjectNames("one")
	_, err := service.GetProjectNames("")
	t.Run("CheckErrorThrown", func(t *testing.T) {
		if err == nil {
			t.Error("Error was not thrown on blank input.")
		}
	})
	t.Run("CheckCacheNotCleared", func(t *testing.T) {
		if _, err = service.GetItem(1); err != nil {
			t.Error("Cache was cleared on blank input. Error: ", err)
		}
	})
}
