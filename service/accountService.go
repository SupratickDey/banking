package service

import (
	"github.com/SupratickDey/banking/domain"
	"github.com/SupratickDey/banking/dto"
	"github.com/SupratickDey/banking/errs"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionRequest, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	newAccount, appError := s.repo.Save(a)
	if appError != nil {
		return nil, appError
	}
	newAccountResponse := newAccount.ToNewAccountResponseDto()
	return &newAccountResponse, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}
