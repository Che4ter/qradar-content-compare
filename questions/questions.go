package questions

import (
	"github.com/AlecAivazis/survey/v2"
	"strings"
)

var initQuestions = []*survey.Question{
	{
		Name: "baseUrlOldQRadar",
		Prompt: &survey.Input{
			Message: "Enter OLD QRadar Base Url:",
		},
		Validate: survey.Required,
	},
	{
		Name: "securityTokenOldQRadar",
		Prompt: &survey.Input{
			Message: "Enter Security Token for OLD QRadar:",
		},
		Validate: survey.Required,
	},
	{
		Name: "baseUrlNewQRadar",
		Prompt: &survey.Input{
			Message: "Enter NEW QRadar Base Url:",
		},
		Validate: survey.Required,
	},
	{
		Name: "securityTokenNewQRadar",
		Prompt: &survey.Input{
			Message: "Enter Security Token for NEW QRadar:",
		},
		Validate: survey.Required,
	},
}

func AskForConnectionDetails() (string, string, string, string, error) {
	answers := struct {
		BaseUrlOldQRadar  string
		SecurityTokenOldQRadar string
		BaseUrlNewQRadar  string
		SecurityTokenNewQRadar string
	}{}

	// ask the question
	if err := survey.Ask(initQuestions, &answers); err != nil {
		return "", "", "", "", err
	}

	if !strings.HasSuffix(answers.BaseUrlOldQRadar, "/"){
		answers.BaseUrlOldQRadar = answers.BaseUrlOldQRadar + "/"
	}
	if !strings.HasPrefix(answers.BaseUrlOldQRadar, "http://") && !strings.HasPrefix(answers.BaseUrlOldQRadar, "https://") {
		answers.BaseUrlOldQRadar = "https://" + answers.BaseUrlOldQRadar
	}

	if !strings.HasSuffix(answers.BaseUrlNewQRadar, "/"){
		answers.BaseUrlNewQRadar = answers.BaseUrlNewQRadar + "/"
	}
	if !strings.HasPrefix(answers.BaseUrlNewQRadar, "http://") && !strings.HasPrefix(answers.BaseUrlNewQRadar, "https://") {
		answers.BaseUrlNewQRadar = "https://" + answers.BaseUrlNewQRadar
	}

	return strings.TrimSpace(answers.BaseUrlOldQRadar), strings.TrimSpace(answers.SecurityTokenOldQRadar), strings.TrimSpace(answers.BaseUrlNewQRadar), strings.TrimSpace(answers.SecurityTokenNewQRadar), nil
}

func AskForReportSelection(reports []string) ([]string, error) {
	var prompt = []*survey.Question{
		{
			Name: "letter",
			Prompt: &survey.MultiSelect{
				Message:  "Select Reports",
				Options:  reports,
				PageSize: 50,
			},
			Validate: survey.Required,
		},
	}
	var selections []string

	if err := survey.Ask(prompt, &selections); err != nil {
		return nil, err
	}

	return selections, nil
}

func AskForFullReport() (bool, error) {
	fullReport := false
	prompt := &survey.Confirm{
		Message: "Do you want to run a full report?",
	}
	if err := survey.AskOne(prompt, &fullReport); err != nil {
		return false, err
	}
	return fullReport, nil
}