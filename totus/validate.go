package totus

import "net/url"

// Validate provides methods for validation services in the Totus API
type Validate struct {
	totus *Totus
}

// Email validates an email address with the specified check level
func (v *Validate) Email(email string, level CheckLevel) (*ValidatedEmail, error) {
	q := url.Values{
		"email": []string{email},
		"level": []string{string(level)},
	}

	var resp ValidatedEmail
	err := v.totus.makeRequest("GET", "/validate/email", q, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
