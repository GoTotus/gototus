package totus

// Utility provides methods for utility services in the Totus API
type Utility struct {
	totus *Totus
}

// Balance fetches the current account credit balance
func (u *Utility) Balance() (*Balance, error) {
	return u.totus.Balance()
}

// Authenticated checks if the API key is valid and working
func (u *Utility) Authenticated() (*Authenticated, error) {
	var resp Authenticated
	err := u.totus.makeRequest("GET", "/authenticated", nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
