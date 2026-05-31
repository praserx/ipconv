// Package ipconv provides basic conversion between IP addresses
// representations. Converter does not require additional settings.
package ipconv

import (
	"encoding/binary"
	"errors"
	"math/big"
	"net"
	"strings"
)

var ErrInvalidIPAddress = errors.New("invalid ip address")
var ErrNotIPv4Address = errors.New("not an IPv4 address")
var ErrNotIPv6Address = errors.New("not an IPv6 address")

// IPv4ToInt converts IP address of version 4 from net.IP to uint32
// representation.
func IPv4ToInt(ipaddr net.IP) (uint32, error) {
	ip4 := ipaddr.To4()
	if ip4 == nil {
		return 0, ErrNotIPv4Address
	}
	return binary.BigEndian.Uint32(ip4), nil
}

// IPv6ToInt converts IP address of version 6 from net.IP to uint64 array
// representation. Return value contains high integer value on the first
// place and low integer value on second place.
func IPv6ToInt(ipaddr net.IP) ([2]uint64, error) {
	if ipaddr == nil {
		return [2]uint64{0, 0}, ErrInvalidIPAddress
	}

	ip16 := ipaddr.To16()
	if ip16 == nil {
		return [2]uint64{0, 0}, ErrNotIPv6Address
	}

	// Get two separates values of integer IP
	ip := [2]uint64{
		binary.BigEndian.Uint64(ip16[0:8]),  // IP high
		binary.BigEndian.Uint64(ip16[8:16]), // IP low
	}

	return ip, nil
}

// IPv6ToBigInt converts IP address of version 6 from net.IP to math big
// integer representation.
func IPv6ToBigInt(ipaddr net.IP) (*big.Int, error) {
	if ipaddr == nil {
		return nil, ErrInvalidIPAddress
	}

	ip16 := ipaddr.To16()
	if ip16 == nil {
		return nil, ErrNotIPv6Address
	}

	// Initialize value as bytes
	var ip big.Int
	ip.SetBytes(ip16)

	return &ip, nil
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

	// Direct zero-allocation write via standard library binary package.
	binary.BigEndian.PutUint64(ip[0:8], high)
	binary.BigEndian.PutUint64(ip[8:16], low)

	return ip
}

// BigIntToIPv6 converts IP address of version 6 from big integer to net.IP
// representation.
func BigIntToIPv6(ipaddr big.Int) net.IP {
	ip := make(net.IP, net.IPv6len)
	ipaddr.FillBytes(ip)
	return ip
}

// ParseIP implements extension of net.ParseIP. It returns additional
// information about IP address bytes length. In general, it works typically
// as standard net.ParseIP. So if IP is not valid, nil is returned.
func ParseIP(s string) (net.IP, int, error) {
	pip := net.ParseIP(s)
	if pip == nil {
		return nil, 0, ErrInvalidIPAddress
	}
	if strings.Contains(s, ":") {
		return pip, 16, nil
	}
	return pip, 4, nil
}
