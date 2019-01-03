package olaf

import (
    "testing"
    "time"
)

func TestUnixMilliseconds(t *testing.T) {
    ms := UnixMilliseconds()
    now := time.Now()
    delta := now.UnixNano() - ms*1000000
    if delta < 0 || delta >= 1000000 {
        t.Errorf("TestUnixMilliseconds failed, ms: %d, now: %d, delta: %d.", ms, now.UnixNano(), delta)
    }
}

func TestWaitTillNextMillisec(t *testing.T) {
    start := time.Now()
    nextMs := WaitTillNextMillisec(start.UnixNano() / 1000000)
    end := time.Now()
    startMs := start.UnixNano() / 1000000
    endMs := end.UnixNano() / 1000000
    delta := nextMs - startMs
    if 0 >= delta {
        t.Errorf("Next milliseconds was incorrect, prevMs: %d, nextMs: %d, delta: %d", startMs, nextMs, delta)
    }
    if endMs < nextMs {
        t.Errorf("Next milliseconds must not greater than now, nextMs: %d, nowMs: %d.", nextMs, endMs)
    }
}

func TestOlaf_ExtractTime64(t *testing.T) {
    o := NewOlaf(1981)
    id64 := o.Id64()
    time := o.ExtractTime64(id64)
    t.Log("id [", id64, "] --> time [", time, "]")
}

func TestOlaf_ExtractTime64Hex(t *testing.T) {
    o := NewOlaf(1981)
    id64Hex := o.Id64Hex()
    time := o.ExtractTime64Hex(id64Hex)
    t.Log("id [", id64Hex, "] --> time [", time, "]")
}

func TestOlaf_ExtractTime64Ascii(t *testing.T) {
    o := NewOlaf(1981)
    id64Ascii := o.Id64Ascii()
    time := o.ExtractTime64Ascii(id64Ascii)
    t.Log("id [", id64Ascii, "] --> time [", time, "]")
}

func TestOlaf_ExtractTime128(t *testing.T) {
    o := NewOlaf(1981)
    id128 := o.Id128()
    id128Str := id128.String()
    time := o.ExtractTime128(id128)
    t.Log("id [", id128Str, "] --> time [", time, "]")
}

func TestOlaf_ExtractTime128Hex(t *testing.T) {
    o := NewOlaf(1981)
    id128Hex := o.Id128Hex()
    time := o.ExtractTime128Hex(id128Hex)
    t.Log("id [", id128Hex, "] --> time [", time, "]")
}

func TestOlaf_ExtractTime128Ascii(t *testing.T) {
    o := NewOlaf(1981)
    id128Ascii := o.Id128Ascii()
    time := o.ExtractTime128Ascii(id128Ascii)
    t.Log("id [", id128Ascii, "] --> time [", time, "]")
}
