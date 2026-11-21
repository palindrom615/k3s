//go:build darwin

package rootlessports

import (
	"context"

	corev1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/core/v1"
)

func Register(ctx context.Context, serviceController corev1.ServiceController, enabled bool, httpsPort int) error {
	return nil
}
