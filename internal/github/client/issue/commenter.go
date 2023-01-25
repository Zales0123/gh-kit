package issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SyliusLabs/gh-kit/internal/factory"
	"github.com/SyliusLabs/gh-kit/internal/github/current_repository"
	"github.com/cli/go-gh/pkg/api"
)

var client api.RESTClient

func init() {
	client = factory.CreateGitHubClient()
}

func AddComment(repo current_repository.CurrentRepository, issueNumber int, message string) error {
	request := struct {
		Body string `json:"body"`
	}{
		Body: message,
	}
	b := new(bytes.Buffer)
	_ = json.NewEncoder(b).Encode(request)

	return client.Post(fmt.Sprintf("repos/%s/%s/issues/%d/comments", repo.Owner, repo.Name, issueNumber), b, nil)
}
