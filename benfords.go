package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	url := "https://sports.yahoo.com/nfl/teams/den/stats"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("%v\n", err)
	}

	if resp.StatusCode != 200 {
		fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	d := html.NewTokenizer(resp.Body)
	// for {
	// 	tt := z.Next()
	// 	if tt == html.ErrorToken {
	// 		fmt.Println(" error on token")
	// 	}

	// }
	for {
		// token type
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			return
		}
		token := d.Token()
		switch tokenType {
		case html.StartTagToken: // <tag>
			// type Token struct {
			//     Type     TokenType
			//     DataAtom atom.Atom
			//     Data     string
			//     Attr     []Attribute
			// }
			//
			// type Attribute struct {
			//     Namespace, Key, Val string
			// }
			if token.DataAtom == 0x4702 { // <p> tag
				d.Next()
				fmt.Println(d.Token())
			}
		case html.TextToken: // text between start and end tag
		case html.EndTagToken: // </tag>
		case html.SelfClosingTagToken: // <tag/>

		}
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
}
