package models

type User struct {
	ID        *int    `json:"id"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	CreatedAt *string `json:"created_at"`
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
