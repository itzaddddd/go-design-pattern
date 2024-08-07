package creational

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	expectedCounter := counter1
	counter1.AddOne()
	if counter1.GetCount() != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", counter1.GetCount())
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but it got a different instance")
	}
}
