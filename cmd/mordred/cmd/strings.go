package cmd

import (
	"log"

	mordred "github.com/juliendoutre/mordred/pkg"
	"github.com/juliendoutre/mordred/pkg/parsing"
	"github.com/spf13/cobra"
)

var detailed bool

var indexStringsCmd = &cobra.Command{
	Use:   "strings",
	Short: "Index strings in a repository",
	Long:  "Scan a repository for blob objects and parse their contents to get sensitive values",
	Run: func(cmd *cobra.Command, args []string) {
		indexStrings(args[0])
	},
	Args: cobra.ExactValidArgs(1),
}

func init() {
	flags := indexStringsCmd.Flags()
	flags.BoolVarP(&detailed, "detailed", "d", false, "include blobs references in the index")
	indexCmd.AddCommand(indexStringsCmd)
}

func indexStrings(target string) {
	repository := setup(target)

	if detailed {
		idx, err := mordred.NewTypedIndex(repository)
		if err != nil {
			log.Fatal(err)
		}

		saveOrPrint(struct {
			Version string              `json:"version"`
			Index   *mordred.TypedIndex `json:"strings"`
		}{hash, idx})
	} else {
		idx, err := mordred.NewBasicIndex(repository)
		if err != nil {
			log.Fatal(err)
		}

		saveOrPrint(struct {
			Version string                          `json:"version"`
			Index   map[parsing.StringType][]string `json:"strings"`
		}{hash, idx.Refine()})
	}
}
