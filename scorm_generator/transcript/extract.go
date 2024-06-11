package transcript

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/symbolsecurity/x/scorm_generator/groq"
)

//go:embed system_prompt.txt
var systemPrompt string

//go:embed system_prompt_es.txt
var systemPromptEs string

func ExtractQuestionsGroq(transcript, lang string) (string, error) {
	llm := groq.NewModel(groq.LLAMA3_8B)

	if lang == "es" {
		llm.SystemPrompt(systemPromptEs)
		llm.UserPrompt("Extrae las preguntas de la siguiente transcripción:\n" + transcript)
	} else {
		llm.SystemPrompt(systemPrompt)
		llm.UserPrompt("Extract the questions from this transcript:\n" + transcript)
	}

	resp, err := llm.Send()
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no questions found")
	}

	d := []byte(resp.Choices[0].Message.Content)

	var questions Questions
	if err := json.Unmarshal(d, &questions); err != nil {
		return "", err
	}

	if err := questions.Validate(); err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
