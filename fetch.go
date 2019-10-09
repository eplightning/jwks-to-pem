package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func fetchJWKSData(input string) ([]byte, error) {
	url, err := url.Parse(input)
	if err != nil {
		return nil, err
	}

	if url.Scheme == "http" || url.Scheme == "https" {
		return fetchJWKSFromHTTP(url.String())
	}

	if len(url.Scheme) == 0 {
		return fetchJWKSFromFile(input)
	}

	return nil, errors.New("unknown JWKS source")
}

func fetchJWKSFromFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// TODO add support for OIDC discovery
func fetchJWKSFromHTTP(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
