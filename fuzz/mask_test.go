package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res, err := Mask("1234567890123456")

		if err != nil {
			t.Log("err should be nil")
			t.FailNow()
		}

		if len(res) != 16 {
			t.Log("len should be 16")
			t.Fail()
		}

		if res != "1234XXXXXXXX3456" {
			t.Log("wrong output")
			t.Fail()
		}

	})
}

func FuzzMask(f *testing.F) {
	cases := []string{"1234567890123456", "", "hello world", "🙅"}
	for _, c := range cases {
		f.Add(c)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		res, err := Mask(orig)

		if err != nil {
			assert.Empty(t, res)
			return
		}

		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		passRegex := validMask.MatchString(res)
		assert.True(t, passRegex, "should pass the regex")
	})
}
