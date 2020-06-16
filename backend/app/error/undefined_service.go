package cerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type UndefinedService struct {
	ServiceName string
}

func NewUndefinedService(ServiceName string) error {
	return errors.WithStack(&UndefinedService{ServiceName})
}

func (e *UndefinedService) Error() string {
	return fmt.Sprintf("undefined service : %s", e.ServiceName)
}
