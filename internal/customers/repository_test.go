package customers

import (
	"database/sql"
	"testing"

	txdb "github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func init() {
	cfg := mysql.Config{
		User:      "user1",
		Passwd:    "secret_password",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "fantasy_products",
		ParseTime: true,
	}

	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}
func TestCustomersRepository(t *testing.T) {
	t.Run("Test happy path conditional Total", func(t *testing.T) {
		//Arrange
		db, err := sql.Open("txdb", "DbRealforTesting")
		assert.NoError(t, err)
		rep := NewRepository(db)
		defer db.Close()
		var expectedr = []ConditionTotal{

			{
				Condition: false,
				Total:     605928.74,
			},
			{
				Condition: true,
				Total:     716791.96,
			},
		}
		//act
		res, err := rep.GetTotalCustomers()

		assert.NoError(t, err)
		//assert
		assert.Equal(t, expectedr, res, "Error")

	})

}
