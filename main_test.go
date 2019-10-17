package main

import (
	"testing"
	"time"
)

func TestUnixTimeSample(t *testing.T) {
	cases := []struct {
		t        time.Time
		expected string
	}{
		{t: time.Unix(1419933531, 0), expected: "Odd"},
		{t: time.Unix(1419933530, 0), expected: "Even"},
	}
	mockTimeManager := &MockTimeManager{}

	for _, c := range cases {
		mockTimeManager.MockTime = &c.t
		result := UnixTimeSample(mockTimeManager)
		if result != c.expected {
			t.Errorf("expected is %s but I am %s", c.expected, result)
		}
	}
}

type MockTimeManager struct {
	MockTime *time.Time
}

func (t *MockTimeManager) Now() time.Time {
	if t.MockTime == nil {
		return time.Now()
	}
	return *t.MockTime
}
