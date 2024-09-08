package errors

type Error interface {
	error
	Status() int
	Print() string
	Error() string
}

// CustomError is a custom error type for Hink
type CustomError struct {
	// StatusCode is the http status code
	StatusCode int `json:"-"`
	// Message is the error message to display
	Message string `json:"message,omitempty"`
	// Details is the error details for debugging
	Details interface{} `json:"details,omitempty"`
}
