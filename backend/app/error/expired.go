package cerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type Expired struct {
}

func NewExpired() error {
	return errors.WithStack(&Expired{})
}

func (e *Expired) Error() string {
	return fmt.Sprintf("expired")
}
