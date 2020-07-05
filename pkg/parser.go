package mordred

import (
	"fmt"
	"regexp"

	"github.com/juliendoutre/mordred/pkg/git"
	"github.com/juliendoutre/mordred/pkg/parsing"
	"github.com/pkg/errors"
)

// Parser returns a slice of strings captured in a blob contents.
type Parser interface {
	Parse([]byte) ([]string, error)
}

// Parse apply a parsing logic to a blob files and returns the strings
// it could extract from it.
func Parse(blob *git.Blob, repository *git.Repository) ([]string, error) {
	goRegexp, err := regexp.Compile(`^.*\.go$`)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error compiling regex %s", goRegexp))
	}

	if goRegexp.MatchString(blob.Name) {
		contents, err := repository.ReadObject(blob)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error reading contents of blob objects %s", blob.Hash()))
		}

		parser := &parsing.GoParser{}
		strings, err := parser.Parse(contents)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error parsing contents of blob objects %s", blob.Hash()))
		}

		return strings, nil
	}

	return nil, nil
}
