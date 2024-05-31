package groq

type Model interface {
	Send() (Response, error)
	UserPrompt(prompt string)
	SystemPrompt(prompt string)
}

func NewModel(name string) Model {
	return models[name]
}

var models = map[string]Model{
	LLAMA3_8B: newLlama3(),
	// LLAMA3_70B: newLlama3(),
	// MIXTRAL: newMixtral(),
	// GEMMA: newGemma(),
}
