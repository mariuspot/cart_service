package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConnection struct {
	db *sql.DB
}

func NewDatabaseConnection(username, password, serverAddress, database string) (*DatabaseConnection, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, serverAddress, database))

	// if there is an error opening the connection, handle it
	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{db: db}, nil
}

func (dc *DatabaseConnection) Close() {
	if dc.db != nil {
		dc.db.Close()
	}
}

func (dc *DatabaseConnection) CreateCart() (int64, error) {
	result, err := dc.db.Exec("INSERT INTO carts(created_at, updated_at) VALUES(NOW(),NOW());")
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// id
// product_id
// cart_id
// created_at
// updated_at
// quantity
// price
// order_id

func (dc *DatabaseConnection) AddLineItem(cart_id, product_id, quantity int64) error {
	stmt, err := dc.db.Prepare("SELECT price FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	rows, err := stmt.Query(product_id)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return fmt.Errorf("Product not found")
	}
	var product_price float64
	rows.Scan(&product_price)
	log.Printf("Product price %f", product_price)

	tx, err := dc.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO line_items(product_id, cart_id, created_at, updated_at, quantity, price) VALUES(?, ?, NOW(), NOW(), ?, ?)", product_id, cart_id, quantity, product_price*float64(quantity))
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE carts SET updated_at = NOW() WHERE id = ?", cart_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (dc *DatabaseConnection) RemoveLineItem(cart_id, product_id, quantity int64) error {
	stmt, err := dc.db.Prepare("SELECT id, quantity FROM line_items WHERE cart_id = ? AND product_id = ?")
	if err != nil {
		return err
	}
	rows, err := stmt.Query(cart_id, product_id)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return fmt.Errorf("Product not found")
	}
	var line_item_id, line_item_quantity int64
	rows.Scan(&line_item_id, &line_item_quantity)

	if quantity > line_item_quantity || quantity == 0 {
		stmt, err := dc.db.Prepare("DELETE FROM line_items WHERE id = ?")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(line_item_id)
		if err != nil {
			return err
		}
	} else {
		stmt, err := dc.db.Prepare("UPDATE line_items SET quantity = ? WHERE id = ?")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(line_item_quantity-quantity, line_item_id)
		if err != nil {
			return err
		}
	}
	return nil
}
