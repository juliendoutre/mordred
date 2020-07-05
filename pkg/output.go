package mordred

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// StdoutJSON prints a JSON representation of a struct to stdout.
func StdoutJSON(v interface{}) error {
	rawBytes, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "error printing identities")
	}

	fmt.Println(string(rawBytes))
	return nil
}

// WriteJSON writes a JSON representation of a struct to a file.
func WriteJSON(path string, v interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "error creating output file")
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(v)
	if err != nil {
		return errors.Wrap(err, "error writing identities to file")
	}

	err = writer.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing writer to output file")
	}

	err = f.Sync()
	if err != nil {
		return errors.Wrap(err, "error syncing outout file")
	}

	return nil
}
