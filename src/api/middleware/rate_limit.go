package api_middleware

import (
	"golang-practive-webapi/src/kvs"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

type rateLimitStore struct {
	kvsClient *kvs.KeyValueStore
}

const (
	// キーをストアしておく期限
	expiry = time.Minute
	// 期限内で許可するリクエスト数
	max = 10
)

/*
	IPアドレスごとに、1分間に10回までのリクエストを許可する。
*/
func (s *rateLimitStore) Allow(identifier string) (bool, error) {
	count, err := s.kvsClient.Client.Get(kvs.Ctx, identifier).Result()
	// まだキーが存在しない場合は、キーを登録してから true を返す
	if err == redis.Nil {
		s.kvsClient.Client.Set(kvs.Ctx, identifier, 1, expiry)
		return true, nil
	}
	// それ以外のエラーはエラーとして返す
	if err != nil {
		return false, err
	}

	countInt, err := strconv.Atoi(count)
	// 登録値のパース失敗時はエラーとして返す
	if err != nil {
		return false, err
	}
	if int(countInt) < max {
		s.kvsClient.Client.Incr(kvs.Ctx, identifier)
		return true, nil
	}

	return false, nil
}

func RateLimit(kvsClient *kvs.KeyValueStore) echo.MiddlewareFunc {
	store := &rateLimitStore{kvsClient: kvsClient}

	return middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,

		Store: store,

		IdentifierExtractor: func(c echo.Context) (string, error) {
			ip := c.RealIP()
			return ip, nil
		},

		ErrorHandler: func(c echo.Context, err error) error {
			return echo.ErrForbidden
		},

		DenyHandler: func(c echo.Context, identifier string, err error) error {
			return echo.ErrTooManyRequests
		},
	})
}
