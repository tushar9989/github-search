# GitHub Search

## Dependencies

* Go >= v1.10
* github.com/PuerkitoBio/goquery. Used to parse the HTML.

## Flow

* The program asks for one single term to search. (ex. google, node, apache)
* After the user enters the term and hits enter button, the first 10 results will be displayed in the command-line with numbers.
* After the first 10 search results are displayed, the user can enter a number from 1-10. The result for that particular number will be printed in detail (all the project name, owner, one-line description, last updated & repo url fields).
* The user can repeatedly enter any numbers from 1-10 to get details of that specific app.
* The user can enter n instead of numbers 1-10 to make a new search.
* The program ends when the user types q instead 1-10.

## Running on a unix environment

* `go get github.com/PuerkitoBio/goquery`
* create a directory for the source code `mkdir -p $HOME/go/src/github.com/tushar9989/github-search`
* copy the source code `cp -r * $HOME/go/src/github.com/tushar9989/github-search`
* move to source directory `cd $HOME/go/src/github.com/tushar9989/github-search`
* _optional_ To run tests execute `go test github.com/tushar9989/github-search/internal/repository`
* _optional_ To run tests execute `go test github.com/tushar9989/github-search/internal/service`
* To run the code `go run cmd/main.go`

OR If you have Docker

* `./docker-runner.sh`