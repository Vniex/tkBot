package utils

import (
	"net/http"
	"net/url"
	"time"
)



func NewHttpClient(timeout int,proxyUrl string)  *http.Client{
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}

	transport := &http.Transport{Proxy: proxy}
	if proxyUrl==""{
		return &http.Client{Timeout:  time.Duration(timeout) * time.Second}
	}
	return &http.Client{Transport: transport, Timeout:  time.Duration(timeout) * time.Second}
}


