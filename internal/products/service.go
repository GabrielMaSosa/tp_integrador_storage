package products

import (
	"log"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Service interface {
	Create(product *domain.Product) error
	ReadAll() ([]*domain.Product, error)
	GetTopProducts() (ret []TopProduct, err error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(product *domain.Product) error {
	_, err := s.r.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Product, error) {
	return s.r.ReadAll()
}

func (s *service) GetTopProducts() (ret []TopProduct, err error) {
	ret, err = s.r.GetTop()
	if err != nil {
		log.Println(err)
		return
	}
	return
}
