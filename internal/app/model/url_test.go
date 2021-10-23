package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	str, err := GenerateNewUrl()
	assert.NoError(t, err)
	assert.Equal(t, 10, len(str))
}
