package settings

import (
	"encoding/json"
	"os"
	"time"
)

type jsonFileFormat struct {
	Interval         int64
	ShowWindowOnOpen bool
}

type jsonFileStorage struct {
	path string
}

func NewJsonFileStorage(path string) SettingsStorage {
	return jsonFileStorage{
		path: path,
	}
}

func (storage jsonFileStorage) Save(pref Settings) error {
	var file *os.File
	var e error

	if file, e = os.Create(storage.path); e != nil {
		return e
	}

	defer file.Close()

	decoded := jsonFileFormat{
		Interval:         int64(pref.Interval),
		ShowWindowOnOpen: pref.ShowWindowOnOpen,
	}

	encoder := json.NewEncoder(file)

	if e := encoder.Encode(&decoded); e != nil {
		return e
	}

	return nil
}

func (storage jsonFileStorage) Load() (Settings, error) {
	var file *os.File
	var e error

	if file, e = os.Open(storage.path); e != nil {
		return Settings{}, e
	}

	defer file.Close()

	var decoded jsonFileFormat

	decoder := json.NewDecoder(file)

	if e = decoder.Decode(&decoded); e != nil {
		return Settings{}, e
	}

	return Settings{
		Interval:         time.Duration(decoded.Interval),
		ShowWindowOnOpen: decoded.ShowWindowOnOpen,
	}, nil
}
