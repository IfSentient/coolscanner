package active

import (
	"coolscanner/pkg/models"
	"net"
	"net/http"
	"time"
)

type Probe[T any] struct {
	Requester      func(addr net.Addr) (*T, error)
	ResponseParser func(response *T) (*models.Problem, bool)
}

func (p *Probe) Check(addr net.Addr) (*models.Problem, bool) {
	resp, err := p.Requester(addr)
	if err != nil {
		return nil, false
	}
	return p.ResponseParser(resp)
}

type HTTPRequester struct {
	client          http.Client
	requestModifier func(r *http.Request)
}

func NewHTTPRequester(requestModifier func(r *http.Request)) *HTTPRequester {
	return &HTTPRequester{
		client: http.Client{
			Timeout: time.Second * 5,
		},
		requestModifier: requestModifier,
	}
}

func (h *HTTPRequester) Request(addr net.Addr) (*http.Response, error) {
	req, err := http.NewRequest("GET", addr.String(), nil)
	if err != nil {
		return nil, err
	}
	if h.requestModifier != nil {
		h.requestModifier(req)
	}
	resp, err := h.client.Do(req)
	if err != nil && resp == nil {
		return nil, err
	}
	return resp, nil
}
