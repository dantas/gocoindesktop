package coindesktop

import (
	"testing"
	"time"
)

func TestPeriodicUpdaterIsStopping(t *testing.T) {
	quitChan := make(chan struct{})

	updater := NewPeriodicCoinsUpdater(3 * time.Second)

	time.AfterFunc(10*time.Second, func() {
		t.Error("Stop() is not stopping the periodic updater")
		close(quitChan)
	})

	time.AfterFunc(6*time.Second, func() {
		updater.Stop()
	})

	go func() {
		for coin := range updater.Channel() {
			coin.Value = coin.Value * 1 // Doing nothing
		}
		close(quitChan)
	}()

	<-quitChan
}

func TestPeriodicUpdaterIsWaitingForTime(t *testing.T) {
	var updater = NewPeriodicCoinsUpdater(2 * time.Second)

	var startLoop = time.Now()

	<-updater.Channel()

	var elapsedTime = time.Since(startLoop)

	if elapsedTime < 2*time.Second {
		t.Errorf("Time between updates is smaller than what it should be")
	}
}

func TestPeriodicUpdaterChangesInterval(t *testing.T) {
	var updater = NewPeriodicCoinsUpdater(2 * time.Second)

	<-updater.Channel()

	updater.SetInterval(5 * time.Second)

	var startLoop = time.Now()

	<-updater.Channel()

	var elapsedTime = time.Since(startLoop)

	if elapsedTime < 5*time.Second {
		t.Errorf("Time between updates is smaller than what it should be")
	}
}
