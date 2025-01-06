package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "goexpert_goorm:goexpert_goorm@tcp(localhost:3306)/goexpert_goorm?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Product{})

	// criar um individualmente
	db.Create(&Product{Name: "Laptop2", Price: 1000})

	// criar um bacth

	products := []Product{
		{Name: "Laptop3", Price: 3000},
		{Name: "Laptop4", Price: 4000},
	}

	db.CreateInBatches(products, len(products))
}
