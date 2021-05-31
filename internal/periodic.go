package coindesktop

import (
	"time"

	"github.com/dantas/gocoindesktop/internal/scrapper"
)

type PeriodicCoinUpdater interface {
	Channel() <-chan scrapper.Coin
	SetInterval(interval time.Duration)
	Stop()
}

type periodicCoinUpdater struct {
	timer   *time.Timer
	channel chan scrapper.Coin
	enabled bool
}

func NewPeriodicCoinsUpdater(interval time.Duration) PeriodicCoinUpdater {
	updater := periodicCoinUpdater{
		timer:   time.NewTimer(interval),
		channel: make(chan scrapper.Coin),
		enabled: true,
	}

	go func() {
		for updater.enabled {
			updater.timer.Reset(interval)

			<-updater.timer.C

			var coin scrapper.Coin // Scan

			updater.channel <- coin
		}

		updater.timer.Stop()
		close(updater.channel)
	}()

	return &updater
}

func (u *periodicCoinUpdater) Channel() <-chan scrapper.Coin {
	return u.channel
}

func (u *periodicCoinUpdater) SetInterval(interval time.Duration) {
	u.timer.Reset(interval)
}

func (u *periodicCoinUpdater) Stop() {
	u.enabled = false
}
