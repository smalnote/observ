package errors_test

import (
	"errors"
	"net"
	"testing"

	errs "github.com/smalnote/observ/errors"
	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	abortedErr := errs.New(errs.Aborted, "aborted")
	fErr := errs.Newf(errs.Aborted, "aborted")
	assert.True(t, errors.Is(fErr, abortedErr))
	assert.Equal(t, "rpc error: code = Aborted desc = aborted", fErr.Error())

	wClosed := errs.Newf(errs.Canceled, "connection closed: %w", net.ErrClosed)
	assert.True(t, errors.Is(wClosed, net.ErrClosed))
	assert.Equal(t,
		"rpc error: code = Canceled desc = connection closed: use of closed network connection",
		wClosed.Error())
}
