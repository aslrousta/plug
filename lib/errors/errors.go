package errors

import (
	"errors"
	"fmt"
)

// ServiceError is the error returned by service functions.
type ServiceError struct {
	Service string
	Func    string
	Message string
	Cause   error
}

func (e *ServiceError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s.%s: %s: %v", e.Service, e.Func, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s.%s: %s", e.Service, e.Func, e.Message)
}

func (e *ServiceError) Unwrap() error {
	return e.Cause
}

// Service creates a new ServiceError.
func Service(name, fun, msg string, cause error) *ServiceError {
	return &ServiceError{
		Service: name,
		Func:    fun,
		Message: msg,
		Cause:   cause,
	}
}

// NotFoundError defines errors which report a missing entity.
type NotFoundError interface {
	error

	// MissingKey returns the missing key.
	MissingKey() string
}

// IsNotFound returns true if err is a NotFoundError.
func IsNotFound(err error) bool {
	if _, ok := err.(NotFoundError); ok {
		return true
	}
	if err := errors.Unwrap(err); err != nil {
		return IsNotFound(err)
	}
	return false
}

type notFoundError struct {
	Message string
	Key     string
}

func (e *notFoundError) Error() string {
	return e.Message
}

func (e *notFoundError) MissingKey() string {
	return e.Key
}

// NotFound creates a new NotFoundError.
func NotFound(message, key string) NotFoundError {
	return &notFoundError{
		Message: message,
		Key:     key,
	}
}

// IntegrityError defines errors which report a data integrity violation.
type IntegrityError interface {
	error

	// CanRetry returns true if the error is temporary and the callee can
	// attempt to perform the operation again.
	CanRetry() bool
}

// IsIntegrity returns true if err is an IntegrityError.
func IsIntegrity(err error) bool {
	if _, ok := err.(IntegrityError); ok {
		return true
	}
	if err := errors.Unwrap(err); err != nil {
		return IsIntegrity(err)
	}
	return false
}

// IsRetryable retyrns true if err is a retryable IntegrityError.
func IsRetryable(err error) bool {
	if err, ok := err.(IntegrityError); ok {
		return err.CanRetry()
	}
	if err := errors.Unwrap(err); err != nil {
		return IsRetryable(err)
	}
	return false
}

type integrityError struct {
	Message     string
	IsRetryable bool
}

func (e *integrityError) Error() string {
	return e.Message
}

func (e *integrityError) CanRetry() bool {
	return e.IsRetryable
}

// Integrity creates a new IntegrityError.
func Integrity(message string, isRetryable bool) IntegrityError {
	return &integrityError{
		Message:     message,
		IsRetryable: isRetryable,
	}
}

// ValidationError defines errors which report an invalid field.
type ValidationError interface {
	error

	// Field returns that name of the invalid field.
	Field() string
}

// IsValidation returns true if err is a ValidationError.
func IsValidation(err error) bool {
	if _, ok := err.(ValidationError); ok {
		return true
	}
	if err := errors.Unwrap(err); err != nil {
		return IsValidation(err)
	}
	return false
}

type validationError struct {
	Message   string
	FieldName string
}

func (e *validationError) Error() string {
	return e.Message
}

func (e *validationError) Field() string {
	return e.FieldName
}

// Validation creates a new ValidationError.
func Validation(message, field string) ValidationError {
	return &validationError{
		Message:   message,
		FieldName: field,
	}
}
