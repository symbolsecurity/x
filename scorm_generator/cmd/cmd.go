package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Name      string
	Output    string
	Subtitles string
	Video     string
}

func ParseFlags() *Config {
	// Define flags with full names
	video := flag.String("video", "", "Path to the video file")
	subtitles := flag.String("subtitles", "", "Path to the subtitles file")
	name := flag.String("name", "output", "Name for the generated questions (optional)")
	output := flag.String("output", "archive.zip", "Name of the output zip file (optional)")

	// Define aliases for the flags
	flag.StringVar(video, "v", "", "Path to the video file (alias for --video)")
	flag.StringVar(subtitles, "s", "", "Path to the subtitles file (alias for --subtitles)")
	flag.StringVar(name, "n", "output", "Name for the generated questions (optional, alias for --name)")
	flag.StringVar(output, "o", "archive.zip", "Name of the output zip file (optional, alias for --output)")

	flag.Parse()

	if *video == "" || *subtitles == "" {
		fmt.Println("Usage: --video <path> --subtitles <path> [--name <name>] [--output <filename>]")
		fmt.Println("Aliases: -v <path> -s <path> [-n <name>] [-o <filename>]")
		os.Exit(1)
	}

	return &Config{
		Video:     *video,
		Subtitles: *subtitles,
		Name:      *name,
		Output:    *output,
	}
}
