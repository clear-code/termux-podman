package env

import (
	"github.com/containers/podman/v5/pkg/rootless"
	"github.com/containers/podman/v5/pkg/util"
)

func getRuntimeDir() (string, error) {
	if !rootless.IsRootless() {
		return "@TERMUX_PREFIX@/run", nil
	}
	return util.GetRootlessRuntimeDir()
}
