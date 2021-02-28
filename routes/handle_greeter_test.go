package routes

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestHandleGreeting(t *testing.T) {
	is := is.New(t)

	srv, err := NewServer()
	is.NoErr(err)

	hsrv := httptest.NewServer(srv)
	defer hsrv.Close()

	req, err := http.NewRequest("post", hsrv.URL+"/api/GreeterService/Greet", strings.NewReader("El"))
	is.NoErr(err)
	client := &http.Client{
		Timeout: 1 * time.Minute,
	}

	resp, err := client.Do(req)
	is.NoErr(err)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	is.NoErr(err)
	is.Equal(string(b), `Well 'ello there, El.`)
}
