package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// menggunakan postgres
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "go-challenge"
)

var (
	db  *sql.DB
	err error
)

type Products struct {
	ID          int64
	Name        string
	CreateDate  time.Time
	UpdatedDate time.Time
}

type Variants struct {
	ID          int
	Name        string
	ProductId   int
	Qty         int
	CreateDate  time.Time
	UpdatedDate time.Time
}

func main() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", config)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Connect to DB")

	// createProduct()
	// updateProduct()
	// getProductById()
	// createVariant()
	// updateVariantById()
	// getProductWithVariant()
	// deleteVariantById()

}

func createProduct() {
	var product = Products{}
	current := time.Now()

	sqlStatement := `INSERT INTO products (name, created_at, updated_at)
	VALUES ($1, $2, $3)
	Returning *`

	err = db.QueryRow(sqlStatement, "Gadget", current, current).Scan(&product.ID, &product.Name, &product.CreateDate, &product.UpdatedDate)
	if err != nil {
		panic(fmt.Sprintf("Error while Create Product %s", err.Error()))

	}
	fmt.Printf("New Product Data : %v\n", product)
}

func updateProduct() {
	current := time.Now()

	sqlStatement := `UPDATE products SET name = $2, updated_at = $3 WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 5, "Lifestyle", current)
	if err != nil {
		panic(fmt.Sprintf("Error while Update Product %s", err.Error()))

	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Updated Row: %v\n", count)
}

func getProductById() {
	var product = Products{}

	sqlStatement := `SELECT * from products WHERE id = $1`

	data, err := db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	row := data.QueryRow(3)
	err = row.Scan(&product.ID, &product.Name, &product.CreateDate, &product.UpdatedDate)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product Data : %v\n", product)
}
func createVariant() {
	var variant = Variants{}
	current := time.Now()

	sqlStatement := `INSERT INTO variants (variant_name,quantity,product_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	Returning *`

	err = db.QueryRow(sqlStatement, "Apa aja", 10, 3, current, current).Scan(&variant.ID, &variant.Name, &variant.Qty, &variant.ProductId, &variant.CreateDate, &variant.UpdatedDate)
	if err != nil {
		panic(fmt.Sprintf("Error while Create variant %s", err.Error()))

	}
	fmt.Printf("New Variant Data Data : %v\n", variant)
}

func updateVariantById() {
	current := time.Now()

	sqlStatement := `UPDATE variants SET quantity = $2, variant_name = $4 ,updated_at = $3 WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 5, 20, current, "Shoes")
	if err != nil {
		panic(fmt.Sprintf("Error while Update Variants %s", err.Error()))

	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Updated Row: %v\n", count)
}

func getProductWithVariant() {
	rows, err := db.Query("SELECT products.id, products.name, variants.variant_name, variants.quantity FROM products JOIN variants ON products.id = variants.product_id")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var product Products
		var variant Variants

		err := rows.Scan(&product.ID, &product.Name, &variant.Name, &variant.Qty)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Variant %s with Quantity %d is Product of %s\n", variant.Name, variant.Qty, product.Name)
	}
}

func deleteVariantById() {
	sqlStatement := `DELETE from variants WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 7)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data variant :", count)
}
