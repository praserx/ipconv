# ipconv

This library provides simple conversion between `net.IP` and integer (`net.IP <--> int`). As new feature, library now contains extension of `net.ParseIP` which returns also byte length of IP address on input.

I hope it will serve you well.

## Example

```go
package main

import (
    "fmt"
    "github.com/praserx/ipconv"
)

func main() {
    ip, version, err := ipconv.ParseIP("192.168.1.1")
    if err == nil && version == 4 {
        val, _ := ipconv.IPv4ToInt(ip)
        fmt.Println(val)
    }
}
```
