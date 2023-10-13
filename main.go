package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ncruces/zenity"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <baseurl>\n", os.Args[0])
		os.Exit(1)
	}
	jiraurl := os.Args[1]
	card, err := zenity.Entry("Enter card ID:", zenity.Title("Open Jira card"))
	if err != nil {
		fmt.Printf("error reading Jira card ID: %v\n", err)
	}
	card = strings.Trim(card, " ")
	if card == "" {
		fmt.Printf("no input. Exiting.\n")
		return
	}
	url := fmt.Sprintf("%s/browse/%s", jiraurl, card)
	fmt.Printf("opening URL %s\n", url)
	err = open.Run(url)
	if err != nil {
		fmt.Printf("error reading opening URL in browser: %v\n", err)
	}
}
