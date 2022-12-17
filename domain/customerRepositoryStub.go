package domain

import "github.com/SupratickDey/banking/errs"

// CustomerRepositoryStub - secondary adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll(status string) ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Supro", "Kolkata", "700156", "1996-07-01", "1"},
		{"1002", "Avi", "Dhanbad", "828202", "1996-10-28", "1"},
	}
	return CustomerRepositoryStub{customers}
}
