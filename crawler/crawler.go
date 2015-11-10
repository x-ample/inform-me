package crawler

import (
	"golang.org/x/net/html"
	"fmt"
	"io"
	"net/http"
)

func Crawler() {
	subject := "politicas"
	engineURL := "https://www.google.com.br/search?tbm=nws"

	url := engineURL + "&" + "q=" + subject
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	links := parser(resp.Body)
	fmt.Printf("%s\n", links)
}

func parser(httpBody io.Reader) []string {
	links := make([]string, 0)
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
	}
}
