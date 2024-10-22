package ginex

import (
	"core/api"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testAPI struct{}

func (m testAPI) Call() (interface{}, error) {
	return "ok", nil
}

func Test_apiPort_Listen(t *testing.T) {
	api.Register("endpoint", "api", testAPI{})

	self := new(apiPort)
	self.req = httptest.NewRequest(
		"POST",
		"/endpoint/api",
		strings.NewReader(""),
	)
	self.resp = httptest.NewRecorder()
	self.Listen()

	res := self.resp.(*httptest.ResponseRecorder).Result()
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(
		t,
		string(body),
		`{"data":"ok","error":0}`,
	)
}
