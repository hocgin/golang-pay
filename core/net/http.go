package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func PostString(url string, data string) (string, error) {
	resp, err := http.Post(url, "application/xml", strings.NewReader(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetString(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetUrl(baseUrl string, values map[string]interface{}) string {
	return GetUrlBase(baseUrl, values, func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	})
}

func GetUrlBase(baseUrl string, values map[string]interface{}, handleValue func(interface{}) string) string {
	i := 0
	for key, value := range values {
		if i == 0 {
			baseUrl += "?"
		} else {
			baseUrl += "&"
		}
		baseUrl += fmt.Sprintf(`%s=%s`, key, handleValue(value))
		i += 1
	}
	return baseUrl
}
func GetUrlEncode(baseUrl string, values map[string]interface{}) string {
	return GetUrlBase(baseUrl, values, func(i interface{}) string {
		return url.QueryEscape(fmt.Sprintf("%s", i))
	})
}
