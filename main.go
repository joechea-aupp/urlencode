package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	encodedText := flag.String("text", "A%", "Text to encode")
	flag.Parse()

	if flag.NArg() > 0 {
		*encodedText = flag.Arg(0)
	}

	encodedString := url.QueryEscape(*encodedText)
	encodedString = strings.Replace(encodedString, "+", "%20", -1)

	fmt.Println("url encoded string:", encodedString)
	clipboard.WriteAll(encodedString)
}
