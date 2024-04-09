package product

import (
	"database/sql"

	"github.com/TimurNiki/go_api_tutorial/v4/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error){
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", productID)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	for rows.Next() {
		p, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}