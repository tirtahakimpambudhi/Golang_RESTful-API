package exception


type ErrorNotFound struct {
	Error string
}

func NewErrorNotFound (err string) ErrorNotFound {
	return 	ErrorNotFound{
		Error: err,
	}
}