package domain

type CustomerRepositoryStub struct {
	Customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			ID:          "1001",
			Name:        "Abdillah",
			Zipcode:     "12312",
			DateOfBirth: "2000-01-01",
			Status:      "1",
		},
		{
			ID:          "1002",
			Name:        "Faisal",
			Zipcode:     "12312",
			DateOfBirth: "2000-01-01",
			Status:      "1",
		},
	}

	return CustomerRepositoryStub{customers}
}
