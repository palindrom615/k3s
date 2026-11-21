//go:build darwin

package containerd

import (
	"net"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/k3s-io/k3s/pkg/daemons/config"
	"github.com/k3s-io/k3s/pkg/util/errors"
	"k8s.io/cri-client/pkg/util"
)

func hostDirectory(host string) string {
	if host, port, err := net.SplitHostPort(host); err == nil && port != "" {
		return host + "_" + port + "_"
	}
	return host
}

func getContainerdArgs(cfg *config.Node) []string {
	return []string{"containerd", "-c", cfg.Containerd.Config}
}

func SetupContainerdConfig(cfg *config.Node) error {
	return errors.ErrUnsupportedPlatform
}

func Client(address string) (*containerd.Client, error) {
	addr, _, err := util.GetAddressAndDialer(address)
	if err != nil {
		return nil, err
	}
	return containerd.New(addr)
}

func OverlaySupported(root string) error {
	return errors.WithMessagef(errors.ErrUnsupportedPlatform, "overlayfs is not supported")
}

func FuseoverlayfsSupported(root string) error {
	return errors.WithMessagef(errors.ErrUnsupportedPlatform, "fuse-overlayfs is not supported")
}

func StargzSupported(root string) error {
	return errors.WithMessagef(errors.ErrUnsupportedPlatform, "stargz is not supported")
}

func NixSupported(root string) error {
	return errors.WithMessagef(errors.ErrUnsupportedPlatform, "nix is not supported")
}
