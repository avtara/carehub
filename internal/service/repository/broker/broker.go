package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/avtara/carehub/internal/service"
	"github.com/hibiken/asynq"
)

type brokerRepository struct {
	conn *asynq.Client
}

func NewBrokerRepository(
	conn *asynq.Client,
) service.BrokerRepository {
	return &brokerRepository{
		conn: conn,
	}
}

func (a *brokerRepository) Publish(ctx context.Context, typename string, payload interface{}) (taskInfo *asynq.TaskInfo, err error) {
	var payloadJSON []byte
	payloadJSON, err = json.Marshal(payload)
	if err != nil {
		err = fmt.Errorf("[Repository][Publish] while marshal payload: %s", err.Error())
		return
	}

	task := asynq.NewTask(typename, payloadJSON)
	taskInfo, err = a.conn.EnqueueContext(ctx, task)
	if err != nil {
		err = fmt.Errorf("[Repository][Publish] while EnqueueContext: %s", err.Error())
		return
	}

	return
}
