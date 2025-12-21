package products

import "ecommerce/domain"

type Service interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List(page, limit int) ([]*domain.Product, error)
	Delete(productID int) error
	Update(p domain.Product) (*domain.Product, error)
	Count() (int, error)
}
