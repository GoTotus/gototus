package totus

import "errors"

// Sentinel errors for specific conditions
var (
	ErrMissingAPIKey  = errors.New("API key must be provided or set in TOTUS_KEY environment variable")
	ErrAuthentication = errors.New("authentication failed")
	ErrNotFound       = errors.New("resource not found")
	ErrClient         = errors.New("client-side error")
	ErrServer         = errors.New("server-side error")
)
