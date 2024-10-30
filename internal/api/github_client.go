package api

import (
	"fmt"
	"github.com/cli/go-gh/v2"
)

type GithubClient struct{}

func NewGithubClient() *GithubClient {
	return &GithubClient{}
}

func (c *GithubClient) CreateIssue(title, body string, label string) error {
	issueUrl, stdErr, err := gh.Exec("issue", "create", "--title", title, "--body", body, "--label", label)
	if err != nil {
		fmt.Println("Error: ", err)
		if stdErr.Len() > 0 {
			fmt.Println("Detail: ", stdErr.String())
		}
		return err
	}
	fmt.Println("issue created, link: ", issueUrl.String())
	return nil
}

func (c *GithubClient) ListIssuesWithLabel(label string) (string, error) {
	issues, stdErr, err := gh.Exec("issue", "list", "--label", label)
	if err != nil {
		fmt.Println("Error: ", err)
		if stdErr.Len() > 0 {
			fmt.Println("Detail: ", stdErr.String())
		}
		return "", err
	}
	return issues.String(), nil
}

func (c *GithubClient) CloseIssue(issueId string) error {
	_, stdErr, err := gh.Exec("issue", "close", issueId)
	if err != nil {
		fmt.Println("Error: ", err)
		if stdErr.Len() > 0 {
			fmt.Println("Detail: ", stdErr.String())
		}
		return err
	}
	return nil
}
