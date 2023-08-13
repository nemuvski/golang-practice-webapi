package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateVars(t *testing.T) {
	t.Run("必要な環境変数がセットされているため、エラーを返却しない", func(t *testing.T) {
		t.Setenv("APP_EXEC_MODE", "prod")
		t.Setenv("APP_API_PORT", "8080")
		t.Setenv("APP_API_VALID_TOKEN", "DUMMY")
		t.Setenv("APP_REDIS_URL", "redis://localhost:6379")
		err := ValidateVars()
		assert.Nil(t, err)
	})

	t.Run("必要な環境変数がセットされていないため、エラーを返却する", func(t *testing.T) {
		err := ValidateVars()
		assert.ErrorContains(t, err, "APP_EXEC_MODE")
		assert.ErrorContains(t, err, "APP_API_PORT")
		assert.ErrorContains(t, err, "APP_API_VALID_TOKEN")
		assert.ErrorContains(t, err, "APP_REDIS_URL")
	})

	t.Run("環境変数APP_EXEC_MODEが無効な値であるため、エラーを返却する", func(t *testing.T) {
		t.Setenv("APP_EXEC_MODE", "DUMMY")
		t.Setenv("APP_API_PORT", "8080")
		t.Setenv("APP_API_VALID_TOKEN", "DUMMY")
		t.Setenv("APP_REDIS_URL", "redis://localhost:6379")
		err := ValidateVars()
		assert.ErrorContains(t, err, "APP_EXEC_MODE")
	})

	t.Run("環境変数APP_API_PORTが無効な値であるため、エラーを返却する", func(t *testing.T) {
		t.Setenv("APP_EXEC_MODE", "prod")
		t.Setenv("APP_API_PORT", "-80")
		t.Setenv("APP_API_VALID_TOKEN", "DUMMY")
		t.Setenv("APP_REDIS_URL", "redis://localhost:6379")
		err := ValidateVars()
		assert.ErrorContains(t, err, "APP_API_PORT")
	})

	t.Run("環境変数APP_REDIS_URLが無効な値であるため、エラーを返却する", func(t *testing.T) {
		t.Setenv("APP_EXEC_MODE", "dev")
		t.Setenv("APP_API_PORT", "8080")
		t.Setenv("APP_API_VALID_TOKEN", "DUMMY")
		t.Setenv("APP_REDIS_URL", "DUMMY://localhost:6379")
		err := ValidateVars()
		assert.ErrorContains(t, err, "APP_REDIS_URL")
	})
}

func TestVar(t *testing.T) {
	t.Run("環境変数APP_REDIS_URLがセットされていないため、空文字を返却する", func(t *testing.T) {
		v := Var("APP_REDIS_URL")
		assert.Equal(t, "", v)
	})

	t.Run("環境変数APP_REDIS_URLがセットされているため、その内容を返却する", func(t *testing.T) {
		t.Setenv("APP_REDIS_URL", "redis://localhost:6379")
		v := Var("APP_REDIS_URL")
		assert.Equal(t, "redis://localhost:6379", v)
	})
}
