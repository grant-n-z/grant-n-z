package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/e2e/e2eclient"
	"github.com/tomoyane/grant-n-z/gnz/entity"
)

type E2eService struct {
	V1Endpoint string
}

var (
	v1ServiceEndpoint = "/api/v1/service"
)

func NewE2eService(url string) E2eService {
	return E2eService{
		V1Endpoint: url + v1ServiceEndpoint,
	}
}

func (e E2eAuth) E2eTestV1Service() {
	req, _ := http.NewRequest(http.MethodGet, e.V1Endpoint, nil)
	res, err := e2eclient.GetHttpClientInstance().Do(req)
	e2eclient.ExpectOk(res, err, e.V1Endpoint, http.MethodGet)
}

func GetService(url string) []entity.Service {
	req, _ := http.NewRequest(http.MethodGet, url+v1ServiceEndpoint, nil)
	res, _ := e2eclient.GetHttpClientInstance().Do(req)

	var services *[]entity.Service
	json.NewDecoder(res.Body).Decode(services)
	return *services
}
