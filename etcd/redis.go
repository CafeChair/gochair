package etcd
import (
	"github.com/hoisie/redis"
)

func LogTo(key ,value string) bool {
	redishost := Config().Redis.Addr
	client := redis.Client{Addr:redishost}
	err := client.Set(key, []byte(value))
	if err != nil {
		return false
	}
	return true
}