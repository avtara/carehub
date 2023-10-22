package service

import (
	"context"
	"github.com/avtara/carehub/internal/models"
	"github.com/hibiken/asynq"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (result models.User, err error)
	GetByID(ctx context.Context, userID int64) (result models.User, err error)
	Insert(ctx context.Context, tx *sqlx.Tx, args models.User) (id int64, err error)
	UpdateProfileByID(ctx context.Context, tx *sqlx.Tx, user models.User) (err error)
}

type BrokerRepository interface {
	Publish(ctx context.Context, typename string, payload interface{}) (taskInfo *asynq.TaskInfo, err error)
}

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) (response []models.Category, err error)
	GetCategoryByID(ctx context.Context, ID int64) (response []models.Category, err error)
}
