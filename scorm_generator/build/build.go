package build

import (
	"fmt"

	"github.com/evanw/esbuild/pkg/api"
)

func BuildJS(result string) error {
	build := api.Build(api.BuildOptions{
		EntryPoints:       []string{"template/app.js"},
		LogLevel:          api.LogLevelDebug,
		TreeShaking:       api.TreeShakingTrue,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Bundle:            true,
		Outdir:            "tmp",
		Write:             true,

		Define: map[string]string{
			"QUESTIONS": result,
		},
	})

	if len(build.Errors) != 0 {
		return fmt.Errorf("Error: %s", build.Errors[0].Text)
	}

	return nil
}
