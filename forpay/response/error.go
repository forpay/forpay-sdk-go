package response

import "fmt"

// ErrorResponse defines error response structure.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

func (e ErrorResponse) Error() string {
	errMsg := "ServerError"

	errMsg += fmt.Sprintf("\nCode: %d", e.Code)
	errMsg += fmt.Sprintf("\nMsg: %s", e.Msg)
	errMsg += fmt.Sprintf("\nSubCode: %s", e.SubCode)
	errMsg += fmt.Sprintf("\nSubMsg: %s", e.SubMsg)

	return errMsg
}

// IsBusinessFailed tells if request business(action) failed.
func (e *ErrorResponse) IsBusinessFailed() bool {
	return e.Code == 1003
}
