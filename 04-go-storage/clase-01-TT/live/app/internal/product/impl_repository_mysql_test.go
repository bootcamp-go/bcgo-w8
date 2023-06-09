package product

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	// env
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	txdb.Register("txdb", "mysql", os.Getenv("DB"))

	// cfg := mysql.Config{
	// 	ParseTime: true,
	// }
	// txdb.Register("txdb", "mysql", "root:password@tcp(localhost:3306)/product_db?parseTime=true")
}

// Test
func TestRepositoryMySQL_GetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// arrange
		// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/product_db?parseTime=true")
		db, err := sql.Open("txdb", "identifier")
		assert.NoError(t, err)
		defer db.Close()

		rp := NewRepositoryMySQL(db)

		exp := []*Product{
			{ID: 1, Name: "Strawberry", Type: "Fruit", Count: 10, Price: 100.00, WarehouseId: 1},
			{ID: 2, Name: "Pepsi", Type: "Drinks", Count: 20, Price: 200.00, WarehouseId: 1},
		}

		// act
		pr, err := rp.GetAll()

		t.Log(pr[0])

		// assert
		assert.NoError(t, err)
		assert.Equal(t, exp, pr)
	})
}

func TestRepositoryMySql_GetFull(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		// arrange
		db, err := sql.Open("txdb", "product_db")
		assert.NoError(t, err)
		defer db.Close()
		
		st := NewRepositoryMySQL(db)

		exp := &ProductFull{
			Product: Product{ID: 1, Name: "Strawberry", Type: "Fruit", Count: 10, Price: 100.00, WarehouseId: 1},
			WarehouseName: "Fresh Products",
			WarehouseAddress: "Av Not Known, 1234",
		}

		// act
		pr, err := st.GetFullData(1)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, exp, pr)
	})

	t.Run("failed, not found", func(t *testing.T) {
		// arrange
		db, err := sql.Open("txdb", "product_db")
		assert.NoError(t, err)
		defer db.Close()

		st := NewRepositoryMySQL(db)

		// act
		pr, err := st.GetFullData(100)

		// assert
		assert.ErrorIs(t, err, ErrRepositoryNotFound)
		assert.Empty(t, pr)
	})
}

func TestRepositoryMySql_Create(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		// arrange
		db, err := sql.Open("txdb", "product_db")
		assert.NoError(t, err)
		defer db.Close()

		rp := NewRepositoryMySQL(db)

		pr := &Product{
			Name: "Coca Cola",
			Type: "Drinks",
			Count: 10,
			Price: 400.00,
			WarehouseId: 1,
		}

		// act
		err = rp.Create(pr)

		// assert
		assert.NoError(t, err)
		// assert.Equal(t, 3, pr.ID)
	})

	t.Run("failed, unexisting foreign key", func(t *testing.T) {
		// arrange
		db, err := sql.Open("txdb", "product_db")
		assert.NoError(t, err)
		defer db.Close()

		rp := NewRepositoryMySQL(db)

		pr := &Product{
			Name: "Coca Cola",
			Type: "Drinks",
			Count: 10,
			Price: 400.00,
			WarehouseId: 100,
		}

		// act
		err = rp.Create(pr)

		// assert
		assert.ErrorIs(t, err, ErrRepositoryForeignKey)
	})
}