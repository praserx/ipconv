![](https://travis-ci.org/PraserX/ipconv.svg?branch=master)

# ipconv

This library provides simple conversion between `net.IP` and integer (`net.IP <--> int`).

I hope that it will serve you well.

## Example

```go
package main

import (
	"fmt"
    "net"
    "github.com/PraserX/ipconv"
)

func main() {
	var ip = net.ParseIP("192.168.1.1")
	fmt.Println(ipconv.IPv4ToInt(ip))
}
```