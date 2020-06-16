package cerror

import (
	"github.com/pkg/errors"
)

type Timeout struct {
	Key string
}

func NewTimeout() error {
	return errors.WithStack(&Timeout{})
}

func (e *Timeout) Error() string {
	return "timeout"
}
