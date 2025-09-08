package filter

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cli/go-gh/v2"
	"github.com/cli/go-gh/v2/pkg/repository"
)

func CreateList(repo repository.Repository, query string, template string) (err error) {
	fmt.Printf("Fetching PRs for %s\n", getRepoName(repo))
	json, err := getPrs()
	if err != nil {
		return
	}

	var queried *bytes.Buffer
	if query == "" {
		queried = json
	} else {
		queried, err = filterJSON(json, query)
		if err != nil {
			return
		}
	}
	applyTemplate(queried, template)
	fmt.Println(queried.String())
	return nil
}

func getRepoName(repo repository.Repository) string {
	result := ""
	if repo.Host != "" {
		result += repo.Host + ":"
	}
	result += repo.Owner + "/" + repo.Name
	return result
}

func getPrs() (*bytes.Buffer, error) {
	stdout, stderr, err := gh.Exec("pr", "list", "--json="+strings.Join(validArgs, ","))
	if err != nil {
		fmt.Println(stderr.String())
		return &stderr, err
	}
	return &stdout, nil

}
