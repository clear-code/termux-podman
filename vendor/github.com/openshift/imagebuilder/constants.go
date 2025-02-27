package imagebuilder

const (
	// in docker/system
	NoBaseImageSpecifier = "scratch"

	// in docker/system
	defaultPathEnv = "@TERMUX_PREFIX@/usr/local/sbin:@TERMUX_PREFIX@/usr/local/bin:@TERMUX_PREFIX@/usr/sbin:@TERMUX_PREFIX@/usr/bin:@TERMUX_PREFIX@/sbin:@TERMUX_PREFIX@/bin"
)
