package customers

import (
	"log"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Service interface {
	Create(customers *domain.Customers) error
	ReadAll() ([]*domain.Customers, error)
	TotalCustomers() (ret []ConditionTotal, err error)
	TopCustomerActive() (ret []CustomerActive, err error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(customers *domain.Customers) error {
	_, err := s.r.Create(customers)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Customers, error) {
	return s.r.ReadAll()
}

func (s *service) TotalCustomers() (ret []ConditionTotal, err error) {
	ret, err = s.r.GetTotalCustomers()
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func (s *service) TopCustomerActive() (ret []CustomerActive, err error) {
	ret, err = s.r.GetTopCustomerActive()
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}
