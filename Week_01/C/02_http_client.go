package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// look for a url flag on the command line
	// go run D.01b_http_client.go -url <myurl>
	// defaults to the Catboy API
	url := flag.String("url", "https://catboys.com/api/baka", "-url <myurl>")
	flag.Parse()

	var body []byte

	// we introduce a new package here
	resp, err := http.Get(*url)
	if err == nil {
		defer resp.Body.Close()
	}

	if err == nil {
		body, err = ioutil.ReadAll(resp.Body)
	}

	if err == nil {
		fmt.Println(string(body))
	} else {
		fmt.Println(err.Error())
	}

}
