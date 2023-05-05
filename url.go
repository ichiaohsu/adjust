package main

import "net/url"

func parseURL(addresses []string) ([]string, error) {
	var results []string
	for _, address := range addresses {
		parsed, err := url.Parse(address)
		if err != nil {
			return nil, err
		}
		if parsed.Scheme == "" {
			parsed.Scheme = "http"
		}
		results = append(results, parsed.String())
	}
	return results, nil
}
