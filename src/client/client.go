package client

import (
	"code.google.com/p/mahonia"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Client struct {
	http.Client
	ProxyAddr *url.URL
	encode    mahonia.Encoder
}

func New() *Client {
	cl := &Client{}

	jar, _ := cookiejar.New(nil)
	cl.Jar = jar
	return cl
}
func (this *Client) UseProxy(addr string) error {
	u, err := url.Parse(addr)
	if err != nil {
		return fmt.Errorf("cannot parse proxy addr: %v", err)
	}
	this.ProxyAddr = u
	transport := &http.Transport{}
	transport.Proxy = func(*http.Request) (*url.URL, error) {
		return this.ProxyAddr, nil
	}

	this.Client.Transport = transport
	return nil
}
func (this *Client) UseEncoder(encoder string) {
	this.encode = mahonia.NewEncoder(encoder)
}

func (this *Client) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	for i, v := range data {
		s := []string{}
		for _, u := range v {
			s = append(s, this.encode.ConvertString(u))
		}
		data[this.encode.ConvertString(i)] = s
	}
	return this.Client.PostForm(url, data)
}
