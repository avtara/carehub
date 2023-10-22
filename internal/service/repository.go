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
	GetAllCategories(ctx context.Context, limit int) (result []models.Category, err error)
	GetCategoryByID(ctx context.Context, ID int64) (result models.Category, err error)
	GetExtraFieldByCategoryID(ctx context.Context, categoryID int64) (result []models.ExtraFieldCategory, err error)
}

type ComplainRepository interface {
	GetAllComplain(ctx context.Context, limit int) (result []models.Complain, err error)
	GetComplainByID(ctx context.Context, ID int64) (result models.Complain, err error)
	GetResolutionByComplainID(ctx context.Context, complainID int64) (result []models.Resolution, err error)
	InsertComplain(ctx context.Context, args models.InsertComplainParams, userID int64) (err error)
	InsertResolution(ctx context.Context, args models.InsertResolutionParams, complainID, adminID int64) (err error)
	UpdateStatus(ctx context.Context, status string, complainID int64) (err error)
}
