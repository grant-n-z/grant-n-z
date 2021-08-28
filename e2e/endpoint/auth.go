package endpoint

import (
	"fmt"
	"net/http"

	"github.com/tomoyane/grant-n-z/e2e/e2eclient"
)

type E2eAuth struct {
	V1Endpoint string
}

var (
	v1Endpoint = "/api/v1/auth"
)

func NewE2eAuth(url string) TestIF {
	return E2eAuth{
		V1Endpoint: url + v1Endpoint,
	}
}

func (e E2eAuth) Init() {
	fmt.Println("Initialized auth test")
}

func (e E2eAuth) Run() {
	e.e2eTestV1auth401()
}

func (e E2eAuth) Close() {
	fmt.Println("Closed")
}

func (e E2eAuth) e2eTestV1auth401() {
	req, _ := http.NewRequest(http.MethodGet, e.V1Endpoint, nil)
	res, err := e2eclient.GetHttpClientInstance().Do(req)
	e2eclient.ExpectUnauthorized(res, err, v1Endpoint, http.MethodGet)
}
