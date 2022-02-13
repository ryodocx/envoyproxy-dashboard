package client

type Route struct {
	Domain string
}

func (c *Client) GetRoutes() ([]*Route, error) {
	// TODO
	return []*Route{}, nil
}
