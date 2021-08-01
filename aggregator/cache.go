package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Client  *redis.Client
	Context context.Context
}

func (c *Cache) Get(key string) (res *PredictionDTO, err error) {
	res = &PredictionDTO{}
	b, err := c.Client.Get(c.Context, key).Result()
	if err != nil {
		return nil, err
	}
	if b != ""{
		c.Client.Del(c.Context, key)
	}
	err = json.Unmarshal([]byte(b), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Cache) Set(val *PredictionDTO) error {
	key := val.TranID
	b, err := json.Marshal(*val)
	if err != nil {
		return err
	}
	return c.Client.Set(c.Context, key, string(b), 0).Err()
}
