package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func extractHref(token html.Token) (ok bool, href string) {
	for _, a := range token.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	return ok, href
}

func exctractAnchors(body io.ReadCloser) {
	tokenizer := html.NewTokenizer(body)

	for {
		tokenType := tokenizer.Next()

		switch {
		case tokenType == html.ErrorToken:
			// We have iterated over the entire doc
			return
		case tokenType == html.StartTagToken:
			token := tokenizer.Token()

			isAnchor := token.Data == "a"
			if !isAnchor {
				continue
			}

			ok, url := extractHref(token)
			if ok {

				isHTTP := strings.Index(url, "http") == 0
				if isHTTP {
					fmt.Println("url: ", url)
				}
			}
		}
	}

}

func crawlURL(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()

		body := resp.Body
		exctractAnchors(body)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing url.")
		os.Exit(1)
	}

	crawlURL(os.Args[1])

	fmt.Println("Done crawling url.")
}
