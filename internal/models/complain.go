package models

import "github.com/jmoiron/sqlx/types"

type (
	Complain struct {
		ID              int64          `json:"id"`
		Category        string         `json:"category"`
		Description     string         `json:"description"`
		ComplaintDetail types.JSONText `json:"complaint_detail"`
		User            User           `json:"user"`
		Status          string         `json:"status"`
		Resolution      []Resolution   `json:"resolution"`
	}

	Resolution struct {
		ID         int64  `json:"id"`
		ResolvedBy User   `json:"resolved_by"`
		Remark     string `json:"remark"`
	}

	InsertComplainParams struct {
		Description     string         `json:"description" validate:"required"`
		ComplaintDetail types.JSONText `json:"complaint_detail" validate:"required"`
		CategoryID      int64          `json:"category_id" validate:"required"`
	}

	InsertResolutionParams struct {
		Remark string `json:"remark" validate:"required"`
	}

	UpdateStatusComplainParams struct {
		Status string `json:"status" validate:"required"`
	}
)
