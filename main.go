package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ryulational/multisearch/engine"

	"github.com/pkg/browser"
)

func main() {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchBing := searchCmd.Bool("bing", false, "Seach Bing")
	searchDdg := searchCmd.Bool("ddg", false, "Search DuckDuckGo")
	searchDuckduckgo := searchCmd.Bool("duckduckgo", false, "Search DuckDuckGo")
	searchGoogle := searchCmd.Bool("google", false, "Search Google")
	searchYahoo := searchCmd.Bool("yahoo", false, "Search Yahoo!")

	urlsCmd := flag.NewFlagSet("urls", flag.ExitOnError)
	urlsBing := urlsCmd.Bool("bing", false, "Use Bing")
	urlsDdg := urlsCmd.Bool("ddg", false, "Use DuckDuckGo")
	urlsDuckduckgo := urlsCmd.Bool("duckduckgo", false, "Use DuckDuckGo")
	urlsGoogle := urlsCmd.Bool("google", false, "Use Google")
	urlsYahoo := urlsCmd.Bool("yahoo", false, "Use Yahoo!")

	if len(os.Args) <= 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println("Use subcommand 'search' or 'get-urls'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "search":
		searchCmd.Parse(os.Args[2:])
		flags := make(map[string]bool)
		flags["bing"] = *searchBing
		flags["ddg"] = *searchDdg
		flags["duckduckgo"] = *searchDuckduckgo
		flags["google"] = *searchGoogle
		flags["yahoo"] = *searchYahoo
		engines := engine.Select_engines(flags)

		query := searchCmd.Arg(len(searchCmd.Args()) - 1)
		urls := engine.Generate_search_urls(query, engines)

		for _, u := range urls {
			fmt.Println("Opening ", u)
			browser.OpenURL(u)
		}
	case "get-urls":
		urlsCmd.Parse(os.Args[2:])
		flags := make(map[string]bool)
		flags["bing"] = *urlsBing
		flags["ddg"] = *urlsDdg
		flags["duckduckgo"] = *urlsDuckduckgo
		flags["google"] = *urlsGoogle
		flags["yahoo"] = *urlsYahoo
		engines := engine.Select_engines(flags)

		query := urlsCmd.Arg(len(urlsCmd.Args()) - 1)
		urls := engine.Generate_search_urls(query, engines)

		for _, u := range urls {
			fmt.Println(u)
		}
	}
}
