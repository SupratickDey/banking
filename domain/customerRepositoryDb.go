package domain

import (
	"database/sql"
	"github.com/SupratickDey/banking/errs"
	"github.com/SupratickDey/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	var err error
	if status == "" {
		findAllSql := `SELECT customer_id, name, date_of_birth, city, zipcode, status FROM freedb_banking.customers;`
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := `SELECT customer_id, name, date_of_birth, city, zipcode, status FROM freedb_banking.customers;`
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while scanning find all customer sql: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	findByIdSql := `SELECT customer_id, name, date_of_birth, city, zipcode, status FROM freedb_banking.customers where customer_id = ?;`
	err := d.client.Get(&c, findByIdSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		log.Println("Error in scanning find customer by id sql: ", err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{
		client: dbClient,
	}
}
