package cmd

import "log"

func logIfVerbose(msg string) {
	if verbose {
		log.Println(msg)
	}
}
