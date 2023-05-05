package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

type clientOpt func(*simpleClient)

type simpleClient struct {
	client *http.Client
}

func newSimpleClient(opts ...clientOpt) *simpleClient {
	sc := simpleClient{
		client: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(&sc)
	}
	return &sc
}

func (c *simpleClient) getHash(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for url: %s, err: %s", url, err.Error())
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to request url: %s, err: %s", url, err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http request to %s not ok: %d", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body error: %s", err.Error())

	}
	// hash
	h := md5.New()
	if _, err = h.Write(body); err != nil {
		return "", fmt.Errorf("failed to hash response for url %s, error: %s", url, err.Error())
	}
	resp.Body.Close()
	return hex.EncodeToString(h.Sum(nil)), nil
}
