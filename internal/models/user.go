package models

type (
	User struct {
		ID       int64  `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		Password string `json:"-"`
		Email    string `json:"email,omitempty"`
		Photo    string `json:"photo,omitempty"`
		Role     string `json:"-"`
	}

	UpdateProfileUserParams struct {
		Name  string
		Photo string
	}
)
