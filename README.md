**Since 10-Apr-2019: Olaf has been merged as a sub-module of [btnguyen2k/consu](https://github.com/btnguyen2k/consu)**

# Olaf

[![Go Report Card](https://goreportcard.com/badge/github.com/btnguyen2k/olaf)](https://goreportcard.com/report/github.com/btnguyen2k/olaf)
[![cover.run](https://cover.run/go/github.com/btnguyen2k/olaf.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fbtnguyen2k%2Folaf)
[![GoDoc](https://godoc.org/github.com/btnguyen2k/olaf?status.svg)](https://godoc.org/github.com/btnguyen2k/olaf)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

Golang implementation of Twitter Snowflake.

## Getting Started

### Install Package

```
go get github.com/btnguyen2k/olaf
```

### Usage

```go
package main

import (
    "fmt"
    "github.com/btnguyen2k/olaf"
)

func main() {
    // use default epoch
    o := olaf.NewOlaf(1981)

    //use custom epoch (note: epoch is in milliseconds)
    //o := olaf.NewOlafWithEpoch(103, 1546543604123)

    id64 := o.Id64()
    id64Hex := o.Id64Hex()
    id64Ascii := o.Id64Ascii()
    fmt.Println("ID 64-bit (int)   : ", id64, " / Timestamp: ", o.ExtractTime64(id64))
    fmt.Println("ID 64-bit (hex)   : ", id64Hex, " / Timestamp: ", o.ExtractTime64Hex(id64Hex))
    fmt.Println("ID 64-bit (ascii) : ", id64Ascii, " / Timestamp: ", o.ExtractTime64Ascii(id64Ascii))

    id128 := o.Id128()
    id128Hex := o.Id128Hex()
    id128Ascii := o.Id128Ascii()
    fmt.Println("ID 128-bit (int)  : ", id128.String(), " / Timestamp: ", o.ExtractTime128(id128))
    fmt.Println("ID 128-bit (hex)  : ", id128Hex, " / Timestamp: ", o.ExtractTime128Hex(id128Hex))
    fmt.Println("ID 128-bit (ascii): ", id128Ascii, " / Timestamp: ", o.ExtractTime128Ascii(id128Ascii))
}
```

## Document - GoDoc

See [GoDoc](https://godoc.org/github.com/btnguyen2k/olaf).


## History

Current version: `0.1.0`.

### 2019-01-03 - v0.1.0

First release:

- Generate 64-bit and 128-bit IDs:
  - Support formats: `integer`, `hex-string` (base 16) and `ascii-string` (base 36).
  - Custom epoch (64-bit ID only).
- Extract the time metadata from generated ID.
