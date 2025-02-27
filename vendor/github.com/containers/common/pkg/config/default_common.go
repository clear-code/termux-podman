//go:build !freebsd && !netbsd

package config

// DefaultInitPath is the default path to the container-init binary.
var DefaultInitPath = "@TERMUX_PREFIX@/usr/libexec/podman/catatonit"
