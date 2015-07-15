package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://www.dn.se"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("%v\n", err)
	}

	if resp.StatusCode != 200 {
		fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// Close the response once we return from the function.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
