package eveapi

import (
	"net/http"
)

type Client struct {
	HttpClient *http.Client
}

type baseResponse struct {
	// restore
}

func parseMoney(s string) (int64, error) {
	panic("restore")
}
