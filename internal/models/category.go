package models

import "github.com/jmoiron/sqlx/types"

type (
	Category struct {
		ID                   int64                `json:"id"`
		Name                 string               `json:"name"`
		ExtraFieldCategories []ExtraFieldCategory `json:"extra_field_categories"`
	}

	ExtraFieldCategory struct {
		ID           int64          `json:"id"`
		CategoryID   int64          `json:"-"`
		FieldType    string         `json:"field_type"`
		FieldLabel   string         `json:"field_label"`
		FieldOptions types.JSONText `json:"field_options"`
	}
)
