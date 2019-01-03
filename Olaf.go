// Package olaf provides methods to generate unique ID using Twitter Snowflake algorithm.
package olaf

// Golang implementation of Twitter Snowflake.
// #author Thanh Nguyen <btnguyen2k@gmail.com>
// #version 0.1.0

import (
	"math/big"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// Twitter snowflake's epoch is set to 2019-01-01 00:00:00 UTC
// You may customize this to set a different epoch for your application.
const Epoch int64 = 1546300800000

const maxRadix = 36

type Olaf struct {
	NodeID     int64 // original node-id
	nodeId64   int64 // node-id  for 64-bit ids
	nodeId128  int64 // node-id  for 128-bit ids
	Epoch      int64 // Twitter snowflake's epoch
	SequenceId int64 // Twitter snowflake's sequence-id
	Timestamp  int64 // last 'touch' UNIX timestamp in milliseconds
}

// NewOlaf creates a new Olaf with default epoch.
func NewOlaf(nodeId int64) *Olaf {
	return NewOlafWithEpoch(nodeId, Epoch)
}

// NewOlafWithEpoch creates a new Olaf with custom epoch.
func NewOlafWithEpoch(nodeId int64, epoch int64) *Olaf {
	olaf := Olaf{}
	olaf.NodeID = nodeId
	olaf.nodeId64 = nodeId & MaskNodeId64
	olaf.nodeId128 = nodeId & MaskNodeId128
	olaf.Epoch = epoch
	olaf.SequenceId = 0
	olaf.Timestamp = 0
	return &olaf
}

const (
	MaskNodeId64     = 0x3FF  // 10 bits
	MaxSequenceId64  = 0x1FFF // 13 bits
	ShiftNodeId64    = 13
	ShiftTimestamp64 = 23

	MaskNodeId128     = 0xFFFFFFFFFFFF // 48 bits
	MaxSequenceId128  = 0xFFFF         // 16 bits
	ShiftNodeId128    = 16
	ShiftTimestamp128 = 64
)

/*----------------------------------------------------------------------*/

// UnixMilliseconds returns current UNIX timestamp in milliseconds.
func UnixMilliseconds() int64 {
	return time.Now().UnixNano() / 1000000
}

// WaitTillNextMillisec waits till clock moves to the next millisecond.
// Returns the "next" millisecond.
func WaitTillNextMillisec(currentMillisec int64) int64 {
	nextMillisec := UnixMilliseconds()
	for ; nextMillisec <= currentMillisec; nextMillisec = UnixMilliseconds() {
		runtime.Gosched()
	}
	return nextMillisec
}

// ExtractTime64 extracts time metadata from a 64-bit id.
func (o *Olaf) ExtractTime64(id64 uint64) time.Time {
	timestamp := id64>>ShiftTimestamp64 + uint64(o.Epoch)
	sec := timestamp / 1000
	nsec := (timestamp % 1000) * 1000000
	return time.Unix(int64(sec), int64(nsec))
}

// ExtractTime64Hex extracts time metadata from a 64-bit id in hex (base 16) format.
func (o *Olaf) ExtractTime64Hex(id64Hex string) time.Time {
	id64, _ := strconv.ParseUint(id64Hex, 16, 64)
	return o.ExtractTime64(id64)
}

// ExtractTime64Ascii extracts time metadata from a 64-bit id in ascii (base 36) format.
func (o *Olaf) ExtractTime64Ascii(id64Ascii string) time.Time {
	id64, _ := strconv.ParseUint(id64Ascii, maxRadix, 64)
	return o.ExtractTime64(id64)
}

// Id64 generates a 64-bit id.
func (o *Olaf) Id64() uint64 {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	timestamp := UnixMilliseconds()
	sequence := int64(0)
	for done := false; !done; {
		done = true
		for timestamp < o.Timestamp {
			timestamp = WaitTillNextMillisec(timestamp)
		}
		if timestamp == o.Timestamp {
			//increase sequence
			sequence = atomic.AddInt64(&o.SequenceId, 1)
			if sequence > MaxSequenceId64 {
				// reset sequence
				o.SequenceId = 0
				timestamp = WaitTillNextMillisec(timestamp)
				done = false
			}
		}
	}
	o.SequenceId = sequence
	o.Timestamp = timestamp
	result := ((timestamp - o.Epoch) << ShiftTimestamp64) | (o.nodeId64 << ShiftNodeId64) | sequence
	return uint64(result)
}

// Id64Hex generates a 64-bit id as a hex (base 16) string.
func (o *Olaf) Id64Hex() string {
	return strconv.FormatUint(o.Id64(), 16)
}

// Id64Ascii generates a 64-bit id as an ascii string (base 36).
func (o *Olaf) Id64Ascii() string {
	return strconv.FormatUint(o.Id64(), maxRadix)
}

/*----------------------------------------------------------------------*/

// ExtractTime128 extracts time metadata from a 128-bit id.
func (o *Olaf) ExtractTime128(id128 *big.Int) time.Time {
	timestamp := id128.Rsh(id128, ShiftTimestamp128).Int64()
	sec := timestamp / 1000
	nsec := (timestamp % 1000) * 1000000
	return time.Unix(sec, nsec)
}

// ExtractTime128Hex extracts time metadata from a 128-bit id in hex (base 16) format.
func (o *Olaf) ExtractTime128Hex(id128Hex string) time.Time {
	id128 := big.NewInt(0)
	id128.SetString(id128Hex, 16)
	return o.ExtractTime128(id128)
}

// ExtractTime128Ascii extracts time metadata from a 128-bit id in ascii (base 36) format.
func (o *Olaf) ExtractTime128Ascii(id128Ascii string) time.Time {
	id128 := big.NewInt(0)
	id128.SetString(id128Ascii, maxRadix)
	return o.ExtractTime128(id128)
}

// Id128 generates a 128-bit id.
func (o *Olaf) Id128() *big.Int {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	timestamp := UnixMilliseconds()
	sequence := int64(0)
	for done := false; !done; {
		done = true
		for timestamp < o.Timestamp {
			timestamp = WaitTillNextMillisec(timestamp)
		}
		if timestamp == o.Timestamp {
			//increase sequence
			sequence = atomic.AddInt64(&o.SequenceId, 1)
			if sequence > MaxSequenceId128 {
				// reset sequence
				o.SequenceId = 0
				timestamp = WaitTillNextMillisec(timestamp)
				done = false
			}
		}
	}
	o.SequenceId = sequence
	o.Timestamp = timestamp
	high := timestamp
	low := (o.nodeId128 << ShiftNodeId128) | sequence
	h := big.NewInt(high)
	h.Lsh(h, ShiftTimestamp128)
	return h.Add(h, big.NewInt(low))
}

// Id128Hex generates a 128-bit id as a hex (base 16) string.
func (o *Olaf) Id128Hex() string {
	return o.Id128().Text(16)
}

// Id128Ascii generates a 128-bit id as an ascii (base 36) string.
func (o *Olaf) Id128Ascii() string {
	return o.Id128().Text(maxRadix)
}
