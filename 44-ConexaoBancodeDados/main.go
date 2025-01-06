package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
)

type Product struct {
	ID       string
	Name     string
	Price    float64
	createAt time.Time `db:"created_at"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:       uuid.New().String(),
		Name:     name,
		Price:    price,
		createAt: time.Now(),
	}

}

func insertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.ID, p.Name, p.Price, p.createAt)

	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Price, p.ID)

	if err != nil {
		return err
	}

	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price, created_at FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price, &p.createAt)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query("SELECT id, name, price, created_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.createAt)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := sql.Open("mysql", "goexpert:goexpert@tcp(127.0.0.1:3306)/goexpert?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Inserindo um produto
	product := NewProduct("camisa", 10.0)
	err = insertProduct(db, product)

	if err != nil {
		panic(err)
	}

	// Atualizando o produto
	product.Price = 20.0
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// Selecionar por id
	p, err := selectProduct(db, product.ID)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Product: %v preço: %v\n", p.Name, p.Price)

	allProducts, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	for _, product := range allProducts {
		fmt.Printf("Product: %v preço: %v criando em: %v\n", product.Name, product.Price, product.createAt)
	}

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

}
