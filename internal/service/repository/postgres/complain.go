package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/avtara/carehub/internal/models"
	"github.com/avtara/carehub/internal/service"
	"github.com/avtara/carehub/internal/service/repository/postgres/queries"
	"github.com/jmoiron/sqlx"
)

type complainRepository struct {
	conn *sqlx.DB
}

func NewComplainRepository(
	conn *sqlx.DB,
) service.ComplainRepository {
	return &complainRepository{
		conn: conn,
	}
}

func (c *complainRepository) GetAllComplain(ctx context.Context, limit int) (result []models.Complain, err error) {
	rows, err := c.conn.QueryContext(ctx, queries.GetAllComplain, limit)
	if err != nil {
		if err != sql.ErrNoRows {
			return
		}
	}
	defer rows.Close()

	for rows.Next() {
		var tmp models.Complain
		err = rows.Scan(
			&tmp.ID,
			&tmp.Description,
			&tmp.ComplaintDetail,
			&tmp.Status,
			&tmp.User.Name,
			&tmp.User.Email,
			&tmp.User.Photo,
			&tmp.Category,
		)
		if err != nil {
			err = fmt.Errorf("[Repository][GetAllComplain] failed scan GetAllComplain: %s", err.Error())
			return
		}
		result = append(result, tmp)
	}

	return
}

func (c *complainRepository) GetComplainByID(ctx context.Context, ID int64) (result models.Complain, err error) {
	rows, err := c.conn.QueryContext(ctx, queries.GetComplainByID, ID)
	if err != nil {
		err = fmt.Errorf("[Repository][GetComplainByID] failed getting complain: %s", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tmp models.Complain
		err = rows.Scan(
			&tmp.ID,
			&tmp.Description,
			&tmp.ComplaintDetail,
			&tmp.Status,
			&tmp.User.Name,
			&tmp.User.Email,
			&tmp.User.Photo,
			&tmp.Category,
		)
		if err != nil {
			err = fmt.Errorf("[Repository][GetCategoryByID] failed scan Complain: %s", err.Error())
			return
		}
		result = tmp
	}

	if result.ID == 0 {
		err = models.ErrorCategoryNotFound
		return
	}

	return
}

func (c *complainRepository) GetResolutionByComplainID(ctx context.Context, complainID int64) (result []models.Resolution, err error) {
	rows, err := c.conn.QueryContext(ctx, queries.GetResolutionByID, complainID)
	if err != nil {
		if err != sql.ErrNoRows {
			err = fmt.Errorf("[Repository][GetResolutionByComplainID] failed getting resolution complain: %s", err.Error())
			return
		}
	}
	defer rows.Close()

	for rows.Next() {
		var tmp models.Resolution
		err = rows.Scan(&tmp.ID, &tmp.Remark, &tmp.ResolvedBy.Name, &tmp.ResolvedBy.Email, &tmp.ResolvedBy.Photo)
		if err != nil {
			err = fmt.Errorf("[Repository][GetExtraFieldByCategoryID] failed scan Resolution: %s", err.Error())
			return
		}
		result = append(result, tmp)
	}

	return
}

func (c *complainRepository) InsertComplain(ctx context.Context, args models.InsertComplainParams, userID int64) (err error) {
	err = c.conn.QueryRowContext(
		ctx,
		queries.InsertComplain,
		args.CategoryID,
		userID,
		args.Description,
		args.ComplaintDetail,
	).Err()

	if err != nil {
		err = fmt.Errorf("[Repository][Insert] failed insert complain: %s", err.Error())
		return
	}
	return
}

func (c *complainRepository) InsertResolution(ctx context.Context, args models.InsertResolutionParams, complainID, adminID int64) (err error) {
	err = c.conn.QueryRowContext(
		ctx,
		queries.InsertResolution,
		complainID,
		adminID,
		args.Remark,
	).Err()

	if err != nil {
		err = fmt.Errorf("[Repository][Insert] failed insert complain: %s", err.Error())
		return
	}
	return
}

func (c *complainRepository) UpdateStatus(ctx context.Context, status string, complainID int64) (err error) {
	fmt.Println("kocak")

	err = c.conn.QueryRowContext(
		ctx,
		queries.UpdateStatusComplain,
		status,
		complainID,
	).Err()

	if err != nil {
		err = fmt.Errorf("[Repository][Insert] failed insert complain: %s", err.Error())
		return
	}
	return

}
