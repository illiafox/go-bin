package public

type InternalError struct {
	s string
}

func (i InternalError) Error() string {
	return i.s
}

var Internal = "internal error, please try again"

func NewInternal(err error) error {
	if err == nil {
		return nil
	}
	return InternalError{err.Error()}

}
