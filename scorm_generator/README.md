# Video SCORM/QUIZ Generator

This Golang application generates quizzes from video subtitles and packages them into a ZIP file. The application uses the GROQ API for processing the subtitles.

## Requirements

- Go 1.16 or higher
- Set the `GROQ_API_KEY` environment variable with your GROQ API key.

## Installation

WIP

## Usage
The application requires the following flags:

`--video` or `-v`: Path to the video file.
`--subtitles` or `-s`: Path to the subtitles file.
`--name` or `-n`: (Optional) Name for the generated questions. Defaults to output.
`--output` or `-o`: (Optional) Name of the output ZIP file. Defaults to archive.zip.

