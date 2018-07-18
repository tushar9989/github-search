package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/tushar9989/github-search/internal/repository"
	"github.com/tushar9989/github-search/internal/service"
)

func main() {
	searchDone := false
	scanner := bufio.NewScanner(os.Stdin)
	service := service.NewSearchService(&repository.GitHubSearchSource{})

	fmt.Print("Enter search term: ")

	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "q":
			os.Exit(0)
		case "n":
			searchDone = false
			fmt.Print("Enter search term: ")
		default:
			if !searchDone {
				projectNames, err := service.GetProjectNames(text)
				if err != nil {
					fmt.Println(err)
				} else {
					searchDone = true
					for i, v := range projectNames {
						fmt.Println(i+1, ". ", v)
					}
				}
			} else {
				i, err := strconv.Atoi(text)
				if err != nil {
					fmt.Println("Invalid Input! ", err)
				} else {
					item, err := service.GetItem(i)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println(item)
					}
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
