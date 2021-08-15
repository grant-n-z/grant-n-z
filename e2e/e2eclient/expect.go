package e2eclient

import (
	"fmt"
	"net/http"
)

func ExpectUnauthorized(res *http.Response, err error, endpoint string, method string) {
	if err != nil{
		FailE2eTest(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnauthorized {
		FailE2eTest(fmt.Sprintf("Fail unauthorized test. Actual: %d, Expect: %d. Endpoint: %s, Method: %s", res.StatusCode, http.StatusUnauthorized, endpoint, method))
	}
	SuccessE2eTest(fmt.Sprintf("Pass unauthorized test. Endpoint: %s, Method: %s", endpoint, method))
}
