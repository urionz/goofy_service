package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type IConnection interface {
	Get(key string) string
	Set(key string, value interface{}, expiration time.Duration) error
	SetEX(key string, value interface{}, expiration time.Duration) error
	SAdd(key string, members ...interface{}) error
}

type Factory interface {
	Connection(name ...string) (*Connection, error)
}

type Connection struct {
	client *redis.Client
	name   string
}

var _ IConnection = (*Connection)(nil)

func NewConnection(client *redis.Client) *Connection {
	return &Connection{
		client: client,
	}
}

func (conn *Connection) SetName(name string) *Connection {
	conn.name = name
	return conn
}

func (conn *Connection) GetName() string {
	return conn.name
}

func (conn *Connection) Get(key string) string {
	return conn.client.Get(context.Background(), key).Val()
}

func (conn *Connection) Set(key string, value interface{}, expiration time.Duration) error {
	return conn.client.Set(context.Background(), key, value, expiration).Err()
}

func (conn *Connection) SetEX(key string, value interface{}, expiration time.Duration) error {
	return conn.client.SetEX(context.Background(), key, value, expiration).Err()
}

func (conn *Connection) SAdd(key string, members ...interface{}) error {
	return conn.client.SAdd(context.Background(), key, members...).Err()
}
