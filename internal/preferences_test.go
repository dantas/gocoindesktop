package coindesktop

import (
	"fmt"
	"testing"
)

func TestPersistance(t *testing.T) {
	if e := SetPeriodicInterval(666); e != nil {
		t.Fatal(e)
	}

	value := GetPeriodicInterval()

	if value != 666 {
		t.Error("Error reading file from storage")
	}

	fmt.Printf("Reading from storage %d\n", value)
}
