//go:build darwin

package containerd

import (
	"net"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/k3s-io/k3s/pkg/daemons/config"
	util3 "github.com/k3s-io/k3s/pkg/util"
	pkgerrors "github.com/pkg/errors"
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
	return util3.ErrUnsupportedPlatform
}

func Client(address string) (*containerd.Client, error) {
	addr, _, err := util.GetAddressAndDialer(address)
	if err != nil {
		return nil, err
	}
	return containerd.New(addr)
}

func OverlaySupported(root string) error {
	return pkgerrors.WithMessagef(util3.ErrUnsupportedPlatform, "overlayfs is not supported")
}

func FuseoverlayfsSupported(root string) error {
	return pkgerrors.WithMessagef(util3.ErrUnsupportedPlatform, "fuse-overlayfs is not supported")
}

func StargzSupported(root string) error {
	return pkgerrors.WithMessagef(util3.ErrUnsupportedPlatform, "stargz is not supported")
}
