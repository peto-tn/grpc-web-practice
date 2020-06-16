package cerror

import "github.com/pkg/errors"

type Unknown struct{}

func NewUnknown() error {
	return errors.WithStack(&Unknown{})
}

func (e *Unknown) Error() string {
	return "unknown"
}
