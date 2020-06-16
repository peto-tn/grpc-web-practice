package cerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type UndefinedMethod struct {
	MethodName string
}

func NewUndefinedMethod(MethodName string) error {
	return errors.WithStack(&UndefinedMethod{MethodName})
}

func (e *UndefinedMethod) Error() string {
	return fmt.Sprintf("undefined method : %s", e.MethodName)
}
