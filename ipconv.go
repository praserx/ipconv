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
func IPv6ToBigInt(ipaddr net.IP) *big.Int {
	var ip big.Int

	// Initialize value as bytes
	ip.SetBytes(ipaddr)

	return &ip
}

// IntToIPv4 converts IP address of version 4 from integer to net.IP
// representation.
func IntToIPv4(ipaddr uint32) net.IP {
	ip := make(net.IP, net.IPv4len)

	// Proceed conversion
	binary.BigEndian.PutUint32(ip, ipaddr)

	return ip
}

// IntToIPv6 converts IP address of version 6 from integer (high and low value)
// to net.IP representation.
func IntToIPv6(high, low uint64) net.IP {
	ip := make(net.IP, net.IPv6len)

	// Allocate 8 bytes arrays for IPs
	ipHigh := make([]byte, 8)
	ipLow := make([]byte, 8)

	// Proceed conversion
	binary.BigEndian.PutUint64(ipHigh, high)
	binary.BigEndian.PutUint64(ipLow, low)

	for i := 0; i < net.IPv6len; i++ {
		if i < 8 {
			ip[i] = ipHigh[i]
		} else if i >= 8 {
			ip[i] = ipLow[i-8]
		}
	}

	return ip
}

// BigIntToIPv6 converts IP address of version 6 from big integer to net.IP
// representation.
func BigIntToIPv6(ipaddr big.Int) net.IP {
	ip := make(net.IP, net.IPv6len)

	ipBytes := ipaddr.Bytes()
	ipBytesLen := len(ipBytes)

	for i := 0; i < net.IPv6len; i++ {
		if i < net.IPv6len-ipBytesLen {
			ip[i] = 0x0
		} else {
			ip[i] = ipBytes[ipBytesLen-net.IPv6len+i]
		}
	}

	return ip
}
