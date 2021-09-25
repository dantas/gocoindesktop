package data

import (
	"testing"
	"time"

	"github.com/dantas/gocoindesktop/domain"
)

func TestPersistance(t *testing.T) {
	preference := domain.Preferences{
		Interval: 3 * time.Hour,
	}

	if e := SavePreferences(preference); e != nil {
		t.Fatal(e)
	}

	if loadedPreference, e := LoadPreferences(); e != nil {
		t.Error("Error reading file from storage")
	} else if preference != loadedPreference {
		t.Error("Loaded preference is different from what is expected")
	}
}
