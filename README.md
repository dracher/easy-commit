# Easy Commit

This CLI tool automatically generates informative Git commit messages.  Using OpenAI models, it analyzes diffs to create concise, descriptive messages adhering to best practices.

## Requirements

- Go version 1.21.1 or higher.
- An OpenAI(other compatible provider) API key.

## Installation

1. Ensure you have Go installed on your system (version 1.21+).
2. Clone the repository to your local machine.
3. Navigate to the cloned directory and run `go build && go install` to compile the application.

## Usage

To use see below example:

```shell
./easy-commit -d -k sk-xxxxxxxx -m o4-mini --url https://api.xxxxxxxxxx.com
```
