package products

import (
	"database/sql"
	"errors"
	"log"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type TopProduct struct {
	Description string `json:"description"`
	Total       int    `json:"total"`
}
type Repository interface {
	Create(product *domain.Product) (int64, error)
	ReadAll() ([]*domain.Product, error)
	GetTop() (ret []TopProduct, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(product *domain.Product) (int64, error) {
	query := `INSERT INTO products (description, price) VALUES (?, ?)`
	row, err := r.db.Exec(query, &product.Description, &product.Price)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Product, error) {
	query := `SELECT id, description, price FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := make([]*domain.Product, 0)
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

/*
SELECT p.description as name , SUM(s.quantity) as total FROM products p INNER JOIN sales s ON s.product_id = p.id
GROUP BY name ORDER BY total DESC LIMIT 5;
*/
func (r *repository) GetTop() (ret []TopProduct, err error) {
	stmt, err := r.db.Prepare(`
	SELECT p.description as description , SUM(s.quantity) as total FROM products p INNER JOIN sales s ON s.product_id = p.id
GROUP BY description ORDER BY total DESC LIMIT 5
;`)
	if err != nil {
		log.Println("error ", err.Error())
	}
	defer stmt.Close()
	res, err := stmt.Query()
	drivererr, ok := err.(*mysql.MySQLError)
	if ok {
		//atrapamos los errores del driver
		log.Println("Error in create product: ", drivererr.Number, drivererr.Message, drivererr.Error())
		err = errors.New("Internal")
		return
	}
	if err != nil {
		log.Println("Error CREATE execute Query", err.Error())
		err = errors.New("Internal")
		return
	}
	for res.Next() {
		var p TopProduct
		if err = res.Scan(&p.Description, &p.Total); err != nil {
			log.Println(err.Error())
			err = errors.New("Internal")
			return
		}
		ret = append(ret, p)

	}

	return
}
