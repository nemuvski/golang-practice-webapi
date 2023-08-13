package kvs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("Redisクライアントに失敗する", func(t *testing.T) {
		client, err := New("THIS_IS_DUMMY_DSN")
		assert.Nil(t, client)
		assert.NotNil(t, err)
	})
}
