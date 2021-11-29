package main

import (
	"math/rand"
	"testing"
	"testing/quick"
	"time"
)

func getConfig(t *testing.T) quick.Config {
	t.Helper()

	src := rand.NewSource(time.Now().Unix())

	return quick.Config{
		MaxCount: 500,
		Rand:     rand.New(src),
	}
}

func TestMask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res, err := Mask("1234567890123456")

		if err != nil {
			t.Log("err shuold be nil")
			t.Fail()
		}

		if len(res) != 16 {
			t.Log("len should be 16")
			t.Fail()
		}

		if res != "1234XXXXXXXX3456" {
			t.Log("wrong output")
			t.Fail()
		}

		if !validMask.MatchString(res) {
			t.Log("res does not pass the regular expression")
			t.Fail()
		}
	})

	t.Run("random cards", func(t *testing.T) {
		f := func(card string) bool {
			res, err := Mask(card)

			if err != nil {
				return len(res) == 0
			}

			expectedLen := len(res) == 16
			passRegex := validMask.MatchString(res)

			return expectedLen && passRegex
		}
		cfg := getConfig(t)

		if err := quick.Check(f, &cfg); err != nil {
			t.Error(err)
		}
	})
}
