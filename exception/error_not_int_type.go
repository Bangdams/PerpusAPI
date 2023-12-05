package exception

type NotIntType struct {
	Error string
}

func NewNotIntType(error string) NotIntType {
	return NotIntType{
		Error: error,
	}
}
