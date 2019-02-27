package errors

// handle represents an error raised by ErrReturn.
// it is here to provide a diference between
type handleError struct {
	err error
}

// only used with Handle to avoid repeating if err != nil
// will panic if err != nil, should be caught by Handle.
func ErrReturn(err error) {
	if err != nil {
		panic(handleError{err})
	}
}

// if there is a panic and that panic is an error then
// set err to it, since err is a named return value this effectively
// changes our function to return this error.
//
// must be called inside defer.
func Handle(errp *error) {
	e := recover()
	if e == nil {
		return
	}

	// we only deal with handleError, everything else is a valid panic.
	if x, ok := e.(handleError); ok {
		*errp = x.err
		return
	}

	panic(e)
}
