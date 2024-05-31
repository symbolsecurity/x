package transcript

import (
	_ "embed"

	"github.com/symbolsecurity/x/scorm_generator/groq"
)

//go:embed system_prompt.txt
var systemPrompt string

func ExtractQuestionsGroq(transcript string) (string, error) {
	llm := groq.NewModel(groq.LLAMA3_8B)

	llm.SystemPrompt(systemPrompt)
	llm.UserPrompt("Extract questions from this transcript:\n" + transcript)

	resp, err := llm.Send()
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
