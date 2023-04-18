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
	Alarms []Alarm
}

func (storage fileStorage) Save(alarms []Alarm) error {
	var file *os.File
	var e error

	if file, e = os.Create(string(storage)); e != nil {
		return e
	}

	defer file.Close()

	decoded := fileFormat{
		Alarms: alarms,
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

	return decoded.Alarms, nil
}
