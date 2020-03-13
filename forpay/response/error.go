package response

import (
	"fmt"
)

// Error represents forpay api error structure.
type Error struct {
	HTTPStatus int

	Code    int
	Msg     string
	SubCode string
	SubMsg  string
}

func (err *Error) Error() string {
	errMsg := "ServerError"

	errMsg += fmt.Sprintf("\nCode: %d", err.Code)
	errMsg += fmt.Sprintf("\nMsg: %s", err.Msg)
	errMsg += fmt.Sprintf("\nSubCode: %s", err.SubCode)
	errMsg += fmt.Sprintf("\nSubMsg: %s", err.SubMsg)

	return errMsg
}

// IsBusinessFailed tells if request business(action) failed.
func (err *Error) IsBusinessFailed() bool {
	return err.Code == 1003
}
