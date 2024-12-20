package client

import (
	"fmt"
	"hsr-api/enums"
	"io"
	"net/http"
)

// create a client struct
type Client struct {
	baseURL    string
	httpClient *http.Client
}

func (client *Client) MakeRequest(endpoint string) (string, error) {
	// request url
	url := client.baseURL + endpoint
	fmt.Printf("Making a request to %s\n", url)

	// create a new request using httpClient
	response, err := client.httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func NewClient(language enums.Language) *Client {
	return &Client{
		baseURL:    fmt.Sprintf("https://sr.yatta.moe/api/v2/%v", language.String()),
		httpClient: &http.Client{},
	}
}
