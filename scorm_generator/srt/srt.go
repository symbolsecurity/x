package srt

import (
	"fmt"
	"strings"

	gosrt "github.com/konifar/go-srt"
)

func ReadTranscript(filename string) (string, error) {
	fmt.Println("Reading transcript file", filename)
	subs, err := gosrt.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var subText []string
	for _, sub := range subs {
		subText = append(subText, sub.Text)
	}

	return strings.Join(subText, "/"), nil
}
