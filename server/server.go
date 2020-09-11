package server

import (
	"context"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

type Queue struct {
	Address string
	Listener net.Listener
}

func (q *Queue) StartServer() error  {
	ln, err := net.Listen("tcp", q.Address)
	if err != nil {
		return err
	}
	q.Listener = ln
	return nil
}

func Over(ctx context.Context) {
    select {
    case <-time.After(1 * time.Second):
		logrus.Info("Oversleep")
	case <-ctx.Done():
    }

}