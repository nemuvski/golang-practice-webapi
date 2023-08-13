package main

import (
	"golang-practive-webapi/src/api"
	"golang-practive-webapi/src/environment"
	"golang-practive-webapi/src/kvs"
	"log"
)

const TARGET_ENV_FILE = ".env.local"

func main() {
	initEnv()

	// Redisサーバー接続
	kvsClient, err := kvs.New(environment.Var(environment.APP_REDIS_URL))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = kvsClient.Flush()
	if err != nil {
		log.Fatalln(err)
	}
	defer kvsClient.Close()

	// APIサーバー起動
	srv := api.New()
	srv.Setup(kvsClient)
	srv.Run(environment.Var(environment.APP_API_PORT))
}

func initEnv() {
	environment.LoadEnvFile(TARGET_ENV_FILE)
	
	if err := environment.ValidateVars(); err != nil {
		log.Fatalln(err)
	}
}
