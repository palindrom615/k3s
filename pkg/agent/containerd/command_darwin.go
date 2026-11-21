//go:build darwin

package containerd

import "os/exec"

func addDeathSig(_ *exec.Cmd) {
}
