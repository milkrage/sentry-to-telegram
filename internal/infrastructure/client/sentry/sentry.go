package sentry

type Client struct {
	Installation installation
}

func New() *Client {
	return &Client{}
}
