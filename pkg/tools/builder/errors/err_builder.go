package error_builder

import "fmt"

func Err(errMsg string, err error) error {
	return fmt.Errorf(errMsg, err)
}
