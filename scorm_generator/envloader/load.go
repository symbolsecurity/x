package envloader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Load loads environment variables from a file.
func Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Split line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("Invalid line:", line)

			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set the environment variable
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
