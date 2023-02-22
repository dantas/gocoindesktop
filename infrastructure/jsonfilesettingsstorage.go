package infrastructure

import (
	"encoding/json"
	"os"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

type jsonFileFormat struct {
	Interval         int64
	ShowWindowOnOpen bool
}

type jsonFileSettingsStorage struct {
	path string
}

func NewJsonFileSettingsStorage(path string) domain.SettingsStorage {
	return jsonFileSettingsStorage{
		path: path,
	}
}

func (storage jsonFileSettingsStorage) Save(pref domain.Settings) error {
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

func (storage jsonFileSettingsStorage) Load() (domain.Settings, error) {
	var file *os.File
	var e error

	if file, e = os.Open(storage.path); e != nil {
		return domain.Settings{}, e
	}

	defer file.Close()

	var decoded jsonFileFormat

	decoder := json.NewDecoder(file)

	if e = decoder.Decode(&decoded); e != nil {
		return domain.Settings{}, e
	}

	return domain.Settings{
		Interval:         time.Duration(decoded.Interval),
		ShowWindowOnOpen: decoded.ShowWindowOnOpen,
	}, nil
}
