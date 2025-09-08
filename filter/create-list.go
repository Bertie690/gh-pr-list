package filter

import (
	"fmt"

	"github.com/cli/go-gh/v2"
	// "github.com/cli/go-gh/v2/pkg/jq"

	"github.com/cli/go-gh/v2/pkg/repository"
	// "github.com/cli/go-gh/v2/pkg/template"
)

func CreateList(repo repository.Repository, args []string) error {
	fmt.Print(repo)
	gh.Exec("pr", "list", "")
	return nil
}
