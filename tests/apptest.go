package tests

import (
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

// Run this before a request
func (t *AppTest) Before() {
	println("Set up")
}

// Run this after request
func (t *AppTest) After() {
	println("Tear down")
}