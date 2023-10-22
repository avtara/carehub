package queries

const (
	GetAllCategory = `
		SELECT 
		    category_id,
		    category_name
        FROM categories
        WHERE is_active=true AND deleted_at IS NULL
        LIMIT $1`

	GetCategoryByID = `
		SELECT 
		    category_id,
		    category_name
        FROM categories
        WHERE category_id = $1 AND is_active=true AND deleted_at IS NULL
        LIMIT 1`

	GetExtraFieldByCategoryID = `
		SELECT 
			field_id,
			category_id,
			field_type,
			field_label,
			field_options
        FROM extra_field_categories
        WHERE category_id = $1 AND is_active=true AND deleted_at IS NULL`
)
