package system

// import "runtime"

const defaultUnixPathEnv = "@TERMUX_PREFIX@/usr/local/sbin:@TERMUX_PREFIX@/usr/local/bin:@TERMUX_PREFIX@/usr/sbin:@TERMUX_PREFIX@/usr/bin:@TERMUX_PREFIX@/sbin:@TERMUX_PREFIX@/bin"

// DefaultPathEnv is unix style list of directories to search for
// executables. Each directory is separated from the next by a colon
// ':' character .
func DefaultPathEnv(platform string) string {
	if "linux" == "windows" {
		if platform != "linux" && LCOWSupported() {
			return defaultUnixPathEnv
		}
		// Deliberately empty on Windows containers on Windows as the default path will be set by
		// the container. Docker has no context of what the default path should be.
		return ""
	}
	return defaultUnixPathEnv
}
