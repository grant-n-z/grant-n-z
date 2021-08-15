package e2eclient

import (
	"net/http"
)

var client = new(http.Client)

func GetHttpClientInstance() *http.Client {
	return client
}
