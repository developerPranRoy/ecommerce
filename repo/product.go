package repo

import (
	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageUrl    string  `json:"imageUrl" db:"image_url"`
}

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) Create(prd domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (title, description, price, image_url)
		VALUES (:title, :description, :price, :image_url)
		RETURNING id;
	`

	rows, err := r.db.NamedQuery(query, prd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&prd.ID)
	}
	return &prd, nil
}

func (r *productRepo) Get(productID int) (*domain.Product, error) {
	query := `SELECT * FROM products WHERE id = $1;`

	var prd domain.Product
	err := r.db.Get(&prd, query, productID)
	if err != nil {
		return nil, err
	}
	return &prd, nil
}

func (r *productRepo) List(page, limit int) ([]*domain.Product, error) {

	query := `SELECT
	id,
	title,
	description,
	price,
	image_url
	FROM products
	LIMIT $1
	OFFSET $2;
	`
	offSet := (((page - 1) * limit) + 1)

	var products []*domain.Product

	err := r.db.Select(&products, query, limit, offSet)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) Delete(productID int) error {
	query := `DELETE FROM products WHERE id = $1;`
	_, err := r.db.Exec(query, productID)
	return err
}

func (r *productRepo) Update(prd domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products
		SET 
			title = :title,
			description = :description,
			price = :price,
			image_url = :image_url
		WHERE id = :id;
	`
	_, err := r.db.NamedExec(query, prd)
	if err != nil {
		return nil, err
	}
	return &prd, nil
}
