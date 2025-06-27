package cmd

import (
	"context"
	"fmt"
	"github.com/dracher/easy-commit/src"
	"github.com/urfave/cli/v3"
	"log"
)

func Cli() *cli.Command {
	return &cli.Command{
		Name:      "easy-commit",
		Usage:     "Let AI generate commit messages based on `git diff` output.",
		UsageText: "easy-commit [options] <valid git repository root>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "dry-run",
				Value:   false,
				Usage:   "If true will only print generated commit message",
				Aliases: []string{"d"},
			},
			&cli.StringFlag{
				Name:     "api-key",
				Value:    "",
				Usage:    "OpenAI(or other compatible service providers) API key",
				Sources:  cli.EnvVars("OPENAI_APIKEY"),
				Aliases:  []string{"k"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "model",
				Value:   "gpt-3.5-turbo",
				Usage:   "The model to use",
				Aliases: []string{"m"},
			},
			&cli.StringFlag{
				Name:    "base-url",
				Value:   "https://api.openai.com",
				Aliases: []string{"url"},
				Usage:   "Can be any openai compatible service",
			},
			&cli.IntFlag{
				Name:    "message-length",
				Value:   30,
				Usage:   "Maximum length of the commit message in words",
				Aliases: []string{"l"},
			},
			&cli.StringFlag{
				Name:    "prompt-size",
				Aliases: []string{"s"},
				Value:   "small",
				Usage:   "The extent of detail and comprehensiveness of the prompt words, 'small' or 'big'",
			},
		},
		UseShortOptionHandling: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			gitAct := src.NewGitAct()
			diffContent, err := gitAct.GetGitDiff()
			if err != nil {
				log.Fatal(err)
			}
			params := src.Params{
				ApiKey:        cmd.String("api-key"),
				BaseURL:       cmd.String("base-url"),
				ModelName:     cmd.String("model"),
				PromptSize:    cmd.String("prompt-size"),
				MessageLength: cmd.Int("message-length"),
			}

			commitMessage, err := src.GenerateCommitMessage(diffContent, &params)
			if err != nil {
				log.Fatal(err)
			}

			if cmd.Bool("dry-run") {
				fmt.Println("========= commitMessage =========")
				fmt.Println(commitMessage)
			} else {
				fmt.Println("========= Gen commitMessage and perform git action =========")
				if err := gitAct.DoGitCommit(commitMessage); err != nil {
					log.Fatal(err)
				}
			}
			return nil
		},
	}
}
