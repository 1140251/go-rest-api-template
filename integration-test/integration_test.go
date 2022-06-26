package integration_test

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

const (
	// Attempts connection
	host = "app:8080"

	// HTTP REST
	basePath = "http://" + host + "/v1"
)

func TestHTTPDoTranslate(t *testing.T) {
	body := `{
		"message": "example note",
	}`
	Test(t,
		Description("Create note Success"),
		Post(basePath+"/note/"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".message").Equal("example note"),
	)
}
