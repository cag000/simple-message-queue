package client

import (
	"net"
)

type Producer struct {
	Con       net.Conn
	QueueName string
}

func (p *Producer) Connect(addresss string) error {
	conn, err := net.Dial("tcp", addresss)
	if err != nil {
		return err
	}
	p.Con = conn
	return nil
}

func (p *Producer) SendMessage(message string) error {
	_, err := p.Con.Write([]byte(p.QueueName + "=" + message))
	if err != nil {
		return err
	}
	return nil
}

