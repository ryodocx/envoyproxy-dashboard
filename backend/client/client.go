package client

type Config struct{}
type Client struct{}

func New() (*Client, error) {
	return &Client{}, nil
}
