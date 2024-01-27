package constant

import "fmt"

var (
	ErrUsernameExist    = fmt.Errorf("user with the same username exist")
	ErrEmailExist       = fmt.Errorf("user with the same email exist")
	ErrUserNotExist     = fmt.Errorf("user not exist")
	ErrPasswordNotMatch = fmt.Errorf("wrong password")
)
