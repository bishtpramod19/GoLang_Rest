package domain

import (
	"Banking/errs"
	"Banking/logger"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRespositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRespositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)

	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=?"
		err = d.client.Select(&customers, findAllSql, status)

	}

	if err != nil {
		logger.Error("Error While querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil

}

func (d CustomerRespositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found !")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")

		}

	}
	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRespositoryDB {
	dbUser := os.Getenv("db_user")
	dbPasswd := os.Getenv("db_password")
	//dbAddr := os.Getenv("db_add")
	//dbPort := os.Getenv("db_port ")
	dbName := os.Getenv("db_name")

	dataSource := fmt.Sprintf("%s:%s@/%s", dbUser, dbPasswd, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	//client, err := sqlx.Open("mysql", "banking:banking@123@/banking")

	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRespositoryDB{client: client}

}
