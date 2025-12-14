package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Review struct {
	ID        int    `json:"id" db:"id"`
	ProductID int    `json:"product_id" db:"product_id"`
	UserID    int    `json:"user_id" db:"user_id"`
	Rating    int    `json:"rating" db:"rating"`
	Comment   string `json:"comment" db:"comment"`
}

type ReviewRepo interface {
	Create(review Review) (*Review, error)
	List(productID int) ([]Review, error)
	Update(review Review) error
	Delete(id int, userID int) error
	Get(id int) (*Review, error)
	ListAll() ([]Review, error)
}

type reviewRepo struct {
	db *sqlx.DB
}

func NewReviewRepo(db *sqlx.DB) ReviewRepo {
	return &reviewRepo{
		db: db,
	}
}

func (r *reviewRepo) List(productID int) ([]Review, error) {
	var reviews []Review
	query := `SELECT * FROM reviews WHERE product_id = $1 ORDER BY id DESC`
	err := r.db.Select(&reviews, query, productID)
	return reviews, err
}

func (r *reviewRepo) ListAll() ([]Review, error) {
	var reviews []Review
	query := `SELECT * FROM reviews ORDER BY id DESC`
	err := r.db.Select(&reviews, query)
	return reviews, err
}
func (r *reviewRepo) Create(review Review) (*Review, error) {
	query := `
        INSERT INTO reviews (product_id, user_id, rating, comment)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	err := r.db.QueryRowx(query, review.ProductID, review.UserID, review.Rating, review.Comment).Scan(&review.ID)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *reviewRepo) Update(review Review) error {
	dbRev, err := r.Get(review.ID)
	if err != nil {
		return err
	}
	if dbRev == nil {
		return fmt.Errorf("Rerview not found")
	}
	if dbRev.UserID != review.UserID {
		return fmt.Errorf("Permission denied")

	}
	query := `
        UPDATE reviews 
        SET rating = $1, comment = $2, updated_at = NOW()
        WHERE id = $3;
    `
	_, err = r.db.Exec(query, review.Rating, review.Comment, review.ID)
	return err
}

func (r *reviewRepo) Delete(id int, userID int) error {
	query := `
        DELETE FROM reviews
        WHERE id = $1 AND user_id = $2;
    `
	res, err := r.db.Exec(query, id, userID)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("Review not found or permission denied")
	}
	return nil
}

func (r *reviewRepo) Get(id int) (*Review, error) {
	var rev Review
	query := `SELECT * FROM reviews WHERE id=$1`

	err := r.db.Get(&rev, query, id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &rev, nil
}
