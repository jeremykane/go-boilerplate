package errorx

type (
	CustomError struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func NewError(e error, code string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: e.Error(),
	}
}

func (c *CustomError) Error() string {

	if c.Message != "" {
		return c.Message
	}
	return ""
}

func (c *CustomError) GetErrorCode() string {
	return c.Code
}
