package models

import "github.com/jmoiron/sqlx/types"

type (
	Complain struct {
		ID              int64          `json:"id"`
		Description     string         `json:"name"`
		ComplaintDetail types.JSONText `json:"complaint_detail""`
		User            User           `json:"user"`
		Status          string         `json:"status"`
		Resolution      []Resolution   `json:"resolution""`
	}

	Resolution struct {
		ID         int64  `json:"id"`
		ResolvedBy User   `json:"resolved_by"`
		Remark     string `json:"remark"`
	}
)
