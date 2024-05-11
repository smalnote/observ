// Package errors implements unify domain errors handling that wraps gRPC status.
package errors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/status"
)

var (

	// As synonym of std errors.As().
	As = errors.As
	// Is synonym of std errors.Is().
	Is = errors.Is
	// Join synonym of std errors.Join().
	Join = errors.Join
	// Unwrap synonym of std errors.Unwrap().
	Unwrap = errors.Unwrap
	// ErrUnsupported synonym of std errors.ErrUnsupported.
	ErrUnsupported = errors.ErrUnsupported
	// Simple synonym of std errors.New().
	Simple = errors.New
)

// New creates a error, code OK will return nil.
func New(c Code, msg string) error {
	return status.Error(c, msg)
}

// Newf returns New(c, fmt.Errorf(format, a).Error())
// with wrapped formatted err.
func Newf(c Code, format string, a ...any) error {
	err := fmt.Errorf(format, a...)
	return &wrapError{
		Status: status.New(c, err.Error()),
		err:    err,
	}
}

type wrapError struct {
	*status.Status
	err error
}

func (e *wrapError) Error() string {
	return e.Status.String()
}

func (e *wrapError) Is(err error) bool {
	switch x := err.(type) {
	case interface{ GRPCStatus() *status.Status }:
		es := x.GRPCStatus()
		return e.Status.Code() == es.Code() && e.Status.Message() == es.Message()
	default:
		return errors.Is(e.err, x)
	}
}

func (e *wrapError) GRPCStatus() *status.Status {
	return e.Status
}
