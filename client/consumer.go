package client

import (
	// "bufio"
	// "fmt"
	
	"context"
	
	
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type MyConsumer struct {
	Queue   string
}

func (m *MyConsumer) ConsumeMessage(ctx context.Context, r *redis.Client) error {
	go func()  {
		for {
			res, err := r.BLPop(ctx, 0*time.Second, m.Queue).Result()
			if err != nil {
				return
			}
			logrus.Println(res[1])
		}
	}()
	select {}
	return nil
}

