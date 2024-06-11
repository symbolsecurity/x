package transcript

import (
	_ "embed"
	"fmt"
)

type Question struct {
	Desc    string   `json:"question"`
	Options []string `json:"options"`
	Answer  int      `json:"answer"`
}

type Questions []Question

func (q Questions) Validate() error {
	for _, question := range q {
		if question.Desc == "" {
			return fmt.Errorf("question description cannot be empty")
		}

		if len(question.Options) != 4 {
			return fmt.Errorf("question must have 4 options")
		}

		if question.Answer < 0 || question.Answer > 3 {
			return fmt.Errorf("question answer cannot be empty or invalid")
		}
	}

	return nil
}
