package storage

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	redis.ClusterClient
}

//	NewRedis is Redis constructor
func NewRedis() *Redis {
	return &Redis{}
}

//	Connect initialize new redis cluster connection
func (r *Redis) Connect(host string) error {
	fmt.Printf("host is %s", host)
	r.ClusterClient = *redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{host}})
	status := r.ClusterClient.Ping()
	if text, err := status.Result(); text != "PONG" {
		return fmt.Errorf("cannot execute ping: %s", err)
	}
	return nil
}
