package main

import (
	"github.com/tomoyane/grant-n-z/gnzserver"
)

func main() {
	server()
}

// Example GrantNZ server
func server() {
	gnzserver.NewGrantNZServer().Run()
}

// Example GrantNZ cache
func cache() {
	gnzserver.NewGrantNZServer().Run()
}
