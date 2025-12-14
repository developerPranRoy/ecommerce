package users

import "ecommerce/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
}
