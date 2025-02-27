//go:build !windows

package ocipull

import (
	"path/filepath"

	"github.com/containers/common/pkg/config"
	_ "github.com/containers/storage/pkg/homedir"
)

func localPolicyOverwrites() []string {
	var dirs []string
	dirs = append(dirs, filepath.Join("@TERMUX_HOME@/.config", "containers", policyfile))
	dirs = append(dirs, config.DefaultSignaturePolicyPath)
	return dirs
}
