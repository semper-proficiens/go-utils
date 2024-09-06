package distributedcache

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type Command string

const (
	CMDSet Command = "SET"
	CMDGet Command = "GET"
)

type MSGSet struct {
	Key   []byte
	Value []byte
	TTL   time.Duration
}

type Message struct {
	Cmd   Command
	Key   []byte
	Value []byte
	TTL   time.Duration
}

func parseMessage(raw []byte) (*Message, error) {
	var (
		rawStr = string(raw)
		parts  = strings.Split(rawStr, " ")
	)
	if len(parts) < 2 {
		return nil, errors.New("invalid protocol format")
	}
	msg := &Message{
		Cmd: Command(parts[0]),
		Key: []byte(parts[1]),
	}
	if msg.Cmd == CMDSet {
		if len(parts) != 4 {
			return nil, errors.New("invalid Set command")
		}
		msg.Value = []byte(parts[2])
		ttl, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, errors.New("invalid Set TTL")
		}
		msg.TTL = time.Duration(ttl) * time.Second
	}
	return msg, nil
}
