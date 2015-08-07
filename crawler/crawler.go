package crawler

import (
	"fmt"
	"io/ioutil"
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

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("%v\n", string(body))
}
