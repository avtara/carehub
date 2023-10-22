package postgres

import (
	"context"
	"fmt"
	"github.com/avtara/carehub/internal/models"
	"github.com/avtara/carehub/internal/service"
	"github.com/avtara/carehub/internal/service/repository/postgres/queries"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"
)

type userRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(
	conn *sqlx.DB,
) service.UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (result models.User, err error) {
	err = u.conn.QueryRowContext(
		ctx,
		queries.GetUserByEmail,
		email,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Password,
		&result.Email,
		&result.Photo,
		&result.Role,
	)

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			err = models.ErrorUserNotFound
			return
		}

		err = fmt.Errorf("[Repository][GetByEmail] failed getting user by email: %s", err.Error())
		return
	}

	return
}

func (u *userRepository) Insert(ctx context.Context, tx *sqlx.Tx, args models.User) (id int64, err error) {

	if tx == nil {
		err = u.conn.QueryRowContext(
			ctx,
			queries.InsertUser,
			args.Name,
			args.Password,
			args.Email,
			args.Photo,
			args.Role).Scan(&id)
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				err = models.ErrorUserDuplicate
				return
			}
		}

		if err != nil {
			err = fmt.Errorf("[Repository][Insert] failed insert user: %s", err.Error())
			return
		}

		return
	}

	err = tx.QueryRowContext(
		ctx,
		queries.InsertUser,
		args.Name,
		args.Password,
		args.Email,
		args.Photo,
		args.Role).Scan(&id)
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "foreign_key_violation", "unique_violation":
			err = models.ErrorUserDuplicate
			return
		}
	}

	if err != nil {
		err = fmt.Errorf("[Repository][Insert] failed insert user: %s", err.Error())
		return
	}

	return
}

func (u *userRepository) GetByID(ctx context.Context, userID int64) (result models.User, err error) {
	err = u.conn.QueryRowContext(
		ctx,
		queries.GetUserByID,
		userID,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Password,
		&result.Email,
		&result.Photo,
		&result.Role,
	)

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			err = models.ErrorUserNotFound
			return
		}

		err = fmt.Errorf("[Repository][GetByEmail] failed getting user by id: %s", err.Error())
		return
	}

	return
}

func (u *userRepository) UpdateProfileByID(ctx context.Context, tx *sqlx.Tx, user models.User) (err error) {
	if tx == nil {
		err = u.conn.QueryRowContext(
			ctx,
			queries.UpdateProfileByID,
			user.Name,
			user.Photo,
			user.ID,
		).Err()

		if err != nil {
			err = fmt.Errorf("[Repository][UpdateProfileByID] failed update profile: %s", err.Error())
			return
		}

		return
	}

	err = tx.QueryRowContext(
		ctx,
		queries.UpdateProfileByID,
		user.Name,
		user.Photo,
		user.ID,
	).Err()

	if err != nil {
		err = fmt.Errorf("[Repository][UpdateProfileByID] failed update profile user: %s", err.Error())
		return
	}

	return
}
