package httptest

import (
	"fmt"
	"net/http"
)

type encodable interface {
	Encode() string
}

type handler struct {
	http.Handler
	Cookies    string
	Headers    map[string]string
	HmaxSecret string
}

func (w *handler) HTML(u string, args ...interface{}) *Request {
	hs := map[string]string{}
	for key, val := range w.Headers {
		hs[key] = val
	}
	return &Request{
		URL:     fmt.Sprintf(u, args...),
		handler: w,
		Headers: hs,
	}
}

func (w *handler) JSON(u string, args ...interface{}) *JSON {
	hs := map[string]string{}
	for key, val := range w.Headers {
		hs[key] = val
	}
	hs["Content-Type"] = "application/json"
	return &JSON{
		URL:     fmt.Sprintf(u, args...),
		handler: w,
		Headers: hs,
	}
}

func (w *handler) XML(u string, args ...interface{}) *XML {
	hs := map[string]string{}
	for key, val := range w.Headers {
		hs[key] = val
	}
	hs["Content-Type"] = "application/xml"
	return &XML{
		URL:     fmt.Sprintf(u, args...),
		handler: w,
		Headers: hs,
	}
}

func New(h http.Handler) *handler {
	return &handler{
		Handler: h,
		Headers: map[string]string{},
	}
}