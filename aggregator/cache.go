package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Client  *redis.Client
	Context context.Context
}

func (c *Cache) Get(key string) (res *PredictionDTO, err error) {
	res = &PredictionDTO{}
	b, err := c.Client.Get(c.Context, key).Result()
	fmt.Println(b)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(b), res)
	if err != nil {
		return nil, err
	}
	fmt.Println(res);
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
