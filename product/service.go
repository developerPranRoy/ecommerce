package product

import (
	"ecommerce/domain"
)

type service struct {
	prdrepo ProductRepo
}

func NewService(prdrepo ProductRepo) Service {
	return &service{
		prdrepo: prdrepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	return svc.prdrepo.Create(p)
}
func (svc *service) Get(productID int) (*domain.Product, error) {
	return svc.prdrepo.Get(productID)

}
func (svc *service) List(page, limit int) ([]*domain.Product, error) {
	return svc.prdrepo.List(page, limit)

}
func (svc *service) Delete(productID int) error {
	return svc.prdrepo.Delete(productID)

}
func (svc *service) Update(p domain.Product) (*domain.Product, error) {
	return svc.prdrepo.Update(p)

}

func (svc *service) Count() (int, error) {
	return svc.prdrepo.Count()

}
