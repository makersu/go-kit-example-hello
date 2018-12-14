package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	http "net/http"
	"net/url"
	"strings"

	endpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	pkgendpoint "github.com/makersu/go-kit-example-hello/hello/pkg/endpoint"
	pkghttp "github.com/makersu/go-kit-example-hello/hello/pkg/http"
	service "github.com/makersu/go-kit-example-hello/hello/pkg/service"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]kithttp.ClientOption) (service.HelloService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var helloEndpoint endpoint.Endpoint
	{
		helloEndpoint = kithttp.NewClient("POST", copyURL(u, "/hello"), encodeHTTPGenericRequest, decodeHelloResponse, options["Hello"]...).Endpoint()
	}

	return pkgendpoint.Endpoints{HelloEndpoint: helloEndpoint}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeHelloResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeHelloResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, pkghttp.ErrorDecoder(r)
	}
	var resp pkgendpoint.HelloResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
