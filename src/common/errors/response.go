package errors

type Response struct {
	HttpCode uint
	Message  string
}

func (err *Response) Error() string {
	return err.Message
}
