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
