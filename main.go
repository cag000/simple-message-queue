package main

/*
	Author: Syahrul Al-Rasyid
*/

import (
	"bufio"
	"context"
	"flag"
	"fmt"

	"strings"

	"github.com/cag000/simple-message-queue/client"
	"github.com/cag000/simple-message-queue/config"
	"github.com/cag000/simple-message-queue/model"
	"github.com/joho/godotenv"

	"github.com/cag000/simple-message-queue/server"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	customFormatterG *logrus.TextFormatter
	connRedis        *redis.Client
	cfg              *config.Config
	ctx              context.Context
	queue            string
	auth             bool
	username         string
	password         string
	err              error
	serverLive       bool
	consumerGroup    bool
	producerGroup    bool
	queueCreate      bool
	queueDelete      bool
	configMe         string
)

func init() {
	// context
	ctx = context.Background()

	// logger
	customFormatter := &logrus.TextFormatter{}
	customFormatter.TimestampFormat = "06/01/02 15:04:05"
	customFormatter.FullTimestamp = true
	customFormatterG = customFormatter

	// args
	flag.StringVar(&configMe, "config", "", "config file .env format")
	flag.BoolVar(&serverLive, "server_queue", false, "Make live server [./app -server_queue=true]")
	flag.StringVar(&queue, "queue", "", "name queue [./app -queue=$name]")
	flag.BoolVar(&queueCreate, "create_queue", false, "create queue [./app -create_queue=true -queue=$name]")
	flag.BoolVar(&queueDelete, "delete_queue", false, "delete queue [./app -delete_queue=true -queue=$name]")
	flag.BoolVar(&consumerGroup, "consumer", false, "consumer [./app -consumer=true -queue=$name]")
	flag.BoolVar(&producerGroup, "producer", false, "producer [./app -producer=true -queue=$name]")
	flag.StringVar(&username, "user", "", "username [./app -user=$user]")
	flag.StringVar(&password, "pwd", "", "password [./app -pwd=$password]")
	flag.BoolVar(&auth, "auth", false, "auth [./app -auth=true -user=$username -pwd=$password]")
	flag.Parse()

	// env load
	err := godotenv.Load(configMe)
	if err != nil {
		logrus.Error(err)
	}

	// set config
	cfg = config.New()
}

func main() {
	logrus.SetFormatter(customFormatterG)

	data := new(client.Dummy)

	redisDb := new(model.DB)
	
	if auth {
		if username == cfg.UserApp.Username && password == cfg.UserApp.Password {
			logrus.Infof("Success Login User [%s]", username)
			if serverLive {
				logrus.Infof("Server start at %s", cfg.Mqueue.Address)
				serv := &server.Queue{
					Address: fmt.Sprintf("%s", cfg.Mqueue.Address),
				}
				err = serv.StartServer()
				if err != nil {
					logrus.Error(err)
				}

				// exit := make(chan bool)
				for {
					if conn, err := serv.Listener.Accept(); err == nil {
						redisDb.Connection(cfg)
						go func() {
							defer redisDb.Conn.Close()
							buf := bufio.NewReader(conn)
							for {
								msg, err := buf.ReadString('\n')
								if err == nil {
									if !strings.Contains(msg, "Delete=") {
										q := msg[:strings.IndexByte(msg, '=')]
										dMsg := strings.Replace(msg, fmt.Sprintf("%s=", q), "", 1)
										redisDb.PushDB(ctx, q, dMsg)
										logrus.Infof("Push to queue [%s] Address [%v]", q, cfg.Mqueue.Address)
									} else {
										qDelte := strings.Replace(msg, "Delete=", "", 1)
										err = redisDb.DeleteQueue(ctx, strings.TrimSuffix(qDelte, "\n"))
										if err != nil {
											logrus.Error(err)
										}
										logrus.Infof("Success Delete Queue %v", qDelte)
									}
								}
							}
						}()
					}
				}
			}

			// use json for sending dummy message
			if producerGroup {
				// producer
				tmpName := queue
				logrus.Infof("Producer start at %s", cfg.Mqueue.Address)
				producer := &client.Producer{
					QueueName: tmpName,
				}
				err = producer.Connect(fmt.Sprintf("%s", cfg.Mqueue.Address))
				if err != nil {
					logrus.Error(err)
				}
				// test data send
				err = data.ReadFile("./client/dummy.json")
				if err != nil {
					logrus.Error(err)
				}
				count := 0
				for _, v := range data.Datas {
					err = producer.SendMessage(v)
					if err != nil {
						logrus.Error(err)
					}
					count++
					logrus.Infof("send message %d", count)
				}
			}

			if consumerGroup {
				// consumer
				err = redisDb.Connection(cfg)
				if err != nil {
					logrus.Error(err)
				}
				logrus.Infof("Consumer start at %s", cfg.Mqueue.Address)
				consumer := &client.MyConsumer{
					Queue: queue,
				}

				err = consumer.ConsumeMessage(ctx, redisDb.Conn)
				if err != nil {
					logrus.Error(err)
				}
			}

			if queueCreate {
				err = redisDb.Connection(cfg)
				if err != nil {
					logrus.Info(err)
				}
				logrus.Infof("Success Delete [%s] Queue", queue)
				err = redisDb.CreateQueue(ctx, queue)
				if err != nil {
					logrus.Error(err)
				}

			}

			if queueDelete {
				err = redisDb.Connection(cfg)
				if err != nil {
					logrus.Info(err)
				}
				err = redisDb.DeleteQueue(ctx, queue)
				if err != nil {
					logrus.Error(err)
				}
				logrus.Infof("Success Delete [%s] Queue", queue)
			}
		} else {
			logrus.Error("Incorrect User authentication")
		}
	} else {
		logrus.Error("Need username & password to continue process !!!!")
	}

}
