package data

import (
	"fmt"
	"testing"
)

func TestPersistance(t *testing.T) {
	if e := SetPeriodicInterval(23); e != nil {
		t.Fatal(e)
	}

	value := GetPeriodicInterval()

	if value != 23 {
		t.Error("Error reading file from storage")
	}

	fmt.Printf("Reading from storage %d\n", value)
}
