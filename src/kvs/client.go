package kvs

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/redis/go-redis/v9"
)

type KeyValueStore struct {
	Client *redis.Client
}

var Ctx = context.TODO()

/*
	Redis クライアントを生成する。
*/
func New(dsn string) (*KeyValueStore, error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse redis dsn")
	}
	client := redis.NewClient(opt)

	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, errors.Wrap(err, "failed to ping redis")
	}

	return &KeyValueStore{ Client: client }, nil
}

func (kvs *KeyValueStore) Flush() (string, error) {
	return kvs.Client.FlushDB(Ctx).Result()
}

func (kvs *KeyValueStore) Close() error {
	return kvs.Client.Close()
}
