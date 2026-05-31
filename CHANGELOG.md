# v1.3.0
## 31-05-2026

1. [](#new)
    * Fix crash in IPv6ToInt and IPv6ToBigInt when given nil or invalid IP addresses
    * Cache To4() result in IPv4ToInt to avoid scanning twice
    * Write directly to the target slice in IntToIPv6 to eliminate extra allocations
    * Use big.Int.FillBytes in BigIntToIPv6 for cleaner, faster conversions
    * Bump Go version to 1.22 in go.mod
    * Update GitHub Actions configuration to test against Go 1.21, 1.22, and 1.23
2. [](#bugfix)
    * Fix typos ("addres" -> "address") in ErrNotIPv4Address and ErrNotIPv6Address messages
    * Fix compiling and runtime errors in README.md example code

# v1.2.2
## 25-07-2024

1. [](#new)
    * Added missing CHANGELOG.md
    * go.mod version bump to 1.22.5

# v1.2.1
## 22-11-2021

1. [](#bugfix)
    * Github username changed (PraserX --> praserx)

# v1.2.0
## 16-11-2021

1. [](#new)
    * Fixed README.md
2. [](#bugfix)
    * Added codeql-analysis.yml

# v1.1.0
## 23-04-2019

1. [](#new)
    * Added function ParseIP

# v1.0.1
## 3-12-2018

1. [](#new)
    * Added big integer support

# v1.0.0
## 23-11-2018

1. [](#new)
    * Initial release