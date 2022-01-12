package main

import (
	"fmt"
	"time"
)

var (
	timeNowCaller = time.Now
)

func formatA(t time.Time) string {
	return t.Format("January")
}

func formatB(t time.Time) string {
	return t.Format("2")
}

func GetTime() string {
	t := timeNowCaller()

	a := formatA(t)
	b := formatB(t)

	return fmt.Sprintf("%s %s", a, b)
}
