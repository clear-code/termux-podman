//go:build !windows && !solaris && !darwin
// +build !windows,!solaris,!darwin

package dbus

import (
	"os"
)

const defaultSystemBusAddress = "unix:path=@TERMUX_PREFIX@/var/run/dbus/system_bus_socket"

func getSystemBusPlatformAddress() string {
	address := os.Getenv("DBUS_SYSTEM_BUS_ADDRESS")
	if address != "" {
		return address
	}
	return defaultSystemBusAddress
}
