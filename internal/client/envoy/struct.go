package envoy

import (
	admin "github.com/envoyproxy/go-control-plane/envoy/admin/v3"
)

type ConfigDump struct {
	BootstrapConfigDump    *admin.BootstrapConfigDump
	ClustersConfigDump     *admin.ClustersConfigDump
	ListenersConfigDump    *admin.ListenersConfigDump
	ScopedRoutesConfigDump *admin.ScopedRoutesConfigDump
	RoutesConfigDump       *admin.RoutesConfigDump
	// SecretsConfigDump      *admin.SecretsConfigDump
	EndpointsConfigDump *admin.EndpointsConfigDump
}
