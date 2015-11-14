package crawler

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type RssItems struct {
	index       *int
	title       string
	link        string
	guid        map[string]string
	pubdate     time.Time
	description string
}

func Crawler() {
	subject := "politicas"
	engineUrl := "https://news.google.com.br/news/section?output=rss"
	url := engineUrl + "&" + "q=" + subject + "&output=rss"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	//rssData := []byte(body)

	fmt.Printf("%s", string(body))

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
