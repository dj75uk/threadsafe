package threadsafe

import (
	"math/rand"
	"time"
)

func SleepRandom(minDuration time.Duration, maxDuration time.Duration) {
	deltaMS := (maxDuration - minDuration).Microseconds()
	duration := maxDuration + time.Duration(rand.Int63n(deltaMS))
	time.Sleep(duration)
}
