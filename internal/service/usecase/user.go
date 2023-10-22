package usecase

import (
	"context"
	"fmt"
	"github.com/avtara/carehub/internal/models"
	"github.com/avtara/carehub/internal/service"
	"github.com/avtara/carehub/utils"
	"github.com/go-playground/validator"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type userUseCase struct {
	userRepository  service.UserRepository
	asyncRepository service.BrokerRepository
	db              *sqlx.DB
}

func NewUserUseCase(userRepository service.UserRepository, asyncRepository service.BrokerRepository, db *sqlx.DB) service.UserUseCase {
	return &userUseCase{
		userRepository:  userRepository,
		asyncRepository: asyncRepository,
		db:              db,
	}
}

func (u userUseCase) UpdateProfile(ctx context.Context, args models.UpdateProfileUserParams, userID int64) (err error) {
	var (
		user models.User
		tx   *sqlx.Tx
	)
	tx, err = u.db.Beginx()

	defer func() {
		if tx == nil {
			return
		}

		if err != nil {
			serr := tx.Rollback()
			if serr != nil {
				log.Errorf("[Usecase][UpdateProfile] failed to rollback: %s", err.Error())
			}
		}
	}()

	user, err = u.userRepository.GetByID(ctx, userID)
	if err != nil && err == models.ErrorUserNotFound {
		return
	}

	if err != nil {
		err = fmt.Errorf("[Usecase][UpdateProfile] failed while get user by id: %s", err.Error())
		return
	}

	if args.Name != "" {
		user.Name = utils.CapitalizeString(args.Name)
	}

	if args.Photo != "" {
		if err = validator.New().VarCtx(ctx, args.Photo, "url"); err == nil {
			user.Photo = args.Photo
		}

		if err = validator.New().VarCtx(ctx, args.Photo, "base64"); err == nil {
			var url string
			url, err = utils.SaveImageToLocalDir(args.Photo, models.GroupFilePhotoUser)
			if err != nil {
				err = fmt.Errorf("[Usecase][UpdateProfile] failed while save file to local dir: %s", err.Error())
				return
			}
			user.Photo = fmt.Sprintf("/%s", url)
		}
	}

	err = u.userRepository.UpdateProfileByID(ctx, tx, user)
	if err != nil && err == models.ErrorUserNotFound {
		err = fmt.Errorf("[Usecase][UpdateProfile] failed while update profile by id: %s", err.Error())
		return
	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("[Usecase][UpdateProfile] failed to commit: %s", err.Error())
		return
	}
	return
}
