package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	assert := assert.New(t)

	var test = []struct {
		input1 float32
		input2 float32
		expect float32
	}{
		{2, 4, 0.5},
		{5, 2, 2.5},
		{11, 2, 5.5},
	}
	for _, test := range test {
		resultado, _ := Division(test.input1, test.input2)
		assert.Equal(resultado, test.expect)
	}
}
