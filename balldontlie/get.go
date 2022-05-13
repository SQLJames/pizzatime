package balldontlie

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Config) Get(url string, queryParams url.Values) []byte {
	c.Logger.Trace(fmt.Sprintf("Performing Http Get %s", url))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Logger.Fatal(err)
	}

	req.URL.RawQuery = queryParams.Encode()
	c.Logger.Debug(req.URL.String())
	resp, err := http.Get(req.URL.String())
	if err != nil {
		c.Logger.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes := c.parseBody(resp.Body)
	return bodyBytes
}

func (c *Config) parseBody(respBody io.Reader) []byte {
	bodyBytes, err := ioutil.ReadAll(respBody)
	if err != nil {
		c.Logger.Fatal(err)
	}
	return bodyBytes
}
