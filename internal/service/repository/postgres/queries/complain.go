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
		    u.photo,
		    ca.category_name
        FROM complains c
        JOIN users u
        ON u.user_id = c.user_id
        JOIN categories ca
        ON ca.category_id = c.category_id
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
		    u.photo,
		    ca.category_name
        FROM complains c
        JOIN users u
        ON u.user_id = c.user_id
         JOIN categories ca
        ON ca.category_id = c.category_id
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

	InsertComplain = `
		INSERT INTO 
		    complains (
		        category_id, 
		        user_id, 
		        complain_description,
		        complain_details
		    ) 
		VALUES ($1, $2, $3, $4) 
	`

	InsertResolution = `
		INSERT INTO 
		    complain_resolutions (
		        complain_id, 
		        resolved_by, 
		        remark
		    ) 
		VALUES ($1, $2, $3) 
	`

	UpdateStatusComplain = `
	UPDATE complains
	SET 
	    status = $1
	WHERE 
	    complain_id = $2;`
)
