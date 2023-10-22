package queries

const (
	GetAllComplain = `
		SELECT 
		    c.complain_id,
		    c.complain_description,
		    c.complain_details,
		    c.status,
		    u.name,
		    u.email,
		    u.photo
        FROM complains c
        JOIN users u
        ON u.user_id = c.user_id
        WHERE c.is_active=true AND c.deleted_at IS NULL
        LIMIT $1`

	GetComplainByID = `
		SELECT 
		    c.complain_id,
		    c.complain_description,
		    c.complain_details,
		    c.status,
		    u.name,
		    u.email,
		    u.photo
        FROM complains c
        JOIN users u
        ON u.user_id = c.user_id
        WHERE complain_id = $1 AND c.is_active=true AND c.deleted_at IS NULL
        LIMIT 1`

	GetResolutionByID = `
		SELECT 
		    cr.resolution_id,
			cr.remark,
			s.name,
			s.email,
			s.photo
        FROM complain_resolutions cr
        JOIN users s
        ON s.user_id = cr.resolved_by
        WHERE cr.complain_id = $1 AND cr.is_active=true AND cr.deleted_at IS NULL
        ORDER BY cr.resolution_id DESC`
)
