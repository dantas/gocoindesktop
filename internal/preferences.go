package coindesktop

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const fileName = "preferences.json"
const defaultInterval = 5 * time.Minute

type preferencesFormat struct {
	Interval int64
}

func GetPeriodicInterval() time.Duration {
	var file *os.File
	var e error

	if file, e = os.Open(fileName); e != nil {
		return defaultInterval
	}

	defer file.Close()

	var decoded preferencesFormat

	decoder := json.NewDecoder(file)
	if e = decoder.Decode(&decoded); e != nil {
		return defaultInterval
	}

	return time.Duration(decoded.Interval)
}

func SetPeriodicInterval(duration time.Duration) error {
	var file *os.File
	var e error

	if file, e = os.Create(fileName); e != nil {
		return fmt.Errorf("error opening preferences file %w", e)
	}

	decoded := preferencesFormat{
		Interval: int64(duration),
	}

	encoder := json.NewEncoder(file)
	if e := encoder.Encode(&decoded); e != nil {
		return e
	}

	return nil
}
