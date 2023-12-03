package helper

import (
	"github.com/AlecAivazis/survey/v2"
)

func Select(label string, options []string) (string, error) {
	var selected string
	prompt := &survey.Select{
		Message: label,
		Options: options,
	}
	err := survey.AskOne(prompt, &selected)
	return selected, err
}
