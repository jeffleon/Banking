package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{ID: "1001", Name: "Jefferson", City: "New Delhi", Zipcode: "11001", DateofBirth: "2000-01-01", Status: "1"},
		{ID: "1002", Name: "Andrea", City: "New Delhi", Zipcode: "11001", DateofBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
