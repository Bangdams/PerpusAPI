package exception

type DuplicateNamee struct {
	Error string
}

func NewDuplicateName(error string) DuplicateNamee {
	return DuplicateNamee{
		Error: error,
	}
}
