package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	_listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	_route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	_hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
)

// GET /routes
func (s *server) routes(w http.ResponseWriter, r *http.Request) {

	e := func(err error) bool {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%v", err)
			return true
		}
		return false
	}

	dump, err := s.conf.Client.GetConfigDump()
	if e(err) {
		return
	}

	if false {
		response := map[string]any{}
		// listener
		listener := _listener.Listener{}
		err = dump.ListenersConfigDump.DynamicListeners[0].ActiveState.Listener.UnmarshalTo(&listener)
		if e(err) {
			return
		}
		response["listener"] = listener

		hcm := _hcm.HttpConnectionManager{}
		listener.FilterChains[0].Filters[0].GetTypedConfig().UnmarshalTo(&hcm)
		err = listener.FilterChains[0].Filters[0].GetTypedConfig().MarshalFrom(&hcm)
		if e(err) {
			return
		}
		response["hcm"] = hcm
		response["route_name"] = hcm.GetRds().GetRouteConfigName()

		// route
		var route *_route.RouteConfiguration
		for _, v := range dump.RoutesConfigDump.DynamicRouteConfigs {
			r := _route.RouteConfiguration{}
			if e(v.RouteConfig.UnmarshalTo(&r)) {
				return
			}
			if r.Name == response["route_name"] {
				route = &r
				break
			}
		}
		response["route"] = route
	}

	route := _route.RouteConfiguration{}
	if e(dump.RoutesConfigDump.DynamicRouteConfigs[0].RouteConfig.UnmarshalTo(&route)) {
		return
	}
	// response["route"] = route

	filterDomains := func(domains []string) []string {
		r := []string{}
		for _, v := range domains {
			if !strings.Contains(v, ":*") {
				r = append(r, v)
			}
		}
		return r
	}
	path := func(m *_route.RouteMatch) string {

		r := func(s string) (string, bool) {
			if len(s) > 0 {
				return s, true
			}
			return s, false
		}

		if s, ok := r(m.GetPath()); ok {
			return s
		} else if s, ok := r(m.GetPathSeparatedPrefix()); ok {
			return path.Join(s, "*")
		} else if s, ok := r(m.GetPathTemplate()); ok {
			return s
		} else if s, ok := r(m.GetPrefix()); ok {
			return s + "*"
		} else if s, ok := r(m.GetSafeRegex().GetRegex()); ok {
			return "/" + s + "/"
		}

		return "parse error"
	}

	//	*Route_Route
	//	*Route_Redirect
	//	*Route_DirectResponse
	//	*Route_FilterAction
	//	*Route_NonForwardingAction
	action := func(r *_route.Route) (actionType, actionValue string) {
		switch v := r.Action.(type) {
		case *_route.Route_Route:
			if s := r.Decorator.String(); strings.HasPrefix(s, `operation:`) {
				return "proxy", strings.TrimSuffix(strings.Trim(strings.TrimPrefix(s, `operation:`), `"`), "/*")
			}
			return "proxy", v.Route.GetCluster()
		case *_route.Route_Redirect:
			return "redirect", v.Redirect.GetHostRedirect() + v.Redirect.GetPathRedirect()
		case *_route.Route_DirectResponse:
			code := int(v.DirectResponse.Status)
			return "direct_response", fmt.Sprintf("%d %s", code, http.StatusText(code))
		case *_route.Route_FilterAction:
		case *_route.Route_NonForwardingAction:
		}
		return "parse error", ""
	}

	response := []any{}

	for _, v := range route.VirtualHosts {

		r := []string{}
		for _, v := range v.Routes {
			actionType, actionValue := action(v)
			r = append(r, fmt.Sprintf("%s → (%s) %s", path(v.Match), actionType, actionValue))
		}

		response = append(response, map[string]any{
			"domains": filterDomains(v.Domains),
			"routes":  r,
		})
	}

	// response
	b, err := json.MarshalIndent(response, "", "    ")
	if e(err) {
		return
	}

	fmt.Fprint(w, string(b))
	w.Header().Set("content-type", "application/json")
}
