package domain_test

import (
	"sync"
	"testing"
	"time"

	"github.com/dantas/gocoindesktop/domain"
	"github.com/dantas/gocoindesktop/system/alarm"
	"github.com/dantas/gocoindesktop/system/settings"
	"github.com/dantas/gocoindesktop/system/timer"
	"github.com/stretchr/testify/assert"
)

/*
	Test Application, which hold the contract for how our app should behave.
	By testing it, we test the whole domain package at once.

	Pending issues:
		- Improve error messages
*/

func TestCoinsAndNoAlarms(t *testing.T) {
	application, fixture := createTestFixture()

	// Arrange
	fixture.coins = []domain.Coin{
		{
			Name:  "Fist Coin",
			Price: 2000,
		},
		{
			Name:  "Second Coin",
			Price: 3500,
		},
		{
			Name:  "Third Coin",
			Price: 5000,
		},
	}

	// Act
	application.Start()
	coinsAndAlarms := syncCoinsAndAlarmsOnce(application)

	// Assert
	assert.Len(t, coinsAndAlarms, len(fixture.coins))
	for i := range coinsAndAlarms {
		assert.Equal(t, coinsAndAlarms[i].Coin, fixture.coins[i])
		assert.Nil(t, coinsAndAlarms[i].Alarm)
	}
}

func TestTimerIsUpdatingCoins(t *testing.T) {
	application, fixture := createTestFixture()

	// Arrange
	// First Step
	// Executed without timer. Start will do the job fetching coins for the first time.
	fixture.coins = []domain.Coin{
		{
			Name:  "Fist Coin",
			Price: 2000,
		},
		{
			Name:  "Second Coin",
			Price: 3500,
		},
		{
			Name:  "Third Coin",
			Price: 5000,
		},
	}

	// Act
	application.Start()
	firstCoinsAndAlarms := syncCoinsAndAlarmsOnce(application)

	// Assert
	assert.Len(t, firstCoinsAndAlarms, len(fixture.coins))
	for i := range firstCoinsAndAlarms {
		assert.Equal(t, firstCoinsAndAlarms[i].Coin, fixture.coins[i])
		assert.Nil(t, firstCoinsAndAlarms[i].Alarm)
	}

	// Arrange
	// Second Step
	// Executed after timer triggered.
	fixture.coins = []domain.Coin{
		{
			Name:  "Timer Fist Coin",
			Price: 12000,
		},
		{
			Name:  "Timer Second Coin",
			Price: 13500,
		},
		{
			Name:  "Timer Third Coin",
			Price: 15000,
		},
	}

	// Act
	fixture.timer.ForceTick()
	secondCoinsAndAlarms := syncCoinsAndAlarmsOnce(application)

	// Assert
	assert.Len(t, secondCoinsAndAlarms, len(fixture.coins))
	for i := range secondCoinsAndAlarms {
		assert.Equal(t, secondCoinsAndAlarms[i].Coin, fixture.coins[i])
		assert.Nil(t, secondCoinsAndAlarms[i].Alarm)
	}
}

func TestIfAlarmsAreCorrectlyAssociatedWithCorrectCoin(t *testing.T) {
	application, fixture := createTestFixture()

	// Arrange
	fixture.coins = []domain.Coin{
		{
			Name:  "First Coin",
			Price: 2000,
		},
		{
			Name:  "Second Coin",
			Price: 3500,
		},
		{
			Name:  "Third Coin",
			Price: 6000,
		},
	}

	fixture.alarmStorage.ToLoadAlarms = []domain.Alarm{
		{
			Name:       "First Coin",
			LowerBound: 1500,
			UpperBound: 3000,
			IsEnabled:  true,
		},
		{
			Name:       "Second Coin",
			LowerBound: 3000,
			UpperBound: 4500,
			IsEnabled:  false,
		},
		{
			Name:       "Third Coin",
			LowerBound: 4500,
			UpperBound: 5500,
			IsEnabled:  true,
		},
		{
			Name:       "Fourth Coin", // Should not be returned
			LowerBound: 6500,
			UpperBound: 7500,
			IsEnabled:  true,
		},
	}

	// Act
	application.Start()
	coinsAndAlarms := syncCoinsAndAlarmsOnce(application)

	// Assert
	assert.Len(t, coinsAndAlarms, len(fixture.coins))
	for i := range coinsAndAlarms {
		assert.NotNil(t, coinsAndAlarms[i].Alarm)
		assert.Equal(t, fixture.alarmStorage.ToLoadAlarms[i], *coinsAndAlarms[i].Alarm)
	}
}

