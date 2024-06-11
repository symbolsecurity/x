package main

import (
	"log"

	"github.com/symbolsecurity/x/scorm_generator/build"
	"github.com/symbolsecurity/x/scorm_generator/cmd"
	"github.com/symbolsecurity/x/scorm_generator/envloader"
	"github.com/symbolsecurity/x/scorm_generator/srt"
	"github.com/symbolsecurity/x/scorm_generator/transcript"
)

func main() {
	// Load environment variables
	if err := envloader.Load(".env"); err != nil {
		log.Fatal(err)
	}

	// Parse command line flags
	config := cmd.ParseFlags()

	// Generate SCORM package
	generateSCORM(config)

	log.Println("SCORM package created successfully")
}

func generateSCORM(cfg *cmd.Config) {
	// Read subtitles from SRT file
	subs, err := srt.ReadTranscript(cfg.Subtitles)
	if err != nil {
		log.Fatal(err)
	}

	// Extract questions from subtitles
	result, err := transcript.ExtractQuestionsGroq(subs, cfg.Lang)
	if err != nil {
		log.Fatal(err)
	}

	// Create SCORM package
	err = build.SCORM(cfg.Video, cfg.Output, result)
	if err != nil {
		log.Fatal(err)
	}
}
