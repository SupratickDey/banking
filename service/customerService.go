package service

import (
	"github.com/SupratickDey/banking/domain"
	"github.com/SupratickDey/banking/dto"
	"github.com/SupratickDey/banking/errs"
)

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// CustomerService - connection between primary to secondary port
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, customer := range customers {
		response = append(response, customer.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := customer.ToDto()
	return &response, nil
}
