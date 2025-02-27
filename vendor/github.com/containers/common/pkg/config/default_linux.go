package config

import (
	"os"
)

func getDefaultCgroupsMode() string {
	return "enabled"
}

// getDefaultTmpDir for linux
func getDefaultTmpDir() string {
	// first check the TMPDIR env var
	if path, found := os.LookupEnv("TMPDIR"); found {
		return path
	}
	return "@TERMUX_PREFIX@/var/tmp"
}

func getDefaultLockType() string {
	return "shm"
}

func getLibpodTmpDir() string {
	return "@TERMUX_PREFIX@/var/run/libpod"
}

// getDefaultMachineVolumes returns default mounted volumes (possibly with env vars, which will be expanded)
func getDefaultMachineVolumes() []string {
	return []string{"$HOME:$HOME"}
}

func getDefaultComposeProviders() []string {
	return defaultUnixComposeProviders
}
