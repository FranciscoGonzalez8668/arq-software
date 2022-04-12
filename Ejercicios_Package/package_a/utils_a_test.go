package package_a

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Hello(), "Hello A")
}
