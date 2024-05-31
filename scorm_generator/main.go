package main

import (
	"log"
	"sym-video/build"
	"sym-video/cmd"
	"sym-video/envloader"
	"sym-video/srt"
	"sym-video/transcript"
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
	result, err := transcript.ExtractQuestionsGroq(subs)
	if err != nil {
		log.Fatal(err)
	}

	// Build JavaScript file
	err = build.BuildJS(result)
	if err != nil {
		log.Fatal(err)
	}

	// Create SCORM package
	err = build.SCORM(cfg.Video, cfg.Output)
	if err != nil {
		log.Fatal(err)
	}
}
