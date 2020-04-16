package v1

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
)

var (
	auth       Auth
	statusCode int
)

func init() {
	log.InitLogger("info")
	ctx.InitContext()

	auth = AuthImpl{tokenProcessor: StubTokenProcessor{}}
}

// Test constructor
func TestGetAuthInstance(t *testing.T) {
	GetAuthInstance()
}

// Test method not allow
func TestAuth_MethodNotAllowed(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodPost}
	auth.Api(response, &request)

	if statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect TestAuth_MethodNotAllow test.")
		t.FailNow()
	}
}

// Test get
func TestAuth_Get_Ok(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, URL: &url.URL{}, Method: http.MethodGet}
	request.Header.Set(middleware.Authorization, "Bearer stub")
	auth.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestAuth_Get_Ok test.")
		t.FailNow()
	}
}

// Less than stub struct
// ResponseWriter
type StubResponseWriter struct {
}

func (w StubResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w StubResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (w StubResponseWriter) WriteHeader(code int) {
	statusCode = code
}
