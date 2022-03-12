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

	e2eAuth := endpoint.NewE2eAuth(*ics)
	e2eAuth.E2eTestV1auth401()
	e2eAuth.E2eTestV1user201()
}
