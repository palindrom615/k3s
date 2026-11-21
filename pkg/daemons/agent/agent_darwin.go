//go:build darwin

package agent

import (
	"github.com/k3s-io/k3s/pkg/daemons/config"
	"github.com/k3s-io/k3s/pkg/util"
	kubeletconfig "k8s.io/kubelet/config/v1beta1"
)

func kubeProxyArgs(cfg *config.Agent) map[string]string {
	return map[string]string{}
}

func kubeletArgsAndConfig(cfg *config.Agent) (map[string]string, *kubeletconfig.KubeletConfiguration, error) {
	return nil, nil, util.ErrUnsupportedPlatform
}
