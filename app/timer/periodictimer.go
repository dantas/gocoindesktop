package timer

import (
	"time"
)

type PeriodicTimer interface {
	Tick() <-chan struct{}
	SetInterval(interval time.Duration)
	Destroy()
}

type periodicTimer struct {
	done   chan any
	tick   chan struct{}
	ticker *time.Ticker
}

// Timer only starts after calling SetInterval
func NewPeriodicTimer() PeriodicTimer {
	p := periodicTimer{
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

func (p *periodicTimer) Tick() <-chan struct{} {
	return p.tick
}

func (p *periodicTimer) SetInterval(interval time.Duration) {
	p.ticker.Reset(interval)
}

func (p *periodicTimer) Destroy() {
	close(p.done)
	p.ticker.Stop()
}
