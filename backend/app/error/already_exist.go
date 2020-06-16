package cerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type AlreadyExist struct {
	Key string
}

func NewAlreadyExist(Key string) error {
	return errors.WithStack(&AlreadyExist{Key})
}

func (e *AlreadyExist) Error() string {
	return fmt.Sprintf("already exist : %s", e.Key)
}
