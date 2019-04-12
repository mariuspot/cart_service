package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	queryCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "cart_service_database_query_counter",
		Help: "The total number of database queries",
	})
	querySummary = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "cart_service_database_query_time_summary",
		Help: "The database query times summary",
	})
)

type DatabaseConnection struct {
	db *sql.DB
}

func NewDatabaseConnection(username, password, ip string, port int, name string) (*DatabaseConnection, error) {
	queryCounter.Inc()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, ip, port, name))

	// if there is an error opening the connection, handle it
	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{db: db}, nil
}

//Helper methods to add metrics to each database call
func (dc *DatabaseConnection) Query(query string, args ...interface{}) (*sql.Rows, error) {
	queryCounter.Inc()
	start := time.Now()
	defer querySummary.Observe(time.Since(start).Seconds())
	return dc.db.Query(query, args...)
}

func (dc *DatabaseConnection) Exec(query string, args ...interface{}) (sql.Result, error) {
	queryCounter.Inc()
	start := time.Now()
	defer querySummary.Observe(time.Since(start).Seconds())
	return dc.db.Exec(query, args...)
}

func (dc *DatabaseConnection) ExecTx(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	queryCounter.Inc()
	start := time.Now()
	defer querySummary.Observe(time.Since(start).Seconds())
	return tx.Exec(query, args...)
}

func (dc *DatabaseConnection) Close() {
	if dc.db != nil {
		dc.db.Close()
	}
}

//Create a new cart
func (dc *DatabaseConnection) CreateCart() (int64, error) {
	queryCounter.Inc()
	start := time.Now()
	result, err := dc.db.Exec("INSERT INTO carts(created_at, updated_at) VALUES(NOW(),NOW());")
	if err != nil {
		return 0, err
	}
	querySummary.Observe(time.Since(start).Seconds())
	return result.LastInsertId()
}

//Add a new line item to a cart
func (dc *DatabaseConnection) AddLineItem(cart_id, product_id, quantity int64) error {
	rows, err := dc.Query("SELECT price FROM products WHERE id = ?", product_id)
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
	_, err = dc.ExecTx(tx, "INSERT INTO line_items(product_id, cart_id, created_at, updated_at, quantity, price) VALUES(?, ?, NOW(), NOW(), ?, ?)", product_id, cart_id, quantity, product_price*float64(quantity))
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = dc.ExecTx(tx, "UPDATE carts SET updated_at = NOW() WHERE id = ?", cart_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

//Remove a line item from a cart
func (dc *DatabaseConnection) RemoveLineItem(cart_id, product_id, quantity int64) error {

	rows, err := dc.Query("SELECT id, quantity FROM line_items WHERE cart_id = ? AND product_id = ?", cart_id, product_id)
	if err != nil {
		return err
	}

	if !rows.Next() {
		return fmt.Errorf("Product not found")
	}

	var line_item_id, line_item_quantity int64
	rows.Scan(&line_item_id, &line_item_quantity)

	if quantity > line_item_quantity || quantity == 0 {
		_, err := dc.Exec("DELETE FROM line_items WHERE id = ?", line_item_id)
		if err != nil {
			return err
		}
	} else {
		_, err := dc.Exec("UPDATE line_items SET quantity = ? WHERE id = ?", line_item_quantity-quantity, line_item_id)
		if err != nil {
			return err
		}
	}
	return nil
}

//Remove all items from a cart
func (dc *DatabaseConnection) EmptyCart(cart_id int64) error {
	_, err := dc.Exec("DELETE FROM line_items WHERE cart_id = ?", cart_id)
	if err != nil {
		return err
	}
	return nil
}

type LineItem struct {
	Title       string
	Description string
	Image_url   sql.NullString
	Quantity    int64
	Price       float32
	Updated_at  time.Time
}

//Get items currently in the cart
func (dc *DatabaseConnection) GetLineItems(cart_id int64) ([]LineItem, error) {
	rows, err := dc.Query("SELECT p.title, p.description, p.image_url, li.quantity, p.price, li.updated_at FROM line_items li INNER JOIN products p ON li.product_id = p.id WHERE li.cart_id = ?", cart_id)
	if err != nil {
		return nil, err
	}

	lineItems := make([]LineItem, 0)
	var title, description string
	var image_url sql.NullString
	var quantity int64
	var price float32
	var updated_at time.Time
	for rows.Next() {
		err = rows.Scan(&title, &description, &image_url, &quantity, &price, &updated_at)
		if err != nil {
			return nil, err
		}
		lineItems = append(lineItems, LineItem{Title: title, Description: description, Image_url: image_url, Quantity: quantity, Price: price, Updated_at: updated_at})
	}
	return lineItems, nil
}

//Convert the cart into an order
func (dc *DatabaseConnection) ConvertCartToOrder(cart_id int64, name, address, email string, pay_type int32) (int64, error) {
	tx, err := dc.db.Begin()
	if err != nil {
		return 0, err
	}
	result, err := dc.ExecTx(tx, "INSERT INTO orders(name, address, email, pay_type, created_at, updated_at) VALUES(?, ?, ?, ?, NOW(), NOW())", name, address, email, pay_type)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	order_id, _ := result.LastInsertId()

	result, err = dc.ExecTx(tx, "UPDATE line_items SET order_id = ?, updated_at = NOW(), cart_id = NULL WHERE cart_id = ?", order_id, cart_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if rows_affected, err := result.RowsAffected(); err != nil {
		tx.Rollback()
		return 0, err
	} else if rows_affected == 0 {
		tx.Rollback()
		return 0, errors.New("Cart is empty")
	}

	result, err = dc.ExecTx(tx, "DELETE FROM carts WHERE id = ?", cart_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if rows_affected, err := result.RowsAffected(); err != nil {
		tx.Rollback()
		return 0, err
	} else if rows_affected == 0 {
		tx.Rollback()
		return 0, errors.New("Cart doesnt exist")
	}

	tx.Commit()
	return order_id, nil

}
