package domain

import (
	"github.com/SupratickDey/banking/dto"
	"github.com/SupratickDey/banking/errs"
)

// Customer - domain object
type Customer struct {
	Id          string `json:"id" db:"customer_id"`
	Name        string `json:"name" db:"name"`
	City        string `json:"city" db:"city"`
	ZipCode     string `json:"zip_code" db:"zipcode"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status" db:"status"`
}

// CustomerRepository - secondary port (driven port)
type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c Customer) StatusToText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.StatusToText(),
	}
}
