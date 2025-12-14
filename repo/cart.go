package repo

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type CartItem struct {
	ID           int       `json:"id" db:"id"`
	UserID       int       `json:"user_id" db:"user_id"`
	ProductID    int       `json:"product_id" db:"product_id"`
	VariantID    *int      `json:"variant_id,omitempty" db:"variant_id"`
	Quantity     int       `json:"quantity" db:"quantity"`
	PriceAtAdded float64   `json:"price_at_added" db:"price_at_added"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type CartRepo interface {
	CreateCart(item CartItem) (*CartItem, error)
	UpdateCartItem(userID int, itemID int, quantity int) (*CartItem, error)
	RemoveCartItem(userID int, itemID int) error
	ClearCart(userID int) error
	GetCartItems(userID int) ([]CartItem, error)
}

type cartRepo struct {
	db *sqlx.DB
}

func NewCartRepo(db *sqlx.DB) CartRepo {
	return &cartRepo{
		db: db,
	}
}

func (r *cartRepo) CreateCart(item CartItem) (*CartItem, error) {
	query := `
    INSERT INTO cart_items (user_id, product_id, variant_id, quantity, price_at_added, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id, created_at, updated_at`

	now := time.Now()
	var id int
	err := r.db.QueryRow(query,
		item.UserID, item.ProductID, item.VariantID, item.Quantity, item.PriceAtAdded, now, now,
	).Scan(&id, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, err
	}
	item.ID = id
	return &item, nil
}

func (r *cartRepo) UpdateCartItem(userID int, itemID int, quantity int) (*CartItem, error) {
	query := `
    UPDATE cart_items
    SET quantity=$1, updated_at=$2
    WHERE id=$3 AND user_id=$4
    RETURNING id, user_id, product_id, variant_id, quantity, price_at_added, created_at, updated_at`

	var item CartItem
	err := r.db.Get(&item, query, quantity, time.Now(), itemID, userID)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *cartRepo) RemoveCartItem(userID int, itemID int) error {
	_, err := r.db.Exec(`DELETE FROM cart_items WHERE id=$1 AND user_id=$2`, itemID, userID)
	return err
}

func (r *cartRepo) ClearCart(userID int) error {
	_, err := r.db.Exec(`DELETE FROM cart_items WHERE user_id=$1`, userID)
	return err
}

func (r *cartRepo) GetCartItems(userID int) ([]CartItem, error) {
	var items []CartItem
	err := r.db.Select(&items, `SELECT * FROM cart_items WHERE user_id=$1`, userID)
	return items, err
}
