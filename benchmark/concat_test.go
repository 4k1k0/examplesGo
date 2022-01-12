package main

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func createSource() []string {
	var source []string
	for i := 0; i < 100_000; i++ {
		var w string
		_ = faker.FakeData(&w)
		source = append(source, w)
	}
	return source
}

func TestConcatA(t *testing.T) {
	source := createSource()
	res := ConcatA(source)
	assert.NotEmpty(t, res)
}

func TestConcatB(t *testing.T) {
	source := createSource()
	res := ConcatB(source)
	assert.NotEmpty(t, res)
}

func BenchmarkConcatA(b *testing.B) {
	source := createSource()
	for i := 0; i < b.N; i++ {
		ConcatA(source)
	}
}

func BenchmarkConcatB(b *testing.B) {
	source := createSource()
	for i := 0; i < b.N; i++ {
		ConcatB(source)
	}
}
