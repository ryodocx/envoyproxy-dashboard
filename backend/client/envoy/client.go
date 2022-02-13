package envoy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	// import for Unmarshal JSON to proto.Message
	admin "github.com/envoyproxy/go-control-plane/envoy/admin/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/common/key_value/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/common/matcher/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/grpc_credential/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/metrics/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/overload/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/tap/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/trace/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/data/cluster/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/data/core/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/data/dns/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/data/tap/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/file/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/grpc/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/open_telemetry/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/stream/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/wasm/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/cache/simple_http_cache/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/clusters/aggregate/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/clusters/dynamic_forward_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/clusters/redis/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/common/dynamic_forward_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/common/matching/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/common/ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/common/tap/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/brotli/compressor/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/brotli/decompressor/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/gzip/compressor/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/gzip/decompressor/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/common/dependency/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/common/fault/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/common/matcher/action/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/adaptive_concurrency/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/admission_control/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/alternate_protocols_cache/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/aws_lambda/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/aws_request_signing/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/bandwidth_limit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/buffer/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/cache/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/cdn_loop/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/composite/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/compressor/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/cors/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/csrf/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/decompressor/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/dynamic_forward_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/dynamo/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ext_authz/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ext_proc/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/fault/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_http1_bridge/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_http1_reverse_bridge/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_json_transcoder/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_stats/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_web/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/gzip/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/header_to_metadata/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/health_check/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ip_tagging/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/jwt_authn/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/kill_request/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/local_ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/lua/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/oauth2/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/on_demand/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/original_src/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/rbac/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/set_metadata/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/tap/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/wasm/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/http_inspector/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/original_dst/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/original_src/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/proxy_protocol/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/tls_inspector/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/client_ssl_auth/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/connection_limit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/direct_response/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/dubbo_proxy/router/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/dubbo_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/echo/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/ext_authz/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/local_ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/mongo_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/rbac/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/redis_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/sni_cluster/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/sni_dynamic_forward_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/filters/header_to_metadata/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/filters/ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/router/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/wasm/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/zookeeper_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/udp/dns_filter/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/udp/udp_proxy/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/formatter/metadata/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/formatter/req_without_query/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/health_checkers/redis/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/http/header_formatters/preserve_case/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/http/original_ip_detection/custom_header/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/http/original_ip_detection/xff/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/internal_redirect/allow_listed_routes/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/internal_redirect/previous_routes/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/internal_redirect/safe_cross_scheme/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/key_value/file_based/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/matching/common_inputs/environment_variable/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/matching/input_matchers/consistent_hashing/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/matching/input_matchers/ip/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/network/dns_resolver/apple/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/network/dns_resolver/cares/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/network/socket_interface/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/quic/crypto_stream/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/quic/proof_source/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/rate_limit_descriptors/expr/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/rbac/matchers/upstream_ip_port/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/request_id/uuid/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/resource_monitors/fixed_heap/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/resource_monitors/injected_resource/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/retry/host/omit_canary_hosts/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/retry/host/omit_host_metadata/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/retry/host/previous_hosts/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/retry/priority/previous_priorities/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/stat_sinks/graphite_statsd/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/stat_sinks/wasm/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/alts/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/proxy_protocol/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/quic/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/raw_buffer/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/s2a/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/starttls/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tap/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tcp_stats/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/generic/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/http/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/tcp/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/tcp/generic/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/wasm/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/watchdog/profile_action/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/cluster/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/event_reporting/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/extension/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/health/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/listener/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/load_stats/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/metrics/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/route/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/status/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/tap/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/service/trace/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/type/http/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/type/metadata/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/type/tracing/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/watchdog/v3"
	_ "istio.io/api/envoy/config/filter/http/alpn/v2alpha1"
	_ "istio.io/api/envoy/config/filter/http/authn/v2alpha1"
	_ "istio.io/api/envoy/config/filter/http/jwt_auth/v2alpha1"
	_ "istio.io/api/envoy/config/filter/network/metadata_exchange"
	_ "istio.io/api/envoy/config/filter/network/tcp_cluster_rewrite/v2alpha1"
	_ "istio.io/api/envoy/extensions/stackdriver/config/v1alpha1"
	_ "istio.io/api/envoy/extensions/stats"
)

type Config struct {
	EnvoyURL *url.URL
}

type Client struct {
	httpClient *http.Client
	envoyURL   url.URL
}

func New(conf Config) (*Client, error) {
	return &Client{
		httpClient: &http.Client{},
		envoyURL:   *conf.EnvoyURL,
	}, nil
}

