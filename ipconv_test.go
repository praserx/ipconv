package ipconv

import (
	"math/big"
	"net"
	"testing"
)

func TestIPv4ToInt(t *testing.T) {
	for _, c := range []struct {
		in   net.IP
		want uint32
	}{
		{net.ParseIP("192.168.1.1"), 3232235777},
		{net.ParseIP("0.0.0.0"), 0},
		{net.ParseIP("8.8.8.8"), 134744072},
		{net.ParseIP("255.255.255.255"), 4294967295},
	} {
		got := IPv4ToInt(c.in)
		if got != c.want {
			t.Errorf("IPv4ToInt(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIPv6ToInt(t *testing.T) {
	for _, c := range []struct {
		in   net.IP
		want [2]uint64
	}{
		{net.ParseIP("0000:0000:0000:0000:0000:0000:0000:0000"), [2]uint64{0, 0}},
		{net.ParseIP("0000:0000:0000:0000:0000:0000:0000:1"), [2]uint64{0, 1}},
		{net.ParseIP("2001:4860:4860::8888"), [2]uint64{2306204062558715904, 34952}},
	} {
		got := IPv6ToInt(c.in)
		if got != c.want {
			t.Errorf("IPv6ToInt(%q) == %q, want %q", c.in.To16(), got, c.want)
		}
	}
}

func TestIPv6ToBigInt(t *testing.T) {
	for _, c := range []struct {
		in   net.IP
		want *big.Int
	}{
		{net.ParseIP("0000:0000:0000:0000:0000:0000:0000:0000"), GetBigInt("0")},
		{net.ParseIP("0000:0000:0000:0000:0000:0000:0000:0001"), GetBigInt("1")},
		{net.ParseIP("2001:0:0:0:0:ffff:c0a8:101"), GetBigInt("42540488161975842760550637899214225665")},
		{net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"), GetBigInt("340282366920938463463374607431768211455")},
	} {
		got := IPv6ToBigInt(c.in)
		if got.Cmp(c.want) != 0 {
			t.Errorf("IPv6ToInt(%q) == %q, want %q", c.in.To16(), got, c.want)
		}
	}
}

func TestIntToIPv4(t *testing.T) {
	for _, c := range []struct {
		in   uint32
		want net.IP
	}{
		{3232235777, net.ParseIP("192.168.1.1")},
		{0, net.ParseIP("0.0.0.0")},
		{134744072, net.ParseIP("8.8.8.8")},
		{4294967295, net.ParseIP("255.255.255.255")},
	} {
		got := IntToIPv4(c.in)
		if !got.Equal(c.want) {
			t.Errorf("IntToIPv4(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIntToIPv6(t *testing.T) {
	for _, c := range []struct {
		in   [2]uint64
		want net.IP
	}{
		{[2]uint64{0, 0}, net.ParseIP("0000:0000:0000:0000:0000:0000:0000:0000")},
		{[2]uint64{0, 1}, net.ParseIP("0000:0000:0000:0000:0000:0000:0000:0001")},
		{[2]uint64{1, 0}, net.ParseIP("0000:0000:0000:0001:0000:0000:0000:0000")},
		{[2]uint64{2306204062558715904, 34952}, net.ParseIP("2001:4860:4860::8888")},
		{[2]uint64{0, 18446744073709551615}, net.ParseIP("0000:0000:0000:0000:ffff:ffff:ffff:ffff")},
		{[2]uint64{18446744073709551615, 18446744073709551615}, net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")},
	} {
		got := IntToIPv6(c.in[0], c.in[1])
		if !got.Equal(c.want) {
			t.Errorf("IntToIPv6(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestBigIntToIPv6(t *testing.T) {
	for _, c := range []struct {
		in   *big.Int
		want net.IP
	}{
		{GetBigInt("0"), net.ParseIP("0000:0000:0000:0000:0000:0000:0000:0000")},
		{GetBigInt("1"), net.ParseIP("0000:0000:0000:0000:0000:0000:0000:0001")},
		{GetBigInt("42540488161975842760550637899214225665"), net.ParseIP("2001:0:0:0:0:ffff:c0a8:101")},
		{GetBigInt("340282366920938463463374607431768211455"), net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")},
	} {
		got := BigIntToIPv6(*c.in)
		if !got.Equal(c.want) {
			t.Errorf("BigIntToIPv6(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func GetBigInt(bi string) *big.Int {
	var bigInt = new(big.Int)
	bigInt.SetString(bi, 10)
	return bigInt
}
