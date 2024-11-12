package errorsx

import "errors"

var (
	DriverDoesNotExistError = errors.New("driver does not exist")
	UserDoesNotExistError   = errors.New("user does not exist")
	PasswordMismatchError   = errors.New("password mismatch")
	WalletDoesNotExistError = errors.New("wallet does not exist")
	NotEnoughMoneyError     = errors.New("not money fund")
	WalletMismatchError     = errors.New("wallet mismatch")
)
