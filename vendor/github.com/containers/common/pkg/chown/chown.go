package chown

import (
	"os"
	"os/user"
	"path/filepath"

	_ "github.com/containers/storage/pkg/homedir"
)

// DangerousHostPath validates if a host path is dangerous and should not be modified
func DangerousHostPath(path string) (bool, error) {
	excludePaths := map[string]bool{
		"/":           true,
		"@TERMUX_PREFIX@/bin":        true,
		"/boot":       true,
		"/dev":        true,
		"@TERMUX_PREFIX@/etc":        true,
		"@TERMUX_PREFIX@/etc/passwd": true,
		"@TERMUX_PREFIX@/etc/pki":    true,
		"@TERMUX_PREFIX@/etc/shadow": true,
		"@TERMUX_PREFIX@/home":       true,
		"@TERMUX_PREFIX@/lib":        true,
		"@TERMUX_PREFIX@/lib64":      true,
		"@TERMUX_PREFIX@/media":      true,
		"@TERMUX_PREFIX@/opt":        true,
		"/proc":       true,
		"/root":       true,
		"@TERMUX_PREFIX@/run":        true,
		"@TERMUX_PREFIX@/sbin":       true,
		"@TERMUX_PREFIX@/srv":        true,
		"/sys":        true,
		"@TERMUX_PREFIX@/tmp":        true,
		"@TERMUX_PREFIX@/usr":        true,
		"@TERMUX_PREFIX@/var":        true,
		"@TERMUX_PREFIX@/var/lib":    true,
		"@TERMUX_PREFIX@/var/log":    true,
	}

	if home := "@TERMUX_HOME@"; home != "" {
		excludePaths[home] = true
	}

	if sudoUser := os.Getenv("SUDO_USER"); sudoUser != "" {
		if usr, err := user.Lookup(sudoUser); err == nil {
			excludePaths[usr.HomeDir] = true
		}
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return true, err
	}

	realPath, err := filepath.EvalSymlinks(absPath)
	if err != nil {
		return true, err
	}

	if excludePaths[realPath] {
		return true, nil
	}

	return false, nil
}
