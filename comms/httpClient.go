package comms

import (
	"io"
	"net/http"
)

type HttpClient struct {
	Client *http.Client
}

func NewHttpClient(client *http.Client) *HttpClient {
	return &HttpClient{Client: client}
}

func (httpClient *HttpClient) GetRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

func (httpClient *HttpClient) DoRequest(r *http.Request) (*http.Response, error) {
	return httpClient.Client.Do(r)
}
