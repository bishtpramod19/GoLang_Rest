package domain

import (
	"Banking/errs"
	"Banking/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

//adapter
func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO banking.accounts (customer_id, opening_date,account_type, amount, status) values (?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database !")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id from new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database !")

	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank acount : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")

	}

	result, _ := tx.Exec("insert into banking.transactions (account_id, amount,transaction_type,transaction_date) values (?,?,?,?)",
		t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	//updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec("update banking.accounts SET amount = amount - ? where account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("update banking.accounts SET amount = amount - ? where account_id = ?", t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error !")
	}

	//commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction : " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error !")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {

		logger.Error("Error while getting last  transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error !")
	}

	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDB) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from banking.accounts where account_id = ?"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return &account, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
