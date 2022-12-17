package active

import (
	"coolscanner/pkg/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type httpAddr struct {
	addr string
}

func (h *httpAddr) String() string {
	return h.addr
}

func (h *httpAddr) Network() string {
	return ""
}

type Bar[T any] interface {
	Do(T)
}

type Foo[T any] struct {
	B func(T)
}

type Bar1[T any] struct {
}

func (b Bar1[T]) Do(T) {
	return
}

func TestFoo(t *testing.T) {
	f := Foo[string]{
		B: func(string) {},
	}
	fmt.Println(f)
}

func TestHTTPRequester_Request(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	var requester Requester //[http.Response]
	requester = NewHTTPRequester(func(r *http.Request) {})
	var probe Probe
	probe = Probe{
		Requester: requester,
		ResponseParser: func(resp *http.Response) (*models.Problem, bool) {
			return &models.Problem{
				Type: "open service",
				Meta: map[string]interface{}{
					"path":   resp.Request.URL.Path,
					"status": resp.StatusCode,
				},
			}, true
		},
	}
	h := httpAddr{server.URL}
	problem, ok := probe.Check(&h)
	fmt.Println(ok)
	fmt.Println(problem)
	t.Fail()
}
