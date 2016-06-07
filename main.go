package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

var str = ""

func main() {
	flag.StringVar(&str, "data", str, "string to encode")
	flag.Parse()
	if str == "" {
		if data, err := ioutil.ReadAll(os.Stdin); err != nil {
			log.Fatal(err.Error())
		} else {
			str = string(data)
		}
	}
	os.Stdout.WriteString(url.QueryEscape(str))
}
