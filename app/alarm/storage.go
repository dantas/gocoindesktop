package alarm

import (
	"encoding/json"
	"os"
)

type AlarmStorage interface {
	Save(alarms []Alarm) error
	Load() ([]Alarm, error)
}

func NewAlarmStorage(path string) AlarmStorage {
	return fileStorage(path)
}

type fileStorage string

type fileFormat struct {
	Alarms []fileAlarm `json:"alarms"`
}

type fileAlarm struct {
	Name       string  `json:"name"`
	LowerBound float64 `json:"lowerBound"`
	UpperBound float64 `json:"upperBound"`
	IsEnabled  bool    `json:"isEnabled"`
}

func (storage fileStorage) Save(alarms []Alarm) error {
	var file *os.File
	var e error

	if file, e = os.Create(string(storage)); e != nil {
		return e
	}

	defer file.Close()

	fileAlarms := make([]fileAlarm, 0, len(alarms))

	for _, alarm := range alarms {
		fileAlarms = append(fileAlarms, fileAlarm(alarm))
	}

	decoded := fileFormat{
		Alarms: fileAlarms,
	}

	encoder := json.NewEncoder(file)

	if e := encoder.Encode(&decoded); e != nil {
		return e
	}

	return nil
}

func (storage fileStorage) Load() ([]Alarm, error) {
	var file *os.File
	var e error

	if file, e = os.Open(string(storage)); e != nil {
		return nil, e
	}

	defer file.Close()

	var decoded fileFormat

	decoder := json.NewDecoder(file)

	if e = decoder.Decode(&decoded); e != nil {
		return nil, e
	}

	alarms := make([]Alarm, 0, len(decoded.Alarms))

	for _, fileAlarm := range decoded.Alarms {
		alarms = append(alarms, Alarm(fileAlarm))
	}

	return alarms, nil
}
