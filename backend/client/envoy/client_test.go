package envoy_test

import (
	"envoyproxy-dashboard/backend/client/envoy"
	"testing"

	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"google.golang.org/protobuf/proto"
)

func TestGetConfigDump(t *testing.T) {
	c := envoy.New()
	configDump, err := c.GetConfigDump()
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
		proto.Unmarshal(v.RouteConfig.Value, &route)

		for _, v := range route.VirtualHosts {
			t.Logf("domains: %v:\n", v.Domains)
		}
	}
}

func TestGetRouteConfigurations(t *testing.T) {
	c := envoy.New()
	configs, err := c.GetRouteConfigurations()
	if err != nil {
		t.Fatal(err)
	}

	for i, v := range configs {
		t.Logf("%d:\n", i)
		t.Logf("%#v:\n", v.VirtualHosts[0].Domains)
	}

}
