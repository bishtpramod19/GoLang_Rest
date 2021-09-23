package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "P123", Name: "Pramod Singh", City: "Noida", Zipcode: "201303", DateofBirth: "19-July-1988", Status: "1"},
		{Id: "P123", Name: "Priyanka Rawat", City: "Noida", Zipcode: "201303", DateofBirth: "08-July-1991", Status: "1"},
		{Id: "P123", Name: "Prateek Singh", City: "Delhi", Zipcode: "110092", DateofBirth: "10-July-1985", Status: ""},
	}

	return CustomerRepositoryStub{customers: customers}
}
