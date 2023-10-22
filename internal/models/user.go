package models

type (
	User struct {
		ID       int64
		Name     string
		Password string
		Email    string
		Photo    string
		Role     string
	}

	UpdateProfileUserParams struct {
		Name  string
		Photo string
	}
)
