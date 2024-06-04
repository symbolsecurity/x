package template

import "embed"

//go:embed *
var f embed.FS

// ReadFile reads the file at the given path from the embedded filesystem.
func ReadFile(path string) ([]byte, error) {
	return f.ReadFile(path)
}
