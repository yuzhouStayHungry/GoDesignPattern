package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	count := GetInstance()

	if count == nil {

		t.Error("A new connection object must have been made")
	}

	expectedCounter := count

	currentCount := count.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)

	}

	count2 := GetInstance()
	if count2 != expectedCounter {
		t.Error("Singleton instances must be different")
	}

	currentCount = count2.AddOne()

	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}
