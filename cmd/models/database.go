package models

import (
	"database/sql"
)

// Anonymously embed the sql.DB connection pool in our Database struct, so we can
// later access its methods from GetProduct().
type Database struct {
	*sql.DB
}

func (db *Database) GetProduct(id int) (*Product, error) {

	stmt := `SELECT product_id, qty FROM stock_levels WHERE product_id = ?`

	row := db.QueryRow(stmt, id)

	s := &Product{}

	err := row.Scan(&s.ProductID, &s.Qty)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return s, nil

}

func (db *Database) GetProducts() (Products, error) {

	stmt := `SELECT product_id, qty FROM stock_levels`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // always close after error checking

	products := Products{}

	for rows.Next() {
		p := &Product{}
		err := rows.Scan(&p.ProductID, &p.Qty)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}
