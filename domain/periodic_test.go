package domain

import (
	"testing"
	"time"
)

func TestPeriodicUpdaterIsStopping(t *testing.T) {
	quitChan := make(chan struct{})

	updater := NewPeriodicUpdater(3 * time.Second)

	isStopped := false

	time.AfterFunc(10*time.Second, func() {
		if !isStopped {
			t.Error("Stop() is not stopping the periodic updater")
			close(quitChan)
		}
	})

	time.AfterFunc(6*time.Second, func() {
		updater.Stop()
	})

	go func() {
		for coin := range updater.Channel() {
			coin.Price = coin.Price * 1 // Doing nothing
		}
		isStopped = true
		close(quitChan)
	}()

	<-quitChan
}

func TestPeriodicUpdaterIsWaitingForTime(t *testing.T) {
	var updater = NewPeriodicUpdater(2 * time.Second)

	var startLoop = time.Now()

	<-updater.Channel()

	var elapsedTime = time.Since(startLoop)

	if elapsedTime < 2*time.Second {
		t.Errorf("Time between updates is smaller than what it should be")
	}
}

func TestPeriodicUpdaterChangesInterval(t *testing.T) {
	var updater = NewPeriodicUpdater(2 * time.Second)

	<-updater.Channel()

	updater.SetInterval(5 * time.Second)

	var startLoop = time.Now()

	<-updater.Channel()

	var elapsedTime = time.Since(startLoop)

	if elapsedTime < 5*time.Second {
		t.Errorf("Time between updates is smaller than what it should be")
	}
}