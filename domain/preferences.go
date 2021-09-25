package domain

import (
	"errors"
	"time"
)

type Preferences struct {
	Interval time.Duration
}

var DefaultPreferences = Preferences{
	Interval: 5 * time.Minute,
}

var ErrLoadingPreferences = errors.New("preference file not available")

var ErrSavingPreferences = errors.New("error saving preference")
