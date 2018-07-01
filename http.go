package main

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []byte{}, errors.New(resp.Status)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	return respBytes, err
}
