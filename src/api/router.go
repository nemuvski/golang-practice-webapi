package api

import (
	api_middleware "golang-practive-webapi/src/api/middleware"
	api_route_private "golang-practive-webapi/src/api/route/private/v1"
	api_route_public "golang-practive-webapi/src/api/route/public/v1"
	"golang-practive-webapi/src/environment"
	"golang-practive-webapi/src/kvs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Api struct {
	Router *echo.Echo
}

/*
	APIインスタンスを生成する。
*/
func New() *Api {
	return &Api{ Router: echo.New() }
}

/*
	APIのルーティング、ミドルウェアを設定する。
*/
func (a *Api) Setup(kvsClient *kvs.KeyValueStore) {
	// ログレベルを実行モードによって分ける
	if environment.IsProdMode() {
		a.Router.Logger.SetLevel(log.INFO)
	} else {
		a.Router.Logger.SetLevel(log.DEBUG)
	}

	// ミドルウェア
	a.Router.Use(middleware.Logger())
	a.Router.Use(middleware.Recover())
	a.Router.Use(middleware.Gzip())
	a.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType},
		AllowMethods: []string{echo.GET},
	}))
	a.Router.Use(api_middleware.RateLimit(kvsClient))

	// ルーティング
	v1 := a.Router.Group("/v1")
	v1.GET("/public/:id", api_route_public.V1GetIndex())
	v1.GET("/private/:id", api_route_private.V1GetIndex(), api_middleware.KeyAuth(environment.Var(environment.APP_API_VALID_TOKEN)))
}

/*
	APIサーバーを起動する。
*/
func (a *Api) Run(port string) {
	a.Router.Logger.Fatal(a.Router.Start(":" + port))
}
