# Olaf

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

## Methods

`func NewOlaf(nodeId int64) *Olaf`

Create a new `Olaf` with default epoch.
The default epoch is `1546300800000` which is UNIX timestamp in milliseconds of `2019-01-01 00:00:00 UTC`.

`func NewOlafWithEpoch(nodeId int64, epoch int64) *Olaf`

Create a new Olaf with custom epoch.

`func UnixMilliseconds() int64`

Return current UNIX timestamp in milliseconds.

`func WaitTillNextMillisec(currentMillisec int64) int64`

Wait till clock moves to the next millisecond and return it.

`func (o *Olaf) ExtractTime64(id64 uint64) time.Time`

Extract time metadata from a 64-bit id.

`func (o *Olaf) ExtractTime64Hex(id64Hex string) time.Time`

Extracts time metadata from a 64-bit id in hex (base 16) format.

`func (o *Olaf) ExtractTime64Ascii(id64Ascii string) time.Time`

Extract time metadata from a 64-bit id in ascii (base 36) format.

`func (o *Olaf) Id64() uint64`

Generate a 64-bit id.

`func (o *Olaf) Id64Hex() string`

Generate a 64-bit id as a hex (base 16) string.

`func (o *Olaf) Id64Ascii() string`

Generate a 64-bit id as an ascii string (base 36).

`func (o *Olaf) ExtractTime128(id128 *big.Int) time.Time`

Extract time metadata from a 128-bit id.

`func (o *Olaf) ExtractTime128Hex(id128Hex string) time.Time`

Extract time metadata from a 128-bit id in hex (base 16) format.

`func (o *Olaf) ExtractTime128Ascii(id128Ascii string) time.Time`

Extract time metadata from a 128-bit id in ascii (base 36) format.

`func (o *Olaf) Id128() *big.Int`

Generate a 128-bit id.

`func (o *Olaf) Id128Hex() string`

Generate a 128-bit id as a hex (base 16) string.

`func (o *Olaf) Id128Ascii() string`

Generate a 128-bit id as an ascii (base 36) string.

## History

Current version: `0.1.0`.

### 2019-01-03 - v0.1.0

First release:

- Generate 64-bit and 128-bit IDs:
  - Support formats: `integer`, `hex-string` (base 16) and `ascii-string` (base 36).
  - Custom epoch (64-bit ID only).
- Extract the time metadata from generated ID.
