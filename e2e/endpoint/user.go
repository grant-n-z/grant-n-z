package endpoint

import (
	"bytes"

	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/e2e/e2eclient"
	"github.com/tomoyane/grant-n-z/gnz/entity"
)

type E2eUser struct {
	BaseUrl    string
	V1Endpoint string
}

var (
	v1UserEndpoint = "/api/v1/auth"
)

func NewE2eUser(url string) E2eUser {
	return E2eUser{
		BaseUrl:    url,
		V1Endpoint: url + v1UserEndpoint,
	}
}

func (e E2eAuth) E2eTestV1user201() {
	svcs := GetService(e.BaseUrl)
	secret := svcs[0].Secret

	user := entity.User{
		Username: "e2e-user",
		Email:    "e2e-email@gmail.com",
		Password: "e2e_password",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	req, _ := http.NewRequest(http.MethodPost, e.V1Endpoint, &buf)
	req.Header.Set("Client-Secret", secret)

	res, err := e2eclient.GetHttpClientInstance().Do(req)
	e2eclient.ExpectCreated(res, err, e.V1Endpoint, http.MethodPost)
}
