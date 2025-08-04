# Snake-Case-Sanitizer
> Yes, the irony of the kebab-case repository name is not lost on me.

This is a simple CLI tool that will take in a given directory and recursively comb through it,
renaming all files and folders to follow `snake_case` conventions. 

The purpose of this project is to explore and get more familiar with the Go language.

## Installation
Make sure Go is installed on your computer, then, simply execute the command:
`go install github.com/Wholesome100/snake-case-sanitizer@latest`

## Usage
The utility can be invoked from the command line like so:
`snake-case-sanitizer -path <FILENAME>`

The `-path` flag specifies the path of the file or directory you want to convert.

⚠️ **NOTE**: Be careful when passing in directory names; the utility will comb through ALL files and subdirectories, converting them to `snake_case`.

### License
All Rights Reserved.
