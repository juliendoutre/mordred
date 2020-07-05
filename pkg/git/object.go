package git

// BlobType and CommitType are the only objects types considered for now.
const (
	BlobType   = "blob"
	CommitType = "commit"
)

// Object is identified by a hash string.
type Object interface {
	Hash() string
}

// Blob references a blob object.
type Blob struct {
	Name string
	hash string
}

// Hash enables Blob to implement Object.
func (b *Blob) Hash() string {
	return b.hash
}

// Commit references a commit object.
type Commit struct {
	hash string
}

// Hash enables Commit to implement Object.
func (c *Commit) Hash() string {
	return c.hash
}
