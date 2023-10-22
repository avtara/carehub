package broker

import (
	"github.com/avtara/carehub/internal/models"
	"github.com/hibiken/asynq"
)

type aqObject struct {
	client      *asynq.Client
	asynqServer *asynq.Server
	asynqMux    *asynq.ServeMux
}

func NewBrokerHandler(
	client *asynq.Client,
	asynqServer *asynq.Server,
	asynqMux *asynq.ServeMux,
) {
	obj := &aqObject{
		asynqMux:    asynqMux,
		client:      client,
		asynqServer: asynqServer,
	}

	asynqMux.HandleFunc(models.TaskSendEmailNewUser, obj.handlerProcessTaskSendEmail)
	asynqServer.Start(asynqMux)
}
