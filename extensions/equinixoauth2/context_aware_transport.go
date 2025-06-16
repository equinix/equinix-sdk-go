// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copied from https://github.com/golang/oauth2/blob/65c15a35147ccc5127e9f8cdf2e07837596e56b4/transport.go
// with modification to support reusing request context

package equinixoauth2

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// ContextAwareTransport is an http.RoundTripper that makes OAuth 2.0 HTTP requests,
// wrapping a base RoundTripper and adding an Authorization header
// with a token from the supplied Sources.
//
// ContextAwareTransport is a low-level mechanism. Most code will use the
// higher-level Config.Client method instead.
type ContextAwareTransport struct {
	// Source supplies the token to add to outgoing requests'
	// Authorization headers.
	Source *ContextAwareTokenSource

	// Base is the base RoundTripper used to make HTTP requests.
	// If nil, http.DefaultTransport is used.
	Base http.RoundTripper
}

func (c *Config) New() *ContextAwareTransport {
	return &ContextAwareTransport{
		Source: c.TokenSource(),
	}
}

// RoundTrip authorizes and authenticates the request with an
// access token from ContextAwareTransport's Source.
func (t *ContextAwareTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqBodyClosed := false
	if req.Body != nil {
		defer func() {
			if !reqBodyClosed {
				//nolint:errcheck // Inherited from upstream; disabling lint to avoid a larger refactor
				req.Body.Close()
			}
		}()
	}

	// Original was modified here to call `TokenWithContext` instead of `Token`,
	// passing in the existing request context
	token, err := t.Source.TokenWithContext(req.Context())
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}

	req2 := cloneRequest(req) // per RoundTripper contract
	token.SetAuthHeader(req2)

	// req.Body is assumed to be closed by the base RoundTripper.
	reqBodyClosed = true
	return t.base().RoundTrip(req2)
}

var cancelOnce sync.Once

// CancelRequest does nothing. It used to be a legacy cancellation mechanism
// but now only it only logs on first use to warn that it's deprecated.
//
// Deprecated: use contexts for cancellation instead.
func (t *ContextAwareTransport) CancelRequest(_ *http.Request) {
	cancelOnce.Do(func() {
		log.Printf("deprecated: golang.org/x/oauth2: ContextAwareTransport.CancelRequest no longer does anything; use contexts")
	})
}

func (t *ContextAwareTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
