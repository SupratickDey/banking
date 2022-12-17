package service

import (
	"github.com/SupratickDey/banking/domain"
	"github.com/SupratickDey/banking/dto"
	"github.com/SupratickDey/banking/errs"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}

func (s DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.NewAccount(request.CustomerId, request.AccountType, request.Amount)
	newAccount, appError := s.repo.Save(a)
	if appError != nil {
		return nil, appError
	}
	newAccountResponse := newAccount.ToNewAccountResponseDto()
	return &newAccountResponse, nil
}

func (s DefaultAccountService) MakeTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	// incoming request validation
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	// server side validation for checking the available balance in the account
	if request.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(request.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(request.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       request.AccountId,
		Amount:          request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}
