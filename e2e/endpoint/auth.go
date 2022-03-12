package endpoint

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/e2e/e2eclient"
)

type E2eAuth struct {
	BaseUrl    string
	V1Endpoint string
}

var (
	v1AuthEndpoint = "/api/v1/auth"
)

func NewE2eAuth(url string) E2eAuth {
	return E2eAuth{
		BaseUrl:    url,
		V1Endpoint: url + v1AuthEndpoint,
	}
}

func (e E2eAuth) E2eTestV1auth401() {
	req, _ := http.NewRequest(http.MethodGet, e.V1Endpoint, nil)
	res, err := e2eclient.GetHttpClientInstance().Do(req)
	e2eclient.ExpectUnauthorized(res, err, e.V1Endpoint, http.MethodGet)
}
