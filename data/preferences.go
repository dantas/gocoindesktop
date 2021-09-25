package data

import (
	"encoding/json"
	"os"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

const fileName = "preferences.json"

type fileFormat struct {
	Interval int64
}

func SavePreferences(pref domain.Preferences) error {
	var file *os.File
	var e error

	if file, e = os.Create(fileName); e != nil {
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

func LoadPreferences() (domain.Preferences, error) {
	var file *os.File
	var e error

	if file, e = os.Open(fileName); e != nil {
		return domain.DefaultPreferences, e
	}

	defer file.Close()

	var decoded fileFormat

	decoder := json.NewDecoder(file)

	if e = decoder.Decode(&decoded); e != nil {
		return domain.DefaultPreferences, e
	}

	return domain.Preferences{
		Interval: time.Duration(decoded.Interval),
	}, nil
}
