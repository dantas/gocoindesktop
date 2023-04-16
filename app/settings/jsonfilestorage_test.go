package settings

import (
	"os"
	"path"
	"testing"
	"time"
)

func TestFileStorageIsCorrectlySavingAndLoading(t *testing.T) {
	location := path.Join(os.TempDir(), "preferences.json")

	if e := os.Remove(location); e != nil && !os.IsNotExist(e) {
		t.Errorf("Error removing file from temporary storage, test setup error: %v", e)
	}

	storage := NewJsonFileStorage(location)

	preference := Settings{
		Interval:         3 * time.Hour,
		ShowWindowOnOpen: true,
	}

	if e := storage.Save(preference); e != nil {
		t.Fatal(e)
	}

	if loadedPreference, e := storage.Load(); e != nil {
		t.Error("Error reading file from storage")
	} else if preference != loadedPreference {
		t.Error("Loaded preference is different from what is expected")
	}
}