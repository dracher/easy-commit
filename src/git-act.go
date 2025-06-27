package src

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type GitAct struct {
}

func NewGitAct() GitAct {
	bin, err := os.Open("/usr/bin/git")
	if err != nil {
		log.Fatal("Can't find git binary: %w", err)
	}
	defer func(bin *os.File) {
		err := bin.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(bin)
	return GitAct{}
}

func (g GitAct) GetGitDiff() (string, error) {
	cmd := exec.Command("git", "diff")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	if string(output) == "" {
		return "Initial commit", nil
	}

	return string(output), nil
}

func (g GitAct) DoGitCommit(commitMessage string) error {
	addCmd := exec.Command("git", "add", ".")
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("fail to add files: %w", err)
	}

	commitCmd := exec.Command("git", "commit", "-m", commitMessage)
	if err := commitCmd.Run(); err != nil {
		return fmt.Errorf("fail to commit: %w", err)
	}
	return nil
}
