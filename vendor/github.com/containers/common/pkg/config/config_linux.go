package config

import (
	selinux "github.com/opencontainers/selinux/go-selinux"
)

const (
	// OverrideContainersConfig holds the default config path overridden by the root user
	OverrideContainersConfig = "@TERMUX_PREFIX@/etc/" + _configPath

	// DefaultContainersConfig holds the default containers config path
	DefaultContainersConfig = "@TERMUX_PREFIX@/share/" + _configPath

	// DefaultSignaturePolicyPath is the default value for the
	// policy.json file.
	DefaultSignaturePolicyPath = "@TERMUX_PREFIX@/etc/containers/policy.json"
)

func selinuxEnabled() bool {
	return selinux.GetEnabled()
}

var defaultHelperBinariesDir = []string{
	"@TERMUX_PREFIX@/local/libexec/podman",
	"@TERMUX_PREFIX@/local/lib/podman",
	"@TERMUX_PREFIX@/libexec/podman",
	"@TERMUX_PREFIX@/lib/podman",
}
