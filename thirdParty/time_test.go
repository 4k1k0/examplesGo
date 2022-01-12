package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func resetMocks(t *testing.T) {
	t.Helper()

	timeNowCaller = time.Now
}

func TestGetTime(t *testing.T) {
	defer resetMocks(t)
	timeNowCaller = func() time.Time {
		return time.Date(2009, 11, 10, 12, 30, 0, 651387237, time.UTC) // Golang release date.
	}

	res := GetTime()
	assert.Equal(t, "November 10", res)
}
