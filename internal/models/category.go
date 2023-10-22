package models

import "github.com/jmoiron/sqlx/types"

type (
	Category struct {
		Name                 string               `json:"name"`
		ExtraFieldCategories []ExtraFieldCategory `json:"extra_field_categories"`
	}

	ExtraFieldCategory struct {
		ID           int64          `json:"id"`
		fieldType    string         `json:"field_type"`
		fieldLabel   string         `json:"field_label"`
		fieldOptions types.JSONText `json:"field_options"`
	}
)
