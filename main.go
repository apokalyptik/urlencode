package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

var str = ""
var decoding = false

func main() {
	flag.StringVar(&str, "data", str, "string to encode")
	flag.BoolVar(&decoding, "decode", decoding, "decode instead of encode")
	flag.Parse()
	if str == "" {
		if data, err := ioutil.ReadAll(os.Stdin); err != nil {
			log.Fatal(err.Error())
		} else {
			str = string(data)
		}
	}
	if !decoding {
		os.Stdout.WriteString(url.QueryEscape(str))
	} else {
		out, err := url.QueryUnescape(str)
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
		}
		os.Stdout.WriteString(out)
		if err != nil {
			os.Exit(1)
		}
	}
}
