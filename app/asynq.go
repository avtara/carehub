package app

import (
	"fmt"
	"github.com/avtara/carehub/utils"
	"github.com/hibiken/asynq"
)

func (cfg *App) InitAsynq() (err error) {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%s", utils.GetEnv("redis.host", "127.0.0.1"), utils.GetEnv("redis.port", "6379")),
	})

	server := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: fmt.Sprintf("%s:%s", utils.GetEnv("redis.host", "127.0.0.1"), utils.GetEnv("redis.port", "6379")),
		},
		asynq.Config{
			Queues: map[string]int{
				"critical": 10,
				"default":  5,
			},
		},
	)

	mux := asynq.NewServeMux()

	cfg.Asynq = client
	cfg.AsynqServer = server
	cfg.AsynqMux = mux
	return
}
