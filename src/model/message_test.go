package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMessageID(t *testing.T) {
	t.Run("IDが1の場合、エラーを返却しない", func(t *testing.T) {
		ok, err := ParseMessageID("1")
		assert.Equal(t, uint8(1), ok)
		assert.Nil(t, err)
	})

	t.Run("IDが256の場合、エラーを返却する", func(t *testing.T) {
		_, err := ParseMessageID("256")
		assert.NotNil(t, err)
	})
}
