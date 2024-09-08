package conn

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/rakib-09/golang-clean-architecture/config"
	"github.com/rakib-09/golang-clean-architecture/utils/errors"
	_ "github.com/rakib-09/golang-clean-architecture/utils/errors"
	"github.com/rakib-09/golang-clean-architecture/utils/logger"
	"github.com/rakib-09/golang-clean-architecture/utils/methodutil"
	"strconv"
	"time"
)

var client *redis.Client

func ConnectRedis() {
	conf := config.Redis()

	logger.Info("connecting to redis at ", conf.Host, ":", conf.Port, "...")

	client = redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Pass,
		DB:       conf.DB,
	})

	if _, err := client.Ping().Result(); err != nil {
		logger.Error(err, "failed to connect redis", conf)
		panic(err)
	}

	logger.Info("redis connection successful...")
}

type RedisClient struct{}

func NewRedisClient() *RedisClient {
	return &RedisClient{}
}

func (rc *RedisClient) Set(key string, value interface{}, ttl int) errors.Error {
	if methodutil.IsEmpty(key) || methodutil.IsEmpty(value) {
		return errors.EmptyRedisValue()
	}

	serializedValue, err := json.Marshal(value)
	if err != nil {
		return errors.InternalServerError(err)
	}

	err = client.Set(key, serializedValue, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		return errors.InternalServerError(err)
	}
	return nil
}

func (rc *RedisClient) SetString(key string, value string, ttl time.Duration) errors.Error {
	if methodutil.IsEmpty(key) || methodutil.IsEmpty(value) {
		return errors.EmptyRedisValue()
	}

	err := client.Set(key, value, ttl*time.Second).Err()
	if err != nil {
		return errors.InternalServerError(err)
	}
	return nil
}

func (rc *RedisClient) SetStruct(key string, value interface{}, ttl time.Duration) error {
	serializedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return client.Set(key, string(serializedValue), ttl*time.Second).Err()
}

func (rc *RedisClient) Get(key string) (string, error) {
	return client.Get(key).Result()
}

func (rc *RedisClient) GetInt(key string) (int, error) {
	str, err := client.Get(key).Result()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(str)
}

func (rc *RedisClient) GetStruct(key string, outputStruct interface{}) error {
	serializedValue, err := client.Get(key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(serializedValue), &outputStruct); err != nil {
		return err
	}

	return nil
}

func (rc *RedisClient) Del(keys ...string) error {
	return client.Del(keys...).Err()
}

func (rc *RedisClient) DelPattern(pattern string) error {
	iter := client.Scan(0, pattern, 0).Iterator()

	for iter.Next() {
		err := client.Del(iter.Val()).Err()
		if err != nil {
			return err
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}

func (rc *RedisClient) Exists(key string) bool {
	exists, err := client.Exists(key).Result()
	if err != nil {
		return false
	}

	return exists == 1
}
