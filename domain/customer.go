package domain

import (
	"Banking/dto"
	"Banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAstext() string {
	statusAstext := "active"
	if c.Status == "0" {
		statusAstext = "inactive"
	}

	return statusAstext
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAstext(),
	}

}

// secondry port
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
