package web

import (
	"net/http"
	"io/ioutil"
)

func LoadPage (url string) (string, error) {
	response, error := http.Get(url)
	if error != nil {
		return "", error
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		return "", error
	}

	return string(body), nil
}