package engine

import "strings"

type searchEngine struct {
	name    string
	pattern string
	divider string
}

var bing = searchEngine{"Bing", "https://www.bing.com/search?q=", "+"}
var duckduckgo = searchEngine{"DuckDuckgo", "https://duckduckgo.com/?q=", "+"}
var google = searchEngine{"Google", "https://www.google.com/search?q=", "+"}
var yahoo = searchEngine{"Yahoo!", "https://search.yahoo.com/search?p=", "+"}

func generate_search_url(query string, engine searchEngine) string {
	q := strings.Join(strings.Split(query, " "), engine.divider)
	return engine.pattern + q
}

func Generate_search_urls(query string, engines []searchEngine) []string {
	var urls []string
	for _, engine := range engines {
		urls = append(urls, generate_search_url(query, engine))
	}
	return urls
}

func Select_engines(flags map[string]bool) []searchEngine {
	var engines []searchEngine

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
