package main

import (
	"flag"
	"fmt"
	"github.com/pkg/browser"
	"os"
	"strings"
)

type SearchEngine struct {
	Name    string
	Pattern string
	Divider string
}

var bing = SearchEngine{"Bing", "https://www.bing.com/search?q=", "+"}
var duckduckgo = SearchEngine{"DuckDuckgo", "https://duckduckgo.com/?q=", "+"}
var google = SearchEngine{"Google", "https://www.google.com/search?q=", "+"}
var yahoo = SearchEngine{"Yahoo!", "https://search.yahoo.com/search?p=", "+"}

func generate_search_url(query string, engine SearchEngine) string {
	q := strings.Join(strings.Split(query, " "), engine.Divider)
	return engine.Pattern + q
}

func generate_search_urls(query string, engines []SearchEngine) []string {
	var urls []string
	for _, engine := range engines {
		urls = append(urls, generate_search_url(query, engine))
	}
	return urls
}

func select_engines(flags map[string]bool) []SearchEngine {
	var engines []SearchEngine

	if flags["bing"] == true {
		engines = append(engines, bing)
	}
	if flags["duckduckgo"] == true {
		engines = append(engines, duckduckgo)
	} else if flags["ddg"] == true {
		engines = append(engines, duckduckgo)
	}
	if flags["google"] == true {
		engines = append(engines, google)
	}
	if flags["yahoo"] == true {
		engines = append(engines, yahoo)
	}

	return engines
}

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
		engines := select_engines(flags)

		query := searchCmd.Arg(len(searchCmd.Args()) - 1)
		urls := generate_search_urls(query, engines)

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
		engines := select_engines(flags)

		query := urlsCmd.Arg(len(urlsCmd.Args()) - 1)
		urls := generate_search_urls(query, engines)

		for _, u := range urls {
			fmt.Println(u)
		}
	}
}
