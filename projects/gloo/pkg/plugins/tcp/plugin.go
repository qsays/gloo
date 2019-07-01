package tcp

import (
	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoytcp "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/tcp_proxy/v2"
	envoyutil "github.com/envoyproxy/go-control-plane/pkg/util"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/tcp"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	translatorutil "github.com/solo-io/gloo/projects/gloo/pkg/translator"
)

const (
	// filter info
	pluginStage = plugins.PostInAuth
)

func NewPlugin() *Plugin {
	return &Plugin{}
}

var _ plugins.Plugin = new(Plugin)
var _ plugins.ListenerPlugin = new(Plugin)

type Plugin struct {
}

func (p *Plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *Plugin) ProcessListener(params plugins.Params, in *v1.Listener, out *envoyapi.Listener) error {
	tl := in.GetTcpListener()
	if tl == nil {
		return nil
	}
	if tl.ListenerPlugins == nil {
		return nil
	}
	tcpSettings := tl.ListenerPlugins.TcpProxySettings
	if tcpSettings == nil {
		return nil
	}
	for _, f := range out.FilterChains {
		for i, filter := range f.Filters {
			if filter.Name == envoyutil.TCPProxy {
				// get config
				var cfg envoytcp.TcpProxy
				err := translatorutil.ParseConfig(&filter, &cfg)
				// this should never error
				if err != nil {
					return err
				}

				copySettings(&cfg, tcpSettings)

				f.Filters[i], err = translatorutil.NewFilterWithConfig(envoyutil.TCPProxy, &cfg)
				// this should never error
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func copySettings(cfg *envoytcp.TcpProxy, tcpSettings *tcp.TcpProxySettings) {
	cfg.IdleTimeout = tcpSettings.IdleTimeout
	cfg.MaxConnectAttempts = tcpSettings.MaxConnectAttempts

}