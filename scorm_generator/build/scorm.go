package build

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/symbolsecurity/x/scorm_generator/template"
)

// SCORM creates a SCORM-compliant zip archive containing the specified video and other required files.
func SCORM(videoPath, outputPath, js string) error {
	filesToAdd := []string{
		"index.html",
		"imsmanifest.xml",
	}

	op := path.Join(outputPath, "archive.zip")

	archive, err := os.Create(op)
	if err != nil {
		return fmt.Errorf("failed to create archive: %w", err)
	}
	defer archive.Close()

	w := zip.NewWriter(archive)
	defer w.Close()

	if err := addJS(w, js); err != nil {
		return fmt.Errorf("failed to add JavaScript code to zip: %w", err)
	}

	if err := addVideo(w, videoPath); err != nil {
		return fmt.Errorf("failed to add video to zip: %w", err)
	}

	for _, file := range filesToAdd {
		f, err := template.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read template file: %w", err)
		}

		writer, err := w.Create(file)
		if err != nil {
			return fmt.Errorf("failed to create file in zip: %w", err)
		}

		if _, err := writer.Write(f); err != nil {
			return fmt.Errorf("failed to write file to zip: %w", err)
		}
	}

	return nil
}

// addJS adds the JavaScript code to the zip archive.
func addJS(w *zip.Writer, js string) error {
	writer, err := w.Create("app.js")
	if err != nil {
		return err
	}

	if _, err := writer.Write([]byte(js)); err != nil {
		return err
	}

	return nil
}

func addVideo(w *zip.Writer, videoPath string) error {
	file, err := os.Open(videoPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer, err := w.Create("video.mp4")
	if err != nil {
		return err
	}

	if _, err = io.Copy(writer, file); err != nil {
		return err
	}

	return nil
}
