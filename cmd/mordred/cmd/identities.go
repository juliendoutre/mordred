package cmd

import (
	"fmt"
	"log"

	mordred "github.com/juliendoutre/mordred/pkg"
	"github.com/spf13/cobra"
)

var indexIdentitiesCmd = &cobra.Command{
	Use:   "identities",
	Short: "Index identities (names and e-emails addresses) of a repository.",
	Long:  "Parse the commit objects of a repository and extract the names and e-mails addresses they contain.",
	Run: func(cmd *cobra.Command, args []string) {
		indexIndentities(args[0])
	},
	Args: cobra.ExactValidArgs(1),
}

func init() {
	indexCmd.AddCommand(indexIdentitiesCmd)
}

func indexIndentities(target string) {
	repository := setup(target)

	logIfVerbose("Listing repository identities...")
	identities, err := mordred.GetIdentities(repository)
	if err != nil {
		log.Fatal(err)
	}
	logIfVerbose(fmt.Sprintf("Found %d distinct identities", len(identities)))

	saveOrPrint(struct {
		Version    string             `json:"version"`
		Identities []mordred.Identity `json:"identities"`
	}{hash, identities})
}
