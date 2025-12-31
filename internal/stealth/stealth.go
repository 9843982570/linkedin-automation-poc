package stealth

import (
	"math/rand"
	"time"
)

func RandomDelay(minMs, maxMs int) {
	delay := rand.Intn(maxMs-minMs) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
