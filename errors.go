package commandjsonio

// ErrEncodeInput indicate failed on encode input into JSON.
type ErrEncodeInput struct {
	err error
}

func (e *ErrEncodeInput) Error() string {
	return "[ErrEncodeInput: " + e.err.Error() + "]"
}

func (e *ErrEncodeInput) Unwrap() error {
	return e.err
}

// ErrDecodeOutput indicate failed on decode output from JSON.
type ErrDecodeOutput struct {
	err error
}

func (e *ErrDecodeOutput) Error() string {
	return "[ErrDecodeOutput: " + e.err.Error() + "]"
}

func (e *ErrDecodeOutput) Unwrap() error {
	return e.err
}
