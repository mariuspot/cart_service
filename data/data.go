package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	queryCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "cart_service_query_counter",
		Help: "The total number of queries",
	})
	querySummary = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "cart_service_query_time_summary",
		Help: "The query times summary",
	})
)

type DatabaseConnection struct {
	db *sql.DB
}

func NewDatabaseConnection(username, password, serverAddress, database string) (*DatabaseConnection, error) {
	queryCounter.Inc()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, serverAddress, database))

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
	queryCounter.Inc()
	start := time.Now()
	result, err := dc.db.Exec("INSERT INTO carts(created_at, updated_at) VALUES(NOW(),NOW());")
	if err != nil {
		return 0, err
	}
	querySummary.Observe(time.Since(start).Seconds())
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
	queryCounter.Inc()
	start := time.Now()
	rows, err := dc.db.Query("SELECT price FROM products WHERE id = ?", product_id)
	if err != nil {
		return err
	}
	querySummary.Observe(time.Since(start).Seconds())
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
	queryCounter.Inc()
	start = time.Now()
	_, err = tx.Exec("INSERT INTO line_items(product_id, cart_id, created_at, updated_at, quantity, price) VALUES(?, ?, NOW(), NOW(), ?, ?)", product_id, cart_id, quantity, product_price*float64(quantity))
	if err != nil {
		tx.Rollback()
		return err
	}
	querySummary.Observe(time.Since(start).Seconds())

	queryCounter.Inc()
	start = time.Now()
	_, err = tx.Exec("UPDATE carts SET updated_at = NOW() WHERE id = ?", cart_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	querySummary.Observe(time.Since(start).Seconds())

	return tx.Commit()
}

func (dc *DatabaseConnection) RemoveLineItem(cart_id, product_id, quantity int64) error {

	queryCounter.Inc()
	start := time.Now()
	rows, err := dc.db.Query("SELECT id, quantity FROM line_items WHERE cart_id = ? AND product_id = ?", cart_id, product_id)
	if err != nil {
		return err
	}
	querySummary.Observe(time.Since(start).Seconds())

	if !rows.Next() {
		return fmt.Errorf("Product not found")
	}

	var line_item_id, line_item_quantity int64
	rows.Scan(&line_item_id, &line_item_quantity)

	if quantity > line_item_quantity || quantity == 0 {
		queryCounter.Inc()
		start = time.Now()
		_, err := dc.db.Exec("DELETE FROM line_items WHERE id = ?", line_item_id)
		if err != nil {
			return err
		}
		querySummary.Observe(time.Since(start).Seconds())
	} else {
		queryCounter.Inc()
		start = time.Now()
		_, err := dc.db.Exec("UPDATE line_items SET quantity = ? WHERE id = ?", line_item_quantity-quantity, line_item_id)
		if err != nil {
			return err
		}
		querySummary.Observe(time.Since(start).Seconds())
	}
	return nil
}

func (dc *DatabaseConnection) EmptyCart(cart_id int64) error {
	queryCounter.Inc()
	start := time.Now()
	_, err := dc.db.Exec("DELETE FROM line_items WHERE cart_id = ?", cart_id)
	if err != nil {
		return err
	}
	querySummary.Observe(time.Since(start).Seconds())
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

func (dc *DatabaseConnection) GetLineItems(cart_id int64) ([]LineItem, error) {
	queryCounter.Inc()
	start := time.Now()
	rows, err := dc.db.Query("SELECT p.title, p.description, p.image_url, li.quantity, p.price, li.updated_at FROM line_items li INNER JOIN products p ON li.product_id = p.id WHERE li.cart_id = ?", cart_id)
	if err != nil {
		return nil, err
	}
	querySummary.Observe(time.Since(start).Seconds())

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
