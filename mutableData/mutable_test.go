package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoo(t *testing.T) {
	t.Run("foo should be foo", func(t *testing.T) {
		res := Foo()
		assert.Equal(t, "foo", res)
	})
}

func TestBar(t *testing.T) {
	t.Run("should get a true value", func(t *testing.T) {
		t.Run("if response is foo", func(t *testing.T) {
			res := Bar("foo")
			assert.True(t, res)
		})

		t.Run("if response is bar", func(t *testing.T) {
			response = "bar"
			res := Bar("bar")
			assert.True(t, res)
		})
	})

	t.Run("should get a false value", func(t *testing.T) {
		res := Bar("fake")
		assert.False(t, res)
	})
}
