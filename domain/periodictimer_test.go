package domain

import (
	"testing"
	"time"
)

func TestPeriodicTimerDestroyStopsTimer(t *testing.T) {
	timer := NewPeriodicTimer()

	timer.SetInterval(3 * time.Second)

	counter := 0

	go func() {
		for range timer.tick {
			counter += 1
		}
	}()

	timer.Destroy()

	time.Sleep(10 * time.Second)

	if counter != 0 {
		t.Error("Timer.Destroy() is not stopping the timer")
	}
}

func TestPeriodicTimerUsesProvidedInterval(t *testing.T) {
	timer := NewPeriodicTimer()

	timer.SetInterval(1 * time.Second)

	counter := 0

	go func() {
		for range timer.tick {
			counter += 1
		}
	}()

	// Tight timing, but it should work
	time.Sleep(10*time.Second + 900*time.Millisecond)

	if counter != 10 {
		t.Error("SetInterval is not using the provided interval", counter)
	}
}
