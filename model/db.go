package model

import (
	"context"
	"strconv"

	"github.com/cag000/simple-message-queue/config"
	"github.com/go-redis/redis/v8"
)

type DB struct {
	KeyQueue string
	Conn *redis.Client
	Address string
}

func (d *DB) Connection(cfg *config.Config) error {
	dbNum, _ := strconv.Atoi(cfg.DbConfigClient.DbUserName)
	d.Conn = redis.NewClient(&redis.Options{
		Addr: d.Address,
		MaxRetries: 5,
		Password: "",
		DB: dbNum,
		PoolSize: 3,
		

	})
	return nil
}


func (d *DB) DeleteQueue(ctx context.Context, queue string) error {
	return d.Conn.Del(ctx, queue).Err()
}

func(d *DB) CreateQueue(ctx context.Context, queue string) error {
	return d.Conn.Set(ctx, queue, "", 0).Err()
}

func (d *DB) PushDB(ctx context.Context, queue string, msg string) error {
	return d.Conn.RPush(ctx, queue, msg).Err()
}