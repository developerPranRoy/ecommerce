package repo

import (
	"fmt"

	"ecommerce/domain"
	"ecommerce/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {

	query := `INSERT INTO users (first_name, last_name, email, password, is_owner)
			  VALUES (:first_name, :last_name, :email, :password, :is_owner)
		RETURNING id;`

	var userID int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&userID)
	}
	user.ID = userID
	return &user, nil
}

func (r *userRepo) Find(email, password string) (*domain.User, error) {

	query := `
		SELECT id, first_name, last_name, email, password, is_owner
		FROM users
		WHERE email = $1 AND password = $2;
	`

	var user domain.User
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	return &user, nil
}
