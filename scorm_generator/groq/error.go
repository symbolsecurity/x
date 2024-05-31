package groq

import "encoding/json"

type Error struct {
	Message string
	Kind    string
}

func newError(r []byte) error {
	var err Error

	if err := json.Unmarshal(r, &err); err != nil {
		return err
	}

	return &err
}

func (e *Error) Error() string {
	return e.Kind + ": " + e.Message
}
