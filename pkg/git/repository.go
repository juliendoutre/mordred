package git

import (
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// Repository stores data about a local repository.
type Repository struct {
	Path string
}

// Clone a remote Git repository.
func Clone(target string) error {
	cmd := exec.Command("git", "clone", target)

	_, err := cmd.Output()
	if err != nil {
		return errors.Wrap(err, "error cloning remote repository")
	}

	return nil
}

// NewLocalRepository returns a Repository from a string local path.
func NewLocalRepository(target string) (*Repository, error) {
	fileInfo, err := os.Stat(target)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("error: target does not exist")
		}

		return nil, errors.Wrap(err, "error checking path")
	}

	if !fileInfo.IsDir() {
		return nil, errors.New("error: target is not a directory")
	}

	gitTarget := path.Join(target, ".git")

	if _, err := os.Stat(gitTarget); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("error: target is not a git repository")
		}

		return nil, errors.Wrap(err, "error checking target is a git repository")
	}

	return &Repository{Path: gitTarget}, nil
}

// NewRepository returns a Repository from a remote endpoint or a local path.
func NewRepository(target string) (*Repository, error) {
	sshRegexp, err := regexp.Compile(`^\w+@\S+:\S+\.git$`)
	if err != nil {
		return nil, errors.Wrap(err, "error compiling ssh regexp")
	}

	httpsRegexp, err := regexp.Compile(`^https:\/\/\S+\/\S+\.git$`)
	if err != nil {
		return nil, errors.Wrap(err, "error compiling https regexp")
	}

	if sshRegexp.MatchString(target) || httpsRegexp.MatchString(target) {
		if err := Clone(target); err != nil {
			return nil, err
		}

		startIndex := strings.LastIndex(target, "/")
		path := target[startIndex+1 : len(target)-4]

		return NewLocalRepository(path)
	}

	return NewLocalRepository(target)
}

// GetObjectType returns the type of an object identified by its hash.
func (r *Repository) GetObjectType(hash string) (string, error) {
	cmd := exec.Command("git", "cat-file", "-t", hash)
	cmd.Dir = r.Path

	tRawOutput, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "error checking object git type")
	}

	return strings.Trim(string(tRawOutput), "\n\t\r "), nil
}

func (r *Repository) listBlobs() ([]string, error) {
	cmd := exec.Command("git", "rev-list", "--objects", "--all")
	cmd.Dir = r.Path

	rawOutput, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "error listing blob objects")
	}

	output := strings.Trim(string(rawOutput), "\n\t\r ")
	return strings.Split(output, "\n"), nil
}

// Blobs returns references to all blobs objects of a repository.
func (r *Repository) Blobs() ([]Blob, error) {
	objects, err := r.listBlobs()
	if err != nil {
		return nil, err
	}

	blobs := make([]Blob, len(objects))
	for _, object := range objects {
		parts := strings.Split(object, " ")

		if len(parts) > 1 {
			objectType, err := r.GetObjectType(parts[0])
			if err != nil {
				return nil, err
			}

			if objectType == BlobType {
				blobs = append(blobs, Blob{hash: parts[0], Name: parts[1]})
			}
		}
	}

	return blobs, nil
}

// ReadObject returns an object contents from its hash.
func (r *Repository) ReadObject(object Object) ([]byte, error) {
	cmd := exec.Command("git", "cat-file", "-p", object.Hash())
	cmd.Dir = r.Path

	rawOutput, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "error reading object contents")
	}

	return rawOutput, nil
}

// ListIdentities returns a strings list of a repository identities.
func (r *Repository) ListIdentities() ([]byte, error) {
	cmd := exec.Command("git", "log", "--pretty=%an,%ae%n%cn,%ce")
	cmd.Dir = r.Path

	rawOutput, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "error listing identities")
	}

	return rawOutput, nil

}
