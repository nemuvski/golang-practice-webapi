package environment

import (
	"os"
	"regexp"

	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

const (
	APP_EXEC_MODE = "APP_EXEC_MODE"
	APP_API_PORT  = "APP_API_PORT"
	APP_API_VALID_TOKEN = "APP_API_VALID_TOKEN"
	APP_REDIS_URL = "APP_REDIS_URL"
)

/*
	環境変数のバリデーションを行う。
*/
func ValidateVars() error {
	err := validation.Errors{
		APP_EXEC_MODE: validation.Validate(
			os.Getenv(APP_EXEC_MODE),
			validation.Required.Error("APP_EXEC_MODEは必須です"),
			validation.In(EXEC_MODE_DEV, EXEC_MODE_PROD).Error("APP_EXEC_MODEはdevかprodのいずれかである必要があります"),
		),

		APP_API_PORT: validation.Validate(
			os.Getenv(APP_API_PORT),
			validation.Required.Error("APP_API_PORTは必須です"),
			is.Port.Error("APP_API_PORTはポート番号である必要があります"),
		),

		APP_API_VALID_TOKEN: validation.Validate(
			os.Getenv(APP_API_VALID_TOKEN),
			validation.Required.Error("APP_API_VALID_TOKENは必須です"),
		),

		APP_REDIS_URL: validation.Validate(
			os.Getenv(APP_REDIS_URL),
			validation.Required.Error("APP_REDIS_URLは必須です"),
			validation.Match(regexp.MustCompile("^redis://.*$")).Error("APP_REDIS_URLはredis://から始まる必要があります"),
		),
	}.Filter()
	return err
}

func Var(key string) string {
	return os.Getenv(key)
}
