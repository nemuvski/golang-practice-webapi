package environment

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	テストに用いる .env ファイルのパス
*/
const TESTING_ENV_FILE = "../../__fixtures__/testing__dotenv.env"

func TestLoadEnvFile(t *testing.T) {
	t.Run("ファイルが存在しないため、読み取りはスキップする", func(t *testing.T) {
		LoadEnvFile("NOT_EXIST_FILE")
	})

	t.Run("ファイルが存在するため、読み取りを行う", func(t *testing.T) {
		LoadEnvFile(TESTING_ENV_FILE)
		assert.Equal(t, "10", os.Getenv("TESTING_DOTENV_ENV1"))
		assert.Equal(t, "Dummy data", os.Getenv("TESTING_DOTENV_ENV2"))

		// テストで追加された環境変数をクリアする
		os.Unsetenv("TESTING_DOTENV_ENV1")
		os.Unsetenv("TESTING_DOTENV_ENV2")
	})
}
