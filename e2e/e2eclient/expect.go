package e2eclient

import (
	"fmt"
	"net/http"
)

func ExpectCreated(res *http.Response, err error, endpoint string, method string) {
	Expect(res, err, endpoint, method, http.StatusCreated)
}

func ExpectOk(res *http.Response, err error, endpoint string, method string) {
	Expect(res, err, endpoint, method, http.StatusOK)
}

func ExpectUnauthorized(res *http.Response, err error, endpoint string, method string) {
	Expect(res, err, endpoint, method, http.StatusUnauthorized)
}

func Expect(res *http.Response, err error, endpoint string, method string, status int) {
	if err != nil {
		FailE2eTest(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != status {
		FailE2eTest(fmt.Sprintf("Fail %d test. Actual: %d, Expect: %d. Endpoint: %s, Method: %s", status, res.StatusCode, status, endpoint, method))
	}
	SuccessE2eTest(fmt.Sprintf("Pass %d test. Endpoint: %s, Method: %s", status, endpoint, method))
}
