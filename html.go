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
			aRetorno := r.FindStringSubmatch(string(html))

			if cap(aRetorno) == 0 {
				c <- "Erro ao ler página " + cUrl
				return
			}

			c <- aRetorno[1]
		}(url)
	}

	return c
}
