//go:build darwin

package config

import (
	"github.com/k3s-io/k3s/pkg/cli/cmds"
	"github.com/k3s-io/k3s/pkg/daemons/config"
)

func applyContainerdOSSpecificConfig(nodeConfig *config.Node) error {
	return nil
}

func applyContainerdQoSClassConfigFileIfPresent(envInfo *cmds.Agent, containerdConfig *config.Containerd) error {
	return nil
}

func applyCRIDockerdOSSpecificConfig(nodeConfig *config.Node) error {
	return nil
}

func configureACL(file string) error {
	return nil
}
