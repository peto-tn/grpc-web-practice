package cerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type NotFound struct {
	Key string
}

func NewNotFound(Key string) error {
	return errors.WithStack(&NotFound{Key})
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("not found : %s", e.Key)
}