type ConfigDump struct {
	BootstrapConfigDump    *admin.BootstrapConfigDump
	ClustersConfigDump     *admin.ClustersConfigDump
	ListenersConfigDump    *admin.ListenersConfigDump
	ScopedRoutesConfigDump *admin.ScopedRoutesConfigDump
	RoutesConfigDump       *admin.RoutesConfigDump
	// SecretsConfigDump      *admin.SecretsConfigDump
	EndpointsConfigDump *admin.EndpointsConfigDump
}

func (c *Client) GetConfigDump() (*ConfigDump, error) {

	resp, err := c.httpClient.Get(c.envoyURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status=%d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var (
		rawConfigDump *admin.ConfigDump = &admin.ConfigDump{} // for Unmarshal JSON
		configDump    *ConfigDump       = &ConfigDump{}       // for return
	)

	// Convert JSON to https://pkg.go.dev/github.com/envoyproxy/go-control-plane/envoy/admin/v3#ConfigDump
	if err := protojson.Unmarshal(body, rawConfigDump); err != nil {
		return nil, err
	} else {
		for _, c := range rawConfigDump.GetConfigs() {
			// Convert Any to each type
			switch typeURL := c.GetTypeUrl(); typeURL {
			case "type.googleapis.com/envoy.admin.v3.BootstrapConfigDump":
				// https://pkg.go.dev/github.com/envoyproxy/go-control-plane/envoy/admin/v3#BootstrapConfigDump
				// Do nothing for security concern
			case "type.googleapis.com/envoy.admin.v3.ClustersConfigDump":
				// https://pkg.go.dev/github.com/envoyproxy/go-control-plane/envoy/admin/v3#ClustersConfigDump
				var d admin.ClustersConfigDump
				if err := proto.Unmarshal(c.Value, &d); err != nil {
					return nil, err
				}
				configDump.ClustersConfigDump = &d
			case "type.googleapis.com/envoy.admin.v3.ListenersConfigDump":
				// https://pkg.go.dev/github.com/envoyproxy/go-control-plane/envoy/admin/v3#ListenersConfigDump
				var d admin.ListenersConfigDump
				if err := proto.Unmarshal(c.Value, &d); err != nil {
					return nil, err
				}
				configDump.ListenersConfigDump = &d
			case "type.googleapis.com/envoy.admin.v3.ScopedRoutesConfigDump":
				// https://pkg.go.dev/github.com/envoyproxy/go-control-plane/envoy/admin/v3#ScopedRoutesConfigDump
				var d admin.ScopedRoutesConfigDump
				if err := proto.Unmarshal(c.Value, &d); err != nil {
					return nil, err
				}
				configDump.ScopedRoutesConfigDump = &d
			case "type.googleapis.com/envoy.admin.v3.RoutesConfigDump":
				// https://pkg.go.dev/github.com/envoyproxy/go-control-plane/envoy/admin/v3#RoutesConfigDump
				var d admin.RoutesConfigDump
				if err := proto.Unmarshal(c.Value, &d); err != nil {
					return nil, err
				}
				configDump.RoutesConfigDump = &d
			case "type.googleapis.com/envoy.admin.v3.SecretsConfigDump":
				// https://pkg.go.dev/github.com/envoyproxy/go-control-plane/envoy/admin/v3#SecretsConfigDump
				// Do nothing for security concern
			case "type.googleapis.com/envoy.admin.v3.EndpointsConfigDump":
				// https://pkg.go.dev/github.com/envoyproxy/go-control-plane/envoy/admin/v3#EndpointsConfigDump
				// Never seen before
				var d admin.EndpointsConfigDump
				if err := proto.Unmarshal(c.Value, &d); err != nil {
					return nil, err
				}
				configDump.EndpointsConfigDump = &d
			default:
				return nil, fmt.Errorf("unexpected TypeURL: %s", typeURL)
			}
		}
	}
	return configDump, nil
}

func (c *Client) GetRouteConfigurations() ([]*route.RouteConfiguration, error) {
	d, err := c.GetConfigDump()
	if err != nil {
		return nil, err
	}

	var routes []*route.RouteConfiguration

	addRoute := func(b []byte) error {
		r := &route.RouteConfiguration{}
		if err := proto.Unmarshal(b, r); err != nil {
			return err
		}
		routes = append(routes, r)
		return nil
	}

	for _, v := range d.RoutesConfigDump.GetStaticRouteConfigs() {
		addRoute(v.RouteConfig.Value)
	}
	for _, v := range d.RoutesConfigDump.GetDynamicRouteConfigs() {
		addRoute(v.RouteConfig.Value)
	}

	return routes, nil
}
