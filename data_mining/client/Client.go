package client

type Client struct {
	name          string
	ageClient     int
	gender        string
	salaryClient  float32
	addressClient []AddressClient
}

type FamilyClient struct {
	quantityFamily   int
	hasParents       int
	hasGrandChildren bool
}

type AddressClient struct {
	street     string
	postalCode int
	city       string
	country    string
	state      string
}
