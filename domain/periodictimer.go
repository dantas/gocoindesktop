package domain

import "time"

type PeriodicTimer interface {
	Tick() <-chan struct{}
	SetInterval(interval time.Duration)
	Destroy()
}
