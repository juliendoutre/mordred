package mordred

import (
	"strings"

	"github.com/juliendoutre/mordred/pkg/git"
)

// Identity contains personal data of a repository author or commiter.
type Identity struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// GetIdentities returns a slice of Identity found in a repository commits.
func GetIdentities(repository *git.Repository) ([]Identity, error) {
	rawOutput, err := repository.ListIdentities()
	if err != nil {
		return nil, err
	}

	cache := map[string]bool{}

	identities := []Identity{}
	for _, line := range strings.Split(string(rawOutput), "\n") {
		if _, ok := cache[line]; !ok {
			cache[line] = true

			tokens := strings.Split(line, ",")
			if len(tokens) > 1 {
				identities = append(identities, Identity{
					Name:  tokens[0],
					Email: tokens[1],
				})
			}
		}
	}

	return identities, nil
}
