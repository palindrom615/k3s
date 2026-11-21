//go:build darwin && !no_embedded_executor

package embed

import (
	daemonconfig "github.com/k3s-io/k3s/pkg/daemons/config"
)

func platformKubeProxyArgs(nodeConfig *daemonconfig.Node) map[string]string {
	return map[string]string{}
}
