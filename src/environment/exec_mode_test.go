package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDevMode(t *testing.T) {
	t.Run("環境変数APP_EXEC_MODEがセットされていないため、falseを返却する", func(t *testing.T) {
		v := IsDevMode()
		assert.Equal(t, false, v)
	})

	t.Run("環境変数APP_EXEC_MODEにprodが設定されているため、falseを返却する", func(t *testing.T) {
		t.Setenv("APP_EXEC_MODE", "prod")
		v := IsDevMode()
		assert.Equal(t, false, v)
	})

	t.Run("環境変数APP_EXEC_MODEにdevが設定されているため、trueを返却する", func(t *testing.T) {
		t.Setenv("APP_EXEC_MODE", "dev")
		v := IsDevMode()
		assert.Equal(t, true, v)
	})
}

func TestIsProdMode(t *testing.T) {
	t.Run("環境変数APP_EXEC_MODEがセットされていないため、falseを返却する", func(t *testing.T) {
		v := IsProdMode()
		assert.Equal(t, false, v)
	})

	t.Run("環境変数APP_EXEC_MODEにprodが設定されているため、trueを返却する", func(t *testing.T) {
		t.Setenv("APP_EXEC_MODE", "prod")
		v := IsProdMode()
		assert.Equal(t, true, v)
	})

	t.Run("環境変数APP_EXEC_MODEにdevが設定されているため、falseを返却する", func(t *testing.T) {
		t.Setenv("APP_EXEC_MODE", "dev")
		v := IsProdMode()
		assert.Equal(t, false, v)
	})
}

