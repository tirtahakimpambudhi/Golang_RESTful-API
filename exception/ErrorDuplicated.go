package exception

type ErrorDuplcated struct {
	Error string
}

func NewErrorDuplcated (err string) ErrorDuplcated {
	return 	ErrorDuplcated{
		Error: err,
	}
}