package settings

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

type fileFormat struct {
	Interval         int64 `json:"interval"`
	ShowWindowOnOpen bool  `json:"showWindowOnOpen"`
}

type fileStorage string

func NewSettingsStorage(path string) domain.SettingsStorage {
	return fileStorage(path)
}

func (storage fileStorage) Save(pref domain.Settings) error {
	var file *os.File
	var e error

	if file, e = os.Create(string(storage)); e != nil {
		return errors.Join(domain.ErrSaveSettings, e)
	}

	defer file.Close()

	decoded := fileFormat{
		Interval:         int64(pref.Interval),
		ShowWindowOnOpen: pref.ShowWindowOnOpen,
	}

	encoder := json.NewEncoder(file)

	if e := encoder.Encode(&decoded); e != nil {
		return errors.Join(domain.ErrSaveSettings, e)
	}

	return nil
}

func (storage fileStorage) Load() (domain.Settings, error) {
	var file *os.File
	var e error

	if file, e = os.Open(string(storage)); e != nil {
		if os.IsNotExist(e) {
			e = nil
		} else {
			e = errors.Join(domain.ErrLoadSettings, e)
		}

		return domain.NewDefaultSettings(), e
	}

	defer file.Close()

	var decoded fileFormat

	decoder := json.NewDecoder(file)

	if e = decoder.Decode(&decoded); e != nil {
		return domain.NewDefaultSettings(), errors.Join(domain.ErrLoadSettings, e)
	}

	return domain.Settings{
		Interval:         time.Duration(decoded.Interval),
		ShowWindowOnOpen: decoded.ShowWindowOnOpen,
	}, nil
}
