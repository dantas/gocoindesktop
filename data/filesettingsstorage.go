package data

import (
	"encoding/json"
	"os"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

type fileFormat struct {
	Interval int64
}

type fileSettingsStorage struct {
	path string
}

func NewFileSettingsStorage(path string) domain.SettingsStorage {
	return fileSettingsStorage{
		path: path,
	}
}

func (storage fileSettingsStorage) Save(pref domain.Settings) error {
	var file *os.File
	var e error

	if file, e = os.Create(storage.path); e != nil {
		return e
	}

	defer file.Close()

	decoded := fileFormat{
		Interval: int64(pref.Interval),
	}

	encoder := json.NewEncoder(file)

	if e := encoder.Encode(&decoded); e != nil {
		return e
	}

	return nil
}

func (storage fileSettingsStorage) Load() (domain.Settings, error) {
	var file *os.File
	var e error

	if file, e = os.Open(storage.path); e != nil {
		return domain.DefaultSettings, e
	}

	defer file.Close()

	var decoded fileFormat

	decoder := json.NewDecoder(file)

	if e = decoder.Decode(&decoded); e != nil {
		return domain.DefaultSettings, e
	}

	return domain.Settings{
		Interval: time.Duration(decoded.Interval),
	}, nil
}
