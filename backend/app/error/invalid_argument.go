package cerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type InvalidArgument struct {
	Key  string
	Data string
}

func NewInvalidArgument(Key, Data string) error {
	return errors.WithStack(&InvalidArgument{Key, Data})
}

func (e *InvalidArgument) Error() string {
	return fmt.Sprintf("invalid argument %s : %s", e.Key, e.Data)
}
