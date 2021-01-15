package gredis

import (
	"coin_price_service/pkg/setting"
	"gopkg.in/redis.v4"
	"sync"
	"time"
)

const (
	TronBlockTime = "tron_block_time"
	EtherBlock    = "ether_block"
)

var RedisSnapshot = make(chan bool)

type redisSnapshotStore struct {
	redisClient *redis.Client
}

var store *redisSnapshotStore
var onceStore sync.Once

// Setup Initialize the Redis instance
func SharedSnapshotStore() *redisSnapshotStore {
	onceStore.Do(func() {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     setting.RedisSetting.Host,
			Password: setting.RedisSetting.Password,
			DB:       0,
		})

		store = &redisSnapshotStore{redisClient: redisClient}
	})
	return store
}

func (s *redisSnapshotStore) Set(key string, data int64, duration time.Duration) error {
	return s.redisClient.Set(key, data, duration).Err()
}

// Get get a key
func (s *redisSnapshotStore) Get(key string) (int64, error) {
	return s.redisClient.Get(key).Int64()
}
