package main

import (
	"json_http/github"
	"json_http/myTemplate"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	//if err := myTemplate.Report.Execute(os.Stdout, result); err != nil {
	//	log.Fatal(err)
	//}
	if err := myTemplate.IssueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
