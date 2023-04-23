package alarm

import (
	"encoding/json"
	"os"

	"github.com/dantas/gocoindesktop/domain"
)

func NewAlarmStorage(path string) domain.AlarmStorage {
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

func (storage fileStorage) Save(alarms []domain.Alarm) error {
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

	toEncode := fileFormat{
		Alarms: fileAlarms,
	}

	encoder := json.NewEncoder(file)

	if e := encoder.Encode(&toEncode); e != nil {
		return e
	}

	return nil
}

func (storage fileStorage) Load() ([]domain.Alarm, error) {
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

	alarms := make([]domain.Alarm, 0, len(decoded.Alarms))

	for _, fileAlarm := range decoded.Alarms {
		alarms = append(alarms, domain.Alarm(fileAlarm))
	}

	return alarms, nil
}
