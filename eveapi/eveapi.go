package eveapi

import (
	"net/http"
)

type Client struct {
	Http *http.Client

	BaseUrl string
	keyID   string
	vCode   string
}

type baseResponse struct {
	CurrentTime EveTime
	CachedUntil EveTime
}

func parseMoney(s string) (int64, error) {
	panic("restore")
}
