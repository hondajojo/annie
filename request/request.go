package request

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hondajojo/annie/config"
)

// Request base request
func Request(
	method, url string, body io.Reader, headers map[string]string,
) *http.Response {
	client := &http.Client{
		Timeout: time.Second * 100,
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Print(url)
		log.Fatal(err)
	}
	for k, v := range config.FakeHeaders {
		req.Header.Set(k, v)
	}
	req.Header.Set("Referer", url)
	if config.Cookie != "" {
		var cookie string
		if _, fileErr := os.Stat(config.Cookie); fileErr == nil {
			// Cookie is a file
			data, _ := ioutil.ReadFile(config.Cookie)
			cookie = string(data)
		} else {
			// Just strings
			cookie = config.Cookie
		}
		req.Header.Set("Cookie", cookie)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if config.Refer != "" {
		req.Header.Set("Referer", config.Refer)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Print(url)
		log.Fatal(err)
	}
	if config.Debug {
		fmt.Println()
		fmt.Printf("URL: %s\n", url)
		fmt.Printf("Method: %s\n", method)
		fmt.Printf("Headers: %s\n", req.Header)
		fmt.Printf("Status Code: %d\n", res.StatusCode)
	}
	return res
}

// Get get request
func Get(url string) string {
	res := Request("GET", url, nil, nil)
	defer res.Body.Close()
	var reader io.ReadCloser
	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, _ = gzip.NewReader(res.Body)
	} else {
		reader = res.Body
	}
	body, _ := ioutil.ReadAll(reader)
	return string(body)
}

// Headers return the HTTP Headers of the url
func Headers(url, refer string) http.Header {
	headers := map[string]string{
		"Referer": refer,
	}
	res := Request("GET", url, nil, headers)
	defer res.Body.Close()
	return res.Header
}

// Size get size of the url
func Size(url, refer string) int64 {
	h := Headers(url, refer)
	s := h.Get("Content-Length")
	size, _ := strconv.ParseInt(s, 10, 64)
	return size
}

// ContentType get Content-Type of the url
func ContentType(url, refer string) string {
	h := Headers(url, refer)
	s := h.Get("Content-Type")
	// handle Content-Type like this: "text/html; charset=utf-8"
	return strings.Split(s, ";")[0]
}
