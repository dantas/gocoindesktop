package domain

import "errors"

type Alarm struct {
	Name       string
	LowerBound float64
	UpperBound float64
	IsEnabled  bool
}

type TriggeredAlarm struct {
	Alarm   Alarm
	Coin    Coin
	InRange bool
}

var (
	ErrSaveAlarm         = errors.New("error saving alarms")
	ErrLoadAlarm         = errors.New("error loading alarms")
	ErrLoadAlarmNotExist = errors.New("error persisted alarms do not exist")
)

type AlarmStorage interface {
	Save(alarms []Alarm) error
	Load() ([]Alarm, error)
}

// Find better name
type AlarmManager struct {
	storage AlarmStorage
	entries map[string]alarmAndStatus
}

type alarmAndStatus struct {
	alarm   Alarm
	inRange bool
}

func NewAlarmManager(storage AlarmStorage) *AlarmManager {
	return &AlarmManager{
		storage: storage,
		entries: make(map[string]alarmAndStatus),
	}
}

func (manager *AlarmManager) Load() error {
	alarms, err := manager.storage.Load()

	if err != nil && !errors.Is(err, ErrLoadAlarmNotExist) {
		return err
	}

	manager.entries = make(map[string]alarmAndStatus)

	for _, a := range alarms {
		manager.entries[a.Name] = alarmAndStatus{
			alarm: a,
		}
	}

	return nil
}

func (manager *AlarmManager) Alarms() []Alarm {
	alarms := make([]Alarm, 0, len(manager.entries))

	for _, entry := range manager.entries {
		alarms = append(alarms, entry.alarm)
	}

	return alarms
}

func (manager *AlarmManager) Set(alarm Alarm) error {
	manager.entries[alarm.Name] = alarmAndStatus{
		alarm:   alarm,
		inRange: manager.entries[alarm.Name].inRange,
	}

	return manager.save()
}

func (manager *AlarmManager) CheckAlarms(coins []Coin) []TriggeredAlarm {
	triggered := make([]TriggeredAlarm, 0)

	for _, coin := range coins {
		entry, exists := manager.entries[coin.Name]

		if !exists {
			continue
		}

		if !entry.alarm.IsEnabled {
			continue
		}

		isInRange := coin.Price >= entry.alarm.LowerBound && coin.Price <= entry.alarm.UpperBound
		update := false

		if entry.inRange && !isInRange {
			entry.inRange = false
			update = true
		}

		if !entry.inRange && isInRange {
			entry.inRange = true
			update = true
		}

		if !update {
			continue
		}

		manager.entries[coin.Name] = entry

		triggered = append(triggered, TriggeredAlarm{
			Alarm:   entry.alarm,
			Coin:    coin,
			InRange: entry.inRange,
		})
	}

	return triggered
}

func (manager *AlarmManager) save() error {
	alarms := make([]Alarm, 0, len(manager.entries))

	for _, v := range manager.entries {
		alarms = append(alarms, v.alarm)
	}

	return manager.storage.Save(alarms)
}
