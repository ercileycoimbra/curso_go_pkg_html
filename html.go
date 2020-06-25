package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Titulo obtem o título de uma janela html
func Titulo(urls ...string) <-chan string {
	c := make(chan string) 

	for _, url := range urls {
		go func(cUrl string) {
			resp, _ := http.Get(cUrl)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}
