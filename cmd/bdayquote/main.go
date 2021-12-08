package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func getQuote(lang string) string {
	var (
		err        error
		response   *http.Response
		retries    = 3
		client     = &http.Client{}
		req        *http.Request
		quoteBytes []byte
		quote      string
	)

	req, err = http.NewRequest("GET", "https://raccoon-ninja-birthday-api.herokuapp.com/", nil)

	if lang != "" {
		req.Header.Set("Accept-Language", lang)
	}

	for retries > 0 {
		response, err = client.Do(req)

		if err != nil {
			log.Println(err)
			retries -= 1
		} else {
			break
		}
	}

	if response == nil {
		panic("Tried a couple of times, but could not get a quote.")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panicf("Failed to close http response reader. Error: %s", err)
		}
	}(response.Body)

	quoteBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicf("Failed to read message from http response body. Error: %s", err)
	}

	quote = fmt.Sprintf("%s\n", quoteBytes)

	return quote
}

func getQuoteLang() string {
	if len(os.Args) <= 1 {
		return ""
	}

	lang := os.Args[1]
	if lang == "" {
		return ""
	}

	return strings.Trim(strings.ToLower(lang), " ")
}

func main() {
	lang := getQuoteLang()
	quote := getQuote(lang)
	fmt.Print(quote)

	err := clipboard.WriteAll(quote)
	if err != nil {
		log.Panicf("Failed to save message to clipboard. Error: %s", err)
	}
}
