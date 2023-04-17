package settings

import (
	"encoding/json"
	"os"
	"time"
)

type fileFormat struct {
	Interval         int64
	ShowWindowOnOpen bool
}

type fileStorage string

func NewSettingsStorage(path string) SettingsStorage {
	return fileStorage(path)
}

func (storage fileStorage) Save(pref Settings) error {
	var file *os.File
	var e error

	if file, e = os.Create(string(storage)); e != nil {
		return e
	}

	defer file.Close()

	decoded := fileFormat{
		Interval:         int64(pref.Interval),
		ShowWindowOnOpen: pref.ShowWindowOnOpen,
	}

	encoder := json.NewEncoder(file)

	if e := encoder.Encode(&decoded); e != nil {
		return e
	}

	return nil
}

func (storage fileStorage) Load() (Settings, error) {
	var file *os.File
	var e error

	if file, e = os.Open(string(storage)); e != nil {
		return newDefaultSettings(), e
	}

	defer file.Close()

	var decoded fileFormat

	decoder := json.NewDecoder(file)

	if e = decoder.Decode(&decoded); e != nil {
		return newDefaultSettings(), e
	}

	return Settings{
		Interval:         time.Duration(decoded.Interval),
		ShowWindowOnOpen: decoded.ShowWindowOnOpen,
	}, nil
}
