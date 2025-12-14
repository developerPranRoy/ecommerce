package repo

import "github.com/jmoiron/sqlx"

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageUrl    string  `json:"imageUrl" db:"image_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) Create(prd Product) (*Product, error) {
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

func (r *productRepo) Get(productID int) (*Product, error) {
	query := `SELECT * FROM products WHERE id = $1;`

	var prd Product
	err := r.db.Get(&prd, query, productID)
	if err != nil {
		return nil, err
	}
	return &prd, nil
}

func (r *productRepo) List() ([]*Product, error) {
	query := `SELECT * FROM products;`

	var products []*Product
	err := r.db.Select(&products, query)
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

func (r *productRepo) Update(prd Product) (*Product, error) {
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
