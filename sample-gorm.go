package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	// gorm.Modelを埋め込むと、ID, CreateAt, UpdateAt, DeleteAt が追加される。
	gorm.Model
	Created int64 `gorm:"autoCreateTime"`
	Code    string
	Price   uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	fmt.Printf("product=%v\n", product)
	db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Printf("product.Price=%v\n", product.Price)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	fmt.Printf("product.Price=%v\n", product.Price)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 300, Code: "F42"}) // non-zero fields
	fmt.Printf("product.Price=%v\n", product.Price)
	db.Model(&product).Updates(map[string]interface{}{"Price": 400, "Code": "F42"})
	fmt.Printf("product.Price=%v\n", product.Price)

	// Delete - delete product
	db.Delete(&product, 1)
}