func TestIfAlarmsAreTriggered(t *testing.T) {
	application, fixture := createTestFixture()

	// Arrange
	fixture.coins = []domain.Coin{
		{
			Name:  "First Coin",
			Price: 2000,
		},
		{
			Name:  "Second Coin",
			Price: 3500,
		},
		{
			Name:  "Third Coin",
			Price: 6000,
		},
	}

	fixture.alarmStorage.ToLoadAlarms = []domain.Alarm{
		{
			Name:       "First Coin",
			LowerBound: 1500,
			UpperBound: 3000,
			IsEnabled:  true,
		},
		{
			Name:       "Second Coin",
			LowerBound: 3000,
			UpperBound: 4500,
			IsEnabled:  false,
		},
		{
			Name:       "Third Coin",
			LowerBound: 4500,
			UpperBound: 5500,
			IsEnabled:  true,
		},
	}

	// Act
	application.Start()
	syncCoinsAndAlarmsOnce(application)

	// Assert
	// Check if alarm entered range
	triggeredAlarmEnteredRange := syncTriggeredAlarmsOnce(application)
	assert.Equal(t, fixture.alarmStorage.ToLoadAlarms[0], triggeredAlarmEnteredRange.Alarm)
	assert.True(t, triggeredAlarmEnteredRange.InRange)

	// Arrange
	// Check if alarm left range
	fixture.coins = []domain.Coin{
		{
			Name:  "First Coin",
			Price: 4000,
		},
	}

	// Act
	fixture.timer.ForceTick()
	syncCoinsAndAlarmsOnce(application)

	// Assert
	triggeredAlarmLeftRange := syncTriggeredAlarmsOnce(application)
	assert.Equal(t, fixture.alarmStorage.ToLoadAlarms[0], triggeredAlarmLeftRange.Alarm)
	assert.False(t, triggeredAlarmLeftRange.InRange)
}

func TestIfSetAlarmPersistsAlarm(t *testing.T) {
	application, fixture := createTestFixture()

	// Arrange
	newAlarm := domain.Alarm{
		Name:       "New Alarm",
		LowerBound: 222,
		UpperBound: 333,
		IsEnabled:  true,
	}

	// Act
	application.SetAlarm(newAlarm)

	// Assert
	assert.ElementsMatch(t, []domain.Alarm{newAlarm}, fixture.alarmStorage.SavedAlarms)
}

func TestIfSetSettingsPersistsSettingsAndChangesTimerInterval(t *testing.T) {
	application, fixture := createTestFixture()

	// Arrange
	newSettings := domain.Settings{
		Interval:         60 * time.Hour,
		ShowWindowOnOpen: false,
	}

	// Act
	application.SetSettings(newSettings)

	// Assert
	assert.Equal(t, newSettings, fixture.settingsStorage.SavedSettings)
	assert.Equal(t, newSettings.Interval, fixture.timer.SetIntervalValue)
}

// ============================================
// Helper methods and structs

type testFixture struct {
	timer           *timer.MockTimer
	settingsStorage settings.MockStorage
	coins           []domain.Coin
	alarmStorage    alarm.MockStorage
}

func createTestFixture() (*domain.Application, *testFixture) {
	fixture := testFixture{
		timer:           timer.NewMockTimer(),
		settingsStorage: settings.MockStorage{},
		alarmStorage:    alarm.MockStorage{},
	}

	coinSource := func() ([]domain.Coin, error) {
		return fixture.coins, nil
	}

	alarmManager := domain.NewAlarmManager(&fixture.alarmStorage)
	application := domain.NewApplication(fixture.timer, &fixture.settingsStorage, coinSource, alarmManager)

	go func() {
		for nothing := range application.Errors() {
			// Do nothing, only consume events
			nothing.Error()
		}
	}()

	return application, &fixture
}

func syncCoinsAndAlarmsOnce(application *domain.Application) []domain.CoinAndAlarm {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	var toReturn []domain.CoinAndAlarm

	go func() {
		for toReturn = range application.CoinsAndAlarms() {
			waitGroup.Done()
			break
		}
	}()

	waitGroup.Wait()

	return toReturn
}

func syncTriggeredAlarmsOnce(application *domain.Application) domain.TriggeredAlarm {
	wait := sync.WaitGroup{}
	wait.Add(1)

	var toReturn domain.TriggeredAlarm

	go func() {
		for toReturn = range application.TriggeredAlarms() {
			wait.Done()
			break
		}
	}()

	wait.Wait()

	return toReturn
}
