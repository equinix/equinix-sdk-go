package transport

import (
	"log"
	"net/http"
	"strings"
)

type deprecationLoggingTransport struct {
	transport http.RoundTripper
}

// NewDeprecationLoggingTransport creates an
// http.RoundTripper that logs headers defined
// by https://tools.ietf.org/html/rfc8594
func NewDeprecationLoggingTransport(transport http.RoundTripper) *deprecationLoggingTransport {
	return &deprecationLoggingTransport{
		transport: transport,
	}
}

// RoundTrip logs Sunset and/or Deprecation headers that appear on the response
func (t *deprecationLoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := t.transport.RoundTrip(req)
	if err != nil {
		return res, err
	}

	dumpDeprecation(res)

	return res, nil
}

func dumpDeprecation(resp *http.Response) {
	uri := ""
	if resp.Request != nil {
		uri = resp.Request.Method + " " + resp.Request.URL.Path
	}

	deprecation := resp.Header.Get("Deprecation")
	if deprecation != "" {
		if deprecation == "true" {
			deprecation = ""
		} else {
			deprecation = " on " + deprecation
		}
		log.Printf("WARNING: %q reported deprecation%s", uri, deprecation)
	}

	sunset := resp.Header.Get("Sunset")
	if sunset != "" {
		log.Printf("WARNING: %q reported sunsetting on %s", uri, sunset)
	}

	links := resp.Header.Values("Link")

	for _, s := range links {
		for _, ss := range strings.Split(s, ",") {
			if strings.Contains(ss, "rel=\"sunset\"") {
				link := strings.Split(ss, ";")[0]
				log.Printf("WARNING: See %s for sunset details", link)
			} else if strings.Contains(ss, "rel=\"deprecation\"") {
				link := strings.Split(ss, ";")[0]
				log.Printf("WARNING: See %s for deprecation details", link)
			}
		}
	}
}
