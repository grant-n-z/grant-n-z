package main

import (
	"flag"
	"github.com/tomoyane/grant-n-z/e2e/endpoint"
)

func main() {
	var (
		ics = flag.String("url", "http://localhost:8080", "Set target url. http(s)://xxxxxxxxx")
	)
	flag.Parse()

	testIF := []endpoint.TestIF{endpoint.NewE2eAuth(*ics)}
	for _, test := range testIF {
		test.Init()
		test.Run()
		test.Close()
	}
}
