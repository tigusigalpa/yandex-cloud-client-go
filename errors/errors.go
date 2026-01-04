package errors

import "fmt"

// YandexCloudError is the base error type for all Yandex Cloud errors
type YandexCloudError struct {
	Message string
	Err     error
}

func (e *YandexCloudError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *YandexCloudError) Unwrap() error {
	return e.Err
}

// AuthenticationError represents authentication-related errors
type AuthenticationError struct {
	YandexCloudError
}

func NewAuthenticationError(message string, err error) *AuthenticationError {
	return &AuthenticationError{
		YandexCloudError: YandexCloudError{
			Message: message,
			Err:     err,
		},
	}
}

// APIError represents API-related errors
type APIError struct {
	YandexCloudError
	StatusCode int
}

func NewAPIError(message string, statusCode int, err error) *APIError {
	return &APIError{
		YandexCloudError: YandexCloudError{
			Message: message,
			Err:     err,
		},
		StatusCode: statusCode,
	}
}

// ValidationError represents validation errors
type ValidationError struct {
	YandexCloudError
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		YandexCloudError: YandexCloudError{
			Message: message,
		},
	}
}
