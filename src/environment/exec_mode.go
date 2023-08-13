package environment

const (
	EXEC_MODE_DEV	= "dev"
	EXEC_MODE_PROD	= "prod"
)

func IsDevMode() bool {
	return Var(APP_EXEC_MODE) == EXEC_MODE_DEV
}

func IsProdMode() bool {
	return Var(APP_EXEC_MODE) == EXEC_MODE_PROD
}
