package client

import "github.com/ryodocx/envoyproxy-dashboard/backend/model"

// type Config struct{}
// type Client struct{}

// func New() (*Client, error) {
// 	// TODO
// 	return &Client{}, nil
// }

// func (c *Client) GetRoutes() ([]*model.Route, error) {
// 	// TODO
// 	return []*model.Route{}, nil
// }

type Client interface {
	GetRoutes() ([]*model.Route, error)
}
