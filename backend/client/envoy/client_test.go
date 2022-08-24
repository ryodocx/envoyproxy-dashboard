package envoy_test

import (
	"net/url"
	"testing"

	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/ryodocx/envoyproxy-dashboard/backend/client/envoy"
	"google.golang.org/protobuf/proto"
)

var envoyClient *envoy.Client

func init() {
	u, err := url.Parse("http://127.0.0.1:15000/config_dump")
	if err != nil {
		panic(err)
	}

	c, err := envoy.New(envoy.Config{
		EnvoyURL: u,
	})
	if err != nil {
		panic(err)
	}
	envoyClient = c
}

func TestGetConfigDump(t *testing.T) {

	configDump, err := envoyClient.GetConfigDump()
	if err != nil {
		t.Fatal(err)
	}

	// t.Logf("%#v\n", configDump.BootstrapConfigDump.LastUpdated.String())

	for i, v := range configDump.RoutesConfigDump.GetDynamicRouteConfigs() {
		t.Logf("%d:\n", i)

		var route route.RouteConfiguration
		proto.Unmarshal(v.RouteConfig.Value, &route)

		for _, v := range route.VirtualHosts {
			t.Logf("domains: %v:\n", v.Domains)
		}
	}
	for i, v := range configDump.RoutesConfigDump.GetStaticRouteConfigs() {
		t.Logf("%d:\n", i)

		var route route.RouteConfiguration
		v.RouteConfig.UnmarshalTo(&route)
		// proto.Unmarshal(v.RouteConfig.Value, &route)

		for _, v := range route.VirtualHosts {
			t.Logf("domains: %v:\n", v.Domains)
		}
	}
}

func TestGetRouteConfigurations(t *testing.T) {
	configs, err := envoyClient.GetRouteConfigurations()
	if err != nil {
		t.Fatal(err)
	}

	for i, v := range configs {
		t.Logf("%d:\n", i)
		t.Logf("%#v:\n", v.VirtualHosts[0].Domains)
	}
}
