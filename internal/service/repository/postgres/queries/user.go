package queries

const (
	GetUserByEmail = `
		SELECT 
		    user_id,
		    name,
		    password,
		    email,
		    photo,
		    role
        FROM users
        WHERE email = $1 AND is_active=true AND deleted_at IS NULL
        LIMIT 1`

	InsertUser = `
		INSERT INTO 
		    users (
		        name, 
		        password, 
		        email,
		        photo,
		        role
		    ) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING user_id
	`

	GetUserByID = `
		SELECT 
		    user_id,
		    name,
		    password,
		    email,
		    photo,
		    role
        FROM users
        WHERE user_id = $1 AND is_active=true AND deleted_at IS NULL
        LIMIT 1`

	UpdateProfileByID = `
	UPDATE users
	SET 
	    name = $1,
    	photo = $2
	WHERE 
	    user_id = $3;`
)
