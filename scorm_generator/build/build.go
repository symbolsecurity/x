package build

import (
	_ "embed"
	"fmt"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/symbolsecurity/x/scorm_generator/template"
)

func BuildJS(result string) (string, error) {
	js, err := template.ReadFile("app.js")
	if err != nil {
		return "", err
	}

	build := api.Transform(string(js), api.TransformOptions{
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Define: map[string]string{
			"QUESTIONS": result,
		},
		TreeShaking: api.TreeShakingTrue,
		LogLevel:    api.LogLevelDebug,
	})

	if len(build.Errors) != 0 {
		return "", fmt.Errorf("Error: %s", build.Errors[0].Text)
	}

	return string(build.Code), nil
}
