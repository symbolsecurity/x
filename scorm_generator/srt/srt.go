package srt

import (
	"strings"

	gosrt "github.com/konifar/go-srt"
)

func ReadTranscript(filename string) (string, error) {
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
