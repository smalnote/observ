package log_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
)

type readOnceCloser struct {
	r      io.Reader
	closed bool
}

// Read check if close, if not delegate it to original reader.
func (r *readOnceCloser) Read(p []byte) (int, error) {
	if r.closed {
		return 0, errors.New("closed already ")
	}
	return r.r.Read(p)
}

// Close just the the flag to closed.
func (r *readOnceCloser) Close() error {
	if !r.closed {
		r.closed = true
		return nil
	}
	return errors.New("closed already ")
}

func TestCloneRequest(t *testing.T) {
	body := readOnceCloser{r: bytes.NewBuffer([]byte("hello, world！"))}
	req, err := http.NewRequest("POST", "http://example.com", &body)
	if err != nil {
		t.Error(err)
		return
	}
	cdata, err := io.ReadAll(&body)
	if err != nil {
		t.Error(err)
		return
	}
	defer body.Close()
	t.Logf("cloned body: %s", string(cdata))

	cbody := bytes.NewReader(cdata)
	req.Body = io.NopCloser(cbody)

	data, err := io.ReadAll(req.Body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("original body: %s", string(data))
}
