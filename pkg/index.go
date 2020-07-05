package mordred

import (
	"github.com/juliendoutre/mordred/pkg/git"
	"github.com/juliendoutre/mordred/pkg/parsing"
	"github.com/pkg/errors"
)

// Index is an inverted index more or less detailed.
type Index interface {
	Update([]string, *git.Blob)
}

// IndexRecord details a string occurence in a specific blob object.
type IndexRecord struct {
	Name string `json:"name,omitempty"`
	Hash string `json:"hash,omitempty"`
}

// InvertedIndex is an inverted index of strings.
type InvertedIndex map[string][]IndexRecord

// BasicIndex contains list of strings sorted following their type.
type BasicIndex map[parsing.StringType]map[string]bool

// Update a BasicIndex with new strings.
func (b *BasicIndex) Update(strings []string, blob *git.Blob) {
	for _, str := range strings {
		strType := parsing.GetType(str)
		if l, ok := (*b)[strType]; ok {
			if _, ok := l[str]; !ok {
				l[str] = true
			}
		} else {
			(*b)[strType] = map[string]bool{str: true}
		}
	}
}

// Refine a BasicIndex to convert its map[string]bool
// (used to avoid duplicated strings at update) in []string
// easier to read for a human.
func (b *BasicIndex) Refine() map[parsing.StringType][]string {
	refinedIdx := map[parsing.StringType][]string{}

	for t, m := range *b {
		l := []string{}
		for s := range m {
			l = append(l, s)
		}

		refinedIdx[t] = l
	}

	return refinedIdx
}

// TypedIndex is an inverted index of strings sorted following their type.
type TypedIndex map[parsing.StringType]InvertedIndex

// Update a TypedIndex with new strings.
func (t *TypedIndex) Update(strings []string, blob *git.Blob) {
	record := IndexRecord{Name: blob.Name, Hash: blob.Hash()}

	for _, str := range strings {
		strType := parsing.GetType(str)
		if i, ok := (*t)[strType]; ok {
			if _, ok := i[str]; ok {
				i[str] = append(i[str], record)
			} else {
				i[str] = []IndexRecord{record}
			}
		} else {
			(*t)[strType] = InvertedIndex{str: []IndexRecord{record}}
		}
	}
}

// NewTypedIndex returns a TypedIndex built over a Repository blobs strings.
func NewTypedIndex(repository *git.Repository) (*TypedIndex, error) {
	blobs, err := repository.Blobs()
	if err != nil {
		return nil, errors.Wrap(err, "error listing repository's blobs")
	}

	idx := &TypedIndex{}

	for _, blob := range blobs {
		strings, err := Parse(&blob, repository)
		if err != nil {
			return nil, err
		}

		idx.Update(strings, &blob)
	}

	return idx, nil
}

// NewBasicIndex returns a BasicIndex built over a Repository blobs strings.
func NewBasicIndex(repository *git.Repository) (*BasicIndex, error) {
	blobs, err := repository.Blobs()
	if err != nil {
		return nil, errors.Wrap(err, "error listing repository's blobs")
	}

	idx := &BasicIndex{}

	for _, blob := range blobs {
		strings, err := Parse(&blob, repository)
		if err != nil {
			return nil, err
		}

		idx.Update(strings, &blob)
	}

	return idx, nil
}
