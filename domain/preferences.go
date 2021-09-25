package domain

import (
	"time"
)

type Preferences struct {
	Interval time.Duration
}

var DefaultPreferences = Preferences{
	Interval: 5 * time.Minute,
}
