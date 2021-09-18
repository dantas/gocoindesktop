package domain

import (
	"time"
)

type PeriodicUpdater interface {
	Channel() <-chan Coin
	SetInterval(interval time.Duration)
	Stop()
}

type periodicUpdater struct {
	timer   *time.Timer
	channel chan Coin
	enabled bool
}

func NewPeriodicUpdater(interval time.Duration) PeriodicUpdater {
	updater := periodicUpdater{
		timer:   time.NewTimer(interval),
		channel: make(chan Coin),
		enabled: true,
	}

	go func() {
		for updater.enabled {
			updater.timer.Reset(interval)

			<-updater.timer.C

			var coin Coin // Scan

			updater.channel <- coin
		}

		updater.timer.Stop()
		close(updater.channel)
	}()

	return &updater
}

func (u *periodicUpdater) Channel() <-chan Coin {
	return u.channel
}

func (u *periodicUpdater) SetInterval(interval time.Duration) {
	u.timer.Reset(interval)
}

func (u *periodicUpdater) Stop() {
	u.enabled = false
}
