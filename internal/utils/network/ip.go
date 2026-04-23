package network

import (
	"fmt"
	"net"
)

// DetectPrimaryIPv4 returns the first non-loopback IPv4 address
// found on an active network interface.
func DetectPrimaryIPv4() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("failed to list network interfaces: %w", err)
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ipv4 := ip.To4()
			if ipv4 == nil {
				continue
			}

			return ipv4.String(), nil
		}
	}

	return "", fmt.Errorf("no active non-loopback IPv4 address found")
}
