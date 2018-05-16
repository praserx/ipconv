// Package ipconv provides basic conversion between IP addresses
// representations. Converter does not require additional settings.
package ipconv

import (
	"encoding/binary"
	"math/big"
	"net"
)

// IPv4ToInt converts IP address of version 4 from net.IP to uint32
// representation.
func IPv4ToInt(ipaddr net.IP) uint32 {
	return binary.BigEndian.Uint32(ipaddr.To4())
}

// IPv6ToInt converts IP address of version 6 from net.IP to uint64 array
// representation. Return value contains high integer value on the first
// place and low integer value on second place.
func IPv6ToInt(ipaddr net.IP) [2]uint64 {
	// Get two separates values of integer IP
	ip := [2]uint64{
		binary.BigEndian.Uint64(ipaddr.To16()[0:8]),  // IP high
		binary.BigEndian.Uint64(ipaddr.To16()[8:16]), // IP low
	}

	return ip
}

// IPv6ToBigInt converts IP address of version 6 from net.IP to math big
// integer representation.
func IPv6ToBigInt(ipaddr net.IP) big.Int {
	// Define big int variable
	var ip big.Int

	// Initialize value as bytes
	ip.SetBytes(ipaddr)

	return ip
}

// IntToIPv4 converts IP address of version 4 from integer to net.IP
// representation.
func IntToIPv4(ipaddr uint32) net.IP {
	// Allocate 4 bytes IP
	ip := make(net.IP, 4)

	// Proceed conversion
	binary.BigEndian.PutUint32(ip, ipaddr)

	return ip
}

// IntToIPv6 converts IP address of version 6 from integer (high and low value)
// to net.IP representation.
func IntToIPv6(high, low uint64) net.IP {
	// Allocate 16 bytes IP
	ip := make(net.IP, 16)

	// Allocate 8 bytes arrays for IPs
	ipHigh := make([]byte, 8)
	ipLow := make([]byte, 8)

	// Proceed conversion
	binary.BigEndian.PutUint64(ipHigh, high)
	binary.BigEndian.PutUint64(ipLow, low)

	for i := 0; i < 16; i++ {
		if i < 8 {
			ip[i] = ipHigh[i]
		} else if i >= 8 {
			ip[i] = ipLow[i-8]
		}
	}

	return ip
}
