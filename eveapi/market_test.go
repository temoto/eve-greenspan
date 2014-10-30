package eveapi

import (
	"bytes"
	"io"
	"net/http"
)

type mockTransport struct {
	r io.Reader
}

func (rt *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func (t *mockTransport) Read(b []byte) (int, error) { return t.r.Read(b) }
func (t *mockTransport) Close() error               { return nil }

func newMockHttpClient(response []byte) *http.Client {
	rt := &mockTransport{r: bytes.NewReader(response)}
	return &http.Client{Transport: rt}
}
