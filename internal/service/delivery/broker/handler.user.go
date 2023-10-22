package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/avtara/carehub/internal/models"
	"github.com/avtara/carehub/utils"
	"github.com/hibiken/asynq"
)

func (so *aqObject) handlerProcessTaskSendEmail(ctx context.Context, task *asynq.Task) (err error) {
	var payload models.User

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("[Delivery][handlerProcessTaskSendEmail] failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	decText, err := utils.Decrypt(ctx, payload.Password, utils.GetEnv("encrypt.secret_key", "!@#SecretBgfast!@#$"))
	if err != nil {
		err = fmt.Errorf("[Delivery][handlerProcessTaskSendEmail] error decrypting your encrypted text: %s", err.Error())
		return
	}

	fmt.Printf("*******Send email to: %[1]v*******\nEmail: %[1]v\nPassword: %[2]v\n**************", payload.Email, decText)
	return
}
