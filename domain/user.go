package domain

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Isowner   bool   `json:"is_owner" db:"is_owner"`
}

func (u *User) ChangePassword(newPass string) {
    u.Password = newPass
}
