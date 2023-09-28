package customers

import (
	"database/sql"
	"errors"
	"log"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type ConditionTotal struct {
	Condition bool    `json:"condition"`
	Total     float64 `json:"total"`
}
type CustomerActive struct {
	LastName  string  `json:"last_name"`
	FirstName string  `json:"first_name"`
	Amount    float64 `json:"amount"`
}

type Repository interface {
	Create(customers *domain.Customers) (int64, error)
	ReadAll() ([]*domain.Customers, error)
	GetTotalCustomers() (ret []ConditionTotal, err error)
	GetTopCustomerActive() (ret []CustomerActive, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(customers *domain.Customers) (int64, error) {
	query := `INSERT INTO customers (first_name, last_name, customers.condition) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &customers.FirstName, &customers.LastName, &customers.Condition)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Customers, error) {
	query := `SELECT id, first_name, last_name, customers.condition FROM customers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customers := make([]*domain.Customers, 0)
	for rows.Next() {
		customer := domain.Customers{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Condition)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	return customers, nil
}

/*
SELECT c.condition as condicion , TRUNCATE(SUM(i.total),2) as total FROM customers c INNER JOIN invoices i ON c.id = i.customer_id
GROUP BY condicion;

*/

func (r *repository) GetTotalCustomers() (ret []ConditionTotal, err error) {
	stmt, err := r.db.Prepare(
		`SELECT c.condition as condicion , TRUNCATE(SUM(i.total),2) as total FROM customers c INNER JOIN invoices i ON c.id = i.customer_id
		GROUP BY condicion;`)
	if err != nil {
		log.Println("error get total customers", err.Error())
	}
	defer stmt.Close()

	res, err := stmt.Query()
	drivererr, ok := err.(*mysql.MySQLError)
	if ok {
		//atrapamos los errores del driver
		log.Println("error get total customers query: ", drivererr.Number, drivererr.Message, drivererr.Error())
		err = errors.New("Internal")
		return
	}
	if err != nil {
		log.Println("error get total customers query", err.Error())
		err = errors.New("Internal")
		return
	}
	for res.Next() {
		var p ConditionTotal
		if res.Scan(&p.Condition, &p.Total); err != nil {
			log.Println("error get total customers scan", err.Error())
			err = errors.New("Internal")
			return
		}
		ret = append(ret, p)

	}

	return
}

/*
SELECT c.first_name as first_name ,c.last_name as last_name, SUM(i.total) as amount FROM customers c
INNER JOIN invoices i ON p.id = i.customer_id
GROUP BY p.id ORDER BY SUM(i.total) DESC
;

*/

func (r *repository) GetTopCustomerActive() (ret []CustomerActive, err error) {
	stmt, err := r.db.Prepare(
		`SELECT c.first_name as first_name ,c.last_name as last_name, SUM(i.total) as amount FROM customers c 
		INNER JOIN invoices i ON c.id = i.customer_id
		GROUP BY c.id ORDER BY SUM(i.total) DESC LIMIT 5;`)
	if err != nil {
		log.Println("error get total customers", err.Error())
	}
	defer stmt.Close()

	res, err := stmt.Query()
	drivererr, ok := err.(*mysql.MySQLError)
	if ok {
		//atrapamos los errores del driver
		log.Println("error get total GetTopCustomerActive query: ", drivererr.Number, drivererr.Message, drivererr.Error())
		err = errors.New("Internal")
		return
	}
	if err != nil {
		log.Println("error get total GetTopCustomerActive query", err.Error())
		err = errors.New("Internal")
		return
	}
	for res.Next() {
		var p CustomerActive
		if res.Scan(&p.FirstName, &p.LastName, &p.Amount); err != nil {
			log.Println("error get total customers scan", err.Error())
			err = errors.New("Internal")
			return
		}
		ret = append(ret, p)

	}

	return

}
