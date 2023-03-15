package timer

import (
	"time"
)

type PeriodicTimer struct {
	done   chan any
	tick   chan struct{}
	ticker *time.Ticker
}

// Timer only starts after calling SetInterval
func NewPeriodicTimer() *PeriodicTimer {
	p := PeriodicTimer{
		done:   make(chan any),
		tick:   make(chan struct{}),
		ticker: time.NewTicker(1 * time.Minute),
	}

	p.ticker.Stop()

	go func() {
		for {
			select {
			case <-p.done:
				close(p.tick)
				return
			case <-p.ticker.C:
				p.tick <- struct{}{}
			}
		}
	}()

	return &p
}

func (p *PeriodicTimer) Tick() <-chan struct{} {
	return p.tick
}

func (p *PeriodicTimer) SetInterval(interval time.Duration) {
	p.ticker.Reset(interval)
}

func (p *PeriodicTimer) Destroy() {
	close(p.done)
	p.ticker.Stop()
}
