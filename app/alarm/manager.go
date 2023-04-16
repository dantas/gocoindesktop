package alarm

import "github.com/dantas/gocoindesktop/app/coin"

// TODO: Find a better name than manager
type AlarmManager struct {
	storage AlarmStorage
	entries map[string]entry
}

type entry struct {
	alarm   Alarm
	inRange bool
}

func NewAlarmManager(storage *AlarmStorage) *AlarmManager {
	return &AlarmManager{
		storage: *storage,
		entries: make(map[string]entry),
	}
}

func (manager *AlarmManager) Load() error {
	alarms, err := manager.storage.Load()

	if err != nil {
		return err
	}

	manager.entries = make(map[string]entry)

	for _, a := range alarms {
		manager.entries[a.Name] = entry{
			alarm: a,
		}
	}

	return nil
}

func (manager *AlarmManager) save() error {
	alarms := make([]Alarm, 0, len(manager.entries))

	for _, v := range manager.entries {
		alarms = append(alarms, v.alarm)
	}

	return manager.storage.Save(alarms)
}

func (manager *AlarmManager) Add(alarm Alarm) error {
	manager.entries[alarm.Name] = entry{
		alarm: alarm,
	}

	return manager.save()
}

func (manager *AlarmManager) Remove(alarm Alarm) error {
	delete(manager.entries, alarm.Name)

	return manager.save()
}

func (manager *AlarmManager) CheckAlarms(coins []coin.Coin) []TriggeredAlarm {
	triggered := make([]TriggeredAlarm, 0)

	for _, coin := range coins {
		entry, exists := manager.entries[coin.Name]

		if !exists {
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
