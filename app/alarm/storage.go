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
	return &JsonFileAlarmStorage{
		path: path,
	}
}

type JsonFileAlarmStorage struct {
	path string
}

type jsonFileFormat struct {
	Alarms []Alarm
}

func (storage *JsonFileAlarmStorage) Save(alarms []Alarm) error {
	var file *os.File
	var e error

	if file, e = os.Create(storage.path); e != nil {
		return e
	}

	defer file.Close()

	decoded := jsonFileFormat{
		Alarms: alarms,
	}

	encoder := json.NewEncoder(file)

	if e := encoder.Encode(&decoded); e != nil {
		return e
	}

	return nil
}

func (storage *JsonFileAlarmStorage) Load() ([]Alarm, error) {
	var file *os.File
	var e error

	if file, e = os.Open(storage.path); e != nil {
		return nil, e
	}

	defer file.Close()

	var decoded jsonFileFormat

	decoder := json.NewDecoder(file)

	if e = decoder.Decode(&decoded); e != nil {
		return nil, e
	}

	return decoded.Alarms, nil
}
