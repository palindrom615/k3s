//go:build darwin

package flannel

import (
	"context"
	"sync"

	"github.com/k3s-io/k3s/pkg/daemons/config"
	"github.com/k3s-io/k3s/pkg/util/errors"
)

const (
	BackendNone            = "none"
	BackendVXLAN           = "vxlan"
	BackendHostGW          = "host-gw"
	BackendWireguardNative = "wireguard-native"
	BackendTailscale       = "tailscale"
)

var (
	BaseAnnotation         = "flannel.alpha.coreos.com"
	ExternalIPv4Annotation = BaseAnnotation + "/public-ip-overwrite"
	ExternalIPv6Annotation = BaseAnnotation + "/public-ipv6-overwrite"
)

func Prepare(ctx context.Context, nodeConfig *config.Node) error {
	return errors.ErrUnsupportedPlatform
}

func Run(ctx context.Context, wg *sync.WaitGroup, nodeConfig *config.Node) error {
	return errors.ErrUnsupportedPlatform
}
