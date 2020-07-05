package cmd

import (
	"fmt"
	"log"

	mordred "github.com/juliendoutre/mordred/pkg"
	"github.com/juliendoutre/mordred/pkg/git"
	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Index a repository",
	Long:  "Scan a repository for its objects and extract useful data from them",
}

var (
	verbose bool
	output  string
)

func init() {
	flags := indexCmd.PersistentFlags()
	flags.BoolVarP(&verbose, "verbose", "v", false, "enable logging")
	flags.StringVarP(&output, "output", "o", "", "file in which to save command outputs, if not set Mordred will prompt the results to sdtout")

	rootCmd.AddCommand(indexCmd)
}

func setup(target string) *git.Repository {
	logIfVerbose("Checking Git install...")
	isGitInstalled, err := git.IsInstalled()
	if err != nil {
		log.Fatal(err)
	}
	logIfVerbose("Git install detected")

	if !isGitInstalled {
		log.Fatal("Please install Git before running this program")
	}

	logIfVerbose(fmt.Sprintf("Retrieving target repository '%s'...", target))
	repository, err := git.NewRepository(target)
	if err != nil {
		log.Fatal(err)
	}
	logIfVerbose(fmt.Sprintf("Repository located at '%s'", repository.Path))

	return repository
}

func saveOrPrint(v interface{}) {
	if output == "" || output == "stdout" {
		if err := mordred.StdoutJSON(v); err != nil {
			log.Fatal(err)
		}
	} else {
		logIfVerbose(fmt.Sprintf("Saving results to %s...", output))
		if err := mordred.WriteJSON(output, v); err != nil {
			log.Fatal(err)
		}
		logIfVerbose(fmt.Sprintf("Results saved to %s", output))
	}
}
