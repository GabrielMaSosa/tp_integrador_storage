package pkg

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

func RecoveryCustomers() ([]domain.Customers, error) {
	data, err := os.Open("./datos/customers.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	slice := []domain.Customers{}
	json.Unmarshal(dataRead, &slice)

	return slice, nil
}

func RecoveryInvoices() ([]domain.Invoices, error) {
	data, err := os.Open("./datos/invoices.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	slice := []domain.Invoices{}
	json.Unmarshal(dataRead, &slice)

	return slice, nil
}

func RecoverySales() ([]domain.Sales, error) {
	data, err := os.Open("./datos/sales.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	slice := []domain.Sales{}
	json.Unmarshal(dataRead, &slice)

	return slice, nil
}

func RecoveryProducts() ([]domain.Product, error) {
	data, err := os.Open("./datos/products.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	slice := []domain.Product{}
	json.Unmarshal(dataRead, &slice)

	return slice, nil
}
