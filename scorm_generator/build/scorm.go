package build

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
)

// SCORM creates a SCORM-compliant zip archive containing the specified video and other required files.
func SCORM(videoPath, outputPath string) error {
	filesToAdd := []struct {
		sourcePath string
		zipPath    string
	}{
		{"template/imsmanifest.xml", "imsmanifest.xml"},
		{"tmp/app.js", "app.js"},
		{"template/index.html", "index_lms.html"},
		{videoPath, "video.mp4"},
	}

	op := path.Join(outputPath, "archive.zip")

	archive, err := os.Create(op)
	if err != nil {
		return fmt.Errorf("failed to create archive: %w", err)
	}
	defer archive.Close()

	w := zip.NewWriter(archive)
	defer func() {
		if cerr := w.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for _, file := range filesToAdd {
		if err := addFileToZip(w, file.sourcePath, file.zipPath); err != nil {
			return fmt.Errorf("failed to add file %s to zip: %w", file.sourcePath, err)
		}
	}

	return nil
}

// addFileToZip adds a file to the given zip.Writer.
func addFileToZip(w *zip.Writer, sourcePath, zipPath string) error {
	file, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", sourcePath, err)
	}
	defer file.Close()

	writer, err := w.Create(zipPath)
	if err != nil {
		return fmt.Errorf("failed to create entry %s in zip: %w", zipPath, err)
	}

	if _, err = io.Copy(writer, file); err != nil {
		return fmt.Errorf("failed to copy file %s to zip: %w", sourcePath, err)
	}

	return nil
}
