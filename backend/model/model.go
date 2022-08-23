package model

type Proxy struct {
	Name string
}

type Instance struct {
	Proxy   Proxy
	Name    string
	Version string
	IP      string
}

type Listener struct {
	Port     int
	BindAddr string
}

type Route struct {
	Proxy Proxy
	Host  string
	Path  string
}
